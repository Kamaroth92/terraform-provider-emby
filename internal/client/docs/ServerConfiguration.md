# ServerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableUPnP** | Pointer to **bool** |  | [optional] 
**PublicPort** | Pointer to **int32** |  | [optional] 
**PublicHttpsPort** | Pointer to **int32** |  | [optional] 
**HttpServerPortNumber** | Pointer to **int32** |  | [optional] 
**HttpsPortNumber** | Pointer to **int32** |  | [optional] 
**EnableHttps** | Pointer to **bool** |  | [optional] 
**CertificatePath** | Pointer to **string** |  | [optional] 
**CertificatePassword** | Pointer to **string** |  | [optional] 
**IsPortAuthorized** | Pointer to **bool** |  | [optional] 
**AutoRunWebApp** | Pointer to **bool** |  | [optional] 
**EnableRemoteAccess** | Pointer to **bool** |  | [optional] 
**LogAllQueryTimes** | Pointer to **bool** |  | [optional] 
**DisableOutgoingIPv6** | Pointer to **bool** |  | [optional] 
**EnableCaseSensitiveItemIds** | Pointer to **bool** |  | [optional] 
**MetadataPath** | Pointer to **string** |  | [optional] 
**MetadataNetworkPath** | Pointer to **string** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **string** |  | [optional] 
**MetadataCountryCode** | Pointer to **string** |  | [optional] 
**SortRemoveWords** | Pointer to **[]string** |  | [optional] 
**LibraryMonitorDelaySeconds** | Pointer to **int32** |  | [optional] 
**EnableDashboardResponseCaching** | Pointer to **bool** |  | [optional] 
**DashboardSourcePath** | Pointer to **string** |  | [optional] 
**ImageSavingConvention** | Pointer to [**ImageSavingConvention**](ImageSavingConvention.md) |  | [optional] 
**EnableAutomaticRestart** | Pointer to **bool** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**PreferredDetectedRemoteAddressFamily** | Pointer to [**NetSocketsAddressFamily**](NetSocketsAddressFamily.md) |  | [optional] 
**WanDdns** | Pointer to **string** |  | [optional] 
**UICulture** | Pointer to **string** |  | [optional] 
**RemoteClientBitrateLimit** | Pointer to **int32** |  | [optional] 
**LocalNetworkSubnets** | Pointer to **[]string** |  | [optional] 
**LocalNetworkAddresses** | Pointer to **[]string** |  | [optional] 
**EnableExternalContentInSuggestions** | Pointer to **bool** |  | [optional] 
**RequireHttps** | Pointer to **bool** |  | [optional] 
**IsBehindProxy** | Pointer to **bool** |  | [optional] 
**RemoteIPFilter** | Pointer to **[]string** |  | [optional] 
**IsRemoteIPFilterBlacklist** | Pointer to **bool** |  | [optional] 
**ImageExtractionTimeoutMs** | Pointer to **int32** |  | [optional] 
**PathSubstitutions** | Pointer to [**[]PathSubstitution**](PathSubstitution.md) |  | [optional] 
**UninstalledPlugins** | Pointer to **[]string** |  | [optional] 
**CollapseVideoFolders** | Pointer to **bool** |  | [optional] 
**EnableOriginalTrackTitles** | Pointer to **bool** |  | [optional] 
**VacuumDatabaseOnStartup** | Pointer to **bool** |  | [optional] 
**SimultaneousStreamLimit** | Pointer to **int32** |  | [optional] 
**DatabaseCacheSizeMB** | Pointer to **int32** |  | [optional] 
**EnableSqLiteMmio** | Pointer to **bool** |  | [optional] 
**PlaylistsUpgradedToM3U** | Pointer to **bool** |  | [optional] 
**ImageExtractorUpgraded1** | Pointer to **bool** |  | [optional] 
**EnablePeopleLetterSubFolders** | Pointer to **bool** |  | [optional] 
**OptimizeDatabaseOnShutdown** | Pointer to **bool** |  | [optional] 
**DatabaseAnalysisLimit** | Pointer to **int32** |  | [optional] 
**MaxLibraryDatabaseConnections** | Pointer to **int32** |  | [optional] 
**MaxAuthDbConnections** | Pointer to **int32** |  | [optional] 
**MaxOtherDbConnections** | Pointer to **int32** |  | [optional] 
**DisableAsyncIO** | Pointer to **bool** |  | [optional] 
**MigratedToUserItemShares8** | Pointer to **bool** |  | [optional] 
**MigratedLibraryOptionsToDb** | Pointer to **bool** |  | [optional] 
**AllowLegacyLocalNetworkPassword** | Pointer to **bool** |  | [optional] 
**EnableSavedMetadataForPeople** | Pointer to **bool** |  | [optional] 
**TvChannelsRefreshed** | Pointer to **bool** |  | [optional] 
**ProxyHeaderMode** | Pointer to [**ProxyHeaderMode**](ProxyHeaderMode.md) |  | [optional] 
**IsInMaintenanceMode** | Pointer to **bool** |  | [optional] 
**MaintenanceModeMessage** | Pointer to **string** |  | [optional] 
**EnableDebugLevelLogging** | Pointer to **bool** |  | [optional] 
**RevertDebugLogging** | Pointer to **string** |  | [optional] 
**EnableAutoUpdate** | Pointer to **bool** |  | [optional] 
**LogFileRetentionDays** | Pointer to **int32** |  | [optional] 
**RunAtStartup** | Pointer to **bool** |  | [optional] 
**IsStartupWizardCompleted** | Pointer to **bool** |  | [optional] 
**CachePath** | Pointer to **string** |  | [optional] 

## Methods

### NewServerConfiguration

`func NewServerConfiguration() *ServerConfiguration`

NewServerConfiguration instantiates a new ServerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigurationWithDefaults

`func NewServerConfigurationWithDefaults() *ServerConfiguration`

NewServerConfigurationWithDefaults instantiates a new ServerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableUPnP

`func (o *ServerConfiguration) GetEnableUPnP() bool`

GetEnableUPnP returns the EnableUPnP field if non-nil, zero value otherwise.

### GetEnableUPnPOk

`func (o *ServerConfiguration) GetEnableUPnPOk() (*bool, bool)`

GetEnableUPnPOk returns a tuple with the EnableUPnP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUPnP

`func (o *ServerConfiguration) SetEnableUPnP(v bool)`

SetEnableUPnP sets EnableUPnP field to given value.

### HasEnableUPnP

`func (o *ServerConfiguration) HasEnableUPnP() bool`

HasEnableUPnP returns a boolean if a field has been set.

### GetPublicPort

`func (o *ServerConfiguration) GetPublicPort() int32`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *ServerConfiguration) GetPublicPortOk() (*int32, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *ServerConfiguration) SetPublicPort(v int32)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *ServerConfiguration) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### GetPublicHttpsPort

`func (o *ServerConfiguration) GetPublicHttpsPort() int32`

GetPublicHttpsPort returns the PublicHttpsPort field if non-nil, zero value otherwise.

### GetPublicHttpsPortOk

`func (o *ServerConfiguration) GetPublicHttpsPortOk() (*int32, bool)`

GetPublicHttpsPortOk returns a tuple with the PublicHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHttpsPort

`func (o *ServerConfiguration) SetPublicHttpsPort(v int32)`

SetPublicHttpsPort sets PublicHttpsPort field to given value.

### HasPublicHttpsPort

