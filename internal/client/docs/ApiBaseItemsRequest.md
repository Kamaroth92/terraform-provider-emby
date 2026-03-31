# ApiBaseItemsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpecialEpisode** | Pointer to **NullableBool** |  | [optional] 
**Is4K** | Pointer to **NullableBool** |  | [optional] 
**MinDateCreated** | Pointer to **NullableTime** |  | [optional] 
**MaxDateCreated** | Pointer to **NullableTime** |  | [optional] 
**EnableTotalRecordCount** | Pointer to **bool** |  | [optional] 
**MatchAnyWord** | Pointer to **bool** |  | [optional] 
**IsDuplicate** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RecordingKeyword** | Pointer to **string** |  | [optional] 
**RecordingKeywordType** | Pointer to [**LiveTvKeywordType**](LiveTvKeywordType.md) |  | [optional] 
**RandomSeed** | Pointer to **int32** |  | [optional] 
**GenreIds** | Pointer to **string** |  | [optional] 
**CollectionIds** | Pointer to **string** |  | [optional] 
**TagIds** | Pointer to **string** |  | [optional] 
**ExcludeTagIds** | Pointer to **string** |  | [optional] 
**ItemPersonTypes** | Pointer to [**[]PersonType**](PersonType.md) |  | [optional] 
**ExcludeArtistIds** | Pointer to **string** |  | [optional] 
**AlbumArtistIds** | Pointer to **string** |  | [optional] 
**ComposerArtistIds** | Pointer to **string** |  | [optional] 
**ContributingArtistIds** | Pointer to **string** |  | [optional] 
**AlbumIds** | Pointer to **string** |  | [optional] 
**OuterIds** | Pointer to **string** |  | [optional] 
**ListItemIds** | Pointer to **string** |  | [optional] 
**AudioLanguages** | Pointer to **string** |  | [optional] 
**SubtitleLanguages** | Pointer to **string** |  | [optional] 
**CanEditItems** | Pointer to **NullableBool** |  | [optional] 
**GroupItemsInto** | Pointer to [**LibraryItemLinkType**](LibraryItemLinkType.md) |  | [optional] 
**IsStandaloneSpecial** | Pointer to **NullableBool** |  | [optional] 
**MinWidth** | Pointer to **NullableInt32** |  | [optional] 
**MinHeight** | Pointer to **NullableInt32** |  | [optional] 
**MaxWidth** | Pointer to **NullableInt32** |  | [optional] 
**MaxHeight** | Pointer to **NullableInt32** |  | [optional] 
**GroupProgramsBySeries** | Pointer to **bool** |  | [optional] 
**GroupByPresentationUniqueKey** | Pointer to **NullableBool** |  | [optional] 
**AirDays** | Pointer to [**[]DayOfWeek**](DayOfWeek.md) |  | [optional] 
**IsAiring** | Pointer to **NullableBool** |  | [optional] 
**HasAired** | Pointer to **NullableBool** |  | [optional] 
**CollectionTypes** | Pointer to **string** |  | [optional] 
**ExcludeSources** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiBaseItemsRequest

`func NewApiBaseItemsRequest() *ApiBaseItemsRequest`

NewApiBaseItemsRequest instantiates a new ApiBaseItemsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBaseItemsRequestWithDefaults

`func NewApiBaseItemsRequestWithDefaults() *ApiBaseItemsRequest`

NewApiBaseItemsRequestWithDefaults instantiates a new ApiBaseItemsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpecialEpisode

`func (o *ApiBaseItemsRequest) GetIsSpecialEpisode() bool`

GetIsSpecialEpisode returns the IsSpecialEpisode field if non-nil, zero value otherwise.

### GetIsSpecialEpisodeOk

`func (o *ApiBaseItemsRequest) GetIsSpecialEpisodeOk() (*bool, bool)`

GetIsSpecialEpisodeOk returns a tuple with the IsSpecialEpisode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialEpisode

`func (o *ApiBaseItemsRequest) SetIsSpecialEpisode(v bool)`

SetIsSpecialEpisode sets IsSpecialEpisode field to given value.

### HasIsSpecialEpisode

`func (o *ApiBaseItemsRequest) HasIsSpecialEpisode() bool`

HasIsSpecialEpisode returns a boolean if a field has been set.

### SetIsSpecialEpisodeNil

`func (o *ApiBaseItemsRequest) SetIsSpecialEpisodeNil(b bool)`

 SetIsSpecialEpisodeNil sets the value for IsSpecialEpisode to be an explicit nil

