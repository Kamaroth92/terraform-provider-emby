# \GenericUIApiServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUIView**](GenericUIApiServiceAPI.md#GetUIView) | **Get** /UI/View | Gets UI view data
[**PostUICommand**](GenericUIApiServiceAPI.md#PostUICommand) | **Post** /UI/Command | Execute a command in the context of tv setup



## GetUIView

> UIViewInfo GetUIView(ctx).PageId(pageId).ClientLocale(clientLocale).Execute()

Gets UI view data



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
	pageId := "pageId_example" // string | Id of the page controller
	clientLocale := "clientLocale_example" // string | Locale identifier of the client

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericUIApiServiceAPI.GetUIView(context.Background()).PageId(pageId).ClientLocale(clientLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericUIApiServiceAPI.GetUIView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUIView`: UIViewInfo
	fmt.Fprintf(os.Stdout, "Response from `GenericUIApiServiceAPI.GetUIView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUIViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageId** | **string** | Id of the page controller | 
 **clientLocale** | **string** | Locale identifier of the client | 

### Return type

[**UIViewInfo**](UIViewInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUICommand

> UIViewInfo PostUICommand(ctx).RunUICommand(runUICommand).Execute()

Execute a command in the context of tv setup



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
	runUICommand := *openapiclient.NewRunUICommand() // RunUICommand | RunUICommand

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericUIApiServiceAPI.PostUICommand(context.Background()).RunUICommand(runUICommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericUIApiServiceAPI.PostUICommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUICommand`: UIViewInfo
	fmt.Fprintf(os.Stdout, "Response from `GenericUIApiServiceAPI.PostUICommand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUICommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runUICommand** | [**RunUICommand**](RunUICommand.md) | RunUICommand | 

### Return type

[**UIViewInfo**](UIViewInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