`func (o *ServerConfiguration) HasPublicHttpsPort() bool`

HasPublicHttpsPort returns a boolean if a field has been set.

### GetHttpServerPortNumber

`func (o *ServerConfiguration) GetHttpServerPortNumber() int32`

GetHttpServerPortNumber returns the HttpServerPortNumber field if non-nil, zero value otherwise.

### GetHttpServerPortNumberOk

`func (o *ServerConfiguration) GetHttpServerPortNumberOk() (*int32, bool)`

GetHttpServerPortNumberOk returns a tuple with the HttpServerPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerPortNumber

`func (o *ServerConfiguration) SetHttpServerPortNumber(v int32)`

SetHttpServerPortNumber sets HttpServerPortNumber field to given value.

### HasHttpServerPortNumber

`func (o *ServerConfiguration) HasHttpServerPortNumber() bool`

HasHttpServerPortNumber returns a boolean if a field has been set.

### GetHttpsPortNumber

`func (o *ServerConfiguration) GetHttpsPortNumber() int32`

GetHttpsPortNumber returns the HttpsPortNumber field if non-nil, zero value otherwise.

### GetHttpsPortNumberOk

`func (o *ServerConfiguration) GetHttpsPortNumberOk() (*int32, bool)`

GetHttpsPortNumberOk returns a tuple with the HttpsPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPortNumber

`func (o *ServerConfiguration) SetHttpsPortNumber(v int32)`

SetHttpsPortNumber sets HttpsPortNumber field to given value.

### HasHttpsPortNumber

`func (o *ServerConfiguration) HasHttpsPortNumber() bool`

HasHttpsPortNumber returns a boolean if a field has been set.

### GetEnableHttps

`func (o *ServerConfiguration) GetEnableHttps() bool`

GetEnableHttps returns the EnableHttps field if non-nil, zero value otherwise.

### GetEnableHttpsOk

`func (o *ServerConfiguration) GetEnableHttpsOk() (*bool, bool)`

GetEnableHttpsOk returns a tuple with the EnableHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttps

`func (o *ServerConfiguration) SetEnableHttps(v bool)`

SetEnableHttps sets EnableHttps field to given value.

### HasEnableHttps

`func (o *ServerConfiguration) HasEnableHttps() bool`

HasEnableHttps returns a boolean if a field has been set.

### GetCertificatePath

`func (o *ServerConfiguration) GetCertificatePath() string`

GetCertificatePath returns the CertificatePath field if non-nil, zero value otherwise.

### GetCertificatePathOk

`func (o *ServerConfiguration) GetCertificatePathOk() (*string, bool)`

GetCertificatePathOk returns a tuple with the CertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePath

`func (o *ServerConfiguration) SetCertificatePath(v string)`

SetCertificatePath sets CertificatePath field to given value.

### HasCertificatePath

`func (o *ServerConfiguration) HasCertificatePath() bool`

HasCertificatePath returns a boolean if a field has been set.

### GetCertificatePassword

`func (o *ServerConfiguration) GetCertificatePassword() string`

GetCertificatePassword returns the CertificatePassword field if non-nil, zero value otherwise.

### GetCertificatePasswordOk

`func (o *ServerConfiguration) GetCertificatePasswordOk() (*string, bool)`

GetCertificatePasswordOk returns a tuple with the CertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePassword

`func (o *ServerConfiguration) SetCertificatePassword(v string)`

SetCertificatePassword sets CertificatePassword field to given value.

### HasCertificatePassword

`func (o *ServerConfiguration) HasCertificatePassword() bool`

HasCertificatePassword returns a boolean if a field has been set.

### GetIsPortAuthorized

`func (o *ServerConfiguration) GetIsPortAuthorized() bool`

GetIsPortAuthorized returns the IsPortAuthorized field if non-nil, zero value otherwise.

### GetIsPortAuthorizedOk

`func (o *ServerConfiguration) GetIsPortAuthorizedOk() (*bool, bool)`

GetIsPortAuthorizedOk returns a tuple with the IsPortAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPortAuthorized

`func (o *ServerConfiguration) SetIsPortAuthorized(v bool)`

SetIsPortAuthorized sets IsPortAuthorized field to given value.

### HasIsPortAuthorized

`func (o *ServerConfiguration) HasIsPortAuthorized() bool`

HasIsPortAuthorized returns a boolean if a field has been set.

### GetAutoRunWebApp

`func (o *ServerConfiguration) GetAutoRunWebApp() bool`

GetAutoRunWebApp returns the AutoRunWebApp field if non-nil, zero value otherwise.

### GetAutoRunWebAppOk

`func (o *ServerConfiguration) GetAutoRunWebAppOk() (*bool, bool)`

GetAutoRunWebAppOk returns a tuple with the AutoRunWebApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRunWebApp

`func (o *ServerConfiguration) SetAutoRunWebApp(v bool)`

SetAutoRunWebApp sets AutoRunWebApp field to given value.

### HasAutoRunWebApp

`func (o *ServerConfiguration) HasAutoRunWebApp() bool`

HasAutoRunWebApp returns a boolean if a field has been set.

### GetEnableRemoteAccess

`func (o *ServerConfiguration) GetEnableRemoteAccess() bool`

GetEnableRemoteAccess returns the EnableRemoteAccess field if non-nil, zero value otherwise.

### GetEnableRemoteAccessOk

`func (o *ServerConfiguration) GetEnableRemoteAccessOk() (*bool, bool)`

GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteAccess

`func (o *ServerConfiguration) SetEnableRemoteAccess(v bool)`

SetEnableRemoteAccess sets EnableRemoteAccess field to given value.

### HasEnableRemoteAccess

`func (o *ServerConfiguration) HasEnableRemoteAccess() bool`

HasEnableRemoteAccess returns a boolean if a field has been set.

### GetLogAllQueryTimes

`func (o *ServerConfiguration) GetLogAllQueryTimes() bool`

GetLogAllQueryTimes returns the LogAllQueryTimes field if non-nil, zero value otherwise.

### GetLogAllQueryTimesOk

`func (o *ServerConfiguration) GetLogAllQueryTimesOk() (*bool, bool)`

GetLogAllQueryTimesOk returns a tuple with the LogAllQueryTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAllQueryTimes

`func (o *ServerConfiguration) SetLogAllQueryTimes(v bool)`

SetLogAllQueryTimes sets LogAllQueryTimes field to given value.

### HasLogAllQueryTimes

`func (o *ServerConfiguration) HasLogAllQueryTimes() bool`

HasLogAllQueryTimes returns a boolean if a field has been set.

### GetDisableOutgoingIPv6

`func (o *ServerConfiguration) GetDisableOutgoingIPv6() bool`

GetDisableOutgoingIPv6 returns the DisableOutgoingIPv6 field if non-nil, zero value otherwise.

### GetDisableOutgoingIPv6Ok

`func (o *ServerConfiguration) GetDisableOutgoingIPv6Ok() (*bool, bool)`

GetDisableOutgoingIPv6Ok returns a tuple with the DisableOutgoingIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOutgoingIPv6

`func (o *ServerConfiguration) SetDisableOutgoingIPv6(v bool)`

SetDisableOutgoingIPv6 sets DisableOutgoingIPv6 field to given value.

### HasDisableOutgoingIPv6

`func (o *ServerConfiguration) HasDisableOutgoingIPv6() bool`

HasDisableOutgoingIPv6 returns a boolean if a field has been set.

### GetEnableCaseSensitiveItemIds

