# VideoCodecBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodecDeviceInfo** | Pointer to [**CommonInterfacesICodecDeviceInfo**](CommonInterfacesICodecDeviceInfo.md) |  | [optional] 
**CodecKind** | Pointer to [**CodecKinds**](CodecKinds.md) |  | [optional] 
**MediaTypeName** | Pointer to **string** |  | [optional] 
**VideoMediaType** | Pointer to [**VideoMediaTypes**](VideoMediaTypes.md) |  | [optional] 
**MinWidth** | Pointer to **NullableInt32** |  | [optional] 
**MaxWidth** | Pointer to **NullableInt32** |  | [optional] 
**MinHeight** | Pointer to **NullableInt32** |  | [optional] 
**MaxHeight** | Pointer to **NullableInt32** |  | [optional] 
**WidthAlignment** | Pointer to **NullableInt32** |  | [optional] 
**HeightAlignment** | Pointer to **NullableInt32** |  | [optional] 
**MaxBitRate** | Pointer to [**BitRate**](BitRate.md) |  | [optional] 
**SupportedColorFormats** | Pointer to [**[]ColorFormats**](ColorFormats.md) |  | [optional] 
**SupportedColorFormatStrings** | Pointer to **[]string** |  | [optional] 
**ProfileAndLevelInformation** | Pointer to [**[]ProfileLevelInformation**](ProfileLevelInformation.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to [**CodecDirections**](CodecDirections.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FrameworkCodec** | Pointer to **string** |  | [optional] 
**IsHardwareCodec** | Pointer to **bool** |  | [optional] 
**SecondaryFramework** | Pointer to [**SecondaryFrameworks**](SecondaryFrameworks.md) |  | [optional] 
**SecondaryFrameworkCodec** | Pointer to **string** |  | [optional] 
**MaxInstanceCount** | Pointer to **NullableInt32** |  | [optional] 
**IsEnabledByDefault** | Pointer to **bool** |  | [optional] 
**DefaultPriority** | Pointer to **int32** |  | [optional] 

## Methods

### NewVideoCodecBase

`func NewVideoCodecBase() *VideoCodecBase`

NewVideoCodecBase instantiates a new VideoCodecBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoCodecBaseWithDefaults

`func NewVideoCodecBaseWithDefaults() *VideoCodecBase`

NewVideoCodecBaseWithDefaults instantiates a new VideoCodecBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodecDeviceInfo

`func (o *VideoCodecBase) GetCodecDeviceInfo() CommonInterfacesICodecDeviceInfo`

GetCodecDeviceInfo returns the CodecDeviceInfo field if non-nil, zero value otherwise.

### GetCodecDeviceInfoOk

`func (o *VideoCodecBase) GetCodecDeviceInfoOk() (*CommonInterfacesICodecDeviceInfo, bool)`

GetCodecDeviceInfoOk returns a tuple with the CodecDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecDeviceInfo

`func (o *VideoCodecBase) SetCodecDeviceInfo(v CommonInterfacesICodecDeviceInfo)`

SetCodecDeviceInfo sets CodecDeviceInfo field to given value.

### HasCodecDeviceInfo

`func (o *VideoCodecBase) HasCodecDeviceInfo() bool`

HasCodecDeviceInfo returns a boolean if a field has been set.

### GetCodecKind

`func (o *VideoCodecBase) GetCodecKind() CodecKinds`

GetCodecKind returns the CodecKind field if non-nil, zero value otherwise.

### GetCodecKindOk

`func (o *VideoCodecBase) GetCodecKindOk() (*CodecKinds, bool)`

GetCodecKindOk returns a tuple with the CodecKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecKind

`func (o *VideoCodecBase) SetCodecKind(v CodecKinds)`

SetCodecKind sets CodecKind field to given value.

### HasCodecKind

`func (o *VideoCodecBase) HasCodecKind() bool`

HasCodecKind returns a boolean if a field has been set.

### GetMediaTypeName

`func (o *VideoCodecBase) GetMediaTypeName() string`

GetMediaTypeName returns the MediaTypeName field if non-nil, zero value otherwise.

### GetMediaTypeNameOk

`func (o *VideoCodecBase) GetMediaTypeNameOk() (*string, bool)`

GetMediaTypeNameOk returns a tuple with the MediaTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypeName

`func (o *VideoCodecBase) SetMediaTypeName(v string)`

SetMediaTypeName sets MediaTypeName field to given value.

### HasMediaTypeName

`func (o *VideoCodecBase) HasMediaTypeName() bool`

HasMediaTypeName returns a boolean if a field has been set.

### GetVideoMediaType

`func (o *VideoCodecBase) GetVideoMediaType() VideoMediaTypes`

GetVideoMediaType returns the VideoMediaType field if non-nil, zero value otherwise.

### GetVideoMediaTypeOk

`func (o *VideoCodecBase) GetVideoMediaTypeOk() (*VideoMediaTypes, bool)`

GetVideoMediaTypeOk returns a tuple with the VideoMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoMediaType

`func (o *VideoCodecBase) SetVideoMediaType(v VideoMediaTypes)`

SetVideoMediaType sets VideoMediaType field to given value.

### HasVideoMediaType

`func (o *VideoCodecBase) HasVideoMediaType() bool`

HasVideoMediaType returns a boolean if a field has been set.

### GetMinWidth

`func (o *VideoCodecBase) GetMinWidth() int32`

GetMinWidth returns the MinWidth field if non-nil, zero value otherwise.

### GetMinWidthOk

`func (o *VideoCodecBase) GetMinWidthOk() (*int32, bool)`

GetMinWidthOk returns a tuple with the MinWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWidth

`func (o *VideoCodecBase) SetMinWidth(v int32)`

SetMinWidth sets MinWidth field to given value.

### HasMinWidth

`func (o *VideoCodecBase) HasMinWidth() bool`

HasMinWidth returns a boolean if a field has been set.

### SetMinWidthNil

`func (o *VideoCodecBase) SetMinWidthNil(b bool)`

 SetMinWidthNil sets the value for MinWidth to be an explicit nil

### UnsetMinWidth
`func (o *VideoCodecBase) UnsetMinWidth()`

UnsetMinWidth ensures that no value is present for MinWidth, not even an explicit nil
### GetMaxWidth

`func (o *VideoCodecBase) GetMaxWidth() int32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *VideoCodecBase) GetMaxWidthOk() (*int32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *VideoCodecBase) SetMaxWidth(v int32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *VideoCodecBase) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### SetMaxWidthNil

`func (o *VideoCodecBase) SetMaxWidthNil(b bool)`

 SetMaxWidthNil sets the value for MaxWidth to be an explicit nil

### UnsetMaxWidth
`func (o *VideoCodecBase) UnsetMaxWidth()`

UnsetMaxWidth ensures that no value is present for MaxWidth, not even an explicit nil
### GetMinHeight

`func (o *VideoCodecBase) GetMinHeight() int32`

GetMinHeight returns the MinHeight field if non-nil, zero value otherwise.

### GetMinHeightOk

`func (o *VideoCodecBase) GetMinHeightOk() (*int32, bool)`

GetMinHeightOk returns a tuple with the MinHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHeight

`func (o *VideoCodecBase) SetMinHeight(v int32)`

SetMinHeight sets MinHeight field to given value.

### HasMinHeight

`func (o *VideoCodecBase) HasMinHeight() bool`

HasMinHeight returns a boolean if a field has been set.

### SetMinHeightNil

`func (o *VideoCodecBase) SetMinHeightNil(b bool)`

 SetMinHeightNil sets the value for MinHeight to be an explicit nil

### UnsetMinHeight
`func (o *VideoCodecBase) UnsetMinHeight()`

UnsetMinHeight ensures that no value is present for MinHeight, not even an explicit nil
### GetMaxHeight

`func (o *VideoCodecBase) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *VideoCodecBase) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *VideoCodecBase) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *VideoCodecBase) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### SetMaxHeightNil

