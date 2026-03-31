# ContentSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Subtitle** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SectionType** | Pointer to **string** |  | [optional] 
**CollectionType** | Pointer to **string** |  | [optional] 
**ViewType** | Pointer to **string** |  | [optional] 
**Monitor** | Pointer to **[]string** |  | [optional] 
**CardSizeOffset** | Pointer to **int32** |  | [optional] 
**ScrollDirection** | Pointer to [**ScrollDirection**](ScrollDirection.md) |  | [optional] 
**ParentItem** | Pointer to [**BaseItemDto**](BaseItemDto.md) |  | [optional] 
**TextInfo** | Pointer to [**TextSectionInfo**](TextSectionInfo.md) |  | [optional] 
**PremiumFeature** | Pointer to **string** |  | [optional] 
**PremiumMessage** | Pointer to **string** |  | [optional] 
**RefreshInterval** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewContentSection

`func NewContentSection() *ContentSection`

NewContentSection instantiates a new ContentSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSectionWithDefaults

`func NewContentSectionWithDefaults() *ContentSection`

NewContentSectionWithDefaults instantiates a new ContentSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContentSection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentSection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentSection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentSection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubtitle

`func (o *ContentSection) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *ContentSection) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *ContentSection) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *ContentSection) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetId

`func (o *ContentSection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentSection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentSection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContentSection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSectionType

`func (o *ContentSection) GetSectionType() string`

GetSectionType returns the SectionType field if non-nil, zero value otherwise.

### GetSectionTypeOk

`func (o *ContentSection) GetSectionTypeOk() (*string, bool)`

GetSectionTypeOk returns a tuple with the SectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionType

`func (o *ContentSection) SetSectionType(v string)`

SetSectionType sets SectionType field to given value.

### HasSectionType

`func (o *ContentSection) HasSectionType() bool`

HasSectionType returns a boolean if a field has been set.

### GetCollectionType

`func (o *ContentSection) GetCollectionType() string`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *ContentSection) GetCollectionTypeOk() (*string, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *ContentSection) SetCollectionType(v string)`

SetCollectionType sets CollectionType field to given value.

### HasCollectionType

`func (o *ContentSection) HasCollectionType() bool`

HasCollectionType returns a boolean if a field has been set.

### GetViewType

`func (o *ContentSection) GetViewType() string`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *ContentSection) GetViewTypeOk() (*string, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *ContentSection) SetViewType(v string)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *ContentSection) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetMonitor

`func (o *ContentSection) GetMonitor() []string`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *ContentSection) GetMonitorOk() (*[]string, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *ContentSection) SetMonitor(v []string)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *ContentSection) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetCardSizeOffset

`func (o *ContentSection) GetCardSizeOffset() int32`

GetCardSizeOffset returns the CardSizeOffset field if non-nil, zero value otherwise.

### GetCardSizeOffsetOk

`func (o *ContentSection) GetCardSizeOffsetOk() (*int32, bool)`

GetCardSizeOffsetOk returns a tuple with the CardSizeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSizeOffset

`func (o *ContentSection) SetCardSizeOffset(v int32)`

SetCardSizeOffset sets CardSizeOffset field to given value.

### HasCardSizeOffset

`func (o *ContentSection) HasCardSizeOffset() bool`

HasCardSizeOffset returns a boolean if a field has been set.

### GetScrollDirection

`func (o *ContentSection) GetScrollDirection() ScrollDirection`

GetScrollDirection returns the ScrollDirection field if non-nil, zero value otherwise.

### GetScrollDirectionOk

`func (o *ContentSection) GetScrollDirectionOk() (*ScrollDirection, bool)`

GetScrollDirectionOk returns a tuple with the ScrollDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollDirection

`func (o *ContentSection) SetScrollDirection(v ScrollDirection)`

SetScrollDirection sets ScrollDirection field to given value.

### HasScrollDirection

`func (o *ContentSection) HasScrollDirection() bool`

HasScrollDirection returns a boolean if a field has been set.

### GetParentItem

`func (o *ContentSection) GetParentItem() BaseItemDto`

GetParentItem returns the ParentItem field if non-nil, zero value otherwise.

### GetParentItemOk

`func (o *ContentSection) GetParentItemOk() (*BaseItemDto, bool)`

GetParentItemOk returns a tuple with the ParentItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentItem

`func (o *ContentSection) SetParentItem(v BaseItemDto)`

SetParentItem sets ParentItem field to given value.

### HasParentItem

`func (o *ContentSection) HasParentItem() bool`

HasParentItem returns a boolean if a field has been set.

### GetTextInfo

`func (o *ContentSection) GetTextInfo() TextSectionInfo`

GetTextInfo returns the TextInfo field if non-nil, zero value otherwise.

### GetTextInfoOk

`func (o *ContentSection) GetTextInfoOk() (*TextSectionInfo, bool)`

GetTextInfoOk returns a tuple with the TextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextInfo

`func (o *ContentSection) SetTextInfo(v TextSectionInfo)`

SetTextInfo sets TextInfo field to given value.

### HasTextInfo

`func (o *ContentSection) HasTextInfo() bool`

HasTextInfo returns a boolean if a field has been set.

### GetPremiumFeature

`func (o *ContentSection) GetPremiumFeature() string`

GetPremiumFeature returns the PremiumFeature field if non-nil, zero value otherwise.

### GetPremiumFeatureOk

`func (o *ContentSection) GetPremiumFeatureOk() (*string, bool)`

GetPremiumFeatureOk returns a tuple with the PremiumFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumFeature

`func (o *ContentSection) SetPremiumFeature(v string)`

SetPremiumFeature sets PremiumFeature field to given value.

### HasPremiumFeature

`func (o *ContentSection) HasPremiumFeature() bool`

HasPremiumFeature returns a boolean if a field has been set.

### GetPremiumMessage

`func (o *ContentSection) GetPremiumMessage() string`

GetPremiumMessage returns the PremiumMessage field if non-nil, zero value otherwise.

### GetPremiumMessageOk

`func (o *ContentSection) GetPremiumMessageOk() (*string, bool)`

GetPremiumMessageOk returns a tuple with the PremiumMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumMessage

`func (o *ContentSection) SetPremiumMessage(v string)`

SetPremiumMessage sets PremiumMessage field to given value.

### HasPremiumMessage

`func (o *ContentSection) HasPremiumMessage() bool`

HasPremiumMessage returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *ContentSection) GetRefreshInterval() int32`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *ContentSection) GetRefreshIntervalOk() (*int32, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *ContentSection) SetRefreshInterval(v int32)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *ContentSection) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### SetRefreshIntervalNil

`func (o *ContentSection) SetRefreshIntervalNil(b bool)`

 SetRefreshIntervalNil sets the value for RefreshInterval to be an explicit nil

### UnsetRefreshInterval
`func (o *ContentSection) UnsetRefreshInterval()`

UnsetRefreshInterval ensures that no value is present for RefreshInterval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


