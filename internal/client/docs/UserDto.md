# UserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**ConnectUserName** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**ConnectLinkType** | Pointer to [**ConnectUserLinkType**](ConnectUserLinkType.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PrimaryImageTag** | Pointer to **string** |  | [optional] 
**HasPassword** | Pointer to **NullableBool** |  | [optional] 
**HasConfiguredPassword** | Pointer to **NullableBool** |  | [optional] 
**EnableAutoLogin** | Pointer to **NullableBool** |  | [optional] 
**LastLoginDate** | Pointer to **NullableTime** |  | [optional] 
**LastActivityDate** | Pointer to **NullableTime** |  | [optional] 
**Configuration** | Pointer to [**UserConfiguration**](UserConfiguration.md) |  | [optional] 
**Policy** | Pointer to [**UserPolicy**](UserPolicy.md) |  | [optional] 
**PrimaryImageAspectRatio** | Pointer to **NullableFloat64** |  | [optional] 
**UserItemShareLevel** | Pointer to [**UserItemShareLevel**](UserItemShareLevel.md) |  | [optional] 

## Methods

### NewUserDto

`func NewUserDto() *UserDto`

NewUserDto instantiates a new UserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDtoWithDefaults

`func NewUserDtoWithDefaults() *UserDto`

NewUserDtoWithDefaults instantiates a new UserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerId

`func (o *UserDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *UserDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *UserDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *UserDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *UserDto) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *UserDto) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *UserDto) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *UserDto) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetPrefix

`func (o *UserDto) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *UserDto) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *UserDto) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *UserDto) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetConnectUserName

`func (o *UserDto) GetConnectUserName() string`

GetConnectUserName returns the ConnectUserName field if non-nil, zero value otherwise.

### GetConnectUserNameOk

`func (o *UserDto) GetConnectUserNameOk() (*string, bool)`

GetConnectUserNameOk returns a tuple with the ConnectUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectUserName

`func (o *UserDto) SetConnectUserName(v string)`

SetConnectUserName sets ConnectUserName field to given value.

### HasConnectUserName

`func (o *UserDto) HasConnectUserName() bool`

HasConnectUserName returns a boolean if a field has been set.

### GetDateCreated

`func (o *UserDto) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UserDto) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UserDto) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UserDto) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *UserDto) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *UserDto) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetConnectLinkType

`func (o *UserDto) GetConnectLinkType() ConnectUserLinkType`

GetConnectLinkType returns the ConnectLinkType field if non-nil, zero value otherwise.

### GetConnectLinkTypeOk

`func (o *UserDto) GetConnectLinkTypeOk() (*ConnectUserLinkType, bool)`

GetConnectLinkTypeOk returns a tuple with the ConnectLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectLinkType

`func (o *UserDto) SetConnectLinkType(v ConnectUserLinkType)`

SetConnectLinkType sets ConnectLinkType field to given value.

### HasConnectLinkType

`func (o *UserDto) HasConnectLinkType() bool`

HasConnectLinkType returns a boolean if a field has been set.

### GetId

`func (o *UserDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *UserDto) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *UserDto) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *UserDto) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *UserDto) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### GetHasPassword

