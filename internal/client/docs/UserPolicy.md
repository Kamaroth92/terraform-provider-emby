# UserPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAdministrator** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**IsHiddenRemotely** | Pointer to **bool** |  | [optional] 
**IsHiddenFromUnusedDevices** | Pointer to **bool** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**LockedOutDate** | Pointer to **int64** |  | [optional] 
**MaxParentalRating** | Pointer to **NullableInt32** |  | [optional] 
**AllowTagOrRating** | Pointer to **bool** |  | [optional] 
**BlockedTags** | Pointer to **[]string** |  | [optional] 
**IsTagBlockingModeInclusive** | Pointer to **bool** |  | [optional] 
**IncludeTags** | Pointer to **[]string** |  | [optional] 
**EnableUserPreferenceAccess** | Pointer to **bool** |  | [optional] 
**AccessSchedules** | Pointer to [**[]AccessSchedule**](AccessSchedule.md) |  | [optional] 
**BlockUnratedItems** | Pointer to [**[]UnratedItem**](UnratedItem.md) |  | [optional] 
**EnableRemoteControlOfOtherUsers** | Pointer to **bool** |  | [optional] 
**EnableSharedDeviceControl** | Pointer to **bool** |  | [optional] 
**EnableRemoteAccess** | Pointer to **bool** |  | [optional] 
**EnableLiveTvManagement** | Pointer to **bool** |  | [optional] 
**EnableLiveTvAccess** | Pointer to **bool** |  | [optional] 
**EnableMediaPlayback** | Pointer to **bool** |  | [optional] 
**EnableAudioPlaybackTranscoding** | Pointer to **bool** |  | [optional] 
**EnableVideoPlaybackTranscoding** | Pointer to **bool** |  | [optional] 
**EnableTranscodingQuality** | Pointer to **bool** |  | [optional] 
**AutoRemoteQuality** | Pointer to **int32** |  | [optional] 
**EnablePlaybackRemuxing** | Pointer to **bool** |  | [optional] 
**EnableContentDeletion** | Pointer to **bool** |  | [optional] 
**RestrictedFeatures** | Pointer to **[]string** |  | [optional] 
**EnableContentDeletionFromFolders** | Pointer to **[]string** |  | [optional] 
**EnableContentDownloading** | Pointer to **bool** |  | [optional] 
**EnableSubtitleDownloading** | Pointer to **bool** |  | [optional] 
**EnableSubtitleManagement** | Pointer to **bool** |  | [optional] 
**EnableSyncTranscoding** | Pointer to **bool** |  | [optional] 
**EnableMediaConversion** | Pointer to **bool** |  | [optional] 
**EnabledChannels** | Pointer to **[]string** |  | [optional] 
**EnableAllChannels** | Pointer to **bool** |  | [optional] 
**EnabledFolders** | Pointer to **[]string** |  | [optional] 
**EnableAllFolders** | Pointer to **bool** |  | [optional] 
**InvalidLoginAttemptCount** | Pointer to **int32** |  | [optional] 
**EnablePublicSharing** | Pointer to **bool** |  | [optional] 
**RemoteClientBitrateLimit** | Pointer to **int32** |  | [optional] 
**AuthenticationProviderId** | Pointer to **string** |  | [optional] 
**ExcludedSubFolders** | Pointer to **[]string** |  | [optional] 
**SimultaneousStreamLimit** | Pointer to **int32** |  | [optional] 
**EnabledDevices** | Pointer to **[]string** |  | [optional] 
**EnableAllDevices** | Pointer to **bool** |  | [optional] 
**AllowCameraUpload** | Pointer to **bool** |  | [optional] 
**AllowSharingPersonalItems** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserPolicy

`func NewUserPolicy() *UserPolicy`

NewUserPolicy instantiates a new UserPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPolicyWithDefaults

`func NewUserPolicyWithDefaults() *UserPolicy`

NewUserPolicyWithDefaults instantiates a new UserPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAdministrator

`func (o *UserPolicy) GetIsAdministrator() bool`

GetIsAdministrator returns the IsAdministrator field if non-nil, zero value otherwise.

### GetIsAdministratorOk

`func (o *UserPolicy) GetIsAdministratorOk() (*bool, bool)`

GetIsAdministratorOk returns a tuple with the IsAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdministrator

`func (o *UserPolicy) SetIsAdministrator(v bool)`

SetIsAdministrator sets IsAdministrator field to given value.

### HasIsAdministrator

`func (o *UserPolicy) HasIsAdministrator() bool`

HasIsAdministrator returns a boolean if a field has been set.

### GetIsHidden

