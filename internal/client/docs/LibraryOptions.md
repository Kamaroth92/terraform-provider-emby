# LibraryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableArchiveMediaFiles** | Pointer to **bool** |  | [optional] 
**EnablePhotos** | Pointer to **bool** |  | [optional] 
**EnableRealtimeMonitor** | Pointer to **bool** |  | [optional] 
**EnableMarkerDetection** | Pointer to **bool** |  | [optional] 
**EnableMarkerDetectionDuringLibraryScan** | Pointer to **bool** |  | [optional] 
**IntroDetectionFingerprintLength** | Pointer to **int32** |  | [optional] 
**EnableChapterImageExtraction** | Pointer to **bool** |  | [optional] 
**ExtractChapterImagesDuringLibraryScan** | Pointer to **bool** |  | [optional] 
**DownloadImagesInAdvance** | Pointer to **bool** |  | [optional] 
**CacheImages** | Pointer to **bool** |  | [optional] 
**ExcludeFromSearch** | Pointer to **bool** |  | [optional] 
**EnablePlexIgnore** | Pointer to **bool** |  | [optional] 
**PathInfos** | Pointer to [**[]MediaPathInfo**](MediaPathInfo.md) |  | [optional] 
**IgnoreHiddenFiles** | Pointer to **bool** |  | [optional] 
**IgnoreFileExtensions** | Pointer to **[]string** |  | [optional] 
**SaveLocalMetadata** | Pointer to **bool** |  | [optional] 
**SaveMetadataHidden** | Pointer to **bool** |  | [optional] 
**SaveLocalThumbnailSets** | Pointer to **bool** |  | [optional] 
**ImportPlaylists** | Pointer to **bool** |  | [optional] 
**EnableAutomaticSeriesGrouping** | Pointer to **bool** |  | [optional] 
**ShareEmbeddedMusicAlbumImages** | Pointer to **bool** |  | [optional] 
**EnableEmbeddedTitles** | Pointer to **bool** |  | [optional] 
**EnableAudioResume** | Pointer to **bool** |  | [optional] 
**AutoGenerateChapters** | Pointer to **bool** |  | [optional] 
**MergeTopLevelFolders** | Pointer to **bool** |  | [optional] 
**AutoGenerateChapterIntervalMinutes** | Pointer to **int32** |  | [optional] 
**AutomaticRefreshIntervalDays** | Pointer to **int32** |  | [optional] 
**PlaceholderMetadataRefreshIntervalDays** | Pointer to **int32** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **string** |  | [optional] 
**PreferredImageLanguage** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**MetadataCountryCode** | Pointer to **string** |  | [optional] 
**MetadataSavers** | Pointer to **[]string** |  | [optional] 
**DisabledLocalMetadataReaders** | Pointer to **[]string** |  | [optional] 
**LocalMetadataReaderOrder** | Pointer to **[]string** |  | [optional] 
**DisabledLyricsFetchers** | Pointer to **[]string** |  | [optional] 
**SaveLyricsWithMedia** | Pointer to **bool** |  | [optional] 
**LyricsDownloadMaxAgeDays** | Pointer to **int32** |  | [optional] 
**LyricsFetcherOrder** | Pointer to **[]string** |  | [optional] 
**LyricsDownloadLanguages** | Pointer to **[]string** |  | [optional] 
**DisabledSubtitleFetchers** | Pointer to **[]string** |  | [optional] 
**SubtitleFetcherOrder** | Pointer to **[]string** |  | [optional] 
**SkipSubtitlesIfEmbeddedSubtitlesPresent** | Pointer to **bool** |  | [optional] 
**SkipSubtitlesIfAudioTrackMatches** | Pointer to **bool** |  | [optional] 
**SubtitleDownloadLanguages** | Pointer to **[]string** |  | [optional] 
**SubtitleDownloadMaxAgeDays** | Pointer to **int32** |  | [optional] 
**RequirePerfectSubtitleMatch** | Pointer to **bool** |  | [optional] 
**SaveSubtitlesWithMedia** | Pointer to **bool** |  | [optional] 
**ForcedSubtitlesOnly** | Pointer to **bool** |  | [optional] 
**HearingImpairedSubtitlesOnly** | Pointer to **bool** |  | [optional] 
**TypeOptions** | Pointer to [**[]TypeOptions**](TypeOptions.md) |  | [optional] 
**CollapseSingleItemFolders** | Pointer to **bool** |  | [optional] 
**ForceCollapseSingleItemFolders** | Pointer to **bool** |  | [optional] 
**EnableAdultMetadata** | Pointer to **bool** |  | [optional] 
**ImportCollections** | Pointer to **bool** |  | [optional] 
**EnableMultiVersionByFiles** | Pointer to **bool** |  | [optional] 
**EnableMultiVersionByMetadata** | Pointer to **bool** |  | [optional] 
**EnableMultiPartItems** | Pointer to **bool** |  | [optional] 
**MinCollectionItems** | Pointer to **int32** |  | [optional] 
**MusicFolderStructure** | Pointer to **string** |  | [optional] 
**MinResumePct** | Pointer to **int32** |  | [optional] 
**MaxResumePct** | Pointer to **int32** |  | [optional] 
**MinResumeDurationSeconds** | Pointer to **int32** |  | [optional] 
**ThumbnailImagesIntervalSeconds** | Pointer to **int32** |  | [optional] 
**SampleIgnoreSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewLibraryOptions

`func NewLibraryOptions() *LibraryOptions`

NewLibraryOptions instantiates a new LibraryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryOptionsWithDefaults

`func NewLibraryOptionsWithDefaults() *LibraryOptions`

NewLibraryOptionsWithDefaults instantiates a new LibraryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableArchiveMediaFiles

`func (o *LibraryOptions) GetEnableArchiveMediaFiles() bool`

GetEnableArchiveMediaFiles returns the EnableArchiveMediaFiles field if non-nil, zero value otherwise.

### GetEnableArchiveMediaFilesOk

`func (o *LibraryOptions) GetEnableArchiveMediaFilesOk() (*bool, bool)`

GetEnableArchiveMediaFilesOk returns a tuple with the EnableArchiveMediaFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableArchiveMediaFiles

`func (o *LibraryOptions) SetEnableArchiveMediaFiles(v bool)`

SetEnableArchiveMediaFiles sets EnableArchiveMediaFiles field to given value.

### HasEnableArchiveMediaFiles

`func (o *LibraryOptions) HasEnableArchiveMediaFiles() bool`

HasEnableArchiveMediaFiles returns a boolean if a field has been set.

### GetEnablePhotos

`func (o *LibraryOptions) GetEnablePhotos() bool`

GetEnablePhotos returns the EnablePhotos field if non-nil, zero value otherwise.

### GetEnablePhotosOk

`func (o *LibraryOptions) GetEnablePhotosOk() (*bool, bool)`

GetEnablePhotosOk returns a tuple with the EnablePhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePhotos

`func (o *LibraryOptions) SetEnablePhotos(v bool)`

SetEnablePhotos sets EnablePhotos field to given value.

### HasEnablePhotos

`func (o *LibraryOptions) HasEnablePhotos() bool`

HasEnablePhotos returns a boolean if a field has been set.

### GetEnableRealtimeMonitor

`func (o *LibraryOptions) GetEnableRealtimeMonitor() bool`

