# UserAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**UserActionType**](UserActionType.md) |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**PositionTicks** | Pointer to **NullableInt64** |  | [optional] 
**Played** | Pointer to **NullableBool** |  | [optional] 
**IsFavorite** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUserAction

`func NewUserAction() *UserAction`

NewUserAction instantiates a new UserAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionWithDefaults

`func NewUserActionWithDefaults() *UserAction`

NewUserActionWithDefaults instantiates a new UserAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServerId

`func (o *UserAction) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *UserAction) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *UserAction) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *UserAction) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetUserId

`func (o *UserAction) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserAction) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserAction) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserAction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetItemId

`func (o *UserAction) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *UserAction) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *UserAction) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *UserAction) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetType

`func (o *UserAction) GetType() UserActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserAction) GetTypeOk() (*UserActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserAction) SetType(v UserActionType)`

SetType sets Type field to given value.

### HasType

`func (o *UserAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDate

`func (o *UserAction) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UserAction) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UserAction) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *UserAction) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPositionTicks

`func (o *UserAction) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *UserAction) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *UserAction) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *UserAction) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *UserAction) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *UserAction) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetPlayed

`func (o *UserAction) GetPlayed() bool`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *UserAction) GetPlayedOk() (*bool, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *UserAction) SetPlayed(v bool)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *UserAction) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### SetPlayedNil

`func (o *UserAction) SetPlayedNil(b bool)`

 SetPlayedNil sets the value for Played to be an explicit nil

### UnsetPlayed
`func (o *UserAction) UnsetPlayed()`

UnsetPlayed ensures that no value is present for Played, not even an explicit nil
### GetIsFavorite

`func (o *UserAction) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *UserAction) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *UserAction) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *UserAction) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *UserAction) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *UserAction) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


