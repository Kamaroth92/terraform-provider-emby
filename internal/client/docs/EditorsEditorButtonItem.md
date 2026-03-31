# EditorsEditorButtonItem

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

### NewEditorsEditorButtonItem

`func NewEditorsEditorButtonItem() *EditorsEditorButtonItem`

NewEditorsEditorButtonItem instantiates a new EditorsEditorButtonItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditorsEditorButtonItemWithDefaults

`func NewEditorsEditorButtonItemWithDefaults() *EditorsEditorButtonItem`

NewEditorsEditorButtonItemWithDefaults instantiates a new EditorsEditorButtonItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditorType

`func (o *EditorsEditorButtonItem) GetEditorType() CommonEditorTypes`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *EditorsEditorButtonItem) GetEditorTypeOk() (*CommonEditorTypes, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *EditorsEditorButtonItem) SetEditorType(v CommonEditorTypes)`

SetEditorType sets EditorType field to given value.

### HasEditorType

`func (o *EditorsEditorButtonItem) HasEditorType() bool`

HasEditorType returns a boolean if a field has been set.

### GetName

`func (o *EditorsEditorButtonItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditorsEditorButtonItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditorsEditorButtonItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditorsEditorButtonItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *EditorsEditorButtonItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditorsEditorButtonItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditorsEditorButtonItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EditorsEditorButtonItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllowEmpty

`func (o *EditorsEditorButtonItem) GetAllowEmpty() bool`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *EditorsEditorButtonItem) GetAllowEmptyOk() (*bool, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *EditorsEditorButtonItem) SetAllowEmpty(v bool)`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *EditorsEditorButtonItem) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *EditorsEditorButtonItem) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *EditorsEditorButtonItem) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *EditorsEditorButtonItem) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *EditorsEditorButtonItem) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetIsAdvanced

`func (o *EditorsEditorButtonItem) GetIsAdvanced() bool`

GetIsAdvanced returns the IsAdvanced field if non-nil, zero value otherwise.

### GetIsAdvancedOk

`func (o *EditorsEditorButtonItem) GetIsAdvancedOk() (*bool, bool)`

GetIsAdvancedOk returns a tuple with the IsAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdvanced

`func (o *EditorsEditorButtonItem) SetIsAdvanced(v bool)`

SetIsAdvanced sets IsAdvanced field to given value.

### HasIsAdvanced

`func (o *EditorsEditorButtonItem) HasIsAdvanced() bool`

HasIsAdvanced returns a boolean if a field has been set.

### GetDisplayName

`func (o *EditorsEditorButtonItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EditorsEditorButtonItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EditorsEditorButtonItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EditorsEditorButtonItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *EditorsEditorButtonItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditorsEditorButtonItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditorsEditorButtonItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditorsEditorButtonItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatureRequiresPremiere

`func (o *EditorsEditorButtonItem) GetFeatureRequiresPremiere() bool`

GetFeatureRequiresPremiere returns the FeatureRequiresPremiere field if non-nil, zero value otherwise.

### GetFeatureRequiresPremiereOk

`func (o *EditorsEditorButtonItem) GetFeatureRequiresPremiereOk() (*bool, bool)`

GetFeatureRequiresPremiereOk returns a tuple with the FeatureRequiresPremiere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureRequiresPremiere

`func (o *EditorsEditorButtonItem) SetFeatureRequiresPremiere(v bool)`

SetFeatureRequiresPremiere sets FeatureRequiresPremiere field to given value.

### HasFeatureRequiresPremiere

`func (o *EditorsEditorButtonItem) HasFeatureRequiresPremiere() bool`

HasFeatureRequiresPremiere returns a boolean if a field has been set.

### GetParentId

`func (o *EditorsEditorButtonItem) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *EditorsEditorButtonItem) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *EditorsEditorButtonItem) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *EditorsEditorButtonItem) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