`func (o *ServerConfiguration) GetEnableCaseSensitiveItemIds() bool`

GetEnableCaseSensitiveItemIds returns the EnableCaseSensitiveItemIds field if non-nil, zero value otherwise.

### GetEnableCaseSensitiveItemIdsOk

`func (o *ServerConfiguration) GetEnableCaseSensitiveItemIdsOk() (*bool, bool)`

GetEnableCaseSensitiveItemIdsOk returns a tuple with the EnableCaseSensitiveItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaseSensitiveItemIds

`func (o *ServerConfiguration) SetEnableCaseSensitiveItemIds(v bool)`

SetEnableCaseSensitiveItemIds sets EnableCaseSensitiveItemIds field to given value.

### HasEnableCaseSensitiveItemIds

`func (o *ServerConfiguration) HasEnableCaseSensitiveItemIds() bool`

HasEnableCaseSensitiveItemIds returns a boolean if a field has been set.

### GetMetadataPath

`func (o *ServerConfiguration) GetMetadataPath() string`

GetMetadataPath returns the MetadataPath field if non-nil, zero value otherwise.

### GetMetadataPathOk

`func (o *ServerConfiguration) GetMetadataPathOk() (*string, bool)`

GetMetadataPathOk returns a tuple with the MetadataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataPath

`func (o *ServerConfiguration) SetMetadataPath(v string)`

SetMetadataPath sets MetadataPath field to given value.

### HasMetadataPath

`func (o *ServerConfiguration) HasMetadataPath() bool`

HasMetadataPath returns a boolean if a field has been set.

### GetMetadataNetworkPath

`func (o *ServerConfiguration) GetMetadataNetworkPath() string`

GetMetadataNetworkPath returns the MetadataNetworkPath field if non-nil, zero value otherwise.

### GetMetadataNetworkPathOk

`func (o *ServerConfiguration) GetMetadataNetworkPathOk() (*string, bool)`

GetMetadataNetworkPathOk returns a tuple with the MetadataNetworkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataNetworkPath

`func (o *ServerConfiguration) SetMetadataNetworkPath(v string)`

SetMetadataNetworkPath sets MetadataNetworkPath field to given value.

### HasMetadataNetworkPath

`func (o *ServerConfiguration) HasMetadataNetworkPath() bool`

HasMetadataNetworkPath returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *ServerConfiguration) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *ServerConfiguration) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *ServerConfiguration) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *ServerConfiguration) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### GetMetadataCountryCode

`func (o *ServerConfiguration) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *ServerConfiguration) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *ServerConfiguration) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *ServerConfiguration) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### GetSortRemoveWords

`func (o *ServerConfiguration) GetSortRemoveWords() []string`

GetSortRemoveWords returns the SortRemoveWords field if non-nil, zero value otherwise.

### GetSortRemoveWordsOk

`func (o *ServerConfiguration) GetSortRemoveWordsOk() (*[]string, bool)`

GetSortRemoveWordsOk returns a tuple with the SortRemoveWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortRemoveWords

`func (o *ServerConfiguration) SetSortRemoveWords(v []string)`

SetSortRemoveWords sets SortRemoveWords field to given value.

### HasSortRemoveWords

`func (o *ServerConfiguration) HasSortRemoveWords() bool`

HasSortRemoveWords returns a boolean if a field has been set.

### GetLibraryMonitorDelaySeconds

`func (o *ServerConfiguration) GetLibraryMonitorDelaySeconds() int32`

GetLibraryMonitorDelaySeconds returns the LibraryMonitorDelaySeconds field if non-nil, zero value otherwise.

### GetLibraryMonitorDelaySecondsOk

`func (o *ServerConfiguration) GetLibraryMonitorDelaySecondsOk() (*int32, bool)`

GetLibraryMonitorDelaySecondsOk returns a tuple with the LibraryMonitorDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryMonitorDelaySeconds

`func (o *ServerConfiguration) SetLibraryMonitorDelaySeconds(v int32)`

SetLibraryMonitorDelaySeconds sets LibraryMonitorDelaySeconds field to given value.

### HasLibraryMonitorDelaySeconds

`func (o *ServerConfiguration) HasLibraryMonitorDelaySeconds() bool`

HasLibraryMonitorDelaySeconds returns a boolean if a field has been set.

### GetEnableDashboardResponseCaching

`func (o *ServerConfiguration) GetEnableDashboardResponseCaching() bool`

GetEnableDashboardResponseCaching returns the EnableDashboardResponseCaching field if non-nil, zero value otherwise.

### GetEnableDashboardResponseCachingOk

`func (o *ServerConfiguration) GetEnableDashboardResponseCachingOk() (*bool, bool)`

GetEnableDashboardResponseCachingOk returns a tuple with the EnableDashboardResponseCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDashboardResponseCaching

`func (o *ServerConfiguration) SetEnableDashboardResponseCaching(v bool)`

SetEnableDashboardResponseCaching sets EnableDashboardResponseCaching field to given value.

### HasEnableDashboardResponseCaching

`func (o *ServerConfiguration) HasEnableDashboardResponseCaching() bool`

HasEnableDashboardResponseCaching returns a boolean if a field has been set.

### GetDashboardSourcePath

`func (o *ServerConfiguration) GetDashboardSourcePath() string`

GetDashboardSourcePath returns the DashboardSourcePath field if non-nil, zero value otherwise.

### GetDashboardSourcePathOk

`func (o *ServerConfiguration) GetDashboardSourcePathOk() (*string, bool)`

GetDashboardSourcePathOk returns a tuple with the DashboardSourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardSourcePath

`func (o *ServerConfiguration) SetDashboardSourcePath(v string)`

SetDashboardSourcePath sets DashboardSourcePath field to given value.

### HasDashboardSourcePath

`func (o *ServerConfiguration) HasDashboardSourcePath() bool`

HasDashboardSourcePath returns a boolean if a field has been set.

### GetImageSavingConvention

`func (o *ServerConfiguration) GetImageSavingConvention() ImageSavingConvention`

GetImageSavingConvention returns the ImageSavingConvention field if non-nil, zero value otherwise.

### GetImageSavingConventionOk

`func (o *ServerConfiguration) GetImageSavingConventionOk() (*ImageSavingConvention, bool)`

GetImageSavingConventionOk returns a tuple with the ImageSavingConvention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSavingConvention

`func (o *ServerConfiguration) SetImageSavingConvention(v ImageSavingConvention)`

SetImageSavingConvention sets ImageSavingConvention field to given value.

### HasImageSavingConvention

`func (o *ServerConfiguration) HasImageSavingConvention() bool`

HasImageSavingConvention returns a boolean if a field has been set.

### GetEnableAutomaticRestart

`func (o *ServerConfiguration) GetEnableAutomaticRestart() bool`

GetEnableAutomaticRestart returns the EnableAutomaticRestart field if non-nil, zero value otherwise.

### GetEnableAutomaticRestartOk

`func (o *ServerConfiguration) GetEnableAutomaticRestartOk() (*bool, bool)`

GetEnableAutomaticRestartOk returns a tuple with the EnableAutomaticRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticRestart

`func (o *ServerConfiguration) SetEnableAutomaticRestart(v bool)`

SetEnableAutomaticRestart sets EnableAutomaticRestart field to given value.

### HasEnableAutomaticRestart

`func (o *ServerConfiguration) HasEnableAutomaticRestart() bool`

HasEnableAutomaticRestart returns a boolean if a field has been set.

### GetServerName

`func (o *ServerConfiguration) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ServerConfiguration) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ServerConfiguration) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ServerConfiguration) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetPreferredDetectedRemoteAddressFamily

