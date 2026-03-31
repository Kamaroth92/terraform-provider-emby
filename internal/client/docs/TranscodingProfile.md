# TranscodingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DlnaProfileType**](DlnaProfileType.md) |  | [optional] 
**VideoCodec** | Pointer to **string** |  | [optional] 
**AudioCodec** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**EstimateContentLength** | Pointer to **bool** |  | [optional] 
**EnableMpegtsM2TsMode** | Pointer to **bool** |  | [optional] 
**TranscodeSeekInfo** | Pointer to [**TranscodeSeekInfo**](TranscodeSeekInfo.md) |  | [optional] 
**CopyTimestamps** | Pointer to **bool** |  | [optional] 
**Context** | Pointer to [**EncodingContext**](EncodingContext.md) |  | [optional] 
**MaxAudioChannels** | Pointer to **string** |  | [optional] 
**MinSegments** | Pointer to **int32** |  | [optional] 
**SegmentLength** | Pointer to **int32** |  | [optional] 
**BreakOnNonKeyFrames** | Pointer to **bool** |  | [optional] 
**AllowInterlacedVideoStreamCopy** | Pointer to **bool** |  | [optional] 
**ManifestSubtitles** | Pointer to **string** |  | [optional] 
**MaxManifestSubtitles** | Pointer to **int32** |  | [optional] 
**MaxWidth** | Pointer to **int32** |  | [optional] 
**MaxHeight** | Pointer to **int32** |  | [optional] 
**FillEmptySubtitleSegments** | Pointer to **bool** |  | [optional] 

## Methods

### NewTranscodingProfile

`func NewTranscodingProfile() *TranscodingProfile`

NewTranscodingProfile instantiates a new TranscodingProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscodingProfileWithDefaults

`func NewTranscodingProfileWithDefaults() *TranscodingProfile`

NewTranscodingProfileWithDefaults instantiates a new TranscodingProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *TranscodingProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *TranscodingProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *TranscodingProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *TranscodingProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetType

`func (o *TranscodingProfile) GetType() DlnaProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TranscodingProfile) GetTypeOk() (*DlnaProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TranscodingProfile) SetType(v DlnaProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *TranscodingProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVideoCodec

`func (o *TranscodingProfile) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *TranscodingProfile) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *TranscodingProfile) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *TranscodingProfile) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetAudioCodec

`func (o *TranscodingProfile) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *TranscodingProfile) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *TranscodingProfile) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *TranscodingProfile) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetProtocol

`func (o *TranscodingProfile) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *TranscodingProfile) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *TranscodingProfile) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *TranscodingProfile) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetEstimateContentLength

`func (o *TranscodingProfile) GetEstimateContentLength() bool`

GetEstimateContentLength returns the EstimateContentLength field if non-nil, zero value otherwise.

### GetEstimateContentLengthOk

`func (o *TranscodingProfile) GetEstimateContentLengthOk() (*bool, bool)`

GetEstimateContentLengthOk returns a tuple with the EstimateContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateContentLength

`func (o *TranscodingProfile) SetEstimateContentLength(v bool)`

SetEstimateContentLength sets EstimateContentLength field to given value.

### HasEstimateContentLength

`func (o *TranscodingProfile) HasEstimateContentLength() bool`

HasEstimateContentLength returns a boolean if a field has been set.

### GetEnableMpegtsM2TsMode

`func (o *TranscodingProfile) GetEnableMpegtsM2TsMode() bool`

GetEnableMpegtsM2TsMode returns the EnableMpegtsM2TsMode field if non-nil, zero value otherwise.

### GetEnableMpegtsM2TsModeOk

`func (o *TranscodingProfile) GetEnableMpegtsM2TsModeOk() (*bool, bool)`

GetEnableMpegtsM2TsModeOk returns a tuple with the EnableMpegtsM2TsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMpegtsM2TsMode

`func (o *TranscodingProfile) SetEnableMpegtsM2TsMode(v bool)`

SetEnableMpegtsM2TsMode sets EnableMpegtsM2TsMode field to given value.

### HasEnableMpegtsM2TsMode

`func (o *TranscodingProfile) HasEnableMpegtsM2TsMode() bool`

HasEnableMpegtsM2TsMode returns a boolean if a field has been set.

