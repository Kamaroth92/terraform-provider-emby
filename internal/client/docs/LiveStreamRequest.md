# LiveStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenToken** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**PlaySessionId** | Pointer to **string** |  | [optional] 
**MaxStreamingBitrate** | Pointer to **NullableInt64** |  | [optional] 
**StartTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**MaxAudioChannels** | Pointer to **NullableInt32** |  | [optional] 
**ItemId** | Pointer to **int64** |  | [optional] 
**DeviceProfile** | Pointer to [**DeviceProfile**](DeviceProfile.md) |  | [optional] 
**EnableDirectPlay** | Pointer to **bool** |  | [optional] 
**EnableDirectStream** | Pointer to **bool** |  | [optional] 
**EnableTranscoding** | Pointer to **bool** |  | [optional] 
**AllowVideoStreamCopy** | Pointer to **bool** |  | [optional] 
**AllowInterlacedVideoStreamCopy** | Pointer to **bool** |  | [optional] 
**AllowAudioStreamCopy** | Pointer to **bool** |  | [optional] 

## Methods

### NewLiveStreamRequest

`func NewLiveStreamRequest() *LiveStreamRequest`

NewLiveStreamRequest instantiates a new LiveStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamRequestWithDefaults

`func NewLiveStreamRequestWithDefaults() *LiveStreamRequest`

NewLiveStreamRequestWithDefaults instantiates a new LiveStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenToken

`func (o *LiveStreamRequest) GetOpenToken() string`

GetOpenToken returns the OpenToken field if non-nil, zero value otherwise.

### GetOpenTokenOk

`func (o *LiveStreamRequest) GetOpenTokenOk() (*string, bool)`

GetOpenTokenOk returns a tuple with the OpenToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenToken

`func (o *LiveStreamRequest) SetOpenToken(v string)`

SetOpenToken sets OpenToken field to given value.

### HasOpenToken

`func (o *LiveStreamRequest) HasOpenToken() bool`

HasOpenToken returns a boolean if a field has been set.

### GetUserId

`func (o *LiveStreamRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LiveStreamRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LiveStreamRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LiveStreamRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPlaySessionId

`func (o *LiveStreamRequest) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *LiveStreamRequest) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *LiveStreamRequest) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *LiveStreamRequest) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### GetMaxStreamingBitrate

`func (o *LiveStreamRequest) GetMaxStreamingBitrate() int64`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *LiveStreamRequest) GetMaxStreamingBitrateOk() (*int64, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *LiveStreamRequest) SetMaxStreamingBitrate(v int64)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *LiveStreamRequest) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *LiveStreamRequest) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *LiveStreamRequest) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetStartTimeTicks

`func (o *LiveStreamRequest) GetStartTimeTicks() int64`

GetStartTimeTicks returns the StartTimeTicks field if non-nil, zero value otherwise.

### GetStartTimeTicksOk

`func (o *LiveStreamRequest) GetStartTimeTicksOk() (*int64, bool)`

GetStartTimeTicksOk returns a tuple with the StartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeTicks

`func (o *LiveStreamRequest) SetStartTimeTicks(v int64)`

SetStartTimeTicks sets StartTimeTicks field to given value.

### HasStartTimeTicks

`func (o *LiveStreamRequest) HasStartTimeTicks() bool`

HasStartTimeTicks returns a boolean if a field has been set.

### SetStartTimeTicksNil

`func (o *LiveStreamRequest) SetStartTimeTicksNil(b bool)`

 SetStartTimeTicksNil sets the value for StartTimeTicks to be an explicit nil

### UnsetStartTimeTicks
`func (o *LiveStreamRequest) UnsetStartTimeTicks()`

UnsetStartTimeTicks ensures that no value is present for StartTimeTicks, not even an explicit nil
### GetAudioStreamIndex

`func (o *LiveStreamRequest) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *LiveStreamRequest) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *LiveStreamRequest) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *LiveStreamRequest) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *LiveStreamRequest) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *LiveStreamRequest) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *LiveStreamRequest) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *LiveStreamRequest) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *LiveStreamRequest) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *LiveStreamRequest) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *LiveStreamRequest) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *LiveStreamRequest) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetMaxAudioChannels

`func (o *LiveStreamRequest) GetMaxAudioChannels() int32`

GetMaxAudioChannels returns the MaxAudioChannels field if non-nil, zero value otherwise.

### GetMaxAudioChannelsOk

`func (o *LiveStreamRequest) GetMaxAudioChannelsOk() (*int32, bool)`

GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudioChannels

`func (o *LiveStreamRequest) SetMaxAudioChannels(v int32)`

SetMaxAudioChannels sets MaxAudioChannels field to given value.

### HasMaxAudioChannels

`func (o *LiveStreamRequest) HasMaxAudioChannels() bool`

HasMaxAudioChannels returns a boolean if a field has been set.

### SetMaxAudioChannelsNil

`func (o *LiveStreamRequest) SetMaxAudioChannelsNil(b bool)`

 SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil

### UnsetMaxAudioChannels
`func (o *LiveStreamRequest) UnsetMaxAudioChannels()`

UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
### GetItemId

`func (o *LiveStreamRequest) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *LiveStreamRequest) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *LiveStreamRequest) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *LiveStreamRequest) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *LiveStreamRequest) GetDeviceProfile() DeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *LiveStreamRequest) GetDeviceProfileOk() (*DeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *LiveStreamRequest) SetDeviceProfile(v DeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *LiveStreamRequest) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### GetEnableDirectPlay

