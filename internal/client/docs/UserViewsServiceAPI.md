# \UserViewsServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsersByUseridViews**](UserViewsServiceAPI.md#GetUsersByUseridViews) | **Get** /Users/{UserId}/Views | 



## GetUsersByUseridViews

> QueryResultBaseItemDto GetUsersByUseridViews(ctx, userId).IncludeExternalContent(includeExternalContent).Execute()





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
	userId := "userId_example" // string | User Id
	includeExternalContent := true // bool | Whether or not to include external views such as channels or live tv

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserViewsServiceAPI.GetUsersByUseridViews(context.Background(), userId).IncludeExternalContent(includeExternalContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserViewsServiceAPI.GetUsersByUseridViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersByUseridViews`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserViewsServiceAPI.GetUsersByUseridViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByUseridViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeExternalContent** | **bool** | Whether or not to include external views such as channels or live tv | 

### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

