# ApiEpgRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**BaseItemDto**](BaseItemDto.md) |  | [optional] 
**Programs** | Pointer to [**[]BaseItemDto**](BaseItemDto.md) |  | [optional] 

## Methods

### NewApiEpgRow

`func NewApiEpgRow() *ApiEpgRow`

NewApiEpgRow instantiates a new ApiEpgRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEpgRowWithDefaults

`func NewApiEpgRowWithDefaults() *ApiEpgRow`

NewApiEpgRowWithDefaults instantiates a new ApiEpgRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ApiEpgRow) GetChannel() BaseItemDto`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApiEpgRow) GetChannelOk() (*BaseItemDto, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApiEpgRow) SetChannel(v BaseItemDto)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApiEpgRow) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetPrograms

`func (o *ApiEpgRow) GetPrograms() []BaseItemDto`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *ApiEpgRow) GetProgramsOk() (*[]BaseItemDto, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *ApiEpgRow) SetPrograms(v []BaseItemDto)`

SetPrograms sets Programs field to given value.

### HasPrograms

`func (o *ApiEpgRow) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


