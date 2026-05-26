# \SystemServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemEndpoint**](SystemServiceAPI.md#GetSystemEndpoint) | **Get** /System/Endpoint | Gets information about the request endpoint
[**GetSystemInfo**](SystemServiceAPI.md#GetSystemInfo) | **Get** /System/Info | Gets information about the server
[**GetSystemInfoPublic**](SystemServiceAPI.md#GetSystemInfoPublic) | **Get** /System/Info/Public | Gets public information about the server
[**GetSystemLogsByName**](SystemServiceAPI.md#GetSystemLogsByName) | **Get** /System/Logs/{Name} | Gets a log file
[**GetSystemLogsByNameLines**](SystemServiceAPI.md#GetSystemLogsByNameLines) | **Get** /System/Logs/{Name}/Lines | Gets a log file
[**GetSystemLogsQuery**](SystemServiceAPI.md#GetSystemLogsQuery) | **Get** /System/Logs/Query | Gets a list of available server log files
[**GetSystemPing**](SystemServiceAPI.md#GetSystemPing) | **Get** /System/Ping | 
[**GetSystemReleasenotes**](SystemServiceAPI.md#GetSystemReleasenotes) | **Get** /System/ReleaseNotes | Gets release notes
[**GetSystemReleasenotesVersions**](SystemServiceAPI.md#GetSystemReleasenotesVersions) | **Get** /System/ReleaseNotes/Versions | Gets release notes
[**GetSystemWakeonlaninfo**](SystemServiceAPI.md#GetSystemWakeonlaninfo) | **Get** /System/WakeOnLanInfo | Gets wake on lan information
[**HeadSystemPing**](SystemServiceAPI.md#HeadSystemPing) | **Head** /System/Ping | 
[**PostSystemPing**](SystemServiceAPI.md#PostSystemPing) | **Post** /System/Ping | 
[**PostSystemRestart**](SystemServiceAPI.md#PostSystemRestart) | **Post** /System/Restart | Restarts the application, if needed
[**PostSystemShutdown**](SystemServiceAPI.md#PostSystemShutdown) | **Post** /System/Shutdown | Shuts down the application



## GetSystemEndpoint

> NetEndPointInfo GetSystemEndpoint(ctx).Execute()

Gets information about the request endpoint



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
	resp, r, err := apiClient.SystemServiceAPI.GetSystemEndpoint(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.GetSystemEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemEndpoint`: NetEndPointInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemServiceAPI.GetSystemEndpoint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemEndpointRequest struct via the builder pattern


### Return type

[**NetEndPointInfo**](NetEndPointInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemInfo

> SystemInfo GetSystemInfo(ctx).Execute()

Gets information about the server



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
	resp, r, err := apiClient.SystemServiceAPI.GetSystemInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.GetSystemInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemInfo`: SystemInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemServiceAPI.GetSystemInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemInfoRequest struct via the builder pattern


### Return type

[**SystemInfo**](SystemInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemInfoPublic

> PublicSystemInfo GetSystemInfoPublic(ctx).Execute()

Gets public information about the server



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
	resp, r, err := apiClient.SystemServiceAPI.GetSystemInfoPublic(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.GetSystemInfoPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemInfoPublic`: PublicSystemInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemServiceAPI.GetSystemInfoPublic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemInfoPublicRequest struct via the builder pattern


### Return type

[**PublicSystemInfo**](PublicSystemInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemLogsByName

> GetSystemLogsByName(ctx, name).Sanitize(sanitize).Execute()

Gets a log file



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
	name := "name_example" // string | The log file name.
	sanitize := true // bool | Return sanitized log (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemServiceAPI.GetSystemLogsByName(context.Background(), name).Sanitize(sanitize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.GetSystemLogsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The log file name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemLogsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sanitize** | **bool** | Return sanitized log | 

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


## GetSystemLogsByNameLines

> QueryResultString GetSystemLogsByNameLines(ctx, name).Execute()

Gets a log file



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
	name := "name_example" // string | The log file name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemServiceAPI.GetSystemLogsByNameLines(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.GetSystemLogsByNameLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemLogsByNameLines`: QueryResultString
	fmt.Fprintf(os.Stdout, "Response from `SystemServiceAPI.GetSystemLogsByNameLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The log file name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemLogsByNameLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryResultString**](QueryResultString.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemLogsQuery

> QueryResultLogFile GetSystemLogsQuery(ctx).StartIndex(startIndex).Limit(limit).Execute()

Gets a list of available server log files



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
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemServiceAPI.GetSystemLogsQuery(context.Background()).StartIndex(startIndex).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.GetSystemLogsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemLogsQuery`: QueryResultLogFile
	fmt.Fprintf(os.Stdout, "Response from `SystemServiceAPI.GetSystemLogsQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemLogsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 

### Return type

[**QueryResultLogFile**](QueryResultLogFile.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemPing

> GetSystemPing(ctx).Execute()





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
	r, err := apiClient.SystemServiceAPI.GetSystemPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.GetSystemPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemPingRequest struct via the builder pattern


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


## GetSystemReleasenotes

> PackageVersionInfo GetSystemReleasenotes(ctx).Execute()

Gets release notes



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
	resp, r, err := apiClient.SystemServiceAPI.GetSystemReleasenotes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.GetSystemReleasenotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemReleasenotes`: PackageVersionInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemServiceAPI.GetSystemReleasenotes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemReleasenotesRequest struct via the builder pattern


### Return type

[**PackageVersionInfo**](PackageVersionInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemReleasenotesVersions

> []PackageVersionInfo GetSystemReleasenotesVersions(ctx).Execute()

Gets release notes



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
	resp, r, err := apiClient.SystemServiceAPI.GetSystemReleasenotesVersions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.GetSystemReleasenotesVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemReleasenotesVersions`: []PackageVersionInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemServiceAPI.GetSystemReleasenotesVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemReleasenotesVersionsRequest struct via the builder pattern


### Return type

[**[]PackageVersionInfo**](PackageVersionInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWakeonlaninfo

> []WakeOnLanInfo GetSystemWakeonlaninfo(ctx).Execute()

Gets wake on lan information



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
	resp, r, err := apiClient.SystemServiceAPI.GetSystemWakeonlaninfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.GetSystemWakeonlaninfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWakeonlaninfo`: []WakeOnLanInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemServiceAPI.GetSystemWakeonlaninfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWakeonlaninfoRequest struct via the builder pattern


### Return type

[**[]WakeOnLanInfo**](WakeOnLanInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadSystemPing

> HeadSystemPing(ctx).Execute()





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
	r, err := apiClient.SystemServiceAPI.HeadSystemPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.HeadSystemPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHeadSystemPingRequest struct via the builder pattern


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


## PostSystemPing

> PostSystemPing(ctx).Execute()





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
	r, err := apiClient.SystemServiceAPI.PostSystemPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.PostSystemPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemPingRequest struct via the builder pattern


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


## PostSystemRestart

> PostSystemRestart(ctx).Execute()

Restarts the application, if needed



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
	r, err := apiClient.SystemServiceAPI.PostSystemRestart(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.PostSystemRestart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemRestartRequest struct via the builder pattern


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


## PostSystemShutdown

> PostSystemShutdown(ctx).Execute()

Shuts down the application



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
	r, err := apiClient.SystemServiceAPI.PostSystemShutdown(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemServiceAPI.PostSystemShutdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemShutdownRequest struct via the builder pattern


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