`func (o *UserPolicy) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *UserPolicy) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *UserPolicy) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *UserPolicy) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsHiddenRemotely

`func (o *UserPolicy) GetIsHiddenRemotely() bool`

GetIsHiddenRemotely returns the IsHiddenRemotely field if non-nil, zero value otherwise.

### GetIsHiddenRemotelyOk

`func (o *UserPolicy) GetIsHiddenRemotelyOk() (*bool, bool)`

GetIsHiddenRemotelyOk returns a tuple with the IsHiddenRemotely field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHiddenRemotely

`func (o *UserPolicy) SetIsHiddenRemotely(v bool)`

SetIsHiddenRemotely sets IsHiddenRemotely field to given value.

### HasIsHiddenRemotely

`func (o *UserPolicy) HasIsHiddenRemotely() bool`

HasIsHiddenRemotely returns a boolean if a field has been set.

### GetIsHiddenFromUnusedDevices

`func (o *UserPolicy) GetIsHiddenFromUnusedDevices() bool`

GetIsHiddenFromUnusedDevices returns the IsHiddenFromUnusedDevices field if non-nil, zero value otherwise.

### GetIsHiddenFromUnusedDevicesOk

`func (o *UserPolicy) GetIsHiddenFromUnusedDevicesOk() (*bool, bool)`

GetIsHiddenFromUnusedDevicesOk returns a tuple with the IsHiddenFromUnusedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHiddenFromUnusedDevices

`func (o *UserPolicy) SetIsHiddenFromUnusedDevices(v bool)`

SetIsHiddenFromUnusedDevices sets IsHiddenFromUnusedDevices field to given value.

### HasIsHiddenFromUnusedDevices

`func (o *UserPolicy) HasIsHiddenFromUnusedDevices() bool`

HasIsHiddenFromUnusedDevices returns a boolean if a field has been set.

### GetIsDisabled

`func (o *UserPolicy) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UserPolicy) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UserPolicy) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UserPolicy) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetLockedOutDate

`func (o *UserPolicy) GetLockedOutDate() int64`

GetLockedOutDate returns the LockedOutDate field if non-nil, zero value otherwise.

### GetLockedOutDateOk

`func (o *UserPolicy) GetLockedOutDateOk() (*int64, bool)`

GetLockedOutDateOk returns a tuple with the LockedOutDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedOutDate

`func (o *UserPolicy) SetLockedOutDate(v int64)`

SetLockedOutDate sets LockedOutDate field to given value.

### HasLockedOutDate

`func (o *UserPolicy) HasLockedOutDate() bool`

HasLockedOutDate returns a boolean if a field has been set.

### GetMaxParentalRating

`func (o *UserPolicy) GetMaxParentalRating() int32`

GetMaxParentalRating returns the MaxParentalRating field if non-nil, zero value otherwise.

### GetMaxParentalRatingOk

`func (o *UserPolicy) GetMaxParentalRatingOk() (*int32, bool)`

GetMaxParentalRatingOk returns a tuple with the MaxParentalRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParentalRating

`func (o *UserPolicy) SetMaxParentalRating(v int32)`

SetMaxParentalRating sets MaxParentalRating field to given value.

### HasMaxParentalRating

`func (o *UserPolicy) HasMaxParentalRating() bool`

HasMaxParentalRating returns a boolean if a field has been set.

### SetMaxParentalRatingNil

`func (o *UserPolicy) SetMaxParentalRatingNil(b bool)`

 SetMaxParentalRatingNil sets the value for MaxParentalRating to be an explicit nil

### UnsetMaxParentalRating
`func (o *UserPolicy) UnsetMaxParentalRating()`

UnsetMaxParentalRating ensures that no value is present for MaxParentalRating, not even an explicit nil
### GetAllowTagOrRating

`func (o *UserPolicy) GetAllowTagOrRating() bool`

GetAllowTagOrRating returns the AllowTagOrRating field if non-nil, zero value otherwise.

### GetAllowTagOrRatingOk

`func (o *UserPolicy) GetAllowTagOrRatingOk() (*bool, bool)`

GetAllowTagOrRatingOk returns a tuple with the AllowTagOrRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTagOrRating

`func (o *UserPolicy) SetAllowTagOrRating(v bool)`

SetAllowTagOrRating sets AllowTagOrRating field to given value.

### HasAllowTagOrRating

`func (o *UserPolicy) HasAllowTagOrRating() bool`

HasAllowTagOrRating returns a boolean if a field has been set.

### GetBlockedTags

`func (o *UserPolicy) GetBlockedTags() []string`

GetBlockedTags returns the BlockedTags field if non-nil, zero value otherwise.

### GetBlockedTagsOk

`func (o *UserPolicy) GetBlockedTagsOk() (*[]string, bool)`

GetBlockedTagsOk returns a tuple with the BlockedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedTags

`func (o *UserPolicy) SetBlockedTags(v []string)`

SetBlockedTags sets BlockedTags field to given value.

### HasBlockedTags

`func (o *UserPolicy) HasBlockedTags() bool`

HasBlockedTags returns a boolean if a field has been set.

### GetIsTagBlockingModeInclusive

`func (o *UserPolicy) GetIsTagBlockingModeInclusive() bool`

