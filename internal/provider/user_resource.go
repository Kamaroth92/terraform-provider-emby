// Copyright (c) HashiCorp, Inc.

package provider

import (
	"context"
	"fmt"
	"sync"

	embyclient "github.com/Kamaroth92/terraform-provider-emby/client"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// userDeleteMutex serializes user deletions to avoid Emby server race conditions.
var userDeleteMutex sync.Mutex

var (
	_ resource.Resource                = &UserResource{}
	_ resource.ResourceWithImportState = &UserResource{}
)

type UserResource struct {
	data *EmbyProviderData
}

type UserResourceModel struct {
	Id                             types.String `tfsdk:"id"`
	Name                           types.String `tfsdk:"name"`
	IsAdministrator                types.Bool   `tfsdk:"is_administrator"`
	IsHidden                       types.Bool   `tfsdk:"is_hidden"`
	IsHiddenRemotely               types.Bool   `tfsdk:"is_hidden_remotely"`
	IsHiddenFromUnusedDevices      types.Bool   `tfsdk:"is_hidden_from_unused_devices"`
	IsDisabled                     types.Bool   `tfsdk:"is_disabled"`
	EnableRemoteAccess             types.Bool   `tfsdk:"enable_remote_access"`
	EnableMediaPlayback            types.Bool   `tfsdk:"enable_media_playback"`
	EnableAllFolders               types.Bool   `tfsdk:"enable_all_folders"`
	EnableAllChannels              types.Bool   `tfsdk:"enable_all_channels"`
	EnableAllDevices               types.Bool   `tfsdk:"enable_all_devices"`
	EnableContentDeletion          types.Bool   `tfsdk:"enable_content_deletion"`
	AuthenticationProviderId       types.String `tfsdk:"authentication_provider_id"`
	EnabledFolders                 types.Set    `tfsdk:"enabled_folders"`
	ExcludedSubFolders             types.Set    `tfsdk:"excluded_sub_folders"`
	EnabledChannels                types.Set    `tfsdk:"enabled_channels"`
	EnabledDevices                 types.Set    `tfsdk:"enabled_devices"`
	BlockedTags                    types.Set    `tfsdk:"blocked_tags"`
	IncludeTags                    types.Set    `tfsdk:"include_tags"`
	IsTagBlockingModeInclusive     types.Bool   `tfsdk:"is_tag_blocking_mode_inclusive"`
	AllowTagOrRating               types.Bool   `tfsdk:"allow_tag_or_rating"`
	MaxParentalRating              types.Int64  `tfsdk:"max_parental_rating"`
	BlockUnratedItems              types.Set    `tfsdk:"block_unrated_items"`
	EnableLiveTvAccess             types.Bool   `tfsdk:"enable_live_tv_access"`
	EnableLiveTvManagement         types.Bool   `tfsdk:"enable_live_tv_management"`
	EnableAudioPlaybackTranscoding types.Bool   `tfsdk:"enable_audio_playback_transcoding"`
	EnableVideoPlaybackTranscoding types.Bool   `tfsdk:"enable_video_playback_transcoding"`
	EnablePlaybackRemuxing         types.Bool   `tfsdk:"enable_playback_remuxing"`
	EnableContentDownloading       types.Bool   `tfsdk:"enable_content_downloading"`
	EnableSubtitleDownloading      types.Bool   `tfsdk:"enable_subtitle_downloading"`
	EnableSubtitleManagement       types.Bool   `tfsdk:"enable_subtitle_management"`
	EnableSyncTranscoding          types.Bool   `tfsdk:"enable_sync_transcoding"`
	EnableMediaConversion          types.Bool   `tfsdk:"enable_media_conversion"`
	EnablePublicSharing            types.Bool   `tfsdk:"enable_public_sharing"`
	RemoteClientBitrateLimit       types.Int64  `tfsdk:"remote_client_bitrate_limit"`
	SimultaneousStreamLimit        types.Int64  `tfsdk:"simultaneous_stream_limit"`
	AllowCameraUpload              types.Bool   `tfsdk:"allow_camera_upload"`
	AllowSharingPersonalItems      types.Bool   `tfsdk:"allow_sharing_personal_items"`
	AudioLanguagePreference        types.String `tfsdk:"audio_language_preference"`
	PlayDefaultAudioTrack          types.Bool   `tfsdk:"play_default_audio_track"`
	SubtitleLanguagePreference     types.String `tfsdk:"subtitle_language_preference"`
	SubtitleMode                   types.String `tfsdk:"subtitle_mode"`
	DisplayMissingEpisodes         types.Bool   `tfsdk:"display_missing_episodes"`
	OrderedViews                   types.List   `tfsdk:"ordered_views"`
	LatestItemsExcludes            types.Set    `tfsdk:"latest_items_excludes"`
	MyMediaExcludes                types.Set    `tfsdk:"my_media_excludes"`
	HidePlayedInLatest             types.Bool   `tfsdk:"hide_played_in_latest"`
	HidePlayedInMoreLikeThis       types.Bool   `tfsdk:"hide_played_in_more_like_this"`
	HidePlayedInSuggestions        types.Bool   `tfsdk:"hide_played_in_suggestions"`
	RememberAudioSelections        types.Bool   `tfsdk:"remember_audio_selections"`
	RememberSubtitleSelections     types.Bool   `tfsdk:"remember_subtitle_selections"`
	EnableNextEpisodeAutoPlay      types.Bool   `tfsdk:"enable_next_episode_auto_play"`
	ResumeRewindSeconds            types.Int64  `tfsdk:"resume_rewind_seconds"`
	IntroSkipMode                  types.String `tfsdk:"intro_skip_mode"`
}

func NewUserResource() resource.Resource {
	return &UserResource{}
}

func (r *UserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *UserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Emby user, including policy, library/channel/device access, parental controls, and profile configuration.",
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
			// --- Policy: general ---
			"is_administrator": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user is an administrator.",
			},
			"is_hidden": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user is hidden from login screens.",
			},
			"is_hidden_remotely": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user is hidden from remote connections.",
			},
			"is_hidden_from_unused_devices": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user is hidden from unused devices.",
			},
			"is_disabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user account is disabled.",
			},
			"enable_remote_access": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can access remotely.",
			},
			"enable_media_playback": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can play media.",
			},
			"authentication_provider_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The authentication provider ID for this user (e.g., an LDAP provider ID). Defaults to the server's default provider.",
			},
			// --- Policy: library access ---
			"enable_all_folders": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can access all library folders. Set to false to restrict to `enabled_folders`.",
			},
			"enabled_folders": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of library IDs the user is allowed to access. Only effective when `enable_all_folders` is false.",
			},
			"excluded_sub_folders": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of sub-folder paths the user is denied access to.",
			},
			// --- Policy: channel access ---
			"enable_all_channels": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can access all channels. Set to false to restrict to `enabled_channels`.",
			},
			"enabled_channels": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of channel IDs the user is allowed to access. Only effective when `enable_all_channels` is false.",
			},
			// --- Policy: device access ---
			"enable_all_devices": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can use all devices. Set to false to restrict to `enabled_devices`.",
			},
			"enabled_devices": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of device IDs the user is allowed to use. Only effective when `enable_all_devices` is false.",
			},
			// --- Policy: parental controls ---
			"blocked_tags": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of content tags to block or allow, depending on `is_tag_blocking_mode_inclusive`.",
			},
			"include_tags": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of content tags to include (only used when `is_tag_blocking_mode_inclusive` is true).",
			},
			"is_tag_blocking_mode_inclusive": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "If false, `blocked_tags` are excluded. If true, only content matching `blocked_tags` or `include_tags` is allowed.",
			},
			"allow_tag_or_rating": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether a single tag OR rating match is sufficient to restrict content.",
			},
			"max_parental_rating": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Maximum allowed parental rating. Omit to leave unchanged (preserves server value). Set to 0 to clear the rating limit. Set to a positive value to enforce that rating as the maximum.",
			},
			"block_unrated_items": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of item types to block when unrated. Valid values: Movie, Trailer, Series, Music, Game, Book, LiveTvChannel, LiveTvProgram, ChannelContent, Other.",
			},
			// --- Policy: content features ---
			"enable_content_deletion": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can delete content.",
			},
			"enable_content_downloading": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can download content.",
			},
			"enable_subtitle_downloading": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can download subtitles.",
			},
			"enable_subtitle_management": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can manage subtitles.",
			},
			"enable_sync_transcoding": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can use sync transcoding.",
			},
			"enable_media_conversion": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can convert media.",
			},
			"enable_public_sharing": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can share content publicly.",
			},
			// --- Policy: playback ---
			"enable_audio_playback_transcoding": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can transcode audio.",
			},
			"enable_video_playback_transcoding": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can transcode video.",
			},
			"enable_playback_remuxing": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can remux playback.",
			},
			"remote_client_bitrate_limit": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Bitrate limit for remote clients in kbps. 0 means unlimited.",
			},
			"simultaneous_stream_limit": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Maximum number of simultaneous streams. 0 means unlimited.",
			},
			// --- Policy: live TV ---
			"enable_live_tv_access": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can access Live TV.",
			},
			"enable_live_tv_management": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can manage Live TV.",
			},
			// --- Policy: other ---
			"allow_camera_upload": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can upload from camera.",
			},
			"allow_sharing_personal_items": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the user can share personal items.",
			},
			// --- Configuration: profile ---
			"audio_language_preference": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Preferred audio language code.",
			},
			"play_default_audio_track": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to play the default audio track.",
			},
			"subtitle_language_preference": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Preferred subtitle language code.",
			},
			"subtitle_mode": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Subtitle playback mode. Valid values: Default, Always, OnlyForced, None, Smart, HearingImpaired.",
			},
			"display_missing_episodes": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to display missing episodes in series views.",
			},
			"ordered_views": schema.ListAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Ordered list of home screen view IDs.",
			},
			"latest_items_excludes": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Item types to exclude from 'Latest' section.",
			},
			"my_media_excludes": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Item types to exclude from 'My Media' section.",
			},
			"hide_played_in_latest": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to hide played items in the 'Latest' section.",
			},
			"hide_played_in_more_like_this": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to hide played items in 'More Like This'.",
			},
			"hide_played_in_suggestions": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to hide played items in suggestions.",
			},
			"remember_audio_selections": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to remember audio track selections.",
			},
			"remember_subtitle_selections": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to remember subtitle track selections.",
			},
			"enable_next_episode_auto_play": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to auto-play the next episode.",
			},
			"resume_rewind_seconds": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Seconds to rewind when resuming playback.",
			},
			"intro_skip_mode": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Intro skip behavior. Valid values: ShowButton, AutoSkip, None.",
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

	r.updateUserPolicy(plan, user, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	r.updateUserConfiguration(plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	r.updateUser(plan, user, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read back the full user so all computed fields are populated.
	reread, httpResp, err := r.data.Client.UserServiceAPI.GetUsersById(r.data.Auth, plan.Id.ValueString()).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Unable to read back user after creation", err.Error())
		return
	}
	mapUserToState(reread, &plan)

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

	user, _, err := r.data.Client.UserServiceAPI.GetUsersById(r.data.Auth, plan.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to read user for update", err.Error())
		return
	}

	r.updateUserPolicy(plan, user, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	r.updateUserConfiguration(plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	r.updateUser(plan, user, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read back the full user so all computed fields are populated.
	reread, httpResp, err := r.data.Client.UserServiceAPI.GetUsersById(r.data.Auth, plan.Id.ValueString()).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Unable to read back user after update", err.Error())
		return
	}
	mapUserToState(reread, &plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	userDeleteMutex.Lock()
	defer userDeleteMutex.Unlock()

	var state UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.data.Client.UserServiceAPI.PostUsersByIdDelete(r.data.Auth, state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to delete user", err.Error())
		return
	}

	// Verify the user is actually gone before confirming deletion.
	verify, verifyResp, verifyErr := r.data.Client.UserServiceAPI.GetUsersById(r.data.Auth, state.Id.ValueString()).Execute()
	if verifyErr == nil && verify != nil && verifyResp.StatusCode == 200 {
		resp.Diagnostics.AddError(
			"User still exists after delete",
			fmt.Sprintf("The Emby API returned success for the delete but user %q was found on verification. A concurrent delete may have silently failed. Retry on next apply.", state.Id.ValueString()),
		)
		return
	}
	if verifyErr != nil && verifyResp != nil && verifyResp.StatusCode != 404 {
		resp.Diagnostics.AddError(
			"Unable to verify user deletion",
			fmt.Sprintf("Delete appeared to succeed but verification failed: %s", verifyErr.Error()),
		)
		return
	}
}

func (r *UserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *UserResource) updateUserPolicy(plan UserResourceModel, user *embyclient.UserDto, diags *diag.Diagnostics) {
	policy := embyclient.UserPolicy{}
	if existing, ok := user.GetPolicyOk(); ok && existing != nil {
		policy = *existing
	}

	// General
	policy.SetIsAdministrator(plan.IsAdministrator.ValueBool())
	policy.SetIsHidden(plan.IsHidden.ValueBool())
	policy.SetIsHiddenRemotely(plan.IsHiddenRemotely.ValueBool())
	policy.SetIsHiddenFromUnusedDevices(plan.IsHiddenFromUnusedDevices.ValueBool())
	policy.SetIsDisabled(plan.IsDisabled.ValueBool())
	policy.SetEnableRemoteAccess(plan.EnableRemoteAccess.ValueBool())
	policy.SetEnableMediaPlayback(plan.EnableMediaPlayback.ValueBool())
	if !plan.AuthenticationProviderId.IsNull() && !plan.AuthenticationProviderId.IsUnknown() {
		policy.SetAuthenticationProviderId(plan.AuthenticationProviderId.ValueString())
	}

	// Library access
	policy.SetEnableAllFolders(plan.EnableAllFolders.ValueBool())
	setStringSliceFromSet(plan.EnabledFolders, policy.SetEnabledFolders, diags)
	setStringSliceFromSet(plan.ExcludedSubFolders, policy.SetExcludedSubFolders, diags)

	// Channel access
	policy.SetEnableAllChannels(plan.EnableAllChannels.ValueBool())
	setStringSliceFromSet(plan.EnabledChannels, policy.SetEnabledChannels, diags)

	// Device access
	policy.SetEnableAllDevices(plan.EnableAllDevices.ValueBool())
	setStringSliceFromSet(plan.EnabledDevices, policy.SetEnabledDevices, diags)

	// Parental controls
	setStringSliceFromSet(plan.BlockedTags, policy.SetBlockedTags, diags)
	setStringSliceFromSet(plan.IncludeTags, policy.SetIncludeTags, diags)
	policy.SetIsTagBlockingModeInclusive(plan.IsTagBlockingModeInclusive.ValueBool())
	policy.SetAllowTagOrRating(plan.AllowTagOrRating.ValueBool())

	if plan.MaxParentalRating.IsNull() || plan.MaxParentalRating.IsUnknown() {
		// Preserve existing value - don't modify
	} else if plan.MaxParentalRating.ValueInt64() == 0 {
		policy.SetMaxParentalRatingNil()
	} else {
		policy.SetMaxParentalRating(int32(plan.MaxParentalRating.ValueInt64()))
	}

	if !plan.BlockUnratedItems.IsNull() && !plan.BlockUnratedItems.IsUnknown() {
		var items []string
		diags.Append(plan.BlockUnratedItems.ElementsAs(context.Background(), &items, false)...)
		if !diags.HasError() {
			unrated := make([]embyclient.UnratedItem, len(items))
			for i, item := range items {
				unrated[i] = embyclient.UnratedItem(item)
			}
			policy.SetBlockUnratedItems(unrated)
		}
	}

	// Content features
	policy.SetEnableContentDeletion(plan.EnableContentDeletion.ValueBool())
	policy.SetEnableContentDownloading(plan.EnableContentDownloading.ValueBool())
	policy.SetEnableSubtitleDownloading(plan.EnableSubtitleDownloading.ValueBool())
	policy.SetEnableSubtitleManagement(plan.EnableSubtitleManagement.ValueBool())
	policy.SetEnableSyncTranscoding(plan.EnableSyncTranscoding.ValueBool())
	policy.SetEnableMediaConversion(plan.EnableMediaConversion.ValueBool())
	policy.SetEnablePublicSharing(plan.EnablePublicSharing.ValueBool())

	// Playback
	policy.SetEnableAudioPlaybackTranscoding(plan.EnableAudioPlaybackTranscoding.ValueBool())
	policy.SetEnableVideoPlaybackTranscoding(plan.EnableVideoPlaybackTranscoding.ValueBool())
	policy.SetEnablePlaybackRemuxing(plan.EnablePlaybackRemuxing.ValueBool())
	policy.SetRemoteClientBitrateLimit(int32(plan.RemoteClientBitrateLimit.ValueInt64()))
	policy.SetSimultaneousStreamLimit(int32(plan.SimultaneousStreamLimit.ValueInt64()))

	// Live TV
	policy.SetEnableLiveTvAccess(plan.EnableLiveTvAccess.ValueBool())
	policy.SetEnableLiveTvManagement(plan.EnableLiveTvManagement.ValueBool())

	// Other
	policy.SetAllowCameraUpload(plan.AllowCameraUpload.ValueBool())
	policy.SetAllowSharingPersonalItems(plan.AllowSharingPersonalItems.ValueBool())

	_, err := r.data.Client.UserServiceAPI.PostUsersByIdPolicy(r.data.Auth, user.GetId()).UserPolicy(policy).Execute()
	if err != nil {
		diags.AddError("Unable to update user policy", err.Error())
	}
}

func (r *UserResource) updateUserConfiguration(plan UserResourceModel, diags *diag.Diagnostics) {
	config := embyclient.UserConfiguration{}

	if !plan.AudioLanguagePreference.IsNull() && !plan.AudioLanguagePreference.IsUnknown() {
		config.SetAudioLanguagePreference(plan.AudioLanguagePreference.ValueString())
	}
	if !plan.PlayDefaultAudioTrack.IsNull() && !plan.PlayDefaultAudioTrack.IsUnknown() {
		config.SetPlayDefaultAudioTrack(plan.PlayDefaultAudioTrack.ValueBool())
	}
	if !plan.SubtitleLanguagePreference.IsNull() && !plan.SubtitleLanguagePreference.IsUnknown() {
		config.SetSubtitleLanguagePreference(plan.SubtitleLanguagePreference.ValueString())
	}
	if !plan.SubtitleMode.IsNull() && !plan.SubtitleMode.IsUnknown() {
		config.SetSubtitleMode(embyclient.SubtitlePlaybackMode(plan.SubtitleMode.ValueString()))
	}
	if !plan.DisplayMissingEpisodes.IsNull() && !plan.DisplayMissingEpisodes.IsUnknown() {
		config.SetDisplayMissingEpisodes(plan.DisplayMissingEpisodes.ValueBool())
	}
	if !plan.OrderedViews.IsNull() && !plan.OrderedViews.IsUnknown() {
		var views []string
		diags.Append(plan.OrderedViews.ElementsAs(context.Background(), &views, false)...)
		if !diags.HasError() {
			config.SetOrderedViews(views)
		}
	}
	if !plan.LatestItemsExcludes.IsNull() && !plan.LatestItemsExcludes.IsUnknown() {
		var items []string
		diags.Append(plan.LatestItemsExcludes.ElementsAs(context.Background(), &items, false)...)
		if !diags.HasError() {
			config.SetLatestItemsExcludes(items)
		}
	}
	if !plan.MyMediaExcludes.IsNull() && !plan.MyMediaExcludes.IsUnknown() {
		var items []string
		diags.Append(plan.MyMediaExcludes.ElementsAs(context.Background(), &items, false)...)
		if !diags.HasError() {
			config.SetMyMediaExcludes(items)
		}
	}
	if !plan.HidePlayedInLatest.IsNull() && !plan.HidePlayedInLatest.IsUnknown() {
		config.SetHidePlayedInLatest(plan.HidePlayedInLatest.ValueBool())
	}
	if !plan.HidePlayedInMoreLikeThis.IsNull() && !plan.HidePlayedInMoreLikeThis.IsUnknown() {
		config.SetHidePlayedInMoreLikeThis(plan.HidePlayedInMoreLikeThis.ValueBool())
	}
	if !plan.HidePlayedInSuggestions.IsNull() && !plan.HidePlayedInSuggestions.IsUnknown() {
		config.SetHidePlayedInSuggestions(plan.HidePlayedInSuggestions.ValueBool())
	}
	if !plan.RememberAudioSelections.IsNull() && !plan.RememberAudioSelections.IsUnknown() {
		config.SetRememberAudioSelections(plan.RememberAudioSelections.ValueBool())
	}
	if !plan.RememberSubtitleSelections.IsNull() && !plan.RememberSubtitleSelections.IsUnknown() {
		config.SetRememberSubtitleSelections(plan.RememberSubtitleSelections.ValueBool())
	}
	if !plan.EnableNextEpisodeAutoPlay.IsNull() && !plan.EnableNextEpisodeAutoPlay.IsUnknown() {
		config.SetEnableNextEpisodeAutoPlay(plan.EnableNextEpisodeAutoPlay.ValueBool())
	}
	if !plan.ResumeRewindSeconds.IsNull() && !plan.ResumeRewindSeconds.IsUnknown() {
		config.SetResumeRewindSeconds(int32(plan.ResumeRewindSeconds.ValueInt64()))
	}
	if !plan.IntroSkipMode.IsNull() && !plan.IntroSkipMode.IsUnknown() {
		config.SetIntroSkipMode(embyclient.SegmentSkipMode(plan.IntroSkipMode.ValueString()))
	}

	_, err := r.data.Client.UserServiceAPI.PostUsersByIdConfiguration(r.data.Auth, plan.Id.ValueString()).UserConfiguration(config).Execute()
	if err != nil {
		diags.AddError("Unable to update user configuration", err.Error())
	}
}

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
	state.IsHiddenRemotely = types.BoolValue(policy.GetIsHiddenRemotely())
	state.IsHiddenFromUnusedDevices = types.BoolValue(policy.GetIsHiddenFromUnusedDevices())
	state.IsDisabled = types.BoolValue(policy.GetIsDisabled())
	state.EnableRemoteAccess = types.BoolValue(policy.GetEnableRemoteAccess())
	state.EnableMediaPlayback = types.BoolValue(policy.GetEnableMediaPlayback())
	state.AuthenticationProviderId = types.StringValue(policy.GetAuthenticationProviderId())

	// Library access
	state.EnableAllFolders = types.BoolValue(policy.GetEnableAllFolders())
	state.EnabledFolders = stringSliceToSetValue(policy.GetEnabledFolders())
	state.ExcludedSubFolders = stringSliceToSetValue(policy.GetExcludedSubFolders())

	// Channel access
	state.EnableAllChannels = types.BoolValue(policy.GetEnableAllChannels())
	state.EnabledChannels = stringSliceToSetValue(policy.GetEnabledChannels())

	// Device access
	state.EnableAllDevices = types.BoolValue(policy.GetEnableAllDevices())
	state.EnabledDevices = stringSliceToSetValue(policy.GetEnabledDevices())

	// Parental controls
	state.BlockedTags = stringSliceToSetValue(policy.GetBlockedTags())
	state.IncludeTags = stringSliceToSetValue(policy.GetIncludeTags())
	state.IsTagBlockingModeInclusive = types.BoolValue(policy.GetIsTagBlockingModeInclusive())
	state.AllowTagOrRating = types.BoolValue(policy.GetAllowTagOrRating())
	if rating, ok := policy.GetMaxParentalRatingOk(); ok {
		state.MaxParentalRating = types.Int64Value(int64(*rating))
	} else {
		state.MaxParentalRating = types.Int64Null()
	}

	unrated := policy.GetBlockUnratedItems()
	unratedStrs := make([]string, len(unrated))
	for i, u := range unrated {
		unratedStrs[i] = string(u)
	}
	state.BlockUnratedItems = stringSliceToSetValue(unratedStrs)

	// Content features
	state.EnableContentDeletion = types.BoolValue(policy.GetEnableContentDeletion())
	state.EnableContentDownloading = types.BoolValue(policy.GetEnableContentDownloading())
	state.EnableSubtitleDownloading = types.BoolValue(policy.GetEnableSubtitleDownloading())
	state.EnableSubtitleManagement = types.BoolValue(policy.GetEnableSubtitleManagement())
	state.EnableSyncTranscoding = types.BoolValue(policy.GetEnableSyncTranscoding())
	state.EnableMediaConversion = types.BoolValue(policy.GetEnableMediaConversion())
	state.EnablePublicSharing = types.BoolValue(policy.GetEnablePublicSharing())

	// Playback
	state.EnableAudioPlaybackTranscoding = types.BoolValue(policy.GetEnableAudioPlaybackTranscoding())
	state.EnableVideoPlaybackTranscoding = types.BoolValue(policy.GetEnableVideoPlaybackTranscoding())
	state.EnablePlaybackRemuxing = types.BoolValue(policy.GetEnablePlaybackRemuxing())
	state.RemoteClientBitrateLimit = types.Int64Value(int64(policy.GetRemoteClientBitrateLimit()))
	state.SimultaneousStreamLimit = types.Int64Value(int64(policy.GetSimultaneousStreamLimit()))

	// Live TV
	state.EnableLiveTvAccess = types.BoolValue(policy.GetEnableLiveTvAccess())
	state.EnableLiveTvManagement = types.BoolValue(policy.GetEnableLiveTvManagement())

	// Other
	state.AllowCameraUpload = types.BoolValue(policy.GetAllowCameraUpload())
	state.AllowSharingPersonalItems = types.BoolValue(policy.GetAllowSharingPersonalItems())

	// Configuration
	config := user.GetConfiguration()
	state.AudioLanguagePreference = types.StringValue(config.GetAudioLanguagePreference())
	state.PlayDefaultAudioTrack = types.BoolValue(config.GetPlayDefaultAudioTrack())
	state.SubtitleLanguagePreference = types.StringValue(config.GetSubtitleLanguagePreference())
	state.SubtitleMode = types.StringValue(string(config.GetSubtitleMode()))
	state.DisplayMissingEpisodes = types.BoolValue(config.GetDisplayMissingEpisodes())
	state.OrderedViews = stringSliceToListValue(config.GetOrderedViews())
	state.LatestItemsExcludes = stringSliceToSetValue(config.GetLatestItemsExcludes())
	state.MyMediaExcludes = stringSliceToSetValue(config.GetMyMediaExcludes())
	state.HidePlayedInLatest = types.BoolValue(config.GetHidePlayedInLatest())
	state.HidePlayedInMoreLikeThis = types.BoolValue(config.GetHidePlayedInMoreLikeThis())
	state.HidePlayedInSuggestions = types.BoolValue(config.GetHidePlayedInSuggestions())
	state.RememberAudioSelections = types.BoolValue(config.GetRememberAudioSelections())
	state.RememberSubtitleSelections = types.BoolValue(config.GetRememberSubtitleSelections())
	state.EnableNextEpisodeAutoPlay = types.BoolValue(config.GetEnableNextEpisodeAutoPlay())
	state.ResumeRewindSeconds = types.Int64Value(int64(config.GetResumeRewindSeconds()))
	state.IntroSkipMode = types.StringValue(string(config.GetIntroSkipMode()))
}

func stringSliceToSetValue(items []string) types.Set {
	if len(items) == 0 {
		return types.SetValueMust(types.StringType, []attr.Value{})
	}
	vals := make([]attr.Value, len(items))
	for i, item := range items {
		vals[i] = types.StringValue(item)
	}
	return types.SetValueMust(types.StringType, vals)
}

func stringSliceToListValue(items []string) types.List {
	if len(items) == 0 {
		return types.ListValueMust(types.StringType, []attr.Value{})
	}
	vals := make([]attr.Value, len(items))
	for i, item := range items {
		vals[i] = types.StringValue(item)
	}
	return types.ListValueMust(types.StringType, vals)
}

func setStringSliceFromSet(set types.Set, setter func([]string), diags *diag.Diagnostics) {
	if set.IsNull() || set.IsUnknown() {
		return
	}
	var items []string
	diags.Append(set.ElementsAs(context.Background(), &items, false)...)
	if !diags.HasError() {
		setter(items)
	}
}
