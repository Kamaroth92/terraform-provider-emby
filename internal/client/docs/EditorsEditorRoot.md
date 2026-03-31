# EditorsEditorRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyConditions** | Pointer to [**[]ConditionsPropertyCondition**](ConditionsPropertyCondition.md) |  | [optional] 
**PostbackActions** | Pointer to [**[]ActionsPostbackAction**](ActionsPostbackAction.md) |  | [optional] 
**TitleButton** | Pointer to [**EditorsEditorButtonItem**](EditorsEditorButtonItem.md) |  | [optional] 
**EditorItems** | Pointer to [**[]EditorsEditorBase**](EditorsEditorBase.md) |  | [optional] 
**EditorType** | Pointer to [**CommonEditorTypes**](CommonEditorTypes.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**AllowEmpty** | Pointer to **bool** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 
**IsAdvanced** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FeatureRequiresPremiere** | Pointer to **bool** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 

## Methods

### NewEditorsEditorRoot

`func NewEditorsEditorRoot() *EditorsEditorRoot`

NewEditorsEditorRoot instantiates a new EditorsEditorRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditorsEditorRootWithDefaults

`func NewEditorsEditorRootWithDefaults() *EditorsEditorRoot`

NewEditorsEditorRootWithDefaults instantiates a new EditorsEditorRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyConditions

`func (o *EditorsEditorRoot) GetPropertyConditions() []ConditionsPropertyCondition`

GetPropertyConditions returns the PropertyConditions field if non-nil, zero value otherwise.

### GetPropertyConditionsOk

`func (o *EditorsEditorRoot) GetPropertyConditionsOk() (*[]ConditionsPropertyCondition, bool)`

GetPropertyConditionsOk returns a tuple with the PropertyConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyConditions

`func (o *EditorsEditorRoot) SetPropertyConditions(v []ConditionsPropertyCondition)`

SetPropertyConditions sets PropertyConditions field to given value.

### HasPropertyConditions

`func (o *EditorsEditorRoot) HasPropertyConditions() bool`

HasPropertyConditions returns a boolean if a field has been set.

### GetPostbackActions

`func (o *EditorsEditorRoot) GetPostbackActions() []ActionsPostbackAction`

GetPostbackActions returns the PostbackActions field if non-nil, zero value otherwise.

### GetPostbackActionsOk

`func (o *EditorsEditorRoot) GetPostbackActionsOk() (*[]ActionsPostbackAction, bool)`

GetPostbackActionsOk returns a tuple with the PostbackActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostbackActions

`func (o *EditorsEditorRoot) SetPostbackActions(v []ActionsPostbackAction)`

SetPostbackActions sets PostbackActions field to given value.

### HasPostbackActions

`func (o *EditorsEditorRoot) HasPostbackActions() bool`

HasPostbackActions returns a boolean if a field has been set.

### GetTitleButton

`func (o *EditorsEditorRoot) GetTitleButton() EditorsEditorButtonItem`

GetTitleButton returns the TitleButton field if non-nil, zero value otherwise.

### GetTitleButtonOk

`func (o *EditorsEditorRoot) GetTitleButtonOk() (*EditorsEditorButtonItem, bool)`

GetTitleButtonOk returns a tuple with the TitleButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleButton

`func (o *EditorsEditorRoot) SetTitleButton(v EditorsEditorButtonItem)`

SetTitleButton sets TitleButton field to given value.

### HasTitleButton

`func (o *EditorsEditorRoot) HasTitleButton() bool`

HasTitleButton returns a boolean if a field has been set.

### GetEditorItems

`func (o *EditorsEditorRoot) GetEditorItems() []EditorsEditorBase`

GetEditorItems returns the EditorItems field if non-nil, zero value otherwise.

### GetEditorItemsOk

`func (o *EditorsEditorRoot) GetEditorItemsOk() (*[]EditorsEditorBase, bool)`

GetEditorItemsOk returns a tuple with the EditorItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorItems

`func (o *EditorsEditorRoot) SetEditorItems(v []EditorsEditorBase)`

SetEditorItems sets EditorItems field to given value.

### HasEditorItems

`func (o *EditorsEditorRoot) HasEditorItems() bool`

HasEditorItems returns a boolean if a field has been set.

### GetEditorType

`func (o *EditorsEditorRoot) GetEditorType() CommonEditorTypes`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *EditorsEditorRoot) GetEditorTypeOk() (*CommonEditorTypes, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *EditorsEditorRoot) SetEditorType(v CommonEditorTypes)`

SetEditorType sets EditorType field to given value.

### HasEditorType

`func (o *EditorsEditorRoot) HasEditorType() bool`

HasEditorType returns a boolean if a field has been set.

### GetName

`func (o *EditorsEditorRoot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditorsEditorRoot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditorsEditorRoot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditorsEditorRoot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *EditorsEditorRoot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditorsEditorRoot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditorsEditorRoot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EditorsEditorRoot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllowEmpty

`func (o *EditorsEditorRoot) GetAllowEmpty() bool`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *EditorsEditorRoot) GetAllowEmptyOk() (*bool, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *EditorsEditorRoot) SetAllowEmpty(v bool)`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *EditorsEditorRoot) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *EditorsEditorRoot) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *EditorsEditorRoot) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *EditorsEditorRoot) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *EditorsEditorRoot) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetIsAdvanced

`func (o *EditorsEditorRoot) GetIsAdvanced() bool`

GetIsAdvanced returns the IsAdvanced field if non-nil, zero value otherwise.

### GetIsAdvancedOk

`func (o *EditorsEditorRoot) GetIsAdvancedOk() (*bool, bool)`

GetIsAdvancedOk returns a tuple with the IsAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdvanced

`func (o *EditorsEditorRoot) SetIsAdvanced(v bool)`

SetIsAdvanced sets IsAdvanced field to given value.

### HasIsAdvanced

`func (o *EditorsEditorRoot) HasIsAdvanced() bool`

HasIsAdvanced returns a boolean if a field has been set.

### GetDisplayName

`func (o *EditorsEditorRoot) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EditorsEditorRoot) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EditorsEditorRoot) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EditorsEditorRoot) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *EditorsEditorRoot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditorsEditorRoot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditorsEditorRoot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditorsEditorRoot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatureRequiresPremiere

`func (o *EditorsEditorRoot) GetFeatureRequiresPremiere() bool`

GetFeatureRequiresPremiere returns the FeatureRequiresPremiere field if non-nil, zero value otherwise.

### GetFeatureRequiresPremiereOk

`func (o *EditorsEditorRoot) GetFeatureRequiresPremiereOk() (*bool, bool)`

GetFeatureRequiresPremiereOk returns a tuple with the FeatureRequiresPremiere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureRequiresPremiere

`func (o *EditorsEditorRoot) SetFeatureRequiresPremiere(v bool)`

SetFeatureRequiresPremiere sets FeatureRequiresPremiere field to given value.

### HasFeatureRequiresPremiere

`func (o *EditorsEditorRoot) HasFeatureRequiresPremiere() bool`

HasFeatureRequiresPremiere returns a boolean if a field has been set.

### GetParentId

`func (o *EditorsEditorRoot) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *EditorsEditorRoot) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *EditorsEditorRoot) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *EditorsEditorRoot) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


