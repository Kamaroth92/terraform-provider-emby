// Copyright (c) HashiCorp, Inc.

package provider

import (
	"context"
	"fmt"

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

var (
	_ resource.Resource                = &LibraryResource{}
	_ resource.ResourceWithImportState = &LibraryResource{}
)

type LibraryResource struct {
	data *EmbyProviderData
}

type LibraryResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	CollectionType types.String `tfsdk:"collection_type"`
	Paths          types.List   `tfsdk:"paths"`
	ItemId         types.String `tfsdk:"item_id"`
	Guid           types.String `tfsdk:"guid"`
	Locations      types.List   `tfsdk:"locations"`

	// --- General ---
	EnableRealtimeMonitor                  types.Bool  `tfsdk:"enable_realtime_monitor"`
	EnablePhotos                           types.Bool  `tfsdk:"enable_photos"`
	EnableArchiveMediaFiles                types.Bool  `tfsdk:"enable_archive_media_files"`
	EnableMarkerDetection                  types.Bool  `tfsdk:"enable_marker_detection"`
	EnableMarkerDetectionDuringLibraryScan types.Bool  `tfsdk:"enable_marker_detection_during_library_scan"`
	IntroDetectionFingerprintLength        types.Int64 `tfsdk:"intro_detection_fingerprint_length"`

	// --- Chapter / Image Extraction ---
	EnableChapterImageExtraction          types.Bool `tfsdk:"enable_chapter_image_extraction"`
	ExtractChapterImagesDuringLibraryScan types.Bool `tfsdk:"extract_chapter_images_during_library_scan"`
	DownloadImagesInAdvance               types.Bool `tfsdk:"download_images_in_advance"`
	CacheImages                           types.Bool `tfsdk:"cache_images"`

	// --- Metadata ---
	PreferredMetadataLanguage      types.String `tfsdk:"preferred_metadata_language"`
	PreferredImageLanguage         types.String `tfsdk:"preferred_image_language"`
	MetadataCountryCode            types.String `tfsdk:"metadata_country_code"`
	ContentType                    types.String `tfsdk:"content_type"`
	MetadataSavers                 types.Set    `tfsdk:"metadata_savers"`
	DisabledLocalMetadataReaders   types.Set    `tfsdk:"disabled_local_metadata_readers"`
	LocalMetadataReaderOrder       types.List   `tfsdk:"local_metadata_reader_order"`
	SaveLocalMetadata              types.Bool   `tfsdk:"save_local_metadata"`
	SaveMetadataHidden             types.Bool   `tfsdk:"save_metadata_hidden"`
	SaveLocalThumbnailSets         types.Bool   `tfsdk:"save_local_thumbnail_sets"`
	EnableAdultMetadata            types.Bool   `tfsdk:"enable_adult_metadata"`

	// --- File Handling ---
	IgnoreHiddenFiles    types.Bool `tfsdk:"ignore_hidden_files"`
	IgnoreFileExtensions types.Set  `tfsdk:"ignore_file_extensions"`
	EnablePlexIgnore     types.Bool `tfsdk:"enable_plex_ignore"`
	ImportPlaylists      types.Bool `tfsdk:"import_playlists"`
	ImportCollections    types.Bool `tfsdk:"import_collections"`

	// --- Subtitles ---
	DisabledSubtitleFetchers                types.Set    `tfsdk:"disabled_subtitle_fetchers"`
	SubtitleFetcherOrder                    types.List   `tfsdk:"subtitle_fetcher_order"`
	SubtitleDownloadLanguages               types.Set    `tfsdk:"subtitle_download_languages"`
	SubtitleDownloadMaxAgeDays              types.Int64  `tfsdk:"subtitle_download_max_age_days"`
	SkipSubtitlesIfEmbeddedSubtitlesPresent types.Bool   `tfsdk:"skip_subtitles_if_embedded_subtitles_present"`
	SkipSubtitlesIfAudioTrackMatches        types.Bool   `tfsdk:"skip_subtitles_if_audio_track_matches"`
	RequirePerfectSubtitleMatch             types.Bool   `tfsdk:"require_perfect_subtitle_match"`
	SaveSubtitlesWithMedia                  types.Bool   `tfsdk:"save_subtitles_with_media"`
	ForcedSubtitlesOnly                     types.Bool   `tfsdk:"forced_subtitles_only"`
	HearingImpairedSubtitlesOnly            types.Bool   `tfsdk:"hearing_impaired_subtitles_only"`

	// --- Lyrics ---
	DisabledLyricsFetchers   types.Set   `tfsdk:"disabled_lyrics_fetchers"`
	LyricsFetcherOrder       types.List  `tfsdk:"lyrics_fetcher_order"`
	LyricsDownloadLanguages  types.Set   `tfsdk:"lyrics_download_languages"`
	LyricsDownloadMaxAgeDays types.Int64 `tfsdk:"lyrics_download_max_age_days"`
	SaveLyricsWithMedia      types.Bool  `tfsdk:"save_lyrics_with_media"`

	// --- Advanced ---
	EnableEmbeddedTitles           types.Bool   `tfsdk:"enable_embedded_titles"`
	EnableAudioResume              types.Bool   `tfsdk:"enable_audio_resume"`
	AutoGenerateChapters           types.Bool   `tfsdk:"auto_generate_chapters"`
	AutoGenerateChapterIntervalMinutes types.Int64 `tfsdk:"auto_generate_chapter_interval_minutes"`
	AutomaticRefreshIntervalDays       types.Int64 `tfsdk:"automatic_refresh_interval_days"`
	PlaceholderMetadataRefreshIntervalDays types.Int64 `tfsdk:"placeholder_metadata_refresh_interval_days"`
	MergeTopLevelFolders               types.Bool  `tfsdk:"merge_top_level_folders"`
	CollapseSingleItemFolders          types.Bool  `tfsdk:"collapse_single_item_folders"`
	ForceCollapseSingleItemFolders     types.Bool  `tfsdk:"force_collapse_single_item_folders"`
	EnableAutomaticSeriesGrouping      types.Bool  `tfsdk:"enable_automatic_series_grouping"`
	ShareEmbeddedMusicAlbumImages      types.Bool  `tfsdk:"share_embedded_music_album_images"`
	EnableMultiVersionByFiles          types.Bool  `tfsdk:"enable_multi_version_by_files"`
	EnableMultiVersionByMetadata       types.Bool  `tfsdk:"enable_multi_version_by_metadata"`
	EnableMultiPartItems               types.Bool  `tfsdk:"enable_multi_part_items"`
	MinCollectionItems                 types.Int64 `tfsdk:"min_collection_items"`
	MusicFolderStructure               types.String `tfsdk:"music_folder_structure"`
	ExcludeFromSearch                  types.Bool  `tfsdk:"exclude_from_search"`

	// --- Resume / Playback ---
	MinResumePct                   types.Int64 `tfsdk:"min_resume_pct"`
	MaxResumePct                   types.Int64 `tfsdk:"max_resume_pct"`
	MinResumeDurationSeconds       types.Int64 `tfsdk:"min_resume_duration_seconds"`
	ThumbnailImagesIntervalSeconds types.Int64 `tfsdk:"thumbnail_images_interval_seconds"`
	SampleIgnoreSize               types.Int64 `tfsdk:"sample_ignore_size"`
}