`func (o *ServerConfiguration) GetPreferredDetectedRemoteAddressFamily() NetSocketsAddressFamily`

GetPreferredDetectedRemoteAddressFamily returns the PreferredDetectedRemoteAddressFamily field if non-nil, zero value otherwise.

### GetPreferredDetectedRemoteAddressFamilyOk

`func (o *ServerConfiguration) GetPreferredDetectedRemoteAddressFamilyOk() (*NetSocketsAddressFamily, bool)`

GetPreferredDetectedRemoteAddressFamilyOk returns a tuple with the PreferredDetectedRemoteAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDetectedRemoteAddressFamily

`func (o *ServerConfiguration) SetPreferredDetectedRemoteAddressFamily(v NetSocketsAddressFamily)`

SetPreferredDetectedRemoteAddressFamily sets PreferredDetectedRemoteAddressFamily field to given value.

### HasPreferredDetectedRemoteAddressFamily

`func (o *ServerConfiguration) HasPreferredDetectedRemoteAddressFamily() bool`

HasPreferredDetectedRemoteAddressFamily returns a boolean if a field has been set.

### GetWanDdns

`func (o *ServerConfiguration) GetWanDdns() string`

GetWanDdns returns the WanDdns field if non-nil, zero value otherwise.

### GetWanDdnsOk

`func (o *ServerConfiguration) GetWanDdnsOk() (*string, bool)`

GetWanDdnsOk returns a tuple with the WanDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanDdns

`func (o *ServerConfiguration) SetWanDdns(v string)`

SetWanDdns sets WanDdns field to given value.

### HasWanDdns

`func (o *ServerConfiguration) HasWanDdns() bool`

HasWanDdns returns a boolean if a field has been set.

### GetUICulture

`func (o *ServerConfiguration) GetUICulture() string`

GetUICulture returns the UICulture field if non-nil, zero value otherwise.

### GetUICultureOk

`func (o *ServerConfiguration) GetUICultureOk() (*string, bool)`

GetUICultureOk returns a tuple with the UICulture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUICulture

`func (o *ServerConfiguration) SetUICulture(v string)`

SetUICulture sets UICulture field to given value.

### HasUICulture

`func (o *ServerConfiguration) HasUICulture() bool`

HasUICulture returns a boolean if a field has been set.

### GetRemoteClientBitrateLimit

`func (o *ServerConfiguration) GetRemoteClientBitrateLimit() int32`

GetRemoteClientBitrateLimit returns the RemoteClientBitrateLimit field if non-nil, zero value otherwise.

### GetRemoteClientBitrateLimitOk

`func (o *ServerConfiguration) GetRemoteClientBitrateLimitOk() (*int32, bool)`

GetRemoteClientBitrateLimitOk returns a tuple with the RemoteClientBitrateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClientBitrateLimit

`func (o *ServerConfiguration) SetRemoteClientBitrateLimit(v int32)`

SetRemoteClientBitrateLimit sets RemoteClientBitrateLimit field to given value.

### HasRemoteClientBitrateLimit

`func (o *ServerConfiguration) HasRemoteClientBitrateLimit() bool`

HasRemoteClientBitrateLimit returns a boolean if a field has been set.

### GetLocalNetworkSubnets

`func (o *ServerConfiguration) GetLocalNetworkSubnets() []string`

GetLocalNetworkSubnets returns the LocalNetworkSubnets field if non-nil, zero value otherwise.

### GetLocalNetworkSubnetsOk

`func (o *ServerConfiguration) GetLocalNetworkSubnetsOk() (*[]string, bool)`

GetLocalNetworkSubnetsOk returns a tuple with the LocalNetworkSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkSubnets

`func (o *ServerConfiguration) SetLocalNetworkSubnets(v []string)`

SetLocalNetworkSubnets sets LocalNetworkSubnets field to given value.

### HasLocalNetworkSubnets

`func (o *ServerConfiguration) HasLocalNetworkSubnets() bool`

HasLocalNetworkSubnets returns a boolean if a field has been set.

### GetLocalNetworkAddresses

`func (o *ServerConfiguration) GetLocalNetworkAddresses() []string`

GetLocalNetworkAddresses returns the LocalNetworkAddresses field if non-nil, zero value otherwise.

### GetLocalNetworkAddressesOk

`func (o *ServerConfiguration) GetLocalNetworkAddressesOk() (*[]string, bool)`

GetLocalNetworkAddressesOk returns a tuple with the LocalNetworkAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkAddresses

`func (o *ServerConfiguration) SetLocalNetworkAddresses(v []string)`

SetLocalNetworkAddresses sets LocalNetworkAddresses field to given value.

### HasLocalNetworkAddresses

`func (o *ServerConfiguration) HasLocalNetworkAddresses() bool`

HasLocalNetworkAddresses returns a boolean if a field has been set.

### GetEnableExternalContentInSuggestions

`func (o *ServerConfiguration) GetEnableExternalContentInSuggestions() bool`

GetEnableExternalContentInSuggestions returns the EnableExternalContentInSuggestions field if non-nil, zero value otherwise.

### GetEnableExternalContentInSuggestionsOk

`func (o *ServerConfiguration) GetEnableExternalContentInSuggestionsOk() (*bool, bool)`

GetEnableExternalContentInSuggestionsOk returns a tuple with the EnableExternalContentInSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalContentInSuggestions

`func (o *ServerConfiguration) SetEnableExternalContentInSuggestions(v bool)`

SetEnableExternalContentInSuggestions sets EnableExternalContentInSuggestions field to given value.

### HasEnableExternalContentInSuggestions

`func (o *ServerConfiguration) HasEnableExternalContentInSuggestions() bool`

HasEnableExternalContentInSuggestions returns a boolean if a field has been set.

### GetRequireHttps

`func (o *ServerConfiguration) GetRequireHttps() bool`

GetRequireHttps returns the RequireHttps field if non-nil, zero value otherwise.

### GetRequireHttpsOk

`func (o *ServerConfiguration) GetRequireHttpsOk() (*bool, bool)`

GetRequireHttpsOk returns a tuple with the RequireHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHttps

`func (o *ServerConfiguration) SetRequireHttps(v bool)`

SetRequireHttps sets RequireHttps field to given value.

### HasRequireHttps

`func (o *ServerConfiguration) HasRequireHttps() bool`

HasRequireHttps returns a boolean if a field has been set.

### GetIsBehindProxy

`func (o *ServerConfiguration) GetIsBehindProxy() bool`

GetIsBehindProxy returns the IsBehindProxy field if non-nil, zero value otherwise.

### GetIsBehindProxyOk

`func (o *ServerConfiguration) GetIsBehindProxyOk() (*bool, bool)`

GetIsBehindProxyOk returns a tuple with the IsBehindProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBehindProxy

`func (o *ServerConfiguration) SetIsBehindProxy(v bool)`

SetIsBehindProxy sets IsBehindProxy field to given value.

### HasIsBehindProxy

`func (o *ServerConfiguration) HasIsBehindProxy() bool`

HasIsBehindProxy returns a boolean if a field has been set.

### GetRemoteIPFilter

`func (o *ServerConfiguration) GetRemoteIPFilter() []string`

