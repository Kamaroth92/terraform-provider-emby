# UICommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | Pointer to [**EnumsUICommandType**](EnumsUICommandType.md) |  | [optional] 
**CommandId** | Pointer to **string** |  | [optional] 
**IsVisible** | Pointer to **bool** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Caption** | Pointer to **string** |  | [optional] 
**SetFocus** | Pointer to **bool** |  | [optional] 
**ConfirmationPrompt** | Pointer to **string** |  | [optional] 

## Methods

### NewUICommand

`func NewUICommand() *UICommand`

NewUICommand instantiates a new UICommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUICommandWithDefaults

`func NewUICommandWithDefaults() *UICommand`

NewUICommandWithDefaults instantiates a new UICommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *UICommand) GetCommandType() EnumsUICommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *UICommand) GetCommandTypeOk() (*EnumsUICommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *UICommand) SetCommandType(v EnumsUICommandType)`

SetCommandType sets CommandType field to given value.

### HasCommandType

`func (o *UICommand) HasCommandType() bool`

HasCommandType returns a boolean if a field has been set.

### GetCommandId

`func (o *UICommand) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *UICommand) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *UICommand) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *UICommand) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetIsVisible

`func (o *UICommand) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *UICommand) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *UICommand) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.

### HasIsVisible

`func (o *UICommand) HasIsVisible() bool`

HasIsVisible returns a boolean if a field has been set.

### GetIsEnabled

`func (o *UICommand) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UICommand) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UICommand) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *UICommand) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetCaption

`func (o *UICommand) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *UICommand) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *UICommand) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *UICommand) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetSetFocus

`func (o *UICommand) GetSetFocus() bool`

GetSetFocus returns the SetFocus field if non-nil, zero value otherwise.

### GetSetFocusOk

`func (o *UICommand) GetSetFocusOk() (*bool, bool)`

GetSetFocusOk returns a tuple with the SetFocus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetFocus

`func (o *UICommand) SetSetFocus(v bool)`

SetSetFocus sets SetFocus field to given value.

### HasSetFocus

`func (o *UICommand) HasSetFocus() bool`

HasSetFocus returns a boolean if a field has been set.

### GetConfirmationPrompt

`func (o *UICommand) GetConfirmationPrompt() string`

GetConfirmationPrompt returns the ConfirmationPrompt field if non-nil, zero value otherwise.

### GetConfirmationPromptOk

`func (o *UICommand) GetConfirmationPromptOk() (*string, bool)`

GetConfirmationPromptOk returns a tuple with the ConfirmationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationPrompt

`func (o *UICommand) SetConfirmationPrompt(v string)`

SetConfirmationPrompt sets ConfirmationPrompt field to given value.

### HasConfirmationPrompt

`func (o *UICommand) HasConfirmationPrompt() bool`

HasConfirmationPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


