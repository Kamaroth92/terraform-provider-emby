# ApiOnPlaybackProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaylistIndex** | Pointer to **int32** |  | [optional] 
**PlaylistLength** | Pointer to **int32** |  | [optional] 
**Shuffle** | Pointer to **bool** |  | [optional] 
**SleepTimerMode** | Pointer to [**SleepTimerMode**](SleepTimerMode.md) |  | [optional] 
**SleepTimerEndTime** | Pointer to **NullableTime** |  | [optional] 
**EventName** | Pointer to [**ProgressEvent**](ProgressEvent.md) |  | [optional] 

## Methods

### NewApiOnPlaybackProgress

`func NewApiOnPlaybackProgress() *ApiOnPlaybackProgress`

NewApiOnPlaybackProgress instantiates a new ApiOnPlaybackProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOnPlaybackProgressWithDefaults

`func NewApiOnPlaybackProgressWithDefaults() *ApiOnPlaybackProgress`

NewApiOnPlaybackProgressWithDefaults instantiates a new ApiOnPlaybackProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaylistIndex

`func (o *ApiOnPlaybackProgress) GetPlaylistIndex() int32`

GetPlaylistIndex returns the PlaylistIndex field if non-nil, zero value otherwise.

### GetPlaylistIndexOk

`func (o *ApiOnPlaybackProgress) GetPlaylistIndexOk() (*int32, bool)`

GetPlaylistIndexOk returns a tuple with the PlaylistIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistIndex

`func (o *ApiOnPlaybackProgress) SetPlaylistIndex(v int32)`

SetPlaylistIndex sets PlaylistIndex field to given value.

### HasPlaylistIndex

`func (o *ApiOnPlaybackProgress) HasPlaylistIndex() bool`

HasPlaylistIndex returns a boolean if a field has been set.

### GetPlaylistLength

`func (o *ApiOnPlaybackProgress) GetPlaylistLength() int32`

GetPlaylistLength returns the PlaylistLength field if non-nil, zero value otherwise.

### GetPlaylistLengthOk

`func (o *ApiOnPlaybackProgress) GetPlaylistLengthOk() (*int32, bool)`

GetPlaylistLengthOk returns a tuple with the PlaylistLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistLength

`func (o *ApiOnPlaybackProgress) SetPlaylistLength(v int32)`

SetPlaylistLength sets PlaylistLength field to given value.

### HasPlaylistLength

`func (o *ApiOnPlaybackProgress) HasPlaylistLength() bool`

HasPlaylistLength returns a boolean if a field has been set.

### GetShuffle

`func (o *ApiOnPlaybackProgress) GetShuffle() bool`

GetShuffle returns the Shuffle field if non-nil, zero value otherwise.

### GetShuffleOk

`func (o *ApiOnPlaybackProgress) GetShuffleOk() (*bool, bool)`

GetShuffleOk returns a tuple with the Shuffle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShuffle

`func (o *ApiOnPlaybackProgress) SetShuffle(v bool)`

SetShuffle sets Shuffle field to given value.

### HasShuffle

`func (o *ApiOnPlaybackProgress) HasShuffle() bool`

HasShuffle returns a boolean if a field has been set.

### GetSleepTimerMode

`func (o *ApiOnPlaybackProgress) GetSleepTimerMode() SleepTimerMode`

GetSleepTimerMode returns the SleepTimerMode field if non-nil, zero value otherwise.

### GetSleepTimerModeOk

`func (o *ApiOnPlaybackProgress) GetSleepTimerModeOk() (*SleepTimerMode, bool)`

GetSleepTimerModeOk returns a tuple with the SleepTimerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepTimerMode

`func (o *ApiOnPlaybackProgress) SetSleepTimerMode(v SleepTimerMode)`

SetSleepTimerMode sets SleepTimerMode field to given value.

### HasSleepTimerMode

`func (o *ApiOnPlaybackProgress) HasSleepTimerMode() bool`

HasSleepTimerMode returns a boolean if a field has been set.

### GetSleepTimerEndTime

`func (o *ApiOnPlaybackProgress) GetSleepTimerEndTime() time.Time`

GetSleepTimerEndTime returns the SleepTimerEndTime field if non-nil, zero value otherwise.

### GetSleepTimerEndTimeOk

`func (o *ApiOnPlaybackProgress) GetSleepTimerEndTimeOk() (*time.Time, bool)`

GetSleepTimerEndTimeOk returns a tuple with the SleepTimerEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepTimerEndTime

`func (o *ApiOnPlaybackProgress) SetSleepTimerEndTime(v time.Time)`

SetSleepTimerEndTime sets SleepTimerEndTime field to given value.

### HasSleepTimerEndTime

`func (o *ApiOnPlaybackProgress) HasSleepTimerEndTime() bool`

HasSleepTimerEndTime returns a boolean if a field has been set.

### SetSleepTimerEndTimeNil

`func (o *ApiOnPlaybackProgress) SetSleepTimerEndTimeNil(b bool)`

 SetSleepTimerEndTimeNil sets the value for SleepTimerEndTime to be an explicit nil

### UnsetSleepTimerEndTime
`func (o *ApiOnPlaybackProgress) UnsetSleepTimerEndTime()`

UnsetSleepTimerEndTime ensures that no value is present for SleepTimerEndTime, not even an explicit nil
### GetEventName

`func (o *ApiOnPlaybackProgress) GetEventName() ProgressEvent`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ApiOnPlaybackProgress) GetEventNameOk() (*ProgressEvent, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ApiOnPlaybackProgress) SetEventName(v ProgressEvent)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *ApiOnPlaybackProgress) HasEventName() bool`

HasEventName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


