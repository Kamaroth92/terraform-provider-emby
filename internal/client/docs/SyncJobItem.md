# SyncJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**JobId** | Pointer to **int64** |  | [optional] 
**ItemId** | Pointer to **int64** |  | [optional] 
**ItemName** | Pointer to **string** |  | [optional] 
**MediaSourceId** | Pointer to **string** |  | [optional] 
**MediaSource** | Pointer to [**MediaSourceInfo**](MediaSourceInfo.md) |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**InternalTargetId** | Pointer to **int64** |  | [optional] 
**OutputPath** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**SyncJobItemStatus**](SyncJobItemStatus.md) |  | [optional] 
**Progress** | Pointer to **NullableFloat64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**PrimaryImageItemId** | Pointer to **string** |  | [optional] 
**PrimaryImageTag** | Pointer to **string** |  | [optional] 
**TemporaryPath** | Pointer to **string** |  | [optional] 
**AdditionalFiles** | Pointer to [**[]ItemFileInfo**](ItemFileInfo.md) |  | [optional] 

## Methods

### NewSyncJobItem

`func NewSyncJobItem() *SyncJobItem`

NewSyncJobItem instantiates a new SyncJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncJobItemWithDefaults

`func NewSyncJobItemWithDefaults() *SyncJobItem`

NewSyncJobItemWithDefaults instantiates a new SyncJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyncJobItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncJobItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncJobItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SyncJobItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobId

`func (o *SyncJobItem) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SyncJobItem) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SyncJobItem) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *SyncJobItem) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetItemId

`func (o *SyncJobItem) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *SyncJobItem) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *SyncJobItem) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *SyncJobItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemName

`func (o *SyncJobItem) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *SyncJobItem) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *SyncJobItem) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *SyncJobItem) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetMediaSourceId

`func (o *SyncJobItem) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *SyncJobItem) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *SyncJobItem) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *SyncJobItem) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### GetMediaSource

`func (o *SyncJobItem) GetMediaSource() MediaSourceInfo`

GetMediaSource returns the MediaSource field if non-nil, zero value otherwise.

### GetMediaSourceOk

`func (o *SyncJobItem) GetMediaSourceOk() (*MediaSourceInfo, bool)`

GetMediaSourceOk returns a tuple with the MediaSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSource

`func (o *SyncJobItem) SetMediaSource(v MediaSourceInfo)`

SetMediaSource sets MediaSource field to given value.

### HasMediaSource

`func (o *SyncJobItem) HasMediaSource() bool`

HasMediaSource returns a boolean if a field has been set.

### GetTargetId

`func (o *SyncJobItem) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *SyncJobItem) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *SyncJobItem) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *SyncJobItem) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetInternalTargetId

`func (o *SyncJobItem) GetInternalTargetId() int64`

GetInternalTargetId returns the InternalTargetId field if non-nil, zero value otherwise.

### GetInternalTargetIdOk

`func (o *SyncJobItem) GetInternalTargetIdOk() (*int64, bool)`

GetInternalTargetIdOk returns a tuple with the InternalTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTargetId

`func (o *SyncJobItem) SetInternalTargetId(v int64)`

SetInternalTargetId sets InternalTargetId field to given value.

### HasInternalTargetId

`func (o *SyncJobItem) HasInternalTargetId() bool`

HasInternalTargetId returns a boolean if a field has been set.

### GetOutputPath

`func (o *SyncJobItem) GetOutputPath() string`

GetOutputPath returns the OutputPath field if non-nil, zero value otherwise.

### GetOutputPathOk

`func (o *SyncJobItem) GetOutputPathOk() (*string, bool)`

GetOutputPathOk returns a tuple with the OutputPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPath

`func (o *SyncJobItem) SetOutputPath(v string)`

SetOutputPath sets OutputPath field to given value.

### HasOutputPath

`func (o *SyncJobItem) HasOutputPath() bool`

HasOutputPath returns a boolean if a field has been set.

### GetStatus

`func (o *SyncJobItem) GetStatus() SyncJobItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncJobItem) GetStatusOk() (*SyncJobItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncJobItem) SetStatus(v SyncJobItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyncJobItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProgress

`func (o *SyncJobItem) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SyncJobItem) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SyncJobItem) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SyncJobItem) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *SyncJobItem) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *SyncJobItem) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetDateCreated

`func (o *SyncJobItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SyncJobItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SyncJobItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SyncJobItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetPrimaryImageItemId

`func (o *SyncJobItem) GetPrimaryImageItemId() string`

GetPrimaryImageItemId returns the PrimaryImageItemId field if non-nil, zero value otherwise.

### GetPrimaryImageItemIdOk

`func (o *SyncJobItem) GetPrimaryImageItemIdOk() (*string, bool)`

GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageItemId

`func (o *SyncJobItem) SetPrimaryImageItemId(v string)`

SetPrimaryImageItemId sets PrimaryImageItemId field to given value.

### HasPrimaryImageItemId

`func (o *SyncJobItem) HasPrimaryImageItemId() bool`

HasPrimaryImageItemId returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *SyncJobItem) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *SyncJobItem) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *SyncJobItem) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *SyncJobItem) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### GetTemporaryPath

`func (o *SyncJobItem) GetTemporaryPath() string`

GetTemporaryPath returns the TemporaryPath field if non-nil, zero value otherwise.

### GetTemporaryPathOk

`func (o *SyncJobItem) GetTemporaryPathOk() (*string, bool)`

GetTemporaryPathOk returns a tuple with the TemporaryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryPath

`func (o *SyncJobItem) SetTemporaryPath(v string)`

SetTemporaryPath sets TemporaryPath field to given value.

### HasTemporaryPath

`func (o *SyncJobItem) HasTemporaryPath() bool`

HasTemporaryPath returns a boolean if a field has been set.

### GetAdditionalFiles

`func (o *SyncJobItem) GetAdditionalFiles() []ItemFileInfo`

GetAdditionalFiles returns the AdditionalFiles field if non-nil, zero value otherwise.

### GetAdditionalFilesOk

`func (o *SyncJobItem) GetAdditionalFilesOk() (*[]ItemFileInfo, bool)`

GetAdditionalFilesOk returns a tuple with the AdditionalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFiles

`func (o *SyncJobItem) SetAdditionalFiles(v []ItemFileInfo)`

SetAdditionalFiles sets AdditionalFiles field to given value.

### HasAdditionalFiles

`func (o *SyncJobItem) HasAdditionalFiles() bool`

HasAdditionalFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


