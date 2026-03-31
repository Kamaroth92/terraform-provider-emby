# DlnaProfilesDlnaProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DlnaProfilesDeviceProfileType**](DlnaProfilesDeviceProfileType.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AlbumArtPn** | Pointer to **string** |  | [optional] 
**MaxAlbumArtWidth** | Pointer to **int32** |  | [optional] 
**MaxAlbumArtHeight** | Pointer to **int32** |  | [optional] 
**MaxIconWidth** | Pointer to **NullableInt32** |  | [optional] 
**MaxIconHeight** | Pointer to **NullableInt32** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**ManufacturerUrl** | Pointer to **string** |  | [optional] 
**ModelName** | Pointer to **string** |  | [optional] 
**ModelDescription** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**ModelUrl** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**EnableAlbumArtInDidl** | Pointer to **bool** |  | [optional] 
**EnableSingleAlbumArtLimit** | Pointer to **bool** |  | [optional] 
**EnableSingleSubtitleLimit** | Pointer to **bool** |  | [optional] 
**ProtocolInfo** | Pointer to **string** |  | [optional] 
**TimelineOffsetSeconds** | Pointer to **int32** |  | [optional] 
**RequiresPlainVideoItems** | Pointer to **bool** |  | [optional] 
**RequiresPlainFolders** | Pointer to **bool** |  | [optional] 
**IgnoreTranscodeByteRangeRequests** | Pointer to **bool** |  | [optional] 
**SupportsSamsungBookmark** | Pointer to **bool** |  | [optional] 
**Identification** | Pointer to [**DlnaProfilesDeviceIdentification**](DlnaProfilesDeviceIdentification.md) |  | [optional] 
**ProtocolInfoDetection** | Pointer to [**DlnaProfilesProtocolInfoDetection**](DlnaProfilesProtocolInfoDetection.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SupportedMediaTypes** | Pointer to **string** |  | [optional] 
**MaxStreamingBitrate** | Pointer to **NullableInt64** |  | [optional] 
**MusicStreamingTranscodingBitrate** | Pointer to **NullableInt32** |  | [optional] 
**MaxStaticMusicBitrate** | Pointer to **NullableInt32** |  | [optional] 
**DeclaredFeatures** | Pointer to **[]string** |  | [optional] 
**DirectPlayProfiles** | Pointer to [**[]DirectPlayProfile**](DirectPlayProfile.md) |  | [optional] 
**TranscodingProfiles** | Pointer to [**[]TranscodingProfile**](TranscodingProfile.md) |  | [optional] 
**ContainerProfiles** | Pointer to [**[]ContainerProfile**](ContainerProfile.md) |  | [optional] 
**CodecProfiles** | Pointer to [**[]CodecProfile**](CodecProfile.md) |  | [optional] 
**ResponseProfiles** | Pointer to [**[]ResponseProfile**](ResponseProfile.md) |  | [optional] 
**SubtitleProfiles** | Pointer to [**[]SubtitleProfile**](SubtitleProfile.md) |  | [optional] 

## Methods

### NewDlnaProfilesDlnaProfile

`func NewDlnaProfilesDlnaProfile() *DlnaProfilesDlnaProfile`

NewDlnaProfilesDlnaProfile instantiates a new DlnaProfilesDlnaProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlnaProfilesDlnaProfileWithDefaults

`func NewDlnaProfilesDlnaProfileWithDefaults() *DlnaProfilesDlnaProfile`

NewDlnaProfilesDlnaProfileWithDefaults instantiates a new DlnaProfilesDlnaProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DlnaProfilesDlnaProfile) GetType() DlnaProfilesDeviceProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DlnaProfilesDlnaProfile) GetTypeOk() (*DlnaProfilesDeviceProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DlnaProfilesDlnaProfile) SetType(v DlnaProfilesDeviceProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *DlnaProfilesDlnaProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPath

`func (o *DlnaProfilesDlnaProfile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DlnaProfilesDlnaProfile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DlnaProfilesDlnaProfile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DlnaProfilesDlnaProfile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUserId

`func (o *DlnaProfilesDlnaProfile) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DlnaProfilesDlnaProfile) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DlnaProfilesDlnaProfile) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DlnaProfilesDlnaProfile) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAlbumArtPn

`func (o *DlnaProfilesDlnaProfile) GetAlbumArtPn() string`

GetAlbumArtPn returns the AlbumArtPn field if non-nil, zero value otherwise.

### GetAlbumArtPnOk

`func (o *DlnaProfilesDlnaProfile) GetAlbumArtPnOk() (*string, bool)`

GetAlbumArtPnOk returns a tuple with the AlbumArtPn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtPn

`func (o *DlnaProfilesDlnaProfile) SetAlbumArtPn(v string)`

SetAlbumArtPn sets AlbumArtPn field to given value.

### HasAlbumArtPn

`func (o *DlnaProfilesDlnaProfile) HasAlbumArtPn() bool`

HasAlbumArtPn returns a boolean if a field has been set.

### GetMaxAlbumArtWidth

`func (o *DlnaProfilesDlnaProfile) GetMaxAlbumArtWidth() int32`

GetMaxAlbumArtWidth returns the MaxAlbumArtWidth field if non-nil, zero value otherwise.

### GetMaxAlbumArtWidthOk

`func (o *DlnaProfilesDlnaProfile) GetMaxAlbumArtWidthOk() (*int32, bool)`

GetMaxAlbumArtWidthOk returns a tuple with the MaxAlbumArtWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlbumArtWidth

`func (o *DlnaProfilesDlnaProfile) SetMaxAlbumArtWidth(v int32)`

SetMaxAlbumArtWidth sets MaxAlbumArtWidth field to given value.

### HasMaxAlbumArtWidth

`func (o *DlnaProfilesDlnaProfile) HasMaxAlbumArtWidth() bool`

HasMaxAlbumArtWidth returns a boolean if a field has been set.

### GetMaxAlbumArtHeight

`func (o *DlnaProfilesDlnaProfile) GetMaxAlbumArtHeight() int32`

GetMaxAlbumArtHeight returns the MaxAlbumArtHeight field if non-nil, zero value otherwise.

### GetMaxAlbumArtHeightOk

`func (o *DlnaProfilesDlnaProfile) GetMaxAlbumArtHeightOk() (*int32, bool)`

GetMaxAlbumArtHeightOk returns a tuple with the MaxAlbumArtHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlbumArtHeight

`func (o *DlnaProfilesDlnaProfile) SetMaxAlbumArtHeight(v int32)`

SetMaxAlbumArtHeight sets MaxAlbumArtHeight field to given value.

### HasMaxAlbumArtHeight

`func (o *DlnaProfilesDlnaProfile) HasMaxAlbumArtHeight() bool`

HasMaxAlbumArtHeight returns a boolean if a field has been set.

### GetMaxIconWidth

`func (o *DlnaProfilesDlnaProfile) GetMaxIconWidth() int32`

GetMaxIconWidth returns the MaxIconWidth field if non-nil, zero value otherwise.

### GetMaxIconWidthOk

`func (o *DlnaProfilesDlnaProfile) GetMaxIconWidthOk() (*int32, bool)`

GetMaxIconWidthOk returns a tuple with the MaxIconWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIconWidth

`func (o *DlnaProfilesDlnaProfile) SetMaxIconWidth(v int32)`

SetMaxIconWidth sets MaxIconWidth field to given value.

### HasMaxIconWidth

`func (o *DlnaProfilesDlnaProfile) HasMaxIconWidth() bool`

HasMaxIconWidth returns a boolean if a field has been set.

### SetMaxIconWidthNil

`func (o *DlnaProfilesDlnaProfile) SetMaxIconWidthNil(b bool)`

 SetMaxIconWidthNil sets the value for MaxIconWidth to be an explicit nil

### UnsetMaxIconWidth
`func (o *DlnaProfilesDlnaProfile) UnsetMaxIconWidth()`

UnsetMaxIconWidth ensures that no value is present for MaxIconWidth, not even an explicit nil
### GetMaxIconHeight

`func (o *DlnaProfilesDlnaProfile) GetMaxIconHeight() int32`

GetMaxIconHeight returns the MaxIconHeight field if non-nil, zero value otherwise.