`func (o *LiveStreamRequest) GetEnableDirectPlay() bool`

GetEnableDirectPlay returns the EnableDirectPlay field if non-nil, zero value otherwise.

### GetEnableDirectPlayOk

`func (o *LiveStreamRequest) GetEnableDirectPlayOk() (*bool, bool)`

GetEnableDirectPlayOk returns a tuple with the EnableDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectPlay

`func (o *LiveStreamRequest) SetEnableDirectPlay(v bool)`

SetEnableDirectPlay sets EnableDirectPlay field to given value.

### HasEnableDirectPlay

`func (o *LiveStreamRequest) HasEnableDirectPlay() bool`

HasEnableDirectPlay returns a boolean if a field has been set.

### GetEnableDirectStream

`func (o *LiveStreamRequest) GetEnableDirectStream() bool`

GetEnableDirectStream returns the EnableDirectStream field if non-nil, zero value otherwise.

### GetEnableDirectStreamOk

`func (o *LiveStreamRequest) GetEnableDirectStreamOk() (*bool, bool)`

GetEnableDirectStreamOk returns a tuple with the EnableDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectStream

`func (o *LiveStreamRequest) SetEnableDirectStream(v bool)`

SetEnableDirectStream sets EnableDirectStream field to given value.

### HasEnableDirectStream

`func (o *LiveStreamRequest) HasEnableDirectStream() bool`

HasEnableDirectStream returns a boolean if a field has been set.

### GetEnableTranscoding

`func (o *LiveStreamRequest) GetEnableTranscoding() bool`

GetEnableTranscoding returns the EnableTranscoding field if non-nil, zero value otherwise.

### GetEnableTranscodingOk

`func (o *LiveStreamRequest) GetEnableTranscodingOk() (*bool, bool)`

GetEnableTranscodingOk returns a tuple with the EnableTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTranscoding

`func (o *LiveStreamRequest) SetEnableTranscoding(v bool)`

SetEnableTranscoding sets EnableTranscoding field to given value.

### HasEnableTranscoding

`func (o *LiveStreamRequest) HasEnableTranscoding() bool`

HasEnableTranscoding returns a boolean if a field has been set.

### GetAllowVideoStreamCopy

`func (o *LiveStreamRequest) GetAllowVideoStreamCopy() bool`

GetAllowVideoStreamCopy returns the AllowVideoStreamCopy field if non-nil, zero value otherwise.

### GetAllowVideoStreamCopyOk

`func (o *LiveStreamRequest) GetAllowVideoStreamCopyOk() (*bool, bool)`

GetAllowVideoStreamCopyOk returns a tuple with the AllowVideoStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVideoStreamCopy

`func (o *LiveStreamRequest) SetAllowVideoStreamCopy(v bool)`

SetAllowVideoStreamCopy sets AllowVideoStreamCopy field to given value.

### HasAllowVideoStreamCopy

`func (o *LiveStreamRequest) HasAllowVideoStreamCopy() bool`

HasAllowVideoStreamCopy returns a boolean if a field has been set.

### GetAllowInterlacedVideoStreamCopy

`func (o *LiveStreamRequest) GetAllowInterlacedVideoStreamCopy() bool`

GetAllowInterlacedVideoStreamCopy returns the AllowInterlacedVideoStreamCopy field if non-nil, zero value otherwise.

### GetAllowInterlacedVideoStreamCopyOk

`func (o *LiveStreamRequest) GetAllowInterlacedVideoStreamCopyOk() (*bool, bool)`

GetAllowInterlacedVideoStreamCopyOk returns a tuple with the AllowInterlacedVideoStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInterlacedVideoStreamCopy

`func (o *LiveStreamRequest) SetAllowInterlacedVideoStreamCopy(v bool)`

SetAllowInterlacedVideoStreamCopy sets AllowInterlacedVideoStreamCopy field to given value.

### HasAllowInterlacedVideoStreamCopy

`func (o *LiveStreamRequest) HasAllowInterlacedVideoStreamCopy() bool`

HasAllowInterlacedVideoStreamCopy returns a boolean if a field has been set.

### GetAllowAudioStreamCopy

`func (o *LiveStreamRequest) GetAllowAudioStreamCopy() bool`

GetAllowAudioStreamCopy returns the AllowAudioStreamCopy field if non-nil, zero value otherwise.

### GetAllowAudioStreamCopyOk

`func (o *LiveStreamRequest) GetAllowAudioStreamCopyOk() (*bool, bool)`

GetAllowAudioStreamCopyOk returns a tuple with the AllowAudioStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAudioStreamCopy

`func (o *LiveStreamRequest) SetAllowAudioStreamCopy(v bool)`

SetAllowAudioStreamCopy sets AllowAudioStreamCopy field to given value.

### HasAllowAudioStreamCopy

`func (o *LiveStreamRequest) HasAllowAudioStreamCopy() bool`

HasAllowAudioStreamCopy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