GetIsTagBlockingModeInclusive returns the IsTagBlockingModeInclusive field if non-nil, zero value otherwise.

### GetIsTagBlockingModeInclusiveOk

`func (o *UserPolicy) GetIsTagBlockingModeInclusiveOk() (*bool, bool)`

GetIsTagBlockingModeInclusiveOk returns a tuple with the IsTagBlockingModeInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTagBlockingModeInclusive

`func (o *UserPolicy) SetIsTagBlockingModeInclusive(v bool)`

SetIsTagBlockingModeInclusive sets IsTagBlockingModeInclusive field to given value.

### HasIsTagBlockingModeInclusive

`func (o *UserPolicy) HasIsTagBlockingModeInclusive() bool`

HasIsTagBlockingModeInclusive returns a boolean if a field has been set.

### GetIncludeTags

`func (o *UserPolicy) GetIncludeTags() []string`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *UserPolicy) GetIncludeTagsOk() (*[]string, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *UserPolicy) SetIncludeTags(v []string)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *UserPolicy) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetEnableUserPreferenceAccess

`func (o *UserPolicy) GetEnableUserPreferenceAccess() bool`

GetEnableUserPreferenceAccess returns the EnableUserPreferenceAccess field if non-nil, zero value otherwise.

### GetEnableUserPreferenceAccessOk

`func (o *UserPolicy) GetEnableUserPreferenceAccessOk() (*bool, bool)`

GetEnableUserPreferenceAccessOk returns a tuple with the EnableUserPreferenceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserPreferenceAccess

`func (o *UserPolicy) SetEnableUserPreferenceAccess(v bool)`

SetEnableUserPreferenceAccess sets EnableUserPreferenceAccess field to given value.

### HasEnableUserPreferenceAccess

`func (o *UserPolicy) HasEnableUserPreferenceAccess() bool`

HasEnableUserPreferenceAccess returns a boolean if a field has been set.

### GetAccessSchedules

`func (o *UserPolicy) GetAccessSchedules() []AccessSchedule`

GetAccessSchedules returns the AccessSchedules field if non-nil, zero value otherwise.

### GetAccessSchedulesOk

`func (o *UserPolicy) GetAccessSchedulesOk() (*[]AccessSchedule, bool)`

GetAccessSchedulesOk returns a tuple with the AccessSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSchedules

`func (o *UserPolicy) SetAccessSchedules(v []AccessSchedule)`

SetAccessSchedules sets AccessSchedules field to given value.

### HasAccessSchedules

`func (o *UserPolicy) HasAccessSchedules() bool`

HasAccessSchedules returns a boolean if a field has been set.

### GetBlockUnratedItems

`func (o *UserPolicy) GetBlockUnratedItems() []UnratedItem`

GetBlockUnratedItems returns the BlockUnratedItems field if non-nil, zero value otherwise.

### GetBlockUnratedItemsOk

`func (o *UserPolicy) GetBlockUnratedItemsOk() (*[]UnratedItem, bool)`

GetBlockUnratedItemsOk returns a tuple with the BlockUnratedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUnratedItems

`func (o *UserPolicy) SetBlockUnratedItems(v []UnratedItem)`

SetBlockUnratedItems sets BlockUnratedItems field to given value.

### HasBlockUnratedItems

`func (o *UserPolicy) HasBlockUnratedItems() bool`

HasBlockUnratedItems returns a boolean if a field has been set.

### GetEnableRemoteControlOfOtherUsers

`func (o *UserPolicy) GetEnableRemoteControlOfOtherUsers() bool`

GetEnableRemoteControlOfOtherUsers returns the EnableRemoteControlOfOtherUsers field if non-nil, zero value otherwise.

### GetEnableRemoteControlOfOtherUsersOk

`func (o *UserPolicy) GetEnableRemoteControlOfOtherUsersOk() (*bool, bool)`

GetEnableRemoteControlOfOtherUsersOk returns a tuple with the EnableRemoteControlOfOtherUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteControlOfOtherUsers

`func (o *UserPolicy) SetEnableRemoteControlOfOtherUsers(v bool)`

SetEnableRemoteControlOfOtherUsers sets EnableRemoteControlOfOtherUsers field to given value.

### HasEnableRemoteControlOfOtherUsers

`func (o *UserPolicy) HasEnableRemoteControlOfOtherUsers() bool`

HasEnableRemoteControlOfOtherUsers returns a boolean if a field has been set.

### GetEnableSharedDeviceControl

`func (o *UserPolicy) GetEnableSharedDeviceControl() bool`

GetEnableSharedDeviceControl returns the EnableSharedDeviceControl field if non-nil, zero value otherwise.

### GetEnableSharedDeviceControlOk

`func (o *UserPolicy) GetEnableSharedDeviceControlOk() (*bool, bool)`

GetEnableSharedDeviceControlOk returns a tuple with the EnableSharedDeviceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSharedDeviceControl

`func (o *UserPolicy) SetEnableSharedDeviceControl(v bool)`

