# DevicesDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **int64** |  | [optional] 
**ReportedDeviceId** | Pointer to **string** |  | [optional] 
**LastUserName** | Pointer to **string** |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**AppVersion** | Pointer to **string** |  | [optional] 
**LastUserId** | Pointer to **string** |  | [optional] 
**DateLastActivity** | Pointer to **time.Time** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewDevicesDeviceInfo

`func NewDevicesDeviceInfo() *DevicesDeviceInfo`

NewDevicesDeviceInfo instantiates a new DevicesDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesDeviceInfoWithDefaults

`func NewDevicesDeviceInfoWithDefaults() *DevicesDeviceInfo`

NewDevicesDeviceInfoWithDefaults instantiates a new DevicesDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DevicesDeviceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevicesDeviceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevicesDeviceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevicesDeviceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *DevicesDeviceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevicesDeviceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevicesDeviceInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevicesDeviceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalId

`func (o *DevicesDeviceInfo) GetInternalId() int64`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *DevicesDeviceInfo) GetInternalIdOk() (*int64, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *DevicesDeviceInfo) SetInternalId(v int64)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *DevicesDeviceInfo) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetReportedDeviceId

`func (o *DevicesDeviceInfo) GetReportedDeviceId() string`

GetReportedDeviceId returns the ReportedDeviceId field if non-nil, zero value otherwise.

### GetReportedDeviceIdOk

`func (o *DevicesDeviceInfo) GetReportedDeviceIdOk() (*string, bool)`

GetReportedDeviceIdOk returns a tuple with the ReportedDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedDeviceId

`func (o *DevicesDeviceInfo) SetReportedDeviceId(v string)`

SetReportedDeviceId sets ReportedDeviceId field to given value.

### HasReportedDeviceId

`func (o *DevicesDeviceInfo) HasReportedDeviceId() bool`

HasReportedDeviceId returns a boolean if a field has been set.

### GetLastUserName

`func (o *DevicesDeviceInfo) GetLastUserName() string`

GetLastUserName returns the LastUserName field if non-nil, zero value otherwise.

### GetLastUserNameOk

`func (o *DevicesDeviceInfo) GetLastUserNameOk() (*string, bool)`

GetLastUserNameOk returns a tuple with the LastUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserName

`func (o *DevicesDeviceInfo) SetLastUserName(v string)`

SetLastUserName sets LastUserName field to given value.

### HasLastUserName

`func (o *DevicesDeviceInfo) HasLastUserName() bool`

HasLastUserName returns a boolean if a field has been set.

### GetAppName

`func (o *DevicesDeviceInfo) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *DevicesDeviceInfo) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *DevicesDeviceInfo) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *DevicesDeviceInfo) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppVersion

`func (o *DevicesDeviceInfo) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *DevicesDeviceInfo) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *DevicesDeviceInfo) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *DevicesDeviceInfo) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetLastUserId

`func (o *DevicesDeviceInfo) GetLastUserId() string`

GetLastUserId returns the LastUserId field if non-nil, zero value otherwise.

### GetLastUserIdOk

`func (o *DevicesDeviceInfo) GetLastUserIdOk() (*string, bool)`

GetLastUserIdOk returns a tuple with the LastUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserId

`func (o *DevicesDeviceInfo) SetLastUserId(v string)`

SetLastUserId sets LastUserId field to given value.

### HasLastUserId

`func (o *DevicesDeviceInfo) HasLastUserId() bool`

HasLastUserId returns a boolean if a field has been set.

### GetDateLastActivity

`func (o *DevicesDeviceInfo) GetDateLastActivity() time.Time`

GetDateLastActivity returns the DateLastActivity field if non-nil, zero value otherwise.

### GetDateLastActivityOk

`func (o *DevicesDeviceInfo) GetDateLastActivityOk() (*time.Time, bool)`

GetDateLastActivityOk returns a tuple with the DateLastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastActivity

`func (o *DevicesDeviceInfo) SetDateLastActivity(v time.Time)`

SetDateLastActivity sets DateLastActivity field to given value.

### HasDateLastActivity

`func (o *DevicesDeviceInfo) HasDateLastActivity() bool`

HasDateLastActivity returns a boolean if a field has been set.

### GetIconUrl

`func (o *DevicesDeviceInfo) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *DevicesDeviceInfo) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *DevicesDeviceInfo) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *DevicesDeviceInfo) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetIpAddress

`func (o *DevicesDeviceInfo) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DevicesDeviceInfo) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DevicesDeviceInfo) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *DevicesDeviceInfo) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


