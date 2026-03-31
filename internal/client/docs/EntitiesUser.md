# EntitiesUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsesIdForConfigurationPath** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**EasyPassword** | Pointer to **string** |  | [optional] 
**Salt** | Pointer to **string** |  | [optional] 
**ConnectUserName** | Pointer to **string** |  | [optional] 
**ConnectUserId** | Pointer to **string** |  | [optional] 
**ConnectLinkType** | Pointer to [**ConnectUserLinkType**](ConnectUserLinkType.md) |  | [optional] 
**ConnectAccessKey** | Pointer to **string** |  | [optional] 
**ImageInfos** | Pointer to [**[]EntitiesItemImageInfo**](EntitiesItemImageInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LastLoginDate** | Pointer to **NullableTime** |  | [optional] 
**LastActivityDate** | Pointer to **NullableTime** |  | [optional] 
**PlayedPercentage** | Pointer to **NullableFloat64** |  | [optional] 
**RecursiveChildCountEqualsChildCount** | Pointer to **bool** |  | [optional] 
**OriginalParsedName** | Pointer to **string** |  | [optional] 
**IsNameParsedFromFolder** | Pointer to **bool** |  | [optional] 
**IdString** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**ImportedCollections** | Pointer to [**[]LinkedItemInfo**](LinkedItemInfo.md) |  | [optional] 
**ResolvedPresentationUniqueKey** | Pointer to **string** |  | [optional] 

## Methods

### NewEntitiesUser

`func NewEntitiesUser() *EntitiesUser`

NewEntitiesUser instantiates a new EntitiesUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitiesUserWithDefaults

`func NewEntitiesUserWithDefaults() *EntitiesUser`

NewEntitiesUserWithDefaults instantiates a new EntitiesUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsesIdForConfigurationPath

`func (o *EntitiesUser) GetUsesIdForConfigurationPath() bool`

GetUsesIdForConfigurationPath returns the UsesIdForConfigurationPath field if non-nil, zero value otherwise.

### GetUsesIdForConfigurationPathOk

`func (o *EntitiesUser) GetUsesIdForConfigurationPathOk() (*bool, bool)`

GetUsesIdForConfigurationPathOk returns a tuple with the UsesIdForConfigurationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesIdForConfigurationPath

`func (o *EntitiesUser) SetUsesIdForConfigurationPath(v bool)`

SetUsesIdForConfigurationPath sets UsesIdForConfigurationPath field to given value.

### HasUsesIdForConfigurationPath

`func (o *EntitiesUser) HasUsesIdForConfigurationPath() bool`

HasUsesIdForConfigurationPath returns a boolean if a field has been set.

### GetPassword

`func (o *EntitiesUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EntitiesUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EntitiesUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EntitiesUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEasyPassword

`func (o *EntitiesUser) GetEasyPassword() string`

GetEasyPassword returns the EasyPassword field if non-nil, zero value otherwise.

### GetEasyPasswordOk

`func (o *EntitiesUser) GetEasyPasswordOk() (*string, bool)`

GetEasyPasswordOk returns a tuple with the EasyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasyPassword

`func (o *EntitiesUser) SetEasyPassword(v string)`

SetEasyPassword sets EasyPassword field to given value.

### HasEasyPassword

`func (o *EntitiesUser) HasEasyPassword() bool`

HasEasyPassword returns a boolean if a field has been set.

### GetSalt

`func (o *EntitiesUser) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *EntitiesUser) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *EntitiesUser) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *EntitiesUser) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetConnectUserName

`func (o *EntitiesUser) GetConnectUserName() string`

GetConnectUserName returns the ConnectUserName field if non-nil, zero value otherwise.

### GetConnectUserNameOk

`func (o *EntitiesUser) GetConnectUserNameOk() (*string, bool)`

GetConnectUserNameOk returns a tuple with the ConnectUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectUserName

`func (o *EntitiesUser) SetConnectUserName(v string)`

SetConnectUserName sets ConnectUserName field to given value.

### HasConnectUserName

`func (o *EntitiesUser) HasConnectUserName() bool`

HasConnectUserName returns a boolean if a field has been set.

### GetConnectUserId

`func (o *EntitiesUser) GetConnectUserId() string`

GetConnectUserId returns the ConnectUserId field if non-nil, zero value otherwise.

### GetConnectUserIdOk

`func (o *EntitiesUser) GetConnectUserIdOk() (*string, bool)`

GetConnectUserIdOk returns a tuple with the ConnectUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectUserId

`func (o *EntitiesUser) SetConnectUserId(v string)`

SetConnectUserId sets ConnectUserId field to given value.

### HasConnectUserId

