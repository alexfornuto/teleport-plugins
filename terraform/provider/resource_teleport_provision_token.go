// Code generated by _gen/main.go DO NOT EDIT
/*
Copyright 2015-2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"
	"fmt"
	"crypto/rand"
	"encoding/hex"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/path"

	"github.com/gravitational/teleport-plugins/lib/backoff"
	"github.com/gravitational/teleport-plugins/terraform/tfschema"
	apitypes "github.com/gravitational/teleport/api/types"
	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"
)

// resourceTeleportProvisionTokenType is the resource metadata type
type resourceTeleportProvisionTokenType struct{}

// resourceTeleportProvisionToken is the resource
type resourceTeleportProvisionToken struct {
	p Provider
}

// GetSchema returns the resource schema
func (r resourceTeleportProvisionTokenType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfschema.GenSchemaProvisionTokenV2(ctx)
}

// NewResource creates the empty resource
func (r resourceTeleportProvisionTokenType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceTeleportProvisionToken{
		p: *(p.(*Provider)),
	}, nil
}

// Create creates the provision token
func (r resourceTeleportProvisionToken) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.IsConfigured(resp.Diagnostics) {
		return
	}

	var plan types.Object
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	provisionToken := &apitypes.ProvisionTokenV2{}
	diags = tfschema.CopyProvisionTokenV2FromTerraform(ctx, plan, provisionToken)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if provisionToken.Metadata.Name == "" {
		b := make([]byte, 32)
		_, err := rand.Read(b)
		if err != nil {
			resp.Diagnostics.AddError("Failed to generate random token", err.Error())
			return
		}
		provisionToken.Metadata.Name = hex.EncodeToString(b)
	}
	

	_, err := r.p.Client.GetToken(ctx, provisionToken.Metadata.Name)
	if !trace.IsNotFound(err) {
		if err == nil {
			n := provisionToken.Metadata.Name
			existErr := fmt.Sprintf("ProvisionToken exists in Teleport. Either remove it (tctl rm token/%v)"+
				" or import it to the existing state (terraform import teleport_app.%v %v)", n, n, n)

			resp.Diagnostics.Append(diagFromErr("ProvisionToken exists in Teleport", trace.Errorf(existErr)))
			return
		}

		resp.Diagnostics.Append(diagFromWrappedErr("Error reading ProvisionToken", trace.Wrap(err), "token"))
		return
	}

	err = provisionToken.CheckAndSetDefaults()
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error setting ProvisionToken defaults", trace.Wrap(err), "token"))
		return
	}

	err = r.p.Client.UpsertToken(ctx, provisionToken)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error creating ProvisionToken", trace.Wrap(err), "token"))
		return
	}

	id := provisionToken.Metadata.Name
	var provisionTokenI apitypes.ProvisionToken

	tries := 0
	backoff := backoff.NewDecorr(r.p.RetryConfig.Base, r.p.RetryConfig.Cap, clockwork.NewRealClock())
	for {
		tries = tries + 1
		provisionTokenI, err = r.p.Client.GetToken(ctx, id)
		if trace.IsNotFound(err) {
			if bErr := backoff.Do(ctx); bErr != nil {
				resp.Diagnostics.Append(diagFromWrappedErr("Error reading ProvisionToken", trace.Wrap(err), "token"))
				return
			}
			if tries >= r.p.RetryConfig.MaxTries {
				diagMessage := fmt.Sprintf("Error reading ProvisionToken (tried %d times) - state outdated, please import resource", tries)
				resp.Diagnostics.Append(diagFromWrappedErr(diagMessage, trace.Wrap(err), "token"))
				return
			}
			continue
		}
		break
	}

	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading ProvisionToken", trace.Wrap(err), "token"))
		return
	}

	provisionToken, ok := provisionTokenI.(*apitypes.ProvisionTokenV2)
	if !ok {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading ProvisionToken", trace.Errorf("Can not convert %T to ProvisionTokenV2", provisionTokenI), "token"))
		return
	}

	diags = tfschema.CopyProvisionTokenV2ToTerraform(ctx, *provisionToken, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Attrs["id"] = types.String{Value: provisionToken.Metadata.Name}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read reads teleport ProvisionToken
func (r resourceTeleportProvisionToken) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state types.Object
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var id types.String
	diags = req.State.GetAttribute(ctx, path.Root("metadata").AtName("name"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	provisionTokenI, err := r.p.Client.GetToken(ctx, id.Value)
	if trace.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading ProvisionToken", trace.Wrap(err), "token"))
		return
	}

	provisionToken := provisionTokenI.(*apitypes.ProvisionTokenV2)
	diags = tfschema.CopyProvisionTokenV2ToTerraform(ctx, *provisionToken, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates teleport ProvisionToken
func (r resourceTeleportProvisionToken) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	if !r.p.IsConfigured(resp.Diagnostics) {
		return
	}

	var plan types.Object
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	provisionToken := &apitypes.ProvisionTokenV2{}
	diags = tfschema.CopyProvisionTokenV2FromTerraform(ctx, plan, provisionToken)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := provisionToken.Metadata.Name

	err := provisionToken.CheckAndSetDefaults()
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error updating ProvisionToken", err, "token"))
		return
	}

	provisionTokenBefore, err := r.p.Client.GetToken(ctx, name)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading ProvisionToken", err, "token"))
		return
	}

	err = r.p.Client.UpsertToken(ctx, provisionToken)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error updating ProvisionToken", err, "token"))
		return
	}

	var provisionTokenI apitypes.ProvisionToken

	tries := 0
	backoff := backoff.NewDecorr(r.p.RetryConfig.Base, r.p.RetryConfig.Cap, clockwork.NewRealClock())
	for {
		tries = tries + 1
		provisionTokenI, err = r.p.Client.GetToken(ctx, name)
		if err != nil {
			resp.Diagnostics.Append(diagFromWrappedErr("Error reading ProvisionToken", err, "token"))
			return
		}
		if provisionTokenBefore.GetMetadata().ID != provisionTokenI.GetMetadata().ID || false {
			break
		}

		if err := backoff.Do(ctx); err != nil {
			resp.Diagnostics.Append(diagFromWrappedErr("Error reading ProvisionToken", trace.Wrap(err), "token"))
			return
		}
		if tries >= r.p.RetryConfig.MaxTries {
			diagMessage := fmt.Sprintf("Error reading ProvisionToken (tried %d times) - state outdated, please import resource", tries)
			resp.Diagnostics.AddError(diagMessage, "token")
			return
		}
	}

	provisionToken = provisionTokenI.(*apitypes.ProvisionTokenV2)
	diags = tfschema.CopyProvisionTokenV2ToTerraform(ctx, *provisionToken, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes Teleport ProvisionToken
func (r resourceTeleportProvisionToken) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var id types.String
	diags := req.State.GetAttribute(ctx, path.Root("metadata").AtName("name"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.p.Client.DeleteToken(ctx, id.Value)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error deleting ProvisionTokenV2", trace.Wrap(err), "token"))
		return
	}

	resp.State.RemoveResource(ctx)
}

// ImportState imports ProvisionToken state
func (r resourceTeleportProvisionToken) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	provisionTokenI, err := r.p.Client.GetToken(ctx, req.ID)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading ProvisionToken", trace.Wrap(err), "token"))
		return
	}

	provisionToken := provisionTokenI.(*apitypes.ProvisionTokenV2)

	var state types.Object

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = tfschema.CopyProvisionTokenV2ToTerraform(ctx, *provisionToken, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Attrs["id"] = types.String{Value: provisionToken.Metadata.Name}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
