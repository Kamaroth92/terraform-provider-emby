# CreateUserByName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**CopyFromUserId** | Pointer to **string** |  | [optional] 
**UserCopyOptions** | Pointer to [**[]LibraryUserCopyOptions**](LibraryUserCopyOptions.md) |  | [optional] 

## Methods

### NewCreateUserByName

`func NewCreateUserByName() *CreateUserByName`

NewCreateUserByName instantiates a new CreateUserByName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserByNameWithDefaults

`func NewCreateUserByNameWithDefaults() *CreateUserByName`

NewCreateUserByNameWithDefaults instantiates a new CreateUserByName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUserByName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserByName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserByName) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUserByName) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCopyFromUserId

`func (o *CreateUserByName) GetCopyFromUserId() string`

GetCopyFromUserId returns the CopyFromUserId field if non-nil, zero value otherwise.

### GetCopyFromUserIdOk

`func (o *CreateUserByName) GetCopyFromUserIdOk() (*string, bool)`

GetCopyFromUserIdOk returns a tuple with the CopyFromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyFromUserId

`func (o *CreateUserByName) SetCopyFromUserId(v string)`

SetCopyFromUserId sets CopyFromUserId field to given value.

### HasCopyFromUserId

`func (o *CreateUserByName) HasCopyFromUserId() bool`

HasCopyFromUserId returns a boolean if a field has been set.

### GetUserCopyOptions

`func (o *CreateUserByName) GetUserCopyOptions() []LibraryUserCopyOptions`

GetUserCopyOptions returns the UserCopyOptions field if non-nil, zero value otherwise.

### GetUserCopyOptionsOk

`func (o *CreateUserByName) GetUserCopyOptionsOk() (*[]LibraryUserCopyOptions, bool)`

GetUserCopyOptionsOk returns a tuple with the UserCopyOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCopyOptions

`func (o *CreateUserByName) SetUserCopyOptions(v []LibraryUserCopyOptions)`

SetUserCopyOptions sets UserCopyOptions field to given value.

### HasUserCopyOptions

`func (o *CreateUserByName) HasUserCopyOptions() bool`

HasUserCopyOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