`func (o *EntitiesUser) HasConnectUserId() bool`

HasConnectUserId returns a boolean if a field has been set.

### GetConnectLinkType

`func (o *EntitiesUser) GetConnectLinkType() ConnectUserLinkType`

GetConnectLinkType returns the ConnectLinkType field if non-nil, zero value otherwise.

### GetConnectLinkTypeOk

`func (o *EntitiesUser) GetConnectLinkTypeOk() (*ConnectUserLinkType, bool)`

GetConnectLinkTypeOk returns a tuple with the ConnectLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectLinkType

`func (o *EntitiesUser) SetConnectLinkType(v ConnectUserLinkType)`

SetConnectLinkType sets ConnectLinkType field to given value.

### HasConnectLinkType

`func (o *EntitiesUser) HasConnectLinkType() bool`

HasConnectLinkType returns a boolean if a field has been set.

### GetConnectAccessKey

`func (o *EntitiesUser) GetConnectAccessKey() string`

GetConnectAccessKey returns the ConnectAccessKey field if non-nil, zero value otherwise.

### GetConnectAccessKeyOk

`func (o *EntitiesUser) GetConnectAccessKeyOk() (*string, bool)`

GetConnectAccessKeyOk returns a tuple with the ConnectAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectAccessKey

`func (o *EntitiesUser) SetConnectAccessKey(v string)`

SetConnectAccessKey sets ConnectAccessKey field to given value.

### HasConnectAccessKey

`func (o *EntitiesUser) HasConnectAccessKey() bool`

HasConnectAccessKey returns a boolean if a field has been set.

### GetImageInfos

`func (o *EntitiesUser) GetImageInfos() []EntitiesItemImageInfo`

GetImageInfos returns the ImageInfos field if non-nil, zero value otherwise.

### GetImageInfosOk

`func (o *EntitiesUser) GetImageInfosOk() (*[]EntitiesItemImageInfo, bool)`

GetImageInfosOk returns a tuple with the ImageInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageInfos

`func (o *EntitiesUser) SetImageInfos(v []EntitiesItemImageInfo)`

SetImageInfos sets ImageInfos field to given value.

### HasImageInfos

`func (o *EntitiesUser) HasImageInfos() bool`

HasImageInfos returns a boolean if a field has been set.

### GetName

`func (o *EntitiesUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitiesUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitiesUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntitiesUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *EntitiesUser) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *EntitiesUser) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *EntitiesUser) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *EntitiesUser) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### SetLastLoginDateNil

`func (o *EntitiesUser) SetLastLoginDateNil(b bool)`

 SetLastLoginDateNil sets the value for LastLoginDate to be an explicit nil

### UnsetLastLoginDate
`func (o *EntitiesUser) UnsetLastLoginDate()`

UnsetLastLoginDate ensures that no value is present for LastLoginDate, not even an explicit nil
### GetLastActivityDate

`func (o *EntitiesUser) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *EntitiesUser) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *EntitiesUser) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *EntitiesUser) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### SetLastActivityDateNil

`func (o *EntitiesUser) SetLastActivityDateNil(b bool)`

 SetLastActivityDateNil sets the value for LastActivityDate to be an explicit nil

### UnsetLastActivityDate
`func (o *EntitiesUser) UnsetLastActivityDate()`

UnsetLastActivityDate ensures that no value is present for LastActivityDate, not even an explicit nil
### GetPlayedPercentage

`func (o *EntitiesUser) GetPlayedPercentage() float64`

GetPlayedPercentage returns the PlayedPercentage field if non-nil, zero value otherwise.

### GetPlayedPercentageOk

`func (o *EntitiesUser) GetPlayedPercentageOk() (*float64, bool)`

GetPlayedPercentageOk returns a tuple with the PlayedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayedPercentage

`func (o *EntitiesUser) SetPlayedPercentage(v float64)`

SetPlayedPercentage sets PlayedPercentage field to given value.

### HasPlayedPercentage

`func (o *EntitiesUser) HasPlayedPercentage() bool`

HasPlayedPercentage returns a boolean if a field has been set.

### SetPlayedPercentageNil

`func (o *EntitiesUser) SetPlayedPercentageNil(b bool)`

 SetPlayedPercentageNil sets the value for PlayedPercentage to be an explicit nil

### UnsetPlayedPercentage
`func (o *EntitiesUser) UnsetPlayedPercentage()`

UnsetPlayedPercentage ensures that no value is present for PlayedPercentage, not even an explicit nil
### GetRecursiveChildCountEqualsChildCount

`func (o *EntitiesUser) GetRecursiveChildCountEqualsChildCount() bool`

