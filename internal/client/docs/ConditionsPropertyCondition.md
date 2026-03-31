# ConditionsPropertyCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedPropertyId** | Pointer to **string** |  | [optional] 
**ConditionType** | Pointer to [**ConditionsPropertyConditionType**](ConditionsPropertyConditionType.md) |  | [optional] 
**TargetPropertyId** | Pointer to **string** |  | [optional] 
**SimpleCondition** | Pointer to [**AttributesSimpleCondition**](AttributesSimpleCondition.md) |  | [optional] 
**ValueCondition** | Pointer to [**AttributesValueCondition**](AttributesValueCondition.md) |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConditionsPropertyCondition

`func NewConditionsPropertyCondition() *ConditionsPropertyCondition`

NewConditionsPropertyCondition instantiates a new ConditionsPropertyCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionsPropertyConditionWithDefaults

`func NewConditionsPropertyConditionWithDefaults() *ConditionsPropertyCondition`

NewConditionsPropertyConditionWithDefaults instantiates a new ConditionsPropertyCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedPropertyId

`func (o *ConditionsPropertyCondition) GetAffectedPropertyId() string`

GetAffectedPropertyId returns the AffectedPropertyId field if non-nil, zero value otherwise.

### GetAffectedPropertyIdOk

`func (o *ConditionsPropertyCondition) GetAffectedPropertyIdOk() (*string, bool)`

GetAffectedPropertyIdOk returns a tuple with the AffectedPropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPropertyId

`func (o *ConditionsPropertyCondition) SetAffectedPropertyId(v string)`

SetAffectedPropertyId sets AffectedPropertyId field to given value.

### HasAffectedPropertyId

`func (o *ConditionsPropertyCondition) HasAffectedPropertyId() bool`

HasAffectedPropertyId returns a boolean if a field has been set.

### GetConditionType

`func (o *ConditionsPropertyCondition) GetConditionType() ConditionsPropertyConditionType`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *ConditionsPropertyCondition) GetConditionTypeOk() (*ConditionsPropertyConditionType, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *ConditionsPropertyCondition) SetConditionType(v ConditionsPropertyConditionType)`

SetConditionType sets ConditionType field to given value.

### HasConditionType

`func (o *ConditionsPropertyCondition) HasConditionType() bool`

HasConditionType returns a boolean if a field has been set.

### GetTargetPropertyId

`func (o *ConditionsPropertyCondition) GetTargetPropertyId() string`

GetTargetPropertyId returns the TargetPropertyId field if non-nil, zero value otherwise.

### GetTargetPropertyIdOk

`func (o *ConditionsPropertyCondition) GetTargetPropertyIdOk() (*string, bool)`

GetTargetPropertyIdOk returns a tuple with the TargetPropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPropertyId

`func (o *ConditionsPropertyCondition) SetTargetPropertyId(v string)`

SetTargetPropertyId sets TargetPropertyId field to given value.

### HasTargetPropertyId

`func (o *ConditionsPropertyCondition) HasTargetPropertyId() bool`

HasTargetPropertyId returns a boolean if a field has been set.

### GetSimpleCondition

`func (o *ConditionsPropertyCondition) GetSimpleCondition() AttributesSimpleCondition`

GetSimpleCondition returns the SimpleCondition field if non-nil, zero value otherwise.

### GetSimpleConditionOk

`func (o *ConditionsPropertyCondition) GetSimpleConditionOk() (*AttributesSimpleCondition, bool)`

GetSimpleConditionOk returns a tuple with the SimpleCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleCondition

`func (o *ConditionsPropertyCondition) SetSimpleCondition(v AttributesSimpleCondition)`

SetSimpleCondition sets SimpleCondition field to given value.

### HasSimpleCondition

`func (o *ConditionsPropertyCondition) HasSimpleCondition() bool`

HasSimpleCondition returns a boolean if a field has been set.

### GetValueCondition

`func (o *ConditionsPropertyCondition) GetValueCondition() AttributesValueCondition`

GetValueCondition returns the ValueCondition field if non-nil, zero value otherwise.

### GetValueConditionOk

`func (o *ConditionsPropertyCondition) GetValueConditionOk() (*AttributesValueCondition, bool)`

GetValueConditionOk returns a tuple with the ValueCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCondition

`func (o *ConditionsPropertyCondition) SetValueCondition(v AttributesValueCondition)`

SetValueCondition sets ValueCondition field to given value.

### HasValueCondition

`func (o *ConditionsPropertyCondition) HasValueCondition() bool`

HasValueCondition returns a boolean if a field has been set.

### GetValue

`func (o *ConditionsPropertyCondition) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConditionsPropertyCondition) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConditionsPropertyCondition) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ConditionsPropertyCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


