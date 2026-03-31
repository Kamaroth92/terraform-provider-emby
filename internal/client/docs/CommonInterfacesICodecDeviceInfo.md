# CommonInterfacesICodecDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**CommonInterfacesICodecDeviceCapabilities**](CommonInterfacesICodecDeviceCapabilities.md) |  | [optional] 
**Adapter** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Desription** | Pointer to **string** |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 
**DriverVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**ApiVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**VendorId** | Pointer to **int32** |  | [optional] 
**DeviceId** | Pointer to **int32** |  | [optional] 
**DeviceIdentifier** | Pointer to **string** |  | [optional] 
**HardwareContextFramework** | Pointer to [**SecondaryFrameworks**](SecondaryFrameworks.md) |  | [optional] 
**DevPath** | Pointer to **string** |  | [optional] 
**DrmNode** | Pointer to **string** |  | [optional] 
**VendorName** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 

## Methods

### NewCommonInterfacesICodecDeviceInfo

`func NewCommonInterfacesICodecDeviceInfo() *CommonInterfacesICodecDeviceInfo`

NewCommonInterfacesICodecDeviceInfo instantiates a new CommonInterfacesICodecDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonInterfacesICodecDeviceInfoWithDefaults

`func NewCommonInterfacesICodecDeviceInfoWithDefaults() *CommonInterfacesICodecDeviceInfo`

NewCommonInterfacesICodecDeviceInfoWithDefaults instantiates a new CommonInterfacesICodecDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *CommonInterfacesICodecDeviceInfo) GetCapabilities() CommonInterfacesICodecDeviceCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CommonInterfacesICodecDeviceInfo) GetCapabilitiesOk() (*CommonInterfacesICodecDeviceCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CommonInterfacesICodecDeviceInfo) SetCapabilities(v CommonInterfacesICodecDeviceCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CommonInterfacesICodecDeviceInfo) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetAdapter

`func (o *CommonInterfacesICodecDeviceInfo) GetAdapter() int32`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *CommonInterfacesICodecDeviceInfo) GetAdapterOk() (*int32, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *CommonInterfacesICodecDeviceInfo) SetAdapter(v int32)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *CommonInterfacesICodecDeviceInfo) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.

### GetName

`func (o *CommonInterfacesICodecDeviceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonInterfacesICodecDeviceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonInterfacesICodecDeviceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonInterfacesICodecDeviceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDesription

`func (o *CommonInterfacesICodecDeviceInfo) GetDesription() string`

GetDesription returns the Desription field if non-nil, zero value otherwise.

### GetDesriptionOk

`func (o *CommonInterfacesICodecDeviceInfo) GetDesriptionOk() (*string, bool)`

GetDesriptionOk returns a tuple with the Desription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesription

`func (o *CommonInterfacesICodecDeviceInfo) SetDesription(v string)`

SetDesription sets Desription field to given value.

### HasDesription

`func (o *CommonInterfacesICodecDeviceInfo) HasDesription() bool`

HasDesription returns a boolean if a field has been set.

### GetDriver

`func (o *CommonInterfacesICodecDeviceInfo) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *CommonInterfacesICodecDeviceInfo) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *CommonInterfacesICodecDeviceInfo) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *CommonInterfacesICodecDeviceInfo) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetDriverVersion

`func (o *CommonInterfacesICodecDeviceInfo) GetDriverVersion() Version`

GetDriverVersion returns the DriverVersion field if non-nil, zero value otherwise.

### GetDriverVersionOk

`func (o *CommonInterfacesICodecDeviceInfo) GetDriverVersionOk() (*Version, bool)`

GetDriverVersionOk returns a tuple with the DriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverVersion

`func (o *CommonInterfacesICodecDeviceInfo) SetDriverVersion(v Version)`

SetDriverVersion sets DriverVersion field to given value.

### HasDriverVersion

`func (o *CommonInterfacesICodecDeviceInfo) HasDriverVersion() bool`

HasDriverVersion returns a boolean if a field has been set.

### GetApiVersion

`func (o *CommonInterfacesICodecDeviceInfo) GetApiVersion() Version`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CommonInterfacesICodecDeviceInfo) GetApiVersionOk() (*Version, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CommonInterfacesICodecDeviceInfo) SetApiVersion(v Version)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CommonInterfacesICodecDeviceInfo) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetVendorId

`func (o *CommonInterfacesICodecDeviceInfo) GetVendorId() int32`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *CommonInterfacesICodecDeviceInfo) GetVendorIdOk() (*int32, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *CommonInterfacesICodecDeviceInfo) SetVendorId(v int32)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *CommonInterfacesICodecDeviceInfo) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetDeviceId

`func (o *CommonInterfacesICodecDeviceInfo) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *CommonInterfacesICodecDeviceInfo) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *CommonInterfacesICodecDeviceInfo) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *CommonInterfacesICodecDeviceInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceIdentifier

`func (o *CommonInterfacesICodecDeviceInfo) GetDeviceIdentifier() string`

GetDeviceIdentifier returns the DeviceIdentifier field if non-nil, zero value otherwise.

### GetDeviceIdentifierOk

`func (o *CommonInterfacesICodecDeviceInfo) GetDeviceIdentifierOk() (*string, bool)`

GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentifier

`func (o *CommonInterfacesICodecDeviceInfo) SetDeviceIdentifier(v string)`

SetDeviceIdentifier sets DeviceIdentifier field to given value.

### HasDeviceIdentifier

`func (o *CommonInterfacesICodecDeviceInfo) HasDeviceIdentifier() bool`

HasDeviceIdentifier returns a boolean if a field has been set.

### GetHardwareContextFramework

`func (o *CommonInterfacesICodecDeviceInfo) GetHardwareContextFramework() SecondaryFrameworks`

GetHardwareContextFramework returns the HardwareContextFramework field if non-nil, zero value otherwise.

### GetHardwareContextFrameworkOk

`func (o *CommonInterfacesICodecDeviceInfo) GetHardwareContextFrameworkOk() (*SecondaryFrameworks, bool)`

GetHardwareContextFrameworkOk returns a tuple with the HardwareContextFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareContextFramework

`func (o *CommonInterfacesICodecDeviceInfo) SetHardwareContextFramework(v SecondaryFrameworks)`

SetHardwareContextFramework sets HardwareContextFramework field to given value.

### HasHardwareContextFramework

`func (o *CommonInterfacesICodecDeviceInfo) HasHardwareContextFramework() bool`

HasHardwareContextFramework returns a boolean if a field has been set.

### GetDevPath

`func (o *CommonInterfacesICodecDeviceInfo) GetDevPath() string`

GetDevPath returns the DevPath field if non-nil, zero value otherwise.

### GetDevPathOk

`func (o *CommonInterfacesICodecDeviceInfo) GetDevPathOk() (*string, bool)`

GetDevPathOk returns a tuple with the DevPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevPath

`func (o *CommonInterfacesICodecDeviceInfo) SetDevPath(v string)`

SetDevPath sets DevPath field to given value.

### HasDevPath

`func (o *CommonInterfacesICodecDeviceInfo) HasDevPath() bool`

HasDevPath returns a boolean if a field has been set.

### GetDrmNode

`func (o *CommonInterfacesICodecDeviceInfo) GetDrmNode() string`

GetDrmNode returns the DrmNode field if non-nil, zero value otherwise.

### GetDrmNodeOk

`func (o *CommonInterfacesICodecDeviceInfo) GetDrmNodeOk() (*string, bool)`

GetDrmNodeOk returns a tuple with the DrmNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrmNode

`func (o *CommonInterfacesICodecDeviceInfo) SetDrmNode(v string)`

SetDrmNode sets DrmNode field to given value.

### HasDrmNode

`func (o *CommonInterfacesICodecDeviceInfo) HasDrmNode() bool`

HasDrmNode returns a boolean if a field has been set.

### GetVendorName

`func (o *CommonInterfacesICodecDeviceInfo) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *CommonInterfacesICodecDeviceInfo) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *CommonInterfacesICodecDeviceInfo) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *CommonInterfacesICodecDeviceInfo) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetDeviceName

`func (o *CommonInterfacesICodecDeviceInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *CommonInterfacesICodecDeviceInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *CommonInterfacesICodecDeviceInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *CommonInterfacesICodecDeviceInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