### UnsetIsSpecialEpisode
`func (o *ApiBaseItemsRequest) UnsetIsSpecialEpisode()`

UnsetIsSpecialEpisode ensures that no value is present for IsSpecialEpisode, not even an explicit nil
### GetIs4K

`func (o *ApiBaseItemsRequest) GetIs4K() bool`

GetIs4K returns the Is4K field if non-nil, zero value otherwise.

### GetIs4KOk

`func (o *ApiBaseItemsRequest) GetIs4KOk() (*bool, bool)`

GetIs4KOk returns a tuple with the Is4K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs4K

`func (o *ApiBaseItemsRequest) SetIs4K(v bool)`

SetIs4K sets Is4K field to given value.

### HasIs4K

`func (o *ApiBaseItemsRequest) HasIs4K() bool`

HasIs4K returns a boolean if a field has been set.

### SetIs4KNil

`func (o *ApiBaseItemsRequest) SetIs4KNil(b bool)`

 SetIs4KNil sets the value for Is4K to be an explicit nil

### UnsetIs4K
`func (o *ApiBaseItemsRequest) UnsetIs4K()`

UnsetIs4K ensures that no value is present for Is4K, not even an explicit nil
### GetMinDateCreated

`func (o *ApiBaseItemsRequest) GetMinDateCreated() time.Time`

GetMinDateCreated returns the MinDateCreated field if non-nil, zero value otherwise.

### GetMinDateCreatedOk

`func (o *ApiBaseItemsRequest) GetMinDateCreatedOk() (*time.Time, bool)`

GetMinDateCreatedOk returns a tuple with the MinDateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDateCreated

`func (o *ApiBaseItemsRequest) SetMinDateCreated(v time.Time)`

SetMinDateCreated sets MinDateCreated field to given value.

### HasMinDateCreated

`func (o *ApiBaseItemsRequest) HasMinDateCreated() bool`

HasMinDateCreated returns a boolean if a field has been set.

### SetMinDateCreatedNil

`func (o *ApiBaseItemsRequest) SetMinDateCreatedNil(b bool)`

 SetMinDateCreatedNil sets the value for MinDateCreated to be an explicit nil

### UnsetMinDateCreated
`func (o *ApiBaseItemsRequest) UnsetMinDateCreated()`

UnsetMinDateCreated ensures that no value is present for MinDateCreated, not even an explicit nil
### GetMaxDateCreated

`func (o *ApiBaseItemsRequest) GetMaxDateCreated() time.Time`

GetMaxDateCreated returns the MaxDateCreated field if non-nil, zero value otherwise.

### GetMaxDateCreatedOk

`func (o *ApiBaseItemsRequest) GetMaxDateCreatedOk() (*time.Time, bool)`

GetMaxDateCreatedOk returns a tuple with the MaxDateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDateCreated

`func (o *ApiBaseItemsRequest) SetMaxDateCreated(v time.Time)`

SetMaxDateCreated sets MaxDateCreated field to given value.

### HasMaxDateCreated

`func (o *ApiBaseItemsRequest) HasMaxDateCreated() bool`

HasMaxDateCreated returns a boolean if a field has been set.

### SetMaxDateCreatedNil

`func (o *ApiBaseItemsRequest) SetMaxDateCreatedNil(b bool)`

 SetMaxDateCreatedNil sets the value for MaxDateCreated to be an explicit nil

### UnsetMaxDateCreated
`func (o *ApiBaseItemsRequest) UnsetMaxDateCreated()`

UnsetMaxDateCreated ensures that no value is present for MaxDateCreated, not even an explicit nil
### GetEnableTotalRecordCount

`func (o *ApiBaseItemsRequest) GetEnableTotalRecordCount() bool`

GetEnableTotalRecordCount returns the EnableTotalRecordCount field if non-nil, zero value otherwise.

### GetEnableTotalRecordCountOk

`func (o *ApiBaseItemsRequest) GetEnableTotalRecordCountOk() (*bool, bool)`

GetEnableTotalRecordCountOk returns a tuple with the EnableTotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTotalRecordCount

`func (o *ApiBaseItemsRequest) SetEnableTotalRecordCount(v bool)`

SetEnableTotalRecordCount sets EnableTotalRecordCount field to given value.

### HasEnableTotalRecordCount

`func (o *ApiBaseItemsRequest) HasEnableTotalRecordCount() bool`

HasEnableTotalRecordCount returns a boolean if a field has been set.

### GetMatchAnyWord