`func (o *VideoCodecBase) SetMaxHeightNil(b bool)`

 SetMaxHeightNil sets the value for MaxHeight to be an explicit nil

### UnsetMaxHeight
`func (o *VideoCodecBase) UnsetMaxHeight()`

UnsetMaxHeight ensures that no value is present for MaxHeight, not even an explicit nil
### GetWidthAlignment

`func (o *VideoCodecBase) GetWidthAlignment() int32`

GetWidthAlignment returns the WidthAlignment field if non-nil, zero value otherwise.

### GetWidthAlignmentOk

`func (o *VideoCodecBase) GetWidthAlignmentOk() (*int32, bool)`

GetWidthAlignmentOk returns a tuple with the WidthAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthAlignment

`func (o *VideoCodecBase) SetWidthAlignment(v int32)`

SetWidthAlignment sets WidthAlignment field to given value.

### HasWidthAlignment

`func (o *VideoCodecBase) HasWidthAlignment() bool`

HasWidthAlignment returns a boolean if a field has been set.

### SetWidthAlignmentNil

`func (o *VideoCodecBase) SetWidthAlignmentNil(b bool)`

 SetWidthAlignmentNil sets the value for WidthAlignment to be an explicit nil

### UnsetWidthAlignment
`func (o *VideoCodecBase) UnsetWidthAlignment()`

