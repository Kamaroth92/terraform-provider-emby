# QueryResultUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]UserDto**](UserDto.md) |  | [optional] 
**TotalRecordCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewQueryResultUserDto

`func NewQueryResultUserDto() *QueryResultUserDto`

NewQueryResultUserDto instantiates a new QueryResultUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResultUserDtoWithDefaults

`func NewQueryResultUserDtoWithDefaults() *QueryResultUserDto`

NewQueryResultUserDtoWithDefaults instantiates a new QueryResultUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *QueryResultUserDto) GetItems() []UserDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *QueryResultUserDto) GetItemsOk() (*[]UserDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *QueryResultUserDto) SetItems(v []UserDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *QueryResultUserDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *QueryResultUserDto) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *QueryResultUserDto) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *QueryResultUserDto) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *QueryResultUserDto) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