`func (o *ApiBaseItemsRequest) GetMatchAnyWord() bool`

GetMatchAnyWord returns the MatchAnyWord field if non-nil, zero value otherwise.

### GetMatchAnyWordOk

`func (o *ApiBaseItemsRequest) GetMatchAnyWordOk() (*bool, bool)`

GetMatchAnyWordOk returns a tuple with the MatchAnyWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAnyWord

`func (o *ApiBaseItemsRequest) SetMatchAnyWord(v bool)`

SetMatchAnyWord sets MatchAnyWord field to given value.

### HasMatchAnyWord

`func (o *ApiBaseItemsRequest) HasMatchAnyWord() bool`

HasMatchAnyWord returns a boolean if a field has been set.

### GetIsDuplicate

`func (o *ApiBaseItemsRequest) GetIsDuplicate() bool`

GetIsDuplicate returns the IsDuplicate field if non-nil, zero value otherwise.

### GetIsDuplicateOk

`func (o *ApiBaseItemsRequest) GetIsDuplicateOk() (*bool, bool)`

GetIsDuplicateOk returns a tuple with the IsDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDuplicate

`func (o *ApiBaseItemsRequest) SetIsDuplicate(v bool)`

SetIsDuplicate sets IsDuplicate field to given value.

### HasIsDuplicate

`func (o *ApiBaseItemsRequest) HasIsDuplicate() bool`

HasIsDuplicate returns a boolean if a field has been set.

### SetIsDuplicateNil

`func (o *ApiBaseItemsRequest) SetIsDuplicateNil(b bool)`

 SetIsDuplicateNil sets the value for IsDuplicate to be an explicit nil

### UnsetIsDuplicate
`func (o *ApiBaseItemsRequest) UnsetIsDuplicate()`

UnsetIsDuplicate ensures that no value is present for IsDuplicate, not even an explicit nil
### GetName

`func (o *ApiBaseItemsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiBaseItemsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiBaseItemsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiBaseItemsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecordingKeyword

`func (o *ApiBaseItemsRequest) GetRecordingKeyword() string`

GetRecordingKeyword returns the RecordingKeyword field if non-nil, zero value otherwise.

### GetRecordingKeywordOk

`func (o *ApiBaseItemsRequest) GetRecordingKeywordOk() (*string, bool)`

GetRecordingKeywordOk returns a tuple with the RecordingKeyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingKeyword

`func (o *ApiBaseItemsRequest) SetRecordingKeyword(v string)`

SetRecordingKeyword sets RecordingKeyword field to given value.

### HasRecordingKeyword

`func (o *ApiBaseItemsRequest) HasRecordingKeyword() bool`

HasRecordingKeyword returns a boolean if a field has been set.

### GetRecordingKeywordType

`func (o *ApiBaseItemsRequest) GetRecordingKeywordType() LiveTvKeywordType`

GetRecordingKeywordType returns the RecordingKeywordType field if non-nil, zero value otherwise.

### GetRecordingKeywordTypeOk

`func (o *ApiBaseItemsRequest) GetRecordingKeywordTypeOk() (*LiveTvKeywordType, bool)`

GetRecordingKeywordTypeOk returns a tuple with the RecordingKeywordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingKeywordType

`func (o *ApiBaseItemsRequest) SetRecordingKeywordType(v LiveTvKeywordType)`

SetRecordingKeywordType sets RecordingKeywordType field to given value.

### HasRecordingKeywordType

`func (o *ApiBaseItemsRequest) HasRecordingKeywordType() bool`

HasRecordingKeywordType returns a boolean if a field has been set.

### GetRandomSeed

`func (o *ApiBaseItemsRequest) GetRandomSeed() int32`

GetRandomSeed returns the RandomSeed field if non-nil, zero value otherwise.

### GetRandomSeedOk

`func (o *ApiBaseItemsRequest) GetRandomSeedOk() (*int32, bool)`

GetRandomSeedOk returns a tuple with the RandomSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomSeed

`func (o *ApiBaseItemsRequest) SetRandomSeed(v int32)`

SetRandomSeed sets RandomSeed field to given value.

### HasRandomSeed

`func (o *ApiBaseItemsRequest) HasRandomSeed() bool`

HasRandomSeed returns a boolean if a field has been set.

### GetGenreIds

`func (o *ApiBaseItemsRequest) GetGenreIds() string`

GetGenreIds returns the GenreIds field if non-nil, zero value otherwise.

### GetGenreIdsOk

`func (o *ApiBaseItemsRequest) GetGenreIdsOk() (*string, bool)`

GetGenreIdsOk returns a tuple with the GenreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreIds

`func (o *ApiBaseItemsRequest) SetGenreIds(v string)`

SetGenreIds sets GenreIds field to given value.

### HasGenreIds

`func (o *ApiBaseItemsRequest) HasGenreIds() bool`

HasGenreIds returns a boolean if a field has been set.

### GetCollectionIds

`func (o *ApiBaseItemsRequest) GetCollectionIds() string`

GetCollectionIds returns the CollectionIds field if non-nil, zero value otherwise.

### GetCollectionIdsOk

`func (o *ApiBaseItemsRequest) GetCollectionIdsOk() (*string, bool)`

GetCollectionIdsOk returns a tuple with the CollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionIds

`func (o *ApiBaseItemsRequest) SetCollectionIds(v string)`

SetCollectionIds sets CollectionIds field to given value.

### HasCollectionIds

`func (o *ApiBaseItemsRequest) HasCollectionIds() bool`

HasCollectionIds returns a boolean if a field has been set.

### GetTagIds

`func (o *ApiBaseItemsRequest) GetTagIds() string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ApiBaseItemsRequest) GetTagIdsOk() (*string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ApiBaseItemsRequest) SetTagIds(v string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ApiBaseItemsRequest) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetExcludeTagIds

`func (o *ApiBaseItemsRequest) GetExcludeTagIds() string`

GetExcludeTagIds returns the ExcludeTagIds field if non-nil, zero value otherwise.

### GetExcludeTagIdsOk

`func (o *ApiBaseItemsRequest) GetExcludeTagIdsOk() (*string, bool)`

GetExcludeTagIdsOk returns a tuple with the ExcludeTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTagIds

`func (o *ApiBaseItemsRequest) SetExcludeTagIds(v string)`

SetExcludeTagIds sets ExcludeTagIds field to given value.

### HasExcludeTagIds

`func (o *ApiBaseItemsRequest) HasExcludeTagIds() bool`

HasExcludeTagIds returns a boolean if a field has been set.

### GetItemPersonTypes

`func (o *ApiBaseItemsRequest) GetItemPersonTypes() []PersonType`

GetItemPersonTypes returns the ItemPersonTypes field if non-nil, zero value otherwise.

### GetItemPersonTypesOk

`func (o *ApiBaseItemsRequest) GetItemPersonTypesOk() (*[]PersonType, bool)`

GetItemPersonTypesOk returns a tuple with the ItemPersonTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPersonTypes

`func (o *ApiBaseItemsRequest) SetItemPersonTypes(v []PersonType)`

SetItemPersonTypes sets ItemPersonTypes field to given value.

### HasItemPersonTypes

`func (o *ApiBaseItemsRequest) HasItemPersonTypes() bool`

HasItemPersonTypes returns a boolean if a field has been set.

### GetExcludeArtistIds

`func (o *ApiBaseItemsRequest) GetExcludeArtistIds() string`

GetExcludeArtistIds returns the ExcludeArtistIds field if non-nil, zero value otherwise.

### GetExcludeArtistIdsOk

`func (o *ApiBaseItemsRequest) GetExcludeArtistIdsOk() (*string, bool)`

GetExcludeArtistIdsOk returns a tuple with the ExcludeArtistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeArtistIds

`func (o *ApiBaseItemsRequest) SetExcludeArtistIds(v string)`

SetExcludeArtistIds sets ExcludeArtistIds field to given value.

### HasExcludeArtistIds

`func (o *ApiBaseItemsRequest) HasExcludeArtistIds() bool`

HasExcludeArtistIds returns a boolean if a field has been set.

### GetAlbumArtistIds

`func (o *ApiBaseItemsRequest) GetAlbumArtistIds() string`

GetAlbumArtistIds returns the AlbumArtistIds field if non-nil, zero value otherwise.

### GetAlbumArtistIdsOk

`func (o *ApiBaseItemsRequest) GetAlbumArtistIdsOk() (*string, bool)`

GetAlbumArtistIdsOk returns a tuple with the AlbumArtistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtistIds

`func (o *ApiBaseItemsRequest) SetAlbumArtistIds(v string)`

SetAlbumArtistIds sets AlbumArtistIds field to given value.

### HasAlbumArtistIds

`func (o *ApiBaseItemsRequest) HasAlbumArtistIds() bool`

HasAlbumArtistIds returns a boolean if a field has been set.

### GetComposerArtistIds

