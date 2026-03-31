# LibraryOptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**SetupUrl** | Pointer to **string** |  | [optional] 
**DefaultEnabled** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to [**[]MetadataFeatures**](MetadataFeatures.md) |  | [optional] 

## Methods

### NewLibraryOptionInfo

`func NewLibraryOptionInfo() *LibraryOptionInfo`

NewLibraryOptionInfo instantiates a new LibraryOptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryOptionInfoWithDefaults

`func NewLibraryOptionInfoWithDefaults() *LibraryOptionInfo`

NewLibraryOptionInfoWithDefaults instantiates a new LibraryOptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LibraryOptionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LibraryOptionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LibraryOptionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LibraryOptionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSetupUrl

`func (o *LibraryOptionInfo) GetSetupUrl() string`

GetSetupUrl returns the SetupUrl field if non-nil, zero value otherwise.

### GetSetupUrlOk

`func (o *LibraryOptionInfo) GetSetupUrlOk() (*string, bool)`

GetSetupUrlOk returns a tuple with the SetupUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupUrl

`func (o *LibraryOptionInfo) SetSetupUrl(v string)`

SetSetupUrl sets SetupUrl field to given value.

### HasSetupUrl

`func (o *LibraryOptionInfo) HasSetupUrl() bool`

HasSetupUrl returns a boolean if a field has been set.

### GetDefaultEnabled

`func (o *LibraryOptionInfo) GetDefaultEnabled() bool`

GetDefaultEnabled returns the DefaultEnabled field if non-nil, zero value otherwise.

### GetDefaultEnabledOk

`func (o *LibraryOptionInfo) GetDefaultEnabledOk() (*bool, bool)`

GetDefaultEnabledOk returns a tuple with the DefaultEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEnabled

`func (o *LibraryOptionInfo) SetDefaultEnabled(v bool)`

SetDefaultEnabled sets DefaultEnabled field to given value.

### HasDefaultEnabled

`func (o *LibraryOptionInfo) HasDefaultEnabled() bool`

HasDefaultEnabled returns a boolean if a field has been set.

### GetFeatures

`func (o *LibraryOptionInfo) GetFeatures() []MetadataFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *LibraryOptionInfo) GetFeaturesOk() (*[]MetadataFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *LibraryOptionInfo) SetFeatures(v []MetadataFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *LibraryOptionInfo) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


