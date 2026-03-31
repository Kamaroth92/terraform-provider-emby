# UIViewInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewId** | Pointer to **string** |  | [optional] 
**PageId** | Pointer to **string** |  | [optional] 
**Caption** | Pointer to **string** |  | [optional] 
**SubCaption** | Pointer to **string** |  | [optional] 
**PluginId** | Pointer to **string** |  | [optional] 
**ViewType** | Pointer to [**EnumsUIViewType**](EnumsUIViewType.md) |  | [optional] 
**ShowDialogFullScreen** | Pointer to **bool** |  | [optional] 
**IsInSequence** | Pointer to **bool** |  | [optional] 
**RedirectViewUrl** | Pointer to **string** |  | [optional] 
**EditObjectContainer** | Pointer to [**GenericEditIEditObjectContainer**](GenericEditIEditObjectContainer.md) |  | [optional] 
**Commands** | Pointer to [**[]UICommand**](UICommand.md) |  | [optional] 
**TabPageInfos** | Pointer to [**[]UITabPageInfo**](UITabPageInfo.md) |  | [optional] 
**IsPageChangeInfo** | Pointer to **bool** |  | [optional] 

## Methods

### NewUIViewInfo

`func NewUIViewInfo() *UIViewInfo`

NewUIViewInfo instantiates a new UIViewInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUIViewInfoWithDefaults

`func NewUIViewInfoWithDefaults() *UIViewInfo`

NewUIViewInfoWithDefaults instantiates a new UIViewInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewId

`func (o *UIViewInfo) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *UIViewInfo) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *UIViewInfo) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *UIViewInfo) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetPageId

`func (o *UIViewInfo) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *UIViewInfo) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *UIViewInfo) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *UIViewInfo) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetCaption

`func (o *UIViewInfo) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *UIViewInfo) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *UIViewInfo) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *UIViewInfo) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetSubCaption

`func (o *UIViewInfo) GetSubCaption() string`

GetSubCaption returns the SubCaption field if non-nil, zero value otherwise.

### GetSubCaptionOk

`func (o *UIViewInfo) GetSubCaptionOk() (*string, bool)`

GetSubCaptionOk returns a tuple with the SubCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCaption

`func (o *UIViewInfo) SetSubCaption(v string)`

SetSubCaption sets SubCaption field to given value.

### HasSubCaption

`func (o *UIViewInfo) HasSubCaption() bool`

HasSubCaption returns a boolean if a field has been set.

### GetPluginId

`func (o *UIViewInfo) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *UIViewInfo) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *UIViewInfo) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *UIViewInfo) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetViewType

`func (o *UIViewInfo) GetViewType() EnumsUIViewType`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *UIViewInfo) GetViewTypeOk() (*EnumsUIViewType, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *UIViewInfo) SetViewType(v EnumsUIViewType)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *UIViewInfo) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetShowDialogFullScreen

`func (o *UIViewInfo) GetShowDialogFullScreen() bool`

GetShowDialogFullScreen returns the ShowDialogFullScreen field if non-nil, zero value otherwise.

### GetShowDialogFullScreenOk

`func (o *UIViewInfo) GetShowDialogFullScreenOk() (*bool, bool)`

GetShowDialogFullScreenOk returns a tuple with the ShowDialogFullScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDialogFullScreen

`func (o *UIViewInfo) SetShowDialogFullScreen(v bool)`

SetShowDialogFullScreen sets ShowDialogFullScreen field to given value.

### HasShowDialogFullScreen

`func (o *UIViewInfo) HasShowDialogFullScreen() bool`

HasShowDialogFullScreen returns a boolean if a field has been set.

### GetIsInSequence

`func (o *UIViewInfo) GetIsInSequence() bool`

GetIsInSequence returns the IsInSequence field if non-nil, zero value otherwise.

### GetIsInSequenceOk

`func (o *UIViewInfo) GetIsInSequenceOk() (*bool, bool)`

GetIsInSequenceOk returns a tuple with the IsInSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInSequence

`func (o *UIViewInfo) SetIsInSequence(v bool)`

SetIsInSequence sets IsInSequence field to given value.

### HasIsInSequence

`func (o *UIViewInfo) HasIsInSequence() bool`

HasIsInSequence returns a boolean if a field has been set.

### GetRedirectViewUrl

`func (o *UIViewInfo) GetRedirectViewUrl() string`

GetRedirectViewUrl returns the RedirectViewUrl field if non-nil, zero value otherwise.

### GetRedirectViewUrlOk

`func (o *UIViewInfo) GetRedirectViewUrlOk() (*string, bool)`

GetRedirectViewUrlOk returns a tuple with the RedirectViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectViewUrl

`func (o *UIViewInfo) SetRedirectViewUrl(v string)`

SetRedirectViewUrl sets RedirectViewUrl field to given value.

### HasRedirectViewUrl

`func (o *UIViewInfo) HasRedirectViewUrl() bool`

HasRedirectViewUrl returns a boolean if a field has been set.

### GetEditObjectContainer

`func (o *UIViewInfo) GetEditObjectContainer() GenericEditIEditObjectContainer`

GetEditObjectContainer returns the EditObjectContainer field if non-nil, zero value otherwise.

### GetEditObjectContainerOk

`func (o *UIViewInfo) GetEditObjectContainerOk() (*GenericEditIEditObjectContainer, bool)`

GetEditObjectContainerOk returns a tuple with the EditObjectContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditObjectContainer

`func (o *UIViewInfo) SetEditObjectContainer(v GenericEditIEditObjectContainer)`

SetEditObjectContainer sets EditObjectContainer field to given value.

### HasEditObjectContainer

`func (o *UIViewInfo) HasEditObjectContainer() bool`

HasEditObjectContainer returns a boolean if a field has been set.

### GetCommands

`func (o *UIViewInfo) GetCommands() []UICommand`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *UIViewInfo) GetCommandsOk() (*[]UICommand, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *UIViewInfo) SetCommands(v []UICommand)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *UIViewInfo) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetTabPageInfos

`func (o *UIViewInfo) GetTabPageInfos() []UITabPageInfo`

GetTabPageInfos returns the TabPageInfos field if non-nil, zero value otherwise.

### GetTabPageInfosOk

`func (o *UIViewInfo) GetTabPageInfosOk() (*[]UITabPageInfo, bool)`

GetTabPageInfosOk returns a tuple with the TabPageInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabPageInfos

`func (o *UIViewInfo) SetTabPageInfos(v []UITabPageInfo)`

SetTabPageInfos sets TabPageInfos field to given value.

### HasTabPageInfos

`func (o *UIViewInfo) HasTabPageInfos() bool`

HasTabPageInfos returns a boolean if a field has been set.

### GetIsPageChangeInfo

`func (o *UIViewInfo) GetIsPageChangeInfo() bool`

GetIsPageChangeInfo returns the IsPageChangeInfo field if non-nil, zero value otherwise.

### GetIsPageChangeInfoOk

`func (o *UIViewInfo) GetIsPageChangeInfoOk() (*bool, bool)`

GetIsPageChangeInfoOk returns a tuple with the IsPageChangeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPageChangeInfo

`func (o *UIViewInfo) SetIsPageChangeInfo(v bool)`

SetIsPageChangeInfo sets IsPageChangeInfo field to given value.

### HasIsPageChangeInfo

`func (o *UIViewInfo) HasIsPageChangeInfo() bool`

HasIsPageChangeInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


