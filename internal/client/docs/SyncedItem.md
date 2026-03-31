# SyncedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** |  | [optional] 
**SyncJobId** | Pointer to **int64** |  | [optional] 
**SyncJobName** | Pointer to **string** |  | [optional] 
**SyncJobDateCreated** | Pointer to **time.Time** |  | [optional] 
**SyncJobItemId** | Pointer to **int64** |  | [optional] 
**OriginalFileName** | Pointer to **string** |  | [optional] 
**Item** | Pointer to [**BaseItemDto**](BaseItemDto.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AdditionalFiles** | Pointer to [**[]ItemFileInfo**](ItemFileInfo.md) |  | [optional] 

## Methods

### NewSyncedItem

`func NewSyncedItem() *SyncedItem`

NewSyncedItem instantiates a new SyncedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncedItemWithDefaults

`func NewSyncedItemWithDefaults() *SyncedItem`

NewSyncedItemWithDefaults instantiates a new SyncedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *SyncedItem) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *SyncedItem) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *SyncedItem) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *SyncedItem) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetSyncJobId

`func (o *SyncedItem) GetSyncJobId() int64`

GetSyncJobId returns the SyncJobId field if non-nil, zero value otherwise.

### GetSyncJobIdOk

`func (o *SyncedItem) GetSyncJobIdOk() (*int64, bool)`

GetSyncJobIdOk returns a tuple with the SyncJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobId

`func (o *SyncedItem) SetSyncJobId(v int64)`

SetSyncJobId sets SyncJobId field to given value.

### HasSyncJobId

`func (o *SyncedItem) HasSyncJobId() bool`

HasSyncJobId returns a boolean if a field has been set.

### GetSyncJobName

`func (o *SyncedItem) GetSyncJobName() string`

GetSyncJobName returns the SyncJobName field if non-nil, zero value otherwise.

### GetSyncJobNameOk

`func (o *SyncedItem) GetSyncJobNameOk() (*string, bool)`

GetSyncJobNameOk returns a tuple with the SyncJobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobName

`func (o *SyncedItem) SetSyncJobName(v string)`

SetSyncJobName sets SyncJobName field to given value.

### HasSyncJobName

`func (o *SyncedItem) HasSyncJobName() bool`

HasSyncJobName returns a boolean if a field has been set.

### GetSyncJobDateCreated

`func (o *SyncedItem) GetSyncJobDateCreated() time.Time`

GetSyncJobDateCreated returns the SyncJobDateCreated field if non-nil, zero value otherwise.

### GetSyncJobDateCreatedOk

`func (o *SyncedItem) GetSyncJobDateCreatedOk() (*time.Time, bool)`

GetSyncJobDateCreatedOk returns a tuple with the SyncJobDateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobDateCreated

`func (o *SyncedItem) SetSyncJobDateCreated(v time.Time)`

SetSyncJobDateCreated sets SyncJobDateCreated field to given value.

### HasSyncJobDateCreated

`func (o *SyncedItem) HasSyncJobDateCreated() bool`

HasSyncJobDateCreated returns a boolean if a field has been set.

### GetSyncJobItemId

`func (o *SyncedItem) GetSyncJobItemId() int64`

GetSyncJobItemId returns the SyncJobItemId field if non-nil, zero value otherwise.

### GetSyncJobItemIdOk

`func (o *SyncedItem) GetSyncJobItemIdOk() (*int64, bool)`

GetSyncJobItemIdOk returns a tuple with the SyncJobItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobItemId

`func (o *SyncedItem) SetSyncJobItemId(v int64)`

SetSyncJobItemId sets SyncJobItemId field to given value.

### HasSyncJobItemId

`func (o *SyncedItem) HasSyncJobItemId() bool`

HasSyncJobItemId returns a boolean if a field has been set.

### GetOriginalFileName

`func (o *SyncedItem) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *SyncedItem) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *SyncedItem) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *SyncedItem) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### GetItem

`func (o *SyncedItem) GetItem() BaseItemDto`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *SyncedItem) GetItemOk() (*BaseItemDto, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *SyncedItem) SetItem(v BaseItemDto)`

SetItem sets Item field to given value.

### HasItem

`func (o *SyncedItem) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetUserId

`func (o *SyncedItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SyncedItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SyncedItem) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SyncedItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAdditionalFiles

`func (o *SyncedItem) GetAdditionalFiles() []ItemFileInfo`

GetAdditionalFiles returns the AdditionalFiles field if non-nil, zero value otherwise.

### GetAdditionalFilesOk

`func (o *SyncedItem) GetAdditionalFilesOk() (*[]ItemFileInfo, bool)`

GetAdditionalFilesOk returns a tuple with the AdditionalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFiles

`func (o *SyncedItem) SetAdditionalFiles(v []ItemFileInfo)`

SetAdditionalFiles sets AdditionalFiles field to given value.

### HasAdditionalFiles

`func (o *SyncedItem) HasAdditionalFiles() bool`

HasAdditionalFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


