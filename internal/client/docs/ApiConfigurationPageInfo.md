# ApiConfigurationPageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**EnableInMainMenu** | Pointer to **bool** |  | [optional] 
**EnableInUserMenu** | Pointer to **bool** |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**MenuSection** | Pointer to **string** |  | [optional] 
**MenuIcon** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ConfigurationPageType** | Pointer to [**PluginsConfigurationPageType**](PluginsConfigurationPageType.md) |  | [optional] 
**PluginId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**NavMenuId** | Pointer to **string** |  | [optional] 
**Plugin** | Pointer to [**CommonPluginsIPlugin**](CommonPluginsIPlugin.md) |  | [optional] 
**Translations** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiConfigurationPageInfo

`func NewApiConfigurationPageInfo() *ApiConfigurationPageInfo`

NewApiConfigurationPageInfo instantiates a new ApiConfigurationPageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiConfigurationPageInfoWithDefaults

`func NewApiConfigurationPageInfoWithDefaults() *ApiConfigurationPageInfo`

NewApiConfigurationPageInfoWithDefaults instantiates a new ApiConfigurationPageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiConfigurationPageInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiConfigurationPageInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiConfigurationPageInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiConfigurationPageInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnableInMainMenu

`func (o *ApiConfigurationPageInfo) GetEnableInMainMenu() bool`

GetEnableInMainMenu returns the EnableInMainMenu field if non-nil, zero value otherwise.

### GetEnableInMainMenuOk

`func (o *ApiConfigurationPageInfo) GetEnableInMainMenuOk() (*bool, bool)`

GetEnableInMainMenuOk returns a tuple with the EnableInMainMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInMainMenu

`func (o *ApiConfigurationPageInfo) SetEnableInMainMenu(v bool)`

SetEnableInMainMenu sets EnableInMainMenu field to given value.

### HasEnableInMainMenu

`func (o *ApiConfigurationPageInfo) HasEnableInMainMenu() bool`

HasEnableInMainMenu returns a boolean if a field has been set.

### GetEnableInUserMenu

`func (o *ApiConfigurationPageInfo) GetEnableInUserMenu() bool`

GetEnableInUserMenu returns the EnableInUserMenu field if non-nil, zero value otherwise.

### GetEnableInUserMenuOk

`func (o *ApiConfigurationPageInfo) GetEnableInUserMenuOk() (*bool, bool)`

GetEnableInUserMenuOk returns a tuple with the EnableInUserMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInUserMenu

`func (o *ApiConfigurationPageInfo) SetEnableInUserMenu(v bool)`

SetEnableInUserMenu sets EnableInUserMenu field to given value.

### HasEnableInUserMenu

`func (o *ApiConfigurationPageInfo) HasEnableInUserMenu() bool`

HasEnableInUserMenu returns a boolean if a field has been set.

### GetFeatureId

`func (o *ApiConfigurationPageInfo) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *ApiConfigurationPageInfo) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *ApiConfigurationPageInfo) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *ApiConfigurationPageInfo) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetMenuSection

`func (o *ApiConfigurationPageInfo) GetMenuSection() string`

GetMenuSection returns the MenuSection field if non-nil, zero value otherwise.

### GetMenuSectionOk

`func (o *ApiConfigurationPageInfo) GetMenuSectionOk() (*string, bool)`

GetMenuSectionOk returns a tuple with the MenuSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuSection

`func (o *ApiConfigurationPageInfo) SetMenuSection(v string)`

SetMenuSection sets MenuSection field to given value.

### HasMenuSection

`func (o *ApiConfigurationPageInfo) HasMenuSection() bool`

HasMenuSection returns a boolean if a field has been set.

### GetMenuIcon

`func (o *ApiConfigurationPageInfo) GetMenuIcon() string`

GetMenuIcon returns the MenuIcon field if non-nil, zero value otherwise.

### GetMenuIconOk

`func (o *ApiConfigurationPageInfo) GetMenuIconOk() (*string, bool)`

GetMenuIconOk returns a tuple with the MenuIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuIcon

`func (o *ApiConfigurationPageInfo) SetMenuIcon(v string)`

SetMenuIcon sets MenuIcon field to given value.

### HasMenuIcon

`func (o *ApiConfigurationPageInfo) HasMenuIcon() bool`

HasMenuIcon returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiConfigurationPageInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiConfigurationPageInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiConfigurationPageInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiConfigurationPageInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetConfigurationPageType

`func (o *ApiConfigurationPageInfo) GetConfigurationPageType() PluginsConfigurationPageType`

GetConfigurationPageType returns the ConfigurationPageType field if non-nil, zero value otherwise.

### GetConfigurationPageTypeOk

`func (o *ApiConfigurationPageInfo) GetConfigurationPageTypeOk() (*PluginsConfigurationPageType, bool)`

GetConfigurationPageTypeOk returns a tuple with the ConfigurationPageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPageType

`func (o *ApiConfigurationPageInfo) SetConfigurationPageType(v PluginsConfigurationPageType)`

SetConfigurationPageType sets ConfigurationPageType field to given value.

### HasConfigurationPageType

`func (o *ApiConfigurationPageInfo) HasConfigurationPageType() bool`

HasConfigurationPageType returns a boolean if a field has been set.

### GetPluginId

`func (o *ApiConfigurationPageInfo) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *ApiConfigurationPageInfo) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *ApiConfigurationPageInfo) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *ApiConfigurationPageInfo) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetHref

`func (o *ApiConfigurationPageInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApiConfigurationPageInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApiConfigurationPageInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApiConfigurationPageInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetNavMenuId

`func (o *ApiConfigurationPageInfo) GetNavMenuId() string`

GetNavMenuId returns the NavMenuId field if non-nil, zero value otherwise.

### GetNavMenuIdOk

`func (o *ApiConfigurationPageInfo) GetNavMenuIdOk() (*string, bool)`

GetNavMenuIdOk returns a tuple with the NavMenuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavMenuId

`func (o *ApiConfigurationPageInfo) SetNavMenuId(v string)`

SetNavMenuId sets NavMenuId field to given value.

### HasNavMenuId

`func (o *ApiConfigurationPageInfo) HasNavMenuId() bool`

HasNavMenuId returns a boolean if a field has been set.

### GetPlugin

`func (o *ApiConfigurationPageInfo) GetPlugin() CommonPluginsIPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *ApiConfigurationPageInfo) GetPluginOk() (*CommonPluginsIPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *ApiConfigurationPageInfo) SetPlugin(v CommonPluginsIPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *ApiConfigurationPageInfo) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetTranslations

`func (o *ApiConfigurationPageInfo) GetTranslations() []string`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *ApiConfigurationPageInfo) GetTranslationsOk() (*[]string, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *ApiConfigurationPageInfo) SetTranslations(v []string)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *ApiConfigurationPageInfo) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