func NewLibraryResource() resource.Resource {
	return &LibraryResource{}
}

func (r *LibraryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_library"
}

func (r *LibraryResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Emby library (virtual folder).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The library ID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The library name.",
			},
			"collection_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The collection type (e.g. movies, tvshows, music, books, homevideos, musicvideos, mixed). Changing this forces the library to be recreated.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"paths": schema.ListAttribute{
				Required:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "The file system paths for the library. At least one path is required.",
			},
			"item_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The internal item ID used for policy bindings (e.g. EnabledFolders).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"guid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The library GUID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"locations": schema.ListAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "The file system paths for the library, as reported by the server.",
			},
			// --- General ---
			"enable_realtime_monitor": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enable real time monitoring of file changes.",
			},
			"enable_photos": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enable photos in this library.",
			},
			"enable_archive_media_files": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enable archive media files (e.g. RAR, ZIP) support.",
			},
			"enable_marker_detection": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enable intro/credits marker detection.",
			},
			"enable_marker_detection_during_library_scan": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Run marker detection during scheduled library scans.",
			},
			"intro_detection_fingerprint_length": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Fingerprint length in seconds for intro detection.",
			},
			// --- Chapter / Image Extraction ---
			"enable_chapter_image_extraction": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Extract chapter images from video files.",
			},
			"extract_chapter_images_during_library_scan": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Extract chapter images during scheduled library scans.",
			},
			"download_images_in_advance": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Download images in advance instead of on-demand.",
			},
			"cache_images": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Cache downloaded images locally.",
			},
			// --- Metadata ---
			"preferred_metadata_language": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Preferred language for metadata (e.g. en, fr, de).",
			},
			"preferred_image_language": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Preferred language for images (e.g. en, fr, de).",
			},
			"metadata_country_code": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Country code used for metadata ratings (e.g. US, GB).",
			},
			"content_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The content type override for the library.",
			},
			"metadata_savers": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Metadata saver plugins to use (e.g. Emby Xml, Nfo).",
			},
			"disabled_local_metadata_readers": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Local metadata reader plugins to disable.",
			},
			"local_metadata_reader_order": schema.ListAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Order of local metadata reader plugins.",
			},
			"save_local_metadata": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Save metadata alongside media files (NFO, etc).",
			},
			"save_metadata_hidden": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Save metadata as hidden files.",
			},
			"save_local_thumbnail_sets": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Save thumbnail images alongside media files.",
			},
			"enable_adult_metadata": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enable metadata from adult sources.",
			},
			// --- File Handling ---
			"ignore_hidden_files": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Ignore hidden files when scanning the library.",
			},
			"ignore_file_extensions": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "File extensions to ignore during library scans.",
			},
			"enable_plex_ignore": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Respect .plexignore files in library folders.",
			},
			"import_playlists": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Import playlists from the library folders.",
			},
			"import_collections": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Import collections from the library folders.",
			},
			// --- Subtitles ---
			"disabled_subtitle_fetchers": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Subtitle fetcher plugins to disable.",
			},
			"subtitle_fetcher_order": schema.ListAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Priority order of subtitle fetcher plugins.",
			},
			"subtitle_download_languages": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Languages to download subtitles for (e.g. eng, fre).",
			},
			"subtitle_download_max_age_days": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Maximum age in days for subtitle downloads. 0 for any age.",
			},
			"skip_subtitles_if_embedded_subtitles_present": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Skip downloading subtitles if the media already has embedded subtitles.",
			},
			"skip_subtitles_if_audio_track_matches": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Skip downloading subtitles if the audio track language matches.",
			},
			"require_perfect_subtitle_match": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Only download subtitles that are a perfect match for the media.",
			},
			"save_subtitles_with_media": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Save downloaded subtitles alongside media files.",
			},
			"forced_subtitles_only": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Only download forced/foreign subtitles.",
			},
			"hearing_impaired_subtitles_only": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Only download subtitles for the hearing impaired (SDH).",
			},
			// --- Lyrics ---
			"disabled_lyrics_fetchers": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Lyrics fetcher plugins to disable.",
			},
			"lyrics_fetcher_order": schema.ListAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Priority order of lyrics fetcher plugins.",
			},
			"lyrics_download_languages": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Languages to download lyrics for.",
			},
			"lyrics_download_max_age_days": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Maximum age in days for lyrics downloads. 0 for any age.",
			},
			"save_lyrics_with_media": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Save downloaded lyrics alongside media files.",
			},
			// --- Advanced ---
			"enable_embedded_titles": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Use embedded titles from media files instead of filenames.",
			},
			"enable_audio_resume": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enable audio playback resume.",
			},
			"auto_generate_chapters": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Automatically generate chapters for media files.",
			},
			"auto_generate_chapter_interval_minutes": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Interval in minutes for auto-generated chapter points.",
			},
			"automatic_refresh_interval_days": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Number of days between automatic library refreshes. 0 to disable.",
			},
			"placeholder_metadata_refresh_interval_days": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Number of days between metadata refreshes for items with placeholder metadata.",
			},
			"merge_top_level_folders": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Merge top-level folders into a single view for multi-path libraries.",
			},
			"collapse_single_item_folders": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Collapse folders that contain only a single item.",
			},
			"force_collapse_single_item_folders": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Force collapse single-item folders even when metadata disagrees.",
			},
			"enable_automatic_series_grouping": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Automatically group movies into collections based on metadata.",
			},
			"share_embedded_music_album_images": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Share embedded album art images across tracks in the same album.",
			},
			"enable_multi_version_by_files": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Group multiple versions of the same media by file naming.",
			},
			"enable_multi_version_by_metadata": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Group multiple versions of the same media by metadata.",
			},
			"enable_multi_part_items": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enable multi-part file support (e.g. CD1, CD2).",
			},
			"min_collection_items": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Minimum number of items required to create a collection.",
			},
			"music_folder_structure": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Folder structure pattern for music libraries.",
			},
			"exclude_from_search": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Exclude this library from global search results.",
			},
			// --- Resume / Playback ---
			"min_resume_pct": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Minimum percentage played to allow resume.",
			},
			"max_resume_pct": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Maximum percentage played to still allow resume.",
			},
			"min_resume_duration_seconds": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Minimum duration in seconds for an item to be resumable.",
			},
			"thumbnail_images_interval_seconds": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Interval in seconds between thumbnail image generation.",
			},
			"sample_ignore_size": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Ignore sample files smaller than this size in bytes.",
			},
		},
	}
}

