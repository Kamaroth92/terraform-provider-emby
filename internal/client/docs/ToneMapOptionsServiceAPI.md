# \ToneMapOptionsServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEncodingFulltonemapoptions**](ToneMapOptionsServiceAPI.md#GetEncodingFulltonemapoptions) | **Get** /Encoding/FullToneMapOptions | Gets the tone mapping options
[**GetEncodingPublictonemapoptions**](ToneMapOptionsServiceAPI.md#GetEncodingPublictonemapoptions) | **Get** /Encoding/PublicToneMapOptions | Gets the tone mapping options
[**PostEncodingFulltonemapoptions**](ToneMapOptionsServiceAPI.md#PostEncodingFulltonemapoptions) | **Post** /Encoding/FullToneMapOptions | Updates the tone mapping options
[**PostEncodingPublictonemapoptions**](ToneMapOptionsServiceAPI.md#PostEncodingPublictonemapoptions) | **Post** /Encoding/PublicToneMapOptions | Updates the tone mapping options



## GetEncodingFulltonemapoptions

> EditObjectContainer GetEncodingFulltonemapoptions(ctx).Execute()

Gets the tone mapping options



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToneMapOptionsServiceAPI.GetEncodingFulltonemapoptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToneMapOptionsServiceAPI.GetEncodingFulltonemapoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingFulltonemapoptions`: EditObjectContainer
	fmt.Fprintf(os.Stdout, "Response from `ToneMapOptionsServiceAPI.GetEncodingFulltonemapoptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingFulltonemapoptionsRequest struct via the builder pattern


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


## GetEncodingPublictonemapoptions

> EditObjectContainer GetEncodingPublictonemapoptions(ctx).Execute()

Gets the tone mapping options



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToneMapOptionsServiceAPI.GetEncodingPublictonemapoptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToneMapOptionsServiceAPI.GetEncodingPublictonemapoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingPublictonemapoptions`: EditObjectContainer
	fmt.Fprintf(os.Stdout, "Response from `ToneMapOptionsServiceAPI.GetEncodingPublictonemapoptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingPublictonemapoptionsRequest struct via the builder pattern


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


## PostEncodingFulltonemapoptions

> PostEncodingFulltonemapoptions(ctx).Body(body).Execute()

Updates the tone mapping options



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ToneMapOptionsServiceAPI.PostEncodingFulltonemapoptions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToneMapOptionsServiceAPI.PostEncodingFulltonemapoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEncodingFulltonemapoptionsRequest struct via the builder pattern


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


## PostEncodingPublictonemapoptions

> PostEncodingPublictonemapoptions(ctx).Body(body).Execute()

Updates the tone mapping options



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ToneMapOptionsServiceAPI.PostEncodingPublictonemapoptions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToneMapOptionsServiceAPI.PostEncodingPublictonemapoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEncodingPublictonemapoptionsRequest struct via the builder pattern


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