GetRecursiveChildCountEqualsChildCount returns the RecursiveChildCountEqualsChildCount field if non-nil, zero value otherwise.

### GetRecursiveChildCountEqualsChildCountOk

`func (o *EntitiesUser) GetRecursiveChildCountEqualsChildCountOk() (*bool, bool)`

GetRecursiveChildCountEqualsChildCountOk returns a tuple with the RecursiveChildCountEqualsChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveChildCountEqualsChildCount

`func (o *EntitiesUser) SetRecursiveChildCountEqualsChildCount(v bool)`

SetRecursiveChildCountEqualsChildCount sets RecursiveChildCountEqualsChildCount field to given value.

### HasRecursiveChildCountEqualsChildCount

`func (o *EntitiesUser) HasRecursiveChildCountEqualsChildCount() bool`

HasRecursiveChildCountEqualsChildCount returns a boolean if a field has been set.

### GetOriginalParsedName

`func (o *EntitiesUser) GetOriginalParsedName() string`

GetOriginalParsedName returns the OriginalParsedName field if non-nil, zero value otherwise.

### GetOriginalParsedNameOk

`func (o *EntitiesUser) GetOriginalParsedNameOk() (*string, bool)`

GetOriginalParsedNameOk returns a tuple with the OriginalParsedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalParsedName

`func (o *EntitiesUser) SetOriginalParsedName(v string)`

SetOriginalParsedName sets OriginalParsedName field to given value.

### HasOriginalParsedName

`func (o *EntitiesUser) HasOriginalParsedName() bool`

HasOriginalParsedName returns a boolean if a field has been set.

### GetIsNameParsedFromFolder

`func (o *EntitiesUser) GetIsNameParsedFromFolder() bool`

GetIsNameParsedFromFolder returns the IsNameParsedFromFolder field if non-nil, zero value otherwise.

### GetIsNameParsedFromFolderOk

`func (o *EntitiesUser) GetIsNameParsedFromFolderOk() (*bool, bool)`

GetIsNameParsedFromFolderOk returns a tuple with the IsNameParsedFromFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNameParsedFromFolder

`func (o *EntitiesUser) SetIsNameParsedFromFolder(v bool)`

SetIsNameParsedFromFolder sets IsNameParsedFromFolder field to given value.

### HasIsNameParsedFromFolder

`func (o *EntitiesUser) HasIsNameParsedFromFolder() bool`

HasIsNameParsedFromFolder returns a boolean if a field has been set.

### GetIdString

`func (o *EntitiesUser) GetIdString() string`

GetIdString returns the IdString field if non-nil, zero value otherwise.

### GetIdStringOk

`func (o *EntitiesUser) GetIdStringOk() (*string, bool)`

GetIdStringOk returns a tuple with the IdString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdString

`func (o *EntitiesUser) SetIdString(v string)`

SetIdString sets IdString field to given value.

### HasIdString

`func (o *EntitiesUser) HasIdString() bool`

HasIdString returns a boolean if a field has been set.

### GetDateCreated

`func (o *EntitiesUser) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *EntitiesUser) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *EntitiesUser) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *EntitiesUser) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetImportedCollections

`func (o *EntitiesUser) GetImportedCollections() []LinkedItemInfo`

GetImportedCollections returns the ImportedCollections field if non-nil, zero value otherwise.

### GetImportedCollectionsOk

`func (o *EntitiesUser) GetImportedCollectionsOk() (*[]LinkedItemInfo, bool)`

GetImportedCollectionsOk returns a tuple with the ImportedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedCollections

`func (o *EntitiesUser) SetImportedCollections(v []LinkedItemInfo)`

SetImportedCollections sets ImportedCollections field to given value.

### HasImportedCollections

`func (o *EntitiesUser) HasImportedCollections() bool`

HasImportedCollections returns a boolean if a field has been set.

### GetResolvedPresentationUniqueKey

`func (o *EntitiesUser) GetResolvedPresentationUniqueKey() string`

GetResolvedPresentationUniqueKey returns the ResolvedPresentationUniqueKey field if non-nil, zero value otherwise.

### GetResolvedPresentationUniqueKeyOk

`func (o *EntitiesUser) GetResolvedPresentationUniqueKeyOk() (*string, bool)`

GetResolvedPresentationUniqueKeyOk returns a tuple with the ResolvedPresentationUniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedPresentationUniqueKey

`func (o *EntitiesUser) SetResolvedPresentationUniqueKey(v string)`

SetResolvedPresentationUniqueKey sets ResolvedPresentationUniqueKey field to given value.

### HasResolvedPresentationUniqueKey

`func (o *EntitiesUser) HasResolvedPresentationUniqueKey() bool`

HasResolvedPresentationUniqueKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