GetEnableRealtimeMonitor returns the EnableRealtimeMonitor field if non-nil, zero value otherwise.

### GetEnableRealtimeMonitorOk

`func (o *LibraryOptions) GetEnableRealtimeMonitorOk() (*bool, bool)`

GetEnableRealtimeMonitorOk returns a tuple with the EnableRealtimeMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRealtimeMonitor

`func (o *LibraryOptions) SetEnableRealtimeMonitor(v bool)`

SetEnableRealtimeMonitor sets EnableRealtimeMonitor field to given value.

### HasEnableRealtimeMonitor

`func (o *LibraryOptions) HasEnableRealtimeMonitor() bool`

HasEnableRealtimeMonitor returns a boolean if a field has been set.

### GetEnableMarkerDetection

`func (o *LibraryOptions) GetEnableMarkerDetection() bool`

GetEnableMarkerDetection returns the EnableMarkerDetection field if non-nil, zero value otherwise.

### GetEnableMarkerDetectionOk

`func (o *LibraryOptions) GetEnableMarkerDetectionOk() (*bool, bool)`

GetEnableMarkerDetectionOk returns a tuple with the EnableMarkerDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMarkerDetection

`func (o *LibraryOptions) SetEnableMarkerDetection(v bool)`

SetEnableMarkerDetection sets EnableMarkerDetection field to given value.

### HasEnableMarkerDetection

`func (o *LibraryOptions) HasEnableMarkerDetection() bool`

HasEnableMarkerDetection returns a boolean if a field has been set.

### GetEnableMarkerDetectionDuringLibraryScan

`func (o *LibraryOptions) GetEnableMarkerDetectionDuringLibraryScan() bool`

GetEnableMarkerDetectionDuringLibraryScan returns the EnableMarkerDetectionDuringLibraryScan field if non-nil, zero value otherwise.

### GetEnableMarkerDetectionDuringLibraryScanOk

`func (o *LibraryOptions) GetEnableMarkerDetectionDuringLibraryScanOk() (*bool, bool)`

GetEnableMarkerDetectionDuringLibraryScanOk returns a tuple with the EnableMarkerDetectionDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMarkerDetectionDuringLibraryScan

`func (o *LibraryOptions) SetEnableMarkerDetectionDuringLibraryScan(v bool)`

SetEnableMarkerDetectionDuringLibraryScan sets EnableMarkerDetectionDuringLibraryScan field to given value.

### HasEnableMarkerDetectionDuringLibraryScan

`func (o *LibraryOptions) HasEnableMarkerDetectionDuringLibraryScan() bool`

HasEnableMarkerDetectionDuringLibraryScan returns a boolean if a field has been set.

### GetIntroDetectionFingerprintLength

`func (o *LibraryOptions) GetIntroDetectionFingerprintLength() int32`

GetIntroDetectionFingerprintLength returns the IntroDetectionFingerprintLength field if non-nil, zero value otherwise.

### GetIntroDetectionFingerprintLengthOk

`func (o *LibraryOptions) GetIntroDetectionFingerprintLengthOk() (*int32, bool)`

GetIntroDetectionFingerprintLengthOk returns a tuple with the IntroDetectionFingerprintLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroDetectionFingerprintLength

`func (o *LibraryOptions) SetIntroDetectionFingerprintLength(v int32)`

SetIntroDetectionFingerprintLength sets IntroDetectionFingerprintLength field to given value.

### HasIntroDetectionFingerprintLength

`func (o *LibraryOptions) HasIntroDetectionFingerprintLength() bool`

HasIntroDetectionFingerprintLength returns a boolean if a field has been set.

### GetEnableChapterImageExtraction

`func (o *LibraryOptions) GetEnableChapterImageExtraction() bool`

GetEnableChapterImageExtraction returns the EnableChapterImageExtraction field if non-nil, zero value otherwise.

### GetEnableChapterImageExtractionOk

`func (o *LibraryOptions) GetEnableChapterImageExtractionOk() (*bool, bool)`

GetEnableChapterImageExtractionOk returns a tuple with the EnableChapterImageExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChapterImageExtraction

`func (o *LibraryOptions) SetEnableChapterImageExtraction(v bool)`

SetEnableChapterImageExtraction sets EnableChapterImageExtraction field to given value.

### HasEnableChapterImageExtraction

`func (o *LibraryOptions) HasEnableChapterImageExtraction() bool`

HasEnableChapterImageExtraction returns a boolean if a field has been set.

### GetExtractChapterImagesDuringLibraryScan

`func (o *LibraryOptions) GetExtractChapterImagesDuringLibraryScan() bool`

GetExtractChapterImagesDuringLibraryScan returns the ExtractChapterImagesDuringLibraryScan field if non-nil, zero value otherwise.

### GetExtractChapterImagesDuringLibraryScanOk

`func (o *LibraryOptions) GetExtractChapterImagesDuringLibraryScanOk() (*bool, bool)`

GetExtractChapterImagesDuringLibraryScanOk returns a tuple with the ExtractChapterImagesDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractChapterImagesDuringLibraryScan

`func (o *LibraryOptions) SetExtractChapterImagesDuringLibraryScan(v bool)`

SetExtractChapterImagesDuringLibraryScan sets ExtractChapterImagesDuringLibraryScan field to given value.

### HasExtractChapterImagesDuringLibraryScan

`func (o *LibraryOptions) HasExtractChapterImagesDuringLibraryScan() bool`

HasExtractChapterImagesDuringLibraryScan returns a boolean if a field has been set.

### GetDownloadImagesInAdvance

`func (o *LibraryOptions) GetDownloadImagesInAdvance() bool`

GetDownloadImagesInAdvance returns the DownloadImagesInAdvance field if non-nil, zero value otherwise.

### GetDownloadImagesInAdvanceOk

`func (o *LibraryOptions) GetDownloadImagesInAdvanceOk() (*bool, bool)`

GetDownloadImagesInAdvanceOk returns a tuple with the DownloadImagesInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadImagesInAdvance

`func (o *LibraryOptions) SetDownloadImagesInAdvance(v bool)`

SetDownloadImagesInAdvance sets DownloadImagesInAdvance field to given value.

### HasDownloadImagesInAdvance

`func (o *LibraryOptions) HasDownloadImagesInAdvance() bool`

HasDownloadImagesInAdvance returns a boolean if a field has been set.

### GetCacheImages

`func (o *LibraryOptions) GetCacheImages() bool`

GetCacheImages returns the CacheImages field if non-nil, zero value otherwise.

### GetCacheImagesOk

`func (o *LibraryOptions) GetCacheImagesOk() (*bool, bool)`

GetCacheImagesOk returns a tuple with the CacheImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheImages

`func (o *LibraryOptions) SetCacheImages(v bool)`

SetCacheImages sets CacheImages field to given value.

### HasCacheImages

`func (o *LibraryOptions) HasCacheImages() bool`

HasCacheImages returns a boolean if a field has been set.

### GetExcludeFromSearch

`func (o *LibraryOptions) GetExcludeFromSearch() bool`

GetExcludeFromSearch returns the ExcludeFromSearch field if non-nil, zero value otherwise.

### GetExcludeFromSearchOk

`func (o *LibraryOptions) GetExcludeFromSearchOk() (*bool, bool)`

