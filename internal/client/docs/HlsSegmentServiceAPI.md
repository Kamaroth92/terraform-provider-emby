# \HlsSegmentServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVideosActiveencodings**](HlsSegmentServiceAPI.md#DeleteVideosActiveencodings) | **Delete** /Videos/ActiveEncodings | 
[**PostVideosActiveencodingsDelete**](HlsSegmentServiceAPI.md#PostVideosActiveencodingsDelete) | **Post** /Videos/ActiveEncodings/Delete | 



## DeleteVideosActiveencodings

> DeleteVideosActiveencodings(ctx).DeviceId(deviceId).PlaySessionId(playSessionId).Execute()





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
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed.
	playSessionId := "playSessionId_example" // string | The play session id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HlsSegmentServiceAPI.DeleteVideosActiveencodings(context.Background()).DeviceId(deviceId).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HlsSegmentServiceAPI.DeleteVideosActiveencodings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideosActiveencodingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **playSessionId** | **string** | The play session id | 

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


## PostVideosActiveencodingsDelete

> PostVideosActiveencodingsDelete(ctx).DeviceId(deviceId).PlaySessionId(playSessionId).Execute()





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
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed.
	playSessionId := "playSessionId_example" // string | The play session id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HlsSegmentServiceAPI.PostVideosActiveencodingsDelete(context.Background()).DeviceId(deviceId).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HlsSegmentServiceAPI.PostVideosActiveencodingsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVideosActiveencodingsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **playSessionId** | **string** | The play session id | 

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