`func (o *ApiBaseItemsRequest) GetComposerArtistIds() string`

GetComposerArtistIds returns the ComposerArtistIds field if non-nil, zero value otherwise.

### GetComposerArtistIdsOk

`func (o *ApiBaseItemsRequest) GetComposerArtistIdsOk() (*string, bool)`

GetComposerArtistIdsOk returns a tuple with the ComposerArtistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposerArtistIds

`func (o *ApiBaseItemsRequest) SetComposerArtistIds(v string)`

SetComposerArtistIds sets ComposerArtistIds field to given value.

### HasComposerArtistIds

`func (o *ApiBaseItemsRequest) HasComposerArtistIds() bool`

HasComposerArtistIds returns a boolean if a field has been set.

### GetContributingArtistIds

`func (o *ApiBaseItemsRequest) GetContributingArtistIds() string`

GetContributingArtistIds returns the ContributingArtistIds field if non-nil, zero value otherwise.

### GetContributingArtistIdsOk

`func (o *ApiBaseItemsRequest) GetContributingArtistIdsOk() (*string, bool)`

GetContributingArtistIdsOk returns a tuple with the ContributingArtistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributingArtistIds

`func (o *ApiBaseItemsRequest) SetContributingArtistIds(v string)`

SetContributingArtistIds sets ContributingArtistIds field to given value.

### HasContributingArtistIds

`func (o *ApiBaseItemsRequest) HasContributingArtistIds() bool`

HasContributingArtistIds returns a boolean if a field has been set.

### GetAlbumIds

`func (o *ApiBaseItemsRequest) GetAlbumIds() string`

GetAlbumIds returns the AlbumIds field if non-nil, zero value otherwise.

### GetAlbumIdsOk

`func (o *ApiBaseItemsRequest) GetAlbumIdsOk() (*string, bool)`

GetAlbumIdsOk returns a tuple with the AlbumIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumIds

`func (o *ApiBaseItemsRequest) SetAlbumIds(v string)`

SetAlbumIds sets AlbumIds field to given value.

### HasAlbumIds

`func (o *ApiBaseItemsRequest) HasAlbumIds() bool`

HasAlbumIds returns a boolean if a field has been set.

### GetOuterIds

`func (o *ApiBaseItemsRequest) GetOuterIds() string`

GetOuterIds returns the OuterIds field if non-nil, zero value otherwise.

### GetOuterIdsOk

`func (o *ApiBaseItemsRequest) GetOuterIdsOk() (*string, bool)`

GetOuterIdsOk returns a tuple with the OuterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterIds

`func (o *ApiBaseItemsRequest) SetOuterIds(v string)`

SetOuterIds sets OuterIds field to given value.

### HasOuterIds

`func (o *ApiBaseItemsRequest) HasOuterIds() bool`

HasOuterIds returns a boolean if a field has been set.

### GetListItemIds

`func (o *ApiBaseItemsRequest) GetListItemIds() string`

GetListItemIds returns the ListItemIds field if non-nil, zero value otherwise.

### GetListItemIdsOk

`func (o *ApiBaseItemsRequest) GetListItemIdsOk() (*string, bool)`

GetListItemIdsOk returns a tuple with the ListItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListItemIds

`func (o *ApiBaseItemsRequest) SetListItemIds(v string)`

SetListItemIds sets ListItemIds field to given value.

### HasListItemIds

`func (o *ApiBaseItemsRequest) HasListItemIds() bool`

HasListItemIds returns a boolean if a field has been set.

### GetAudioLanguages

`func (o *ApiBaseItemsRequest) GetAudioLanguages() string`

GetAudioLanguages returns the AudioLanguages field if non-nil, zero value otherwise.

### GetAudioLanguagesOk

`func (o *ApiBaseItemsRequest) GetAudioLanguagesOk() (*string, bool)`

GetAudioLanguagesOk returns a tuple with the AudioLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioLanguages

`func (o *ApiBaseItemsRequest) SetAudioLanguages(v string)`

SetAudioLanguages sets AudioLanguages field to given value.

### HasAudioLanguages

`func (o *ApiBaseItemsRequest) HasAudioLanguages() bool`

HasAudioLanguages returns a boolean if a field has been set.

### GetSubtitleLanguages

`func (o *ApiBaseItemsRequest) GetSubtitleLanguages() string`

GetSubtitleLanguages returns the SubtitleLanguages field if non-nil, zero value otherwise.

### GetSubtitleLanguagesOk

`func (o *ApiBaseItemsRequest) GetSubtitleLanguagesOk() (*string, bool)`

