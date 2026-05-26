# \EncodingInfoServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEncodingCodecconfigurationDefaults**](EncodingInfoServiceAPI.md#GetEncodingCodecconfigurationDefaults) | **Get** /Encoding/CodecConfiguration/Defaults | Gets default codec configurations
[**GetEncodingCodecinformationVideo**](EncodingInfoServiceAPI.md#GetEncodingCodecinformationVideo) | **Get** /Encoding/CodecInformation/Video | Gets details about available video encoders and decoders
[**GetEncodingTonemapoptions**](EncodingInfoServiceAPI.md#GetEncodingTonemapoptions) | **Get** /Encoding/ToneMapOptions | Gets available tone mapping options



## GetEncodingCodecconfigurationDefaults

> []CodecConfiguration GetEncodingCodecconfigurationDefaults(ctx).Execute()

Gets default codec configurations



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
	resp, r, err := apiClient.EncodingInfoServiceAPI.GetEncodingCodecconfigurationDefaults(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncodingInfoServiceAPI.GetEncodingCodecconfigurationDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingCodecconfigurationDefaults`: []CodecConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EncodingInfoServiceAPI.GetEncodingCodecconfigurationDefaults`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingCodecconfigurationDefaultsRequest struct via the builder pattern


### Return type

[**[]CodecConfiguration**](CodecConfiguration.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncodingCodecinformationVideo

> []VideoCodecBase GetEncodingCodecinformationVideo(ctx).Execute()

Gets details about available video encoders and decoders



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
	resp, r, err := apiClient.EncodingInfoServiceAPI.GetEncodingCodecinformationVideo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncodingInfoServiceAPI.GetEncodingCodecinformationVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingCodecinformationVideo`: []VideoCodecBase
	fmt.Fprintf(os.Stdout, "Response from `EncodingInfoServiceAPI.GetEncodingCodecinformationVideo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingCodecinformationVideoRequest struct via the builder pattern


### Return type

[**[]VideoCodecBase**](VideoCodecBase.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncodingTonemapoptions

> ConfigurationToneMappingToneMapOptionsVisibility GetEncodingTonemapoptions(ctx).Execute()

Gets available tone mapping options



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
	resp, r, err := apiClient.EncodingInfoServiceAPI.GetEncodingTonemapoptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncodingInfoServiceAPI.GetEncodingTonemapoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingTonemapoptions`: ConfigurationToneMappingToneMapOptionsVisibility
	fmt.Fprintf(os.Stdout, "Response from `EncodingInfoServiceAPI.GetEncodingTonemapoptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingTonemapoptionsRequest struct via the builder pattern


### Return type

[**ConfigurationToneMappingToneMapOptionsVisibility**](ConfigurationToneMappingToneMapOptionsVisibility.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

