# TranscodingVpStepInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepType** | Pointer to [**TranscodingVpStepTypes**](TranscodingVpStepTypes.md) |  | [optional] 
**StepTypeName** | Pointer to **string** |  | [optional] 
**HardwareContextName** | Pointer to **string** |  | [optional] 
**IsHardwareContext** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Short** | Pointer to **string** |  | [optional] 
**FfmpegName** | Pointer to **string** |  | [optional] 
**FfmpegDescription** | Pointer to **string** |  | [optional] 
**FfmpegOptions** | Pointer to **string** |  | [optional] 
**Param** | Pointer to **string** |  | [optional] 
**ParamShort** | Pointer to **string** |  | [optional] 

## Methods

### NewTranscodingVpStepInfo

`func NewTranscodingVpStepInfo() *TranscodingVpStepInfo`

NewTranscodingVpStepInfo instantiates a new TranscodingVpStepInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscodingVpStepInfoWithDefaults

`func NewTranscodingVpStepInfoWithDefaults() *TranscodingVpStepInfo`

NewTranscodingVpStepInfoWithDefaults instantiates a new TranscodingVpStepInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepType

`func (o *TranscodingVpStepInfo) GetStepType() TranscodingVpStepTypes`

GetStepType returns the StepType field if non-nil, zero value otherwise.

### GetStepTypeOk

`func (o *TranscodingVpStepInfo) GetStepTypeOk() (*TranscodingVpStepTypes, bool)`

GetStepTypeOk returns a tuple with the StepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepType

`func (o *TranscodingVpStepInfo) SetStepType(v TranscodingVpStepTypes)`

SetStepType sets StepType field to given value.

### HasStepType

`func (o *TranscodingVpStepInfo) HasStepType() bool`

HasStepType returns a boolean if a field has been set.

### GetStepTypeName

`func (o *TranscodingVpStepInfo) GetStepTypeName() string`

GetStepTypeName returns the StepTypeName field if non-nil, zero value otherwise.

### GetStepTypeNameOk

`func (o *TranscodingVpStepInfo) GetStepTypeNameOk() (*string, bool)`

GetStepTypeNameOk returns a tuple with the StepTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepTypeName

`func (o *TranscodingVpStepInfo) SetStepTypeName(v string)`

SetStepTypeName sets StepTypeName field to given value.

### HasStepTypeName

`func (o *TranscodingVpStepInfo) HasStepTypeName() bool`

HasStepTypeName returns a boolean if a field has been set.

### GetHardwareContextName

`func (o *TranscodingVpStepInfo) GetHardwareContextName() string`

GetHardwareContextName returns the HardwareContextName field if non-nil, zero value otherwise.

### GetHardwareContextNameOk

`func (o *TranscodingVpStepInfo) GetHardwareContextNameOk() (*string, bool)`

GetHardwareContextNameOk returns a tuple with the HardwareContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareContextName

`func (o *TranscodingVpStepInfo) SetHardwareContextName(v string)`

SetHardwareContextName sets HardwareContextName field to given value.

### HasHardwareContextName

`func (o *TranscodingVpStepInfo) HasHardwareContextName() bool`

HasHardwareContextName returns a boolean if a field has been set.

### GetIsHardwareContext

`func (o *TranscodingVpStepInfo) GetIsHardwareContext() bool`

GetIsHardwareContext returns the IsHardwareContext field if non-nil, zero value otherwise.

### GetIsHardwareContextOk

`func (o *TranscodingVpStepInfo) GetIsHardwareContextOk() (*bool, bool)`

GetIsHardwareContextOk returns a tuple with the IsHardwareContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHardwareContext

`func (o *TranscodingVpStepInfo) SetIsHardwareContext(v bool)`

SetIsHardwareContext sets IsHardwareContext field to given value.

### HasIsHardwareContext

`func (o *TranscodingVpStepInfo) HasIsHardwareContext() bool`

HasIsHardwareContext returns a boolean if a field has been set.

### GetName

`func (o *TranscodingVpStepInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TranscodingVpStepInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TranscodingVpStepInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TranscodingVpStepInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShort

`func (o *TranscodingVpStepInfo) GetShort() string`

GetShort returns the Short field if non-nil, zero value otherwise.

### GetShortOk

`func (o *TranscodingVpStepInfo) GetShortOk() (*string, bool)`

GetShortOk returns a tuple with the Short field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShort

`func (o *TranscodingVpStepInfo) SetShort(v string)`

SetShort sets Short field to given value.

### HasShort

`func (o *TranscodingVpStepInfo) HasShort() bool`

HasShort returns a boolean if a field has been set.

### GetFfmpegName

`func (o *TranscodingVpStepInfo) GetFfmpegName() string`

GetFfmpegName returns the FfmpegName field if non-nil, zero value otherwise.

### GetFfmpegNameOk

`func (o *TranscodingVpStepInfo) GetFfmpegNameOk() (*string, bool)`

GetFfmpegNameOk returns a tuple with the FfmpegName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfmpegName

`func (o *TranscodingVpStepInfo) SetFfmpegName(v string)`

SetFfmpegName sets FfmpegName field to given value.

### HasFfmpegName

`func (o *TranscodingVpStepInfo) HasFfmpegName() bool`

HasFfmpegName returns a boolean if a field has been set.

### GetFfmpegDescription

`func (o *TranscodingVpStepInfo) GetFfmpegDescription() string`

GetFfmpegDescription returns the FfmpegDescription field if non-nil, zero value otherwise.

### GetFfmpegDescriptionOk

`func (o *TranscodingVpStepInfo) GetFfmpegDescriptionOk() (*string, bool)`

GetFfmpegDescriptionOk returns a tuple with the FfmpegDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfmpegDescription

`func (o *TranscodingVpStepInfo) SetFfmpegDescription(v string)`

SetFfmpegDescription sets FfmpegDescription field to given value.

### HasFfmpegDescription

`func (o *TranscodingVpStepInfo) HasFfmpegDescription() bool`

HasFfmpegDescription returns a boolean if a field has been set.

### GetFfmpegOptions

`func (o *TranscodingVpStepInfo) GetFfmpegOptions() string`

GetFfmpegOptions returns the FfmpegOptions field if non-nil, zero value otherwise.

### GetFfmpegOptionsOk

`func (o *TranscodingVpStepInfo) GetFfmpegOptionsOk() (*string, bool)`

GetFfmpegOptionsOk returns a tuple with the FfmpegOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfmpegOptions

`func (o *TranscodingVpStepInfo) SetFfmpegOptions(v string)`

SetFfmpegOptions sets FfmpegOptions field to given value.

### HasFfmpegOptions

`func (o *TranscodingVpStepInfo) HasFfmpegOptions() bool`

HasFfmpegOptions returns a boolean if a field has been set.

### GetParam

`func (o *TranscodingVpStepInfo) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *TranscodingVpStepInfo) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *TranscodingVpStepInfo) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *TranscodingVpStepInfo) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetParamShort

`func (o *TranscodingVpStepInfo) GetParamShort() string`

GetParamShort returns the ParamShort field if non-nil, zero value otherwise.

### GetParamShortOk

`func (o *TranscodingVpStepInfo) GetParamShortOk() (*string, bool)`

GetParamShortOk returns a tuple with the ParamShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamShort

`func (o *TranscodingVpStepInfo) SetParamShort(v string)`

SetParamShort sets ParamShort field to given value.

### HasParamShort

`func (o *TranscodingVpStepInfo) HasParamShort() bool`

HasParamShort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


