# PackageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**IsPremium** | Pointer to **bool** |  | [optional] 
**Adult** | Pointer to **bool** |  | [optional] 
**RichDescUrl** | Pointer to **string** |  | [optional] 
**ThumbImage** | Pointer to **string** |  | [optional] 
**PreviewImage** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TargetFilename** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**TileColor** | Pointer to **string** |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**TargetSystem** | Pointer to [**PackageTargetSystem**](PackageTargetSystem.md) |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**IsRegistered** | Pointer to **bool** |  | [optional] 
**ExpDate** | Pointer to **time.Time** |  | [optional] 
**Versions** | Pointer to [**[]PackageVersionInfo**](PackageVersionInfo.md) |  | [optional] 
**EnableInAppStore** | Pointer to **bool** |  | [optional] 
**Installs** | Pointer to **int32** |  | [optional] 

## Methods

### NewPackageInfo

`func NewPackageInfo() *PackageInfo`

NewPackageInfo instantiates a new PackageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageInfoWithDefaults

`func NewPackageInfoWithDefaults() *PackageInfo`

NewPackageInfoWithDefaults instantiates a new PackageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PackageInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PackageInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PackageInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PackageInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PackageInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackageInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortDescription

`func (o *PackageInfo) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *PackageInfo) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *PackageInfo) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *PackageInfo) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetOverview

`func (o *PackageInfo) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *PackageInfo) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *PackageInfo) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *PackageInfo) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetIsPremium

`func (o *PackageInfo) GetIsPremium() bool`

GetIsPremium returns the IsPremium field if non-nil, zero value otherwise.

### GetIsPremiumOk

`func (o *PackageInfo) GetIsPremiumOk() (*bool, bool)`

GetIsPremiumOk returns a tuple with the IsPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremium

`func (o *PackageInfo) SetIsPremium(v bool)`

SetIsPremium sets IsPremium field to given value.

### HasIsPremium

`func (o *PackageInfo) HasIsPremium() bool`

HasIsPremium returns a boolean if a field has been set.

### GetAdult

`func (o *PackageInfo) GetAdult() bool`

GetAdult returns the Adult field if non-nil, zero value otherwise.

### GetAdultOk

`func (o *PackageInfo) GetAdultOk() (*bool, bool)`

GetAdultOk returns a tuple with the Adult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdult

`func (o *PackageInfo) SetAdult(v bool)`

SetAdult sets Adult field to given value.

### HasAdult

`func (o *PackageInfo) HasAdult() bool`

HasAdult returns a boolean if a field has been set.

### GetRichDescUrl

`func (o *PackageInfo) GetRichDescUrl() string`

GetRichDescUrl returns the RichDescUrl field if non-nil, zero value otherwise.

### GetRichDescUrlOk

`func (o *PackageInfo) GetRichDescUrlOk() (*string, bool)`

GetRichDescUrlOk returns a tuple with the RichDescUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichDescUrl

`func (o *PackageInfo) SetRichDescUrl(v string)`

SetRichDescUrl sets RichDescUrl field to given value.

### HasRichDescUrl

`func (o *PackageInfo) HasRichDescUrl() bool`

HasRichDescUrl returns a boolean if a field has been set.

### GetThumbImage

`func (o *PackageInfo) GetThumbImage() string`

GetThumbImage returns the ThumbImage field if non-nil, zero value otherwise.

### GetThumbImageOk

`func (o *PackageInfo) GetThumbImageOk() (*string, bool)`

GetThumbImageOk returns a tuple with the ThumbImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbImage

`func (o *PackageInfo) SetThumbImage(v string)`

SetThumbImage sets ThumbImage field to given value.

### HasThumbImage

`func (o *PackageInfo) HasThumbImage() bool`

HasThumbImage returns a boolean if a field has been set.

### GetPreviewImage

`func (o *PackageInfo) GetPreviewImage() string`

GetPreviewImage returns the PreviewImage field if non-nil, zero value otherwise.

### GetPreviewImageOk

`func (o *PackageInfo) GetPreviewImageOk() (*string, bool)`

GetPreviewImageOk returns a tuple with the PreviewImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImage

`func (o *PackageInfo) SetPreviewImage(v string)`

SetPreviewImage sets PreviewImage field to given value.

### HasPreviewImage

`func (o *PackageInfo) HasPreviewImage() bool`

HasPreviewImage returns a boolean if a field has been set.

### GetType

