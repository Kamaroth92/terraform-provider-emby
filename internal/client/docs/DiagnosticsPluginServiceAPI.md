# \DiagnosticsPluginServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEncodingdiagnosticsDiagnosticoptions**](DiagnosticsPluginServiceAPI.md#GetEncodingdiagnosticsDiagnosticoptions) | **Get** /EncodingDiagnostics/DiagnosticOptions | Gets the EncodingDiagnosticOptions for generic editor
[**PostEncodingdiagnosticsDiagnosticoptions**](DiagnosticsPluginServiceAPI.md#PostEncodingdiagnosticsDiagnosticoptions) | **Post** /EncodingDiagnostics/DiagnosticOptions | Updates EncodingDiagnosticOptions from generic editor



## GetEncodingdiagnosticsDiagnosticoptions

> EditObjectContainer GetEncodingdiagnosticsDiagnosticoptions(ctx).Execute()

Gets the EncodingDiagnosticOptions for generic editor



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
	resp, r, err := apiClient.DiagnosticsPluginServiceAPI.GetEncodingdiagnosticsDiagnosticoptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiagnosticsPluginServiceAPI.GetEncodingdiagnosticsDiagnosticoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingdiagnosticsDiagnosticoptions`: EditObjectContainer
	fmt.Fprintf(os.Stdout, "Response from `DiagnosticsPluginServiceAPI.GetEncodingdiagnosticsDiagnosticoptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingdiagnosticsDiagnosticoptionsRequest struct via the builder pattern


### Return type

[**EditObjectContainer**](EditObjectContainer.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEncodingdiagnosticsDiagnosticoptions

> PostEncodingdiagnosticsDiagnosticoptions(ctx).Body(body).Execute()

Updates EncodingDiagnosticOptions from generic editor



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
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiagnosticsPluginServiceAPI.PostEncodingdiagnosticsDiagnosticoptions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiagnosticsPluginServiceAPI.PostEncodingdiagnosticsDiagnosticoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEncodingdiagnosticsDiagnosticoptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

