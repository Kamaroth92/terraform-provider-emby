# QueryResultSyncJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]SyncJobItem**](SyncJobItem.md) |  | [optional] 
**TotalRecordCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewQueryResultSyncJobItem

`func NewQueryResultSyncJobItem() *QueryResultSyncJobItem`

NewQueryResultSyncJobItem instantiates a new QueryResultSyncJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResultSyncJobItemWithDefaults

`func NewQueryResultSyncJobItemWithDefaults() *QueryResultSyncJobItem`

NewQueryResultSyncJobItemWithDefaults instantiates a new QueryResultSyncJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *QueryResultSyncJobItem) GetItems() []SyncJobItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *QueryResultSyncJobItem) GetItemsOk() (*[]SyncJobItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *QueryResultSyncJobItem) SetItems(v []SyncJobItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *QueryResultSyncJobItem) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *QueryResultSyncJobItem) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *QueryResultSyncJobItem) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *QueryResultSyncJobItem) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *QueryResultSyncJobItem) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


