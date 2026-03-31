package provider

import (
	"context"
	"fmt"

	embyclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &UserResource{}
	_ resource.ResourceWithImportState = &UserResource{}
)

type UserResource struct {
	data *EmbyProviderData
}

type UserResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	IsAdministrator       types.Bool   `tfsdk:"is_administrator"`
	IsHidden              types.Bool   `tfsdk:"is_hidden"`
	IsDisabled            types.Bool   `tfsdk:"is_disabled"`
	EnableRemoteAccess    types.Bool   `tfsdk:"enable_remote_access"`
	EnableMediaPlayback   types.Bool   `tfsdk:"enable_media_playback"`
	EnableAllFolders      types.Bool   `tfsdk:"enable_all_folders"`
	EnableAllChannels     types.Bool   `tfsdk:"enable_all_channels"`
	EnableAllDevices      types.Bool   `tfsdk:"enable_all_devices"`
	EnableContentDeletion types.Bool   `tfsdk:"enable_content_deletion"`
}

func NewUserResource() resource.Resource {
	return &UserResource{}
}

func (r *UserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *UserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Emby user.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The user ID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The username.",
			},
			"is_administrator": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Whether the user is an administrator.",
			},
			"is_hidden": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Whether the user is hidden from login screens.",
			},
			"is_disabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Whether the user account is disabled.",
			},
			"enable_remote_access": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Whether the user can access remotely.",
			},
			"enable_media_playback": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Whether the user can play media.",
			},
			"enable_all_folders": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Whether the user can access all library folders.",
			},
			"enable_all_channels": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Whether the user can access all channels.",
			},
			"enable_all_devices": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Whether the user can use all devices.",
			},
			"enable_content_deletion": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Whether the user can delete content.",
			},
		},
	}
}

func (r *UserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *UserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := embyclient.NewCreateUserByName()
	createReq.SetName(plan.Name.ValueString())

	user, _, err := r.data.Client.UserServiceAPI.PostUsersNew(r.data.Auth).CreateUserByName(*createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to create user", err.Error())
		return
	}

	plan.Id = types.StringValue(user.GetId())

	// Update the user policy with desired settings
	r.updateUserPolicy(plan, user, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update the user name if needed (PostUsersById updates user details)
	r.updateUser(plan, user, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	user, httpResp, err := r.data.Client.UserServiceAPI.GetUsersById(r.data.Auth, state.Id.ValueString()).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Unable to read user", err.Error())
		return
	}

	mapUserToState(user, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan UserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = state.Id

	// Read current user to get full object for update
	user, _, err := r.data.Client.UserServiceAPI.GetUsersById(r.data.Auth, plan.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to read user for update", err.Error())
		return
	}

	r.updateUserPolicy(plan, user, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	r.updateUser(plan, user, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.data.Client.UserServiceAPI.DeleteUsersById(r.data.Auth, state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to delete user", err.Error())
		return
	}
}

func (r *UserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// updateUserPolicy applies the planned policy settings to the Emby user via PostUsersByIdPolicy.
func (r *UserResource) updateUserPolicy(plan UserResourceModel, user *embyclient.UserDto, diags *diag.Diagnostics) {
	policy := embyclient.UserPolicy{}
	// Preserve existing policy fields if available
	if existing, ok := user.GetPolicyOk(); ok && existing != nil {
		policy = *existing
	}
	policy.SetIsAdministrator(plan.IsAdministrator.ValueBool())
	policy.SetIsHidden(plan.IsHidden.ValueBool())
	policy.SetIsDisabled(plan.IsDisabled.ValueBool())
	policy.SetEnableRemoteAccess(plan.EnableRemoteAccess.ValueBool())
	policy.SetEnableMediaPlayback(plan.EnableMediaPlayback.ValueBool())
	policy.SetEnableAllFolders(plan.EnableAllFolders.ValueBool())
	policy.SetEnableAllChannels(plan.EnableAllChannels.ValueBool())
	policy.SetEnableAllDevices(plan.EnableAllDevices.ValueBool())
	policy.SetEnableContentDeletion(plan.EnableContentDeletion.ValueBool())

	_, err := r.data.Client.UserServiceAPI.PostUsersByIdPolicy(r.data.Auth, user.GetId()).UserPolicy(policy).Execute()
	if err != nil {
		diags.AddError("Unable to update user policy", err.Error())
	}
}

// updateUser applies name changes to the user via PostUsersById.
func (r *UserResource) updateUser(plan UserResourceModel, user *embyclient.UserDto, diags *diag.Diagnostics) {
	user.SetName(plan.Name.ValueString())

	_, err := r.data.Client.UserServiceAPI.PostUsersById(r.data.Auth, user.GetId()).UserDto(*user).Execute()
	if err != nil {
		diags.AddError("Unable to update user", err.Error())
	}
}

func mapUserToState(user *embyclient.UserDto, state *UserResourceModel) {
	state.Id = types.StringValue(user.GetId())
	state.Name = types.StringValue(user.GetName())

	policy := user.GetPolicy()
	state.IsAdministrator = types.BoolValue(policy.GetIsAdministrator())
	state.IsHidden = types.BoolValue(policy.GetIsHidden())
	state.IsDisabled = types.BoolValue(policy.GetIsDisabled())
	state.EnableRemoteAccess = types.BoolValue(policy.GetEnableRemoteAccess())
	state.EnableMediaPlayback = types.BoolValue(policy.GetEnableMediaPlayback())
	state.EnableAllFolders = types.BoolValue(policy.GetEnableAllFolders())
	state.EnableAllChannels = types.BoolValue(policy.GetEnableAllChannels())
	state.EnableAllDevices = types.BoolValue(policy.GetEnableAllDevices())
	state.EnableContentDeletion = types.BoolValue(policy.GetEnableContentDeletion())
}
