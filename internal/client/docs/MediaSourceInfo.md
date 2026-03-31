# MediaSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chapters** | Pointer to [**[]ChapterInfo**](ChapterInfo.md) |  | [optional] 
**Protocol** | Pointer to [**MediaProtocol**](MediaProtocol.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**EncoderPath** | Pointer to **string** |  | [optional] 
**EncoderProtocol** | Pointer to [**MediaProtocol**](MediaProtocol.md) |  | [optional] 
**Type** | Pointer to [**MediaSourceType**](MediaSourceType.md) |  | [optional] 
**ProbePath** | Pointer to **string** |  | [optional] 
**ProbeProtocol** | Pointer to [**MediaProtocol**](MediaProtocol.md) |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SortName** | Pointer to **string** |  | [optional] 
**IsRemote** | Pointer to **bool** |  | [optional] 
**HasMixedProtocols** | Pointer to **bool** |  | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**ContainerStartTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**SupportsTranscoding** | Pointer to **bool** |  | [optional] 
**TrancodeLiveStartIndex** | Pointer to **NullableInt32** |  | [optional] 
**WallClockStart** | Pointer to **NullableTime** |  | [optional] 
**SupportsDirectStream** | Pointer to **bool** |  | [optional] 
**SupportsDirectPlay** | Pointer to **bool** |  | [optional] 
**IsInfiniteStream** | Pointer to **bool** |  | [optional] 
**RequiresOpening** | Pointer to **bool** |  | [optional] 
**OpenToken** | Pointer to **string** |  | [optional] 
**RequiresClosing** | Pointer to **bool** |  | [optional] 
**LiveStreamId** | Pointer to **string** |  | [optional] 
**BufferMs** | Pointer to **NullableInt32** |  | [optional] 
**RequiresLooping** | Pointer to **bool** |  | [optional] 
**SupportsProbing** | Pointer to **bool** |  | [optional] 
**Video3DFormat** | Pointer to [**Video3DFormat**](Video3DFormat.md) |  | [optional] 
**MediaStreams** | Pointer to [**[]MediaStream**](MediaStream.md) |  | [optional] 
**Formats** | Pointer to **[]string** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**Timestamp** | Pointer to [**TransportStreamTimestamp**](TransportStreamTimestamp.md) |  | [optional] 
**RequiredHttpHeaders** | Pointer to **map[string]string** |  | [optional] 
**DirectStreamUrl** | Pointer to **string** |  | [optional] 
**AddApiKeyToDirectStreamUrl** | Pointer to **bool** |  | [optional] 
**TranscodingUrl** | Pointer to **string** |  | [optional] 
**TranscodingSubProtocol** | Pointer to **string** |  | [optional] 
**TranscodingContainer** | Pointer to **string** |  | [optional] 
**AnalyzeDurationMs** | Pointer to **NullableInt32** |  | [optional] 
**ReadAtNativeFramerate** | Pointer to **bool** |  | [optional] 
**DefaultAudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**DefaultSubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 

## Methods

### NewMediaSourceInfo

`func NewMediaSourceInfo() *MediaSourceInfo`

NewMediaSourceInfo instantiates a new MediaSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaSourceInfoWithDefaults

`func NewMediaSourceInfoWithDefaults() *MediaSourceInfo`

NewMediaSourceInfoWithDefaults instantiates a new MediaSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChapters

`func (o *MediaSourceInfo) GetChapters() []ChapterInfo`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *MediaSourceInfo) GetChaptersOk() (*[]ChapterInfo, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *MediaSourceInfo) SetChapters(v []ChapterInfo)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *MediaSourceInfo) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### GetProtocol

