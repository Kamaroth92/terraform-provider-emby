# PlayerStateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PositionTicks** | Pointer to **NullableInt64** |  | [optional] 
**CanSeek** | Pointer to **bool** |  | [optional] 
**IsPaused** | Pointer to **bool** |  | [optional] 
**IsMuted** | Pointer to **bool** |  | [optional] 
**VolumeLevel** | Pointer to **NullableInt32** |  | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**MediaSourceId** | Pointer to **string** |  | [optional] 
**MediaSource** | Pointer to [**MediaSourceInfo**](MediaSourceInfo.md) |  | [optional] 
**PlayMethod** | Pointer to [**PlayMethod**](PlayMethod.md) |  | [optional] 
**RepeatMode** | Pointer to [**RepeatMode**](RepeatMode.md) |  | [optional] 
**SleepTimerMode** | Pointer to [**SleepTimerMode**](SleepTimerMode.md) |  | [optional] 
**SleepTimerEndTime** | Pointer to **NullableTime** |  | [optional] 
**SubtitleOffset** | Pointer to **int32** |  | [optional] 
**Shuffle** | Pointer to **bool** |  | [optional] 
**PlaybackRate** | Pointer to **float64** |  | [optional] 

## Methods

### NewPlayerStateInfo

`func NewPlayerStateInfo() *PlayerStateInfo`

NewPlayerStateInfo instantiates a new PlayerStateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerStateInfoWithDefaults

`func NewPlayerStateInfoWithDefaults() *PlayerStateInfo`

NewPlayerStateInfoWithDefaults instantiates a new PlayerStateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionTicks

`func (o *PlayerStateInfo) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *PlayerStateInfo) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *PlayerStateInfo) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *PlayerStateInfo) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *PlayerStateInfo) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *PlayerStateInfo) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetCanSeek

`func (o *PlayerStateInfo) GetCanSeek() bool`

GetCanSeek returns the CanSeek field if non-nil, zero value otherwise.

### GetCanSeekOk

`func (o *PlayerStateInfo) GetCanSeekOk() (*bool, bool)`

GetCanSeekOk returns a tuple with the CanSeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeek

`func (o *PlayerStateInfo) SetCanSeek(v bool)`

SetCanSeek sets CanSeek field to given value.

### HasCanSeek

`func (o *PlayerStateInfo) HasCanSeek() bool`

HasCanSeek returns a boolean if a field has been set.

### GetIsPaused

`func (o *PlayerStateInfo) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *PlayerStateInfo) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *PlayerStateInfo) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *PlayerStateInfo) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### GetIsMuted

`func (o *PlayerStateInfo) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *PlayerStateInfo) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *PlayerStateInfo) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *PlayerStateInfo) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetVolumeLevel

`func (o *PlayerStateInfo) GetVolumeLevel() int32`

GetVolumeLevel returns the VolumeLevel field if non-nil, zero value otherwise.

### GetVolumeLevelOk

`func (o *PlayerStateInfo) GetVolumeLevelOk() (*int32, bool)`

GetVolumeLevelOk returns a tuple with the VolumeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLevel

`func (o *PlayerStateInfo) SetVolumeLevel(v int32)`

SetVolumeLevel sets VolumeLevel field to given value.

### HasVolumeLevel

`func (o *PlayerStateInfo) HasVolumeLevel() bool`

HasVolumeLevel returns a boolean if a field has been set.

### SetVolumeLevelNil

`func (o *PlayerStateInfo) SetVolumeLevelNil(b bool)`

 SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil

### UnsetVolumeLevel
`func (o *PlayerStateInfo) UnsetVolumeLevel()`

UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
### GetAudioStreamIndex

`func (o *PlayerStateInfo) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *PlayerStateInfo) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *PlayerStateInfo) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *PlayerStateInfo) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *PlayerStateInfo) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *PlayerStateInfo) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *PlayerStateInfo) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *PlayerStateInfo) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *PlayerStateInfo) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *PlayerStateInfo) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *PlayerStateInfo) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *PlayerStateInfo) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetMediaSourceId

`func (o *PlayerStateInfo) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *PlayerStateInfo) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *PlayerStateInfo) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *PlayerStateInfo) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### GetMediaSource

`func (o *PlayerStateInfo) GetMediaSource() MediaSourceInfo`

GetMediaSource returns the MediaSource field if non-nil, zero value otherwise.

### GetMediaSourceOk

`func (o *PlayerStateInfo) GetMediaSourceOk() (*MediaSourceInfo, bool)`

GetMediaSourceOk returns a tuple with the MediaSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSource

`func (o *PlayerStateInfo) SetMediaSource(v MediaSourceInfo)`

SetMediaSource sets MediaSource field to given value.

### HasMediaSource

`func (o *PlayerStateInfo) HasMediaSource() bool`

HasMediaSource returns a boolean if a field has been set.

### GetPlayMethod

`func (o *PlayerStateInfo) GetPlayMethod() PlayMethod`

GetPlayMethod returns the PlayMethod field if non-nil, zero value otherwise.

### GetPlayMethodOk

`func (o *PlayerStateInfo) GetPlayMethodOk() (*PlayMethod, bool)`

GetPlayMethodOk returns a tuple with the PlayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayMethod

`func (o *PlayerStateInfo) SetPlayMethod(v PlayMethod)`

SetPlayMethod sets PlayMethod field to given value.

### HasPlayMethod

`func (o *PlayerStateInfo) HasPlayMethod() bool`

HasPlayMethod returns a boolean if a field has been set.

### GetRepeatMode

`func (o *PlayerStateInfo) GetRepeatMode() RepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *PlayerStateInfo) GetRepeatModeOk() (*RepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *PlayerStateInfo) SetRepeatMode(v RepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *PlayerStateInfo) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.

