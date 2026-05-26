# \ItemLookupServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemsByIdExternalidinfos**](ItemLookupServiceAPI.md#GetItemsByIdExternalidinfos) | **Get** /Items/{Id}/ExternalIdInfos | Gets external id infos for an item
[**GetItemsRemotesearchImage**](ItemLookupServiceAPI.md#GetItemsRemotesearchImage) | **Get** /Items/RemoteSearch/Image | Gets a remote image
[**PostItemsMetadataReset**](ItemLookupServiceAPI.md#PostItemsMetadataReset) | **Post** /Items/Metadata/Reset | Resets metadata for one or more items
[**PostItemsRemotesearchApplyById**](ItemLookupServiceAPI.md#PostItemsRemotesearchApplyById) | **Post** /Items/RemoteSearch/Apply/{Id} | Applies search criteria to an item and refreshes metadata
[**PostItemsRemotesearchBook**](ItemLookupServiceAPI.md#PostItemsRemotesearchBook) | **Post** /Items/RemoteSearch/Book | 
[**PostItemsRemotesearchBoxset**](ItemLookupServiceAPI.md#PostItemsRemotesearchBoxset) | **Post** /Items/RemoteSearch/BoxSet | 
[**PostItemsRemotesearchGame**](ItemLookupServiceAPI.md#PostItemsRemotesearchGame) | **Post** /Items/RemoteSearch/Game | 
[**PostItemsRemotesearchMovie**](ItemLookupServiceAPI.md#PostItemsRemotesearchMovie) | **Post** /Items/RemoteSearch/Movie | 
[**PostItemsRemotesearchMusicalbum**](ItemLookupServiceAPI.md#PostItemsRemotesearchMusicalbum) | **Post** /Items/RemoteSearch/MusicAlbum | 
[**PostItemsRemotesearchMusicartist**](ItemLookupServiceAPI.md#PostItemsRemotesearchMusicartist) | **Post** /Items/RemoteSearch/MusicArtist | 
[**PostItemsRemotesearchMusicvideo**](ItemLookupServiceAPI.md#PostItemsRemotesearchMusicvideo) | **Post** /Items/RemoteSearch/MusicVideo | 
[**PostItemsRemotesearchPerson**](ItemLookupServiceAPI.md#PostItemsRemotesearchPerson) | **Post** /Items/RemoteSearch/Person | 
[**PostItemsRemotesearchSeries**](ItemLookupServiceAPI.md#PostItemsRemotesearchSeries) | **Post** /Items/RemoteSearch/Series | 
[**PostItemsRemotesearchTrailer**](ItemLookupServiceAPI.md#PostItemsRemotesearchTrailer) | **Post** /Items/RemoteSearch/Trailer | 



## GetItemsByIdExternalidinfos

> []ExternalIdInfo GetItemsByIdExternalidinfos(ctx, id).Execute()

Gets external id infos for an item



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
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.GetItemsByIdExternalidinfos(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.GetItemsByIdExternalidinfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdExternalidinfos`: []ExternalIdInfo
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.GetItemsByIdExternalidinfos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdExternalidinfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ExternalIdInfo**](ExternalIdInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsRemotesearchImage

> GetItemsRemotesearchImage(ctx).ImageUrl(imageUrl).ProviderName(providerName).Execute()

Gets a remote image



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
	imageUrl := "imageUrl_example" // string | The image url
	providerName := "providerName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemLookupServiceAPI.GetItemsRemotesearchImage(context.Background()).ImageUrl(imageUrl).ProviderName(providerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.GetItemsRemotesearchImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsRemotesearchImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageUrl** | **string** | The image url | 
 **providerName** | **string** |  | 

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


## PostItemsMetadataReset

> PostItemsMetadataReset(ctx).ItemIds(itemIds).Execute()

Resets metadata for one or more items



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
	itemIds := "itemIds_example" // string | The item ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemLookupServiceAPI.PostItemsMetadataReset(context.Background()).ItemIds(itemIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsMetadataReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsMetadataResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemIds** | **string** | The item ids | 

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


## PostItemsRemotesearchApplyById

> PostItemsRemotesearchApplyById(ctx, id).RemoteSearchResult(remoteSearchResult).ReplaceAllImages(replaceAllImages).Execute()

Applies search criteria to an item and refreshes metadata



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
	id := "id_example" // string | The item id
	remoteSearchResult := *openapiclient.NewRemoteSearchResult() // RemoteSearchResult | RemoteSearchResult: 
	replaceAllImages := true // bool | Whether or not to replace all images (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchApplyById(context.Background(), id).RemoteSearchResult(remoteSearchResult).ReplaceAllImages(replaceAllImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchApplyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchApplyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteSearchResult** | [**RemoteSearchResult**](RemoteSearchResult.md) | RemoteSearchResult:  | 
 **replaceAllImages** | **bool** | Whether or not to replace all images | 

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


## PostItemsRemotesearchBook

> []RemoteSearchResult PostItemsRemotesearchBook(ctx).RemoteSearchQueryBookInfo(remoteSearchQueryBookInfo).Execute()





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
	remoteSearchQueryBookInfo := *openapiclient.NewRemoteSearchQueryBookInfo() // RemoteSearchQueryBookInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchBook(context.Background()).RemoteSearchQueryBookInfo(remoteSearchQueryBookInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchBook`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteSearchQueryBookInfo** | [**RemoteSearchQueryBookInfo**](RemoteSearchQueryBookInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchBoxset

> []RemoteSearchResult PostItemsRemotesearchBoxset(ctx).RemoteSearchQueryItemLookupInfo(remoteSearchQueryItemLookupInfo).Execute()





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
	remoteSearchQueryItemLookupInfo := *openapiclient.NewRemoteSearchQueryItemLookupInfo() // RemoteSearchQueryItemLookupInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchBoxset(context.Background()).RemoteSearchQueryItemLookupInfo(remoteSearchQueryItemLookupInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchBoxset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchBoxset`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchBoxset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchBoxsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteSearchQueryItemLookupInfo** | [**RemoteSearchQueryItemLookupInfo**](RemoteSearchQueryItemLookupInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchGame

> []RemoteSearchResult PostItemsRemotesearchGame(ctx).RemoteSearchQueryGameInfo(remoteSearchQueryGameInfo).Execute()





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
	remoteSearchQueryGameInfo := *openapiclient.NewRemoteSearchQueryGameInfo() // RemoteSearchQueryGameInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchGame(context.Background()).RemoteSearchQueryGameInfo(remoteSearchQueryGameInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchGame`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchGame`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteSearchQueryGameInfo** | [**RemoteSearchQueryGameInfo**](RemoteSearchQueryGameInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchMovie

> []RemoteSearchResult PostItemsRemotesearchMovie(ctx).RemoteSearchQueryMovieInfo(remoteSearchQueryMovieInfo).Execute()





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
	remoteSearchQueryMovieInfo := *openapiclient.NewRemoteSearchQueryMovieInfo() // RemoteSearchQueryMovieInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchMovie(context.Background()).RemoteSearchQueryMovieInfo(remoteSearchQueryMovieInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchMovie``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchMovie`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchMovie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteSearchQueryMovieInfo** | [**RemoteSearchQueryMovieInfo**](RemoteSearchQueryMovieInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchMusicalbum

> []RemoteSearchResult PostItemsRemotesearchMusicalbum(ctx).RemoteSearchQueryAlbumInfo(remoteSearchQueryAlbumInfo).Execute()





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
	remoteSearchQueryAlbumInfo := *openapiclient.NewRemoteSearchQueryAlbumInfo() // RemoteSearchQueryAlbumInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchMusicalbum(context.Background()).RemoteSearchQueryAlbumInfo(remoteSearchQueryAlbumInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchMusicalbum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchMusicalbum`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchMusicalbum`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchMusicalbumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteSearchQueryAlbumInfo** | [**RemoteSearchQueryAlbumInfo**](RemoteSearchQueryAlbumInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchMusicartist

> []RemoteSearchResult PostItemsRemotesearchMusicartist(ctx).RemoteSearchQueryArtistInfo(remoteSearchQueryArtistInfo).Execute()





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
	remoteSearchQueryArtistInfo := *openapiclient.NewRemoteSearchQueryArtistInfo() // RemoteSearchQueryArtistInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchMusicartist(context.Background()).RemoteSearchQueryArtistInfo(remoteSearchQueryArtistInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchMusicartist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchMusicartist`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchMusicartist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchMusicartistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteSearchQueryArtistInfo** | [**RemoteSearchQueryArtistInfo**](RemoteSearchQueryArtistInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchMusicvideo

> []RemoteSearchResult PostItemsRemotesearchMusicvideo(ctx).RemoteSearchQueryMusicVideoInfo(remoteSearchQueryMusicVideoInfo).Execute()





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
	remoteSearchQueryMusicVideoInfo := *openapiclient.NewRemoteSearchQueryMusicVideoInfo() // RemoteSearchQueryMusicVideoInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchMusicvideo(context.Background()).RemoteSearchQueryMusicVideoInfo(remoteSearchQueryMusicVideoInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchMusicvideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchMusicvideo`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchMusicvideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchMusicvideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteSearchQueryMusicVideoInfo** | [**RemoteSearchQueryMusicVideoInfo**](RemoteSearchQueryMusicVideoInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchPerson

> []RemoteSearchResult PostItemsRemotesearchPerson(ctx).RemoteSearchQueryPersonLookupInfo(remoteSearchQueryPersonLookupInfo).Execute()





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
	remoteSearchQueryPersonLookupInfo := *openapiclient.NewRemoteSearchQueryPersonLookupInfo() // RemoteSearchQueryPersonLookupInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchPerson(context.Background()).RemoteSearchQueryPersonLookupInfo(remoteSearchQueryPersonLookupInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchPerson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchPerson`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchPerson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteSearchQueryPersonLookupInfo** | [**RemoteSearchQueryPersonLookupInfo**](RemoteSearchQueryPersonLookupInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchSeries

> []RemoteSearchResult PostItemsRemotesearchSeries(ctx).RemoteSearchQuerySeriesInfo(remoteSearchQuerySeriesInfo).Execute()





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
	remoteSearchQuerySeriesInfo := *openapiclient.NewRemoteSearchQuerySeriesInfo() // RemoteSearchQuerySeriesInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchSeries(context.Background()).RemoteSearchQuerySeriesInfo(remoteSearchQuerySeriesInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchSeries`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteSearchQuerySeriesInfo** | [**RemoteSearchQuerySeriesInfo**](RemoteSearchQuerySeriesInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchTrailer

> []RemoteSearchResult PostItemsRemotesearchTrailer(ctx).RemoteSearchQueryTrailerInfo(remoteSearchQueryTrailerInfo).Execute()





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
	remoteSearchQueryTrailerInfo := *openapiclient.NewRemoteSearchQueryTrailerInfo() // RemoteSearchQueryTrailerInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchTrailer(context.Background()).RemoteSearchQueryTrailerInfo(remoteSearchQueryTrailerInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchTrailer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchTrailer`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchTrailer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchTrailerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteSearchQueryTrailerInfo** | [**RemoteSearchQueryTrailerInfo**](RemoteSearchQueryTrailerInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