`func (o *MediaSourceInfo) GetProtocol() MediaProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *MediaSourceInfo) GetProtocolOk() (*MediaProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *MediaSourceInfo) SetProtocol(v MediaProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *MediaSourceInfo) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetId

`func (o *MediaSourceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaSourceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaSourceInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MediaSourceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *MediaSourceInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MediaSourceInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MediaSourceInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MediaSourceInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetEncoderPath

`func (o *MediaSourceInfo) GetEncoderPath() string`

GetEncoderPath returns the EncoderPath field if non-nil, zero value otherwise.

### GetEncoderPathOk

`func (o *MediaSourceInfo) GetEncoderPathOk() (*string, bool)`

GetEncoderPathOk returns a tuple with the EncoderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderPath

`func (o *MediaSourceInfo) SetEncoderPath(v string)`

SetEncoderPath sets EncoderPath field to given value.

### HasEncoderPath

`func (o *MediaSourceInfo) HasEncoderPath() bool`

HasEncoderPath returns a boolean if a field has been set.

### GetEncoderProtocol

`func (o *MediaSourceInfo) GetEncoderProtocol() MediaProtocol`

GetEncoderProtocol returns the EncoderProtocol field if non-nil, zero value otherwise.

### GetEncoderProtocolOk

`func (o *MediaSourceInfo) GetEncoderProtocolOk() (*MediaProtocol, bool)`

GetEncoderProtocolOk returns a tuple with the EncoderProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderProtocol

`func (o *MediaSourceInfo) SetEncoderProtocol(v MediaProtocol)`

SetEncoderProtocol sets EncoderProtocol field to given value.

### HasEncoderProtocol

`func (o *MediaSourceInfo) HasEncoderProtocol() bool`

HasEncoderProtocol returns a boolean if a field has been set.

### GetType

`func (o *MediaSourceInfo) GetType() MediaSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaSourceInfo) GetTypeOk() (*MediaSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaSourceInfo) SetType(v MediaSourceType)`

SetType sets Type field to given value.

### HasType

`func (o *MediaSourceInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProbePath

`func (o *MediaSourceInfo) GetProbePath() string`

GetProbePath returns the ProbePath field if non-nil, zero value otherwise.

### GetProbePathOk

`func (o *MediaSourceInfo) GetProbePathOk() (*string, bool)`

GetProbePathOk returns a tuple with the ProbePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbePath

`func (o *MediaSourceInfo) SetProbePath(v string)`

SetProbePath sets ProbePath field to given value.

### HasProbePath

`func (o *MediaSourceInfo) HasProbePath() bool`

HasProbePath returns a boolean if a field has been set.

### GetProbeProtocol

`func (o *MediaSourceInfo) GetProbeProtocol() MediaProtocol`

GetProbeProtocol returns the ProbeProtocol field if non-nil, zero value otherwise.

### GetProbeProtocolOk

`func (o *MediaSourceInfo) GetProbeProtocolOk() (*MediaProtocol, bool)`

GetProbeProtocolOk returns a tuple with the ProbeProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeProtocol

`func (o *MediaSourceInfo) SetProbeProtocol(v MediaProtocol)`

SetProbeProtocol sets ProbeProtocol field to given value.

### HasProbeProtocol

`func (o *MediaSourceInfo) HasProbeProtocol() bool`

HasProbeProtocol returns a boolean if a field has been set.

### GetContainer

`func (o *MediaSourceInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *MediaSourceInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *MediaSourceInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *MediaSourceInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetSize

`func (o *MediaSourceInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MediaSourceInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MediaSourceInfo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MediaSourceInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MediaSourceInfo) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MediaSourceInfo) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetName

`func (o *MediaSourceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MediaSourceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MediaSourceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MediaSourceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSortName

`func (o *MediaSourceInfo) GetSortName() string`

GetSortName returns the SortName field if non-nil, zero value otherwise.

### GetSortNameOk

`func (o *MediaSourceInfo) GetSortNameOk() (*string, bool)`

GetSortNameOk returns a tuple with the SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortName

`func (o *MediaSourceInfo) SetSortName(v string)`

SetSortName sets SortName field to given value.

### HasSortName

`func (o *MediaSourceInfo) HasSortName() bool`

HasSortName returns a boolean if a field has been set.

### GetIsRemote

`func (o *MediaSourceInfo) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *MediaSourceInfo) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *MediaSourceInfo) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *MediaSourceInfo) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetHasMixedProtocols

`func (o *MediaSourceInfo) GetHasMixedProtocols() bool`

GetHasMixedProtocols returns the HasMixedProtocols field if non-nil, zero value otherwise.

### GetHasMixedProtocolsOk

`func (o *MediaSourceInfo) GetHasMixedProtocolsOk() (*bool, bool)`

GetHasMixedProtocolsOk returns a tuple with the HasMixedProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMixedProtocols

`func (o *MediaSourceInfo) SetHasMixedProtocols(v bool)`

SetHasMixedProtocols sets HasMixedProtocols field to given value.

### HasHasMixedProtocols

`func (o *MediaSourceInfo) HasHasMixedProtocols() bool`

HasHasMixedProtocols returns a boolean if a field has been set.

### GetRunTimeTicks

`func (o *MediaSourceInfo) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *MediaSourceInfo) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *MediaSourceInfo) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *MediaSourceInfo) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *MediaSourceInfo) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *MediaSourceInfo) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetContainerStartTimeTicks

`func (o *MediaSourceInfo) GetContainerStartTimeTicks() int64`

GetContainerStartTimeTicks returns the ContainerStartTimeTicks field if non-nil, zero value otherwise.

### GetContainerStartTimeTicksOk

`func (o *MediaSourceInfo) GetContainerStartTimeTicksOk() (*int64, bool)`

GetContainerStartTimeTicksOk returns a tuple with the ContainerStartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerStartTimeTicks

`func (o *MediaSourceInfo) SetContainerStartTimeTicks(v int64)`

SetContainerStartTimeTicks sets ContainerStartTimeTicks field to given value.

### HasContainerStartTimeTicks

`func (o *MediaSourceInfo) HasContainerStartTimeTicks() bool`

HasContainerStartTimeTicks returns a boolean if a field has been set.

### SetContainerStartTimeTicksNil

`func (o *MediaSourceInfo) SetContainerStartTimeTicksNil(b bool)`

 SetContainerStartTimeTicksNil sets the value for ContainerStartTimeTicks to be an explicit nil

### UnsetContainerStartTimeTicks
`func (o *MediaSourceInfo) UnsetContainerStartTimeTicks()`

UnsetContainerStartTimeTicks ensures that no value is present for ContainerStartTimeTicks, not even an explicit nil
### GetSupportsTranscoding

`func (o *MediaSourceInfo) GetSupportsTranscoding() bool`

GetSupportsTranscoding returns the SupportsTranscoding field if non-nil, zero value otherwise.

### GetSupportsTranscodingOk

`func (o *MediaSourceInfo) GetSupportsTranscodingOk() (*bool, bool)`

GetSupportsTranscodingOk returns a tuple with the SupportsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsTranscoding

`func (o *MediaSourceInfo) SetSupportsTranscoding(v bool)`

SetSupportsTranscoding sets SupportsTranscoding field to given value.

### HasSupportsTranscoding

`func (o *MediaSourceInfo) HasSupportsTranscoding() bool`

HasSupportsTranscoding returns a boolean if a field has been set.

### GetTrancodeLiveStartIndex

`func (o *MediaSourceInfo) GetTrancodeLiveStartIndex() int32`

GetTrancodeLiveStartIndex returns the TrancodeLiveStartIndex field if non-nil, zero value otherwise.

### GetTrancodeLiveStartIndexOk

`func (o *MediaSourceInfo) GetTrancodeLiveStartIndexOk() (*int32, bool)`

GetTrancodeLiveStartIndexOk returns a tuple with the TrancodeLiveStartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrancodeLiveStartIndex

`func (o *MediaSourceInfo) SetTrancodeLiveStartIndex(v int32)`

SetTrancodeLiveStartIndex sets TrancodeLiveStartIndex field to given value.

### HasTrancodeLiveStartIndex

`func (o *MediaSourceInfo) HasTrancodeLiveStartIndex() bool`

HasTrancodeLiveStartIndex returns a boolean if a field has been set.

### SetTrancodeLiveStartIndexNil

`func (o *MediaSourceInfo) SetTrancodeLiveStartIndexNil(b bool)`

 SetTrancodeLiveStartIndexNil sets the value for TrancodeLiveStartIndex to be an explicit nil

### UnsetTrancodeLiveStartIndex
`func (o *MediaSourceInfo) UnsetTrancodeLiveStartIndex()`

