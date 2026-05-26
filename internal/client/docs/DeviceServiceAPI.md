# \DeviceServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDevices**](DeviceServiceAPI.md#DeleteDevices) | **Delete** /Devices | Deletes a device
[**GetDevices**](DeviceServiceAPI.md#GetDevices) | **Get** /Devices | Gets all devices
[**GetDevicesCamerauploads**](DeviceServiceAPI.md#GetDevicesCamerauploads) | **Get** /Devices/CameraUploads | Gets camera upload history for a device
[**GetDevicesInfo**](DeviceServiceAPI.md#GetDevicesInfo) | **Get** /Devices/Info | Gets info for a device
[**GetDevicesOptions**](DeviceServiceAPI.md#GetDevicesOptions) | **Get** /Devices/Options | Gets options for a device
[**PostDevicesCamerauploads**](DeviceServiceAPI.md#PostDevicesCamerauploads) | **Post** /Devices/CameraUploads | Uploads content
[**PostDevicesDelete**](DeviceServiceAPI.md#PostDevicesDelete) | **Post** /Devices/Delete | Deletes a device
[**PostDevicesOptions**](DeviceServiceAPI.md#PostDevicesOptions) | **Post** /Devices/Options | Updates device options



## DeleteDevices

> DeleteDevices(ctx).Id(id).Execute()

Deletes a device



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
	id := "id_example" // string | Device Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceServiceAPI.DeleteDevices(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.DeleteDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id | 

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


## GetDevices

> QueryResultDevicesDeviceInfo GetDevices(ctx).SortOrder(sortOrder).Execute()

Gets all devices



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
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceServiceAPI.GetDevices(context.Background()).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.GetDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevices`: QueryResultDevicesDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `DeviceServiceAPI.GetDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 

### Return type

[**QueryResultDevicesDeviceInfo**](QueryResultDevicesDeviceInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesCamerauploads

> DevicesContentUploadHistory GetDevicesCamerauploads(ctx).Execute()

Gets camera upload history for a device



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceServiceAPI.GetDevicesCamerauploads(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.GetDevicesCamerauploads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesCamerauploads`: DevicesContentUploadHistory
	fmt.Fprintf(os.Stdout, "Response from `DeviceServiceAPI.GetDevicesCamerauploads`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesCamerauploadsRequest struct via the builder pattern


### Return type

[**DevicesContentUploadHistory**](DevicesContentUploadHistory.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesInfo

> DevicesDeviceInfo GetDevicesInfo(ctx).Id(id).Execute()

Gets info for a device



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
	id := "id_example" // string | Device Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceServiceAPI.GetDevicesInfo(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.GetDevicesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesInfo`: DevicesDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `DeviceServiceAPI.GetDevicesInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id | 

### Return type

[**DevicesDeviceInfo**](DevicesDeviceInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesOptions

> DevicesDeviceOptions GetDevicesOptions(ctx).Id(id).Execute()

Gets options for a device



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
	id := "id_example" // string | Device Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceServiceAPI.GetDevicesOptions(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.GetDevicesOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesOptions`: DevicesDeviceOptions
	fmt.Fprintf(os.Stdout, "Response from `DeviceServiceAPI.GetDevicesOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id | 

### Return type

[**DevicesDeviceOptions**](DevicesDeviceOptions.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDevicesCamerauploads

> PostDevicesCamerauploads(ctx).Album(album).Name(name).Id(id).Body(body).Execute()

Uploads content



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
	album := "album_example" // string | Album
	name := "name_example" // string | Name
	id := "id_example" // string | Id
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceServiceAPI.PostDevicesCamerauploads(context.Background()).Album(album).Name(name).Id(id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.PostDevicesCamerauploads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDevicesCamerauploadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **album** | **string** | Album | 
 **name** | **string** | Name | 
 **id** | **string** | Id | 
 **body** | ***os.File** | Binary stream | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDevicesDelete

> PostDevicesDelete(ctx).Id(id).Execute()

Deletes a device



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
	id := "id_example" // string | Device Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceServiceAPI.PostDevicesDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.PostDevicesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDevicesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id | 

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


## PostDevicesOptions

> PostDevicesOptions(ctx).Id(id).DevicesDeviceOptions(devicesDeviceOptions).Execute()

Updates device options



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
	id := "id_example" // string | Device Id
	devicesDeviceOptions := *openapiclient.NewDevicesDeviceOptions() // DevicesDeviceOptions | DeviceOptions: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceServiceAPI.PostDevicesOptions(context.Background()).Id(id).DevicesDeviceOptions(devicesDeviceOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.PostDevicesOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDevicesOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id | 
 **devicesDeviceOptions** | [**DevicesDeviceOptions**](DevicesDeviceOptions.md) | DeviceOptions:  | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