GetRemoteIPFilter returns the RemoteIPFilter field if non-nil, zero value otherwise.

### GetRemoteIPFilterOk

`func (o *ServerConfiguration) GetRemoteIPFilterOk() (*[]string, bool)`

GetRemoteIPFilterOk returns a tuple with the RemoteIPFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIPFilter

`func (o *ServerConfiguration) SetRemoteIPFilter(v []string)`

SetRemoteIPFilter sets RemoteIPFilter field to given value.

### HasRemoteIPFilter

`func (o *ServerConfiguration) HasRemoteIPFilter() bool`

HasRemoteIPFilter returns a boolean if a field has been set.

### GetIsRemoteIPFilterBlacklist

`func (o *ServerConfiguration) GetIsRemoteIPFilterBlacklist() bool`

GetIsRemoteIPFilterBlacklist returns the IsRemoteIPFilterBlacklist field if non-nil, zero value otherwise.

### GetIsRemoteIPFilterBlacklistOk

`func (o *ServerConfiguration) GetIsRemoteIPFilterBlacklistOk() (*bool, bool)`

GetIsRemoteIPFilterBlacklistOk returns a tuple with the IsRemoteIPFilterBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteIPFilterBlacklist

`func (o *ServerConfiguration) SetIsRemoteIPFilterBlacklist(v bool)`

SetIsRemoteIPFilterBlacklist sets IsRemoteIPFilterBlacklist field to given value.

### HasIsRemoteIPFilterBlacklist

`func (o *ServerConfiguration) HasIsRemoteIPFilterBlacklist() bool`

HasIsRemoteIPFilterBlacklist returns a boolean if a field has been set.

### GetImageExtractionTimeoutMs

`func (o *ServerConfiguration) GetImageExtractionTimeoutMs() int32`

GetImageExtractionTimeoutMs returns the ImageExtractionTimeoutMs field if non-nil, zero value otherwise.

### GetImageExtractionTimeoutMsOk

`func (o *ServerConfiguration) GetImageExtractionTimeoutMsOk() (*int32, bool)`

GetImageExtractionTimeoutMsOk returns a tuple with the ImageExtractionTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageExtractionTimeoutMs

`func (o *ServerConfiguration) SetImageExtractionTimeoutMs(v int32)`

SetImageExtractionTimeoutMs sets ImageExtractionTimeoutMs field to given value.

### HasImageExtractionTimeoutMs

`func (o *ServerConfiguration) HasImageExtractionTimeoutMs() bool`

HasImageExtractionTimeoutMs returns a boolean if a field has been set.

### GetPathSubstitutions

`func (o *ServerConfiguration) GetPathSubstitutions() []PathSubstitution`

GetPathSubstitutions returns the PathSubstitutions field if non-nil, zero value otherwise.

### GetPathSubstitutionsOk

`func (o *ServerConfiguration) GetPathSubstitutionsOk() (*[]PathSubstitution, bool)`

GetPathSubstitutionsOk returns a tuple with the PathSubstitutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSubstitutions

`func (o *ServerConfiguration) SetPathSubstitutions(v []PathSubstitution)`

SetPathSubstitutions sets PathSubstitutions field to given value.

### HasPathSubstitutions

`func (o *ServerConfiguration) HasPathSubstitutions() bool`

HasPathSubstitutions returns a boolean if a field has been set.

### GetUninstalledPlugins

`func (o *ServerConfiguration) GetUninstalledPlugins() []string`

GetUninstalledPlugins returns the UninstalledPlugins field if non-nil, zero value otherwise.

### GetUninstalledPluginsOk

`func (o *ServerConfiguration) GetUninstalledPluginsOk() (*[]string, bool)`

GetUninstalledPluginsOk returns a tuple with the UninstalledPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninstalledPlugins

`func (o *ServerConfiguration) SetUninstalledPlugins(v []string)`

SetUninstalledPlugins sets UninstalledPlugins field to given value.

### HasUninstalledPlugins

`func (o *ServerConfiguration) HasUninstalledPlugins() bool`

HasUninstalledPlugins returns a boolean if a field has been set.

### GetCollapseVideoFolders

`func (o *ServerConfiguration) GetCollapseVideoFolders() bool`

GetCollapseVideoFolders returns the CollapseVideoFolders field if non-nil, zero value otherwise.

### GetCollapseVideoFoldersOk

`func (o *ServerConfiguration) GetCollapseVideoFoldersOk() (*bool, bool)`

GetCollapseVideoFoldersOk returns a tuple with the CollapseVideoFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseVideoFolders

`func (o *ServerConfiguration) SetCollapseVideoFolders(v bool)`

SetCollapseVideoFolders sets CollapseVideoFolders field to given value.

### HasCollapseVideoFolders

`func (o *ServerConfiguration) HasCollapseVideoFolders() bool`

HasCollapseVideoFolders returns a boolean if a field has been set.

### GetEnableOriginalTrackTitles

`func (o *ServerConfiguration) GetEnableOriginalTrackTitles() bool`

GetEnableOriginalTrackTitles returns the EnableOriginalTrackTitles field if non-nil, zero value otherwise.

### GetEnableOriginalTrackTitlesOk

`func (o *ServerConfiguration) GetEnableOriginalTrackTitlesOk() (*bool, bool)`

GetEnableOriginalTrackTitlesOk returns a tuple with the EnableOriginalTrackTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOriginalTrackTitles

`func (o *ServerConfiguration) SetEnableOriginalTrackTitles(v bool)`

SetEnableOriginalTrackTitles sets EnableOriginalTrackTitles field to given value.

### HasEnableOriginalTrackTitles

`func (o *ServerConfiguration) HasEnableOriginalTrackTitles() bool`

HasEnableOriginalTrackTitles returns a boolean if a field has been set.

### GetVacuumDatabaseOnStartup

`func (o *ServerConfiguration) GetVacuumDatabaseOnStartup() bool`

GetVacuumDatabaseOnStartup returns the VacuumDatabaseOnStartup field if non-nil, zero value otherwise.

### GetVacuumDatabaseOnStartupOk

`func (o *ServerConfiguration) GetVacuumDatabaseOnStartupOk() (*bool, bool)`

GetVacuumDatabaseOnStartupOk returns a tuple with the VacuumDatabaseOnStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacuumDatabaseOnStartup

`func (o *ServerConfiguration) SetVacuumDatabaseOnStartup(v bool)`

SetVacuumDatabaseOnStartup sets VacuumDatabaseOnStartup field to given value.

### HasVacuumDatabaseOnStartup

`func (o *ServerConfiguration) HasVacuumDatabaseOnStartup() bool`

HasVacuumDatabaseOnStartup returns a boolean if a field has been set.

### GetSimultaneousStreamLimit

`func (o *ServerConfiguration) GetSimultaneousStreamLimit() int32`

GetSimultaneousStreamLimit returns the SimultaneousStreamLimit field if non-nil, zero value otherwise.

### GetSimultaneousStreamLimitOk

`func (o *ServerConfiguration) GetSimultaneousStreamLimitOk() (*int32, bool)`

GetSimultaneousStreamLimitOk returns a tuple with the SimultaneousStreamLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimultaneousStreamLimit

`func (o *ServerConfiguration) SetSimultaneousStreamLimit(v int32)`

SetSimultaneousStreamLimit sets SimultaneousStreamLimit field to given value.

### HasSimultaneousStreamLimit

`func (o *ServerConfiguration) HasSimultaneousStreamLimit() bool`

