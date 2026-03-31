# SessionSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayState** | Pointer to [**PlayerStateInfo**](PlayerStateInfo.md) |  | [optional] 
**AdditionalUsers** | Pointer to [**[]SessionUserInfo**](SessionUserInfo.md) |  | [optional] 
**RemoteEndPoint** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**PlayableMediaTypes** | Pointer to **[]string** |  | [optional] 
**PlaylistItemId** | Pointer to **string** |  | [optional] 
**PlaylistIndex** | Pointer to **int32** |  | [optional] 
**PlaylistLength** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**PartyId** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**UserPrimaryImageTag** | Pointer to **string** |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 
**LastActivityDate** | Pointer to **time.Time** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**NowPlayingItem** | Pointer to [**BaseItemDto**](BaseItemDto.md) |  | [optional] 
**InternalDeviceId** | Pointer to **int64** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**ApplicationVersion** | Pointer to **string** |  | [optional] 
**AppIconUrl** | Pointer to **string** |  | [optional] 
**SupportedCommands** | Pointer to **[]string** |  | [optional] 
**TranscodingInfo** | Pointer to [**TranscodingInfo**](TranscodingInfo.md) |  | [optional] 
**SupportsRemoteControl** | Pointer to **bool** |  | [optional] 

## Methods

### NewSessionSessionInfo

`func NewSessionSessionInfo() *SessionSessionInfo`

NewSessionSessionInfo instantiates a new SessionSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionSessionInfoWithDefaults

`func NewSessionSessionInfoWithDefaults() *SessionSessionInfo`

NewSessionSessionInfoWithDefaults instantiates a new SessionSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayState

`func (o *SessionSessionInfo) GetPlayState() PlayerStateInfo`

GetPlayState returns the PlayState field if non-nil, zero value otherwise.

### GetPlayStateOk

`func (o *SessionSessionInfo) GetPlayStateOk() (*PlayerStateInfo, bool)`

GetPlayStateOk returns a tuple with the PlayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayState

`func (o *SessionSessionInfo) SetPlayState(v PlayerStateInfo)`

SetPlayState sets PlayState field to given value.

### HasPlayState

`func (o *SessionSessionInfo) HasPlayState() bool`

HasPlayState returns a boolean if a field has been set.

### GetAdditionalUsers

`func (o *SessionSessionInfo) GetAdditionalUsers() []SessionUserInfo`

GetAdditionalUsers returns the AdditionalUsers field if non-nil, zero value otherwise.

### GetAdditionalUsersOk

`func (o *SessionSessionInfo) GetAdditionalUsersOk() (*[]SessionUserInfo, bool)`

GetAdditionalUsersOk returns a tuple with the AdditionalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUsers

`func (o *SessionSessionInfo) SetAdditionalUsers(v []SessionUserInfo)`

SetAdditionalUsers sets AdditionalUsers field to given value.

### HasAdditionalUsers

`func (o *SessionSessionInfo) HasAdditionalUsers() bool`

HasAdditionalUsers returns a boolean if a field has been set.

### GetRemoteEndPoint

`func (o *SessionSessionInfo) GetRemoteEndPoint() string`

GetRemoteEndPoint returns the RemoteEndPoint field if non-nil, zero value otherwise.

### GetRemoteEndPointOk

`func (o *SessionSessionInfo) GetRemoteEndPointOk() (*string, bool)`

GetRemoteEndPointOk returns a tuple with the RemoteEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndPoint

`func (o *SessionSessionInfo) SetRemoteEndPoint(v string)`

SetRemoteEndPoint sets RemoteEndPoint field to given value.

### HasRemoteEndPoint

`func (o *SessionSessionInfo) HasRemoteEndPoint() bool`

HasRemoteEndPoint returns a boolean if a field has been set.

### GetProtocol

`func (o *SessionSessionInfo) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SessionSessionInfo) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SessionSessionInfo) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SessionSessionInfo) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPlayableMediaTypes

`func (o *SessionSessionInfo) GetPlayableMediaTypes() []string`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *SessionSessionInfo) GetPlayableMediaTypesOk() (*[]string, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *SessionSessionInfo) SetPlayableMediaTypes(v []string)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *SessionSessionInfo) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *SessionSessionInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *SessionSessionInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *SessionSessionInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *SessionSessionInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### GetPlaylistIndex

`func (o *SessionSessionInfo) GetPlaylistIndex() int32`

GetPlaylistIndex returns the PlaylistIndex field if non-nil, zero value otherwise.

### GetPlaylistIndexOk

`func (o *SessionSessionInfo) GetPlaylistIndexOk() (*int32, bool)`

GetPlaylistIndexOk returns a tuple with the PlaylistIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistIndex

`func (o *SessionSessionInfo) SetPlaylistIndex(v int32)`

SetPlaylistIndex sets PlaylistIndex field to given value.

### HasPlaylistIndex

`func (o *SessionSessionInfo) HasPlaylistIndex() bool`

HasPlaylistIndex returns a boolean if a field has been set.

### GetPlaylistLength

`func (o *SessionSessionInfo) GetPlaylistLength() int32`

GetPlaylistLength returns the PlaylistLength field if non-nil, zero value otherwise.

### GetPlaylistLengthOk

`func (o *SessionSessionInfo) GetPlaylistLengthOk() (*int32, bool)`

GetPlaylistLengthOk returns a tuple with the PlaylistLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistLength

`func (o *SessionSessionInfo) SetPlaylistLength(v int32)`

SetPlaylistLength sets PlaylistLength field to given value.

### HasPlaylistLength

`func (o *SessionSessionInfo) HasPlaylistLength() bool`

HasPlaylistLength returns a boolean if a field has been set.

### GetId

`func (o *SessionSessionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionSessionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionSessionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionSessionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServerId

`func (o *SessionSessionInfo) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *SessionSessionInfo) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *SessionSessionInfo) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *SessionSessionInfo) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetUserId

`func (o *SessionSessionInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SessionSessionInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SessionSessionInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SessionSessionInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPartyId

`func (o *SessionSessionInfo) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *SessionSessionInfo) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *SessionSessionInfo) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *SessionSessionInfo) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetUserName

`func (o *SessionSessionInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SessionSessionInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SessionSessionInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SessionSessionInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserPrimaryImageTag

`func (o *SessionSessionInfo) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *SessionSessionInfo) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *SessionSessionInfo) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *SessionSessionInfo) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### GetClient

`func (o *SessionSessionInfo) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *SessionSessionInfo) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *SessionSessionInfo) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *SessionSessionInfo) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetLastActivityDate

