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

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/gravitational/teleport-plugins/lib/backoff"
	"github.com/gravitational/teleport-plugins/terraform/tfschema"
	apitypes "github.com/gravitational/teleport/api/types"
	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"
)

// resourceTeleportTrustedClusterType is the resource metadata type
type resourceTeleportTrustedClusterType struct{}

// resourceTeleportTrustedCluster is the resource
type resourceTeleportTrustedCluster struct {
	p Provider
}

// GetSchema returns the resource schema
func (r resourceTeleportTrustedClusterType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfschema.GenSchemaTrustedClusterV2(ctx)
}

// NewResource creates the empty resource
func (r resourceTeleportTrustedClusterType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceTeleportTrustedCluster{
		p: *(p.(*Provider)),
	}, nil
}

// Create creates the provision token
func (r resourceTeleportTrustedCluster) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.IsConfigured(resp.Diagnostics) {
		return
	}

	var plan types.Object
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	trustedCluster := &apitypes.TrustedClusterV2{}
	diags = tfschema.CopyTrustedClusterV2FromTerraform(ctx, plan, trustedCluster)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	

	_, err := r.p.Client.GetTrustedCluster(ctx, trustedCluster.Metadata.Name)
	if !trace.IsNotFound(err) {
		if err == nil {
			n := trustedCluster.Metadata.Name
			existErr := fmt.Sprintf("TrustedCluster exists in Teleport. Either remove it (tctl rm trusted_cluster/%v)"+
				" or import it to the existing state (terraform import teleport_app.%v %v)", n, n, n)

			resp.Diagnostics.Append(diagFromErr("TrustedCluster exists in Teleport", trace.Errorf(existErr)))
			return
		}

		resp.Diagnostics.Append(diagFromWrappedErr("Error reading TrustedCluster", trace.Wrap(err), "trusted_cluster"))
		return
	}

	err = trustedCluster.CheckAndSetDefaults()
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error setting TrustedCluster defaults", trace.Wrap(err), "trusted_cluster"))
		return
	}

	_, err = r.p.Client.UpsertTrustedCluster(ctx, trustedCluster)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error creating TrustedCluster", trace.Wrap(err), "trusted_cluster"))
		return
	}

	id := trustedCluster.Metadata.Name
	var trustedClusterI apitypes.TrustedCluster

	tries := 0
	backoff := backoff.NewDecorr(r.p.RetryConfig.Base, r.p.RetryConfig.Cap, clockwork.NewRealClock())
	for {
		tries = tries + 1
		trustedClusterI, err = r.p.Client.GetTrustedCluster(ctx, id)
		if trace.IsNotFound(err) {
			if bErr := backoff.Do(ctx); bErr != nil {
				resp.Diagnostics.Append(diagFromWrappedErr("Error reading TrustedCluster", trace.Wrap(err), "trusted_cluster"))
				return
			}
			if tries >= r.p.RetryConfig.MaxTries {
				diagMessage := fmt.Sprintf("Error reading TrustedCluster (tried %d times)", tries)
				resp.Diagnostics.Append(diagFromWrappedErr(diagMessage, trace.Wrap(err), "trusted_cluster"))
				return
			}
			continue
		}
		break
	}

	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading TrustedCluster", trace.Wrap(err), "trusted_cluster"))
		return
	}

	trustedCluster, ok := trustedClusterI.(*apitypes.TrustedClusterV2)
	if !ok {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading TrustedCluster", trace.Errorf("Can not convert %T to TrustedClusterV2", trustedClusterI), "trusted_cluster"))
		return
	}

	diags = tfschema.CopyTrustedClusterV2ToTerraform(ctx, *trustedCluster, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Attrs["id"] = types.String{Value: trustedCluster.Metadata.Name}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read reads teleport TrustedCluster
func (r resourceTeleportTrustedCluster) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state types.Object
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var id types.String
	diags = req.State.GetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("metadata").WithAttributeName("name"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	trustedClusterI, err := r.p.Client.GetTrustedCluster(ctx, id.Value)
	if trace.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading TrustedCluster", trace.Wrap(err), "trusted_cluster"))
		return
	}

	trustedCluster := trustedClusterI.(*apitypes.TrustedClusterV2)
	diags = tfschema.CopyTrustedClusterV2ToTerraform(ctx, *trustedCluster, &state)
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

// Update updates teleport TrustedCluster
func (r resourceTeleportTrustedCluster) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	if !r.p.IsConfigured(resp.Diagnostics) {
		return
	}

	var plan types.Object
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	trustedCluster := &apitypes.TrustedClusterV2{}
	diags = tfschema.CopyTrustedClusterV2FromTerraform(ctx, plan, trustedCluster)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := trustedCluster.Metadata.Name

	err := trustedCluster.CheckAndSetDefaults()
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error updating TrustedCluster", err, "trusted_cluster"))
		return
	}

	_, err = r.p.Client.UpsertTrustedCluster(ctx, trustedCluster)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error updating TrustedCluster", err, "trusted_cluster"))
		return
	}

	trustedClusterI, err := r.p.Client.GetTrustedCluster(ctx, name)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading TrustedCluster", err, "trusted_cluster"))
		return
	}

	trustedCluster = trustedClusterI.(*apitypes.TrustedClusterV2)
	diags = tfschema.CopyTrustedClusterV2ToTerraform(ctx, *trustedCluster, &plan)
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

// Delete deletes Teleport TrustedCluster
func (r resourceTeleportTrustedCluster) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var id types.String
	diags := req.State.GetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("metadata").WithAttributeName("name"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.p.Client.DeleteTrustedCluster(ctx, id.Value)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error deleting TrustedClusterV2", trace.Wrap(err), "trusted_cluster"))
		return
	}

	resp.State.RemoveResource(ctx)
}

// ImportState imports TrustedCluster state
func (r resourceTeleportTrustedCluster) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	trustedClusterI, err := r.p.Client.GetTrustedCluster(ctx, req.ID)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading TrustedCluster", trace.Wrap(err), "trusted_cluster"))
		return
	}

	trustedCluster := trustedClusterI.(*apitypes.TrustedClusterV2)

	var state types.Object

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = tfschema.CopyTrustedClusterV2ToTerraform(ctx, *trustedCluster, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Attrs["id"] = types.String{Value: trustedCluster.Metadata.Name}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
