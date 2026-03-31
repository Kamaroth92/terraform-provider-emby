# EditorsEditorBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewEditorsEditorBase

`func NewEditorsEditorBase() *EditorsEditorBase`

NewEditorsEditorBase instantiates a new EditorsEditorBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditorsEditorBaseWithDefaults

`func NewEditorsEditorBaseWithDefaults() *EditorsEditorBase`

NewEditorsEditorBaseWithDefaults instantiates a new EditorsEditorBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditorType

`func (o *EditorsEditorBase) GetEditorType() CommonEditorTypes`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *EditorsEditorBase) GetEditorTypeOk() (*CommonEditorTypes, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *EditorsEditorBase) SetEditorType(v CommonEditorTypes)`

SetEditorType sets EditorType field to given value.

### HasEditorType

`func (o *EditorsEditorBase) HasEditorType() bool`

HasEditorType returns a boolean if a field has been set.

### GetName

`func (o *EditorsEditorBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditorsEditorBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditorsEditorBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditorsEditorBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *EditorsEditorBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditorsEditorBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditorsEditorBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EditorsEditorBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllowEmpty

`func (o *EditorsEditorBase) GetAllowEmpty() bool`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *EditorsEditorBase) GetAllowEmptyOk() (*bool, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *EditorsEditorBase) SetAllowEmpty(v bool)`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *EditorsEditorBase) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *EditorsEditorBase) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *EditorsEditorBase) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *EditorsEditorBase) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *EditorsEditorBase) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetIsAdvanced

`func (o *EditorsEditorBase) GetIsAdvanced() bool`

GetIsAdvanced returns the IsAdvanced field if non-nil, zero value otherwise.

### GetIsAdvancedOk

`func (o *EditorsEditorBase) GetIsAdvancedOk() (*bool, bool)`

GetIsAdvancedOk returns a tuple with the IsAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdvanced

`func (o *EditorsEditorBase) SetIsAdvanced(v bool)`

SetIsAdvanced sets IsAdvanced field to given value.

### HasIsAdvanced

`func (o *EditorsEditorBase) HasIsAdvanced() bool`

HasIsAdvanced returns a boolean if a field has been set.

### GetDisplayName

`func (o *EditorsEditorBase) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EditorsEditorBase) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EditorsEditorBase) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EditorsEditorBase) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *EditorsEditorBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditorsEditorBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditorsEditorBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditorsEditorBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatureRequiresPremiere

`func (o *EditorsEditorBase) GetFeatureRequiresPremiere() bool`

GetFeatureRequiresPremiere returns the FeatureRequiresPremiere field if non-nil, zero value otherwise.

### GetFeatureRequiresPremiereOk

`func (o *EditorsEditorBase) GetFeatureRequiresPremiereOk() (*bool, bool)`

GetFeatureRequiresPremiereOk returns a tuple with the FeatureRequiresPremiere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureRequiresPremiere

`func (o *EditorsEditorBase) SetFeatureRequiresPremiere(v bool)`

SetFeatureRequiresPremiere sets FeatureRequiresPremiere field to given value.

### HasFeatureRequiresPremiere

`func (o *EditorsEditorBase) HasFeatureRequiresPremiere() bool`

HasFeatureRequiresPremiere returns a boolean if a field has been set.

### GetParentId

`func (o *EditorsEditorBase) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *EditorsEditorBase) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *EditorsEditorBase) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *EditorsEditorBase) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