### GetTranscodeSeekInfo

`func (o *TranscodingProfile) GetTranscodeSeekInfo() TranscodeSeekInfo`

GetTranscodeSeekInfo returns the TranscodeSeekInfo field if non-nil, zero value otherwise.

### GetTranscodeSeekInfoOk

`func (o *TranscodingProfile) GetTranscodeSeekInfoOk() (*TranscodeSeekInfo, bool)`

GetTranscodeSeekInfoOk returns a tuple with the TranscodeSeekInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodeSeekInfo

`func (o *TranscodingProfile) SetTranscodeSeekInfo(v TranscodeSeekInfo)`

SetTranscodeSeekInfo sets TranscodeSeekInfo field to given value.

### HasTranscodeSeekInfo

`func (o *TranscodingProfile) HasTranscodeSeekInfo() bool`

HasTranscodeSeekInfo returns a boolean if a field has been set.

### GetCopyTimestamps

`func (o *TranscodingProfile) GetCopyTimestamps() bool`

GetCopyTimestamps returns the CopyTimestamps field if non-nil, zero value otherwise.

### GetCopyTimestampsOk

`func (o *TranscodingProfile) GetCopyTimestampsOk() (*bool, bool)`

GetCopyTimestampsOk returns a tuple with the CopyTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTimestamps

`func (o *TranscodingProfile) SetCopyTimestamps(v bool)`

SetCopyTimestamps sets CopyTimestamps field to given value.

### HasCopyTimestamps

`func (o *TranscodingProfile) HasCopyTimestamps() bool`

HasCopyTimestamps returns a boolean if a field has been set.

### GetContext

