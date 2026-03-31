# SyncJobCreationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**SyncJob**](SyncJob.md) |  | [optional] 
**JobItems** | Pointer to [**[]SyncJobItem**](SyncJobItem.md) |  | [optional] 

## Methods

### NewSyncJobCreationResult

`func NewSyncJobCreationResult() *SyncJobCreationResult`

NewSyncJobCreationResult instantiates a new SyncJobCreationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncJobCreationResultWithDefaults

`func NewSyncJobCreationResultWithDefaults() *SyncJobCreationResult`

NewSyncJobCreationResultWithDefaults instantiates a new SyncJobCreationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *SyncJobCreationResult) GetJob() SyncJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *SyncJobCreationResult) GetJobOk() (*SyncJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *SyncJobCreationResult) SetJob(v SyncJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *SyncJobCreationResult) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobItems

`func (o *SyncJobCreationResult) GetJobItems() []SyncJobItem`

GetJobItems returns the JobItems field if non-nil, zero value otherwise.

### GetJobItemsOk

`func (o *SyncJobCreationResult) GetJobItemsOk() (*[]SyncJobItem, bool)`

GetJobItemsOk returns a tuple with the JobItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobItems

`func (o *SyncJobCreationResult) SetJobItems(v []SyncJobItem)`

SetJobItems sets JobItems field to given value.

### HasJobItems

`func (o *SyncJobCreationResult) HasJobItems() bool`

HasJobItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


