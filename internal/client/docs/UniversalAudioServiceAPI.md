# \UniversalAudioServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAudioByIdUniversal**](UniversalAudioServiceAPI.md#GetAudioByIdUniversal) | **Get** /Audio/{Id}/universal | Gets an audio stream
[**GetAudioByIdUniversalByContainer**](UniversalAudioServiceAPI.md#GetAudioByIdUniversalByContainer) | **Get** /Audio/{Id}/universal.{Container} | Gets an audio stream
[**HeadAudioByIdUniversal**](UniversalAudioServiceAPI.md#HeadAudioByIdUniversal) | **Head** /Audio/{Id}/universal | Gets an audio stream
[**HeadAudioByIdUniversalByContainer**](UniversalAudioServiceAPI.md#HeadAudioByIdUniversalByContainer) | **Head** /Audio/{Id}/universal.{Container} | Gets an audio stream



## GetAudioByIdUniversal

> GetAudioByIdUniversal(ctx, id).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()

Gets an audio stream



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Kamaroth92/terraform-provider-emby/client"
)

func main() {
	id := "id_example" // string | Item Id
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1ms = 10000 ticks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniversalAudioServiceAPI.GetAudioByIdUniversal(context.Background(), id).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversalAudioServiceAPI.GetAudioByIdUniversal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioByIdUniversalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1ms &#x3D; 10000 ticks. | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudioByIdUniversalByContainer

> GetAudioByIdUniversalByContainer(ctx, id, container).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()

Gets an audio stream



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Kamaroth92/terraform-provider-emby/client"
)

func main() {
	id := "id_example" // string | Item Id
	container := "container_example" // string | 
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1ms = 10000 ticks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniversalAudioServiceAPI.GetAudioByIdUniversalByContainer(context.Background(), id, container).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversalAudioServiceAPI.GetAudioByIdUniversalByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**container** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioByIdUniversalByContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1ms &#x3D; 10000 ticks. | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadAudioByIdUniversal

> HeadAudioByIdUniversal(ctx, id).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()

Gets an audio stream



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Kamaroth92/terraform-provider-emby/client"
)

func main() {
	id := "id_example" // string | Item Id
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1ms = 10000 ticks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniversalAudioServiceAPI.HeadAudioByIdUniversal(context.Background(), id).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversalAudioServiceAPI.HeadAudioByIdUniversal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadAudioByIdUniversalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1ms &#x3D; 10000 ticks. | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadAudioByIdUniversalByContainer

> HeadAudioByIdUniversalByContainer(ctx, id, container).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()

Gets an audio stream



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Kamaroth92/terraform-provider-emby/client"
)

func main() {
	id := "id_example" // string | Item Id
	container := "container_example" // string | 
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1ms = 10000 ticks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniversalAudioServiceAPI.HeadAudioByIdUniversalByContainer(context.Background(), id, container).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversalAudioServiceAPI.HeadAudioByIdUniversalByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**container** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadAudioByIdUniversalByContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1ms &#x3D; 10000 ticks. | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