`func (o *SessionSessionInfo) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *SessionSessionInfo) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *SessionSessionInfo) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *SessionSessionInfo) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### GetDeviceName

`func (o *SessionSessionInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SessionSessionInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SessionSessionInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SessionSessionInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *SessionSessionInfo) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SessionSessionInfo) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SessionSessionInfo) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *SessionSessionInfo) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetNowPlayingItem

`func (o *SessionSessionInfo) GetNowPlayingItem() BaseItemDto`

GetNowPlayingItem returns the NowPlayingItem field if non-nil, zero value otherwise.

### GetNowPlayingItemOk

`func (o *SessionSessionInfo) GetNowPlayingItemOk() (*BaseItemDto, bool)`

GetNowPlayingItemOk returns a tuple with the NowPlayingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingItem

`func (o *SessionSessionInfo) SetNowPlayingItem(v BaseItemDto)`

SetNowPlayingItem sets NowPlayingItem field to given value.

### HasNowPlayingItem

`func (o *SessionSessionInfo) HasNowPlayingItem() bool`

HasNowPlayingItem returns a boolean if a field has been set.

### GetInternalDeviceId

`func (o *SessionSessionInfo) GetInternalDeviceId() int64`

GetInternalDeviceId returns the InternalDeviceId field if non-nil, zero value otherwise.

### GetInternalDeviceIdOk

`func (o *SessionSessionInfo) GetInternalDeviceIdOk() (*int64, bool)`

GetInternalDeviceIdOk returns a tuple with the InternalDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDeviceId

`func (o *SessionSessionInfo) SetInternalDeviceId(v int64)`

SetInternalDeviceId sets InternalDeviceId field to given value.

### HasInternalDeviceId

`func (o *SessionSessionInfo) HasInternalDeviceId() bool`

HasInternalDeviceId returns a boolean if a field has been set.

### GetDeviceId

`func (o *SessionSessionInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SessionSessionInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SessionSessionInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SessionSessionInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetApplicationVersion

`func (o *SessionSessionInfo) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *SessionSessionInfo) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *SessionSessionInfo) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *SessionSessionInfo) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### GetAppIconUrl

`func (o *SessionSessionInfo) GetAppIconUrl() string`

GetAppIconUrl returns the AppIconUrl field if non-nil, zero value otherwise.

### GetAppIconUrlOk

`func (o *SessionSessionInfo) GetAppIconUrlOk() (*string, bool)`

GetAppIconUrlOk returns a tuple with the AppIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIconUrl

`func (o *SessionSessionInfo) SetAppIconUrl(v string)`

SetAppIconUrl sets AppIconUrl field to given value.

### HasAppIconUrl

`func (o *SessionSessionInfo) HasAppIconUrl() bool`

HasAppIconUrl returns a boolean if a field has been set.

### GetSupportedCommands

`func (o *SessionSessionInfo) GetSupportedCommands() []string`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *SessionSessionInfo) GetSupportedCommandsOk() (*[]string, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *SessionSessionInfo) SetSupportedCommands(v []string)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *SessionSessionInfo) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### GetTranscodingInfo

`func (o *SessionSessionInfo) GetTranscodingInfo() TranscodingInfo`

GetTranscodingInfo returns the TranscodingInfo field if non-nil, zero value otherwise.

### GetTranscodingInfoOk

`func (o *SessionSessionInfo) GetTranscodingInfoOk() (*TranscodingInfo, bool)`

GetTranscodingInfoOk returns a tuple with the TranscodingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingInfo

`func (o *SessionSessionInfo) SetTranscodingInfo(v TranscodingInfo)`

SetTranscodingInfo sets TranscodingInfo field to given value.

### HasTranscodingInfo

`func (o *SessionSessionInfo) HasTranscodingInfo() bool`

HasTranscodingInfo returns a boolean if a field has been set.

### GetSupportsRemoteControl

`func (o *SessionSessionInfo) GetSupportsRemoteControl() bool`

GetSupportsRemoteControl returns the SupportsRemoteControl field if non-nil, zero value otherwise.

### GetSupportsRemoteControlOk

`func (o *SessionSessionInfo) GetSupportsRemoteControlOk() (*bool, bool)`

GetSupportsRemoteControlOk returns a tuple with the SupportsRemoteControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRemoteControl

`func (o *SessionSessionInfo) SetSupportsRemoteControl(v bool)`

SetSupportsRemoteControl sets SupportsRemoteControl field to given value.

### HasSupportsRemoteControl

`func (o *SessionSessionInfo) HasSupportsRemoteControl() bool`

HasSupportsRemoteControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