UnsetWidthAlignment ensures that no value is present for WidthAlignment, not even an explicit nil
### GetHeightAlignment

`func (o *VideoCodecBase) GetHeightAlignment() int32`

GetHeightAlignment returns the HeightAlignment field if non-nil, zero value otherwise.

### GetHeightAlignmentOk

`func (o *VideoCodecBase) GetHeightAlignmentOk() (*int32, bool)`

GetHeightAlignmentOk returns a tuple with the HeightAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightAlignment

`func (o *VideoCodecBase) SetHeightAlignment(v int32)`

SetHeightAlignment sets HeightAlignment field to given value.

### HasHeightAlignment

`func (o *VideoCodecBase) HasHeightAlignment() bool`

HasHeightAlignment returns a boolean if a field has been set.

### SetHeightAlignmentNil

`func (o *VideoCodecBase) SetHeightAlignmentNil(b bool)`

 SetHeightAlignmentNil sets the value for HeightAlignment to be an explicit nil

### UnsetHeightAlignment
`func (o *VideoCodecBase) UnsetHeightAlignment()`

UnsetHeightAlignment ensures that no value is present for HeightAlignment, not even an explicit nil
### GetMaxBitRate

`func (o *VideoCodecBase) GetMaxBitRate() BitRate`

GetMaxBitRate returns the MaxBitRate field if non-nil, zero value otherwise.

### GetMaxBitRateOk

`func (o *VideoCodecBase) GetMaxBitRateOk() (*BitRate, bool)`

GetMaxBitRateOk returns a tuple with the MaxBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRate

`func (o *VideoCodecBase) SetMaxBitRate(v BitRate)`

SetMaxBitRate sets MaxBitRate field to given value.

### HasMaxBitRate

`func (o *VideoCodecBase) HasMaxBitRate() bool`

HasMaxBitRate returns a boolean if a field has been set.

### GetSupportedColorFormats

`func (o *VideoCodecBase) GetSupportedColorFormats() []ColorFormats`

GetSupportedColorFormats returns the SupportedColorFormats field if non-nil, zero value otherwise.

### GetSupportedColorFormatsOk

`func (o *VideoCodecBase) GetSupportedColorFormatsOk() (*[]ColorFormats, bool)`

GetSupportedColorFormatsOk returns a tuple with the SupportedColorFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedColorFormats

`func (o *VideoCodecBase) SetSupportedColorFormats(v []ColorFormats)`

SetSupportedColorFormats sets SupportedColorFormats field to given value.

### HasSupportedColorFormats

`func (o *VideoCodecBase) HasSupportedColorFormats() bool`

HasSupportedColorFormats returns a boolean if a field has been set.