UnsetTrancodeLiveStartIndex ensures that no value is present for TrancodeLiveStartIndex, not even an explicit nil
### GetWallClockStart

`func (o *MediaSourceInfo) GetWallClockStart() time.Time`

GetWallClockStart returns the WallClockStart field if non-nil, zero value otherwise.

### GetWallClockStartOk

`func (o *MediaSourceInfo) GetWallClockStartOk() (*time.Time, bool)`

GetWallClockStartOk returns a tuple with the WallClockStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallClockStart

`func (o *MediaSourceInfo) SetWallClockStart(v time.Time)`

SetWallClockStart sets WallClockStart field to given value.

### HasWallClockStart

`func (o *MediaSourceInfo) HasWallClockStart() bool`

HasWallClockStart returns a boolean if a field has been set.

### SetWallClockStartNil

`func (o *MediaSourceInfo) SetWallClockStartNil(b bool)`

 SetWallClockStartNil sets the value for WallClockStart to be an explicit nil

### UnsetWallClockStart
`func (o *MediaSourceInfo) UnsetWallClockStart()`

UnsetWallClockStart ensures that no value is present for WallClockStart, not even an explicit nil
### GetSupportsDirectStream

`func (o *MediaSourceInfo) GetSupportsDirectStream() bool`

GetSupportsDirectStream returns the SupportsDirectStream field if non-nil, zero value otherwise.

### GetSupportsDirectStreamOk

`func (o *MediaSourceInfo) GetSupportsDirectStreamOk() (*bool, bool)`

GetSupportsDirectStreamOk returns a tuple with the SupportsDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDirectStream

`func (o *MediaSourceInfo) SetSupportsDirectStream(v bool)`

SetSupportsDirectStream sets SupportsDirectStream field to given value.

### HasSupportsDirectStream

`func (o *MediaSourceInfo) HasSupportsDirectStream() bool`

HasSupportsDirectStream returns a boolean if a field has been set.

### GetSupportsDirectPlay

`func (o *MediaSourceInfo) GetSupportsDirectPlay() bool`

GetSupportsDirectPlay returns the SupportsDirectPlay field if non-nil, zero value otherwise.

### GetSupportsDirectPlayOk

`func (o *MediaSourceInfo) GetSupportsDirectPlayOk() (*bool, bool)`

GetSupportsDirectPlayOk returns a tuple with the SupportsDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDirectPlay

`func (o *MediaSourceInfo) SetSupportsDirectPlay(v bool)`

SetSupportsDirectPlay sets SupportsDirectPlay field to given value.

### HasSupportsDirectPlay

`func (o *MediaSourceInfo) HasSupportsDirectPlay() bool`

HasSupportsDirectPlay returns a boolean if a field has been set.

### GetIsInfiniteStream

`func (o *MediaSourceInfo) GetIsInfiniteStream() bool`

GetIsInfiniteStream returns the IsInfiniteStream field if non-nil, zero value otherwise.

### GetIsInfiniteStreamOk

`func (o *MediaSourceInfo) GetIsInfiniteStreamOk() (*bool, bool)`

GetIsInfiniteStreamOk returns a tuple with the IsInfiniteStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInfiniteStream

`func (o *MediaSourceInfo) SetIsInfiniteStream(v bool)`

SetIsInfiniteStream sets IsInfiniteStream field to given value.

### HasIsInfiniteStream

`func (o *MediaSourceInfo) HasIsInfiniteStream() bool`

HasIsInfiniteStream returns a boolean if a field has been set.

### GetRequiresOpening

`func (o *MediaSourceInfo) GetRequiresOpening() bool`

GetRequiresOpening returns the RequiresOpening field if non-nil, zero value otherwise.

### GetRequiresOpeningOk

`func (o *MediaSourceInfo) GetRequiresOpeningOk() (*bool, bool)`

GetRequiresOpeningOk returns a tuple with the RequiresOpening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresOpening

`func (o *MediaSourceInfo) SetRequiresOpening(v bool)`

SetRequiresOpening sets RequiresOpening field to given value.

### HasRequiresOpening

`func (o *MediaSourceInfo) HasRequiresOpening() bool`

