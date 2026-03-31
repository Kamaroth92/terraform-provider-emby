# ResponseProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** |  | [optional] 
**AudioCodec** | Pointer to **string** |  | [optional] 
**VideoCodec** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DlnaProfileType**](DlnaProfileType.md) |  | [optional] 
**OrgPn** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]ProfileCondition**](ProfileCondition.md) |  | [optional] 

## Methods

### NewResponseProfile

`func NewResponseProfile() *ResponseProfile`

NewResponseProfile instantiates a new ResponseProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseProfileWithDefaults

`func NewResponseProfileWithDefaults() *ResponseProfile`

NewResponseProfileWithDefaults instantiates a new ResponseProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *ResponseProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ResponseProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ResponseProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ResponseProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetAudioCodec

`func (o *ResponseProfile) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *ResponseProfile) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *ResponseProfile) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *ResponseProfile) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetVideoCodec

`func (o *ResponseProfile) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *ResponseProfile) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *ResponseProfile) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *ResponseProfile) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetType

`func (o *ResponseProfile) GetType() DlnaProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseProfile) GetTypeOk() (*DlnaProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseProfile) SetType(v DlnaProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrgPn

`func (o *ResponseProfile) GetOrgPn() string`

GetOrgPn returns the OrgPn field if non-nil, zero value otherwise.

### GetOrgPnOk

`func (o *ResponseProfile) GetOrgPnOk() (*string, bool)`

GetOrgPnOk returns a tuple with the OrgPn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPn

`func (o *ResponseProfile) SetOrgPn(v string)`

SetOrgPn sets OrgPn field to given value.

### HasOrgPn

`func (o *ResponseProfile) HasOrgPn() bool`

HasOrgPn returns a boolean if a field has been set.

### GetMimeType

`func (o *ResponseProfile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResponseProfile) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResponseProfile) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ResponseProfile) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetConditions

`func (o *ResponseProfile) GetConditions() []ProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ResponseProfile) GetConditionsOk() (*[]ProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ResponseProfile) SetConditions(v []ProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ResponseProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


