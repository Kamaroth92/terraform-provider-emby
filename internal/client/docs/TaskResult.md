# TaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimeUtc** | Pointer to **time.Time** |  | [optional] 
**EndTimeUtc** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**TaskCompletionStatus**](TaskCompletionStatus.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**LongErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskResult

`func NewTaskResult() *TaskResult`

NewTaskResult instantiates a new TaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResultWithDefaults

`func NewTaskResultWithDefaults() *TaskResult`

NewTaskResultWithDefaults instantiates a new TaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimeUtc

`func (o *TaskResult) GetStartTimeUtc() time.Time`

GetStartTimeUtc returns the StartTimeUtc field if non-nil, zero value otherwise.

### GetStartTimeUtcOk

`func (o *TaskResult) GetStartTimeUtcOk() (*time.Time, bool)`

GetStartTimeUtcOk returns a tuple with the StartTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUtc

`func (o *TaskResult) SetStartTimeUtc(v time.Time)`

SetStartTimeUtc sets StartTimeUtc field to given value.

### HasStartTimeUtc

`func (o *TaskResult) HasStartTimeUtc() bool`

HasStartTimeUtc returns a boolean if a field has been set.

### GetEndTimeUtc

`func (o *TaskResult) GetEndTimeUtc() time.Time`

GetEndTimeUtc returns the EndTimeUtc field if non-nil, zero value otherwise.

### GetEndTimeUtcOk

`func (o *TaskResult) GetEndTimeUtcOk() (*time.Time, bool)`

GetEndTimeUtcOk returns a tuple with the EndTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUtc

`func (o *TaskResult) SetEndTimeUtc(v time.Time)`

SetEndTimeUtc sets EndTimeUtc field to given value.

### HasEndTimeUtc

`func (o *TaskResult) HasEndTimeUtc() bool`

HasEndTimeUtc returns a boolean if a field has been set.

### GetStatus

`func (o *TaskResult) GetStatus() TaskCompletionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskResult) GetStatusOk() (*TaskCompletionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskResult) SetStatus(v TaskCompletionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *TaskResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *TaskResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TaskResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TaskResult) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TaskResult) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetId

`func (o *TaskResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TaskResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TaskResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TaskResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TaskResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetLongErrorMessage

`func (o *TaskResult) GetLongErrorMessage() string`

GetLongErrorMessage returns the LongErrorMessage field if non-nil, zero value otherwise.

### GetLongErrorMessageOk

`func (o *TaskResult) GetLongErrorMessageOk() (*string, bool)`

GetLongErrorMessageOk returns a tuple with the LongErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongErrorMessage

`func (o *TaskResult) SetLongErrorMessage(v string)`

SetLongErrorMessage sets LongErrorMessage field to given value.

### HasLongErrorMessage

`func (o *TaskResult) HasLongErrorMessage() bool`

HasLongErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


