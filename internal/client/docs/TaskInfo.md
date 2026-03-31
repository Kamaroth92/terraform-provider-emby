# TaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**TaskState**](TaskState.md) |  | [optional] 
**CurrentProgressPercentage** | Pointer to **NullableFloat64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastExecutionResult** | Pointer to [**TaskResult**](TaskResult.md) |  | [optional] 
**Triggers** | Pointer to [**[]TaskTriggerInfo**](TaskTriggerInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskInfo

`func NewTaskInfo() *TaskInfo`

NewTaskInfo instantiates a new TaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInfoWithDefaults

`func NewTaskInfoWithDefaults() *TaskInfo`

NewTaskInfoWithDefaults instantiates a new TaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaskInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *TaskInfo) GetState() TaskState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskInfo) GetStateOk() (*TaskState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskInfo) SetState(v TaskState)`

SetState sets State field to given value.

### HasState

`func (o *TaskInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCurrentProgressPercentage

`func (o *TaskInfo) GetCurrentProgressPercentage() float64`

GetCurrentProgressPercentage returns the CurrentProgressPercentage field if non-nil, zero value otherwise.

### GetCurrentProgressPercentageOk

`func (o *TaskInfo) GetCurrentProgressPercentageOk() (*float64, bool)`

GetCurrentProgressPercentageOk returns a tuple with the CurrentProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgressPercentage

`func (o *TaskInfo) SetCurrentProgressPercentage(v float64)`

SetCurrentProgressPercentage sets CurrentProgressPercentage field to given value.

### HasCurrentProgressPercentage

`func (o *TaskInfo) HasCurrentProgressPercentage() bool`

HasCurrentProgressPercentage returns a boolean if a field has been set.

### SetCurrentProgressPercentageNil

`func (o *TaskInfo) SetCurrentProgressPercentageNil(b bool)`

 SetCurrentProgressPercentageNil sets the value for CurrentProgressPercentage to be an explicit nil

### UnsetCurrentProgressPercentage
`func (o *TaskInfo) UnsetCurrentProgressPercentage()`

UnsetCurrentProgressPercentage ensures that no value is present for CurrentProgressPercentage, not even an explicit nil
### GetId

`func (o *TaskInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastExecutionResult

`func (o *TaskInfo) GetLastExecutionResult() TaskResult`

GetLastExecutionResult returns the LastExecutionResult field if non-nil, zero value otherwise.

### GetLastExecutionResultOk

`func (o *TaskInfo) GetLastExecutionResultOk() (*TaskResult, bool)`

GetLastExecutionResultOk returns a tuple with the LastExecutionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionResult

`func (o *TaskInfo) SetLastExecutionResult(v TaskResult)`

SetLastExecutionResult sets LastExecutionResult field to given value.

### HasLastExecutionResult

`func (o *TaskInfo) HasLastExecutionResult() bool`

HasLastExecutionResult returns a boolean if a field has been set.

### GetTriggers

`func (o *TaskInfo) GetTriggers() []TaskTriggerInfo`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *TaskInfo) GetTriggersOk() (*[]TaskTriggerInfo, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *TaskInfo) SetTriggers(v []TaskTriggerInfo)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *TaskInfo) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetDescription

`func (o *TaskInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *TaskInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TaskInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TaskInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TaskInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetIsHidden

`func (o *TaskInfo) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *TaskInfo) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *TaskInfo) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *TaskInfo) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetKey

`func (o *TaskInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TaskInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TaskInfo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TaskInfo) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