HasRequiresOpening returns a boolean if a field has been set.

### GetOpenToken

`func (o *MediaSourceInfo) GetOpenToken() string`

GetOpenToken returns the OpenToken field if non-nil, zero value otherwise.

### GetOpenTokenOk

`func (o *MediaSourceInfo) GetOpenTokenOk() (*string, bool)`

GetOpenTokenOk returns a tuple with the OpenToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenToken

`func (o *MediaSourceInfo) SetOpenToken(v string)`

SetOpenToken sets OpenToken field to given value.

### HasOpenToken

`func (o *MediaSourceInfo) HasOpenToken() bool`

HasOpenToken returns a boolean if a field has been set.

### GetRequiresClosing

`func (o *MediaSourceInfo) GetRequiresClosing() bool`

GetRequiresClosing returns the RequiresClosing field if non-nil, zero value otherwise.

### GetRequiresClosingOk

`func (o *MediaSourceInfo) GetRequiresClosingOk() (*bool, bool)`

GetRequiresClosingOk returns a tuple with the RequiresClosing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresClosing

`func (o *MediaSourceInfo) SetRequiresClosing(v bool)`

SetRequiresClosing sets RequiresClosing field to given value.

### HasRequiresClosing

`func (o *MediaSourceInfo) HasRequiresClosing() bool`

HasRequiresClosing returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *MediaSourceInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *MediaSourceInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *MediaSourceInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *MediaSourceInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### GetBufferMs

`func (o *MediaSourceInfo) GetBufferMs() int32`

GetBufferMs returns the BufferMs field if non-nil, zero value otherwise.

### GetBufferMsOk

`func (o *MediaSourceInfo) GetBufferMsOk() (*int32, bool)`

GetBufferMsOk returns a tuple with the BufferMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferMs

`func (o *MediaSourceInfo) SetBufferMs(v int32)`

SetBufferMs sets BufferMs field to given value.

### HasBufferMs

`func (o *MediaSourceInfo) HasBufferMs() bool`

HasBufferMs returns a boolean if a field has been set.

### SetBufferMsNil

`func (o *MediaSourceInfo) SetBufferMsNil(b bool)`

 SetBufferMsNil sets the value for BufferMs to be an explicit nil

### UnsetBufferMs
`func (o *MediaSourceInfo) UnsetBufferMs()`

UnsetBufferMs ensures that no value is present for BufferMs, not even an explicit nil
### GetRequiresLooping

`func (o *MediaSourceInfo) GetRequiresLooping() bool`

GetRequiresLooping returns the RequiresLooping field if non-nil, zero value otherwise.

### GetRequiresLoopingOk

`func (o *MediaSourceInfo) GetRequiresLoopingOk() (*bool, bool)`

GetRequiresLoopingOk returns a tuple with the RequiresLooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresLooping

`func (o *MediaSourceInfo) SetRequiresLooping(v bool)`

SetRequiresLooping sets RequiresLooping field to given value.

### HasRequiresLooping

`func (o *MediaSourceInfo) HasRequiresLooping() bool`

HasRequiresLooping returns a boolean if a field has been set.

### GetSupportsProbing

`func (o *MediaSourceInfo) GetSupportsProbing() bool`

GetSupportsProbing returns the SupportsProbing field if non-nil, zero value otherwise.

### GetSupportsProbingOk

`func (o *MediaSourceInfo) GetSupportsProbingOk() (*bool, bool)`

GetSupportsProbingOk returns a tuple with the SupportsProbing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsProbing

`func (o *MediaSourceInfo) SetSupportsProbing(v bool)`

SetSupportsProbing sets SupportsProbing field to given value.

### HasSupportsProbing

`func (o *MediaSourceInfo) HasSupportsProbing() bool`

HasSupportsProbing returns a boolean if a field has been set.

### GetVideo3DFormat

`func (o *MediaSourceInfo) GetVideo3DFormat() Video3DFormat`

GetVideo3DFormat returns the Video3DFormat field if non-nil, zero value otherwise.

### GetVideo3DFormatOk

`func (o *MediaSourceInfo) GetVideo3DFormatOk() (*Video3DFormat, bool)`

