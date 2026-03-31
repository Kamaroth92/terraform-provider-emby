# \LiveStreamServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLivetvLiverecordingsByIdHlsBySegment**](LiveStreamServiceAPI.md#GetLivetvLiverecordingsByIdHlsBySegment) | **Get** /LiveTv/LiveRecordings/{Id}/hls/{Segment} | Gets a live recording
[**GetLivetvLiverecordingsByIdHlsLiveM3u8**](LiveStreamServiceAPI.md#GetLivetvLiverecordingsByIdHlsLiveM3u8) | **Get** /LiveTv/LiveRecordings/{Id}/hls/live.m3u8 | Gets a live recording
[**GetLivetvLiverecordingsByIdHlsMasterM3u8**](LiveStreamServiceAPI.md#GetLivetvLiverecordingsByIdHlsMasterM3u8) | **Get** /LiveTv/LiveRecordings/{Id}/hls/master.m3u8 | Gets a live recording
[**GetLivetvLiverecordingsByIdStream**](LiveStreamServiceAPI.md#GetLivetvLiverecordingsByIdStream) | **Get** /LiveTv/LiveRecordings/{Id}/stream | Gets a live tv channel
[**GetLivetvLivestreamfilesByIdHlsBySegment**](LiveStreamServiceAPI.md#GetLivetvLivestreamfilesByIdHlsBySegment) | **Get** /LiveTv/LiveStreamFiles/{Id}/hls/{Segment} | Gets a live tv channel
[**GetLivetvLivestreamfilesByIdHlsLiveM3u8**](LiveStreamServiceAPI.md#GetLivetvLivestreamfilesByIdHlsLiveM3u8) | **Get** /LiveTv/LiveStreamFiles/{Id}/hls/live.m3u8 | Gets a live tv channel
[**GetLivetvLivestreamfilesByIdHlsMasterM3u8**](LiveStreamServiceAPI.md#GetLivetvLivestreamfilesByIdHlsMasterM3u8) | **Get** /LiveTv/LiveStreamFiles/{Id}/hls/master.m3u8 | Gets a live tv channel
[**GetLivetvLivestreamfilesByIdStreamByContainer**](LiveStreamServiceAPI.md#GetLivetvLivestreamfilesByIdStreamByContainer) | **Get** /LiveTv/LiveStreamFiles/{Id}/stream.{Container} | Gets a live tv channel
[**HeadLivetvLiverecordingsByIdHlsBySegment**](LiveStreamServiceAPI.md#HeadLivetvLiverecordingsByIdHlsBySegment) | **Head** /LiveTv/LiveRecordings/{Id}/hls/{Segment} | Gets a live recording
[**HeadLivetvLiverecordingsByIdHlsLiveM3u8**](LiveStreamServiceAPI.md#HeadLivetvLiverecordingsByIdHlsLiveM3u8) | **Head** /LiveTv/LiveRecordings/{Id}/hls/live.m3u8 | Gets a live recording
[**HeadLivetvLiverecordingsByIdHlsMasterM3u8**](LiveStreamServiceAPI.md#HeadLivetvLiverecordingsByIdHlsMasterM3u8) | **Head** /LiveTv/LiveRecordings/{Id}/hls/master.m3u8 | Gets a live recording
[**HeadLivetvLivestreamfilesByIdHlsBySegment**](LiveStreamServiceAPI.md#HeadLivetvLivestreamfilesByIdHlsBySegment) | **Head** /LiveTv/LiveStreamFiles/{Id}/hls/{Segment} | Gets a live tv channel
[**HeadLivetvLivestreamfilesByIdHlsLiveM3u8**](LiveStreamServiceAPI.md#HeadLivetvLivestreamfilesByIdHlsLiveM3u8) | **Head** /LiveTv/LiveStreamFiles/{Id}/hls/live.m3u8 | Gets a live tv channel
[**HeadLivetvLivestreamfilesByIdHlsMasterM3u8**](LiveStreamServiceAPI.md#HeadLivetvLivestreamfilesByIdHlsMasterM3u8) | **Head** /LiveTv/LiveStreamFiles/{Id}/hls/master.m3u8 | Gets a live tv channel



## GetLivetvLiverecordingsByIdHlsBySegment

> GetLivetvLiverecordingsByIdHlsBySegment(ctx, id, segment).Execute()

Gets a live recording



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
	id := "id_example" // string | 
	segment := "segment_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.GetLivetvLiverecordingsByIdHlsBySegment(context.Background(), id, segment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.GetLivetvLiverecordingsByIdHlsBySegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**segment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvLiverecordingsByIdHlsBySegmentRequest struct via the builder pattern


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


## GetLivetvLiverecordingsByIdHlsLiveM3u8

> GetLivetvLiverecordingsByIdHlsLiveM3u8(ctx, id).Execute()

Gets a live recording



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.GetLivetvLiverecordingsByIdHlsLiveM3u8(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.GetLivetvLiverecordingsByIdHlsLiveM3u8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvLiverecordingsByIdHlsLiveM3u8Request struct via the builder pattern


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


## GetLivetvLiverecordingsByIdHlsMasterM3u8

> GetLivetvLiverecordingsByIdHlsMasterM3u8(ctx, id).Execute()

Gets a live recording



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.GetLivetvLiverecordingsByIdHlsMasterM3u8(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.GetLivetvLiverecordingsByIdHlsMasterM3u8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvLiverecordingsByIdHlsMasterM3u8Request struct via the builder pattern


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


## GetLivetvLiverecordingsByIdStream

> GetLivetvLiverecordingsByIdStream(ctx, id).Execute()

Gets a live tv channel



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.GetLivetvLiverecordingsByIdStream(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.GetLivetvLiverecordingsByIdStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvLiverecordingsByIdStreamRequest struct via the builder pattern


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


## GetLivetvLivestreamfilesByIdHlsBySegment

> GetLivetvLivestreamfilesByIdHlsBySegment(ctx, id, segment).Execute()

Gets a live tv channel



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
	id := "id_example" // string | 
	segment := "segment_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.GetLivetvLivestreamfilesByIdHlsBySegment(context.Background(), id, segment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.GetLivetvLivestreamfilesByIdHlsBySegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**segment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvLivestreamfilesByIdHlsBySegmentRequest struct via the builder pattern


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


## GetLivetvLivestreamfilesByIdHlsLiveM3u8

> GetLivetvLivestreamfilesByIdHlsLiveM3u8(ctx, id).Execute()

Gets a live tv channel



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.GetLivetvLivestreamfilesByIdHlsLiveM3u8(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.GetLivetvLivestreamfilesByIdHlsLiveM3u8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvLivestreamfilesByIdHlsLiveM3u8Request struct via the builder pattern


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


## GetLivetvLivestreamfilesByIdHlsMasterM3u8

> GetLivetvLivestreamfilesByIdHlsMasterM3u8(ctx, id).Execute()

Gets a live tv channel



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.GetLivetvLivestreamfilesByIdHlsMasterM3u8(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.GetLivetvLivestreamfilesByIdHlsMasterM3u8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvLivestreamfilesByIdHlsMasterM3u8Request struct via the builder pattern


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


## GetLivetvLivestreamfilesByIdStreamByContainer

> GetLivetvLivestreamfilesByIdStreamByContainer(ctx, id, container).Execute()

Gets a live tv channel



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
	id := "id_example" // string | 
	container := "container_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.GetLivetvLivestreamfilesByIdStreamByContainer(context.Background(), id, container).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.GetLivetvLivestreamfilesByIdStreamByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**container** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvLivestreamfilesByIdStreamByContainerRequest struct via the builder pattern


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


## HeadLivetvLiverecordingsByIdHlsBySegment

> HeadLivetvLiverecordingsByIdHlsBySegment(ctx, id, segment).Execute()

Gets a live recording



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
	id := "id_example" // string | 
	segment := "segment_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.HeadLivetvLiverecordingsByIdHlsBySegment(context.Background(), id, segment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.HeadLivetvLiverecordingsByIdHlsBySegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**segment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadLivetvLiverecordingsByIdHlsBySegmentRequest struct via the builder pattern


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


## HeadLivetvLiverecordingsByIdHlsLiveM3u8

> HeadLivetvLiverecordingsByIdHlsLiveM3u8(ctx, id).Execute()

Gets a live recording



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.HeadLivetvLiverecordingsByIdHlsLiveM3u8(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.HeadLivetvLiverecordingsByIdHlsLiveM3u8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadLivetvLiverecordingsByIdHlsLiveM3u8Request struct via the builder pattern


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


## HeadLivetvLiverecordingsByIdHlsMasterM3u8

> HeadLivetvLiverecordingsByIdHlsMasterM3u8(ctx, id).Execute()

Gets a live recording



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.HeadLivetvLiverecordingsByIdHlsMasterM3u8(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.HeadLivetvLiverecordingsByIdHlsMasterM3u8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadLivetvLiverecordingsByIdHlsMasterM3u8Request struct via the builder pattern


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


## HeadLivetvLivestreamfilesByIdHlsBySegment

> HeadLivetvLivestreamfilesByIdHlsBySegment(ctx, id, segment).Execute()

Gets a live tv channel



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
	id := "id_example" // string | 
	segment := "segment_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.HeadLivetvLivestreamfilesByIdHlsBySegment(context.Background(), id, segment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.HeadLivetvLivestreamfilesByIdHlsBySegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**segment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadLivetvLivestreamfilesByIdHlsBySegmentRequest struct via the builder pattern


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


## HeadLivetvLivestreamfilesByIdHlsLiveM3u8

> HeadLivetvLivestreamfilesByIdHlsLiveM3u8(ctx, id).Execute()

Gets a live tv channel



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.HeadLivetvLivestreamfilesByIdHlsLiveM3u8(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.HeadLivetvLivestreamfilesByIdHlsLiveM3u8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadLivetvLivestreamfilesByIdHlsLiveM3u8Request struct via the builder pattern


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


## HeadLivetvLivestreamfilesByIdHlsMasterM3u8

> HeadLivetvLivestreamfilesByIdHlsMasterM3u8(ctx, id).Execute()

Gets a live tv channel



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveStreamServiceAPI.HeadLivetvLivestreamfilesByIdHlsMasterM3u8(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamServiceAPI.HeadLivetvLivestreamfilesByIdHlsMasterM3u8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadLivetvLivestreamfilesByIdHlsMasterM3u8Request struct via the builder pattern


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