GetSubtitleLanguagesOk returns a tuple with the SubtitleLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleLanguages

`func (o *ApiBaseItemsRequest) SetSubtitleLanguages(v string)`

SetSubtitleLanguages sets SubtitleLanguages field to given value.

### HasSubtitleLanguages

`func (o *ApiBaseItemsRequest) HasSubtitleLanguages() bool`

HasSubtitleLanguages returns a boolean if a field has been set.

### GetCanEditItems

`func (o *ApiBaseItemsRequest) GetCanEditItems() bool`

GetCanEditItems returns the CanEditItems field if non-nil, zero value otherwise.

### GetCanEditItemsOk

`func (o *ApiBaseItemsRequest) GetCanEditItemsOk() (*bool, bool)`

GetCanEditItemsOk returns a tuple with the CanEditItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditItems

`func (o *ApiBaseItemsRequest) SetCanEditItems(v bool)`

SetCanEditItems sets CanEditItems field to given value.

### HasCanEditItems

`func (o *ApiBaseItemsRequest) HasCanEditItems() bool`

HasCanEditItems returns a boolean if a field has been set.

### SetCanEditItemsNil

`func (o *ApiBaseItemsRequest) SetCanEditItemsNil(b bool)`

 SetCanEditItemsNil sets the value for CanEditItems to be an explicit nil

### UnsetCanEditItems
`func (o *ApiBaseItemsRequest) UnsetCanEditItems()`

UnsetCanEditItems ensures that no value is present for CanEditItems, not even an explicit nil
### GetGroupItemsInto

`func (o *ApiBaseItemsRequest) GetGroupItemsInto() LibraryItemLinkType`

GetGroupItemsInto returns the GroupItemsInto field if non-nil, zero value otherwise.

### GetGroupItemsIntoOk

`func (o *ApiBaseItemsRequest) GetGroupItemsIntoOk() (*LibraryItemLinkType, bool)`

GetGroupItemsIntoOk returns a tuple with the GroupItemsInto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupItemsInto

`func (o *ApiBaseItemsRequest) SetGroupItemsInto(v LibraryItemLinkType)`

SetGroupItemsInto sets GroupItemsInto field to given value.

### HasGroupItemsInto

`func (o *ApiBaseItemsRequest) HasGroupItemsInto() bool`

HasGroupItemsInto returns a boolean if a field has been set.

### GetIsStandaloneSpecial

`func (o *ApiBaseItemsRequest) GetIsStandaloneSpecial() bool`

GetIsStandaloneSpecial returns the IsStandaloneSpecial field if non-nil, zero value otherwise.

### GetIsStandaloneSpecialOk

`func (o *ApiBaseItemsRequest) GetIsStandaloneSpecialOk() (*bool, bool)`

GetIsStandaloneSpecialOk returns a tuple with the IsStandaloneSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandaloneSpecial

`func (o *ApiBaseItemsRequest) SetIsStandaloneSpecial(v bool)`

SetIsStandaloneSpecial sets IsStandaloneSpecial field to given value.

### HasIsStandaloneSpecial

`func (o *ApiBaseItemsRequest) HasIsStandaloneSpecial() bool`

HasIsStandaloneSpecial returns a boolean if a field has been set.

### SetIsStandaloneSpecialNil

`func (o *ApiBaseItemsRequest) SetIsStandaloneSpecialNil(b bool)`

 SetIsStandaloneSpecialNil sets the value for IsStandaloneSpecial to be an explicit nil

### UnsetIsStandaloneSpecial
`func (o *ApiBaseItemsRequest) UnsetIsStandaloneSpecial()`

UnsetIsStandaloneSpecial ensures that no value is present for IsStandaloneSpecial, not even an explicit nil
### GetMinWidth

`func (o *ApiBaseItemsRequest) GetMinWidth() int32`

GetMinWidth returns the MinWidth field if non-nil, zero value otherwise.

### GetMinWidthOk

`func (o *ApiBaseItemsRequest) GetMinWidthOk() (*int32, bool)`

GetMinWidthOk returns a tuple with the MinWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWidth

`func (o *ApiBaseItemsRequest) SetMinWidth(v int32)`

SetMinWidth sets MinWidth field to given value.

### HasMinWidth

`func (o *ApiBaseItemsRequest) HasMinWidth() bool`

HasMinWidth returns a boolean if a field has been set.

### SetMinWidthNil

