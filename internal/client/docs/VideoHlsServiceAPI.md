# \VideoHlsServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer**](VideoHlsServiceAPI.md#GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer) | **Get** /Audio/{Id}/hls/{PlaylistId}/{SegmentId}.{SegmentContainer} | 
[**GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer**](VideoHlsServiceAPI.md#GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer) | **Get** /Videos/{Id}/hls/{PlaylistId}/{SegmentId}.{SegmentContainer} | 



## GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer

> GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer(ctx, segmentContainer, segmentId, id, playlistId).Execute()





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
	segmentContainer := "segmentContainer_example" // string | 
	segmentId := "segmentId_example" // string | 
	id := "id_example" // string | 
	playlistId := "playlistId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoHlsServiceAPI.GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer(context.Background(), segmentContainer, segmentId, id, playlistId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoHlsServiceAPI.GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**segmentContainer** | **string** |  | 
**segmentId** | **string** |  | 
**id** | **string** |  | 
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainerRequest struct via the builder pattern


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


## GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer

> GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer(ctx, segmentContainer, segmentId, id, playlistId).Execute()





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
	segmentContainer := "segmentContainer_example" // string | 
	segmentId := "segmentId_example" // string | 
	id := "id_example" // string | 
	playlistId := "playlistId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoHlsServiceAPI.GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer(context.Background(), segmentContainer, segmentId, id, playlistId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoHlsServiceAPI.GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**segmentContainer** | **string** |  | 
**segmentId** | **string** |  | 
**id** | **string** |  | 
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainerRequest struct via the builder pattern


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

