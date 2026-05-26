# \PlaystateServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsersByUseridPlayeditemsById**](PlaystateServiceAPI.md#DeleteUsersByUseridPlayeditemsById) | **Delete** /Users/{UserId}/PlayedItems/{Id} | Marks an item as unplayed
[**DeleteUsersByUseridPlayingitemsById**](PlaystateServiceAPI.md#DeleteUsersByUseridPlayingitemsById) | **Delete** /Users/{UserId}/PlayingItems/{Id} | Reports that a user has stopped playing an item
[**PostSessionsPlaying**](PlaystateServiceAPI.md#PostSessionsPlaying) | **Post** /Sessions/Playing | Reports playback has started within a session
[**PostSessionsPlayingPing**](PlaystateServiceAPI.md#PostSessionsPlayingPing) | **Post** /Sessions/Playing/Ping | Pings a playback session
[**PostSessionsPlayingProgress**](PlaystateServiceAPI.md#PostSessionsPlayingProgress) | **Post** /Sessions/Playing/Progress | Reports playback progress within a session
[**PostSessionsPlayingStopped**](PlaystateServiceAPI.md#PostSessionsPlayingStopped) | **Post** /Sessions/Playing/Stopped | Reports playback has stopped within a session
[**PostUsersByUseridItemsByItemidUserdata**](PlaystateServiceAPI.md#PostUsersByUseridItemsByItemidUserdata) | **Post** /Users/{UserId}/Items/{ItemId}/UserData | Updates userdata for an item
[**PostUsersByUseridPlayeditemsById**](PlaystateServiceAPI.md#PostUsersByUseridPlayeditemsById) | **Post** /Users/{UserId}/PlayedItems/{Id} | Marks an item as played
[**PostUsersByUseridPlayeditemsByIdDelete**](PlaystateServiceAPI.md#PostUsersByUseridPlayeditemsByIdDelete) | **Post** /Users/{UserId}/PlayedItems/{Id}/Delete | Marks an item as unplayed
[**PostUsersByUseridPlayingitemsById**](PlaystateServiceAPI.md#PostUsersByUseridPlayingitemsById) | **Post** /Users/{UserId}/PlayingItems/{Id} | Reports that a user has begun playing an item
[**PostUsersByUseridPlayingitemsByIdDelete**](PlaystateServiceAPI.md#PostUsersByUseridPlayingitemsByIdDelete) | **Post** /Users/{UserId}/PlayingItems/{Id}/Delete | Reports that a user has stopped playing an item
[**PostUsersByUseridPlayingitemsByIdProgress**](PlaystateServiceAPI.md#PostUsersByUseridPlayingitemsByIdProgress) | **Post** /Users/{UserId}/PlayingItems/{Id}/Progress | Reports a user&#39;s playback progress



## DeleteUsersByUseridPlayeditemsById

> UserItemDataDto DeleteUsersByUseridPlayeditemsById(ctx, userId, id).Execute()

Marks an item as unplayed



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
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaystateServiceAPI.DeleteUsersByUseridPlayeditemsById(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.DeleteUsersByUseridPlayeditemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUsersByUseridPlayeditemsById`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `PlaystateServiceAPI.DeleteUsersByUseridPlayeditemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersByUseridPlayeditemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsersByUseridPlayingitemsById

> DeleteUsersByUseridPlayingitemsById(ctx, userId, id).MediaSourceId(mediaSourceId).NextMediaType(nextMediaType).PositionTicks(positionTicks).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()

Reports that a user has stopped playing an item



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
	id := "id_example" // string | Item Id
	mediaSourceId := "mediaSourceId_example" // string | The id of the MediaSource
	nextMediaType := "nextMediaType_example" // string | The next media type that will play
	positionTicks := int64(789) // int64 | Optional. The position, in ticks, where playback stopped. 1ms = 10000 ticks. (optional)
	liveStreamId := "liveStreamId_example" // string |  (optional)
	playSessionId := "playSessionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.DeleteUsersByUseridPlayingitemsById(context.Background(), userId, id).MediaSourceId(mediaSourceId).NextMediaType(nextMediaType).PositionTicks(positionTicks).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.DeleteUsersByUseridPlayingitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersByUseridPlayingitemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **string** | The id of the MediaSource | 
 **nextMediaType** | **string** | The next media type that will play | 
 **positionTicks** | **int64** | Optional. The position, in ticks, where playback stopped. 1ms &#x3D; 10000 ticks. | 
 **liveStreamId** | **string** |  | 
 **playSessionId** | **string** |  | 

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


## PostSessionsPlaying

> PostSessionsPlaying(ctx).PlaybackStartInfo(playbackStartInfo).Execute()

Reports playback has started within a session



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
	playbackStartInfo := *openapiclient.NewPlaybackStartInfo() // PlaybackStartInfo | PlaybackStartInfo: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostSessionsPlaying(context.Background()).PlaybackStartInfo(playbackStartInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostSessionsPlaying``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsPlayingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackStartInfo** | [**PlaybackStartInfo**](PlaybackStartInfo.md) | PlaybackStartInfo:  | 

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


## PostSessionsPlayingPing

> PostSessionsPlayingPing(ctx).PlaySessionId(playSessionId).Execute()

Pings a playback session



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
	playSessionId := "playSessionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostSessionsPlayingPing(context.Background()).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostSessionsPlayingPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsPlayingPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playSessionId** | **string** |  | 

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


## PostSessionsPlayingProgress

> PostSessionsPlayingProgress(ctx).PlaybackProgressInfo(playbackProgressInfo).Execute()

Reports playback progress within a session



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
	playbackProgressInfo := *openapiclient.NewPlaybackProgressInfo() // PlaybackProgressInfo | PlaybackProgressInfo: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostSessionsPlayingProgress(context.Background()).PlaybackProgressInfo(playbackProgressInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostSessionsPlayingProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsPlayingProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackProgressInfo** | [**PlaybackProgressInfo**](PlaybackProgressInfo.md) | PlaybackProgressInfo:  | 

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


## PostSessionsPlayingStopped

> PostSessionsPlayingStopped(ctx).PlaybackStopInfo(playbackStopInfo).Execute()

Reports playback has stopped within a session



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
	playbackStopInfo := *openapiclient.NewPlaybackStopInfo() // PlaybackStopInfo | PlaybackStopInfo: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostSessionsPlayingStopped(context.Background()).PlaybackStopInfo(playbackStopInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostSessionsPlayingStopped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsPlayingStoppedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackStopInfo** | [**PlaybackStopInfo**](PlaybackStopInfo.md) | PlaybackStopInfo:  | 

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


## PostUsersByUseridItemsByItemidUserdata

> PostUsersByUseridItemsByItemidUserdata(ctx, userId, itemId).UserItemDataDto(userItemDataDto).Execute()

Updates userdata for an item



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
	itemId := "itemId_example" // string | 
	userItemDataDto := *openapiclient.NewUserItemDataDto() // UserItemDataDto | UserItemDataDto: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostUsersByUseridItemsByItemidUserdata(context.Background(), userId, itemId).UserItemDataDto(userItemDataDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostUsersByUseridItemsByItemidUserdata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridItemsByItemidUserdataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userItemDataDto** | [**UserItemDataDto**](UserItemDataDto.md) | UserItemDataDto:  | 

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


## PostUsersByUseridPlayeditemsById

> UserItemDataDto PostUsersByUseridPlayeditemsById(ctx, userId, id).DatePlayed(datePlayed).Execute()

Marks an item as played



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
	id := "id_example" // string | Item Id
	datePlayed := "datePlayed_example" // string | The date the item was played (if any). Format = yyyyMMddHHmmss (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaystateServiceAPI.PostUsersByUseridPlayeditemsById(context.Background(), userId, id).DatePlayed(datePlayed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostUsersByUseridPlayeditemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersByUseridPlayeditemsById`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `PlaystateServiceAPI.PostUsersByUseridPlayeditemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridPlayeditemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datePlayed** | **string** | The date the item was played (if any). Format &#x3D; yyyyMMddHHmmss | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersByUseridPlayeditemsByIdDelete

> UserItemDataDto PostUsersByUseridPlayeditemsByIdDelete(ctx, userId, id).Execute()

Marks an item as unplayed



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
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaystateServiceAPI.PostUsersByUseridPlayeditemsByIdDelete(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostUsersByUseridPlayeditemsByIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersByUseridPlayeditemsByIdDelete`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `PlaystateServiceAPI.PostUsersByUseridPlayeditemsByIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridPlayeditemsByIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersByUseridPlayingitemsById

> PostUsersByUseridPlayingitemsById(ctx, userId, id).MediaSourceId(mediaSourceId).CanSeek(canSeek).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()

Reports that a user has begun playing an item



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
	id := "id_example" // string | Item Id
	mediaSourceId := "mediaSourceId_example" // string | The id of the MediaSource
	canSeek := true // bool | Indicates if the client can seek (optional)
	audioStreamIndex := int32(56) // int32 |  (optional)
	subtitleStreamIndex := int32(56) // int32 |  (optional)
	playMethod := openapiclient.PlayMethod("Transcode") // PlayMethod |  (optional)
	liveStreamId := "liveStreamId_example" // string |  (optional)
	playSessionId := "playSessionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostUsersByUseridPlayingitemsById(context.Background(), userId, id).MediaSourceId(mediaSourceId).CanSeek(canSeek).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostUsersByUseridPlayingitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridPlayingitemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **string** | The id of the MediaSource | 
 **canSeek** | **bool** | Indicates if the client can seek | 
 **audioStreamIndex** | **int32** |  | 
 **subtitleStreamIndex** | **int32** |  | 
 **playMethod** | [**PlayMethod**](PlayMethod.md) |  | 
 **liveStreamId** | **string** |  | 
 **playSessionId** | **string** |  | 

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


## PostUsersByUseridPlayingitemsByIdDelete

> PostUsersByUseridPlayingitemsByIdDelete(ctx, userId, id).MediaSourceId(mediaSourceId).NextMediaType(nextMediaType).PositionTicks(positionTicks).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()

Reports that a user has stopped playing an item



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
	id := "id_example" // string | Item Id
	mediaSourceId := "mediaSourceId_example" // string | The id of the MediaSource
	nextMediaType := "nextMediaType_example" // string | The next media type that will play
	positionTicks := int64(789) // int64 | Optional. The position, in ticks, where playback stopped. 1ms = 10000 ticks. (optional)
	liveStreamId := "liveStreamId_example" // string |  (optional)
	playSessionId := "playSessionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostUsersByUseridPlayingitemsByIdDelete(context.Background(), userId, id).MediaSourceId(mediaSourceId).NextMediaType(nextMediaType).PositionTicks(positionTicks).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostUsersByUseridPlayingitemsByIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridPlayingitemsByIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **string** | The id of the MediaSource | 
 **nextMediaType** | **string** | The next media type that will play | 
 **positionTicks** | **int64** | Optional. The position, in ticks, where playback stopped. 1ms &#x3D; 10000 ticks. | 
 **liveStreamId** | **string** |  | 
 **playSessionId** | **string** |  | 

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


## PostUsersByUseridPlayingitemsByIdProgress

> PostUsersByUseridPlayingitemsByIdProgress(ctx, userId, id).MediaSourceId(mediaSourceId).ApiOnPlaybackProgress(apiOnPlaybackProgress).PositionTicks(positionTicks).IsPaused(isPaused).IsMuted(isMuted).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).VolumeLevel(volumeLevel).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).RepeatMode(repeatMode).SubtitleOffset(subtitleOffset).PlaybackRate(playbackRate).Execute()

Reports a user's playback progress



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
	id := "id_example" // string | Item Id
	mediaSourceId := "mediaSourceId_example" // string | The id of the MediaSource
	apiOnPlaybackProgress := *openapiclient.NewApiOnPlaybackProgress() // ApiOnPlaybackProgress | OnPlaybackProgress
	positionTicks := int64(789) // int64 | Optional. The current position, in ticks. 1ms = 10000 ticks. (optional)
	isPaused := true // bool | Indicates if the player is paused. (optional)
	isMuted := true // bool | Indicates if the player is muted. (optional)
	audioStreamIndex := int32(56) // int32 |  (optional)
	subtitleStreamIndex := int32(56) // int32 |  (optional)
	volumeLevel := int32(56) // int32 | Scale of 0-100 (optional)
	playMethod := openapiclient.PlayMethod("Transcode") // PlayMethod |  (optional)
	liveStreamId := "liveStreamId_example" // string |  (optional)
	playSessionId := "playSessionId_example" // string |  (optional)
	repeatMode := openapiclient.RepeatMode("RepeatNone") // RepeatMode |  (optional)
	subtitleOffset := int32(56) // int32 |  (optional)
	playbackRate := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostUsersByUseridPlayingitemsByIdProgress(context.Background(), userId, id).MediaSourceId(mediaSourceId).ApiOnPlaybackProgress(apiOnPlaybackProgress).PositionTicks(positionTicks).IsPaused(isPaused).IsMuted(isMuted).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).VolumeLevel(volumeLevel).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).RepeatMode(repeatMode).SubtitleOffset(subtitleOffset).PlaybackRate(playbackRate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostUsersByUseridPlayingitemsByIdProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridPlayingitemsByIdProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **string** | The id of the MediaSource | 
 **apiOnPlaybackProgress** | [**ApiOnPlaybackProgress**](ApiOnPlaybackProgress.md) | OnPlaybackProgress | 
 **positionTicks** | **int64** | Optional. The current position, in ticks. 1ms &#x3D; 10000 ticks. | 
 **isPaused** | **bool** | Indicates if the player is paused. | 
 **isMuted** | **bool** | Indicates if the player is muted. | 
 **audioStreamIndex** | **int32** |  | 
 **subtitleStreamIndex** | **int32** |  | 
 **volumeLevel** | **int32** | Scale of 0-100 | 
 **playMethod** | [**PlayMethod**](PlayMethod.md) |  | 
 **liveStreamId** | **string** |  | 
 **playSessionId** | **string** |  | 
 **repeatMode** | [**RepeatMode**](RepeatMode.md) |  | 
 **subtitleOffset** | **int32** |  | 
 **playbackRate** | **float64** |  | 

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

