# TextSectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Level** | Pointer to [**NotificationsNotificationLevel**](NotificationsNotificationLevel.md) |  | [optional] 

## Methods

### NewTextSectionInfo

`func NewTextSectionInfo() *TextSectionInfo`

NewTextSectionInfo instantiates a new TextSectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextSectionInfoWithDefaults

`func NewTextSectionInfoWithDefaults() *TextSectionInfo`

NewTextSectionInfoWithDefaults instantiates a new TextSectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *TextSectionInfo) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TextSectionInfo) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TextSectionInfo) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TextSectionInfo) HasText() bool`

HasText returns a boolean if a field has been set.

### GetName

`func (o *TextSectionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TextSectionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TextSectionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TextSectionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *TextSectionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TextSectionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TextSectionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TextSectionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *TextSectionInfo) GetLevel() NotificationsNotificationLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *TextSectionInfo) GetLevelOk() (*NotificationsNotificationLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *TextSectionInfo) SetLevel(v NotificationsNotificationLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *TextSectionInfo) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


