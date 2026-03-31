# LevelInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Ordinal** | Pointer to **NullableInt32** |  | [optional] 
**MaxBitRate** | Pointer to [**BitRate**](BitRate.md) |  | [optional] 
**MaxBitRateDisplay** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ResolutionRates** | Pointer to [**[]ResolutionWithRate**](ResolutionWithRate.md) |  | [optional] 
**ResolutionRateStrings** | Pointer to **[]string** |  | [optional] 
**ResolutionRatesDisplay** | Pointer to **string** |  | [optional] 

## Methods

### NewLevelInformation

`func NewLevelInformation() *LevelInformation`

NewLevelInformation instantiates a new LevelInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLevelInformationWithDefaults

`func NewLevelInformationWithDefaults() *LevelInformation`

NewLevelInformationWithDefaults instantiates a new LevelInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortName

`func (o *LevelInformation) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *LevelInformation) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *LevelInformation) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *LevelInformation) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetDescription

`func (o *LevelInformation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LevelInformation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LevelInformation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LevelInformation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrdinal

`func (o *LevelInformation) GetOrdinal() int32`

GetOrdinal returns the Ordinal field if non-nil, zero value otherwise.

### GetOrdinalOk

`func (o *LevelInformation) GetOrdinalOk() (*int32, bool)`

GetOrdinalOk returns a tuple with the Ordinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinal

`func (o *LevelInformation) SetOrdinal(v int32)`

SetOrdinal sets Ordinal field to given value.

### HasOrdinal

`func (o *LevelInformation) HasOrdinal() bool`

HasOrdinal returns a boolean if a field has been set.

### SetOrdinalNil

`func (o *LevelInformation) SetOrdinalNil(b bool)`

 SetOrdinalNil sets the value for Ordinal to be an explicit nil

### UnsetOrdinal
`func (o *LevelInformation) UnsetOrdinal()`

UnsetOrdinal ensures that no value is present for Ordinal, not even an explicit nil
### GetMaxBitRate

`func (o *LevelInformation) GetMaxBitRate() BitRate`

GetMaxBitRate returns the MaxBitRate field if non-nil, zero value otherwise.

### GetMaxBitRateOk

`func (o *LevelInformation) GetMaxBitRateOk() (*BitRate, bool)`

GetMaxBitRateOk returns a tuple with the MaxBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRate

`func (o *LevelInformation) SetMaxBitRate(v BitRate)`

SetMaxBitRate sets MaxBitRate field to given value.

### HasMaxBitRate

`func (o *LevelInformation) HasMaxBitRate() bool`

HasMaxBitRate returns a boolean if a field has been set.

### GetMaxBitRateDisplay

`func (o *LevelInformation) GetMaxBitRateDisplay() string`

GetMaxBitRateDisplay returns the MaxBitRateDisplay field if non-nil, zero value otherwise.

### GetMaxBitRateDisplayOk

`func (o *LevelInformation) GetMaxBitRateDisplayOk() (*string, bool)`

GetMaxBitRateDisplayOk returns a tuple with the MaxBitRateDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRateDisplay

`func (o *LevelInformation) SetMaxBitRateDisplay(v string)`

SetMaxBitRateDisplay sets MaxBitRateDisplay field to given value.

### HasMaxBitRateDisplay

`func (o *LevelInformation) HasMaxBitRateDisplay() bool`

HasMaxBitRateDisplay returns a boolean if a field has been set.

### GetId

`func (o *LevelInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LevelInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LevelInformation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LevelInformation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResolutionRates

`func (o *LevelInformation) GetResolutionRates() []ResolutionWithRate`

GetResolutionRates returns the ResolutionRates field if non-nil, zero value otherwise.

### GetResolutionRatesOk

`func (o *LevelInformation) GetResolutionRatesOk() (*[]ResolutionWithRate, bool)`

GetResolutionRatesOk returns a tuple with the ResolutionRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionRates

`func (o *LevelInformation) SetResolutionRates(v []ResolutionWithRate)`

SetResolutionRates sets ResolutionRates field to given value.

### HasResolutionRates

`func (o *LevelInformation) HasResolutionRates() bool`

HasResolutionRates returns a boolean if a field has been set.

### GetResolutionRateStrings

`func (o *LevelInformation) GetResolutionRateStrings() []string`

GetResolutionRateStrings returns the ResolutionRateStrings field if non-nil, zero value otherwise.

### GetResolutionRateStringsOk

`func (o *LevelInformation) GetResolutionRateStringsOk() (*[]string, bool)`

GetResolutionRateStringsOk returns a tuple with the ResolutionRateStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionRateStrings

`func (o *LevelInformation) SetResolutionRateStrings(v []string)`

SetResolutionRateStrings sets ResolutionRateStrings field to given value.

### HasResolutionRateStrings

`func (o *LevelInformation) HasResolutionRateStrings() bool`

HasResolutionRateStrings returns a boolean if a field has been set.

### GetResolutionRatesDisplay

`func (o *LevelInformation) GetResolutionRatesDisplay() string`

GetResolutionRatesDisplay returns the ResolutionRatesDisplay field if non-nil, zero value otherwise.

### GetResolutionRatesDisplayOk

`func (o *LevelInformation) GetResolutionRatesDisplayOk() (*string, bool)`

GetResolutionRatesDisplayOk returns a tuple with the ResolutionRatesDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionRatesDisplay

`func (o *LevelInformation) SetResolutionRatesDisplay(v string)`

SetResolutionRatesDisplay sets ResolutionRatesDisplay field to given value.

### HasResolutionRatesDisplay

`func (o *LevelInformation) HasResolutionRatesDisplay() bool`

HasResolutionRatesDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