`func (o *PackageInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackageInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackageInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PackageInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTargetFilename

`func (o *PackageInfo) GetTargetFilename() string`

GetTargetFilename returns the TargetFilename field if non-nil, zero value otherwise.

### GetTargetFilenameOk

`func (o *PackageInfo) GetTargetFilenameOk() (*string, bool)`

GetTargetFilenameOk returns a tuple with the TargetFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFilename

`func (o *PackageInfo) SetTargetFilename(v string)`

SetTargetFilename sets TargetFilename field to given value.

### HasTargetFilename

`func (o *PackageInfo) HasTargetFilename() bool`

HasTargetFilename returns a boolean if a field has been set.

### GetOwner

`func (o *PackageInfo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PackageInfo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PackageInfo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PackageInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCategory

`func (o *PackageInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PackageInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PackageInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PackageInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTileColor

`func (o *PackageInfo) GetTileColor() string`

GetTileColor returns the TileColor field if non-nil, zero value otherwise.

### GetTileColorOk

`func (o *PackageInfo) GetTileColorOk() (*string, bool)`

GetTileColorOk returns a tuple with the TileColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTileColor

`func (o *PackageInfo) SetTileColor(v string)`

SetTileColor sets TileColor field to given value.

### HasTileColor

`func (o *PackageInfo) HasTileColor() bool`

HasTileColor returns a boolean if a field has been set.

### GetFeatureId

`func (o *PackageInfo) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *PackageInfo) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *PackageInfo) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *PackageInfo) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetPrice

`func (o *PackageInfo) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PackageInfo) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PackageInfo) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PackageInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *PackageInfo) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *PackageInfo) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetTargetSystem

`func (o *PackageInfo) GetTargetSystem() PackageTargetSystem`

GetTargetSystem returns the TargetSystem field if non-nil, zero value otherwise.

### GetTargetSystemOk

`func (o *PackageInfo) GetTargetSystemOk() (*PackageTargetSystem, bool)`

GetTargetSystemOk returns a tuple with the TargetSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystem

`func (o *PackageInfo) SetTargetSystem(v PackageTargetSystem)`

SetTargetSystem sets TargetSystem field to given value.

### HasTargetSystem

`func (o *PackageInfo) HasTargetSystem() bool`

HasTargetSystem returns a boolean if a field has been set.

### GetGuid

`func (o *PackageInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *PackageInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *PackageInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *PackageInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetIsRegistered

`func (o *PackageInfo) GetIsRegistered() bool`

GetIsRegistered returns the IsRegistered field if non-nil, zero value otherwise.

### GetIsRegisteredOk

`func (o *PackageInfo) GetIsRegisteredOk() (*bool, bool)`

GetIsRegisteredOk returns a tuple with the IsRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegistered

`func (o *PackageInfo) SetIsRegistered(v bool)`

SetIsRegistered sets IsRegistered field to given value.

### HasIsRegistered

`func (o *PackageInfo) HasIsRegistered() bool`

HasIsRegistered returns a boolean if a field has been set.

### GetExpDate

`func (o *PackageInfo) GetExpDate() time.Time`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *PackageInfo) GetExpDateOk() (*time.Time, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *PackageInfo) SetExpDate(v time.Time)`

SetExpDate sets ExpDate field to given value.

### HasExpDate

`func (o *PackageInfo) HasExpDate() bool`

HasExpDate returns a boolean if a field has been set.

### GetVersions

`func (o *PackageInfo) GetVersions() []PackageVersionInfo`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PackageInfo) GetVersionsOk() (*[]PackageVersionInfo, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PackageInfo) SetVersions(v []PackageVersionInfo)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PackageInfo) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetEnableInAppStore

`func (o *PackageInfo) GetEnableInAppStore() bool`

GetEnableInAppStore returns the EnableInAppStore field if non-nil, zero value otherwise.

### GetEnableInAppStoreOk

`func (o *PackageInfo) GetEnableInAppStoreOk() (*bool, bool)`

GetEnableInAppStoreOk returns a tuple with the EnableInAppStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInAppStore

`func (o *PackageInfo) SetEnableInAppStore(v bool)`

SetEnableInAppStore sets EnableInAppStore field to given value.

### HasEnableInAppStore

`func (o *PackageInfo) HasEnableInAppStore() bool`

HasEnableInAppStore returns a boolean if a field has been set.

### GetInstalls

`func (o *PackageInfo) GetInstalls() int32`

GetInstalls returns the Installs field if non-nil, zero value otherwise.

### GetInstallsOk

`func (o *PackageInfo) GetInstallsOk() (*int32, bool)`

GetInstallsOk returns a tuple with the Installs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalls

`func (o *PackageInfo) SetInstalls(v int32)`

SetInstalls sets Installs field to given value.

### HasInstalls

`func (o *PackageInfo) HasInstalls() bool`

HasInstalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