SetEnableSharedDeviceControl sets EnableSharedDeviceControl field to given value.

### HasEnableSharedDeviceControl

`func (o *UserPolicy) HasEnableSharedDeviceControl() bool`

HasEnableSharedDeviceControl returns a boolean if a field has been set.

### GetEnableRemoteAccess

`func (o *UserPolicy) GetEnableRemoteAccess() bool`

GetEnableRemoteAccess returns the EnableRemoteAccess field if non-nil, zero value otherwise.

### GetEnableRemoteAccessOk

`func (o *UserPolicy) GetEnableRemoteAccessOk() (*bool, bool)`

GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteAccess

`func (o *UserPolicy) SetEnableRemoteAccess(v bool)`

SetEnableRemoteAccess sets EnableRemoteAccess field to given value.

### HasEnableRemoteAccess

`func (o *UserPolicy) HasEnableRemoteAccess() bool`

HasEnableRemoteAccess returns a boolean if a field has been set.

### GetEnableLiveTvManagement

`func (o *UserPolicy) GetEnableLiveTvManagement() bool`

GetEnableLiveTvManagement returns the EnableLiveTvManagement field if non-nil, zero value otherwise.

### GetEnableLiveTvManagementOk

`func (o *UserPolicy) GetEnableLiveTvManagementOk() (*bool, bool)`

GetEnableLiveTvManagementOk returns a tuple with the EnableLiveTvManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveTvManagement

`func (o *UserPolicy) SetEnableLiveTvManagement(v bool)`

SetEnableLiveTvManagement sets EnableLiveTvManagement field to given value.

### HasEnableLiveTvManagement

`func (o *UserPolicy) HasEnableLiveTvManagement() bool`

HasEnableLiveTvManagement returns a boolean if a field has been set.

### GetEnableLiveTvAccess

`func (o *UserPolicy) GetEnableLiveTvAccess() bool`

GetEnableLiveTvAccess returns the EnableLiveTvAccess field if non-nil, zero value otherwise.

### GetEnableLiveTvAccessOk

`func (o *UserPolicy) GetEnableLiveTvAccessOk() (*bool, bool)`

GetEnableLiveTvAccessOk returns a tuple with the EnableLiveTvAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveTvAccess

`func (o *UserPolicy) SetEnableLiveTvAccess(v bool)`

SetEnableLiveTvAccess sets EnableLiveTvAccess field to given value.

### HasEnableLiveTvAccess

`func (o *UserPolicy) HasEnableLiveTvAccess() bool`

HasEnableLiveTvAccess returns a boolean if a field has been set.

### GetEnableMediaPlayback

`func (o *UserPolicy) GetEnableMediaPlayback() bool`

GetEnableMediaPlayback returns the EnableMediaPlayback field if non-nil, zero value otherwise.

### GetEnableMediaPlaybackOk

`func (o *UserPolicy) GetEnableMediaPlaybackOk() (*bool, bool)`

GetEnableMediaPlaybackOk returns a tuple with the EnableMediaPlayback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMediaPlayback

`func (o *UserPolicy) SetEnableMediaPlayback(v bool)`

SetEnableMediaPlayback sets EnableMediaPlayback field to given value.

### HasEnableMediaPlayback

`func (o *UserPolicy) HasEnableMediaPlayback() bool`

HasEnableMediaPlayback returns a boolean if a field has been set.

### GetEnableAudioPlaybackTranscoding

`func (o *UserPolicy) GetEnableAudioPlaybackTranscoding() bool`

GetEnableAudioPlaybackTranscoding returns the EnableAudioPlaybackTranscoding field if non-nil, zero value otherwise.

### GetEnableAudioPlaybackTranscodingOk

`func (o *UserPolicy) GetEnableAudioPlaybackTranscodingOk() (*bool, bool)`

GetEnableAudioPlaybackTranscodingOk returns a tuple with the EnableAudioPlaybackTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioPlaybackTranscoding

`func (o *UserPolicy) SetEnableAudioPlaybackTranscoding(v bool)`

SetEnableAudioPlaybackTranscoding sets EnableAudioPlaybackTranscoding field to given value.

### HasEnableAudioPlaybackTranscoding

`func (o *UserPolicy) HasEnableAudioPlaybackTranscoding() bool`

HasEnableAudioPlaybackTranscoding returns a boolean if a field has been set.

### GetEnableVideoPlaybackTranscoding

`func (o *UserPolicy) GetEnableVideoPlaybackTranscoding() bool`

GetEnableVideoPlaybackTranscoding returns the EnableVideoPlaybackTranscoding field if non-nil, zero value otherwise.

### GetEnableVideoPlaybackTranscodingOk

`func (o *UserPolicy) GetEnableVideoPlaybackTranscodingOk() (*bool, bool)`

GetEnableVideoPlaybackTranscodingOk returns a tuple with the EnableVideoPlaybackTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideoPlaybackTranscoding

`func (o *UserPolicy) SetEnableVideoPlaybackTranscoding(v bool)`

SetEnableVideoPlaybackTranscoding sets EnableVideoPlaybackTranscoding field to given value.

### HasEnableVideoPlaybackTranscoding

`func (o *UserPolicy) HasEnableVideoPlaybackTranscoding() bool`

HasEnableVideoPlaybackTranscoding returns a boolean if a field has been set.

### GetEnableTranscodingQuality

`func (o *UserPolicy) GetEnableTranscodingQuality() bool`

GetEnableTranscodingQuality returns the EnableTranscodingQuality field if non-nil, zero value otherwise.

### GetEnableTranscodingQualityOk

`func (o *UserPolicy) GetEnableTranscodingQualityOk() (*bool, bool)`

GetEnableTranscodingQualityOk returns a tuple with the EnableTranscodingQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTranscodingQuality

`func (o *UserPolicy) SetEnableTranscodingQuality(v bool)`

SetEnableTranscodingQuality sets EnableTranscodingQuality field to given value.

### HasEnableTranscodingQuality

`func (o *UserPolicy) HasEnableTranscodingQuality() bool`

HasEnableTranscodingQuality returns a boolean if a field has been set.

### GetAutoRemoteQuality

`func (o *UserPolicy) GetAutoRemoteQuality() int32`

GetAutoRemoteQuality returns the AutoRemoteQuality field if non-nil, zero value otherwise.

### GetAutoRemoteQualityOk

`func (o *UserPolicy) GetAutoRemoteQualityOk() (*int32, bool)`

GetAutoRemoteQualityOk returns a tuple with the AutoRemoteQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRemoteQuality

`func (o *UserPolicy) SetAutoRemoteQuality(v int32)`

SetAutoRemoteQuality sets AutoRemoteQuality field to given value.

### HasAutoRemoteQuality

`func (o *UserPolicy) HasAutoRemoteQuality() bool`

HasAutoRemoteQuality returns a boolean if a field has been set.

### GetEnablePlaybackRemuxing

`func (o *UserPolicy) GetEnablePlaybackRemuxing() bool`

GetEnablePlaybackRemuxing returns the EnablePlaybackRemuxing field if non-nil, zero value otherwise.

### GetEnablePlaybackRemuxingOk

`func (o *UserPolicy) GetEnablePlaybackRemuxingOk() (*bool, bool)`

GetEnablePlaybackRemuxingOk returns a tuple with the EnablePlaybackRemuxing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePlaybackRemuxing

`func (o *UserPolicy) SetEnablePlaybackRemuxing(v bool)`

SetEnablePlaybackRemuxing sets EnablePlaybackRemuxing field to given value.

### HasEnablePlaybackRemuxing

`func (o *UserPolicy) HasEnablePlaybackRemuxing() bool`

HasEnablePlaybackRemuxing returns a boolean if a field has been set.

### GetEnableContentDeletion

`func (o *UserPolicy) GetEnableContentDeletion() bool`

GetEnableContentDeletion returns the EnableContentDeletion field if non-nil, zero value otherwise.

### GetEnableContentDeletionOk

`func (o *UserPolicy) GetEnableContentDeletionOk() (*bool, bool)`

GetEnableContentDeletionOk returns a tuple with the EnableContentDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDeletion

`func (o *UserPolicy) SetEnableContentDeletion(v bool)`

SetEnableContentDeletion sets EnableContentDeletion field to given value.

### HasEnableContentDeletion

`func (o *UserPolicy) HasEnableContentDeletion() bool`

HasEnableContentDeletion returns a boolean if a field has been set.

### GetRestrictedFeatures

`func (o *UserPolicy) GetRestrictedFeatures() []string`

GetRestrictedFeatures returns the RestrictedFeatures field if non-nil, zero value otherwise.

### GetRestrictedFeaturesOk

`func (o *UserPolicy) GetRestrictedFeaturesOk() (*[]string, bool)`

GetRestrictedFeaturesOk returns a tuple with the RestrictedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedFeatures

`func (o *UserPolicy) SetRestrictedFeatures(v []string)`

SetRestrictedFeatures sets RestrictedFeatures field to given value.

### HasRestrictedFeatures

`func (o *UserPolicy) HasRestrictedFeatures() bool`

HasRestrictedFeatures returns a boolean if a field has been set.

### GetEnableContentDeletionFromFolders

`func (o *UserPolicy) GetEnableContentDeletionFromFolders() []string`

GetEnableContentDeletionFromFolders returns the EnableContentDeletionFromFolders field if non-nil, zero value otherwise.

### GetEnableContentDeletionFromFoldersOk

`func (o *UserPolicy) GetEnableContentDeletionFromFoldersOk() (*[]string, bool)`

GetEnableContentDeletionFromFoldersOk returns a tuple with the EnableContentDeletionFromFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDeletionFromFolders

`func (o *UserPolicy) SetEnableContentDeletionFromFolders(v []string)`

