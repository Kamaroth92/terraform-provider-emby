// Copyright (c) HashiCorp, Inc.

package provider

import (
	"context"
	"fmt"

	embyclient "github.com/Kamaroth92/terraform-provider-emby/client"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &UserLibraryAccessResource{}

type UserLibraryAccessResource struct {
	data *EmbyProviderData
}

type UserLibraryAccessResourceModel struct {
	UserId     types.String `tfsdk:"user_id"`
	LibraryIds types.Set    `tfsdk:"library_ids"`
}

func NewUserLibraryAccessResource() resource.Resource {
	return &UserLibraryAccessResource{}
}

func (r *UserLibraryAccessResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_library_access"
}

func (r *UserLibraryAccessResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages which libraries a user can access. Sets `EnableAllFolders` to false and assigns the specified library IDs. Destroying this resource restores full library access for the user.",
		Attributes: map[string]schema.Attribute{
			"user_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The ID of the user to manage library access for.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"library_ids": schema.SetAttribute{
				Required:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of library IDs the user is allowed to access.",
			},
		},
	}
}

func (r *UserLibraryAccessResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(*EmbyProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *EmbyProviderData, got: %T.", req.ProviderData),
		)
		return
	}

	r.data = data
}

func (r *UserLibraryAccessResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserLibraryAccessResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.applyLibraryAccess(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserLibraryAccessResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserLibraryAccessResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	user, httpResp, err := r.data.Client.UserServiceAPI.GetUsersById(r.data.Auth, state.UserId.ValueString()).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Unable to read user", err.Error())
		return
	}

	policy := user.GetPolicy()

	// If EnableAllFolders was set back to true outside Terraform, remove this resource from state
	if policy.GetEnableAllFolders() {
		resp.State.RemoveResource(ctx)
		return
	}

	enabledFolders := policy.GetEnabledFolders()
	folderValues := make([]types.String, len(enabledFolders))
	for i, id := range enabledFolders {
		folderValues[i] = types.StringValue(id)
	}

	libraryIds, diags := types.SetValueFrom(ctx, types.StringType, enabledFolders)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.LibraryIds = libraryIds
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserLibraryAccessResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan UserLibraryAccessResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.applyLibraryAccess(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserLibraryAccessResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UserLibraryAccessResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Restore full library access
	user, httpResp, err := r.data.Client.UserServiceAPI.GetUsersById(r.data.Auth, state.UserId.ValueString()).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			return // user already gone, nothing to restore
		}
		resp.Diagnostics.AddError("Unable to read user", err.Error())
		return
	}

	policy := embyclient.UserPolicy{}
	if existing, ok := user.GetPolicyOk(); ok && existing != nil {
		policy = *existing
	}
	policy.SetEnableAllFolders(true)
	policy.SetEnabledFolders([]string{})

	_, err = r.data.Client.UserServiceAPI.PostUsersByIdPolicy(r.data.Auth, user.GetId()).UserPolicy(policy).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to restore user library access", err.Error())
		return
	}
}

// applyLibraryAccess sets EnableAllFolders=false and updates EnabledFolders on the user policy.
func (r *UserLibraryAccessResource) applyLibraryAccess(ctx context.Context, plan UserLibraryAccessResourceModel, diags *diag.Diagnostics) {
	user, _, err := r.data.Client.UserServiceAPI.GetUsersById(r.data.Auth, plan.UserId.ValueString()).Execute()
	if err != nil {
		diags.AddError("Unable to read user", err.Error())
		return
	}

	var libraryIds []string
	diags.Append(plan.LibraryIds.ElementsAs(ctx, &libraryIds, false)...)
	if diags.HasError() {
		return
	}

	policy := embyclient.UserPolicy{}
	if existing, ok := user.GetPolicyOk(); ok && existing != nil {
		policy = *existing
	}
	policy.SetEnableAllFolders(false)
	policy.SetEnabledFolders(libraryIds)

	_, err = r.data.Client.UserServiceAPI.PostUsersByIdPolicy(r.data.Auth, user.GetId()).UserPolicy(policy).Execute()
	if err != nil {
		diags.AddError("Unable to update user library access", err.Error())
	}
}
