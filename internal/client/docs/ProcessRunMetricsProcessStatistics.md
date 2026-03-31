# ProcessRunMetricsProcessStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentCpu** | Pointer to **float64** |  | [optional] 
**AverageCpu** | Pointer to **float64** |  | [optional] 
**CurrentVirtualMemory** | Pointer to **float64** |  | [optional] 
**CurrentWorkingSet** | Pointer to **float64** |  | [optional] 
**Metrics** | Pointer to [**[]ProcessRunMetricsProcessMetricPoint**](ProcessRunMetricsProcessMetricPoint.md) |  | [optional] 

## Methods

### NewProcessRunMetricsProcessStatistics

`func NewProcessRunMetricsProcessStatistics() *ProcessRunMetricsProcessStatistics`

NewProcessRunMetricsProcessStatistics instantiates a new ProcessRunMetricsProcessStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessRunMetricsProcessStatisticsWithDefaults

`func NewProcessRunMetricsProcessStatisticsWithDefaults() *ProcessRunMetricsProcessStatistics`

NewProcessRunMetricsProcessStatisticsWithDefaults instantiates a new ProcessRunMetricsProcessStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentCpu

`func (o *ProcessRunMetricsProcessStatistics) GetCurrentCpu() float64`

GetCurrentCpu returns the CurrentCpu field if non-nil, zero value otherwise.

### GetCurrentCpuOk

`func (o *ProcessRunMetricsProcessStatistics) GetCurrentCpuOk() (*float64, bool)`

GetCurrentCpuOk returns a tuple with the CurrentCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCpu

`func (o *ProcessRunMetricsProcessStatistics) SetCurrentCpu(v float64)`

SetCurrentCpu sets CurrentCpu field to given value.

### HasCurrentCpu

`func (o *ProcessRunMetricsProcessStatistics) HasCurrentCpu() bool`

HasCurrentCpu returns a boolean if a field has been set.

### GetAverageCpu

`func (o *ProcessRunMetricsProcessStatistics) GetAverageCpu() float64`

GetAverageCpu returns the AverageCpu field if non-nil, zero value otherwise.

### GetAverageCpuOk

`func (o *ProcessRunMetricsProcessStatistics) GetAverageCpuOk() (*float64, bool)`

GetAverageCpuOk returns a tuple with the AverageCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCpu

`func (o *ProcessRunMetricsProcessStatistics) SetAverageCpu(v float64)`

SetAverageCpu sets AverageCpu field to given value.

### HasAverageCpu

`func (o *ProcessRunMetricsProcessStatistics) HasAverageCpu() bool`

HasAverageCpu returns a boolean if a field has been set.

### GetCurrentVirtualMemory

`func (o *ProcessRunMetricsProcessStatistics) GetCurrentVirtualMemory() float64`

GetCurrentVirtualMemory returns the CurrentVirtualMemory field if non-nil, zero value otherwise.

### GetCurrentVirtualMemoryOk

`func (o *ProcessRunMetricsProcessStatistics) GetCurrentVirtualMemoryOk() (*float64, bool)`

GetCurrentVirtualMemoryOk returns a tuple with the CurrentVirtualMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVirtualMemory

`func (o *ProcessRunMetricsProcessStatistics) SetCurrentVirtualMemory(v float64)`

SetCurrentVirtualMemory sets CurrentVirtualMemory field to given value.

### HasCurrentVirtualMemory

`func (o *ProcessRunMetricsProcessStatistics) HasCurrentVirtualMemory() bool`

HasCurrentVirtualMemory returns a boolean if a field has been set.

### GetCurrentWorkingSet

`func (o *ProcessRunMetricsProcessStatistics) GetCurrentWorkingSet() float64`

GetCurrentWorkingSet returns the CurrentWorkingSet field if non-nil, zero value otherwise.

### GetCurrentWorkingSetOk

`func (o *ProcessRunMetricsProcessStatistics) GetCurrentWorkingSetOk() (*float64, bool)`

GetCurrentWorkingSetOk returns a tuple with the CurrentWorkingSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWorkingSet

`func (o *ProcessRunMetricsProcessStatistics) SetCurrentWorkingSet(v float64)`

SetCurrentWorkingSet sets CurrentWorkingSet field to given value.

### HasCurrentWorkingSet

`func (o *ProcessRunMetricsProcessStatistics) HasCurrentWorkingSet() bool`

HasCurrentWorkingSet returns a boolean if a field has been set.

### GetMetrics

`func (o *ProcessRunMetricsProcessStatistics) GetMetrics() []ProcessRunMetricsProcessMetricPoint`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ProcessRunMetricsProcessStatistics) GetMetricsOk() (*[]ProcessRunMetricsProcessMetricPoint, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ProcessRunMetricsProcessStatistics) SetMetrics(v []ProcessRunMetricsProcessMetricPoint)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ProcessRunMetricsProcessStatistics) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


