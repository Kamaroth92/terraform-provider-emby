# LibraryTypeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**MetadataFetchers** | Pointer to [**[]LibraryOptionInfo**](LibraryOptionInfo.md) |  | [optional] 
**ImageFetchers** | Pointer to [**[]LibraryOptionInfo**](LibraryOptionInfo.md) |  | [optional] 
**SupportedImageTypes** | Pointer to [**[]ImageType**](ImageType.md) |  | [optional] 
**DefaultImageOptions** | Pointer to [**[]ImageOption**](ImageOption.md) |  | [optional] 

## Methods

### NewLibraryTypeOptions

`func NewLibraryTypeOptions() *LibraryTypeOptions`

NewLibraryTypeOptions instantiates a new LibraryTypeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryTypeOptionsWithDefaults

`func NewLibraryTypeOptionsWithDefaults() *LibraryTypeOptions`

NewLibraryTypeOptionsWithDefaults instantiates a new LibraryTypeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LibraryTypeOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LibraryTypeOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LibraryTypeOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LibraryTypeOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetadataFetchers

`func (o *LibraryTypeOptions) GetMetadataFetchers() []LibraryOptionInfo`

GetMetadataFetchers returns the MetadataFetchers field if non-nil, zero value otherwise.

### GetMetadataFetchersOk

`func (o *LibraryTypeOptions) GetMetadataFetchersOk() (*[]LibraryOptionInfo, bool)`

GetMetadataFetchersOk returns a tuple with the MetadataFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetchers

`func (o *LibraryTypeOptions) SetMetadataFetchers(v []LibraryOptionInfo)`

SetMetadataFetchers sets MetadataFetchers field to given value.

### HasMetadataFetchers

`func (o *LibraryTypeOptions) HasMetadataFetchers() bool`

HasMetadataFetchers returns a boolean if a field has been set.

### GetImageFetchers

`func (o *LibraryTypeOptions) GetImageFetchers() []LibraryOptionInfo`

GetImageFetchers returns the ImageFetchers field if non-nil, zero value otherwise.

### GetImageFetchersOk

`func (o *LibraryTypeOptions) GetImageFetchersOk() (*[]LibraryOptionInfo, bool)`

GetImageFetchersOk returns a tuple with the ImageFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetchers

`func (o *LibraryTypeOptions) SetImageFetchers(v []LibraryOptionInfo)`

SetImageFetchers sets ImageFetchers field to given value.

### HasImageFetchers

`func (o *LibraryTypeOptions) HasImageFetchers() bool`

HasImageFetchers returns a boolean if a field has been set.

### GetSupportedImageTypes

`func (o *LibraryTypeOptions) GetSupportedImageTypes() []ImageType`

GetSupportedImageTypes returns the SupportedImageTypes field if non-nil, zero value otherwise.

### GetSupportedImageTypesOk

`func (o *LibraryTypeOptions) GetSupportedImageTypesOk() (*[]ImageType, bool)`

GetSupportedImageTypesOk returns a tuple with the SupportedImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedImageTypes

`func (o *LibraryTypeOptions) SetSupportedImageTypes(v []ImageType)`

SetSupportedImageTypes sets SupportedImageTypes field to given value.

### HasSupportedImageTypes

`func (o *LibraryTypeOptions) HasSupportedImageTypes() bool`

HasSupportedImageTypes returns a boolean if a field has been set.

### GetDefaultImageOptions

`func (o *LibraryTypeOptions) GetDefaultImageOptions() []ImageOption`

GetDefaultImageOptions returns the DefaultImageOptions field if non-nil, zero value otherwise.

### GetDefaultImageOptionsOk

`func (o *LibraryTypeOptions) GetDefaultImageOptionsOk() (*[]ImageOption, bool)`

GetDefaultImageOptionsOk returns a tuple with the DefaultImageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageOptions

`func (o *LibraryTypeOptions) SetDefaultImageOptions(v []ImageOption)`

SetDefaultImageOptions sets DefaultImageOptions field to given value.

### HasDefaultImageOptions

`func (o *LibraryTypeOptions) HasDefaultImageOptions() bool`

HasDefaultImageOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


