# NotificationCategoryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**[]NotificationTypeInfo**](NotificationTypeInfo.md) |  | [optional] 

## Methods

### NewNotificationCategoryInfo

`func NewNotificationCategoryInfo() *NotificationCategoryInfo`

NewNotificationCategoryInfo instantiates a new NotificationCategoryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationCategoryInfoWithDefaults

`func NewNotificationCategoryInfoWithDefaults() *NotificationCategoryInfo`

NewNotificationCategoryInfoWithDefaults instantiates a new NotificationCategoryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationCategoryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationCategoryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationCategoryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationCategoryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *NotificationCategoryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationCategoryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationCategoryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationCategoryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEvents

`func (o *NotificationCategoryInfo) GetEvents() []NotificationTypeInfo`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *NotificationCategoryInfo) GetEventsOk() (*[]NotificationTypeInfo, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *NotificationCategoryInfo) SetEvents(v []NotificationTypeInfo)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *NotificationCategoryInfo) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


