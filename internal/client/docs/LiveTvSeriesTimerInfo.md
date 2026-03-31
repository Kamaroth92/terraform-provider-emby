# LiveTvSeriesTimerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**ChannelIds** | Pointer to **[]string** |  | [optional] 
**ParentFolderId** | Pointer to **int64** |  | [optional] 
**ProgramId** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**RecordAnyTime** | Pointer to **bool** |  | [optional] 
**KeepUpTo** | Pointer to **int32** |  | [optional] 
**KeepUntil** | Pointer to [**LiveTvKeepUntil**](LiveTvKeepUntil.md) |  | [optional] 
**SkipEpisodesInLibrary** | Pointer to **bool** |  | [optional] 
**MatchExistingItemsWithAnyLibrary** | Pointer to **bool** |  | [optional] 
**RecordNewOnly** | Pointer to **bool** |  | [optional] 
**Days** | Pointer to [**[]DayOfWeek**](DayOfWeek.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**PrePaddingSeconds** | Pointer to **int32** |  | [optional] 
**PostPaddingSeconds** | Pointer to **int32** |  | [optional] 
**IsPrePaddingRequired** | Pointer to **bool** |  | [optional] 
**IsPostPaddingRequired** | Pointer to **bool** |  | [optional] 
**SeriesId** | Pointer to **string** |  | [optional] 
**ProviderIds** | Pointer to **map[string]string** |  | [optional] 
**MaxRecordingSeconds** | Pointer to **int32** |  | [optional] 
**Keywords** | Pointer to [**[]LiveTvKeywordInfo**](LiveTvKeywordInfo.md) |  | [optional] 
**TimerType** | Pointer to [**LiveTvTimerType**](LiveTvTimerType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewLiveTvSeriesTimerInfo

`func NewLiveTvSeriesTimerInfo() *LiveTvSeriesTimerInfo`

NewLiveTvSeriesTimerInfo instantiates a new LiveTvSeriesTimerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveTvSeriesTimerInfoWithDefaults

`func NewLiveTvSeriesTimerInfoWithDefaults() *LiveTvSeriesTimerInfo`

NewLiveTvSeriesTimerInfoWithDefaults instantiates a new LiveTvSeriesTimerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LiveTvSeriesTimerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveTvSeriesTimerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveTvSeriesTimerInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveTvSeriesTimerInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChannelId

`func (o *LiveTvSeriesTimerInfo) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *LiveTvSeriesTimerInfo) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *LiveTvSeriesTimerInfo) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *LiveTvSeriesTimerInfo) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetChannelIds

`func (o *LiveTvSeriesTimerInfo) GetChannelIds() []string`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *LiveTvSeriesTimerInfo) GetChannelIdsOk() (*[]string, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *LiveTvSeriesTimerInfo) SetChannelIds(v []string)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *LiveTvSeriesTimerInfo) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.

### GetParentFolderId

`func (o *LiveTvSeriesTimerInfo) GetParentFolderId() int64`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *LiveTvSeriesTimerInfo) GetParentFolderIdOk() (*int64, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *LiveTvSeriesTimerInfo) SetParentFolderId(v int64)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *LiveTvSeriesTimerInfo) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetProgramId

`func (o *LiveTvSeriesTimerInfo) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *LiveTvSeriesTimerInfo) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *LiveTvSeriesTimerInfo) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *LiveTvSeriesTimerInfo) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetServiceName

`func (o *LiveTvSeriesTimerInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *LiveTvSeriesTimerInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *LiveTvSeriesTimerInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *LiveTvSeriesTimerInfo) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetOverview

`func (o *LiveTvSeriesTimerInfo) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *LiveTvSeriesTimerInfo) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *LiveTvSeriesTimerInfo) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *LiveTvSeriesTimerInfo) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetStartDate

`func (o *LiveTvSeriesTimerInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LiveTvSeriesTimerInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LiveTvSeriesTimerInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LiveTvSeriesTimerInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *LiveTvSeriesTimerInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *LiveTvSeriesTimerInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *LiveTvSeriesTimerInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *LiveTvSeriesTimerInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRecordAnyTime

`func (o *LiveTvSeriesTimerInfo) GetRecordAnyTime() bool`

GetRecordAnyTime returns the RecordAnyTime field if non-nil, zero value otherwise.

### GetRecordAnyTimeOk

`func (o *LiveTvSeriesTimerInfo) GetRecordAnyTimeOk() (*bool, bool)`

GetRecordAnyTimeOk returns a tuple with the RecordAnyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordAnyTime

`func (o *LiveTvSeriesTimerInfo) SetRecordAnyTime(v bool)`

SetRecordAnyTime sets RecordAnyTime field to given value.

### HasRecordAnyTime

`func (o *LiveTvSeriesTimerInfo) HasRecordAnyTime() bool`

HasRecordAnyTime returns a boolean if a field has been set.

### GetKeepUpTo

`func (o *LiveTvSeriesTimerInfo) GetKeepUpTo() int32`

GetKeepUpTo returns the KeepUpTo field if non-nil, zero value otherwise.

### GetKeepUpToOk

`func (o *LiveTvSeriesTimerInfo) GetKeepUpToOk() (*int32, bool)`

GetKeepUpToOk returns a tuple with the KeepUpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUpTo

`func (o *LiveTvSeriesTimerInfo) SetKeepUpTo(v int32)`

SetKeepUpTo sets KeepUpTo field to given value.

### HasKeepUpTo

`func (o *LiveTvSeriesTimerInfo) HasKeepUpTo() bool`

HasKeepUpTo returns a boolean if a field has been set.

### GetKeepUntil

`func (o *LiveTvSeriesTimerInfo) GetKeepUntil() LiveTvKeepUntil`

GetKeepUntil returns the KeepUntil field if non-nil, zero value otherwise.

### GetKeepUntilOk

`func (o *LiveTvSeriesTimerInfo) GetKeepUntilOk() (*LiveTvKeepUntil, bool)`

GetKeepUntilOk returns a tuple with the KeepUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUntil

`func (o *LiveTvSeriesTimerInfo) SetKeepUntil(v LiveTvKeepUntil)`

SetKeepUntil sets KeepUntil field to given value.

### HasKeepUntil

`func (o *LiveTvSeriesTimerInfo) HasKeepUntil() bool`

HasKeepUntil returns a boolean if a field has been set.

### GetSkipEpisodesInLibrary

`func (o *LiveTvSeriesTimerInfo) GetSkipEpisodesInLibrary() bool`

GetSkipEpisodesInLibrary returns the SkipEpisodesInLibrary field if non-nil, zero value otherwise.

### GetSkipEpisodesInLibraryOk

`func (o *LiveTvSeriesTimerInfo) GetSkipEpisodesInLibraryOk() (*bool, bool)`

GetSkipEpisodesInLibraryOk returns a tuple with the SkipEpisodesInLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEpisodesInLibrary

`func (o *LiveTvSeriesTimerInfo) SetSkipEpisodesInLibrary(v bool)`

SetSkipEpisodesInLibrary sets SkipEpisodesInLibrary field to given value.

### HasSkipEpisodesInLibrary

`func (o *LiveTvSeriesTimerInfo) HasSkipEpisodesInLibrary() bool`

HasSkipEpisodesInLibrary returns a boolean if a field has been set.

### GetMatchExistingItemsWithAnyLibrary

`func (o *LiveTvSeriesTimerInfo) GetMatchExistingItemsWithAnyLibrary() bool`

GetMatchExistingItemsWithAnyLibrary returns the MatchExistingItemsWithAnyLibrary field if non-nil, zero value otherwise.

### GetMatchExistingItemsWithAnyLibraryOk

`func (o *LiveTvSeriesTimerInfo) GetMatchExistingItemsWithAnyLibraryOk() (*bool, bool)`

GetMatchExistingItemsWithAnyLibraryOk returns a tuple with the MatchExistingItemsWithAnyLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExistingItemsWithAnyLibrary

`func (o *LiveTvSeriesTimerInfo) SetMatchExistingItemsWithAnyLibrary(v bool)`

SetMatchExistingItemsWithAnyLibrary sets MatchExistingItemsWithAnyLibrary field to given value.

### HasMatchExistingItemsWithAnyLibrary

`func (o *LiveTvSeriesTimerInfo) HasMatchExistingItemsWithAnyLibrary() bool`

HasMatchExistingItemsWithAnyLibrary returns a boolean if a field has been set.

### GetRecordNewOnly

`func (o *LiveTvSeriesTimerInfo) GetRecordNewOnly() bool`

GetRecordNewOnly returns the RecordNewOnly field if non-nil, zero value otherwise.

### GetRecordNewOnlyOk

`func (o *LiveTvSeriesTimerInfo) GetRecordNewOnlyOk() (*bool, bool)`

GetRecordNewOnlyOk returns a tuple with the RecordNewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordNewOnly

`func (o *LiveTvSeriesTimerInfo) SetRecordNewOnly(v bool)`

SetRecordNewOnly sets RecordNewOnly field to given value.

### HasRecordNewOnly

`func (o *LiveTvSeriesTimerInfo) HasRecordNewOnly() bool`

HasRecordNewOnly returns a boolean if a field has been set.

### GetDays

`func (o *LiveTvSeriesTimerInfo) GetDays() []DayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *LiveTvSeriesTimerInfo) GetDaysOk() (*[]DayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *LiveTvSeriesTimerInfo) SetDays(v []DayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *LiveTvSeriesTimerInfo) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetPriority

