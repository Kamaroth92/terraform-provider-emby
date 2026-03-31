# DisplayPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SortBy** | Pointer to **string** |  | [optional] 
**CustomPrefs** | Pointer to **map[string]string** |  | [optional] 
**SortOrder** | Pointer to [**SortOrder**](SortOrder.md) |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 

## Methods

### NewDisplayPreferences

`func NewDisplayPreferences() *DisplayPreferences`

NewDisplayPreferences instantiates a new DisplayPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisplayPreferencesWithDefaults

`func NewDisplayPreferencesWithDefaults() *DisplayPreferences`

NewDisplayPreferencesWithDefaults instantiates a new DisplayPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DisplayPreferences) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisplayPreferences) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisplayPreferences) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DisplayPreferences) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSortBy

`func (o *DisplayPreferences) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *DisplayPreferences) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *DisplayPreferences) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *DisplayPreferences) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetCustomPrefs

`func (o *DisplayPreferences) GetCustomPrefs() map[string]string`

GetCustomPrefs returns the CustomPrefs field if non-nil, zero value otherwise.

### GetCustomPrefsOk

`func (o *DisplayPreferences) GetCustomPrefsOk() (*map[string]string, bool)`

GetCustomPrefsOk returns a tuple with the CustomPrefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrefs

`func (o *DisplayPreferences) SetCustomPrefs(v map[string]string)`

SetCustomPrefs sets CustomPrefs field to given value.

### HasCustomPrefs

`func (o *DisplayPreferences) HasCustomPrefs() bool`

HasCustomPrefs returns a boolean if a field has been set.

### GetSortOrder

`func (o *DisplayPreferences) GetSortOrder() SortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *DisplayPreferences) GetSortOrderOk() (*SortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *DisplayPreferences) SetSortOrder(v SortOrder)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *DisplayPreferences) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetClient

`func (o *DisplayPreferences) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *DisplayPreferences) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *DisplayPreferences) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *DisplayPreferences) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