SetEnableContentDeletionFromFolders sets EnableContentDeletionFromFolders field to given value.

### HasEnableContentDeletionFromFolders

`func (o *UserPolicy) HasEnableContentDeletionFromFolders() bool`

HasEnableContentDeletionFromFolders returns a boolean if a field has been set.

### GetEnableContentDownloading

`func (o *UserPolicy) GetEnableContentDownloading() bool`

GetEnableContentDownloading returns the EnableContentDownloading field if non-nil, zero value otherwise.

### GetEnableContentDownloadingOk

`func (o *UserPolicy) GetEnableContentDownloadingOk() (*bool, bool)`

GetEnableContentDownloadingOk returns a tuple with the EnableContentDownloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDownloading

`func (o *UserPolicy) SetEnableContentDownloading(v bool)`

SetEnableContentDownloading sets EnableContentDownloading field to given value.

### HasEnableContentDownloading

`func (o *UserPolicy) HasEnableContentDownloading() bool`

HasEnableContentDownloading returns a boolean if a field has been set.

### GetEnableSubtitleDownloading

`func (o *UserPolicy) GetEnableSubtitleDownloading() bool`

GetEnableSubtitleDownloading returns the EnableSubtitleDownloading field if non-nil, zero value otherwise.

### GetEnableSubtitleDownloadingOk

`func (o *UserPolicy) GetEnableSubtitleDownloadingOk() (*bool, bool)`

GetEnableSubtitleDownloadingOk returns a tuple with the EnableSubtitleDownloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubtitleDownloading

`func (o *UserPolicy) SetEnableSubtitleDownloading(v bool)`

SetEnableSubtitleDownloading sets EnableSubtitleDownloading field to given value.

### HasEnableSubtitleDownloading

`func (o *UserPolicy) HasEnableSubtitleDownloading() bool`

HasEnableSubtitleDownloading returns a boolean if a field has been set.

### GetEnableSubtitleManagement

`func (o *UserPolicy) GetEnableSubtitleManagement() bool`

GetEnableSubtitleManagement returns the EnableSubtitleManagement field if non-nil, zero value otherwise.

### GetEnableSubtitleManagementOk

`func (o *UserPolicy) GetEnableSubtitleManagementOk() (*bool, bool)`

GetEnableSubtitleManagementOk returns a tuple with the EnableSubtitleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubtitleManagement

`func (o *UserPolicy) SetEnableSubtitleManagement(v bool)`

SetEnableSubtitleManagement sets EnableSubtitleManagement field to given value.

### HasEnableSubtitleManagement

`func (o *UserPolicy) HasEnableSubtitleManagement() bool`

HasEnableSubtitleManagement returns a boolean if a field has been set.

### GetEnableSyncTranscoding

`func (o *UserPolicy) GetEnableSyncTranscoding() bool`

GetEnableSyncTranscoding returns the EnableSyncTranscoding field if non-nil, zero value otherwise.

### GetEnableSyncTranscodingOk

`func (o *UserPolicy) GetEnableSyncTranscodingOk() (*bool, bool)`

GetEnableSyncTranscodingOk returns a tuple with the EnableSyncTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSyncTranscoding

`func (o *UserPolicy) SetEnableSyncTranscoding(v bool)`

SetEnableSyncTranscoding sets EnableSyncTranscoding field to given value.

### HasEnableSyncTranscoding

`func (o *UserPolicy) HasEnableSyncTranscoding() bool`

HasEnableSyncTranscoding returns a boolean if a field has been set.

### GetEnableMediaConversion

`func (o *UserPolicy) GetEnableMediaConversion() bool`

GetEnableMediaConversion returns the EnableMediaConversion field if non-nil, zero value otherwise.

### GetEnableMediaConversionOk

`func (o *UserPolicy) GetEnableMediaConversionOk() (*bool, bool)`

GetEnableMediaConversionOk returns a tuple with the EnableMediaConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMediaConversion

`func (o *UserPolicy) SetEnableMediaConversion(v bool)`

SetEnableMediaConversion sets EnableMediaConversion field to given value.

### HasEnableMediaConversion

`func (o *UserPolicy) HasEnableMediaConversion() bool`

HasEnableMediaConversion returns a boolean if a field has been set.

### GetEnabledChannels

`func (o *UserPolicy) GetEnabledChannels() []string`

GetEnabledChannels returns the EnabledChannels field if non-nil, zero value otherwise.

### GetEnabledChannelsOk

`func (o *UserPolicy) GetEnabledChannelsOk() (*[]string, bool)`

GetEnabledChannelsOk returns a tuple with the EnabledChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledChannels

`func (o *UserPolicy) SetEnabledChannels(v []string)`

SetEnabledChannels sets EnabledChannels field to given value.

### HasEnabledChannels

`func (o *UserPolicy) HasEnabledChannels() bool`

HasEnabledChannels returns a boolean if a field has been set.