HasSimultaneousStreamLimit returns a boolean if a field has been set.

### GetDatabaseCacheSizeMB

`func (o *ServerConfiguration) GetDatabaseCacheSizeMB() int32`

GetDatabaseCacheSizeMB returns the DatabaseCacheSizeMB field if non-nil, zero value otherwise.

### GetDatabaseCacheSizeMBOk

`func (o *ServerConfiguration) GetDatabaseCacheSizeMBOk() (*int32, bool)`

GetDatabaseCacheSizeMBOk returns a tuple with the DatabaseCacheSizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCacheSizeMB

`func (o *ServerConfiguration) SetDatabaseCacheSizeMB(v int32)`

SetDatabaseCacheSizeMB sets DatabaseCacheSizeMB field to given value.

### HasDatabaseCacheSizeMB

`func (o *ServerConfiguration) HasDatabaseCacheSizeMB() bool`

HasDatabaseCacheSizeMB returns a boolean if a field has been set.

### GetEnableSqLiteMmio

`func (o *ServerConfiguration) GetEnableSqLiteMmio() bool`

GetEnableSqLiteMmio returns the EnableSqLiteMmio field if non-nil, zero value otherwise.

### GetEnableSqLiteMmioOk

`func (o *ServerConfiguration) GetEnableSqLiteMmioOk() (*bool, bool)`

GetEnableSqLiteMmioOk returns a tuple with the EnableSqLiteMmio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSqLiteMmio

`func (o *ServerConfiguration) SetEnableSqLiteMmio(v bool)`

SetEnableSqLiteMmio sets EnableSqLiteMmio field to given value.

### HasEnableSqLiteMmio

`func (o *ServerConfiguration) HasEnableSqLiteMmio() bool`

HasEnableSqLiteMmio returns a boolean if a field has been set.

### GetPlaylistsUpgradedToM3U

`func (o *ServerConfiguration) GetPlaylistsUpgradedToM3U() bool`

GetPlaylistsUpgradedToM3U returns the PlaylistsUpgradedToM3U field if non-nil, zero value otherwise.

### GetPlaylistsUpgradedToM3UOk

`func (o *ServerConfiguration) GetPlaylistsUpgradedToM3UOk() (*bool, bool)`

GetPlaylistsUpgradedToM3UOk returns a tuple with the PlaylistsUpgradedToM3U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistsUpgradedToM3U

`func (o *ServerConfiguration) SetPlaylistsUpgradedToM3U(v bool)`

SetPlaylistsUpgradedToM3U sets PlaylistsUpgradedToM3U field to given value.

### HasPlaylistsUpgradedToM3U

`func (o *ServerConfiguration) HasPlaylistsUpgradedToM3U() bool`

HasPlaylistsUpgradedToM3U returns a boolean if a field has been set.

### GetImageExtractorUpgraded1

`func (o *ServerConfiguration) GetImageExtractorUpgraded1() bool`

GetImageExtractorUpgraded1 returns the ImageExtractorUpgraded1 field if non-nil, zero value otherwise.

### GetImageExtractorUpgraded1Ok

`func (o *ServerConfiguration) GetImageExtractorUpgraded1Ok() (*bool, bool)`

GetImageExtractorUpgraded1Ok returns a tuple with the ImageExtractorUpgraded1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageExtractorUpgraded1

`func (o *ServerConfiguration) SetImageExtractorUpgraded1(v bool)`

SetImageExtractorUpgraded1 sets ImageExtractorUpgraded1 field to given value.

### HasImageExtractorUpgraded1

`func (o *ServerConfiguration) HasImageExtractorUpgraded1() bool`

HasImageExtractorUpgraded1 returns a boolean if a field has been set.

### GetEnablePeopleLetterSubFolders

`func (o *ServerConfiguration) GetEnablePeopleLetterSubFolders() bool`

GetEnablePeopleLetterSubFolders returns the EnablePeopleLetterSubFolders field if non-nil, zero value otherwise.

### GetEnablePeopleLetterSubFoldersOk

`func (o *ServerConfiguration) GetEnablePeopleLetterSubFoldersOk() (*bool, bool)`

GetEnablePeopleLetterSubFoldersOk returns a tuple with the EnablePeopleLetterSubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePeopleLetterSubFolders

`func (o *ServerConfiguration) SetEnablePeopleLetterSubFolders(v bool)`

SetEnablePeopleLetterSubFolders sets EnablePeopleLetterSubFolders field to given value.

### HasEnablePeopleLetterSubFolders

`func (o *ServerConfiguration) HasEnablePeopleLetterSubFolders() bool`

HasEnablePeopleLetterSubFolders returns a boolean if a field has been set.

### GetOptimizeDatabaseOnShutdown

`func (o *ServerConfiguration) GetOptimizeDatabaseOnShutdown() bool`

GetOptimizeDatabaseOnShutdown returns the OptimizeDatabaseOnShutdown field if non-nil, zero value otherwise.

### GetOptimizeDatabaseOnShutdownOk

`func (o *ServerConfiguration) GetOptimizeDatabaseOnShutdownOk() (*bool, bool)`

GetOptimizeDatabaseOnShutdownOk returns a tuple with the OptimizeDatabaseOnShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeDatabaseOnShutdown

`func (o *ServerConfiguration) SetOptimizeDatabaseOnShutdown(v bool)`

SetOptimizeDatabaseOnShutdown sets OptimizeDatabaseOnShutdown field to given value.

### HasOptimizeDatabaseOnShutdown

`func (o *ServerConfiguration) HasOptimizeDatabaseOnShutdown() bool`

HasOptimizeDatabaseOnShutdown returns a boolean if a field has been set.

### GetDatabaseAnalysisLimit

`func (o *ServerConfiguration) GetDatabaseAnalysisLimit() int32`

GetDatabaseAnalysisLimit returns the DatabaseAnalysisLimit field if non-nil, zero value otherwise.

### GetDatabaseAnalysisLimitOk

`func (o *ServerConfiguration) GetDatabaseAnalysisLimitOk() (*int32, bool)`

GetDatabaseAnalysisLimitOk returns a tuple with the DatabaseAnalysisLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseAnalysisLimit

`func (o *ServerConfiguration) SetDatabaseAnalysisLimit(v int32)`

SetDatabaseAnalysisLimit sets DatabaseAnalysisLimit field to given value.

### HasDatabaseAnalysisLimit

`func (o *ServerConfiguration) HasDatabaseAnalysisLimit() bool`

HasDatabaseAnalysisLimit returns a boolean if a field has been set.

### GetMaxLibraryDatabaseConnections

`func (o *ServerConfiguration) GetMaxLibraryDatabaseConnections() int32`

GetMaxLibraryDatabaseConnections returns the MaxLibraryDatabaseConnections field if non-nil, zero value otherwise.

### GetMaxLibraryDatabaseConnectionsOk

`func (o *ServerConfiguration) GetMaxLibraryDatabaseConnectionsOk() (*int32, bool)`

GetMaxLibraryDatabaseConnectionsOk returns a tuple with the MaxLibraryDatabaseConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLibraryDatabaseConnections

`func (o *ServerConfiguration) SetMaxLibraryDatabaseConnections(v int32)`

SetMaxLibraryDatabaseConnections sets MaxLibraryDatabaseConnections field to given value.

### HasMaxLibraryDatabaseConnections

`func (o *ServerConfiguration) HasMaxLibraryDatabaseConnections() bool`

