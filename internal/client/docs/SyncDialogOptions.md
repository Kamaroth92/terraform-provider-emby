# SyncDialogOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to [**[]SyncTarget**](SyncTarget.md) |  | [optional] 
**Options** | Pointer to [**[]SyncJobOption**](SyncJobOption.md) |  | [optional] 
**QualityOptions** | Pointer to [**[]SyncQualityOption**](SyncQualityOption.md) |  | [optional] 
**ProfileOptions** | Pointer to [**[]SyncProfileOption**](SyncProfileOption.md) |  | [optional] 

## Methods

### NewSyncDialogOptions

`func NewSyncDialogOptions() *SyncDialogOptions`

NewSyncDialogOptions instantiates a new SyncDialogOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncDialogOptionsWithDefaults

`func NewSyncDialogOptionsWithDefaults() *SyncDialogOptions`

NewSyncDialogOptionsWithDefaults instantiates a new SyncDialogOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *SyncDialogOptions) GetTargets() []SyncTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *SyncDialogOptions) GetTargetsOk() (*[]SyncTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *SyncDialogOptions) SetTargets(v []SyncTarget)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *SyncDialogOptions) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetOptions

`func (o *SyncDialogOptions) GetOptions() []SyncJobOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SyncDialogOptions) GetOptionsOk() (*[]SyncJobOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SyncDialogOptions) SetOptions(v []SyncJobOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SyncDialogOptions) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetQualityOptions

`func (o *SyncDialogOptions) GetQualityOptions() []SyncQualityOption`

GetQualityOptions returns the QualityOptions field if non-nil, zero value otherwise.

### GetQualityOptionsOk

`func (o *SyncDialogOptions) GetQualityOptionsOk() (*[]SyncQualityOption, bool)`

GetQualityOptionsOk returns a tuple with the QualityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityOptions

`func (o *SyncDialogOptions) SetQualityOptions(v []SyncQualityOption)`

SetQualityOptions sets QualityOptions field to given value.

### HasQualityOptions

`func (o *SyncDialogOptions) HasQualityOptions() bool`

HasQualityOptions returns a boolean if a field has been set.

### GetProfileOptions

`func (o *SyncDialogOptions) GetProfileOptions() []SyncProfileOption`

GetProfileOptions returns the ProfileOptions field if non-nil, zero value otherwise.

### GetProfileOptionsOk

`func (o *SyncDialogOptions) GetProfileOptionsOk() (*[]SyncProfileOption, bool)`

GetProfileOptionsOk returns a tuple with the ProfileOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOptions

`func (o *SyncDialogOptions) SetProfileOptions(v []SyncProfileOption)`

SetProfileOptions sets ProfileOptions field to given value.

### HasProfileOptions

`func (o *SyncDialogOptions) HasProfileOptions() bool`

HasProfileOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


