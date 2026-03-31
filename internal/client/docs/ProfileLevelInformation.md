# ProfileLevelInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**ProfileInformation**](ProfileInformation.md) |  | [optional] 
**Level** | Pointer to [**LevelInformation**](LevelInformation.md) |  | [optional] 

## Methods

### NewProfileLevelInformation

`func NewProfileLevelInformation() *ProfileLevelInformation`

NewProfileLevelInformation instantiates a new ProfileLevelInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileLevelInformationWithDefaults

`func NewProfileLevelInformationWithDefaults() *ProfileLevelInformation`

NewProfileLevelInformationWithDefaults instantiates a new ProfileLevelInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *ProfileLevelInformation) GetProfile() ProfileInformation`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ProfileLevelInformation) GetProfileOk() (*ProfileInformation, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ProfileLevelInformation) SetProfile(v ProfileInformation)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ProfileLevelInformation) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetLevel

`func (o *ProfileLevelInformation) GetLevel() LevelInformation`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ProfileLevelInformation) GetLevelOk() (*LevelInformation, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ProfileLevelInformation) SetLevel(v LevelInformation)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ProfileLevelInformation) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


