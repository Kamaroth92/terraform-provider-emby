# \ExtractApiEndpointAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChapterApiExtractTheme**](ExtractApiEndpointAPI.md#GetChapterApiExtractTheme) | **Get** /chapter_api/extract_theme | Extract the Theme chromaprint



## GetChapterApiExtractTheme

> map[string]interface{} GetChapterApiExtractTheme(ctx).Id(id).Type_(type_).Execute()

Extract the Theme chromaprint



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
	id := int32(56) // int32 | item id
	type_ := int32(56) // int32 | extraction type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtractApiEndpointAPI.GetChapterApiExtractTheme(context.Background()).Id(id).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtractApiEndpointAPI.GetChapterApiExtractTheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiExtractTheme`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ExtractApiEndpointAPI.GetChapterApiExtractTheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiExtractThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | item id | 
 **type_** | **int32** | extraction type | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