`func (o *UserDto) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *UserDto) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *UserDto) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *UserDto) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### SetHasPasswordNil

`func (o *UserDto) SetHasPasswordNil(b bool)`

 SetHasPasswordNil sets the value for HasPassword to be an explicit nil

### UnsetHasPassword
`func (o *UserDto) UnsetHasPassword()`

UnsetHasPassword ensures that no value is present for HasPassword, not even an explicit nil
### GetHasConfiguredPassword

`func (o *UserDto) GetHasConfiguredPassword() bool`

GetHasConfiguredPassword returns the HasConfiguredPassword field if non-nil, zero value otherwise.

### GetHasConfiguredPasswordOk

`func (o *UserDto) GetHasConfiguredPasswordOk() (*bool, bool)`

GetHasConfiguredPasswordOk returns a tuple with the HasConfiguredPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfiguredPassword

`func (o *UserDto) SetHasConfiguredPassword(v bool)`

SetHasConfiguredPassword sets HasConfiguredPassword field to given value.

### HasHasConfiguredPassword

`func (o *UserDto) HasHasConfiguredPassword() bool`

HasHasConfiguredPassword returns a boolean if a field has been set.

### SetHasConfiguredPasswordNil

`func (o *UserDto) SetHasConfiguredPasswordNil(b bool)`

 SetHasConfiguredPasswordNil sets the value for HasConfiguredPassword to be an explicit nil

### UnsetHasConfiguredPassword
`func (o *UserDto) UnsetHasConfiguredPassword()`

UnsetHasConfiguredPassword ensures that no value is present for HasConfiguredPassword, not even an explicit nil
### GetEnableAutoLogin

`func (o *UserDto) GetEnableAutoLogin() bool`

GetEnableAutoLogin returns the EnableAutoLogin field if non-nil, zero value otherwise.

### GetEnableAutoLoginOk

`func (o *UserDto) GetEnableAutoLoginOk() (*bool, bool)`

GetEnableAutoLoginOk returns a tuple with the EnableAutoLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoLogin

`func (o *UserDto) SetEnableAutoLogin(v bool)`

SetEnableAutoLogin sets EnableAutoLogin field to given value.

### HasEnableAutoLogin

`func (o *UserDto) HasEnableAutoLogin() bool`

HasEnableAutoLogin returns a boolean if a field has been set.

### SetEnableAutoLoginNil

`func (o *UserDto) SetEnableAutoLoginNil(b bool)`

 SetEnableAutoLoginNil sets the value for EnableAutoLogin to be an explicit nil

### UnsetEnableAutoLogin
`func (o *UserDto) UnsetEnableAutoLogin()`

UnsetEnableAutoLogin ensures that no value is present for EnableAutoLogin, not even an explicit nil
### GetLastLoginDate

`func (o *UserDto) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *UserDto) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *UserDto) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *UserDto) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### SetLastLoginDateNil

`func (o *UserDto) SetLastLoginDateNil(b bool)`

 SetLastLoginDateNil sets the value for LastLoginDate to be an explicit nil

### UnsetLastLoginDate
`func (o *UserDto) UnsetLastLoginDate()`

UnsetLastLoginDate ensures that no value is present for LastLoginDate, not even an explicit nil
### GetLastActivityDate

`func (o *UserDto) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *UserDto) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *UserDto) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *UserDto) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### SetLastActivityDateNil

`func (o *UserDto) SetLastActivityDateNil(b bool)`

 SetLastActivityDateNil sets the value for LastActivityDate to be an explicit nil

### UnsetLastActivityDate
`func (o *UserDto) UnsetLastActivityDate()`

UnsetLastActivityDate ensures that no value is present for LastActivityDate, not even an explicit nil
### GetConfiguration

`func (o *UserDto) GetConfiguration() UserConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UserDto) GetConfigurationOk() (*UserConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UserDto) SetConfiguration(v UserConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UserDto) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPolicy

`func (o *UserDto) GetPolicy() UserPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UserDto) GetPolicyOk() (*UserPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UserDto) SetPolicy(v UserPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UserDto) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPrimaryImageAspectRatio

`func (o *UserDto) GetPrimaryImageAspectRatio() float64`

GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field if non-nil, zero value otherwise.

### GetPrimaryImageAspectRatioOk

`func (o *UserDto) GetPrimaryImageAspectRatioOk() (*float64, bool)`

GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageAspectRatio

`func (o *UserDto) SetPrimaryImageAspectRatio(v float64)`

SetPrimaryImageAspectRatio sets PrimaryImageAspectRatio field to given value.

### HasPrimaryImageAspectRatio

`func (o *UserDto) HasPrimaryImageAspectRatio() bool`

HasPrimaryImageAspectRatio returns a boolean if a field has been set.

### SetPrimaryImageAspectRatioNil

`func (o *UserDto) SetPrimaryImageAspectRatioNil(b bool)`

 SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil

### UnsetPrimaryImageAspectRatio
`func (o *UserDto) UnsetPrimaryImageAspectRatio()`

UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil
### GetUserItemShareLevel

`func (o *UserDto) GetUserItemShareLevel() UserItemShareLevel`

GetUserItemShareLevel returns the UserItemShareLevel field if non-nil, zero value otherwise.

### GetUserItemShareLevelOk

`func (o *UserDto) GetUserItemShareLevelOk() (*UserItemShareLevel, bool)`

GetUserItemShareLevelOk returns a tuple with the UserItemShareLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserItemShareLevel

`func (o *UserDto) SetUserItemShareLevel(v UserItemShareLevel)`

SetUserItemShareLevel sets UserItemShareLevel field to given value.

### HasUserItemShareLevel

`func (o *UserDto) HasUserItemShareLevel() bool`

HasUserItemShareLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


