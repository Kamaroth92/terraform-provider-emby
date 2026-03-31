# SyncDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalItemIds** | Pointer to **[]string** |  | [optional] 
**InternalTargetIds** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewSyncDataRequest

`func NewSyncDataRequest() *SyncDataRequest`

NewSyncDataRequest instantiates a new SyncDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncDataRequestWithDefaults

`func NewSyncDataRequestWithDefaults() *SyncDataRequest`

NewSyncDataRequestWithDefaults instantiates a new SyncDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalItemIds

`func (o *SyncDataRequest) GetLocalItemIds() []string`

GetLocalItemIds returns the LocalItemIds field if non-nil, zero value otherwise.

### GetLocalItemIdsOk

`func (o *SyncDataRequest) GetLocalItemIdsOk() (*[]string, bool)`

GetLocalItemIdsOk returns a tuple with the LocalItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalItemIds

`func (o *SyncDataRequest) SetLocalItemIds(v []string)`

SetLocalItemIds sets LocalItemIds field to given value.

### HasLocalItemIds

`func (o *SyncDataRequest) HasLocalItemIds() bool`

HasLocalItemIds returns a boolean if a field has been set.

### GetInternalTargetIds

`func (o *SyncDataRequest) GetInternalTargetIds() []int64`

GetInternalTargetIds returns the InternalTargetIds field if non-nil, zero value otherwise.

### GetInternalTargetIdsOk

`func (o *SyncDataRequest) GetInternalTargetIdsOk() (*[]int64, bool)`

GetInternalTargetIdsOk returns a tuple with the InternalTargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTargetIds

`func (o *SyncDataRequest) SetInternalTargetIds(v []int64)`

SetInternalTargetIds sets InternalTargetIds field to given value.

### HasInternalTargetIds

`func (o *SyncDataRequest) HasInternalTargetIds() bool`

HasInternalTargetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


