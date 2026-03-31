# UserNotificationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierKey** | Pointer to **string** |  | [optional] 
**SetupModuleUrl** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**PluginId** | Pointer to **string** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 
**DeviceIds** | Pointer to **[]string** |  | [optional] 
**LibraryIds** | Pointer to **[]string** |  | [optional] 
**EventIds** | Pointer to **[]string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**IsSelfNotification** | Pointer to **bool** |  | [optional] 
**GroupItems** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUserNotificationInfo

`func NewUserNotificationInfo() *UserNotificationInfo`

NewUserNotificationInfo instantiates a new UserNotificationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNotificationInfoWithDefaults

`func NewUserNotificationInfoWithDefaults() *UserNotificationInfo`

NewUserNotificationInfoWithDefaults instantiates a new UserNotificationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierKey

`func (o *UserNotificationInfo) GetNotifierKey() string`

GetNotifierKey returns the NotifierKey field if non-nil, zero value otherwise.

### GetNotifierKeyOk

`func (o *UserNotificationInfo) GetNotifierKeyOk() (*string, bool)`

GetNotifierKeyOk returns a tuple with the NotifierKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierKey

`func (o *UserNotificationInfo) SetNotifierKey(v string)`

SetNotifierKey sets NotifierKey field to given value.

### HasNotifierKey

`func (o *UserNotificationInfo) HasNotifierKey() bool`

HasNotifierKey returns a boolean if a field has been set.

### GetSetupModuleUrl

`func (o *UserNotificationInfo) GetSetupModuleUrl() string`

GetSetupModuleUrl returns the SetupModuleUrl field if non-nil, zero value otherwise.

### GetSetupModuleUrlOk

`func (o *UserNotificationInfo) GetSetupModuleUrlOk() (*string, bool)`

GetSetupModuleUrlOk returns a tuple with the SetupModuleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupModuleUrl

`func (o *UserNotificationInfo) SetSetupModuleUrl(v string)`

SetSetupModuleUrl sets SetupModuleUrl field to given value.

### HasSetupModuleUrl

`func (o *UserNotificationInfo) HasSetupModuleUrl() bool`

HasSetupModuleUrl returns a boolean if a field has been set.

### GetServiceName

`func (o *UserNotificationInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *UserNotificationInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *UserNotificationInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *UserNotificationInfo) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetPluginId

`func (o *UserNotificationInfo) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *UserNotificationInfo) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *UserNotificationInfo) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *UserNotificationInfo) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetFriendlyName

`func (o *UserNotificationInfo) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *UserNotificationInfo) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *UserNotificationInfo) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *UserNotificationInfo) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetId

`func (o *UserNotificationInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserNotificationInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserNotificationInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserNotificationInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *UserNotificationInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserNotificationInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserNotificationInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserNotificationInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserIds

`func (o *UserNotificationInfo) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UserNotificationInfo) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UserNotificationInfo) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UserNotificationInfo) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetDeviceIds

`func (o *UserNotificationInfo) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *UserNotificationInfo) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *UserNotificationInfo) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *UserNotificationInfo) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetLibraryIds

`func (o *UserNotificationInfo) GetLibraryIds() []string`

GetLibraryIds returns the LibraryIds field if non-nil, zero value otherwise.

### GetLibraryIdsOk

`func (o *UserNotificationInfo) GetLibraryIdsOk() (*[]string, bool)`

GetLibraryIdsOk returns a tuple with the LibraryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryIds

`func (o *UserNotificationInfo) SetLibraryIds(v []string)`

SetLibraryIds sets LibraryIds field to given value.

### HasLibraryIds

`func (o *UserNotificationInfo) HasLibraryIds() bool`

HasLibraryIds returns a boolean if a field has been set.

### GetEventIds

`func (o *UserNotificationInfo) GetEventIds() []string`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *UserNotificationInfo) GetEventIdsOk() (*[]string, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *UserNotificationInfo) SetEventIds(v []string)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *UserNotificationInfo) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.

### GetUserId

`func (o *UserNotificationInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserNotificationInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserNotificationInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserNotificationInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetIsSelfNotification

`func (o *UserNotificationInfo) GetIsSelfNotification() bool`

GetIsSelfNotification returns the IsSelfNotification field if non-nil, zero value otherwise.

### GetIsSelfNotificationOk

`func (o *UserNotificationInfo) GetIsSelfNotificationOk() (*bool, bool)`

GetIsSelfNotificationOk returns a tuple with the IsSelfNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfNotification

`func (o *UserNotificationInfo) SetIsSelfNotification(v bool)`

SetIsSelfNotification sets IsSelfNotification field to given value.

### HasIsSelfNotification

`func (o *UserNotificationInfo) HasIsSelfNotification() bool`

HasIsSelfNotification returns a boolean if a field has been set.

### GetGroupItems

`func (o *UserNotificationInfo) GetGroupItems() bool`

GetGroupItems returns the GroupItems field if non-nil, zero value otherwise.

### GetGroupItemsOk

`func (o *UserNotificationInfo) GetGroupItemsOk() (*bool, bool)`

GetGroupItemsOk returns a tuple with the GroupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupItems

`func (o *UserNotificationInfo) SetGroupItems(v bool)`

SetGroupItems sets GroupItems field to given value.

### HasGroupItems

`func (o *UserNotificationInfo) HasGroupItems() bool`

HasGroupItems returns a boolean if a field has been set.

### GetOptions

`func (o *UserNotificationInfo) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UserNotificationInfo) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UserNotificationInfo) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UserNotificationInfo) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