### GetSleepTimerMode

`func (o *PlayerStateInfo) GetSleepTimerMode() SleepTimerMode`

GetSleepTimerMode returns the SleepTimerMode field if non-nil, zero value otherwise.

### GetSleepTimerModeOk

`func (o *PlayerStateInfo) GetSleepTimerModeOk() (*SleepTimerMode, bool)`

GetSleepTimerModeOk returns a tuple with the SleepTimerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepTimerMode

`func (o *PlayerStateInfo) SetSleepTimerMode(v SleepTimerMode)`

SetSleepTimerMode sets SleepTimerMode field to given value.

### HasSleepTimerMode

`func (o *PlayerStateInfo) HasSleepTimerMode() bool`

HasSleepTimerMode returns a boolean if a field has been set.

### GetSleepTimerEndTime

`func (o *PlayerStateInfo) GetSleepTimerEndTime() time.Time`

GetSleepTimerEndTime returns the SleepTimerEndTime field if non-nil, zero value otherwise.

### GetSleepTimerEndTimeOk

`func (o *PlayerStateInfo) GetSleepTimerEndTimeOk() (*time.Time, bool)`

GetSleepTimerEndTimeOk returns a tuple with the SleepTimerEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepTimerEndTime

`func (o *PlayerStateInfo) SetSleepTimerEndTime(v time.Time)`

SetSleepTimerEndTime sets SleepTimerEndTime field to given value.

### HasSleepTimerEndTime

`func (o *PlayerStateInfo) HasSleepTimerEndTime() bool`

HasSleepTimerEndTime returns a boolean if a field has been set.

### SetSleepTimerEndTimeNil

`func (o *PlayerStateInfo) SetSleepTimerEndTimeNil(b bool)`

 SetSleepTimerEndTimeNil sets the value for SleepTimerEndTime to be an explicit nil

### UnsetSleepTimerEndTime
`func (o *PlayerStateInfo) UnsetSleepTimerEndTime()`

UnsetSleepTimerEndTime ensures that no value is present for SleepTimerEndTime, not even an explicit nil
### GetSubtitleOffset

`func (o *PlayerStateInfo) GetSubtitleOffset() int32`

GetSubtitleOffset returns the SubtitleOffset field if non-nil, zero value otherwise.

### GetSubtitleOffsetOk

`func (o *PlayerStateInfo) GetSubtitleOffsetOk() (*int32, bool)`

GetSubtitleOffsetOk returns a tuple with the SubtitleOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleOffset

`func (o *PlayerStateInfo) SetSubtitleOffset(v int32)`

SetSubtitleOffset sets SubtitleOffset field to given value.

### HasSubtitleOffset

`func (o *PlayerStateInfo) HasSubtitleOffset() bool`

HasSubtitleOffset returns a boolean if a field has been set.

### GetShuffle

`func (o *PlayerStateInfo) GetShuffle() bool`

GetShuffle returns the Shuffle field if non-nil, zero value otherwise.

### GetShuffleOk

`func (o *PlayerStateInfo) GetShuffleOk() (*bool, bool)`

GetShuffleOk returns a tuple with the Shuffle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShuffle

`func (o *PlayerStateInfo) SetShuffle(v bool)`

SetShuffle sets Shuffle field to given value.

### HasShuffle

`func (o *PlayerStateInfo) HasShuffle() bool`

HasShuffle returns a boolean if a field has been set.

### GetPlaybackRate

`func (o *PlayerStateInfo) GetPlaybackRate() float64`

GetPlaybackRate returns the PlaybackRate field if non-nil, zero value otherwise.

### GetPlaybackRateOk

`func (o *PlayerStateInfo) GetPlaybackRateOk() (*float64, bool)`

GetPlaybackRateOk returns a tuple with the PlaybackRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackRate

`func (o *PlayerStateInfo) SetPlaybackRate(v float64)`

SetPlaybackRate sets PlaybackRate field to given value.

### HasPlaybackRate

`func (o *PlayerStateInfo) HasPlaybackRate() bool`

HasPlaybackRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


