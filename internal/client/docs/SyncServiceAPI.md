# \SyncServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSyncByTargetidItems**](SyncServiceAPI.md#DeleteSyncByTargetidItems) | **Delete** /Sync/{TargetId}/Items | Cancels items from a sync target
[**DeleteSyncJobitemsById**](SyncServiceAPI.md#DeleteSyncJobitemsById) | **Delete** /Sync/JobItems/{Id} | Cancels a sync job item
[**DeleteSyncJobsById**](SyncServiceAPI.md#DeleteSyncJobsById) | **Delete** /Sync/Jobs/{Id} | Cancels a sync job.
[**GetSyncItemsReady**](SyncServiceAPI.md#GetSyncItemsReady) | **Get** /Sync/Items/Ready | Gets ready to download sync items.
[**GetSyncJobitems**](SyncServiceAPI.md#GetSyncJobitems) | **Get** /Sync/JobItems | Gets sync job items.
[**GetSyncJobitemsByIdAdditionalfiles**](SyncServiceAPI.md#GetSyncJobitemsByIdAdditionalfiles) | **Get** /Sync/JobItems/{Id}/AdditionalFiles | Gets a sync job item file
[**GetSyncJobitemsByIdFile**](SyncServiceAPI.md#GetSyncJobitemsByIdFile) | **Get** /Sync/JobItems/{Id}/File | Gets a sync job item file
[**GetSyncJobs**](SyncServiceAPI.md#GetSyncJobs) | **Get** /Sync/Jobs | Gets sync jobs.
[**GetSyncJobsById**](SyncServiceAPI.md#GetSyncJobsById) | **Get** /Sync/Jobs/{Id} | Gets a sync job.
[**GetSyncOptions**](SyncServiceAPI.md#GetSyncOptions) | **Get** /Sync/Options | Gets a list of available sync targets.
[**GetSyncTargets**](SyncServiceAPI.md#GetSyncTargets) | **Get** /Sync/Targets | Gets a list of available sync targets.
[**HeadSyncJobitemsByIdFile**](SyncServiceAPI.md#HeadSyncJobitemsByIdFile) | **Head** /Sync/JobItems/{Id}/File | Gets a sync job item file
[**PostSyncByItemidStatus**](SyncServiceAPI.md#PostSyncByItemidStatus) | **Post** /Sync/{ItemId}/Status | Gets sync status for an item.
[**PostSyncByTargetidItemsDelete**](SyncServiceAPI.md#PostSyncByTargetidItemsDelete) | **Post** /Sync/{TargetId}/Items/Delete | Cancels items from a sync target
[**PostSyncData**](SyncServiceAPI.md#PostSyncData) | **Post** /Sync/Data | Syncs data between device and server
[**PostSyncItemsCancel**](SyncServiceAPI.md#PostSyncItemsCancel) | **Post** /Sync/Items/Cancel | Cancels items from a sync target
[**PostSyncJobitemsByIdDelete**](SyncServiceAPI.md#PostSyncJobitemsByIdDelete) | **Post** /Sync/JobItems/{Id}/Delete | Cancels a sync job item
[**PostSyncJobitemsByIdEnable**](SyncServiceAPI.md#PostSyncJobitemsByIdEnable) | **Post** /Sync/JobItems/{Id}/Enable | Enables a cancelled or queued sync job item
[**PostSyncJobitemsByIdMarkforremoval**](SyncServiceAPI.md#PostSyncJobitemsByIdMarkforremoval) | **Post** /Sync/JobItems/{Id}/MarkForRemoval | Marks a job item for removal
[**PostSyncJobitemsByIdTransferred**](SyncServiceAPI.md#PostSyncJobitemsByIdTransferred) | **Post** /Sync/JobItems/{Id}/Transferred | Reports that a sync job item has successfully been transferred.
[**PostSyncJobitemsByIdUnmarkforremoval**](SyncServiceAPI.md#PostSyncJobitemsByIdUnmarkforremoval) | **Post** /Sync/JobItems/{Id}/UnmarkForRemoval | Unmarks a job item for removal
[**PostSyncJobs**](SyncServiceAPI.md#PostSyncJobs) | **Post** /Sync/Jobs | Gets sync jobs.
[**PostSyncJobsById**](SyncServiceAPI.md#PostSyncJobsById) | **Post** /Sync/Jobs/{Id} | Updates a sync job.
[**PostSyncJobsByIdDelete**](SyncServiceAPI.md#PostSyncJobsByIdDelete) | **Post** /Sync/Jobs/{Id}/Delete | Cancels a sync job.
[**PostSyncOfflineactions**](SyncServiceAPI.md#PostSyncOfflineactions) | **Post** /Sync/OfflineActions | Reports an action that occurred while offline.



## DeleteSyncByTargetidItems

> DeleteSyncByTargetidItems(ctx, targetId).ItemIds(itemIds).Execute()

Cancels items from a sync target



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
	targetId := "targetId_example" // string | TargetId
	itemIds := "itemIds_example" // string | ItemIds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.DeleteSyncByTargetidItems(context.Background(), targetId).ItemIds(itemIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.DeleteSyncByTargetidItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetId** | **string** | TargetId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyncByTargetidItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemIds** | **string** | ItemIds | 

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


## DeleteSyncJobitemsById

> DeleteSyncJobitemsById(ctx, id).Execute()

Cancels a sync job item



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.DeleteSyncJobitemsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.DeleteSyncJobitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyncJobitemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteSyncJobsById

> DeleteSyncJobsById(ctx, id).Execute()

Cancels a sync job.



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.DeleteSyncJobsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.DeleteSyncJobsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyncJobsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSyncItemsReady

> []SyncedItem GetSyncItemsReady(ctx).TargetId(targetId).Execute()

Gets ready to download sync items.



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
	targetId := "targetId_example" // string | TargetId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.GetSyncItemsReady(context.Background()).TargetId(targetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncItemsReady``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncItemsReady`: []SyncedItem
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncItemsReady`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncItemsReadyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetId** | **string** | TargetId | 

### Return type

[**[]SyncedItem**](SyncedItem.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncJobitems

> QueryResultSyncJobItem GetSyncJobitems(ctx).TargetId(targetId).Execute()

Gets sync job items.



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
	targetId := "targetId_example" // string | TargetId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.GetSyncJobitems(context.Background()).TargetId(targetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncJobitems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncJobitems`: QueryResultSyncJobItem
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncJobitems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncJobitemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetId** | **string** | TargetId | 

### Return type

[**QueryResultSyncJobItem**](QueryResultSyncJobItem.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncJobitemsByIdAdditionalfiles

> GetSyncJobitemsByIdAdditionalfiles(ctx, id).Name(name).Execute()

Gets a sync job item file



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
	id := "id_example" // string | Id
	name := "name_example" // string | Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.GetSyncJobitemsByIdAdditionalfiles(context.Background(), id).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncJobitemsByIdAdditionalfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncJobitemsByIdAdditionalfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name | 

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


## GetSyncJobitemsByIdFile

> GetSyncJobitemsByIdFile(ctx, id).Execute()

Gets a sync job item file



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.GetSyncJobitemsByIdFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncJobitemsByIdFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncJobitemsByIdFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSyncJobs

> QueryResultSyncJob GetSyncJobs(ctx).Execute()

Gets sync jobs.



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
	resp, r, err := apiClient.SyncServiceAPI.GetSyncJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncJobs`: QueryResultSyncJob
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncJobsRequest struct via the builder pattern


### Return type

[**QueryResultSyncJob**](QueryResultSyncJob.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncJobsById

> SyncJob GetSyncJobsById(ctx, id).Execute()

Gets a sync job.



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.GetSyncJobsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncJobsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncJobsById`: SyncJob
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncJobsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncJobsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncJob**](SyncJob.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncOptions

> SyncDialogOptions GetSyncOptions(ctx).UserId(userId).ItemIds(itemIds).ParentId(parentId).TargetId(targetId).Category(category).Execute()

Gets a list of available sync targets.



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
	userId := "userId_example" // string | UserId
	itemIds := "itemIds_example" // string | ItemIds (optional)
	parentId := "parentId_example" // string | ParentId (optional)
	targetId := "targetId_example" // string | TargetId (optional)
	category := openapiclient.SyncCategory("Latest") // SyncCategory | Category (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.GetSyncOptions(context.Background()).UserId(userId).ItemIds(itemIds).ParentId(parentId).TargetId(targetId).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncOptions`: SyncDialogOptions
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | UserId | 
 **itemIds** | **string** | ItemIds | 
 **parentId** | **string** | ParentId | 
 **targetId** | **string** | TargetId | 
 **category** | [**SyncCategory**](SyncCategory.md) | Category | 

### Return type

[**SyncDialogOptions**](SyncDialogOptions.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncTargets

> []SyncTarget GetSyncTargets(ctx).UserId(userId).Execute()

Gets a list of available sync targets.



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
	userId := "userId_example" // string | UserId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.GetSyncTargets(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncTargets`: []SyncTarget
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | UserId | 

### Return type

[**[]SyncTarget**](SyncTarget.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadSyncJobitemsByIdFile

> HeadSyncJobitemsByIdFile(ctx, id).Execute()

Gets a sync job item file



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.HeadSyncJobitemsByIdFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.HeadSyncJobitemsByIdFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadSyncJobitemsByIdFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostSyncByItemidStatus

> PostSyncByItemidStatus(ctx, itemId).SyncedItemProgress(syncedItemProgress).Execute()

Gets sync status for an item.



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
	itemId := "itemId_example" // string | 
	syncedItemProgress := *openapiclient.NewSyncedItemProgress() // SyncedItemProgress | SyncedItemProgress: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncByItemidStatus(context.Background(), itemId).SyncedItemProgress(syncedItemProgress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncByItemidStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncByItemidStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncedItemProgress** | [**SyncedItemProgress**](SyncedItemProgress.md) | SyncedItemProgress:  | 

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


## PostSyncByTargetidItemsDelete

> PostSyncByTargetidItemsDelete(ctx, targetId).ItemIds(itemIds).Execute()

Cancels items from a sync target



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
	targetId := "targetId_example" // string | TargetId
	itemIds := "itemIds_example" // string | ItemIds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncByTargetidItemsDelete(context.Background(), targetId).ItemIds(itemIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncByTargetidItemsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetId** | **string** | TargetId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncByTargetidItemsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemIds** | **string** | ItemIds | 

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


## PostSyncData

> SyncDataResponse PostSyncData(ctx).TargetId(targetId).SyncDataRequest(syncDataRequest).Execute()

Syncs data between device and server



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
	targetId := "targetId_example" // string | TargetId
	syncDataRequest := *openapiclient.NewSyncDataRequest() // SyncDataRequest | SyncDataRequest: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.PostSyncData(context.Background()).TargetId(targetId).SyncDataRequest(syncDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSyncData`: SyncDataResponse
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.PostSyncData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetId** | **string** | TargetId | 
 **syncDataRequest** | [**SyncDataRequest**](SyncDataRequest.md) | SyncDataRequest:  | 

### Return type

[**SyncDataResponse**](SyncDataResponse.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSyncItemsCancel

> PostSyncItemsCancel(ctx).ItemIds(itemIds).Execute()

Cancels items from a sync target



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
	itemIds := "itemIds_example" // string | ItemIds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncItemsCancel(context.Background()).ItemIds(itemIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncItemsCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncItemsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemIds** | **string** | ItemIds | 

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


## PostSyncJobitemsByIdDelete

> PostSyncJobitemsByIdDelete(ctx, id).Execute()

Cancels a sync job item



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobitemsByIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobitemsByIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobitemsByIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostSyncJobitemsByIdEnable

> PostSyncJobitemsByIdEnable(ctx, id).Execute()

Enables a cancelled or queued sync job item



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobitemsByIdEnable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobitemsByIdEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobitemsByIdEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostSyncJobitemsByIdMarkforremoval

> PostSyncJobitemsByIdMarkforremoval(ctx, id).Execute()

Marks a job item for removal



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobitemsByIdMarkforremoval(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobitemsByIdMarkforremoval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobitemsByIdMarkforremovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostSyncJobitemsByIdTransferred

> PostSyncJobitemsByIdTransferred(ctx, id).Execute()

Reports that a sync job item has successfully been transferred.



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobitemsByIdTransferred(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobitemsByIdTransferred``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobitemsByIdTransferredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostSyncJobitemsByIdUnmarkforremoval

> PostSyncJobitemsByIdUnmarkforremoval(ctx, id).Execute()

Unmarks a job item for removal



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobitemsByIdUnmarkforremoval(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobitemsByIdUnmarkforremoval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobitemsByIdUnmarkforremovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostSyncJobs

> SyncJobCreationResult PostSyncJobs(ctx).SyncJobRequest(syncJobRequest).Execute()

Gets sync jobs.



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
	syncJobRequest := *openapiclient.NewSyncJobRequest() // SyncJobRequest | SyncJobRequest: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.PostSyncJobs(context.Background()).SyncJobRequest(syncJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSyncJobs`: SyncJobCreationResult
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.PostSyncJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncJobRequest** | [**SyncJobRequest**](SyncJobRequest.md) | SyncJobRequest:  | 

### Return type

[**SyncJobCreationResult**](SyncJobCreationResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSyncJobsById

> PostSyncJobsById(ctx, id).SyncJob(syncJob).Execute()

Updates a sync job.



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
	id := int64(789) // int64 | 
	syncJob := *openapiclient.NewSyncJob() // SyncJob | SyncJob: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobsById(context.Background(), id).SyncJob(syncJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncJob** | [**SyncJob**](SyncJob.md) | SyncJob:  | 

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


## PostSyncJobsByIdDelete

> PostSyncJobsByIdDelete(ctx, id).Execute()

Cancels a sync job.



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobsByIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobsByIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobsByIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostSyncOfflineactions

> PostSyncOfflineactions(ctx).UserAction(userAction).Execute()

Reports an action that occurred while offline.



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
	userAction := []openapiclient.UserAction{*openapiclient.NewUserAction()} // []UserAction | List`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncOfflineactions(context.Background()).UserAction(userAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncOfflineactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncOfflineactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAction** | [**[]UserAction**](UserAction.md) | List&#x60;1:  | 

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

