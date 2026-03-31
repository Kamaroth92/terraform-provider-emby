# SessionPartyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Sessions** | Pointer to [**[]SessionSessionInfo**](SessionSessionInfo.md) |  | [optional] 
**Users** | Pointer to [**[]EntitiesUser**](EntitiesUser.md) |  | [optional] 

## Methods

### NewSessionPartyInfo

`func NewSessionPartyInfo() *SessionPartyInfo`

NewSessionPartyInfo instantiates a new SessionPartyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionPartyInfoWithDefaults

`func NewSessionPartyInfoWithDefaults() *SessionPartyInfo`

NewSessionPartyInfoWithDefaults instantiates a new SessionPartyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionPartyInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionPartyInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionPartyInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionPartyInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SessionPartyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionPartyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionPartyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SessionPartyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSessions

`func (o *SessionPartyInfo) GetSessions() []SessionSessionInfo`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *SessionPartyInfo) GetSessionsOk() (*[]SessionSessionInfo, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *SessionPartyInfo) SetSessions(v []SessionSessionInfo)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *SessionPartyInfo) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetUsers

`func (o *SessionPartyInfo) GetUsers() []EntitiesUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SessionPartyInfo) GetUsersOk() (*[]EntitiesUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SessionPartyInfo) SetUsers(v []EntitiesUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *SessionPartyInfo) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