GetExcludeFromSearchOk returns a tuple with the ExcludeFromSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromSearch

`func (o *LibraryOptions) SetExcludeFromSearch(v bool)`

SetExcludeFromSearch sets ExcludeFromSearch field to given value.

### HasExcludeFromSearch

`func (o *LibraryOptions) HasExcludeFromSearch() bool`

HasExcludeFromSearch returns a boolean if a field has been set.

### GetEnablePlexIgnore

`func (o *LibraryOptions) GetEnablePlexIgnore() bool`

GetEnablePlexIgnore returns the EnablePlexIgnore field if non-nil, zero value otherwise.

### GetEnablePlexIgnoreOk

`func (o *LibraryOptions) GetEnablePlexIgnoreOk() (*bool, bool)`

GetEnablePlexIgnoreOk returns a tuple with the EnablePlexIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePlexIgnore

`func (o *LibraryOptions) SetEnablePlexIgnore(v bool)`

SetEnablePlexIgnore sets EnablePlexIgnore field to given value.

### HasEnablePlexIgnore

`func (o *LibraryOptions) HasEnablePlexIgnore() bool`

HasEnablePlexIgnore returns a boolean if a field has been set.

### GetPathInfos

`func (o *LibraryOptions) GetPathInfos() []MediaPathInfo`

GetPathInfos returns the PathInfos field if non-nil, zero value otherwise.

### GetPathInfosOk

`func (o *LibraryOptions) GetPathInfosOk() (*[]MediaPathInfo, bool)`

GetPathInfosOk returns a tuple with the PathInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathInfos

`func (o *LibraryOptions) SetPathInfos(v []MediaPathInfo)`

SetPathInfos sets PathInfos field to given value.

### HasPathInfos

`func (o *LibraryOptions) HasPathInfos() bool`

HasPathInfos returns a boolean if a field has been set.

### GetIgnoreHiddenFiles

`func (o *LibraryOptions) GetIgnoreHiddenFiles() bool`

GetIgnoreHiddenFiles returns the IgnoreHiddenFiles field if non-nil, zero value otherwise.

### GetIgnoreHiddenFilesOk

`func (o *LibraryOptions) GetIgnoreHiddenFilesOk() (*bool, bool)`

GetIgnoreHiddenFilesOk returns a tuple with the IgnoreHiddenFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreHiddenFiles

`func (o *LibraryOptions) SetIgnoreHiddenFiles(v bool)`

SetIgnoreHiddenFiles sets IgnoreHiddenFiles field to given value.

### HasIgnoreHiddenFiles

`func (o *LibraryOptions) HasIgnoreHiddenFiles() bool`

HasIgnoreHiddenFiles returns a boolean if a field has been set.

### GetIgnoreFileExtensions

`func (o *LibraryOptions) GetIgnoreFileExtensions() []string`

GetIgnoreFileExtensions returns the IgnoreFileExtensions field if non-nil, zero value otherwise.

### GetIgnoreFileExtensionsOk

`func (o *LibraryOptions) GetIgnoreFileExtensionsOk() (*[]string, bool)`

GetIgnoreFileExtensionsOk returns a tuple with the IgnoreFileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreFileExtensions

`func (o *LibraryOptions) SetIgnoreFileExtensions(v []string)`

SetIgnoreFileExtensions sets IgnoreFileExtensions field to given value.

### HasIgnoreFileExtensions

`func (o *LibraryOptions) HasIgnoreFileExtensions() bool`

HasIgnoreFileExtensions returns a boolean if a field has been set.

### GetSaveLocalMetadata

`func (o *LibraryOptions) GetSaveLocalMetadata() bool`

GetSaveLocalMetadata returns the SaveLocalMetadata field if non-nil, zero value otherwise.

### GetSaveLocalMetadataOk

`func (o *LibraryOptions) GetSaveLocalMetadataOk() (*bool, bool)`

GetSaveLocalMetadataOk returns a tuple with the SaveLocalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLocalMetadata

`func (o *LibraryOptions) SetSaveLocalMetadata(v bool)`

SetSaveLocalMetadata sets SaveLocalMetadata field to given value.

### HasSaveLocalMetadata

`func (o *LibraryOptions) HasSaveLocalMetadata() bool`

HasSaveLocalMetadata returns a boolean if a field has been set.

### GetSaveMetadataHidden

`func (o *LibraryOptions) GetSaveMetadataHidden() bool`

GetSaveMetadataHidden returns the SaveMetadataHidden field if non-nil, zero value otherwise.

### GetSaveMetadataHiddenOk

`func (o *LibraryOptions) GetSaveMetadataHiddenOk() (*bool, bool)`

GetSaveMetadataHiddenOk returns a tuple with the SaveMetadataHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveMetadataHidden

`func (o *LibraryOptions) SetSaveMetadataHidden(v bool)`

SetSaveMetadataHidden sets SaveMetadataHidden field to given value.

### HasSaveMetadataHidden

`func (o *LibraryOptions) HasSaveMetadataHidden() bool`

HasSaveMetadataHidden returns a boolean if a field has been set.

### GetSaveLocalThumbnailSets

`func (o *LibraryOptions) GetSaveLocalThumbnailSets() bool`

GetSaveLocalThumbnailSets returns the SaveLocalThumbnailSets field if non-nil, zero value otherwise.

### GetSaveLocalThumbnailSetsOk

`func (o *LibraryOptions) GetSaveLocalThumbnailSetsOk() (*bool, bool)`

GetSaveLocalThumbnailSetsOk returns a tuple with the SaveLocalThumbnailSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLocalThumbnailSets

`func (o *LibraryOptions) SetSaveLocalThumbnailSets(v bool)`

SetSaveLocalThumbnailSets sets SaveLocalThumbnailSets field to given value.

### HasSaveLocalThumbnailSets

`func (o *LibraryOptions) HasSaveLocalThumbnailSets() bool`

HasSaveLocalThumbnailSets returns a boolean if a field has been set.

### GetImportPlaylists

`func (o *LibraryOptions) GetImportPlaylists() bool`

GetImportPlaylists returns the ImportPlaylists field if non-nil, zero value otherwise.

### GetImportPlaylistsOk

`func (o *LibraryOptions) GetImportPlaylistsOk() (*bool, bool)`

GetImportPlaylistsOk returns a tuple with the ImportPlaylists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPlaylists

`func (o *LibraryOptions) SetImportPlaylists(v bool)`

SetImportPlaylists sets ImportPlaylists field to given value.

### HasImportPlaylists

`func (o *LibraryOptions) HasImportPlaylists() bool`

HasImportPlaylists returns a boolean if a field has been set.

### GetEnableAutomaticSeriesGrouping

`func (o *LibraryOptions) GetEnableAutomaticSeriesGrouping() bool`

GetEnableAutomaticSeriesGrouping returns the EnableAutomaticSeriesGrouping field if non-nil, zero value otherwise.

### GetEnableAutomaticSeriesGroupingOk

`func (o *LibraryOptions) GetEnableAutomaticSeriesGroupingOk() (*bool, bool)`

GetEnableAutomaticSeriesGroupingOk returns a tuple with the EnableAutomaticSeriesGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticSeriesGrouping

`func (o *LibraryOptions) SetEnableAutomaticSeriesGrouping(v bool)`