### GetSupportedColorFormatStrings

`func (o *VideoCodecBase) GetSupportedColorFormatStrings() []string`

GetSupportedColorFormatStrings returns the SupportedColorFormatStrings field if non-nil, zero value otherwise.

### GetSupportedColorFormatStringsOk

`func (o *VideoCodecBase) GetSupportedColorFormatStringsOk() (*[]string, bool)`

GetSupportedColorFormatStringsOk returns a tuple with the SupportedColorFormatStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedColorFormatStrings

`func (o *VideoCodecBase) SetSupportedColorFormatStrings(v []string)`

SetSupportedColorFormatStrings sets SupportedColorFormatStrings field to given value.

### HasSupportedColorFormatStrings

`func (o *VideoCodecBase) HasSupportedColorFormatStrings() bool`

HasSupportedColorFormatStrings returns a boolean if a field has been set.

### GetProfileAndLevelInformation

`func (o *VideoCodecBase) GetProfileAndLevelInformation() []ProfileLevelInformation`

GetProfileAndLevelInformation returns the ProfileAndLevelInformation field if non-nil, zero value otherwise.

### GetProfileAndLevelInformationOk

`func (o *VideoCodecBase) GetProfileAndLevelInformationOk() (*[]ProfileLevelInformation, bool)`

GetProfileAndLevelInformationOk returns a tuple with the ProfileAndLevelInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAndLevelInformation

`func (o *VideoCodecBase) SetProfileAndLevelInformation(v []ProfileLevelInformation)`

SetProfileAndLevelInformation sets ProfileAndLevelInformation field to given value.

### HasProfileAndLevelInformation

`func (o *VideoCodecBase) HasProfileAndLevelInformation() bool`

HasProfileAndLevelInformation returns a boolean if a field has been set.

### GetId

