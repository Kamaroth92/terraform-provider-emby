# \SubtitleOptionsServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEncodingSubtitleoptions**](SubtitleOptionsServiceAPI.md#GetEncodingSubtitleoptions) | **Get** /Encoding/SubtitleOptions | Gets the subtitle options
[**PostEncodingSubtitleoptions**](SubtitleOptionsServiceAPI.md#PostEncodingSubtitleoptions) | **Post** /Encoding/SubtitleOptions | Updates the subtitle options



## GetEncodingSubtitleoptions

> EditObjectContainer GetEncodingSubtitleoptions(ctx).Execute()

Gets the subtitle options



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
	resp, r, err := apiClient.SubtitleOptionsServiceAPI.GetEncodingSubtitleoptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleOptionsServiceAPI.GetEncodingSubtitleoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingSubtitleoptions`: EditObjectContainer
	fmt.Fprintf(os.Stdout, "Response from `SubtitleOptionsServiceAPI.GetEncodingSubtitleoptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingSubtitleoptionsRequest struct via the builder pattern


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


## PostEncodingSubtitleoptions

> PostEncodingSubtitleoptions(ctx).Body(body).Execute()

Updates the subtitle options



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
	r, err := apiClient.SubtitleOptionsServiceAPI.PostEncodingSubtitleoptions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleOptionsServiceAPI.PostEncodingSubtitleoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEncodingSubtitleoptionsRequest struct via the builder pattern


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