### GetEnableAllChannels

`func (o *UserPolicy) GetEnableAllChannels() bool`

GetEnableAllChannels returns the EnableAllChannels field if non-nil, zero value otherwise.

### GetEnableAllChannelsOk

`func (o *UserPolicy) GetEnableAllChannelsOk() (*bool, bool)`

GetEnableAllChannelsOk returns a tuple with the EnableAllChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllChannels

`func (o *UserPolicy) SetEnableAllChannels(v bool)`

SetEnableAllChannels sets EnableAllChannels field to given value.

### HasEnableAllChannels

`func (o *UserPolicy) HasEnableAllChannels() bool`

HasEnableAllChannels returns a boolean if a field has been set.

### GetEnabledFolders

`func (o *UserPolicy) GetEnabledFolders() []string`

GetEnabledFolders returns the EnabledFolders field if non-nil, zero value otherwise.

### GetEnabledFoldersOk

`func (o *UserPolicy) GetEnabledFoldersOk() (*[]string, bool)`

GetEnabledFoldersOk returns a tuple with the EnabledFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFolders

`func (o *UserPolicy) SetEnabledFolders(v []string)`

SetEnabledFolders sets EnabledFolders field to given value.

### HasEnabledFolders

`func (o *UserPolicy) HasEnabledFolders() bool`

HasEnabledFolders returns a boolean if a field has been set.

### GetEnableAllFolders

`func (o *UserPolicy) GetEnableAllFolders() bool`

GetEnableAllFolders returns the EnableAllFolders field if non-nil, zero value otherwise.

### GetEnableAllFoldersOk

`func (o *UserPolicy) GetEnableAllFoldersOk() (*bool, bool)`

GetEnableAllFoldersOk returns a tuple with the EnableAllFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllFolders

`func (o *UserPolicy) SetEnableAllFolders(v bool)`

SetEnableAllFolders sets EnableAllFolders field to given value.

### HasEnableAllFolders

`func (o *UserPolicy) HasEnableAllFolders() bool`

HasEnableAllFolders returns a boolean if a field has been set.

### GetInvalidLoginAttemptCount

`func (o *UserPolicy) GetInvalidLoginAttemptCount() int32`

GetInvalidLoginAttemptCount returns the InvalidLoginAttemptCount field if non-nil, zero value otherwise.

### GetInvalidLoginAttemptCountOk

`func (o *UserPolicy) GetInvalidLoginAttemptCountOk() (*int32, bool)`

GetInvalidLoginAttemptCountOk returns a tuple with the InvalidLoginAttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidLoginAttemptCount

`func (o *UserPolicy) SetInvalidLoginAttemptCount(v int32)`

SetInvalidLoginAttemptCount sets InvalidLoginAttemptCount field to given value.

### HasInvalidLoginAttemptCount

`func (o *UserPolicy) HasInvalidLoginAttemptCount() bool`

HasInvalidLoginAttemptCount returns a boolean if a field has been set.

### GetEnablePublicSharing

`func (o *UserPolicy) GetEnablePublicSharing() bool`

GetEnablePublicSharing returns the EnablePublicSharing field if non-nil, zero value otherwise.

### GetEnablePublicSharingOk

`func (o *UserPolicy) GetEnablePublicSharingOk() (*bool, bool)`

GetEnablePublicSharingOk returns a tuple with the EnablePublicSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublicSharing

`func (o *UserPolicy) SetEnablePublicSharing(v bool)`

SetEnablePublicSharing sets EnablePublicSharing field to given value.

### HasEnablePublicSharing

`func (o *UserPolicy) HasEnablePublicSharing() bool`

HasEnablePublicSharing returns a boolean if a field has been set.

### GetRemoteClientBitrateLimit

`func (o *UserPolicy) GetRemoteClientBitrateLimit() int32`

GetRemoteClientBitrateLimit returns the RemoteClientBitrateLimit field if non-nil, zero value otherwise.

### GetRemoteClientBitrateLimitOk

`func (o *UserPolicy) GetRemoteClientBitrateLimitOk() (*int32, bool)`

GetRemoteClientBitrateLimitOk returns a tuple with the RemoteClientBitrateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClientBitrateLimit

`func (o *UserPolicy) SetRemoteClientBitrateLimit(v int32)`

SetRemoteClientBitrateLimit sets RemoteClientBitrateLimit field to given value.

### HasRemoteClientBitrateLimit

`func (o *UserPolicy) HasRemoteClientBitrateLimit() bool`

HasRemoteClientBitrateLimit returns a boolean if a field has been set.

### GetAuthenticationProviderId

`func (o *UserPolicy) GetAuthenticationProviderId() string`

GetAuthenticationProviderId returns the AuthenticationProviderId field if non-nil, zero value otherwise.

### GetAuthenticationProviderIdOk

`func (o *UserPolicy) GetAuthenticationProviderIdOk() (*string, bool)`

