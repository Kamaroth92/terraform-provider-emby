# \LibraryStructureServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLibraryVirtualfolders**](LibraryStructureServiceAPI.md#DeleteLibraryVirtualfolders) | **Delete** /Library/VirtualFolders | 
[**DeleteLibraryVirtualfoldersPaths**](LibraryStructureServiceAPI.md#DeleteLibraryVirtualfoldersPaths) | **Delete** /Library/VirtualFolders/Paths | 
[**GetLibraryVirtualfoldersQuery**](LibraryStructureServiceAPI.md#GetLibraryVirtualfoldersQuery) | **Get** /Library/VirtualFolders/Query | 
[**PostLibraryVirtualfolders**](LibraryStructureServiceAPI.md#PostLibraryVirtualfolders) | **Post** /Library/VirtualFolders | 
[**PostLibraryVirtualfoldersDelete**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersDelete) | **Post** /Library/VirtualFolders/Delete | 
[**PostLibraryVirtualfoldersLibraryoptions**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersLibraryoptions) | **Post** /Library/VirtualFolders/LibraryOptions | 
[**PostLibraryVirtualfoldersName**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersName) | **Post** /Library/VirtualFolders/Name | 
[**PostLibraryVirtualfoldersPaths**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersPaths) | **Post** /Library/VirtualFolders/Paths | 
[**PostLibraryVirtualfoldersPathsDelete**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersPathsDelete) | **Post** /Library/VirtualFolders/Paths/Delete | 
[**PostLibraryVirtualfoldersPathsUpdate**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersPathsUpdate) | **Post** /Library/VirtualFolders/Paths/Update | 



## DeleteLibraryVirtualfolders

> DeleteLibraryVirtualfolders(ctx).Execute()





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
	r, err := apiClient.LibraryStructureServiceAPI.DeleteLibraryVirtualfolders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.DeleteLibraryVirtualfolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLibraryVirtualfoldersRequest struct via the builder pattern


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


## DeleteLibraryVirtualfoldersPaths

> DeleteLibraryVirtualfoldersPaths(ctx).Execute()





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
	r, err := apiClient.LibraryStructureServiceAPI.DeleteLibraryVirtualfoldersPaths(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.DeleteLibraryVirtualfoldersPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLibraryVirtualfoldersPathsRequest struct via the builder pattern


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


## GetLibraryVirtualfoldersQuery

> QueryResultVirtualFolderInfo GetLibraryVirtualfoldersQuery(ctx).StartIndex(startIndex).Limit(limit).Execute()





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
	resp, r, err := apiClient.LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery(context.Background()).StartIndex(startIndex).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryVirtualfoldersQuery`: QueryResultVirtualFolderInfo
	fmt.Fprintf(os.Stdout, "Response from `LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryVirtualfoldersQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 

### Return type

[**QueryResultVirtualFolderInfo**](QueryResultVirtualFolderInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLibraryVirtualfolders

> PostLibraryVirtualfolders(ctx).LibraryAddVirtualFolder(libraryAddVirtualFolder).Execute()





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
	libraryAddVirtualFolder := *openapiclient.NewLibraryAddVirtualFolder() // LibraryAddVirtualFolder | AddVirtualFolder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfolders(context.Background()).LibraryAddVirtualFolder(libraryAddVirtualFolder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryAddVirtualFolder** | [**LibraryAddVirtualFolder**](LibraryAddVirtualFolder.md) | AddVirtualFolder | 

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


## PostLibraryVirtualfoldersDelete

> PostLibraryVirtualfoldersDelete(ctx).LibraryRemoveVirtualFolder(libraryRemoveVirtualFolder).Execute()





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
	libraryRemoveVirtualFolder := *openapiclient.NewLibraryRemoveVirtualFolder() // LibraryRemoveVirtualFolder | RemoveVirtualFolder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersDelete(context.Background()).LibraryRemoveVirtualFolder(libraryRemoveVirtualFolder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryRemoveVirtualFolder** | [**LibraryRemoveVirtualFolder**](LibraryRemoveVirtualFolder.md) | RemoveVirtualFolder | 

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


## PostLibraryVirtualfoldersLibraryoptions

> PostLibraryVirtualfoldersLibraryoptions(ctx).LibraryUpdateLibraryOptions(libraryUpdateLibraryOptions).Execute()





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
	libraryUpdateLibraryOptions := *openapiclient.NewLibraryUpdateLibraryOptions() // LibraryUpdateLibraryOptions | UpdateLibraryOptions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersLibraryoptions(context.Background()).LibraryUpdateLibraryOptions(libraryUpdateLibraryOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersLibraryoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersLibraryoptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryUpdateLibraryOptions** | [**LibraryUpdateLibraryOptions**](LibraryUpdateLibraryOptions.md) | UpdateLibraryOptions | 

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


## PostLibraryVirtualfoldersName

> PostLibraryVirtualfoldersName(ctx).LibraryRenameVirtualFolder(libraryRenameVirtualFolder).Execute()





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
	libraryRenameVirtualFolder := *openapiclient.NewLibraryRenameVirtualFolder() // LibraryRenameVirtualFolder | RenameVirtualFolder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersName(context.Background()).LibraryRenameVirtualFolder(libraryRenameVirtualFolder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryRenameVirtualFolder** | [**LibraryRenameVirtualFolder**](LibraryRenameVirtualFolder.md) | RenameVirtualFolder | 

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


## PostLibraryVirtualfoldersPaths

> PostLibraryVirtualfoldersPaths(ctx).LibraryAddMediaPath(libraryAddMediaPath).Execute()





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
	libraryAddMediaPath := *openapiclient.NewLibraryAddMediaPath() // LibraryAddMediaPath | AddMediaPath

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPaths(context.Background()).LibraryAddMediaPath(libraryAddMediaPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryAddMediaPath** | [**LibraryAddMediaPath**](LibraryAddMediaPath.md) | AddMediaPath | 

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


## PostLibraryVirtualfoldersPathsDelete

> PostLibraryVirtualfoldersPathsDelete(ctx).LibraryRemoveMediaPath(libraryRemoveMediaPath).Execute()





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
	libraryRemoveMediaPath := *openapiclient.NewLibraryRemoveMediaPath() // LibraryRemoveMediaPath | RemoveMediaPath

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsDelete(context.Background()).LibraryRemoveMediaPath(libraryRemoveMediaPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersPathsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryRemoveMediaPath** | [**LibraryRemoveMediaPath**](LibraryRemoveMediaPath.md) | RemoveMediaPath | 

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


## PostLibraryVirtualfoldersPathsUpdate

> PostLibraryVirtualfoldersPathsUpdate(ctx).LibraryUpdateMediaPath(libraryUpdateMediaPath).Execute()





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
	libraryUpdateMediaPath := *openapiclient.NewLibraryUpdateMediaPath() // LibraryUpdateMediaPath | UpdateMediaPath

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsUpdate(context.Background()).LibraryUpdateMediaPath(libraryUpdateMediaPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersPathsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryUpdateMediaPath** | [**LibraryUpdateMediaPath**](LibraryUpdateMediaPath.md) | UpdateMediaPath | 

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