GetVideo3DFormatOk returns a tuple with the Video3DFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo3DFormat

`func (o *MediaSourceInfo) SetVideo3DFormat(v Video3DFormat)`

SetVideo3DFormat sets Video3DFormat field to given value.

### HasVideo3DFormat

`func (o *MediaSourceInfo) HasVideo3DFormat() bool`

HasVideo3DFormat returns a boolean if a field has been set.

### GetMediaStreams

`func (o *MediaSourceInfo) GetMediaStreams() []MediaStream`

GetMediaStreams returns the MediaStreams field if non-nil, zero value otherwise.

### GetMediaStreamsOk

`func (o *MediaSourceInfo) GetMediaStreamsOk() (*[]MediaStream, bool)`

GetMediaStreamsOk returns a tuple with the MediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStreams

`func (o *MediaSourceInfo) SetMediaStreams(v []MediaStream)`

SetMediaStreams sets MediaStreams field to given value.

### HasMediaStreams

`func (o *MediaSourceInfo) HasMediaStreams() bool`

HasMediaStreams returns a boolean if a field has been set.

### GetFormats

`func (o *MediaSourceInfo) GetFormats() []string`

GetFormats returns the Formats field if non-nil, zero value otherwise.

### GetFormatsOk

`func (o *MediaSourceInfo) GetFormatsOk() (*[]string, bool)`

GetFormatsOk returns a tuple with the Formats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormats

`func (o *MediaSourceInfo) SetFormats(v []string)`

SetFormats sets Formats field to given value.

### HasFormats

`func (o *MediaSourceInfo) HasFormats() bool`

HasFormats returns a boolean if a field has been set.

### GetBitrate