GetAuthenticationProviderIdOk returns a tuple with the AuthenticationProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProviderId

`func (o *UserPolicy) SetAuthenticationProviderId(v string)`

SetAuthenticationProviderId sets AuthenticationProviderId field to given value.

### HasAuthenticationProviderId

`func (o *UserPolicy) HasAuthenticationProviderId() bool`

HasAuthenticationProviderId returns a boolean if a field has been set.

### GetExcludedSubFolders

`func (o *UserPolicy) GetExcludedSubFolders() []string`

GetExcludedSubFolders returns the ExcludedSubFolders field if non-nil, zero value otherwise.

### GetExcludedSubFoldersOk

`func (o *UserPolicy) GetExcludedSubFoldersOk() (*[]string, bool)`

GetExcludedSubFoldersOk returns a tuple with the ExcludedSubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSubFolders

`func (o *UserPolicy) SetExcludedSubFolders(v []string)`

SetExcludedSubFolders sets ExcludedSubFolders field to given value.

### HasExcludedSubFolders

`func (o *UserPolicy) HasExcludedSubFolders() bool`

HasExcludedSubFolders returns a boolean if a field has been set.

### GetSimultaneousStreamLimit

`func (o *UserPolicy) GetSimultaneousStreamLimit() int32`

GetSimultaneousStreamLimit returns the SimultaneousStreamLimit field if non-nil, zero value otherwise.

### GetSimultaneousStreamLimitOk

`func (o *UserPolicy) GetSimultaneousStreamLimitOk() (*int32, bool)`

GetSimultaneousStreamLimitOk returns a tuple with the SimultaneousStreamLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimultaneousStreamLimit

`func (o *UserPolicy) SetSimultaneousStreamLimit(v int32)`

SetSimultaneousStreamLimit sets SimultaneousStreamLimit field to given value.

### HasSimultaneousStreamLimit

`func (o *UserPolicy) HasSimultaneousStreamLimit() bool`

HasSimultaneousStreamLimit returns a boolean if a field has been set.

### GetEnabledDevices

`func (o *UserPolicy) GetEnabledDevices() []string`

GetEnabledDevices returns the EnabledDevices field if non-nil, zero value otherwise.

### GetEnabledDevicesOk

`func (o *UserPolicy) GetEnabledDevicesOk() (*[]string, bool)`

GetEnabledDevicesOk returns a tuple with the EnabledDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledDevices

`func (o *UserPolicy) SetEnabledDevices(v []string)`

SetEnabledDevices sets EnabledDevices field to given value.

### HasEnabledDevices

`func (o *UserPolicy) HasEnabledDevices() bool`

HasEnabledDevices returns a boolean if a field has been set.

### GetEnableAllDevices

`func (o *UserPolicy) GetEnableAllDevices() bool`

GetEnableAllDevices returns the EnableAllDevices field if non-nil, zero value otherwise.

### GetEnableAllDevicesOk

`func (o *UserPolicy) GetEnableAllDevicesOk() (*bool, bool)`

GetEnableAllDevicesOk returns a tuple with the EnableAllDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllDevices

`func (o *UserPolicy) SetEnableAllDevices(v bool)`

SetEnableAllDevices sets EnableAllDevices field to given value.

### HasEnableAllDevices

`func (o *UserPolicy) HasEnableAllDevices() bool`

HasEnableAllDevices returns a boolean if a field has been set.

### GetAllowCameraUpload

`func (o *UserPolicy) GetAllowCameraUpload() bool`

GetAllowCameraUpload returns the AllowCameraUpload field if non-nil, zero value otherwise.

### GetAllowCameraUploadOk

`func (o *UserPolicy) GetAllowCameraUploadOk() (*bool, bool)`

GetAllowCameraUploadOk returns a tuple with the AllowCameraUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCameraUpload

`func (o *UserPolicy) SetAllowCameraUpload(v bool)`

SetAllowCameraUpload sets AllowCameraUpload field to given value.

### HasAllowCameraUpload

`func (o *UserPolicy) HasAllowCameraUpload() bool`

HasAllowCameraUpload returns a boolean if a field has been set.

### GetAllowSharingPersonalItems

`func (o *UserPolicy) GetAllowSharingPersonalItems() bool`

GetAllowSharingPersonalItems returns the AllowSharingPersonalItems field if non-nil, zero value otherwise.

### GetAllowSharingPersonalItemsOk

`func (o *UserPolicy) GetAllowSharingPersonalItemsOk() (*bool, bool)`

GetAllowSharingPersonalItemsOk returns a tuple with the AllowSharingPersonalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSharingPersonalItems

`func (o *UserPolicy) SetAllowSharingPersonalItems(v bool)`

SetAllowSharingPersonalItems sets AllowSharingPersonalItems field to given value.

### HasAllowSharingPersonalItems

`func (o *UserPolicy) HasAllowSharingPersonalItems() bool`

HasAllowSharingPersonalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