`func (o *VideoCodecBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoCodecBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoCodecBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VideoCodecBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDirection

`func (o *VideoCodecBase) GetDirection() CodecDirections`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *VideoCodecBase) GetDirectionOk() (*CodecDirections, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *VideoCodecBase) SetDirection(v CodecDirections)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *VideoCodecBase) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetName

`func (o *VideoCodecBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoCodecBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoCodecBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VideoCodecBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VideoCodecBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoCodecBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoCodecBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoCodecBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFrameworkCodec

`func (o *VideoCodecBase) GetFrameworkCodec() string`

GetFrameworkCodec returns the FrameworkCodec field if non-nil, zero value otherwise.

### GetFrameworkCodecOk

`func (o *VideoCodecBase) GetFrameworkCodecOk() (*string, bool)`

GetFrameworkCodecOk returns a tuple with the FrameworkCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkCodec

`func (o *VideoCodecBase) SetFrameworkCodec(v string)`

SetFrameworkCodec sets FrameworkCodec field to given value.

### HasFrameworkCodec

`func (o *VideoCodecBase) HasFrameworkCodec() bool`

HasFrameworkCodec returns a boolean if a field has been set.

### GetIsHardwareCodec

`func (o *VideoCodecBase) GetIsHardwareCodec() bool`

GetIsHardwareCodec returns the IsHardwareCodec field if non-nil, zero value otherwise.

### GetIsHardwareCodecOk

`func (o *VideoCodecBase) GetIsHardwareCodecOk() (*bool, bool)`

GetIsHardwareCodecOk returns a tuple with the IsHardwareCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHardwareCodec

`func (o *VideoCodecBase) SetIsHardwareCodec(v bool)`

SetIsHardwareCodec sets IsHardwareCodec field to given value.

### HasIsHardwareCodec

`func (o *VideoCodecBase) HasIsHardwareCodec() bool`

HasIsHardwareCodec returns a boolean if a field has been set.

### GetSecondaryFramework

`func (o *VideoCodecBase) GetSecondaryFramework() SecondaryFrameworks`

GetSecondaryFramework returns the SecondaryFramework field if non-nil, zero value otherwise.

### GetSecondaryFrameworkOk

`func (o *VideoCodecBase) GetSecondaryFrameworkOk() (*SecondaryFrameworks, bool)`

GetSecondaryFrameworkOk returns a tuple with the SecondaryFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryFramework

`func (o *VideoCodecBase) SetSecondaryFramework(v SecondaryFrameworks)`

SetSecondaryFramework sets SecondaryFramework field to given value.

### HasSecondaryFramework

`func (o *VideoCodecBase) HasSecondaryFramework() bool`

HasSecondaryFramework returns a boolean if a field has been set.

### GetSecondaryFrameworkCodec

`func (o *VideoCodecBase) GetSecondaryFrameworkCodec() string`

GetSecondaryFrameworkCodec returns the SecondaryFrameworkCodec field if non-nil, zero value otherwise.

### GetSecondaryFrameworkCodecOk

`func (o *VideoCodecBase) GetSecondaryFrameworkCodecOk() (*string, bool)`

GetSecondaryFrameworkCodecOk returns a tuple with the SecondaryFrameworkCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryFrameworkCodec

`func (o *VideoCodecBase) SetSecondaryFrameworkCodec(v string)`

SetSecondaryFrameworkCodec sets SecondaryFrameworkCodec field to given value.

### HasSecondaryFrameworkCodec

`func (o *VideoCodecBase) HasSecondaryFrameworkCodec() bool`

HasSecondaryFrameworkCodec returns a boolean if a field has been set.

### GetMaxInstanceCount

`func (o *VideoCodecBase) GetMaxInstanceCount() int32`

GetMaxInstanceCount returns the MaxInstanceCount field if non-nil, zero value otherwise.

### GetMaxInstanceCountOk

`func (o *VideoCodecBase) GetMaxInstanceCountOk() (*int32, bool)`

GetMaxInstanceCountOk returns a tuple with the MaxInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstanceCount

`func (o *VideoCodecBase) SetMaxInstanceCount(v int32)`

SetMaxInstanceCount sets MaxInstanceCount field to given value.

### HasMaxInstanceCount

`func (o *VideoCodecBase) HasMaxInstanceCount() bool`

HasMaxInstanceCount returns a boolean if a field has been set.

### SetMaxInstanceCountNil

`func (o *VideoCodecBase) SetMaxInstanceCountNil(b bool)`

 SetMaxInstanceCountNil sets the value for MaxInstanceCount to be an explicit nil

### UnsetMaxInstanceCount
`func (o *VideoCodecBase) UnsetMaxInstanceCount()`

UnsetMaxInstanceCount ensures that no value is present for MaxInstanceCount, not even an explicit nil
### GetIsEnabledByDefault

`func (o *VideoCodecBase) GetIsEnabledByDefault() bool`

GetIsEnabledByDefault returns the IsEnabledByDefault field if non-nil, zero value otherwise.

### GetIsEnabledByDefaultOk

`func (o *VideoCodecBase) GetIsEnabledByDefaultOk() (*bool, bool)`

GetIsEnabledByDefaultOk returns a tuple with the IsEnabledByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledByDefault

`func (o *VideoCodecBase) SetIsEnabledByDefault(v bool)`

SetIsEnabledByDefault sets IsEnabledByDefault field to given value.

### HasIsEnabledByDefault

`func (o *VideoCodecBase) HasIsEnabledByDefault() bool`

HasIsEnabledByDefault returns a boolean if a field has been set.

### GetDefaultPriority

`func (o *VideoCodecBase) GetDefaultPriority() int32`

GetDefaultPriority returns the DefaultPriority field if non-nil, zero value otherwise.

### GetDefaultPriorityOk

`func (o *VideoCodecBase) GetDefaultPriorityOk() (*int32, bool)`

GetDefaultPriorityOk returns a tuple with the DefaultPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPriority

`func (o *VideoCodecBase) SetDefaultPriority(v int32)`

SetDefaultPriority sets DefaultPriority field to given value.

### HasDefaultPriority

`func (o *VideoCodecBase) HasDefaultPriority() bool`

HasDefaultPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