HasMaxLibraryDatabaseConnections returns a boolean if a field has been set.

### GetMaxAuthDbConnections

`func (o *ServerConfiguration) GetMaxAuthDbConnections() int32`

GetMaxAuthDbConnections returns the MaxAuthDbConnections field if non-nil, zero value otherwise.

### GetMaxAuthDbConnectionsOk

`func (o *ServerConfiguration) GetMaxAuthDbConnectionsOk() (*int32, bool)`

GetMaxAuthDbConnectionsOk returns a tuple with the MaxAuthDbConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAuthDbConnections

`func (o *ServerConfiguration) SetMaxAuthDbConnections(v int32)`

SetMaxAuthDbConnections sets MaxAuthDbConnections field to given value.

### HasMaxAuthDbConnections

`func (o *ServerConfiguration) HasMaxAuthDbConnections() bool`

HasMaxAuthDbConnections returns a boolean if a field has been set.

### GetMaxOtherDbConnections

`func (o *ServerConfiguration) GetMaxOtherDbConnections() int32`

GetMaxOtherDbConnections returns the MaxOtherDbConnections field if non-nil, zero value otherwise.

### GetMaxOtherDbConnectionsOk

`func (o *ServerConfiguration) GetMaxOtherDbConnectionsOk() (*int32, bool)`

GetMaxOtherDbConnectionsOk returns a tuple with the MaxOtherDbConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOtherDbConnections

`func (o *ServerConfiguration) SetMaxOtherDbConnections(v int32)`

SetMaxOtherDbConnections sets MaxOtherDbConnections field to given value.

### HasMaxOtherDbConnections

`func (o *ServerConfiguration) HasMaxOtherDbConnections() bool`

HasMaxOtherDbConnections returns a boolean if a field has been set.

### GetDisableAsyncIO

`func (o *ServerConfiguration) GetDisableAsyncIO() bool`

GetDisableAsyncIO returns the DisableAsyncIO field if non-nil, zero value otherwise.

### GetDisableAsyncIOOk

`func (o *ServerConfiguration) GetDisableAsyncIOOk() (*bool, bool)`

GetDisableAsyncIOOk returns a tuple with the DisableAsyncIO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAsyncIO

`func (o *ServerConfiguration) SetDisableAsyncIO(v bool)`

SetDisableAsyncIO sets DisableAsyncIO field to given value.

### HasDisableAsyncIO

`func (o *ServerConfiguration) HasDisableAsyncIO() bool`

HasDisableAsyncIO returns a boolean if a field has been set.

### GetMigratedToUserItemShares8

`func (o *ServerConfiguration) GetMigratedToUserItemShares8() bool`

GetMigratedToUserItemShares8 returns the MigratedToUserItemShares8 field if non-nil, zero value otherwise.

### GetMigratedToUserItemShares8Ok

`func (o *ServerConfiguration) GetMigratedToUserItemShares8Ok() (*bool, bool)`

GetMigratedToUserItemShares8Ok returns a tuple with the MigratedToUserItemShares8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedToUserItemShares8

`func (o *ServerConfiguration) SetMigratedToUserItemShares8(v bool)`

SetMigratedToUserItemShares8 sets MigratedToUserItemShares8 field to given value.

### HasMigratedToUserItemShares8

`func (o *ServerConfiguration) HasMigratedToUserItemShares8() bool`

HasMigratedToUserItemShares8 returns a boolean if a field has been set.

### GetMigratedLibraryOptionsToDb

`func (o *ServerConfiguration) GetMigratedLibraryOptionsToDb() bool`

GetMigratedLibraryOptionsToDb returns the MigratedLibraryOptionsToDb field if non-nil, zero value otherwise.

### GetMigratedLibraryOptionsToDbOk

`func (o *ServerConfiguration) GetMigratedLibraryOptionsToDbOk() (*bool, bool)`

GetMigratedLibraryOptionsToDbOk returns a tuple with the MigratedLibraryOptionsToDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedLibraryOptionsToDb

`func (o *ServerConfiguration) SetMigratedLibraryOptionsToDb(v bool)`

SetMigratedLibraryOptionsToDb sets MigratedLibraryOptionsToDb field to given value.

### HasMigratedLibraryOptionsToDb

`func (o *ServerConfiguration) HasMigratedLibraryOptionsToDb() bool`

HasMigratedLibraryOptionsToDb returns a boolean if a field has been set.

### GetAllowLegacyLocalNetworkPassword

`func (o *ServerConfiguration) GetAllowLegacyLocalNetworkPassword() bool`

GetAllowLegacyLocalNetworkPassword returns the AllowLegacyLocalNetworkPassword field if non-nil, zero value otherwise.

### GetAllowLegacyLocalNetworkPasswordOk

`func (o *ServerConfiguration) GetAllowLegacyLocalNetworkPasswordOk() (*bool, bool)`

GetAllowLegacyLocalNetworkPasswordOk returns a tuple with the AllowLegacyLocalNetworkPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLegacyLocalNetworkPassword

`func (o *ServerConfiguration) SetAllowLegacyLocalNetworkPassword(v bool)`

SetAllowLegacyLocalNetworkPassword sets AllowLegacyLocalNetworkPassword field to given value.

### HasAllowLegacyLocalNetworkPassword

`func (o *ServerConfiguration) HasAllowLegacyLocalNetworkPassword() bool`

HasAllowLegacyLocalNetworkPassword returns a boolean if a field has been set.

### GetEnableSavedMetadataForPeople

`func (o *ServerConfiguration) GetEnableSavedMetadataForPeople() bool`

GetEnableSavedMetadataForPeople returns the EnableSavedMetadataForPeople field if non-nil, zero value otherwise.

### GetEnableSavedMetadataForPeopleOk

`func (o *ServerConfiguration) GetEnableSavedMetadataForPeopleOk() (*bool, bool)`

GetEnableSavedMetadataForPeopleOk returns a tuple with the EnableSavedMetadataForPeople field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSavedMetadataForPeople

`func (o *ServerConfiguration) SetEnableSavedMetadataForPeople(v bool)`

SetEnableSavedMetadataForPeople sets EnableSavedMetadataForPeople field to given value.

### HasEnableSavedMetadataForPeople

`func (o *ServerConfiguration) HasEnableSavedMetadataForPeople() bool`

HasEnableSavedMetadataForPeople returns a boolean if a field has been set.

### GetTvChannelsRefreshed

`func (o *ServerConfiguration) GetTvChannelsRefreshed() bool`

GetTvChannelsRefreshed returns the TvChannelsRefreshed field if non-nil, zero value otherwise.

### GetTvChannelsRefreshedOk

`func (o *ServerConfiguration) GetTvChannelsRefreshedOk() (*bool, bool)`

GetTvChannelsRefreshedOk returns a tuple with the TvChannelsRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvChannelsRefreshed

`func (o *ServerConfiguration) SetTvChannelsRefreshed(v bool)`

SetTvChannelsRefreshed sets TvChannelsRefreshed field to given value.

### HasTvChannelsRefreshed

`func (o *ServerConfiguration) HasTvChannelsRefreshed() bool`

HasTvChannelsRefreshed returns a boolean if a field has been set.

### GetProxyHeaderMode

`func (o *ServerConfiguration) GetProxyHeaderMode() ProxyHeaderMode`

GetProxyHeaderMode returns the ProxyHeaderMode field if non-nil, zero value otherwise.

