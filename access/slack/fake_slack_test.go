package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gravitational/trace"
	"github.com/julienschmidt/httprouter"

	log "github.com/sirupsen/logrus"
)

type FakeSlack struct {
	srv *httptest.Server

	botUser                    SlackUser
	objects                    sync.Map
	newMessages                chan SlackMsg
	messageUpdatesByAPI        chan SlackMsg
	messageUpdatesByResponding chan SlackMsg
	messageCounter             uint64
	userIDCounter              uint64
	startTime                  time.Time
}

func NewFakeSlack(botUser SlackUser, concurrency int) *FakeSlack {
	router := httprouter.New()

	s := &FakeSlack{
		newMessages:                make(chan SlackMsg, concurrency*6),
		messageUpdatesByAPI:        make(chan SlackMsg, concurrency*2),
		messageUpdatesByResponding: make(chan SlackMsg, concurrency),
		startTime:                  time.Now(),
		srv:                        httptest.NewServer(router),
	}

	s.botUser = s.StoreUser(botUser)

	router.POST("/auth.test", func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rw.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(rw).Encode(SlackResponse{Ok: true})
		panicIf(err)
	})

	router.POST("/chat.postMessage", func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rw.Header().Add("Content-Type", "application/json")

		var payload SlackMsg
		err := json.NewDecoder(r.Body).Decode(&payload)
		panicIf(err)

		// text limit and block text limit as per
		// https://api.slack.com/methods/chat.postMessage and
		// https://api.slack.com/reference/block-kit/blocks#section
		if len(payload.Text) > 4000 || func() bool {
			for _, block := range payload.BlockItems {
				sectionBlock, ok := block.Block.(SectionBlock)
				if !ok {
					continue
				}
				if len(sectionBlock.Text.GetText()) > 3000 {
					return true
				}
			}
			return false
		}() {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		msg := s.StoreMessage(SlackMsg{Msg: Msg{
			Type:     "message",
			Channel:  payload.Channel,
			ThreadTs: payload.ThreadTs,
			User:     s.botUser.ID,
			Username: s.botUser.Name,
		},
			BlockItems: payload.BlockItems,
			Text:       payload.Text,
		})
		s.newMessages <- msg

		response := ChatMsgResponse{
			SlackResponse: SlackResponse{Ok: true},
			Channel:       msg.Channel,
			Timestamp:     msg.Timestamp,
			Text:          msg.Text,
		}
		err = json.NewEncoder(rw).Encode(response)
		panicIf(err)
	})

	router.POST("/chat.update", func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rw.Header().Add("Content-Type", "application/json")

		var payload SlackMsg
		err := json.NewDecoder(r.Body).Decode(&payload)
		panicIf(err)

		msg, found := s.GetMessage(payload.Timestamp)
		if !found {
			err := json.NewEncoder(rw).Encode(SlackResponse{Ok: false, Error: "message_not_found"})
			panicIf(err)
			return
		}

		msg.Text = payload.Text
		msg.BlockItems = payload.BlockItems

		s.messageUpdatesByAPI <- s.StoreMessage(msg)

		response := ChatMsgResponse{
			SlackResponse: SlackResponse{Ok: true},
			Channel:       msg.Channel,
			Timestamp:     msg.Timestamp,
			Text:          msg.Text,
		}
		err = json.NewEncoder(rw).Encode(&response)
		panicIf(err)
	})

	router.POST("/_response/:ts", func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		rw.Header().Add("Content-Type", "application/json")

		var payload struct {
			SlackMsg
			ReplaceOriginal bool `json:"replace_original"`
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		panicIf(err)

		timestamp := ps.ByName("ts")
		msg, found := s.GetMessage(timestamp)
		if !found {
			err := json.NewEncoder(rw).Encode(SlackResponse{Ok: false, Error: "message_not_found"})
			panicIf(err)
			return
		}

		if payload.ReplaceOriginal {
			msg.BlockItems = payload.BlockItems
			s.messageUpdatesByResponding <- s.StoreMessage(msg)
		} else {
			newMsg := s.StoreMessage(SlackMsg{Msg: Msg{
				Type:     "message",
				Channel:  msg.Channel,
				User:     s.botUser.ID,
				Username: s.botUser.Name,
			},
				BlockItems: payload.BlockItems,
			})
			s.newMessages <- newMsg
		}
		err = json.NewEncoder(rw).Encode(SlackResponse{Ok: true})
		panicIf(err)
	})

	router.GET("/users.info", func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rw.Header().Add("Content-Type", "application/json")

		id := r.URL.Query().Get("user")
		if id == "" {
			err := json.NewEncoder(rw).Encode(SlackResponse{Ok: false, Error: "invalid_arguments"})
			panicIf(err)
			return
		}

		user, found := s.GetUser(id)
		if !found {
			err := json.NewEncoder(rw).Encode(SlackResponse{Ok: false, Error: "user_not_found"})
			panicIf(err)
			return
		}

		err := json.NewEncoder(rw).Encode(struct {
			User SlackUser `json:"user"`
			Ok   bool      `json:"ok"`
		}{user, true})
		panicIf(err)
	})

	router.GET("/users.lookupByEmail", func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rw.Header().Add("Content-Type", "application/json")

		email := r.URL.Query().Get("email")
		if email == "" {
			err := json.NewEncoder(rw).Encode(SlackResponse{Ok: false, Error: "invalid_arguments"})
			panicIf(err)
			return
		}

		user, found := s.GetUserByEmail(email)
		if !found {
			err := json.NewEncoder(rw).Encode(SlackResponse{Ok: false, Error: "users_not_found"})
			panicIf(err)
			return
		}

		err := json.NewEncoder(rw).Encode(struct {
			User SlackUser `json:"user"`
			Ok   bool      `json:"ok"`
		}{user, true})
		panicIf(err)
	})

	return s
}