SetEnableAutomaticSeriesGrouping sets EnableAutomaticSeriesGrouping field to given value.

### HasEnableAutomaticSeriesGrouping

`func (o *LibraryOptions) HasEnableAutomaticSeriesGrouping() bool`

HasEnableAutomaticSeriesGrouping returns a boolean if a field has been set.

### GetShareEmbeddedMusicAlbumImages

`func (o *LibraryOptions) GetShareEmbeddedMusicAlbumImages() bool`

GetShareEmbeddedMusicAlbumImages returns the ShareEmbeddedMusicAlbumImages field if non-nil, zero value otherwise.

### GetShareEmbeddedMusicAlbumImagesOk

`func (o *LibraryOptions) GetShareEmbeddedMusicAlbumImagesOk() (*bool, bool)`

GetShareEmbeddedMusicAlbumImagesOk returns a tuple with the ShareEmbeddedMusicAlbumImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareEmbeddedMusicAlbumImages

`func (o *LibraryOptions) SetShareEmbeddedMusicAlbumImages(v bool)`

SetShareEmbeddedMusicAlbumImages sets ShareEmbeddedMusicAlbumImages field to given value.

### HasShareEmbeddedMusicAlbumImages

`func (o *LibraryOptions) HasShareEmbeddedMusicAlbumImages() bool`

HasShareEmbeddedMusicAlbumImages returns a boolean if a field has been set.

### GetEnableEmbeddedTitles

`func (o *LibraryOptions) GetEnableEmbeddedTitles() bool`

GetEnableEmbeddedTitles returns the EnableEmbeddedTitles field if non-nil, zero value otherwise.

### GetEnableEmbeddedTitlesOk

`func (o *LibraryOptions) GetEnableEmbeddedTitlesOk() (*bool, bool)`

GetEnableEmbeddedTitlesOk returns a tuple with the EnableEmbeddedTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedTitles

`func (o *LibraryOptions) SetEnableEmbeddedTitles(v bool)`

SetEnableEmbeddedTitles sets EnableEmbeddedTitles field to given value.

### HasEnableEmbeddedTitles

`func (o *LibraryOptions) HasEnableEmbeddedTitles() bool`

HasEnableEmbeddedTitles returns a boolean if a field has been set.

### GetEnableAudioResume

`func (o *LibraryOptions) GetEnableAudioResume() bool`

GetEnableAudioResume returns the EnableAudioResume field if non-nil, zero value otherwise.

### GetEnableAudioResumeOk

`func (o *LibraryOptions) GetEnableAudioResumeOk() (*bool, bool)`

GetEnableAudioResumeOk returns a tuple with the EnableAudioResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioResume

`func (o *LibraryOptions) SetEnableAudioResume(v bool)`

SetEnableAudioResume sets EnableAudioResume field to given value.

### HasEnableAudioResume

`func (o *LibraryOptions) HasEnableAudioResume() bool`

HasEnableAudioResume returns a boolean if a field has been set.

### GetAutoGenerateChapters

`func (o *LibraryOptions) GetAutoGenerateChapters() bool`

GetAutoGenerateChapters returns the AutoGenerateChapters field if non-nil, zero value otherwise.

### GetAutoGenerateChaptersOk

`func (o *LibraryOptions) GetAutoGenerateChaptersOk() (*bool, bool)`

GetAutoGenerateChaptersOk returns a tuple with the AutoGenerateChapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoGenerateChapters

`func (o *LibraryOptions) SetAutoGenerateChapters(v bool)`

SetAutoGenerateChapters sets AutoGenerateChapters field to given value.

### HasAutoGenerateChapters

`func (o *LibraryOptions) HasAutoGenerateChapters() bool`

HasAutoGenerateChapters returns a boolean if a field has been set.

### GetMergeTopLevelFolders

`func (o *LibraryOptions) GetMergeTopLevelFolders() bool`

GetMergeTopLevelFolders returns the MergeTopLevelFolders field if non-nil, zero value otherwise.

### GetMergeTopLevelFoldersOk

`func (o *LibraryOptions) GetMergeTopLevelFoldersOk() (*bool, bool)`

GetMergeTopLevelFoldersOk returns a tuple with the MergeTopLevelFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeTopLevelFolders

`func (o *LibraryOptions) SetMergeTopLevelFolders(v bool)`

SetMergeTopLevelFolders sets MergeTopLevelFolders field to given value.

### HasMergeTopLevelFolders

`func (o *LibraryOptions) HasMergeTopLevelFolders() bool`

HasMergeTopLevelFolders returns a boolean if a field has been set.

### GetAutoGenerateChapterIntervalMinutes

`func (o *LibraryOptions) GetAutoGenerateChapterIntervalMinutes() int32`

GetAutoGenerateChapterIntervalMinutes returns the AutoGenerateChapterIntervalMinutes field if non-nil, zero value otherwise.

### GetAutoGenerateChapterIntervalMinutesOk

`func (o *LibraryOptions) GetAutoGenerateChapterIntervalMinutesOk() (*int32, bool)`

GetAutoGenerateChapterIntervalMinutesOk returns a tuple with the AutoGenerateChapterIntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoGenerateChapterIntervalMinutes

`func (o *LibraryOptions) SetAutoGenerateChapterIntervalMinutes(v int32)`

SetAutoGenerateChapterIntervalMinutes sets AutoGenerateChapterIntervalMinutes field to given value.

### HasAutoGenerateChapterIntervalMinutes

`func (o *LibraryOptions) HasAutoGenerateChapterIntervalMinutes() bool`

HasAutoGenerateChapterIntervalMinutes returns a boolean if a field has been set.

### GetAutomaticRefreshIntervalDays

`func (o *LibraryOptions) GetAutomaticRefreshIntervalDays() int32`

GetAutomaticRefreshIntervalDays returns the AutomaticRefreshIntervalDays field if non-nil, zero value otherwise.

### GetAutomaticRefreshIntervalDaysOk

`func (o *LibraryOptions) GetAutomaticRefreshIntervalDaysOk() (*int32, bool)`

GetAutomaticRefreshIntervalDaysOk returns a tuple with the AutomaticRefreshIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticRefreshIntervalDays

`func (o *LibraryOptions) SetAutomaticRefreshIntervalDays(v int32)`

SetAutomaticRefreshIntervalDays sets AutomaticRefreshIntervalDays field to given value.

### HasAutomaticRefreshIntervalDays

`func (o *LibraryOptions) HasAutomaticRefreshIntervalDays() bool`

HasAutomaticRefreshIntervalDays returns a boolean if a field has been set.

### GetPlaceholderMetadataRefreshIntervalDays

`func (o *LibraryOptions) GetPlaceholderMetadataRefreshIntervalDays() int32`

GetPlaceholderMetadataRefreshIntervalDays returns the PlaceholderMetadataRefreshIntervalDays field if non-nil, zero value otherwise.

### GetPlaceholderMetadataRefreshIntervalDaysOk

`func (o *LibraryOptions) GetPlaceholderMetadataRefreshIntervalDaysOk() (*int32, bool)`

GetPlaceholderMetadataRefreshIntervalDaysOk returns a tuple with the PlaceholderMetadataRefreshIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholderMetadataRefreshIntervalDays

`func (o *LibraryOptions) SetPlaceholderMetadataRefreshIntervalDays(v int32)`

