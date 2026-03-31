# LibraryOptionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataSavers** | Pointer to [**[]LibraryOptionInfo**](LibraryOptionInfo.md) |  | [optional] 
**MetadataReaders** | Pointer to [**[]LibraryOptionInfo**](LibraryOptionInfo.md) |  | [optional] 
**SubtitleFetchers** | Pointer to [**[]LibraryOptionInfo**](LibraryOptionInfo.md) |  | [optional] 
**LyricsFetchers** | Pointer to [**[]LibraryOptionInfo**](LibraryOptionInfo.md) |  | [optional] 
**TypeOptions** | Pointer to [**[]LibraryTypeOptions**](LibraryTypeOptions.md) |  | [optional] 
**DefaultLibraryOptions** | Pointer to [**LibraryOptions**](LibraryOptions.md) |  | [optional] 

## Methods

### NewLibraryOptionsResult

`func NewLibraryOptionsResult() *LibraryOptionsResult`

NewLibraryOptionsResult instantiates a new LibraryOptionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryOptionsResultWithDefaults

`func NewLibraryOptionsResultWithDefaults() *LibraryOptionsResult`

NewLibraryOptionsResultWithDefaults instantiates a new LibraryOptionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataSavers

`func (o *LibraryOptionsResult) GetMetadataSavers() []LibraryOptionInfo`

GetMetadataSavers returns the MetadataSavers field if non-nil, zero value otherwise.

### GetMetadataSaversOk

`func (o *LibraryOptionsResult) GetMetadataSaversOk() (*[]LibraryOptionInfo, bool)`

GetMetadataSaversOk returns a tuple with the MetadataSavers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSavers

`func (o *LibraryOptionsResult) SetMetadataSavers(v []LibraryOptionInfo)`

SetMetadataSavers sets MetadataSavers field to given value.

### HasMetadataSavers

`func (o *LibraryOptionsResult) HasMetadataSavers() bool`

HasMetadataSavers returns a boolean if a field has been set.

### GetMetadataReaders

`func (o *LibraryOptionsResult) GetMetadataReaders() []LibraryOptionInfo`

GetMetadataReaders returns the MetadataReaders field if non-nil, zero value otherwise.

### GetMetadataReadersOk

`func (o *LibraryOptionsResult) GetMetadataReadersOk() (*[]LibraryOptionInfo, bool)`

GetMetadataReadersOk returns a tuple with the MetadataReaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataReaders

`func (o *LibraryOptionsResult) SetMetadataReaders(v []LibraryOptionInfo)`

SetMetadataReaders sets MetadataReaders field to given value.

### HasMetadataReaders

`func (o *LibraryOptionsResult) HasMetadataReaders() bool`

HasMetadataReaders returns a boolean if a field has been set.

### GetSubtitleFetchers

`func (o *LibraryOptionsResult) GetSubtitleFetchers() []LibraryOptionInfo`

GetSubtitleFetchers returns the SubtitleFetchers field if non-nil, zero value otherwise.

### GetSubtitleFetchersOk

`func (o *LibraryOptionsResult) GetSubtitleFetchersOk() (*[]LibraryOptionInfo, bool)`

GetSubtitleFetchersOk returns a tuple with the SubtitleFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleFetchers

`func (o *LibraryOptionsResult) SetSubtitleFetchers(v []LibraryOptionInfo)`

SetSubtitleFetchers sets SubtitleFetchers field to given value.

### HasSubtitleFetchers

`func (o *LibraryOptionsResult) HasSubtitleFetchers() bool`

HasSubtitleFetchers returns a boolean if a field has been set.

### GetLyricsFetchers

`func (o *LibraryOptionsResult) GetLyricsFetchers() []LibraryOptionInfo`

GetLyricsFetchers returns the LyricsFetchers field if non-nil, zero value otherwise.

### GetLyricsFetchersOk

`func (o *LibraryOptionsResult) GetLyricsFetchersOk() (*[]LibraryOptionInfo, bool)`

GetLyricsFetchersOk returns a tuple with the LyricsFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyricsFetchers

`func (o *LibraryOptionsResult) SetLyricsFetchers(v []LibraryOptionInfo)`

SetLyricsFetchers sets LyricsFetchers field to given value.

### HasLyricsFetchers

`func (o *LibraryOptionsResult) HasLyricsFetchers() bool`

HasLyricsFetchers returns a boolean if a field has been set.

### GetTypeOptions

`func (o *LibraryOptionsResult) GetTypeOptions() []LibraryTypeOptions`

GetTypeOptions returns the TypeOptions field if non-nil, zero value otherwise.

### GetTypeOptionsOk

`func (o *LibraryOptionsResult) GetTypeOptionsOk() (*[]LibraryTypeOptions, bool)`

GetTypeOptionsOk returns a tuple with the TypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOptions

`func (o *LibraryOptionsResult) SetTypeOptions(v []LibraryTypeOptions)`

SetTypeOptions sets TypeOptions field to given value.

### HasTypeOptions

`func (o *LibraryOptionsResult) HasTypeOptions() bool`

HasTypeOptions returns a boolean if a field has been set.

### GetDefaultLibraryOptions

`func (o *LibraryOptionsResult) GetDefaultLibraryOptions() LibraryOptions`

GetDefaultLibraryOptions returns the DefaultLibraryOptions field if non-nil, zero value otherwise.

### GetDefaultLibraryOptionsOk

`func (o *LibraryOptionsResult) GetDefaultLibraryOptionsOk() (*LibraryOptions, bool)`

GetDefaultLibraryOptionsOk returns a tuple with the DefaultLibraryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLibraryOptions

`func (o *LibraryOptionsResult) SetDefaultLibraryOptions(v LibraryOptions)`

SetDefaultLibraryOptions sets DefaultLibraryOptions field to given value.

### HasDefaultLibraryOptions

`func (o *LibraryOptionsResult) HasDefaultLibraryOptions() bool`

HasDefaultLibraryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