### GetMaxIconHeightOk

`func (o *DlnaProfilesDlnaProfile) GetMaxIconHeightOk() (*int32, bool)`

GetMaxIconHeightOk returns a tuple with the MaxIconHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIconHeight

`func (o *DlnaProfilesDlnaProfile) SetMaxIconHeight(v int32)`

SetMaxIconHeight sets MaxIconHeight field to given value.

### HasMaxIconHeight

`func (o *DlnaProfilesDlnaProfile) HasMaxIconHeight() bool`

HasMaxIconHeight returns a boolean if a field has been set.

### SetMaxIconHeightNil

`func (o *DlnaProfilesDlnaProfile) SetMaxIconHeightNil(b bool)`

 SetMaxIconHeightNil sets the value for MaxIconHeight to be an explicit nil

### UnsetMaxIconHeight
`func (o *DlnaProfilesDlnaProfile) UnsetMaxIconHeight()`

UnsetMaxIconHeight ensures that no value is present for MaxIconHeight, not even an explicit nil
### GetFriendlyName

`func (o *DlnaProfilesDlnaProfile) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *DlnaProfilesDlnaProfile) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *DlnaProfilesDlnaProfile) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *DlnaProfilesDlnaProfile) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetManufacturer

`func (o *DlnaProfilesDlnaProfile) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *DlnaProfilesDlnaProfile) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *DlnaProfilesDlnaProfile) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *DlnaProfilesDlnaProfile) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetManufacturerUrl

`func (o *DlnaProfilesDlnaProfile) GetManufacturerUrl() string`

GetManufacturerUrl returns the ManufacturerUrl field if non-nil, zero value otherwise.

### GetManufacturerUrlOk

`func (o *DlnaProfilesDlnaProfile) GetManufacturerUrlOk() (*string, bool)`

GetManufacturerUrlOk returns a tuple with the ManufacturerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerUrl

`func (o *DlnaProfilesDlnaProfile) SetManufacturerUrl(v string)`

SetManufacturerUrl sets ManufacturerUrl field to given value.

### HasManufacturerUrl

`func (o *DlnaProfilesDlnaProfile) HasManufacturerUrl() bool`

HasManufacturerUrl returns a boolean if a field has been set.

### GetModelName

`func (o *DlnaProfilesDlnaProfile) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *DlnaProfilesDlnaProfile) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *DlnaProfilesDlnaProfile) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *DlnaProfilesDlnaProfile) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModelDescription

`func (o *DlnaProfilesDlnaProfile) GetModelDescription() string`

GetModelDescription returns the ModelDescription field if non-nil, zero value otherwise.

### GetModelDescriptionOk

`func (o *DlnaProfilesDlnaProfile) GetModelDescriptionOk() (*string, bool)`

GetModelDescriptionOk returns a tuple with the ModelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelDescription

`func (o *DlnaProfilesDlnaProfile) SetModelDescription(v string)`

SetModelDescription sets ModelDescription field to given value.

### HasModelDescription

`func (o *DlnaProfilesDlnaProfile) HasModelDescription() bool`

HasModelDescription returns a boolean if a field has been set.

### GetModelNumber

`func (o *DlnaProfilesDlnaProfile) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *DlnaProfilesDlnaProfile) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *DlnaProfilesDlnaProfile) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *DlnaProfilesDlnaProfile) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetModelUrl

`func (o *DlnaProfilesDlnaProfile) GetModelUrl() string`

GetModelUrl returns the ModelUrl field if non-nil, zero value otherwise.

### GetModelUrlOk

`func (o *DlnaProfilesDlnaProfile) GetModelUrlOk() (*string, bool)`

GetModelUrlOk returns a tuple with the ModelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelUrl

`func (o *DlnaProfilesDlnaProfile) SetModelUrl(v string)`

SetModelUrl sets ModelUrl field to given value.

### HasModelUrl

`func (o *DlnaProfilesDlnaProfile) HasModelUrl() bool`

HasModelUrl returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DlnaProfilesDlnaProfile) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DlnaProfilesDlnaProfile) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DlnaProfilesDlnaProfile) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DlnaProfilesDlnaProfile) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetEnableAlbumArtInDidl

`func (o *DlnaProfilesDlnaProfile) GetEnableAlbumArtInDidl() bool`

GetEnableAlbumArtInDidl returns the EnableAlbumArtInDidl field if non-nil, zero value otherwise.

### GetEnableAlbumArtInDidlOk

`func (o *DlnaProfilesDlnaProfile) GetEnableAlbumArtInDidlOk() (*bool, bool)`

GetEnableAlbumArtInDidlOk returns a tuple with the EnableAlbumArtInDidl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAlbumArtInDidl

`func (o *DlnaProfilesDlnaProfile) SetEnableAlbumArtInDidl(v bool)`

SetEnableAlbumArtInDidl sets EnableAlbumArtInDidl field to given value.

### HasEnableAlbumArtInDidl

`func (o *DlnaProfilesDlnaProfile) HasEnableAlbumArtInDidl() bool`

HasEnableAlbumArtInDidl returns a boolean if a field has been set.

### GetEnableSingleAlbumArtLimit

`func (o *DlnaProfilesDlnaProfile) GetEnableSingleAlbumArtLimit() bool`

GetEnableSingleAlbumArtLimit returns the EnableSingleAlbumArtLimit field if non-nil, zero value otherwise.

### GetEnableSingleAlbumArtLimitOk

`func (o *DlnaProfilesDlnaProfile) GetEnableSingleAlbumArtLimitOk() (*bool, bool)`

GetEnableSingleAlbumArtLimitOk returns a tuple with the EnableSingleAlbumArtLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleAlbumArtLimit

`func (o *DlnaProfilesDlnaProfile) SetEnableSingleAlbumArtLimit(v bool)`

SetEnableSingleAlbumArtLimit sets EnableSingleAlbumArtLimit field to given value.

### HasEnableSingleAlbumArtLimit

`func (o *DlnaProfilesDlnaProfile) HasEnableSingleAlbumArtLimit() bool`

HasEnableSingleAlbumArtLimit returns a boolean if a field has been set.

### GetEnableSingleSubtitleLimit

`func (o *DlnaProfilesDlnaProfile) GetEnableSingleSubtitleLimit() bool`

GetEnableSingleSubtitleLimit returns the EnableSingleSubtitleLimit field if non-nil, zero value otherwise.

### GetEnableSingleSubtitleLimitOk

`func (o *DlnaProfilesDlnaProfile) GetEnableSingleSubtitleLimitOk() (*bool, bool)`

GetEnableSingleSubtitleLimitOk returns a tuple with the EnableSingleSubtitleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleSubtitleLimit

`func (o *DlnaProfilesDlnaProfile) SetEnableSingleSubtitleLimit(v bool)`

SetEnableSingleSubtitleLimit sets EnableSingleSubtitleLimit field to given value.

### HasEnableSingleSubtitleLimit

`func (o *DlnaProfilesDlnaProfile) HasEnableSingleSubtitleLimit() bool`

HasEnableSingleSubtitleLimit returns a boolean if a field has been set.

### GetProtocolInfo

`func (o *DlnaProfilesDlnaProfile) GetProtocolInfo() string`

GetProtocolInfo returns the ProtocolInfo field if non-nil, zero value otherwise.

### GetProtocolInfoOk

`func (o *DlnaProfilesDlnaProfile) GetProtocolInfoOk() (*string, bool)`

GetProtocolInfoOk returns a tuple with the ProtocolInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolInfo

`func (o *DlnaProfilesDlnaProfile) SetProtocolInfo(v string)`

SetProtocolInfo sets ProtocolInfo field to given value.

### HasProtocolInfo

`func (o *DlnaProfilesDlnaProfile) HasProtocolInfo() bool`

HasProtocolInfo returns a boolean if a field has been set.

### GetTimelineOffsetSeconds

`func (o *DlnaProfilesDlnaProfile) GetTimelineOffsetSeconds() int32`

GetTimelineOffsetSeconds returns the TimelineOffsetSeconds field if non-nil, zero value otherwise.

### GetTimelineOffsetSecondsOk

`func (o *DlnaProfilesDlnaProfile) GetTimelineOffsetSecondsOk() (*int32, bool)`