SetPlaceholderMetadataRefreshIntervalDays sets PlaceholderMetadataRefreshIntervalDays field to given value.

### HasPlaceholderMetadataRefreshIntervalDays

`func (o *LibraryOptions) HasPlaceholderMetadataRefreshIntervalDays() bool`

HasPlaceholderMetadataRefreshIntervalDays returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *LibraryOptions) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *LibraryOptions) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *LibraryOptions) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *LibraryOptions) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### GetPreferredImageLanguage

`func (o *LibraryOptions) GetPreferredImageLanguage() string`

GetPreferredImageLanguage returns the PreferredImageLanguage field if non-nil, zero value otherwise.

### GetPreferredImageLanguageOk

`func (o *LibraryOptions) GetPreferredImageLanguageOk() (*string, bool)`

GetPreferredImageLanguageOk returns a tuple with the PreferredImageLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredImageLanguage

`func (o *LibraryOptions) SetPreferredImageLanguage(v string)`

SetPreferredImageLanguage sets PreferredImageLanguage field to given value.

### HasPreferredImageLanguage

`func (o *LibraryOptions) HasPreferredImageLanguage() bool`

HasPreferredImageLanguage returns a boolean if a field has been set.

### GetContentType

`func (o *LibraryOptions) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LibraryOptions) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LibraryOptions) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *LibraryOptions) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetMetadataCountryCode

`func (o *LibraryOptions) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *LibraryOptions) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *LibraryOptions) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *LibraryOptions) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### GetMetadataSavers

`func (o *LibraryOptions) GetMetadataSavers() []string`

GetMetadataSavers returns the MetadataSavers field if non-nil, zero value otherwise.

### GetMetadataSaversOk

`func (o *LibraryOptions) GetMetadataSaversOk() (*[]string, bool)`

GetMetadataSaversOk returns a tuple with the MetadataSavers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSavers

`func (o *LibraryOptions) SetMetadataSavers(v []string)`

SetMetadataSavers sets MetadataSavers field to given value.

### HasMetadataSavers

`func (o *LibraryOptions) HasMetadataSavers() bool`

HasMetadataSavers returns a boolean if a field has been set.

### GetDisabledLocalMetadataReaders

`func (o *LibraryOptions) GetDisabledLocalMetadataReaders() []string`

GetDisabledLocalMetadataReaders returns the DisabledLocalMetadataReaders field if non-nil, zero value otherwise.

### GetDisabledLocalMetadataReadersOk

`func (o *LibraryOptions) GetDisabledLocalMetadataReadersOk() (*[]string, bool)`

GetDisabledLocalMetadataReadersOk returns a tuple with the DisabledLocalMetadataReaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLocalMetadataReaders

`func (o *LibraryOptions) SetDisabledLocalMetadataReaders(v []string)`

SetDisabledLocalMetadataReaders sets DisabledLocalMetadataReaders field to given value.

### HasDisabledLocalMetadataReaders

`func (o *LibraryOptions) HasDisabledLocalMetadataReaders() bool`

HasDisabledLocalMetadataReaders returns a boolean if a field has been set.

### GetLocalMetadataReaderOrder

`func (o *LibraryOptions) GetLocalMetadataReaderOrder() []string`

GetLocalMetadataReaderOrder returns the LocalMetadataReaderOrder field if non-nil, zero value otherwise.

### GetLocalMetadataReaderOrderOk

`func (o *LibraryOptions) GetLocalMetadataReaderOrderOk() (*[]string, bool)`

GetLocalMetadataReaderOrderOk returns a tuple with the LocalMetadataReaderOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMetadataReaderOrder

`func (o *LibraryOptions) SetLocalMetadataReaderOrder(v []string)`

SetLocalMetadataReaderOrder sets LocalMetadataReaderOrder field to given value.

### HasLocalMetadataReaderOrder

`func (o *LibraryOptions) HasLocalMetadataReaderOrder() bool`

HasLocalMetadataReaderOrder returns a boolean if a field has been set.

### GetDisabledLyricsFetchers

`func (o *LibraryOptions) GetDisabledLyricsFetchers() []string`

GetDisabledLyricsFetchers returns the DisabledLyricsFetchers field if non-nil, zero value otherwise.

### GetDisabledLyricsFetchersOk

`func (o *LibraryOptions) GetDisabledLyricsFetchersOk() (*[]string, bool)`

GetDisabledLyricsFetchersOk returns a tuple with the DisabledLyricsFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLyricsFetchers

`func (o *LibraryOptions) SetDisabledLyricsFetchers(v []string)`

SetDisabledLyricsFetchers sets DisabledLyricsFetchers field to given value.

### HasDisabledLyricsFetchers

`func (o *LibraryOptions) HasDisabledLyricsFetchers() bool`

HasDisabledLyricsFetchers returns a boolean if a field has been set.

### GetSaveLyricsWithMedia

`func (o *LibraryOptions) GetSaveLyricsWithMedia() bool`

GetSaveLyricsWithMedia returns the SaveLyricsWithMedia field if non-nil, zero value otherwise.

### GetSaveLyricsWithMediaOk

`func (o *LibraryOptions) GetSaveLyricsWithMediaOk() (*bool, bool)`

GetSaveLyricsWithMediaOk returns a tuple with the SaveLyricsWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLyricsWithMedia

`func (o *LibraryOptions) SetSaveLyricsWithMedia(v bool)`

SetSaveLyricsWithMedia sets SaveLyricsWithMedia field to given value.

### HasSaveLyricsWithMedia

`func (o *LibraryOptions) HasSaveLyricsWithMedia() bool`

HasSaveLyricsWithMedia returns a boolean if a field has been set.

### GetLyricsDownloadMaxAgeDays

`func (o *LibraryOptions) GetLyricsDownloadMaxAgeDays() int32`

GetLyricsDownloadMaxAgeDays returns the LyricsDownloadMaxAgeDays field if non-nil, zero value otherwise.

### GetLyricsDownloadMaxAgeDaysOk

`func (o *LibraryOptions) GetLyricsDownloadMaxAgeDaysOk() (*int32, bool)`

GetLyricsDownloadMaxAgeDaysOk returns a tuple with the LyricsDownloadMaxAgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyricsDownloadMaxAgeDays

`func (o *LibraryOptions) SetLyricsDownloadMaxAgeDays(v int32)`

SetLyricsDownloadMaxAgeDays sets LyricsDownloadMaxAgeDays field to given value.

### HasLyricsDownloadMaxAgeDays

`func (o *LibraryOptions) HasLyricsDownloadMaxAgeDays() bool`

HasLyricsDownloadMaxAgeDays returns a boolean if a field has been set.

### GetLyricsFetcherOrder

`func (o *LibraryOptions) GetLyricsFetcherOrder() []string`

GetLyricsFetcherOrder returns the LyricsFetcherOrder field if non-nil, zero value otherwise.

### GetLyricsFetcherOrderOk

`func (o *LibraryOptions) GetLyricsFetcherOrderOk() (*[]string, bool)`

GetLyricsFetcherOrderOk returns a tuple with the LyricsFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyricsFetcherOrder

`func (o *LibraryOptions) SetLyricsFetcherOrder(v []string)`

SetLyricsFetcherOrder sets LyricsFetcherOrder field to given value.