func (r *LibraryResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *LibraryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan LibraryResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var paths []string
	resp.Diagnostics.Append(plan.Paths.ElementsAs(ctx, &paths, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := embyclient.NewLibraryAddVirtualFolder()
	createReq.SetName(plan.Name.ValueString())
	createReq.SetCollectionType(plan.CollectionType.ValueString())
	createReq.SetPaths(paths)
	createReq.SetRefreshLibrary(true)

	// Only include LibraryOptions if at least one option is explicitly set.
	// Sending an empty LibraryOptions object can cause the server to reset
	// defaults (e.g. clearing CollectionType).
	opts := buildLibraryOptions(plan)
	if !libraryOptionsEmpty(opts) {
		createReq.SetLibraryOptions(opts)
	}

	httpResp, err := r.data.Client.LibraryStructureServiceAPI.PostLibraryVirtualfolders(r.data.Auth).
		LibraryAddVirtualFolder(*createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to create library", err.Error())
		return
	}
	if httpResp != nil && httpResp.StatusCode >= 300 {
		resp.Diagnostics.AddError("Unexpected status creating library", fmt.Sprintf("HTTP %d", httpResp.StatusCode))
		return
	}

	found := r.findLibraryByName(plan.Name.ValueString(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if found == nil {
		resp.Diagnostics.AddError("Library not found after creation", fmt.Sprintf("Library %q was created but could not be found in the library list.", plan.Name.ValueString()))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, found)...)
}

func (r *LibraryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LibraryResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	found := r.findLibraryById(state.Id.ValueString(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if found == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, found)...)
}

func (r *LibraryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan LibraryResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state LibraryResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = state.Id

	// Handle name change.
	if plan.Name.ValueString() != state.Name.ValueString() {
		renameReq := embyclient.NewLibraryRenameVirtualFolder()
		renameReq.SetId(plan.Id.ValueString())
		renameReq.SetNewName(plan.Name.ValueString())

		_, err := r.data.Client.LibraryStructureServiceAPI.PostLibraryVirtualfoldersName(r.data.Auth).
			LibraryRenameVirtualFolder(*renameReq).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Unable to rename library", err.Error())
			return
		}
	}

	// Handle path changes.
	var planPaths, statePaths []string
	resp.Diagnostics.Append(plan.Paths.ElementsAs(ctx, &planPaths, false)...)
	resp.Diagnostics.Append(state.Paths.ElementsAs(ctx, &statePaths, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	statePathSet := make(map[string]bool, len(statePaths))
	for _, p := range statePaths {
		statePathSet[p] = true
	}
	planPathSet := make(map[string]bool, len(planPaths))
	for _, p := range planPaths {
		planPathSet[p] = true
	}

	for _, p := range planPaths {
		if !statePathSet[p] {
			addReq := embyclient.NewLibraryAddMediaPath()
			addReq.SetId(plan.Id.ValueString())
			addReq.SetPath(p)
			addReq.SetRefreshLibrary(true)

			_, err := r.data.Client.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPaths(r.data.Auth).
				LibraryAddMediaPath(*addReq).Execute()
			if err != nil {
				resp.Diagnostics.AddError("Unable to add library path", err.Error())
				return
			}
		}
	}

	for _, p := range statePaths {
		if !planPathSet[p] {
			removeReq := embyclient.NewLibraryRemoveMediaPath()
			removeReq.SetId(plan.Id.ValueString())
			removeReq.SetPath(p)
			removeReq.SetRefreshLibrary(true)

			_, err := r.data.Client.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsDelete(r.data.Auth).
				LibraryRemoveMediaPath(*removeReq).Execute()
			if err != nil {
				resp.Diagnostics.AddError("Unable to remove library path", err.Error())
				return
			}
		}
	}

	// Update library options only if something changed.
	planOpts := buildLibraryOptions(plan)
	stateOpts := buildLibraryOptions(state)
	if !libraryOptionsEqual(planOpts, stateOpts) {
		updateOpts := embyclient.NewLibraryUpdateLibraryOptions()
		updateOpts.SetId(plan.Id.ValueString())
		updateOpts.SetLibraryOptions(planOpts)

		_, err := r.data.Client.LibraryStructureServiceAPI.PostLibraryVirtualfoldersLibraryoptions(r.data.Auth).
			LibraryUpdateLibraryOptions(*updateOpts).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Unable to update library options", err.Error())
			return
		}
	}

	// Read back the updated state.
	found := r.findLibraryById(plan.Id.ValueString(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if found == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, found)...)
}

func (r *LibraryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LibraryResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteReq := embyclient.NewLibraryRemoveVirtualFolder()
	deleteReq.SetId(state.Id.ValueString())

	_, err := r.data.Client.LibraryStructureServiceAPI.PostLibraryVirtualfoldersDelete(r.data.Auth).
		LibraryRemoveVirtualFolder(*deleteReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to delete library", err.Error())
		return
	}
}

func (r *LibraryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LibraryResource) findLibraryByName(name string, diags *diag.Diagnostics) *LibraryResourceModel {
	result, _, err := r.data.Client.LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery(r.data.Auth).Execute()
	if err != nil {
		diags.AddError("Unable to read libraries", err.Error())
		return nil
	}

	for _, folder := range result.GetItems() {
		if folder.GetName() == name {
			return mapVirtualFolderToState(&folder)
		}
	}
	return nil
}

func (r *LibraryResource) findLibraryById(id string, diags *diag.Diagnostics) *LibraryResourceModel {
	result, _, err := r.data.Client.LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery(r.data.Auth).Execute()
	if err != nil {
		diags.AddError("Unable to read libraries", err.Error())
		return nil
	}

	for _, folder := range result.GetItems() {
		if folder.GetId() == id {
			diags.AddWarning("DEBUG findLibraryById",
				fmt.Sprintf("Found library id=%s name=%q collectionType=%q libraryOptions=%v",
					folder.GetId(), folder.GetName(), folder.GetCollectionType(), folder.HasLibraryOptions()))
			return mapVirtualFolderToState(&folder)
		}
	}
	return nil
}

func mapVirtualFolderToState(folder *embyclient.VirtualFolderInfo) *LibraryResourceModel {
	locations := make([]types.String, 0, len(folder.GetLocations()))
	paths := make([]types.String, 0, len(folder.GetLocations()))
	for _, loc := range folder.GetLocations() {
		locations = append(locations, types.StringValue(loc))
		paths = append(paths, types.StringValue(loc))
	}

	state := &LibraryResourceModel{
		Id:             types.StringValue(folder.GetId()),
		Name:           types.StringValue(folder.GetName()),
		CollectionType: types.StringValue(folder.GetCollectionType()),
		Paths:          types.ListValueMust(types.StringType, toAttrValues(paths)),
		ItemId:         types.StringValue(folder.GetItemId()),
		Guid:           types.StringValue(folder.GetGuid()),
		Locations:      types.ListValueMust(types.StringType, toAttrValues(locations)),
	}

	if opts, ok := folder.GetLibraryOptionsOk(); ok {
		mapLibraryOptionsToState(opts, state)
	}

	return state
}

func toAttrValues(strs []types.String) []attr.Value {
	vals := make([]attr.Value, len(strs))
	for i, s := range strs {
		vals[i] = s
	}
	return vals
}

// buildLibraryOptions converts Terraform plan fields into an Emby LibraryOptions struct.
// Only fields that are explicitly set (non-null, non-unknown) are included.
func buildLibraryOptions(plan LibraryResourceModel) embyclient.LibraryOptions {
	opts := embyclient.NewLibraryOptions()

	// --- General ---
	if !plan.EnableRealtimeMonitor.IsNull() && !plan.EnableRealtimeMonitor.IsUnknown() {
		opts.SetEnableRealtimeMonitor(plan.EnableRealtimeMonitor.ValueBool())
	}
	if !plan.EnablePhotos.IsNull() && !plan.EnablePhotos.IsUnknown() {
		opts.SetEnablePhotos(plan.EnablePhotos.ValueBool())
	}
	if !plan.EnableArchiveMediaFiles.IsNull() && !plan.EnableArchiveMediaFiles.IsUnknown() {
		opts.SetEnableArchiveMediaFiles(plan.EnableArchiveMediaFiles.ValueBool())
	}
	if !plan.EnableMarkerDetection.IsNull() && !plan.EnableMarkerDetection.IsUnknown() {
		opts.SetEnableMarkerDetection(plan.EnableMarkerDetection.ValueBool())
	}
	if !plan.EnableMarkerDetectionDuringLibraryScan.IsNull() && !plan.EnableMarkerDetectionDuringLibraryScan.IsUnknown() {
		opts.SetEnableMarkerDetectionDuringLibraryScan(plan.EnableMarkerDetectionDuringLibraryScan.ValueBool())
	}
	if !plan.IntroDetectionFingerprintLength.IsNull() && !plan.IntroDetectionFingerprintLength.IsUnknown() {
		opts.SetIntroDetectionFingerprintLength(int32(plan.IntroDetectionFingerprintLength.ValueInt64()))
	}

	// --- Chapter / Image Extraction ---
	if !plan.EnableChapterImageExtraction.IsNull() && !plan.EnableChapterImageExtraction.IsUnknown() {
		opts.SetEnableChapterImageExtraction(plan.EnableChapterImageExtraction.ValueBool())
	}
	if !plan.ExtractChapterImagesDuringLibraryScan.IsNull() && !plan.ExtractChapterImagesDuringLibraryScan.IsUnknown() {
		opts.SetExtractChapterImagesDuringLibraryScan(plan.ExtractChapterImagesDuringLibraryScan.ValueBool())
	}
	if !plan.DownloadImagesInAdvance.IsNull() && !plan.DownloadImagesInAdvance.IsUnknown() {
		opts.SetDownloadImagesInAdvance(plan.DownloadImagesInAdvance.ValueBool())
	}
	if !plan.CacheImages.IsNull() && !plan.CacheImages.IsUnknown() {
		opts.SetCacheImages(plan.CacheImages.ValueBool())
	}

	// --- Metadata ---
	if !plan.PreferredMetadataLanguage.IsNull() && !plan.PreferredMetadataLanguage.IsUnknown() {
		opts.SetPreferredMetadataLanguage(plan.PreferredMetadataLanguage.ValueString())
	}
	if !plan.PreferredImageLanguage.IsNull() && !plan.PreferredImageLanguage.IsUnknown() {
		opts.SetPreferredImageLanguage(plan.PreferredImageLanguage.ValueString())
	}
	if !plan.MetadataCountryCode.IsNull() && !plan.MetadataCountryCode.IsUnknown() {
		opts.SetMetadataCountryCode(plan.MetadataCountryCode.ValueString())
	}
	if !plan.ContentType.IsNull() && !plan.ContentType.IsUnknown() {
		opts.SetContentType(plan.ContentType.ValueString())
	}
	if !plan.SaveLocalMetadata.IsNull() && !plan.SaveLocalMetadata.IsUnknown() {
		opts.SetSaveLocalMetadata(plan.SaveLocalMetadata.ValueBool())
	}
	if !plan.SaveMetadataHidden.IsNull() && !plan.SaveMetadataHidden.IsUnknown() {
		opts.SetSaveMetadataHidden(plan.SaveMetadataHidden.ValueBool())
	}
	if !plan.SaveLocalThumbnailSets.IsNull() && !plan.SaveLocalThumbnailSets.IsUnknown() {
		opts.SetSaveLocalThumbnailSets(plan.SaveLocalThumbnailSets.ValueBool())
	}
	if !plan.EnableAdultMetadata.IsNull() && !plan.EnableAdultMetadata.IsUnknown() {
		opts.SetEnableAdultMetadata(plan.EnableAdultMetadata.ValueBool())
	}

	// --- File Handling ---
	if !plan.IgnoreHiddenFiles.IsNull() && !plan.IgnoreHiddenFiles.IsUnknown() {
		opts.SetIgnoreHiddenFiles(plan.IgnoreHiddenFiles.ValueBool())
	}
	if !plan.EnablePlexIgnore.IsNull() && !plan.EnablePlexIgnore.IsUnknown() {
		opts.SetEnablePlexIgnore(plan.EnablePlexIgnore.ValueBool())
	}
	if !plan.ImportPlaylists.IsNull() && !plan.ImportPlaylists.IsUnknown() {
		opts.SetImportPlaylists(plan.ImportPlaylists.ValueBool())
	}
	if !plan.ImportCollections.IsNull() && !plan.ImportCollections.IsUnknown() {
		opts.SetImportCollections(plan.ImportCollections.ValueBool())
	}

	// --- Subtitles ---
	if !plan.SkipSubtitlesIfEmbeddedSubtitlesPresent.IsNull() && !plan.SkipSubtitlesIfEmbeddedSubtitlesPresent.IsUnknown() {
		opts.SetSkipSubtitlesIfEmbeddedSubtitlesPresent(plan.SkipSubtitlesIfEmbeddedSubtitlesPresent.ValueBool())
	}
	if !plan.SkipSubtitlesIfAudioTrackMatches.IsNull() && !plan.SkipSubtitlesIfAudioTrackMatches.IsUnknown() {
		opts.SetSkipSubtitlesIfAudioTrackMatches(plan.SkipSubtitlesIfAudioTrackMatches.ValueBool())
	}
	if !plan.RequirePerfectSubtitleMatch.IsNull() && !plan.RequirePerfectSubtitleMatch.IsUnknown() {
		opts.SetRequirePerfectSubtitleMatch(plan.RequirePerfectSubtitleMatch.ValueBool())
	}
	if !plan.SaveSubtitlesWithMedia.IsNull() && !plan.SaveSubtitlesWithMedia.IsUnknown() {
		opts.SetSaveSubtitlesWithMedia(plan.SaveSubtitlesWithMedia.ValueBool())
	}
	if !plan.ForcedSubtitlesOnly.IsNull() && !plan.ForcedSubtitlesOnly.IsUnknown() {
		opts.SetForcedSubtitlesOnly(plan.ForcedSubtitlesOnly.ValueBool())
	}
	if !plan.HearingImpairedSubtitlesOnly.IsNull() && !plan.HearingImpairedSubtitlesOnly.IsUnknown() {
		opts.SetHearingImpairedSubtitlesOnly(plan.HearingImpairedSubtitlesOnly.ValueBool())
	}
	if !plan.SubtitleDownloadMaxAgeDays.IsNull() && !plan.SubtitleDownloadMaxAgeDays.IsUnknown() {
		opts.SetSubtitleDownloadMaxAgeDays(int32(plan.SubtitleDownloadMaxAgeDays.ValueInt64()))
	}

	// --- Lyrics ---
	if !plan.SaveLyricsWithMedia.IsNull() && !plan.SaveLyricsWithMedia.IsUnknown() {
		opts.SetSaveLyricsWithMedia(plan.SaveLyricsWithMedia.ValueBool())
	}
	if !plan.LyricsDownloadMaxAgeDays.IsNull() && !plan.LyricsDownloadMaxAgeDays.IsUnknown() {
		opts.SetLyricsDownloadMaxAgeDays(int32(plan.LyricsDownloadMaxAgeDays.ValueInt64()))
	}

	// --- Advanced ---
	if !plan.EnableEmbeddedTitles.IsNull() && !plan.EnableEmbeddedTitles.IsUnknown() {
		opts.SetEnableEmbeddedTitles(plan.EnableEmbeddedTitles.ValueBool())
	}
	if !plan.EnableAudioResume.IsNull() && !plan.EnableAudioResume.IsUnknown() {
		opts.SetEnableAudioResume(plan.EnableAudioResume.ValueBool())
	}
	if !plan.AutoGenerateChapters.IsNull() && !plan.AutoGenerateChapters.IsUnknown() {
		opts.SetAutoGenerateChapters(plan.AutoGenerateChapters.ValueBool())
	}
	if !plan.AutoGenerateChapterIntervalMinutes.IsNull() && !plan.AutoGenerateChapterIntervalMinutes.IsUnknown() {
		opts.SetAutoGenerateChapterIntervalMinutes(int32(plan.AutoGenerateChapterIntervalMinutes.ValueInt64()))
	}
	if !plan.AutomaticRefreshIntervalDays.IsNull() && !plan.AutomaticRefreshIntervalDays.IsUnknown() {
		opts.SetAutomaticRefreshIntervalDays(int32(plan.AutomaticRefreshIntervalDays.ValueInt64()))
	}
	if !plan.PlaceholderMetadataRefreshIntervalDays.IsNull() && !plan.PlaceholderMetadataRefreshIntervalDays.IsUnknown() {
		opts.SetPlaceholderMetadataRefreshIntervalDays(int32(plan.PlaceholderMetadataRefreshIntervalDays.ValueInt64()))
	}
	if !plan.MergeTopLevelFolders.IsNull() && !plan.MergeTopLevelFolders.IsUnknown() {
		opts.SetMergeTopLevelFolders(plan.MergeTopLevelFolders.ValueBool())
	}
	if !plan.CollapseSingleItemFolders.IsNull() && !plan.CollapseSingleItemFolders.IsUnknown() {
		opts.SetCollapseSingleItemFolders(plan.CollapseSingleItemFolders.ValueBool())
	}
	if !plan.ForceCollapseSingleItemFolders.IsNull() && !plan.ForceCollapseSingleItemFolders.IsUnknown() {
		opts.SetForceCollapseSingleItemFolders(plan.ForceCollapseSingleItemFolders.ValueBool())
	}
	if !plan.EnableAutomaticSeriesGrouping.IsNull() && !plan.EnableAutomaticSeriesGrouping.IsUnknown() {
		opts.SetEnableAutomaticSeriesGrouping(plan.EnableAutomaticSeriesGrouping.ValueBool())
	}
	if !plan.ShareEmbeddedMusicAlbumImages.IsNull() && !plan.ShareEmbeddedMusicAlbumImages.IsUnknown() {
		opts.SetShareEmbeddedMusicAlbumImages(plan.ShareEmbeddedMusicAlbumImages.ValueBool())
	}
	if !plan.EnableMultiVersionByFiles.IsNull() && !plan.EnableMultiVersionByFiles.IsUnknown() {
		opts.SetEnableMultiVersionByFiles(plan.EnableMultiVersionByFiles.ValueBool())
	}
	if !plan.EnableMultiVersionByMetadata.IsNull() && !plan.EnableMultiVersionByMetadata.IsUnknown() {
		opts.SetEnableMultiVersionByMetadata(plan.EnableMultiVersionByMetadata.ValueBool())
	}
	if !plan.EnableMultiPartItems.IsNull() && !plan.EnableMultiPartItems.IsUnknown() {
		opts.SetEnableMultiPartItems(plan.EnableMultiPartItems.ValueBool())
	}
	if !plan.MinCollectionItems.IsNull() && !plan.MinCollectionItems.IsUnknown() {
		opts.SetMinCollectionItems(int32(plan.MinCollectionItems.ValueInt64()))
	}
	if !plan.MusicFolderStructure.IsNull() && !plan.MusicFolderStructure.IsUnknown() {
		opts.SetMusicFolderStructure(plan.MusicFolderStructure.ValueString())
	}
	if !plan.ExcludeFromSearch.IsNull() && !plan.ExcludeFromSearch.IsUnknown() {
		opts.SetExcludeFromSearch(plan.ExcludeFromSearch.ValueBool())
	}

	// --- Resume / Playback ---
	if !plan.MinResumePct.IsNull() && !plan.MinResumePct.IsUnknown() {
		opts.SetMinResumePct(int32(plan.MinResumePct.ValueInt64()))
	}
	if !plan.MaxResumePct.IsNull() && !plan.MaxResumePct.IsUnknown() {
		opts.SetMaxResumePct(int32(plan.MaxResumePct.ValueInt64()))
	}
	if !plan.MinResumeDurationSeconds.IsNull() && !plan.MinResumeDurationSeconds.IsUnknown() {
		opts.SetMinResumeDurationSeconds(int32(plan.MinResumeDurationSeconds.ValueInt64()))
	}
	if !plan.ThumbnailImagesIntervalSeconds.IsNull() && !plan.ThumbnailImagesIntervalSeconds.IsUnknown() {
		opts.SetThumbnailImagesIntervalSeconds(int32(plan.ThumbnailImagesIntervalSeconds.ValueInt64()))
	}
	if !plan.SampleIgnoreSize.IsNull() && !plan.SampleIgnoreSize.IsUnknown() {
		opts.SetSampleIgnoreSize(int32(plan.SampleIgnoreSize.ValueInt64()))
	}

	return *opts
}

// mapLibraryOptionsToState copies Emby LibraryOptions fields into the Terraform state model.
func mapLibraryOptionsToState(opts *embyclient.LibraryOptions, state *LibraryResourceModel) {
	// --- General ---
	state.EnableRealtimeMonitor = types.BoolValue(opts.GetEnableRealtimeMonitor())
	state.EnablePhotos = types.BoolValue(opts.GetEnablePhotos())
	state.EnableArchiveMediaFiles = types.BoolValue(opts.GetEnableArchiveMediaFiles())
	state.EnableMarkerDetection = types.BoolValue(opts.GetEnableMarkerDetection())
	state.EnableMarkerDetectionDuringLibraryScan = types.BoolValue(opts.GetEnableMarkerDetectionDuringLibraryScan())
	state.IntroDetectionFingerprintLength = types.Int64Value(int64(opts.GetIntroDetectionFingerprintLength()))

	// --- Chapter / Image Extraction ---
	state.EnableChapterImageExtraction = types.BoolValue(opts.GetEnableChapterImageExtraction())
	state.ExtractChapterImagesDuringLibraryScan = types.BoolValue(opts.GetExtractChapterImagesDuringLibraryScan())
	state.DownloadImagesInAdvance = types.BoolValue(opts.GetDownloadImagesInAdvance())
	state.CacheImages = types.BoolValue(opts.GetCacheImages())

	// --- Metadata ---
	state.PreferredMetadataLanguage = types.StringValue(opts.GetPreferredMetadataLanguage())
	state.PreferredImageLanguage = types.StringValue(opts.GetPreferredImageLanguage())
	state.MetadataCountryCode = types.StringValue(opts.GetMetadataCountryCode())
	state.ContentType = types.StringValue(opts.GetContentType())
	state.MetadataSavers = stringSliceToSetValue(opts.GetMetadataSavers())
	state.DisabledLocalMetadataReaders = stringSliceToSetValue(opts.GetDisabledLocalMetadataReaders())
	state.LocalMetadataReaderOrder = stringSliceToListValue(opts.GetLocalMetadataReaderOrder())
	state.SaveLocalMetadata = types.BoolValue(opts.GetSaveLocalMetadata())
	state.SaveMetadataHidden = types.BoolValue(opts.GetSaveMetadataHidden())
	state.SaveLocalThumbnailSets = types.BoolValue(opts.GetSaveLocalThumbnailSets())
	state.EnableAdultMetadata = types.BoolValue(opts.GetEnableAdultMetadata())

	// --- File Handling ---
	state.IgnoreHiddenFiles = types.BoolValue(opts.GetIgnoreHiddenFiles())
	state.IgnoreFileExtensions = stringSliceToSetValue(opts.GetIgnoreFileExtensions())
	state.EnablePlexIgnore = types.BoolValue(opts.GetEnablePlexIgnore())
	state.ImportPlaylists = types.BoolValue(opts.GetImportPlaylists())
	state.ImportCollections = types.BoolValue(opts.GetImportCollections())

	// --- Subtitles ---
	state.DisabledSubtitleFetchers = stringSliceToSetValue(opts.GetDisabledSubtitleFetchers())
	state.SubtitleFetcherOrder = stringSliceToListValue(opts.GetSubtitleFetcherOrder())
	state.SubtitleDownloadLanguages = stringSliceToSetValue(opts.GetSubtitleDownloadLanguages())
	state.SubtitleDownloadMaxAgeDays = types.Int64Value(int64(opts.GetSubtitleDownloadMaxAgeDays()))
	state.SkipSubtitlesIfEmbeddedSubtitlesPresent = types.BoolValue(opts.GetSkipSubtitlesIfEmbeddedSubtitlesPresent())
	state.SkipSubtitlesIfAudioTrackMatches = types.BoolValue(opts.GetSkipSubtitlesIfAudioTrackMatches())
	state.RequirePerfectSubtitleMatch = types.BoolValue(opts.GetRequirePerfectSubtitleMatch())
	state.SaveSubtitlesWithMedia = types.BoolValue(opts.GetSaveSubtitlesWithMedia())
	state.ForcedSubtitlesOnly = types.BoolValue(opts.GetForcedSubtitlesOnly())
	state.HearingImpairedSubtitlesOnly = types.BoolValue(opts.GetHearingImpairedSubtitlesOnly())

	// --- Lyrics ---
	state.DisabledLyricsFetchers = stringSliceToSetValue(opts.GetDisabledLyricsFetchers())
	state.LyricsFetcherOrder = stringSliceToListValue(opts.GetLyricsFetcherOrder())
	state.LyricsDownloadLanguages = stringSliceToSetValue(opts.GetLyricsDownloadLanguages())
	state.LyricsDownloadMaxAgeDays = types.Int64Value(int64(opts.GetLyricsDownloadMaxAgeDays()))
	state.SaveLyricsWithMedia = types.BoolValue(opts.GetSaveLyricsWithMedia())

	// --- Advanced ---
	state.EnableEmbeddedTitles = types.BoolValue(opts.GetEnableEmbeddedTitles())
	state.EnableAudioResume = types.BoolValue(opts.GetEnableAudioResume())
	state.AutoGenerateChapters = types.BoolValue(opts.GetAutoGenerateChapters())
	state.AutoGenerateChapterIntervalMinutes = types.Int64Value(int64(opts.GetAutoGenerateChapterIntervalMinutes()))
	state.AutomaticRefreshIntervalDays = types.Int64Value(int64(opts.GetAutomaticRefreshIntervalDays()))
	state.PlaceholderMetadataRefreshIntervalDays = types.Int64Value(int64(opts.GetPlaceholderMetadataRefreshIntervalDays()))
	state.MergeTopLevelFolders = types.BoolValue(opts.GetMergeTopLevelFolders())
	state.CollapseSingleItemFolders = types.BoolValue(opts.GetCollapseSingleItemFolders())
	state.ForceCollapseSingleItemFolders = types.BoolValue(opts.GetForceCollapseSingleItemFolders())
	state.EnableAutomaticSeriesGrouping = types.BoolValue(opts.GetEnableAutomaticSeriesGrouping())
	state.ShareEmbeddedMusicAlbumImages = types.BoolValue(opts.GetShareEmbeddedMusicAlbumImages())
	state.EnableMultiVersionByFiles = types.BoolValue(opts.GetEnableMultiVersionByFiles())
	state.EnableMultiVersionByMetadata = types.BoolValue(opts.GetEnableMultiVersionByMetadata())
	state.EnableMultiPartItems = types.BoolValue(opts.GetEnableMultiPartItems())
	state.MinCollectionItems = types.Int64Value(int64(opts.GetMinCollectionItems()))
	state.MusicFolderStructure = types.StringValue(opts.GetMusicFolderStructure())
	state.ExcludeFromSearch = types.BoolValue(opts.GetExcludeFromSearch())

	// --- Resume / Playback ---
	state.MinResumePct = types.Int64Value(int64(opts.GetMinResumePct()))
	state.MaxResumePct = types.Int64Value(int64(opts.GetMaxResumePct()))
	state.MinResumeDurationSeconds = types.Int64Value(int64(opts.GetMinResumeDurationSeconds()))
	state.ThumbnailImagesIntervalSeconds = types.Int64Value(int64(opts.GetThumbnailImagesIntervalSeconds()))
	state.SampleIgnoreSize = types.Int64Value(int64(opts.GetSampleIgnoreSize()))
}

// libraryOptionsEqual returns true if both LibraryOptions serialize to the same JSON.
func libraryOptionsEqual(a, b embyclient.LibraryOptions) bool {
	ja, _ := a.MarshalJSON()
	jb, _ := b.MarshalJSON()
	return string(ja) == string(jb)
}

// libraryOptionsEmpty returns true if no fields are set in the LibraryOptions.
func libraryOptionsEmpty(opts embyclient.LibraryOptions) bool {
	j, _ := opts.MarshalJSON()
	return string(j) == "{}"
}