`func (o *TranscodingProfile) GetContext() EncodingContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TranscodingProfile) GetContextOk() (*EncodingContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TranscodingProfile) SetContext(v EncodingContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TranscodingProfile) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMaxAudioChannels

`func (o *TranscodingProfile) GetMaxAudioChannels() string`

GetMaxAudioChannels returns the MaxAudioChannels field if non-nil, zero value otherwise.

### GetMaxAudioChannelsOk

`func (o *TranscodingProfile) GetMaxAudioChannelsOk() (*string, bool)`

GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudioChannels

`func (o *TranscodingProfile) SetMaxAudioChannels(v string)`

SetMaxAudioChannels sets MaxAudioChannels field to given value.

### HasMaxAudioChannels

`func (o *TranscodingProfile) HasMaxAudioChannels() bool`

HasMaxAudioChannels returns a boolean if a field has been set.

### GetMinSegments

`func (o *TranscodingProfile) GetMinSegments() int32`

GetMinSegments returns the MinSegments field if non-nil, zero value otherwise.

### GetMinSegmentsOk

`func (o *TranscodingProfile) GetMinSegmentsOk() (*int32, bool)`

GetMinSegmentsOk returns a tuple with the MinSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSegments

`func (o *TranscodingProfile) SetMinSegments(v int32)`

SetMinSegments sets MinSegments field to given value.

### HasMinSegments

`func (o *TranscodingProfile) HasMinSegments() bool`

HasMinSegments returns a boolean if a field has been set.

### GetSegmentLength

`func (o *TranscodingProfile) GetSegmentLength() int32`

GetSegmentLength returns the SegmentLength field if non-nil, zero value otherwise.

### GetSegmentLengthOk

`func (o *TranscodingProfile) GetSegmentLengthOk() (*int32, bool)`

GetSegmentLengthOk returns a tuple with the SegmentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentLength

`func (o *TranscodingProfile) SetSegmentLength(v int32)`

SetSegmentLength sets SegmentLength field to given value.

### HasSegmentLength

`func (o *TranscodingProfile) HasSegmentLength() bool`

HasSegmentLength returns a boolean if a field has been set.

### GetBreakOnNonKeyFrames

`func (o *TranscodingProfile) GetBreakOnNonKeyFrames() bool`

GetBreakOnNonKeyFrames returns the BreakOnNonKeyFrames field if non-nil, zero value otherwise.

### GetBreakOnNonKeyFramesOk

`func (o *TranscodingProfile) GetBreakOnNonKeyFramesOk() (*bool, bool)`

GetBreakOnNonKeyFramesOk returns a tuple with the BreakOnNonKeyFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakOnNonKeyFrames

`func (o *TranscodingProfile) SetBreakOnNonKeyFrames(v bool)`

SetBreakOnNonKeyFrames sets BreakOnNonKeyFrames field to given value.

### HasBreakOnNonKeyFrames

`func (o *TranscodingProfile) HasBreakOnNonKeyFrames() bool`

HasBreakOnNonKeyFrames returns a boolean if a field has been set.

### GetAllowInterlacedVideoStreamCopy

`func (o *TranscodingProfile) GetAllowInterlacedVideoStreamCopy() bool`

GetAllowInterlacedVideoStreamCopy returns the AllowInterlacedVideoStreamCopy field if non-nil, zero value otherwise.

### GetAllowInterlacedVideoStreamCopyOk

`func (o *TranscodingProfile) GetAllowInterlacedVideoStreamCopyOk() (*bool, bool)`

GetAllowInterlacedVideoStreamCopyOk returns a tuple with the AllowInterlacedVideoStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInterlacedVideoStreamCopy

`func (o *TranscodingProfile) SetAllowInterlacedVideoStreamCopy(v bool)`

SetAllowInterlacedVideoStreamCopy sets AllowInterlacedVideoStreamCopy field to given value.

### HasAllowInterlacedVideoStreamCopy

`func (o *TranscodingProfile) HasAllowInterlacedVideoStreamCopy() bool`

HasAllowInterlacedVideoStreamCopy returns a boolean if a field has been set.

### GetManifestSubtitles

`func (o *TranscodingProfile) GetManifestSubtitles() string`

GetManifestSubtitles returns the ManifestSubtitles field if non-nil, zero value otherwise.

### GetManifestSubtitlesOk

`func (o *TranscodingProfile) GetManifestSubtitlesOk() (*string, bool)`

GetManifestSubtitlesOk returns a tuple with the ManifestSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestSubtitles

`func (o *TranscodingProfile) SetManifestSubtitles(v string)`

SetManifestSubtitles sets ManifestSubtitles field to given value.

### HasManifestSubtitles

`func (o *TranscodingProfile) HasManifestSubtitles() bool`

HasManifestSubtitles returns a boolean if a field has been set.

### GetMaxManifestSubtitles

`func (o *TranscodingProfile) GetMaxManifestSubtitles() int32`

GetMaxManifestSubtitles returns the MaxManifestSubtitles field if non-nil, zero value otherwise.

### GetMaxManifestSubtitlesOk

`func (o *TranscodingProfile) GetMaxManifestSubtitlesOk() (*int32, bool)`

GetMaxManifestSubtitlesOk returns a tuple with the MaxManifestSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxManifestSubtitles

`func (o *TranscodingProfile) SetMaxManifestSubtitles(v int32)`

SetMaxManifestSubtitles sets MaxManifestSubtitles field to given value.

### HasMaxManifestSubtitles

`func (o *TranscodingProfile) HasMaxManifestSubtitles() bool`

HasMaxManifestSubtitles returns a boolean if a field has been set.

### GetMaxWidth

`func (o *TranscodingProfile) GetMaxWidth() int32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *TranscodingProfile) GetMaxWidthOk() (*int32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *TranscodingProfile) SetMaxWidth(v int32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *TranscodingProfile) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### GetMaxHeight

`func (o *TranscodingProfile) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *TranscodingProfile) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *TranscodingProfile) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *TranscodingProfile) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### GetFillEmptySubtitleSegments

`func (o *TranscodingProfile) GetFillEmptySubtitleSegments() bool`

GetFillEmptySubtitleSegments returns the FillEmptySubtitleSegments field if non-nil, zero value otherwise.

### GetFillEmptySubtitleSegmentsOk

`func (o *TranscodingProfile) GetFillEmptySubtitleSegmentsOk() (*bool, bool)`

GetFillEmptySubtitleSegmentsOk returns a tuple with the FillEmptySubtitleSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillEmptySubtitleSegments

`func (o *TranscodingProfile) SetFillEmptySubtitleSegments(v bool)`

SetFillEmptySubtitleSegments sets FillEmptySubtitleSegments field to given value.

### HasFillEmptySubtitleSegments

`func (o *TranscodingProfile) HasFillEmptySubtitleSegments() bool`

HasFillEmptySubtitleSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


