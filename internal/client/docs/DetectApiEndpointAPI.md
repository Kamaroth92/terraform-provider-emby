# \DetectApiEndpointAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChapterApiCancelJob**](DetectApiEndpointAPI.md#GetChapterApiCancelJob) | **Get** /chapter_api/cancel_job | Cancel a job
[**GetChapterApiDownloadIntroData**](DetectApiEndpointAPI.md#GetChapterApiDownloadIntroData) | **Get** /chapter_api/download_intro_data | Intro DB download
[**GetChapterApiGetEpisodeList**](DetectApiEndpointAPI.md#GetChapterApiGetEpisodeList) | **Get** /chapter_api/get_episode_list | Get list of episodes
[**GetChapterApiGetJobInfo**](DetectApiEndpointAPI.md#GetChapterApiGetJobInfo) | **Get** /chapter_api/get_job_info | Get job info
[**GetChapterApiGetJobItem**](DetectApiEndpointAPI.md#GetChapterApiGetJobItem) | **Get** /chapter_api/get_job_item | Gets info for a job work item
[**GetChapterApiGetJobList**](DetectApiEndpointAPI.md#GetChapterApiGetJobList) | **Get** /chapter_api/get_job_list | Get list of jobs
[**GetChapterApiGetSeasonList**](DetectApiEndpointAPI.md#GetChapterApiGetSeasonList) | **Get** /chapter_api/get_season_list | Get list of seasons
[**GetChapterApiGetSeriesList**](DetectApiEndpointAPI.md#GetChapterApiGetSeriesList) | **Get** /chapter_api/get_series_list | Get list of series
[**GetChapterApiInsertChapters**](DetectApiEndpointAPI.md#GetChapterApiInsertChapters) | **Get** /chapter_api/insert_chapters | Insert detected chapters
[**GetChapterApiIntroDataStats**](DetectApiEndpointAPI.md#GetChapterApiIntroDataStats) | **Get** /chapter_api/intro_data_stats | Intro DB stats
[**GetChapterApiReloadIntroData**](DetectApiEndpointAPI.md#GetChapterApiReloadIntroData) | **Get** /chapter_api/reload_intro_data | Reloads the intro DB form the data path
[**GetChapterApiRemoveJob**](DetectApiEndpointAPI.md#GetChapterApiRemoveJob) | **Get** /chapter_api/remove_job | Remove a job
[**PostChapterApiAddDetectionJob**](DetectApiEndpointAPI.md#PostChapterApiAddDetectionJob) | **Post** /chapter_api/add_detection_job | Add detection job



## GetChapterApiCancelJob

> map[string]interface{} GetChapterApiCancelJob(ctx).Execute()

Cancel a job



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiCancelJob(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiCancelJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiCancelJob`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiCancelJob`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiCancelJobRequest struct via the builder pattern


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


## GetChapterApiDownloadIntroData

> map[string]interface{} GetChapterApiDownloadIntroData(ctx).Execute()

Intro DB download



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiDownloadIntroData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiDownloadIntroData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiDownloadIntroData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiDownloadIntroData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiDownloadIntroDataRequest struct via the builder pattern


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


## GetChapterApiGetEpisodeList

> map[string]interface{} GetChapterApiGetEpisodeList(ctx).Execute()

Get list of episodes



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiGetEpisodeList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiGetEpisodeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiGetEpisodeList`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiGetEpisodeList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiGetEpisodeListRequest struct via the builder pattern


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


## GetChapterApiGetJobInfo

> map[string]interface{} GetChapterApiGetJobInfo(ctx).Execute()

Get job info



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiGetJobInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiGetJobInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiGetJobInfo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiGetJobInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiGetJobInfoRequest struct via the builder pattern


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


## GetChapterApiGetJobItem

> map[string]interface{} GetChapterApiGetJobItem(ctx).Execute()

Gets info for a job work item



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiGetJobItem(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiGetJobItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiGetJobItem`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiGetJobItem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiGetJobItemRequest struct via the builder pattern


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


## GetChapterApiGetJobList

> map[string]interface{} GetChapterApiGetJobList(ctx).Execute()

Get list of jobs



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiGetJobList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiGetJobList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiGetJobList`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiGetJobList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiGetJobListRequest struct via the builder pattern


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


## GetChapterApiGetSeasonList

> map[string]interface{} GetChapterApiGetSeasonList(ctx).Execute()

Get list of seasons



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiGetSeasonList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiGetSeasonList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiGetSeasonList`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiGetSeasonList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiGetSeasonListRequest struct via the builder pattern


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


## GetChapterApiGetSeriesList

> map[string]interface{} GetChapterApiGetSeriesList(ctx).Execute()

Get list of series



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiGetSeriesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiGetSeriesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiGetSeriesList`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiGetSeriesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiGetSeriesListRequest struct via the builder pattern


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


## GetChapterApiInsertChapters

> map[string]interface{} GetChapterApiInsertChapters(ctx).Execute()

Insert detected chapters



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiInsertChapters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiInsertChapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiInsertChapters`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiInsertChapters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiInsertChaptersRequest struct via the builder pattern


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


## GetChapterApiIntroDataStats

> map[string]interface{} GetChapterApiIntroDataStats(ctx).Execute()

Intro DB stats



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiIntroDataStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiIntroDataStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiIntroDataStats`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiIntroDataStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiIntroDataStatsRequest struct via the builder pattern


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


## GetChapterApiReloadIntroData

> map[string]interface{} GetChapterApiReloadIntroData(ctx).Execute()

Reloads the intro DB form the data path



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiReloadIntroData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiReloadIntroData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiReloadIntroData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiReloadIntroData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiReloadIntroDataRequest struct via the builder pattern


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


## GetChapterApiRemoveJob

> map[string]interface{} GetChapterApiRemoveJob(ctx).Execute()

Remove a job



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
	resp, r, err := apiClient.DetectApiEndpointAPI.GetChapterApiRemoveJob(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.GetChapterApiRemoveJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChapterApiRemoveJob`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.GetChapterApiRemoveJob`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChapterApiRemoveJobRequest struct via the builder pattern


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


## PostChapterApiAddDetectionJob

> map[string]interface{} PostChapterApiAddDetectionJob(ctx).ChapterApiAddDetectionJob(chapterApiAddDetectionJob).Execute()

Add detection job



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
	chapterApiAddDetectionJob := *openapiclient.NewChapterApiAddDetectionJob() // ChapterApiAddDetectionJob | AddDetectionJob

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DetectApiEndpointAPI.PostChapterApiAddDetectionJob(context.Background()).ChapterApiAddDetectionJob(chapterApiAddDetectionJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectApiEndpointAPI.PostChapterApiAddDetectionJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostChapterApiAddDetectionJob`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DetectApiEndpointAPI.PostChapterApiAddDetectionJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostChapterApiAddDetectionJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chapterApiAddDetectionJob** | [**ChapterApiAddDetectionJob**](ChapterApiAddDetectionJob.md) | AddDetectionJob | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