GetTimelineOffsetSecondsOk returns a tuple with the TimelineOffsetSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineOffsetSeconds

`func (o *DlnaProfilesDlnaProfile) SetTimelineOffsetSeconds(v int32)`

SetTimelineOffsetSeconds sets TimelineOffsetSeconds field to given value.

### HasTimelineOffsetSeconds

`func (o *DlnaProfilesDlnaProfile) HasTimelineOffsetSeconds() bool`

HasTimelineOffsetSeconds returns a boolean if a field has been set.

### GetRequiresPlainVideoItems

`func (o *DlnaProfilesDlnaProfile) GetRequiresPlainVideoItems() bool`

GetRequiresPlainVideoItems returns the RequiresPlainVideoItems field if non-nil, zero value otherwise.

### GetRequiresPlainVideoItemsOk

`func (o *DlnaProfilesDlnaProfile) GetRequiresPlainVideoItemsOk() (*bool, bool)`

GetRequiresPlainVideoItemsOk returns a tuple with the RequiresPlainVideoItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPlainVideoItems

`func (o *DlnaProfilesDlnaProfile) SetRequiresPlainVideoItems(v bool)`

SetRequiresPlainVideoItems sets RequiresPlainVideoItems field to given value.

### HasRequiresPlainVideoItems

`func (o *DlnaProfilesDlnaProfile) HasRequiresPlainVideoItems() bool`

HasRequiresPlainVideoItems returns a boolean if a field has been set.

### GetRequiresPlainFolders

`func (o *DlnaProfilesDlnaProfile) GetRequiresPlainFolders() bool`

GetRequiresPlainFolders returns the RequiresPlainFolders field if non-nil, zero value otherwise.

### GetRequiresPlainFoldersOk

`func (o *DlnaProfilesDlnaProfile) GetRequiresPlainFoldersOk() (*bool, bool)`

GetRequiresPlainFoldersOk returns a tuple with the RequiresPlainFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPlainFolders

`func (o *DlnaProfilesDlnaProfile) SetRequiresPlainFolders(v bool)`

SetRequiresPlainFolders sets RequiresPlainFolders field to given value.

### HasRequiresPlainFolders

`func (o *DlnaProfilesDlnaProfile) HasRequiresPlainFolders() bool`

HasRequiresPlainFolders returns a boolean if a field has been set.

### GetIgnoreTranscodeByteRangeRequests

`func (o *DlnaProfilesDlnaProfile) GetIgnoreTranscodeByteRangeRequests() bool`

GetIgnoreTranscodeByteRangeRequests returns the IgnoreTranscodeByteRangeRequests field if non-nil, zero value otherwise.

### GetIgnoreTranscodeByteRangeRequestsOk

`func (o *DlnaProfilesDlnaProfile) GetIgnoreTranscodeByteRangeRequestsOk() (*bool, bool)`

GetIgnoreTranscodeByteRangeRequestsOk returns a tuple with the IgnoreTranscodeByteRangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTranscodeByteRangeRequests

`func (o *DlnaProfilesDlnaProfile) SetIgnoreTranscodeByteRangeRequests(v bool)`

SetIgnoreTranscodeByteRangeRequests sets IgnoreTranscodeByteRangeRequests field to given value.

### HasIgnoreTranscodeByteRangeRequests

`func (o *DlnaProfilesDlnaProfile) HasIgnoreTranscodeByteRangeRequests() bool`

HasIgnoreTranscodeByteRangeRequests returns a boolean if a field has been set.

### GetSupportsSamsungBookmark

`func (o *DlnaProfilesDlnaProfile) GetSupportsSamsungBookmark() bool`

GetSupportsSamsungBookmark returns the SupportsSamsungBookmark field if non-nil, zero value otherwise.

### GetSupportsSamsungBookmarkOk

`func (o *DlnaProfilesDlnaProfile) GetSupportsSamsungBookmarkOk() (*bool, bool)`

GetSupportsSamsungBookmarkOk returns a tuple with the SupportsSamsungBookmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSamsungBookmark

`func (o *DlnaProfilesDlnaProfile) SetSupportsSamsungBookmark(v bool)`

