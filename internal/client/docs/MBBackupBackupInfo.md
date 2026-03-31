# MBBackupBackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerVersion** | Pointer to **string** |  | [optional] 
**PluginVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CanRestore** | Pointer to **bool** |  | [optional] 
**IsFullBackup** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**Users** | Pointer to [**[]NameIdPair**](NameIdPair.md) |  | [optional] 

## Methods

### NewMBBackupBackupInfo

`func NewMBBackupBackupInfo() *MBBackupBackupInfo`

NewMBBackupBackupInfo instantiates a new MBBackupBackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBBackupBackupInfoWithDefaults

`func NewMBBackupBackupInfoWithDefaults() *MBBackupBackupInfo`

NewMBBackupBackupInfoWithDefaults instantiates a new MBBackupBackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerVersion

`func (o *MBBackupBackupInfo) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *MBBackupBackupInfo) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *MBBackupBackupInfo) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *MBBackupBackupInfo) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetPluginVersion

`func (o *MBBackupBackupInfo) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *MBBackupBackupInfo) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *MBBackupBackupInfo) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.

### HasPluginVersion

`func (o *MBBackupBackupInfo) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.

### GetName

`func (o *MBBackupBackupInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MBBackupBackupInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MBBackupBackupInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MBBackupBackupInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCanRestore

`func (o *MBBackupBackupInfo) GetCanRestore() bool`

GetCanRestore returns the CanRestore field if non-nil, zero value otherwise.

### GetCanRestoreOk

`func (o *MBBackupBackupInfo) GetCanRestoreOk() (*bool, bool)`

GetCanRestoreOk returns a tuple with the CanRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRestore

`func (o *MBBackupBackupInfo) SetCanRestore(v bool)`

SetCanRestore sets CanRestore field to given value.

### HasCanRestore

`func (o *MBBackupBackupInfo) HasCanRestore() bool`

HasCanRestore returns a boolean if a field has been set.

### GetIsFullBackup

`func (o *MBBackupBackupInfo) GetIsFullBackup() bool`

GetIsFullBackup returns the IsFullBackup field if non-nil, zero value otherwise.

### GetIsFullBackupOk

`func (o *MBBackupBackupInfo) GetIsFullBackupOk() (*bool, bool)`

GetIsFullBackupOk returns a tuple with the IsFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullBackup

`func (o *MBBackupBackupInfo) SetIsFullBackup(v bool)`

SetIsFullBackup sets IsFullBackup field to given value.

### HasIsFullBackup

`func (o *MBBackupBackupInfo) HasIsFullBackup() bool`

HasIsFullBackup returns a boolean if a field has been set.

### GetDateCreated

`func (o *MBBackupBackupInfo) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *MBBackupBackupInfo) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *MBBackupBackupInfo) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *MBBackupBackupInfo) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetUsers

`func (o *MBBackupBackupInfo) GetUsers() []NameIdPair`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *MBBackupBackupInfo) GetUsersOk() (*[]NameIdPair, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *MBBackupBackupInfo) SetUsers(v []NameIdPair)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *MBBackupBackupInfo) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


