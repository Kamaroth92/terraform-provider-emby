# \ChapterApiEndpointAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChapterApiGetChapters**](ChapterApiEndpointAPI.md#GetChapterApiGetChapters) | **Get** /chapter_api/get_chapters | Get a list of items for type and filtered
[**GetChapterApiGetItemPath**](ChapterApiEndpointAPI.md#GetChapterApiGetItemPath) | **Get** /chapter_api/get_item_path | Get a list of items for type and filtered
[**GetChapterApiGetItems**](ChapterApiEndpointAPI.md#GetChapterApiGetItems) | **Get** /chapter_api/get_items | Get a list of items for type and filtered
[**GetChapterApiGetSummary**](ChapterApiEndpointAPI.md#GetChapterApiGetSummary) | **Get** /chapter_api/get_summary | Get a list of items for type and filtered
[**GetChapterApiUpdateChapters**](ChapterApiEndpointAPI.md#GetChapterApiUpdateChapters) | **Get** /chapter_api/update_chapters | Updates chapters



## GetChapterApiGetChapters

> map[string]interface{} GetChapterApiGetChapters(ctx).Id(id).Execute()

Get a list of items for type and filtered



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
	id := int32(56) // int32 | item id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChapterApiEndpointAPI.GetChapterApiGetChapters(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChapterApiEndpointAPI.GetChapterApiGetChapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiGetChapters`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChapterApiEndpointAPI.GetChapterApiGetChapters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiGetChaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | item id | 

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


## GetChapterApiGetItemPath

> map[string]interface{} GetChapterApiGetItemPath(ctx).Id(id).Execute()

Get a list of items for type and filtered



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
	id := int32(56) // int32 | item id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChapterApiEndpointAPI.GetChapterApiGetItemPath(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChapterApiEndpointAPI.GetChapterApiGetItemPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiGetItemPath`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChapterApiEndpointAPI.GetChapterApiGetItemPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiGetItemPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | item id | 

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


## GetChapterApiGetItems

> map[string]interface{} GetChapterApiGetItems(ctx).Filter(filter).ItemType(itemType).Parent(parent).Execute()

Get a list of items for type and filtered



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
	filter := "filter_example" // string | filter string (optional)
	itemType := "itemType_example" // string | type of items to return (optional)
	parent := int32(56) // int32 | parent id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChapterApiEndpointAPI.GetChapterApiGetItems(context.Background()).Filter(filter).ItemType(itemType).Parent(parent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChapterApiEndpointAPI.GetChapterApiGetItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiGetItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChapterApiEndpointAPI.GetChapterApiGetItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiGetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter string | 
 **itemType** | **string** | type of items to return | 
 **parent** | **int32** | parent id | 

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


## GetChapterApiGetSummary

> map[string]interface{} GetChapterApiGetSummary(ctx).Type_(type_).Execute()

Get a list of items for type and filtered



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
	type_ := "type__example" // string | summary type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChapterApiEndpointAPI.GetChapterApiGetSummary(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChapterApiEndpointAPI.GetChapterApiGetSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiGetSummary`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChapterApiEndpointAPI.GetChapterApiGetSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiGetSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | summary type | 

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


## GetChapterApiUpdateChapters

> map[string]interface{} GetChapterApiUpdateChapters(ctx).Id(id).IndexList(indexList).Action(action).Name(name).Type_(type_).Time(time).AutoInterval(autoInterval).Execute()

Updates chapters



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
	id := int64(789) // int64 | item id (optional)
	indexList := "indexList_example" // string | list if indexes (optional)
	action := "action_example" // string | action to take (optional)
	name := "name_example" // string | chapter name (optional)
	type_ := "type__example" // string | chapter type (optional)
	time := "time_example" // string | time string of start time hh:mm:ss (optional)
	autoInterval := int32(56) // int32 | auto create interval (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChapterApiEndpointAPI.GetChapterApiUpdateChapters(context.Background()).Id(id).IndexList(indexList).Action(action).Name(name).Type_(type_).Time(time).AutoInterval(autoInterval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChapterApiEndpointAPI.GetChapterApiUpdateChapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiUpdateChapters`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChapterApiEndpointAPI.GetChapterApiUpdateChapters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiUpdateChaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | item id | 
 **indexList** | **string** | list if indexes | 
 **action** | **string** | action to take | 
 **name** | **string** | chapter name | 
 **type_** | **string** | chapter type | 
 **time** | **string** | time string of start time hh:mm:ss | 
 **autoInterval** | **int32** | auto create interval | 

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

