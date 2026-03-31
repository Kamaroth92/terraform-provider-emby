# SyncedItemProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Progress** | Pointer to **NullableFloat64** |  | [optional] 
**Status** | Pointer to [**SyncJobItemStatus**](SyncJobItemStatus.md) |  | [optional] 

## Methods

### NewSyncedItemProgress

`func NewSyncedItemProgress() *SyncedItemProgress`

NewSyncedItemProgress instantiates a new SyncedItemProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncedItemProgressWithDefaults

`func NewSyncedItemProgressWithDefaults() *SyncedItemProgress`

NewSyncedItemProgressWithDefaults instantiates a new SyncedItemProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgress

`func (o *SyncedItemProgress) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SyncedItemProgress) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SyncedItemProgress) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SyncedItemProgress) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *SyncedItemProgress) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *SyncedItemProgress) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetStatus

`func (o *SyncedItemProgress) GetStatus() SyncJobItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncedItemProgress) GetStatusOk() (*SyncJobItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncedItemProgress) SetStatus(v SyncJobItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyncedItemProgress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