`func (o *MediaSourceInfo) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *MediaSourceInfo) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *MediaSourceInfo) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *MediaSourceInfo) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *MediaSourceInfo) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *MediaSourceInfo) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetTimestamp

`func (o *MediaSourceInfo) GetTimestamp() TransportStreamTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MediaSourceInfo) GetTimestampOk() (*TransportStreamTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MediaSourceInfo) SetTimestamp(v TransportStreamTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MediaSourceInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRequiredHttpHeaders

`func (o *MediaSourceInfo) GetRequiredHttpHeaders() map[string]string`

GetRequiredHttpHeaders returns the RequiredHttpHeaders field if non-nil, zero value otherwise.

### GetRequiredHttpHeadersOk

`func (o *MediaSourceInfo) GetRequiredHttpHeadersOk() (*map[string]string, bool)`

GetRequiredHttpHeadersOk returns a tuple with the RequiredHttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredHttpHeaders

`func (o *MediaSourceInfo) SetRequiredHttpHeaders(v map[string]string)`

SetRequiredHttpHeaders sets RequiredHttpHeaders field to given value.

### HasRequiredHttpHeaders

`func (o *MediaSourceInfo) HasRequiredHttpHeaders() bool`

HasRequiredHttpHeaders returns a boolean if a field has been set.

### GetDirectStreamUrl

`func (o *MediaSourceInfo) GetDirectStreamUrl() string`

GetDirectStreamUrl returns the DirectStreamUrl field if non-nil, zero value otherwise.

### GetDirectStreamUrlOk

`func (o *MediaSourceInfo) GetDirectStreamUrlOk() (*string, bool)`

GetDirectStreamUrlOk returns a tuple with the DirectStreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectStreamUrl

`func (o *MediaSourceInfo) SetDirectStreamUrl(v string)`

SetDirectStreamUrl sets DirectStreamUrl field to given value.

### HasDirectStreamUrl

`func (o *MediaSourceInfo) HasDirectStreamUrl() bool`

HasDirectStreamUrl returns a boolean if a field has been set.

### GetAddApiKeyToDirectStreamUrl

`func (o *MediaSourceInfo) GetAddApiKeyToDirectStreamUrl() bool`

GetAddApiKeyToDirectStreamUrl returns the AddApiKeyToDirectStreamUrl field if non-nil, zero value otherwise.

### GetAddApiKeyToDirectStreamUrlOk

`func (o *MediaSourceInfo) GetAddApiKeyToDirectStreamUrlOk() (*bool, bool)`

GetAddApiKeyToDirectStreamUrlOk returns a tuple with the AddApiKeyToDirectStreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddApiKeyToDirectStreamUrl

`func (o *MediaSourceInfo) SetAddApiKeyToDirectStreamUrl(v bool)`

SetAddApiKeyToDirectStreamUrl sets AddApiKeyToDirectStreamUrl field to given value.

### HasAddApiKeyToDirectStreamUrl

`func (o *MediaSourceInfo) HasAddApiKeyToDirectStreamUrl() bool`

HasAddApiKeyToDirectStreamUrl returns a boolean if a field has been set.

### GetTranscodingUrl

`func (o *MediaSourceInfo) GetTranscodingUrl() string`

GetTranscodingUrl returns the TranscodingUrl field if non-nil, zero value otherwise.

### GetTranscodingUrlOk

`func (o *MediaSourceInfo) GetTranscodingUrlOk() (*string, bool)`

GetTranscodingUrlOk returns a tuple with the TranscodingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingUrl

`func (o *MediaSourceInfo) SetTranscodingUrl(v string)`

SetTranscodingUrl sets TranscodingUrl field to given value.

### HasTranscodingUrl

`func (o *MediaSourceInfo) HasTranscodingUrl() bool`

HasTranscodingUrl returns a boolean if a field has been set.

### GetTranscodingSubProtocol

`func (o *MediaSourceInfo) GetTranscodingSubProtocol() string`

GetTranscodingSubProtocol returns the TranscodingSubProtocol field if non-nil, zero value otherwise.

### GetTranscodingSubProtocolOk

`func (o *MediaSourceInfo) GetTranscodingSubProtocolOk() (*string, bool)`

GetTranscodingSubProtocolOk returns a tuple with the TranscodingSubProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingSubProtocol

`func (o *MediaSourceInfo) SetTranscodingSubProtocol(v string)`

SetTranscodingSubProtocol sets TranscodingSubProtocol field to given value.

### HasTranscodingSubProtocol

`func (o *MediaSourceInfo) HasTranscodingSubProtocol() bool`

HasTranscodingSubProtocol returns a boolean if a field has been set.

### GetTranscodingContainer

`func (o *MediaSourceInfo) GetTranscodingContainer() string`

GetTranscodingContainer returns the TranscodingContainer field if non-nil, zero value otherwise.

### GetTranscodingContainerOk

`func (o *MediaSourceInfo) GetTranscodingContainerOk() (*string, bool)`

GetTranscodingContainerOk returns a tuple with the TranscodingContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingContainer

`func (o *MediaSourceInfo) SetTranscodingContainer(v string)`

SetTranscodingContainer sets TranscodingContainer field to given value.

### HasTranscodingContainer

`func (o *MediaSourceInfo) HasTranscodingContainer() bool`

HasTranscodingContainer returns a boolean if a field has been set.

### GetAnalyzeDurationMs

`func (o *MediaSourceInfo) GetAnalyzeDurationMs() int32`

GetAnalyzeDurationMs returns the AnalyzeDurationMs field if non-nil, zero value otherwise.

### GetAnalyzeDurationMsOk

`func (o *MediaSourceInfo) GetAnalyzeDurationMsOk() (*int32, bool)`

GetAnalyzeDurationMsOk returns a tuple with the AnalyzeDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzeDurationMs

`func (o *MediaSourceInfo) SetAnalyzeDurationMs(v int32)`

SetAnalyzeDurationMs sets AnalyzeDurationMs field to given value.

### HasAnalyzeDurationMs

`func (o *MediaSourceInfo) HasAnalyzeDurationMs() bool`

HasAnalyzeDurationMs returns a boolean if a field has been set.

### SetAnalyzeDurationMsNil

`func (o *MediaSourceInfo) SetAnalyzeDurationMsNil(b bool)`

 SetAnalyzeDurationMsNil sets the value for AnalyzeDurationMs to be an explicit nil

### UnsetAnalyzeDurationMs
`func (o *MediaSourceInfo) UnsetAnalyzeDurationMs()`

UnsetAnalyzeDurationMs ensures that no value is present for AnalyzeDurationMs, not even an explicit nil
### GetReadAtNativeFramerate

`func (o *MediaSourceInfo) GetReadAtNativeFramerate() bool`

GetReadAtNativeFramerate returns the ReadAtNativeFramerate field if non-nil, zero value otherwise.

### GetReadAtNativeFramerateOk

`func (o *MediaSourceInfo) GetReadAtNativeFramerateOk() (*bool, bool)`

GetReadAtNativeFramerateOk returns a tuple with the ReadAtNativeFramerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAtNativeFramerate

`func (o *MediaSourceInfo) SetReadAtNativeFramerate(v bool)`

SetReadAtNativeFramerate sets ReadAtNativeFramerate field to given value.

### HasReadAtNativeFramerate

`func (o *MediaSourceInfo) HasReadAtNativeFramerate() bool`

HasReadAtNativeFramerate returns a boolean if a field has been set.

### GetDefaultAudioStreamIndex

`func (o *MediaSourceInfo) GetDefaultAudioStreamIndex() int32`

GetDefaultAudioStreamIndex returns the DefaultAudioStreamIndex field if non-nil, zero value otherwise.

### GetDefaultAudioStreamIndexOk

`func (o *MediaSourceInfo) GetDefaultAudioStreamIndexOk() (*int32, bool)`

GetDefaultAudioStreamIndexOk returns a tuple with the DefaultAudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAudioStreamIndex

`func (o *MediaSourceInfo) SetDefaultAudioStreamIndex(v int32)`

SetDefaultAudioStreamIndex sets DefaultAudioStreamIndex field to given value.

### HasDefaultAudioStreamIndex

`func (o *MediaSourceInfo) HasDefaultAudioStreamIndex() bool`

HasDefaultAudioStreamIndex returns a boolean if a field has been set.

### SetDefaultAudioStreamIndexNil

`func (o *MediaSourceInfo) SetDefaultAudioStreamIndexNil(b bool)`

 SetDefaultAudioStreamIndexNil sets the value for DefaultAudioStreamIndex to be an explicit nil

### UnsetDefaultAudioStreamIndex
`func (o *MediaSourceInfo) UnsetDefaultAudioStreamIndex()`

UnsetDefaultAudioStreamIndex ensures that no value is present for DefaultAudioStreamIndex, not even an explicit nil
### GetDefaultSubtitleStreamIndex

`func (o *MediaSourceInfo) GetDefaultSubtitleStreamIndex() int32`

GetDefaultSubtitleStreamIndex returns the DefaultSubtitleStreamIndex field if non-nil, zero value otherwise.

### GetDefaultSubtitleStreamIndexOk

`func (o *MediaSourceInfo) GetDefaultSubtitleStreamIndexOk() (*int32, bool)`

GetDefaultSubtitleStreamIndexOk returns a tuple with the DefaultSubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubtitleStreamIndex

`func (o *MediaSourceInfo) SetDefaultSubtitleStreamIndex(v int32)`

SetDefaultSubtitleStreamIndex sets DefaultSubtitleStreamIndex field to given value.

### HasDefaultSubtitleStreamIndex

`func (o *MediaSourceInfo) HasDefaultSubtitleStreamIndex() bool`

HasDefaultSubtitleStreamIndex returns a boolean if a field has been set.

### SetDefaultSubtitleStreamIndexNil

`func (o *MediaSourceInfo) SetDefaultSubtitleStreamIndexNil(b bool)`

 SetDefaultSubtitleStreamIndexNil sets the value for DefaultSubtitleStreamIndex to be an explicit nil

### UnsetDefaultSubtitleStreamIndex
`func (o *MediaSourceInfo) UnsetDefaultSubtitleStreamIndex()`

UnsetDefaultSubtitleStreamIndex ensures that no value is present for DefaultSubtitleStreamIndex, not even an explicit nil
### GetItemId

`func (o *MediaSourceInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *MediaSourceInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *MediaSourceInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *MediaSourceInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetServerId

`func (o *MediaSourceInfo) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *MediaSourceInfo) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *MediaSourceInfo) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *MediaSourceInfo) HasServerId() bool`

HasServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


