# MBBackupApiRestoreOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreServerId** | Pointer to **bool** |  | [optional] 
**UseFiles** | Pointer to **string** |  | [optional] 

## Methods

### NewMBBackupApiRestoreOptions

`func NewMBBackupApiRestoreOptions() *MBBackupApiRestoreOptions`

NewMBBackupApiRestoreOptions instantiates a new MBBackupApiRestoreOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBBackupApiRestoreOptionsWithDefaults

`func NewMBBackupApiRestoreOptionsWithDefaults() *MBBackupApiRestoreOptions`

NewMBBackupApiRestoreOptionsWithDefaults instantiates a new MBBackupApiRestoreOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreServerId

`func (o *MBBackupApiRestoreOptions) GetRestoreServerId() bool`

GetRestoreServerId returns the RestoreServerId field if non-nil, zero value otherwise.

### GetRestoreServerIdOk

`func (o *MBBackupApiRestoreOptions) GetRestoreServerIdOk() (*bool, bool)`

GetRestoreServerIdOk returns a tuple with the RestoreServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreServerId

`func (o *MBBackupApiRestoreOptions) SetRestoreServerId(v bool)`

SetRestoreServerId sets RestoreServerId field to given value.

### HasRestoreServerId

`func (o *MBBackupApiRestoreOptions) HasRestoreServerId() bool`

HasRestoreServerId returns a boolean if a field has been set.

### GetUseFiles

`func (o *MBBackupApiRestoreOptions) GetUseFiles() string`

GetUseFiles returns the UseFiles field if non-nil, zero value otherwise.

### GetUseFilesOk

`func (o *MBBackupApiRestoreOptions) GetUseFilesOk() (*string, bool)`

GetUseFilesOk returns a tuple with the UseFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFiles

`func (o *MBBackupApiRestoreOptions) SetUseFiles(v string)`

SetUseFiles sets UseFiles field to given value.

### HasUseFiles

`func (o *MBBackupApiRestoreOptions) HasUseFiles() bool`

HasUseFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