`func (o *ApiBaseItemsRequest) SetMinWidthNil(b bool)`

 SetMinWidthNil sets the value for MinWidth to be an explicit nil

### UnsetMinWidth
`func (o *ApiBaseItemsRequest) UnsetMinWidth()`

UnsetMinWidth ensures that no value is present for MinWidth, not even an explicit nil
### GetMinHeight

`func (o *ApiBaseItemsRequest) GetMinHeight() int32`

GetMinHeight returns the MinHeight field if non-nil, zero value otherwise.

### GetMinHeightOk

`func (o *ApiBaseItemsRequest) GetMinHeightOk() (*int32, bool)`

GetMinHeightOk returns a tuple with the MinHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHeight

`func (o *ApiBaseItemsRequest) SetMinHeight(v int32)`

SetMinHeight sets MinHeight field to given value.

### HasMinHeight

`func (o *ApiBaseItemsRequest) HasMinHeight() bool`

HasMinHeight returns a boolean if a field has been set.

### SetMinHeightNil

`func (o *ApiBaseItemsRequest) SetMinHeightNil(b bool)`

 SetMinHeightNil sets the value for MinHeight to be an explicit nil

### UnsetMinHeight
`func (o *ApiBaseItemsRequest) UnsetMinHeight()`

UnsetMinHeight ensures that no value is present for MinHeight, not even an explicit nil
### GetMaxWidth

`func (o *ApiBaseItemsRequest) GetMaxWidth() int32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *ApiBaseItemsRequest) GetMaxWidthOk() (*int32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *ApiBaseItemsRequest) SetMaxWidth(v int32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *ApiBaseItemsRequest) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### SetMaxWidthNil

`func (o *ApiBaseItemsRequest) SetMaxWidthNil(b bool)`

 SetMaxWidthNil sets the value for MaxWidth to be an explicit nil

### UnsetMaxWidth
`func (o *ApiBaseItemsRequest) UnsetMaxWidth()`

UnsetMaxWidth ensures that no value is present for MaxWidth, not even an explicit nil
### GetMaxHeight

`func (o *ApiBaseItemsRequest) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *ApiBaseItemsRequest) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *ApiBaseItemsRequest) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *ApiBaseItemsRequest) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### SetMaxHeightNil

`func (o *ApiBaseItemsRequest) SetMaxHeightNil(b bool)`

 SetMaxHeightNil sets the value for MaxHeight to be an explicit nil

### UnsetMaxHeight
`func (o *ApiBaseItemsRequest) UnsetMaxHeight()`

UnsetMaxHeight ensures that no value is present for MaxHeight, not even an explicit nil
### GetGroupProgramsBySeries

`func (o *ApiBaseItemsRequest) GetGroupProgramsBySeries() bool`

GetGroupProgramsBySeries returns the GroupProgramsBySeries field if non-nil, zero value otherwise.

### GetGroupProgramsBySeriesOk

`func (o *ApiBaseItemsRequest) GetGroupProgramsBySeriesOk() (*bool, bool)`

GetGroupProgramsBySeriesOk returns a tuple with the GroupProgramsBySeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupProgramsBySeries

`func (o *ApiBaseItemsRequest) SetGroupProgramsBySeries(v bool)`

SetGroupProgramsBySeries sets GroupProgramsBySeries field to given value.

### HasGroupProgramsBySeries

`func (o *ApiBaseItemsRequest) HasGroupProgramsBySeries() bool`

HasGroupProgramsBySeries returns a boolean if a field has been set.

### GetGroupByPresentationUniqueKey

`func (o *ApiBaseItemsRequest) GetGroupByPresentationUniqueKey() bool`

GetGroupByPresentationUniqueKey returns the GroupByPresentationUniqueKey field if non-nil, zero value otherwise.

### GetGroupByPresentationUniqueKeyOk

`func (o *ApiBaseItemsRequest) GetGroupByPresentationUniqueKeyOk() (*bool, bool)`

GetGroupByPresentationUniqueKeyOk returns a tuple with the GroupByPresentationUniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByPresentationUniqueKey

`func (o *ApiBaseItemsRequest) SetGroupByPresentationUniqueKey(v bool)`

SetGroupByPresentationUniqueKey sets GroupByPresentationUniqueKey field to given value.

### HasGroupByPresentationUniqueKey

`func (o *ApiBaseItemsRequest) HasGroupByPresentationUniqueKey() bool`

HasGroupByPresentationUniqueKey returns a boolean if a field has been set.

### SetGroupByPresentationUniqueKeyNil