func (s *FakeSlack) URL() string {
	return s.srv.URL
}

func (s *FakeSlack) Close() {
	s.srv.Close()
	close(s.newMessages)
	close(s.messageUpdatesByAPI)
	close(s.messageUpdatesByResponding)
}

func (s *FakeSlack) StoreMessage(msg SlackMsg) SlackMsg {
	if msg.Timestamp == "" {
		now := s.startTime.Add(time.Since(s.startTime)) // get monotonic timestamp
		uniq := atomic.AddUint64(&s.messageCounter, 1)  // generate uniq int to prevent races
		msg.Timestamp = fmt.Sprintf("%d.%d", now.UnixNano(), uniq)
	}
	s.objects.Store(fmt.Sprintf("msg-%s", msg.Timestamp), msg)
	return msg
}

func (s *FakeSlack) GetMessage(id string) (SlackMsg, bool) {
	if obj, ok := s.objects.Load(fmt.Sprintf("msg-%s", id)); ok {
		msg, ok := obj.(SlackMsg)
		return msg, ok
	}
	return SlackMsg{}, false
}

func (s *FakeSlack) StoreUser(user SlackUser) SlackUser {
	if user.ID == "" {
		user.ID = fmt.Sprintf("U%d", atomic.AddUint64(&s.userIDCounter, 1))
	}
	s.objects.Store(fmt.Sprintf("user-%s", user.ID), user)
	s.objects.Store(fmt.Sprintf("userByEmail-%s", user.Profile.Email), user)
	return user
}

func (s *FakeSlack) GetUser(id string) (SlackUser, bool) {
	if obj, ok := s.objects.Load(fmt.Sprintf("user-%s", id)); ok {
		user, ok := obj.(SlackUser)
		return user, ok
	}
	return SlackUser{}, false
}

func (s *FakeSlack) GetUserByEmail(email string) (SlackUser, bool) {
	if obj, ok := s.objects.Load(fmt.Sprintf("userByEmail-%s", email)); ok {
		user, ok := obj.(SlackUser)
		return user, ok
	}
	return SlackUser{}, false
}

func (s *FakeSlack) CheckNewMessage(ctx context.Context) (SlackMsg, error) {
	select {
	case message := <-s.newMessages:
		return message, nil
	case <-ctx.Done():
		return SlackMsg{}, trace.Wrap(ctx.Err())
	}
}

func (s *FakeSlack) CheckMessageUpdateByAPI(ctx context.Context) (SlackMsg, error) {
	select {
	case message := <-s.messageUpdatesByAPI:
		return message, nil
	case <-ctx.Done():
		return SlackMsg{}, trace.Wrap(ctx.Err())
	}
}

func (s *FakeSlack) CheckMessageUpdateByResponding(ctx context.Context) (SlackMsg, error) {
	select {
	case message := <-s.messageUpdatesByResponding:
		return message, nil
	case <-ctx.Done():
		return SlackMsg{}, trace.Wrap(ctx.Err())
	}
}

func panicIf(err error) {
	if err != nil {
		log.Panicf("%v at %v", err, string(debug.Stack()))
	}
}