### GetProxyHeaderModeOk

`func (o *ServerConfiguration) GetProxyHeaderModeOk() (*ProxyHeaderMode, bool)`

GetProxyHeaderModeOk returns a tuple with the ProxyHeaderMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHeaderMode

`func (o *ServerConfiguration) SetProxyHeaderMode(v ProxyHeaderMode)`

SetProxyHeaderMode sets ProxyHeaderMode field to given value.

### HasProxyHeaderMode

`func (o *ServerConfiguration) HasProxyHeaderMode() bool`

HasProxyHeaderMode returns a boolean if a field has been set.

### GetIsInMaintenanceMode

`func (o *ServerConfiguration) GetIsInMaintenanceMode() bool`

GetIsInMaintenanceMode returns the IsInMaintenanceMode field if non-nil, zero value otherwise.

### GetIsInMaintenanceModeOk

`func (o *ServerConfiguration) GetIsInMaintenanceModeOk() (*bool, bool)`

GetIsInMaintenanceModeOk returns a tuple with the IsInMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInMaintenanceMode

`func (o *ServerConfiguration) SetIsInMaintenanceMode(v bool)`

SetIsInMaintenanceMode sets IsInMaintenanceMode field to given value.

### HasIsInMaintenanceMode

`func (o *ServerConfiguration) HasIsInMaintenanceMode() bool`

HasIsInMaintenanceMode returns a boolean if a field has been set.

### GetMaintenanceModeMessage

`func (o *ServerConfiguration) GetMaintenanceModeMessage() string`

GetMaintenanceModeMessage returns the MaintenanceModeMessage field if non-nil, zero value otherwise.

### GetMaintenanceModeMessageOk

`func (o *ServerConfiguration) GetMaintenanceModeMessageOk() (*string, bool)`

GetMaintenanceModeMessageOk returns a tuple with the MaintenanceModeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceModeMessage

`func (o *ServerConfiguration) SetMaintenanceModeMessage(v string)`

SetMaintenanceModeMessage sets MaintenanceModeMessage field to given value.

### HasMaintenanceModeMessage

`func (o *ServerConfiguration) HasMaintenanceModeMessage() bool`

HasMaintenanceModeMessage returns a boolean if a field has been set.

### GetEnableDebugLevelLogging

`func (o *ServerConfiguration) GetEnableDebugLevelLogging() bool`

GetEnableDebugLevelLogging returns the EnableDebugLevelLogging field if non-nil, zero value otherwise.

### GetEnableDebugLevelLoggingOk

`func (o *ServerConfiguration) GetEnableDebugLevelLoggingOk() (*bool, bool)`

GetEnableDebugLevelLoggingOk returns a tuple with the EnableDebugLevelLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebugLevelLogging

`func (o *ServerConfiguration) SetEnableDebugLevelLogging(v bool)`

SetEnableDebugLevelLogging sets EnableDebugLevelLogging field to given value.

### HasEnableDebugLevelLogging

`func (o *ServerConfiguration) HasEnableDebugLevelLogging() bool`

HasEnableDebugLevelLogging returns a boolean if a field has been set.

### GetRevertDebugLogging

`func (o *ServerConfiguration) GetRevertDebugLogging() string`

GetRevertDebugLogging returns the RevertDebugLogging field if non-nil, zero value otherwise.

### GetRevertDebugLoggingOk

`func (o *ServerConfiguration) GetRevertDebugLoggingOk() (*string, bool)`

GetRevertDebugLoggingOk returns a tuple with the RevertDebugLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertDebugLogging

`func (o *ServerConfiguration) SetRevertDebugLogging(v string)`

SetRevertDebugLogging sets RevertDebugLogging field to given value.

### HasRevertDebugLogging

`func (o *ServerConfiguration) HasRevertDebugLogging() bool`

HasRevertDebugLogging returns a boolean if a field has been set.

### GetEnableAutoUpdate

`func (o *ServerConfiguration) GetEnableAutoUpdate() bool`

GetEnableAutoUpdate returns the EnableAutoUpdate field if non-nil, zero value otherwise.

### GetEnableAutoUpdateOk

`func (o *ServerConfiguration) GetEnableAutoUpdateOk() (*bool, bool)`

GetEnableAutoUpdateOk returns a tuple with the EnableAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoUpdate

`func (o *ServerConfiguration) SetEnableAutoUpdate(v bool)`

SetEnableAutoUpdate sets EnableAutoUpdate field to given value.

### HasEnableAutoUpdate

`func (o *ServerConfiguration) HasEnableAutoUpdate() bool`

HasEnableAutoUpdate returns a boolean if a field has been set.

### GetLogFileRetentionDays

`func (o *ServerConfiguration) GetLogFileRetentionDays() int32`

GetLogFileRetentionDays returns the LogFileRetentionDays field if non-nil, zero value otherwise.

### GetLogFileRetentionDaysOk

`func (o *ServerConfiguration) GetLogFileRetentionDaysOk() (*int32, bool)`

GetLogFileRetentionDaysOk returns a tuple with the LogFileRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileRetentionDays

`func (o *ServerConfiguration) SetLogFileRetentionDays(v int32)`

SetLogFileRetentionDays sets LogFileRetentionDays field to given value.

### HasLogFileRetentionDays

`func (o *ServerConfiguration) HasLogFileRetentionDays() bool`

HasLogFileRetentionDays returns a boolean if a field has been set.

### GetRunAtStartup

`func (o *ServerConfiguration) GetRunAtStartup() bool`

GetRunAtStartup returns the RunAtStartup field if non-nil, zero value otherwise.

### GetRunAtStartupOk

`func (o *ServerConfiguration) GetRunAtStartupOk() (*bool, bool)`

GetRunAtStartupOk returns a tuple with the RunAtStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAtStartup

`func (o *ServerConfiguration) SetRunAtStartup(v bool)`

SetRunAtStartup sets RunAtStartup field to given value.

### HasRunAtStartup

`func (o *ServerConfiguration) HasRunAtStartup() bool`

HasRunAtStartup returns a boolean if a field has been set.

### GetIsStartupWizardCompleted

`func (o *ServerConfiguration) GetIsStartupWizardCompleted() bool`

GetIsStartupWizardCompleted returns the IsStartupWizardCompleted field if non-nil, zero value otherwise.

### GetIsStartupWizardCompletedOk

`func (o *ServerConfiguration) GetIsStartupWizardCompletedOk() (*bool, bool)`

GetIsStartupWizardCompletedOk returns a tuple with the IsStartupWizardCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStartupWizardCompleted

`func (o *ServerConfiguration) SetIsStartupWizardCompleted(v bool)`

SetIsStartupWizardCompleted sets IsStartupWizardCompleted field to given value.

### HasIsStartupWizardCompleted

`func (o *ServerConfiguration) HasIsStartupWizardCompleted() bool`

HasIsStartupWizardCompleted returns a boolean if a field has been set.

### GetCachePath

`func (o *ServerConfiguration) GetCachePath() string`

GetCachePath returns the CachePath field if non-nil, zero value otherwise.

### GetCachePathOk

`func (o *ServerConfiguration) GetCachePathOk() (*string, bool)`

GetCachePathOk returns a tuple with the CachePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePath

`func (o *ServerConfiguration) SetCachePath(v string)`

SetCachePath sets CachePath field to given value.

### HasCachePath

`func (o *ServerConfiguration) HasCachePath() bool`

HasCachePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