`func (o *LiveTvSeriesTimerInfo) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *LiveTvSeriesTimerInfo) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *LiveTvSeriesTimerInfo) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *LiveTvSeriesTimerInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrePaddingSeconds

`func (o *LiveTvSeriesTimerInfo) GetPrePaddingSeconds() int32`

GetPrePaddingSeconds returns the PrePaddingSeconds field if non-nil, zero value otherwise.

### GetPrePaddingSecondsOk

`func (o *LiveTvSeriesTimerInfo) GetPrePaddingSecondsOk() (*int32, bool)`

GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaddingSeconds

`func (o *LiveTvSeriesTimerInfo) SetPrePaddingSeconds(v int32)`

SetPrePaddingSeconds sets PrePaddingSeconds field to given value.

### HasPrePaddingSeconds

`func (o *LiveTvSeriesTimerInfo) HasPrePaddingSeconds() bool`

HasPrePaddingSeconds returns a boolean if a field has been set.

### GetPostPaddingSeconds

`func (o *LiveTvSeriesTimerInfo) GetPostPaddingSeconds() int32`

GetPostPaddingSeconds returns the PostPaddingSeconds field if non-nil, zero value otherwise.

### GetPostPaddingSecondsOk

`func (o *LiveTvSeriesTimerInfo) GetPostPaddingSecondsOk() (*int32, bool)`

GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaddingSeconds

`func (o *LiveTvSeriesTimerInfo) SetPostPaddingSeconds(v int32)`

SetPostPaddingSeconds sets PostPaddingSeconds field to given value.

### HasPostPaddingSeconds

`func (o *LiveTvSeriesTimerInfo) HasPostPaddingSeconds() bool`

HasPostPaddingSeconds returns a boolean if a field has been set.

### GetIsPrePaddingRequired

`func (o *LiveTvSeriesTimerInfo) GetIsPrePaddingRequired() bool`

GetIsPrePaddingRequired returns the IsPrePaddingRequired field if non-nil, zero value otherwise.

### GetIsPrePaddingRequiredOk

`func (o *LiveTvSeriesTimerInfo) GetIsPrePaddingRequiredOk() (*bool, bool)`

GetIsPrePaddingRequiredOk returns a tuple with the IsPrePaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrePaddingRequired

`func (o *LiveTvSeriesTimerInfo) SetIsPrePaddingRequired(v bool)`

SetIsPrePaddingRequired sets IsPrePaddingRequired field to given value.

### HasIsPrePaddingRequired

`func (o *LiveTvSeriesTimerInfo) HasIsPrePaddingRequired() bool`

HasIsPrePaddingRequired returns a boolean if a field has been set.

### GetIsPostPaddingRequired

`func (o *LiveTvSeriesTimerInfo) GetIsPostPaddingRequired() bool`

GetIsPostPaddingRequired returns the IsPostPaddingRequired field if non-nil, zero value otherwise.

### GetIsPostPaddingRequiredOk

`func (o *LiveTvSeriesTimerInfo) GetIsPostPaddingRequiredOk() (*bool, bool)`

GetIsPostPaddingRequiredOk returns a tuple with the IsPostPaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPostPaddingRequired

`func (o *LiveTvSeriesTimerInfo) SetIsPostPaddingRequired(v bool)`

SetIsPostPaddingRequired sets IsPostPaddingRequired field to given value.

### HasIsPostPaddingRequired

`func (o *LiveTvSeriesTimerInfo) HasIsPostPaddingRequired() bool`

HasIsPostPaddingRequired returns a boolean if a field has been set.

### GetSeriesId

`func (o *LiveTvSeriesTimerInfo) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *LiveTvSeriesTimerInfo) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *LiveTvSeriesTimerInfo) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *LiveTvSeriesTimerInfo) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetProviderIds

`func (o *LiveTvSeriesTimerInfo) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *LiveTvSeriesTimerInfo) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *LiveTvSeriesTimerInfo) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *LiveTvSeriesTimerInfo) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### GetMaxRecordingSeconds

`func (o *LiveTvSeriesTimerInfo) GetMaxRecordingSeconds() int32`

GetMaxRecordingSeconds returns the MaxRecordingSeconds field if non-nil, zero value otherwise.

### GetMaxRecordingSecondsOk

`func (o *LiveTvSeriesTimerInfo) GetMaxRecordingSecondsOk() (*int32, bool)`

GetMaxRecordingSecondsOk returns a tuple with the MaxRecordingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecordingSeconds

`func (o *LiveTvSeriesTimerInfo) SetMaxRecordingSeconds(v int32)`

SetMaxRecordingSeconds sets MaxRecordingSeconds field to given value.

### HasMaxRecordingSeconds

`func (o *LiveTvSeriesTimerInfo) HasMaxRecordingSeconds() bool`

HasMaxRecordingSeconds returns a boolean if a field has been set.

### GetKeywords

`func (o *LiveTvSeriesTimerInfo) GetKeywords() []LiveTvKeywordInfo`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *LiveTvSeriesTimerInfo) GetKeywordsOk() (*[]LiveTvKeywordInfo, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *LiveTvSeriesTimerInfo) SetKeywords(v []LiveTvKeywordInfo)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *LiveTvSeriesTimerInfo) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetTimerType

`func (o *LiveTvSeriesTimerInfo) GetTimerType() LiveTvTimerType`

GetTimerType returns the TimerType field if non-nil, zero value otherwise.

### GetTimerTypeOk

`func (o *LiveTvSeriesTimerInfo) GetTimerTypeOk() (*LiveTvTimerType, bool)`

GetTimerTypeOk returns a tuple with the TimerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerType

`func (o *LiveTvSeriesTimerInfo) SetTimerType(v LiveTvTimerType)`

SetTimerType sets TimerType field to given value.

### HasTimerType

`func (o *LiveTvSeriesTimerInfo) HasTimerType() bool`

HasTimerType returns a boolean if a field has been set.

### GetName

`func (o *LiveTvSeriesTimerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LiveTvSeriesTimerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LiveTvSeriesTimerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LiveTvSeriesTimerInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