### HasLyricsFetcherOrder

`func (o *LibraryOptions) HasLyricsFetcherOrder() bool`

HasLyricsFetcherOrder returns a boolean if a field has been set.

### GetLyricsDownloadLanguages

`func (o *LibraryOptions) GetLyricsDownloadLanguages() []string`

GetLyricsDownloadLanguages returns the LyricsDownloadLanguages field if non-nil, zero value otherwise.

### GetLyricsDownloadLanguagesOk

`func (o *LibraryOptions) GetLyricsDownloadLanguagesOk() (*[]string, bool)`

GetLyricsDownloadLanguagesOk returns a tuple with the LyricsDownloadLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyricsDownloadLanguages

`func (o *LibraryOptions) SetLyricsDownloadLanguages(v []string)`

SetLyricsDownloadLanguages sets LyricsDownloadLanguages field to given value.

### HasLyricsDownloadLanguages

`func (o *LibraryOptions) HasLyricsDownloadLanguages() bool`

HasLyricsDownloadLanguages returns a boolean if a field has been set.

### GetDisabledSubtitleFetchers

`func (o *LibraryOptions) GetDisabledSubtitleFetchers() []string`

GetDisabledSubtitleFetchers returns the DisabledSubtitleFetchers field if non-nil, zero value otherwise.

### GetDisabledSubtitleFetchersOk

`func (o *LibraryOptions) GetDisabledSubtitleFetchersOk() (*[]string, bool)`

GetDisabledSubtitleFetchersOk returns a tuple with the DisabledSubtitleFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledSubtitleFetchers

`func (o *LibraryOptions) SetDisabledSubtitleFetchers(v []string)`

SetDisabledSubtitleFetchers sets DisabledSubtitleFetchers field to given value.

### HasDisabledSubtitleFetchers

`func (o *LibraryOptions) HasDisabledSubtitleFetchers() bool`

HasDisabledSubtitleFetchers returns a boolean if a field has been set.

### GetSubtitleFetcherOrder

`func (o *LibraryOptions) GetSubtitleFetcherOrder() []string`

GetSubtitleFetcherOrder returns the SubtitleFetcherOrder field if non-nil, zero value otherwise.

### GetSubtitleFetcherOrderOk

`func (o *LibraryOptions) GetSubtitleFetcherOrderOk() (*[]string, bool)`

GetSubtitleFetcherOrderOk returns a tuple with the SubtitleFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleFetcherOrder

`func (o *LibraryOptions) SetSubtitleFetcherOrder(v []string)`

SetSubtitleFetcherOrder sets SubtitleFetcherOrder field to given value.

### HasSubtitleFetcherOrder

`func (o *LibraryOptions) HasSubtitleFetcherOrder() bool`

HasSubtitleFetcherOrder returns a boolean if a field has been set.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *LibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

GetSkipSubtitlesIfEmbeddedSubtitlesPresent returns the SkipSubtitlesIfEmbeddedSubtitlesPresent field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk

`func (o *LibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk() (*bool, bool)`

GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipSubtitlesIfEmbeddedSubtitlesPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *LibraryOptions) SetSkipSubtitlesIfEmbeddedSubtitlesPresent(v bool)`

SetSkipSubtitlesIfEmbeddedSubtitlesPresent sets SkipSubtitlesIfEmbeddedSubtitlesPresent field to given value.

### HasSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *LibraryOptions) HasSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

HasSkipSubtitlesIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.

### GetSkipSubtitlesIfAudioTrackMatches

`func (o *LibraryOptions) GetSkipSubtitlesIfAudioTrackMatches() bool`

GetSkipSubtitlesIfAudioTrackMatches returns the SkipSubtitlesIfAudioTrackMatches field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfAudioTrackMatchesOk

`func (o *LibraryOptions) GetSkipSubtitlesIfAudioTrackMatchesOk() (*bool, bool)`

GetSkipSubtitlesIfAudioTrackMatchesOk returns a tuple with the SkipSubtitlesIfAudioTrackMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfAudioTrackMatches

`func (o *LibraryOptions) SetSkipSubtitlesIfAudioTrackMatches(v bool)`

SetSkipSubtitlesIfAudioTrackMatches sets SkipSubtitlesIfAudioTrackMatches field to given value.

### HasSkipSubtitlesIfAudioTrackMatches

`func (o *LibraryOptions) HasSkipSubtitlesIfAudioTrackMatches() bool`

HasSkipSubtitlesIfAudioTrackMatches returns a boolean if a field has been set.

### GetSubtitleDownloadLanguages

`func (o *LibraryOptions) GetSubtitleDownloadLanguages() []string`

GetSubtitleDownloadLanguages returns the SubtitleDownloadLanguages field if non-nil, zero value otherwise.

### GetSubtitleDownloadLanguagesOk

`func (o *LibraryOptions) GetSubtitleDownloadLanguagesOk() (*[]string, bool)`

GetSubtitleDownloadLanguagesOk returns a tuple with the SubtitleDownloadLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleDownloadLanguages

`func (o *LibraryOptions) SetSubtitleDownloadLanguages(v []string)`

SetSubtitleDownloadLanguages sets SubtitleDownloadLanguages field to given value.

### HasSubtitleDownloadLanguages

`func (o *LibraryOptions) HasSubtitleDownloadLanguages() bool`

HasSubtitleDownloadLanguages returns a boolean if a field has been set.

### GetSubtitleDownloadMaxAgeDays

`func (o *LibraryOptions) GetSubtitleDownloadMaxAgeDays() int32`

GetSubtitleDownloadMaxAgeDays returns the SubtitleDownloadMaxAgeDays field if non-nil, zero value otherwise.

### GetSubtitleDownloadMaxAgeDaysOk

`func (o *LibraryOptions) GetSubtitleDownloadMaxAgeDaysOk() (*int32, bool)`

GetSubtitleDownloadMaxAgeDaysOk returns a tuple with the SubtitleDownloadMaxAgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleDownloadMaxAgeDays

`func (o *LibraryOptions) SetSubtitleDownloadMaxAgeDays(v int32)`

SetSubtitleDownloadMaxAgeDays sets SubtitleDownloadMaxAgeDays field to given value.

### HasSubtitleDownloadMaxAgeDays

`func (o *LibraryOptions) HasSubtitleDownloadMaxAgeDays() bool`

HasSubtitleDownloadMaxAgeDays returns a boolean if a field has been set.

### GetRequirePerfectSubtitleMatch

`func (o *LibraryOptions) GetRequirePerfectSubtitleMatch() bool`

GetRequirePerfectSubtitleMatch returns the RequirePerfectSubtitleMatch field if non-nil, zero value otherwise.

### GetRequirePerfectSubtitleMatchOk

`func (o *LibraryOptions) GetRequirePerfectSubtitleMatchOk() (*bool, bool)`

GetRequirePerfectSubtitleMatchOk returns a tuple with the RequirePerfectSubtitleMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePerfectSubtitleMatch

`func (o *LibraryOptions) SetRequirePerfectSubtitleMatch(v bool)`

SetRequirePerfectSubtitleMatch sets RequirePerfectSubtitleMatch field to given value.

### HasRequirePerfectSubtitleMatch

`func (o *LibraryOptions) HasRequirePerfectSubtitleMatch() bool`

HasRequirePerfectSubtitleMatch returns a boolean if a field has been set.

### GetSaveSubtitlesWithMedia

`func (o *LibraryOptions) GetSaveSubtitlesWithMedia() bool`

GetSaveSubtitlesWithMedia returns the SaveSubtitlesWithMedia field if non-nil, zero value otherwise.

### GetSaveSubtitlesWithMediaOk

`func (o *LibraryOptions) GetSaveSubtitlesWithMediaOk() (*bool, bool)`

GetSaveSubtitlesWithMediaOk returns a tuple with the SaveSubtitlesWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveSubtitlesWithMedia

`func (o *LibraryOptions) SetSaveSubtitlesWithMedia(v bool)`

SetSaveSubtitlesWithMedia sets SaveSubtitlesWithMedia field to given value.

### HasSaveSubtitlesWithMedia

`func (o *LibraryOptions) HasSaveSubtitlesWithMedia() bool`

HasSaveSubtitlesWithMedia returns a boolean if a field has been set.

### GetForcedSubtitlesOnly

`func (o *LibraryOptions) GetForcedSubtitlesOnly() bool`

GetForcedSubtitlesOnly returns the ForcedSubtitlesOnly field if non-nil, zero value otherwise.

### GetForcedSubtitlesOnlyOk

`func (o *LibraryOptions) GetForcedSubtitlesOnlyOk() (*bool, bool)`

GetForcedSubtitlesOnlyOk returns a tuple with the ForcedSubtitlesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSubtitlesOnly

`func (o *LibraryOptions) SetForcedSubtitlesOnly(v bool)`

SetForcedSubtitlesOnly sets ForcedSubtitlesOnly field to given value.

### HasForcedSubtitlesOnly

`func (o *LibraryOptions) HasForcedSubtitlesOnly() bool`

HasForcedSubtitlesOnly returns a boolean if a field has been set.

### GetHearingImpairedSubtitlesOnly

`func (o *LibraryOptions) GetHearingImpairedSubtitlesOnly() bool`

GetHearingImpairedSubtitlesOnly returns the HearingImpairedSubtitlesOnly field if non-nil, zero value otherwise.

### GetHearingImpairedSubtitlesOnlyOk

`func (o *LibraryOptions) GetHearingImpairedSubtitlesOnlyOk() (*bool, bool)`

GetHearingImpairedSubtitlesOnlyOk returns a tuple with the HearingImpairedSubtitlesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHearingImpairedSubtitlesOnly

`func (o *LibraryOptions) SetHearingImpairedSubtitlesOnly(v bool)`

SetHearingImpairedSubtitlesOnly sets HearingImpairedSubtitlesOnly field to given value.

### HasHearingImpairedSubtitlesOnly

`func (o *LibraryOptions) HasHearingImpairedSubtitlesOnly() bool`

HasHearingImpairedSubtitlesOnly returns a boolean if a field has been set.

### GetTypeOptions

`func (o *LibraryOptions) GetTypeOptions() []TypeOptions`

GetTypeOptions returns the TypeOptions field if non-nil, zero value otherwise.

### GetTypeOptionsOk

`func (o *LibraryOptions) GetTypeOptionsOk() (*[]TypeOptions, bool)`

GetTypeOptionsOk returns a tuple with the TypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOptions

`func (o *LibraryOptions) SetTypeOptions(v []TypeOptions)`

SetTypeOptions sets TypeOptions field to given value.

### HasTypeOptions

`func (o *LibraryOptions) HasTypeOptions() bool`

HasTypeOptions returns a boolean if a field has been set.

### GetCollapseSingleItemFolders

`func (o *LibraryOptions) GetCollapseSingleItemFolders() bool`

GetCollapseSingleItemFolders returns the CollapseSingleItemFolders field if non-nil, zero value otherwise.

### GetCollapseSingleItemFoldersOk

`func (o *LibraryOptions) GetCollapseSingleItemFoldersOk() (*bool, bool)`

GetCollapseSingleItemFoldersOk returns a tuple with the CollapseSingleItemFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseSingleItemFolders

`func (o *LibraryOptions) SetCollapseSingleItemFolders(v bool)`

SetCollapseSingleItemFolders sets CollapseSingleItemFolders field to given value.

### HasCollapseSingleItemFolders

`func (o *LibraryOptions) HasCollapseSingleItemFolders() bool`

HasCollapseSingleItemFolders returns a boolean if a field has been set.

### GetForceCollapseSingleItemFolders

`func (o *LibraryOptions) GetForceCollapseSingleItemFolders() bool`

GetForceCollapseSingleItemFolders returns the ForceCollapseSingleItemFolders field if non-nil, zero value otherwise.

### GetForceCollapseSingleItemFoldersOk

`func (o *LibraryOptions) GetForceCollapseSingleItemFoldersOk() (*bool, bool)`

GetForceCollapseSingleItemFoldersOk returns a tuple with the ForceCollapseSingleItemFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceCollapseSingleItemFolders

`func (o *LibraryOptions) SetForceCollapseSingleItemFolders(v bool)`

SetForceCollapseSingleItemFolders sets ForceCollapseSingleItemFolders field to given value.

### HasForceCollapseSingleItemFolders

`func (o *LibraryOptions) HasForceCollapseSingleItemFolders() bool`

HasForceCollapseSingleItemFolders returns a boolean if a field has been set.

### GetEnableAdultMetadata

`func (o *LibraryOptions) GetEnableAdultMetadata() bool`

GetEnableAdultMetadata returns the EnableAdultMetadata field if non-nil, zero value otherwise.

### GetEnableAdultMetadataOk

`func (o *LibraryOptions) GetEnableAdultMetadataOk() (*bool, bool)`

GetEnableAdultMetadataOk returns a tuple with the EnableAdultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdultMetadata

`func (o *LibraryOptions) SetEnableAdultMetadata(v bool)`

SetEnableAdultMetadata sets EnableAdultMetadata field to given value.

### HasEnableAdultMetadata

`func (o *LibraryOptions) HasEnableAdultMetadata() bool`

HasEnableAdultMetadata returns a boolean if a field has been set.

### GetImportCollections

`func (o *LibraryOptions) GetImportCollections() bool`

GetImportCollections returns the ImportCollections field if non-nil, zero value otherwise.

### GetImportCollectionsOk

`func (o *LibraryOptions) GetImportCollectionsOk() (*bool, bool)`

GetImportCollectionsOk returns a tuple with the ImportCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportCollections

`func (o *LibraryOptions) SetImportCollections(v bool)`

SetImportCollections sets ImportCollections field to given value.

### HasImportCollections

`func (o *LibraryOptions) HasImportCollections() bool`

HasImportCollections returns a boolean if a field has been set.

### GetEnableMultiVersionByFiles

`func (o *LibraryOptions) GetEnableMultiVersionByFiles() bool`

GetEnableMultiVersionByFiles returns the EnableMultiVersionByFiles field if non-nil, zero value otherwise.

### GetEnableMultiVersionByFilesOk

`func (o *LibraryOptions) GetEnableMultiVersionByFilesOk() (*bool, bool)`

GetEnableMultiVersionByFilesOk returns a tuple with the EnableMultiVersionByFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultiVersionByFiles

`func (o *LibraryOptions) SetEnableMultiVersionByFiles(v bool)`

SetEnableMultiVersionByFiles sets EnableMultiVersionByFiles field to given value.

### HasEnableMultiVersionByFiles

`func (o *LibraryOptions) HasEnableMultiVersionByFiles() bool`

HasEnableMultiVersionByFiles returns a boolean if a field has been set.

### GetEnableMultiVersionByMetadata

`func (o *LibraryOptions) GetEnableMultiVersionByMetadata() bool`

GetEnableMultiVersionByMetadata returns the EnableMultiVersionByMetadata field if non-nil, zero value otherwise.

### GetEnableMultiVersionByMetadataOk

`func (o *LibraryOptions) GetEnableMultiVersionByMetadataOk() (*bool, bool)`

GetEnableMultiVersionByMetadataOk returns a tuple with the EnableMultiVersionByMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultiVersionByMetadata

`func (o *LibraryOptions) SetEnableMultiVersionByMetadata(v bool)`

SetEnableMultiVersionByMetadata sets EnableMultiVersionByMetadata field to given value.

### HasEnableMultiVersionByMetadata

`func (o *LibraryOptions) HasEnableMultiVersionByMetadata() bool`

HasEnableMultiVersionByMetadata returns a boolean if a field has been set.

### GetEnableMultiPartItems

`func (o *LibraryOptions) GetEnableMultiPartItems() bool`

GetEnableMultiPartItems returns the EnableMultiPartItems field if non-nil, zero value otherwise.

### GetEnableMultiPartItemsOk

`func (o *LibraryOptions) GetEnableMultiPartItemsOk() (*bool, bool)`

GetEnableMultiPartItemsOk returns a tuple with the EnableMultiPartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultiPartItems

`func (o *LibraryOptions) SetEnableMultiPartItems(v bool)`

SetEnableMultiPartItems sets EnableMultiPartItems field to given value.

### HasEnableMultiPartItems

`func (o *LibraryOptions) HasEnableMultiPartItems() bool`

HasEnableMultiPartItems returns a boolean if a field has been set.

### GetMinCollectionItems

`func (o *LibraryOptions) GetMinCollectionItems() int32`

GetMinCollectionItems returns the MinCollectionItems field if non-nil, zero value otherwise.

### GetMinCollectionItemsOk

`func (o *LibraryOptions) GetMinCollectionItemsOk() (*int32, bool)`

GetMinCollectionItemsOk returns a tuple with the MinCollectionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCollectionItems

`func (o *LibraryOptions) SetMinCollectionItems(v int32)`

SetMinCollectionItems sets MinCollectionItems field to given value.

### HasMinCollectionItems

`func (o *LibraryOptions) HasMinCollectionItems() bool`

HasMinCollectionItems returns a boolean if a field has been set.

### GetMusicFolderStructure

`func (o *LibraryOptions) GetMusicFolderStructure() string`

GetMusicFolderStructure returns the MusicFolderStructure field if non-nil, zero value otherwise.

### GetMusicFolderStructureOk

`func (o *LibraryOptions) GetMusicFolderStructureOk() (*string, bool)`

GetMusicFolderStructureOk returns a tuple with the MusicFolderStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicFolderStructure

`func (o *LibraryOptions) SetMusicFolderStructure(v string)`

SetMusicFolderStructure sets MusicFolderStructure field to given value.

### HasMusicFolderStructure

`func (o *LibraryOptions) HasMusicFolderStructure() bool`

HasMusicFolderStructure returns a boolean if a field has been set.

### GetMinResumePct

`func (o *LibraryOptions) GetMinResumePct() int32`

GetMinResumePct returns the MinResumePct field if non-nil, zero value otherwise.

### GetMinResumePctOk

`func (o *LibraryOptions) GetMinResumePctOk() (*int32, bool)`

GetMinResumePctOk returns a tuple with the MinResumePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinResumePct

`func (o *LibraryOptions) SetMinResumePct(v int32)`

SetMinResumePct sets MinResumePct field to given value.

### HasMinResumePct

`func (o *LibraryOptions) HasMinResumePct() bool`

HasMinResumePct returns a boolean if a field has been set.

### GetMaxResumePct

`func (o *LibraryOptions) GetMaxResumePct() int32`

GetMaxResumePct returns the MaxResumePct field if non-nil, zero value otherwise.

### GetMaxResumePctOk

`func (o *LibraryOptions) GetMaxResumePctOk() (*int32, bool)`

GetMaxResumePctOk returns a tuple with the MaxResumePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResumePct

`func (o *LibraryOptions) SetMaxResumePct(v int32)`

SetMaxResumePct sets MaxResumePct field to given value.

### HasMaxResumePct

`func (o *LibraryOptions) HasMaxResumePct() bool`

HasMaxResumePct returns a boolean if a field has been set.

### GetMinResumeDurationSeconds

`func (o *LibraryOptions) GetMinResumeDurationSeconds() int32`

GetMinResumeDurationSeconds returns the MinResumeDurationSeconds field if non-nil, zero value otherwise.

### GetMinResumeDurationSecondsOk

`func (o *LibraryOptions) GetMinResumeDurationSecondsOk() (*int32, bool)`

GetMinResumeDurationSecondsOk returns a tuple with the MinResumeDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinResumeDurationSeconds

`func (o *LibraryOptions) SetMinResumeDurationSeconds(v int32)`

SetMinResumeDurationSeconds sets MinResumeDurationSeconds field to given value.

### HasMinResumeDurationSeconds

`func (o *LibraryOptions) HasMinResumeDurationSeconds() bool`

HasMinResumeDurationSeconds returns a boolean if a field has been set.

### GetThumbnailImagesIntervalSeconds

`func (o *LibraryOptions) GetThumbnailImagesIntervalSeconds() int32`

GetThumbnailImagesIntervalSeconds returns the ThumbnailImagesIntervalSeconds field if non-nil, zero value otherwise.

### GetThumbnailImagesIntervalSecondsOk

`func (o *LibraryOptions) GetThumbnailImagesIntervalSecondsOk() (*int32, bool)`

GetThumbnailImagesIntervalSecondsOk returns a tuple with the ThumbnailImagesIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImagesIntervalSeconds

`func (o *LibraryOptions) SetThumbnailImagesIntervalSeconds(v int32)`

SetThumbnailImagesIntervalSeconds sets ThumbnailImagesIntervalSeconds field to given value.

### HasThumbnailImagesIntervalSeconds

`func (o *LibraryOptions) HasThumbnailImagesIntervalSeconds() bool`

HasThumbnailImagesIntervalSeconds returns a boolean if a field has been set.

### GetSampleIgnoreSize

`func (o *LibraryOptions) GetSampleIgnoreSize() int32`

GetSampleIgnoreSize returns the SampleIgnoreSize field if non-nil, zero value otherwise.

### GetSampleIgnoreSizeOk

`func (o *LibraryOptions) GetSampleIgnoreSizeOk() (*int32, bool)`

GetSampleIgnoreSizeOk returns a tuple with the SampleIgnoreSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIgnoreSize

`func (o *LibraryOptions) SetSampleIgnoreSize(v int32)`

SetSampleIgnoreSize sets SampleIgnoreSize field to given value.

### HasSampleIgnoreSize

`func (o *LibraryOptions) HasSampleIgnoreSize() bool`

HasSampleIgnoreSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


