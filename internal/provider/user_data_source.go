// Copyright (c) HashiCorp, Inc.

package provider

import (
	"context"
	"fmt"

	embyclient "github.com/Kamaroth92/terraform-provider-emby/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &UserDataSource{}

type UserDataSource struct {
	data *EmbyProviderData
}

type UserDataSourceModel struct {
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

func NewUserDataSource() datasource.DataSource {
	return &UserDataSource{}
}

func (d *UserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (d *UserDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Gets information about an Emby user by ID or name. Returns all policy and configuration fields.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The user ID.",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The username.",
			},
			// --- Policy: general ---
			"is_administrator": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user is an administrator.",
			},
			"is_hidden": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user is hidden from login screens.",
			},
			"is_hidden_remotely": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user is hidden from remote connections.",
			},
			"is_hidden_from_unused_devices": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user is hidden from unused devices.",
			},
			"is_disabled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user account is disabled.",
			},
			"enable_remote_access": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can access remotely.",
			},
			"enable_media_playback": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can play media.",
			},
			"authentication_provider_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The authentication provider ID for this user.",
			},
			// --- Policy: library access ---
			"enable_all_folders": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can access all library folders.",
			},
			"enabled_folders": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of library IDs the user is allowed to access.",
			},
			"excluded_sub_folders": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of sub-folder paths the user is denied access to.",
			},
			// --- Policy: channel access ---
			"enable_all_channels": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can access all channels.",
			},
			"enabled_channels": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of channel IDs the user is allowed to access.",
			},
			// --- Policy: device access ---
			"enable_all_devices": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can use all devices.",
			},
			"enabled_devices": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of device IDs the user is allowed to use.",
			},
			// --- Policy: parental controls ---
			"blocked_tags": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of content tags to block or allow.",
			},
			"include_tags": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Set of content tags to include.",
			},
			"is_tag_blocking_mode_inclusive": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Tag blocking mode. false = blocklist, true = allowlist.",
			},
			"allow_tag_or_rating": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether a single tag OR rating match is sufficient to restrict content.",
			},
			"max_parental_rating": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Maximum allowed parental rating. Null means no rating limit.",
			},
			"block_unrated_items": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Item types to block when unrated.",
			},
			// --- Policy: content features ---
			"enable_content_deletion": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can delete content.",
			},
			"enable_content_downloading": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can download content.",
			},
			"enable_subtitle_downloading": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can download subtitles.",
			},
			"enable_subtitle_management": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can manage subtitles.",
			},
			"enable_sync_transcoding": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can use sync transcoding.",
			},
			"enable_media_conversion": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can convert media.",
			},
			"enable_public_sharing": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can share content publicly.",
			},
			// --- Policy: playback ---
			"enable_audio_playback_transcoding": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can transcode audio.",
			},
			"enable_video_playback_transcoding": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can transcode video.",
			},
			"enable_playback_remuxing": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can remux playback.",
			},
			"remote_client_bitrate_limit": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Bitrate limit for remote clients in kbps.",
			},
			"simultaneous_stream_limit": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Maximum number of simultaneous streams.",
			},
			// --- Policy: live TV ---
			"enable_live_tv_access": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can access Live TV.",
			},
			"enable_live_tv_management": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can manage Live TV.",
			},
			// --- Policy: other ---
			"allow_camera_upload": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can upload from camera.",
			},
			"allow_sharing_personal_items": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user can share personal items.",
			},
			// --- Configuration: profile ---
			"audio_language_preference": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Preferred audio language code.",
			},
			"play_default_audio_track": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether to play the default audio track.",
			},
			"subtitle_language_preference": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Preferred subtitle language code.",
			},
			"subtitle_mode": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Subtitle playback mode.",
			},
			"display_missing_episodes": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether to display missing episodes.",
			},
			"ordered_views": schema.ListAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Ordered list of home screen view IDs.",
			},
			"latest_items_excludes": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Item types excluded from 'Latest' section.",
			},
			"my_media_excludes": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Item types excluded from 'My Media' section.",
			},
			"hide_played_in_latest": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether to hide played items in 'Latest'.",
			},
			"hide_played_in_more_like_this": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether to hide played items in 'More Like This'.",
			},
			"hide_played_in_suggestions": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether to hide played items in suggestions.",
			},
			"remember_audio_selections": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether to remember audio track selections.",
			},
			"remember_subtitle_selections": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether to remember subtitle track selections.",
			},
			"enable_next_episode_auto_play": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether to auto-play the next episode.",
			},
			"resume_rewind_seconds": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Seconds to rewind when resuming playback.",
			},
			"intro_skip_mode": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Intro skip behavior.",
			},
		},
	}
}

func (d *UserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(*EmbyProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *EmbyProviderData, got: %T.", req.ProviderData),
		)
		return
	}

	d.data = data
}

func (d *UserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config UserDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !config.Name.IsNull() && !config.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Multiple identifiers specified",
			"Only one of name or id should be specified to identify the user.",
		)
		return
	}

	if config.Name.IsNull() && config.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Missing user identifier",
			"Either name or id must be specified.",
		)
		return
	}

	if !config.Id.IsNull() {
		user, _, err := d.data.Client.UserServiceAPI.GetUsersById(d.data.Auth, config.Id.ValueString()).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Unable to read user", err.Error())
			return
		}
		mapUserDtoToDataSourceModel(user, &config)
		resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
		return
	}

	result, _, err := d.data.Client.UserServiceAPI.GetUsersQuery(d.data.Auth).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to query users", err.Error())
		return
	}

	userName := config.Name.ValueString()
	var found *embyclient.UserDto
	for _, user := range result.GetItems() {
		if user.GetName() == userName {
			found = &user
			break
		}
	}

	if found == nil {
		resp.Diagnostics.AddError(
			"User not found",
			fmt.Sprintf("No user found with name %q", userName),
		)
		return
	}

	mapUserDtoToDataSourceModel(found, &config)
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

func mapUserDtoToDataSourceModel(user *embyclient.UserDto, state *UserDataSourceModel) {
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

	// Other policy
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