`func (o *ApiBaseItemsRequest) SetGroupByPresentationUniqueKeyNil(b bool)`

 SetGroupByPresentationUniqueKeyNil sets the value for GroupByPresentationUniqueKey to be an explicit nil

### UnsetGroupByPresentationUniqueKey
`func (o *ApiBaseItemsRequest) UnsetGroupByPresentationUniqueKey()`

UnsetGroupByPresentationUniqueKey ensures that no value is present for GroupByPresentationUniqueKey, not even an explicit nil
### GetAirDays

`func (o *ApiBaseItemsRequest) GetAirDays() []DayOfWeek`

GetAirDays returns the AirDays field if non-nil, zero value otherwise.

### GetAirDaysOk

`func (o *ApiBaseItemsRequest) GetAirDaysOk() (*[]DayOfWeek, bool)`

GetAirDaysOk returns a tuple with the AirDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDays

`func (o *ApiBaseItemsRequest) SetAirDays(v []DayOfWeek)`

SetAirDays sets AirDays field to given value.

### HasAirDays

`func (o *ApiBaseItemsRequest) HasAirDays() bool`

HasAirDays returns a boolean if a field has been set.

### GetIsAiring

`func (o *ApiBaseItemsRequest) GetIsAiring() bool`

GetIsAiring returns the IsAiring field if non-nil, zero value otherwise.

### GetIsAiringOk

`func (o *ApiBaseItemsRequest) GetIsAiringOk() (*bool, bool)`

GetIsAiringOk returns a tuple with the IsAiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAiring

`func (o *ApiBaseItemsRequest) SetIsAiring(v bool)`

SetIsAiring sets IsAiring field to given value.

### HasIsAiring

`func (o *ApiBaseItemsRequest) HasIsAiring() bool`

HasIsAiring returns a boolean if a field has been set.

### SetIsAiringNil

`func (o *ApiBaseItemsRequest) SetIsAiringNil(b bool)`

 SetIsAiringNil sets the value for IsAiring to be an explicit nil

### UnsetIsAiring
`func (o *ApiBaseItemsRequest) UnsetIsAiring()`

UnsetIsAiring ensures that no value is present for IsAiring, not even an explicit nil
### GetHasAired

`func (o *ApiBaseItemsRequest) GetHasAired() bool`

GetHasAired returns the HasAired field if non-nil, zero value otherwise.

### GetHasAiredOk

`func (o *ApiBaseItemsRequest) GetHasAiredOk() (*bool, bool)`

GetHasAiredOk returns a tuple with the HasAired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAired

`func (o *ApiBaseItemsRequest) SetHasAired(v bool)`

SetHasAired sets HasAired field to given value.

### HasHasAired

`func (o *ApiBaseItemsRequest) HasHasAired() bool`

HasHasAired returns a boolean if a field has been set.

### SetHasAiredNil

`func (o *ApiBaseItemsRequest) SetHasAiredNil(b bool)`

 SetHasAiredNil sets the value for HasAired to be an explicit nil

### UnsetHasAired
`func (o *ApiBaseItemsRequest) UnsetHasAired()`

UnsetHasAired ensures that no value is present for HasAired, not even an explicit nil
### GetCollectionTypes

`func (o *ApiBaseItemsRequest) GetCollectionTypes() string`

GetCollectionTypes returns the CollectionTypes field if non-nil, zero value otherwise.

### GetCollectionTypesOk

`func (o *ApiBaseItemsRequest) GetCollectionTypesOk() (*string, bool)`

GetCollectionTypesOk returns a tuple with the CollectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTypes

`func (o *ApiBaseItemsRequest) SetCollectionTypes(v string)`

SetCollectionTypes sets CollectionTypes field to given value.

### HasCollectionTypes

`func (o *ApiBaseItemsRequest) HasCollectionTypes() bool`

HasCollectionTypes returns a boolean if a field has been set.

### GetExcludeSources

`func (o *ApiBaseItemsRequest) GetExcludeSources() []string`

GetExcludeSources returns the ExcludeSources field if non-nil, zero value otherwise.

### GetExcludeSourcesOk

`func (o *ApiBaseItemsRequest) GetExcludeSourcesOk() (*[]string, bool)`

GetExcludeSourcesOk returns a tuple with the ExcludeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSources

`func (o *ApiBaseItemsRequest) SetExcludeSources(v []string)`

SetExcludeSources sets ExcludeSources field to given value.

### HasExcludeSources

`func (o *ApiBaseItemsRequest) HasExcludeSources() bool`

HasExcludeSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


