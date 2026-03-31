# \FfmpegOptionsServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEncodingFfmpegoptions**](FfmpegOptionsServiceAPI.md#GetEncodingFfmpegoptions) | **Get** /Encoding/FfmpegOptions | Gets the ffmpeg options
[**PostEncodingFfmpegoptions**](FfmpegOptionsServiceAPI.md#PostEncodingFfmpegoptions) | **Post** /Encoding/FfmpegOptions | Updates the ffmpeg options



## GetEncodingFfmpegoptions

> EditObjectContainer GetEncodingFfmpegoptions(ctx).Execute()

Gets the ffmpeg options



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
	resp, r, err := apiClient.FfmpegOptionsServiceAPI.GetEncodingFfmpegoptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FfmpegOptionsServiceAPI.GetEncodingFfmpegoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingFfmpegoptions`: EditObjectContainer
	fmt.Fprintf(os.Stdout, "Response from `FfmpegOptionsServiceAPI.GetEncodingFfmpegoptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingFfmpegoptionsRequest struct via the builder pattern


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


## PostEncodingFfmpegoptions

> PostEncodingFfmpegoptions(ctx).Body(body).Execute()

Updates the ffmpeg options



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
	r, err := apiClient.FfmpegOptionsServiceAPI.PostEncodingFfmpegoptions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FfmpegOptionsServiceAPI.PostEncodingFfmpegoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEncodingFfmpegoptionsRequest struct via the builder pattern


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

