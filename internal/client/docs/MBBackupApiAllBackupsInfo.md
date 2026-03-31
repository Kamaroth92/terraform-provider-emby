# MBBackupApiAllBackupsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullBackupInfo** | Pointer to [**MBBackupBackupInfo**](MBBackupBackupInfo.md) |  | [optional] 
**LightBackups** | Pointer to [**[]MBBackupBackupInfo**](MBBackupBackupInfo.md) |  | [optional] 

## Methods

### NewMBBackupApiAllBackupsInfo

`func NewMBBackupApiAllBackupsInfo() *MBBackupApiAllBackupsInfo`

NewMBBackupApiAllBackupsInfo instantiates a new MBBackupApiAllBackupsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBBackupApiAllBackupsInfoWithDefaults

`func NewMBBackupApiAllBackupsInfoWithDefaults() *MBBackupApiAllBackupsInfo`

NewMBBackupApiAllBackupsInfoWithDefaults instantiates a new MBBackupApiAllBackupsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullBackupInfo

`func (o *MBBackupApiAllBackupsInfo) GetFullBackupInfo() MBBackupBackupInfo`

GetFullBackupInfo returns the FullBackupInfo field if non-nil, zero value otherwise.

### GetFullBackupInfoOk

`func (o *MBBackupApiAllBackupsInfo) GetFullBackupInfoOk() (*MBBackupBackupInfo, bool)`

GetFullBackupInfoOk returns a tuple with the FullBackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupInfo

`func (o *MBBackupApiAllBackupsInfo) SetFullBackupInfo(v MBBackupBackupInfo)`

SetFullBackupInfo sets FullBackupInfo field to given value.

### HasFullBackupInfo

`func (o *MBBackupApiAllBackupsInfo) HasFullBackupInfo() bool`

HasFullBackupInfo returns a boolean if a field has been set.

### GetLightBackups

`func (o *MBBackupApiAllBackupsInfo) GetLightBackups() []MBBackupBackupInfo`

GetLightBackups returns the LightBackups field if non-nil, zero value otherwise.

### GetLightBackupsOk

`func (o *MBBackupApiAllBackupsInfo) GetLightBackupsOk() (*[]MBBackupBackupInfo, bool)`

GetLightBackupsOk returns a tuple with the LightBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightBackups

`func (o *MBBackupApiAllBackupsInfo) SetLightBackups(v []MBBackupBackupInfo)`

SetLightBackups sets LightBackups field to given value.

### HasLightBackups

`func (o *MBBackupApiAllBackupsInfo) HasLightBackups() bool`

HasLightBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