SetSupportsSamsungBookmark sets SupportsSamsungBookmark field to given value.

### HasSupportsSamsungBookmark

`func (o *DlnaProfilesDlnaProfile) HasSupportsSamsungBookmark() bool`

HasSupportsSamsungBookmark returns a boolean if a field has been set.

### GetIdentification

`func (o *DlnaProfilesDlnaProfile) GetIdentification() DlnaProfilesDeviceIdentification`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *DlnaProfilesDlnaProfile) GetIdentificationOk() (*DlnaProfilesDeviceIdentification, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *DlnaProfilesDlnaProfile) SetIdentification(v DlnaProfilesDeviceIdentification)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *DlnaProfilesDlnaProfile) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### GetProtocolInfoDetection

`func (o *DlnaProfilesDlnaProfile) GetProtocolInfoDetection() DlnaProfilesProtocolInfoDetection`

GetProtocolInfoDetection returns the ProtocolInfoDetection field if non-nil, zero value otherwise.

### GetProtocolInfoDetectionOk

`func (o *DlnaProfilesDlnaProfile) GetProtocolInfoDetectionOk() (*DlnaProfilesProtocolInfoDetection, bool)`

GetProtocolInfoDetectionOk returns a tuple with the ProtocolInfoDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolInfoDetection

`func (o *DlnaProfilesDlnaProfile) SetProtocolInfoDetection(v DlnaProfilesProtocolInfoDetection)`

SetProtocolInfoDetection sets ProtocolInfoDetection field to given value.

### HasProtocolInfoDetection

`func (o *DlnaProfilesDlnaProfile) HasProtocolInfoDetection() bool`

HasProtocolInfoDetection returns a boolean if a field has been set.

### GetName

`func (o *DlnaProfilesDlnaProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DlnaProfilesDlnaProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DlnaProfilesDlnaProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DlnaProfilesDlnaProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *DlnaProfilesDlnaProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DlnaProfilesDlnaProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DlnaProfilesDlnaProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DlnaProfilesDlnaProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSupportedMediaTypes

`func (o *DlnaProfilesDlnaProfile) GetSupportedMediaTypes() string`

GetSupportedMediaTypes returns the SupportedMediaTypes field if non-nil, zero value otherwise.

### GetSupportedMediaTypesOk

`func (o *DlnaProfilesDlnaProfile) GetSupportedMediaTypesOk() (*string, bool)`

GetSupportedMediaTypesOk returns a tuple with the SupportedMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMediaTypes

`func (o *DlnaProfilesDlnaProfile) SetSupportedMediaTypes(v string)`

SetSupportedMediaTypes sets SupportedMediaTypes field to given value.

### HasSupportedMediaTypes

`func (o *DlnaProfilesDlnaProfile) HasSupportedMediaTypes() bool`

HasSupportedMediaTypes returns a boolean if a field has been set.

### GetMaxStreamingBitrate

`func (o *DlnaProfilesDlnaProfile) GetMaxStreamingBitrate() int64`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *DlnaProfilesDlnaProfile) GetMaxStreamingBitrateOk() (*int64, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *DlnaProfilesDlnaProfile) SetMaxStreamingBitrate(v int64)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *DlnaProfilesDlnaProfile) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *DlnaProfilesDlnaProfile) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *DlnaProfilesDlnaProfile) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetMusicStreamingTranscodingBitrate

`func (o *DlnaProfilesDlnaProfile) GetMusicStreamingTranscodingBitrate() int32`

GetMusicStreamingTranscodingBitrate returns the MusicStreamingTranscodingBitrate field if non-nil, zero value otherwise.

### GetMusicStreamingTranscodingBitrateOk

`func (o *DlnaProfilesDlnaProfile) GetMusicStreamingTranscodingBitrateOk() (*int32, bool)`

GetMusicStreamingTranscodingBitrateOk returns a tuple with the MusicStreamingTranscodingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicStreamingTranscodingBitrate

`func (o *DlnaProfilesDlnaProfile) SetMusicStreamingTranscodingBitrate(v int32)`

SetMusicStreamingTranscodingBitrate sets MusicStreamingTranscodingBitrate field to given value.

### HasMusicStreamingTranscodingBitrate

`func (o *DlnaProfilesDlnaProfile) HasMusicStreamingTranscodingBitrate() bool`

HasMusicStreamingTranscodingBitrate returns a boolean if a field has been set.

### SetMusicStreamingTranscodingBitrateNil

`func (o *DlnaProfilesDlnaProfile) SetMusicStreamingTranscodingBitrateNil(b bool)`

 SetMusicStreamingTranscodingBitrateNil sets the value for MusicStreamingTranscodingBitrate to be an explicit nil

### UnsetMusicStreamingTranscodingBitrate
`func (o *DlnaProfilesDlnaProfile) UnsetMusicStreamingTranscodingBitrate()`

UnsetMusicStreamingTranscodingBitrate ensures that no value is present for MusicStreamingTranscodingBitrate, not even an explicit nil
### GetMaxStaticMusicBitrate

`func (o *DlnaProfilesDlnaProfile) GetMaxStaticMusicBitrate() int32`

GetMaxStaticMusicBitrate returns the MaxStaticMusicBitrate field if non-nil, zero value otherwise.

### GetMaxStaticMusicBitrateOk

`func (o *DlnaProfilesDlnaProfile) GetMaxStaticMusicBitrateOk() (*int32, bool)`

GetMaxStaticMusicBitrateOk returns a tuple with the MaxStaticMusicBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaticMusicBitrate

`func (o *DlnaProfilesDlnaProfile) SetMaxStaticMusicBitrate(v int32)`

SetMaxStaticMusicBitrate sets MaxStaticMusicBitrate field to given value.

### HasMaxStaticMusicBitrate

`func (o *DlnaProfilesDlnaProfile) HasMaxStaticMusicBitrate() bool`

HasMaxStaticMusicBitrate returns a boolean if a field has been set.

### SetMaxStaticMusicBitrateNil

`func (o *DlnaProfilesDlnaProfile) SetMaxStaticMusicBitrateNil(b bool)`

 SetMaxStaticMusicBitrateNil sets the value for MaxStaticMusicBitrate to be an explicit nil

### UnsetMaxStaticMusicBitrate
`func (o *DlnaProfilesDlnaProfile) UnsetMaxStaticMusicBitrate()`

UnsetMaxStaticMusicBitrate ensures that no value is present for MaxStaticMusicBitrate, not even an explicit nil
### GetDeclaredFeatures

`func (o *DlnaProfilesDlnaProfile) GetDeclaredFeatures() []string`

GetDeclaredFeatures returns the DeclaredFeatures field if non-nil, zero value otherwise.

### GetDeclaredFeaturesOk

`func (o *DlnaProfilesDlnaProfile) GetDeclaredFeaturesOk() (*[]string, bool)`

GetDeclaredFeaturesOk returns a tuple with the DeclaredFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredFeatures

`func (o *DlnaProfilesDlnaProfile) SetDeclaredFeatures(v []string)`

SetDeclaredFeatures sets DeclaredFeatures field to given value.

### HasDeclaredFeatures

`func (o *DlnaProfilesDlnaProfile) HasDeclaredFeatures() bool`

HasDeclaredFeatures returns a boolean if a field has been set.

### GetDirectPlayProfiles

`func (o *DlnaProfilesDlnaProfile) GetDirectPlayProfiles() []DirectPlayProfile`

GetDirectPlayProfiles returns the DirectPlayProfiles field if non-nil, zero value otherwise.

### GetDirectPlayProfilesOk

`func (o *DlnaProfilesDlnaProfile) GetDirectPlayProfilesOk() (*[]DirectPlayProfile, bool)`

GetDirectPlayProfilesOk returns a tuple with the DirectPlayProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProfiles

`func (o *DlnaProfilesDlnaProfile) SetDirectPlayProfiles(v []DirectPlayProfile)`

SetDirectPlayProfiles sets DirectPlayProfiles field to given value.

### HasDirectPlayProfiles

`func (o *DlnaProfilesDlnaProfile) HasDirectPlayProfiles() bool`

HasDirectPlayProfiles returns a boolean if a field has been set.

### GetTranscodingProfiles

`func (o *DlnaProfilesDlnaProfile) GetTranscodingProfiles() []TranscodingProfile`

GetTranscodingProfiles returns the TranscodingProfiles field if non-nil, zero value otherwise.

### GetTranscodingProfilesOk

`func (o *DlnaProfilesDlnaProfile) GetTranscodingProfilesOk() (*[]TranscodingProfile, bool)`

GetTranscodingProfilesOk returns a tuple with the TranscodingProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingProfiles

`func (o *DlnaProfilesDlnaProfile) SetTranscodingProfiles(v []TranscodingProfile)`

SetTranscodingProfiles sets TranscodingProfiles field to given value.

### HasTranscodingProfiles

`func (o *DlnaProfilesDlnaProfile) HasTranscodingProfiles() bool`

HasTranscodingProfiles returns a boolean if a field has been set.

### GetContainerProfiles

`func (o *DlnaProfilesDlnaProfile) GetContainerProfiles() []ContainerProfile`

GetContainerProfiles returns the ContainerProfiles field if non-nil, zero value otherwise.

### GetContainerProfilesOk

`func (o *DlnaProfilesDlnaProfile) GetContainerProfilesOk() (*[]ContainerProfile, bool)`

GetContainerProfilesOk returns a tuple with the ContainerProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerProfiles

`func (o *DlnaProfilesDlnaProfile) SetContainerProfiles(v []ContainerProfile)`

SetContainerProfiles sets ContainerProfiles field to given value.

### HasContainerProfiles

`func (o *DlnaProfilesDlnaProfile) HasContainerProfiles() bool`

HasContainerProfiles returns a boolean if a field has been set.

### GetCodecProfiles

`func (o *DlnaProfilesDlnaProfile) GetCodecProfiles() []CodecProfile`

GetCodecProfiles returns the CodecProfiles field if non-nil, zero value otherwise.

### GetCodecProfilesOk

`func (o *DlnaProfilesDlnaProfile) GetCodecProfilesOk() (*[]CodecProfile, bool)`

GetCodecProfilesOk returns a tuple with the CodecProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecProfiles

`func (o *DlnaProfilesDlnaProfile) SetCodecProfiles(v []CodecProfile)`

SetCodecProfiles sets CodecProfiles field to given value.

### HasCodecProfiles

`func (o *DlnaProfilesDlnaProfile) HasCodecProfiles() bool`

HasCodecProfiles returns a boolean if a field has been set.

### GetResponseProfiles

`func (o *DlnaProfilesDlnaProfile) GetResponseProfiles() []ResponseProfile`

GetResponseProfiles returns the ResponseProfiles field if non-nil, zero value otherwise.

### GetResponseProfilesOk

`func (o *DlnaProfilesDlnaProfile) GetResponseProfilesOk() (*[]ResponseProfile, bool)`

GetResponseProfilesOk returns a tuple with the ResponseProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseProfiles

`func (o *DlnaProfilesDlnaProfile) SetResponseProfiles(v []ResponseProfile)`

SetResponseProfiles sets ResponseProfiles field to given value.

### HasResponseProfiles

`func (o *DlnaProfilesDlnaProfile) HasResponseProfiles() bool`

HasResponseProfiles returns a boolean if a field has been set.

### GetSubtitleProfiles

`func (o *DlnaProfilesDlnaProfile) GetSubtitleProfiles() []SubtitleProfile`

GetSubtitleProfiles returns the SubtitleProfiles field if non-nil, zero value otherwise.

### GetSubtitleProfilesOk

`func (o *DlnaProfilesDlnaProfile) GetSubtitleProfilesOk() (*[]SubtitleProfile, bool)`

GetSubtitleProfilesOk returns a tuple with the SubtitleProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleProfiles

`func (o *DlnaProfilesDlnaProfile) SetSubtitleProfiles(v []SubtitleProfile)`

SetSubtitleProfiles sets SubtitleProfiles field to given value.

### HasSubtitleProfiles

`func (o *DlnaProfilesDlnaProfile) HasSubtitleProfiles() bool`

HasSubtitleProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


