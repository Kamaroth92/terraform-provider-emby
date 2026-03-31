# EditObjectContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultObject** | Pointer to **map[string]interface{}** |  | [optional] 
**TypeName** | Pointer to **string** |  | [optional] 
**EditorRoot** | Pointer to [**EditorsEditorRoot**](EditorsEditorRoot.md) |  | [optional] 

## Methods

### NewEditObjectContainer

`func NewEditObjectContainer() *EditObjectContainer`

NewEditObjectContainer instantiates a new EditObjectContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditObjectContainerWithDefaults

`func NewEditObjectContainerWithDefaults() *EditObjectContainer`

NewEditObjectContainerWithDefaults instantiates a new EditObjectContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *EditObjectContainer) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *EditObjectContainer) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *EditObjectContainer) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.

### HasObject

`func (o *EditObjectContainer) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetDefaultObject

`func (o *EditObjectContainer) GetDefaultObject() map[string]interface{}`

GetDefaultObject returns the DefaultObject field if non-nil, zero value otherwise.

### GetDefaultObjectOk

`func (o *EditObjectContainer) GetDefaultObjectOk() (*map[string]interface{}, bool)`

GetDefaultObjectOk returns a tuple with the DefaultObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultObject

`func (o *EditObjectContainer) SetDefaultObject(v map[string]interface{})`

SetDefaultObject sets DefaultObject field to given value.

### HasDefaultObject

`func (o *EditObjectContainer) HasDefaultObject() bool`

HasDefaultObject returns a boolean if a field has been set.

### GetTypeName

`func (o *EditObjectContainer) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *EditObjectContainer) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *EditObjectContainer) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *EditObjectContainer) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetEditorRoot

`func (o *EditObjectContainer) GetEditorRoot() EditorsEditorRoot`

GetEditorRoot returns the EditorRoot field if non-nil, zero value otherwise.

### GetEditorRootOk

`func (o *EditObjectContainer) GetEditorRootOk() (*EditorsEditorRoot, bool)`

GetEditorRootOk returns a tuple with the EditorRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorRoot

`func (o *EditObjectContainer) SetEditorRoot(v EditorsEditorRoot)`

SetEditorRoot sets EditorRoot field to given value.

### HasEditorRoot

`func (o *EditObjectContainer) HasEditorRoot() bool`

HasEditorRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


