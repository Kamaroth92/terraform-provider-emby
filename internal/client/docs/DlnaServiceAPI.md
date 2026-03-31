# \DlnaServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDlnaProfilesById**](DlnaServiceAPI.md#DeleteDlnaProfilesById) | **Delete** /Dlna/Profiles/{Id} | Deletes a profile
[**GetDlnaProfileinfos**](DlnaServiceAPI.md#GetDlnaProfileinfos) | **Get** /Dlna/ProfileInfos | Gets a list of profiles
[**GetDlnaProfilesById**](DlnaServiceAPI.md#GetDlnaProfilesById) | **Get** /Dlna/Profiles/{Id} | Gets a single profile
[**GetDlnaProfilesDefault**](DlnaServiceAPI.md#GetDlnaProfilesDefault) | **Get** /Dlna/Profiles/Default | Gets the default profile
[**PostDlnaProfiles**](DlnaServiceAPI.md#PostDlnaProfiles) | **Post** /Dlna/Profiles | Creates a profile
[**PostDlnaProfilesById**](DlnaServiceAPI.md#PostDlnaProfilesById) | **Post** /Dlna/Profiles/{Id} | Updates a profile



## DeleteDlnaProfilesById

> DeleteDlnaProfilesById(ctx, id).Execute()

Deletes a profile



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
	id := "id_example" // string | Profile Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServiceAPI.DeleteDlnaProfilesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServiceAPI.DeleteDlnaProfilesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDlnaProfilesByIdRequest struct via the builder pattern


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


## GetDlnaProfileinfos

> []DlnaProfilesDlnaProfile GetDlnaProfileinfos(ctx).Execute()

Gets a list of profiles



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServiceAPI.GetDlnaProfileinfos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServiceAPI.GetDlnaProfileinfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDlnaProfileinfos`: []DlnaProfilesDlnaProfile
	fmt.Fprintf(os.Stdout, "Response from `DlnaServiceAPI.GetDlnaProfileinfos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlnaProfileinfosRequest struct via the builder pattern


### Return type

[**[]DlnaProfilesDlnaProfile**](DlnaProfilesDlnaProfile.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDlnaProfilesById

> DlnaProfilesDlnaProfile GetDlnaProfilesById(ctx, id).Execute()

Gets a single profile



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
	id := "id_example" // string | Profile Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServiceAPI.GetDlnaProfilesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServiceAPI.GetDlnaProfilesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDlnaProfilesById`: DlnaProfilesDlnaProfile
	fmt.Fprintf(os.Stdout, "Response from `DlnaServiceAPI.GetDlnaProfilesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlnaProfilesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DlnaProfilesDlnaProfile**](DlnaProfilesDlnaProfile.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDlnaProfilesDefault

> DlnaProfilesDlnaProfile GetDlnaProfilesDefault(ctx).Execute()

Gets the default profile



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServiceAPI.GetDlnaProfilesDefault(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServiceAPI.GetDlnaProfilesDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDlnaProfilesDefault`: DlnaProfilesDlnaProfile
	fmt.Fprintf(os.Stdout, "Response from `DlnaServiceAPI.GetDlnaProfilesDefault`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlnaProfilesDefaultRequest struct via the builder pattern


### Return type

[**DlnaProfilesDlnaProfile**](DlnaProfilesDlnaProfile.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDlnaProfiles

> PostDlnaProfiles(ctx).DlnaProfilesDlnaProfile(dlnaProfilesDlnaProfile).Execute()

Creates a profile



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
	dlnaProfilesDlnaProfile := *openapiclient.NewDlnaProfilesDlnaProfile() // DlnaProfilesDlnaProfile | DlnaProfile: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServiceAPI.PostDlnaProfiles(context.Background()).DlnaProfilesDlnaProfile(dlnaProfilesDlnaProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServiceAPI.PostDlnaProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDlnaProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dlnaProfilesDlnaProfile** | [**DlnaProfilesDlnaProfile**](DlnaProfilesDlnaProfile.md) | DlnaProfile:  | 

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


## PostDlnaProfilesById

> PostDlnaProfilesById(ctx, id).DlnaProfilesDlnaProfile(dlnaProfilesDlnaProfile).Execute()

Updates a profile



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
	dlnaProfilesDlnaProfile := *openapiclient.NewDlnaProfilesDlnaProfile() // DlnaProfilesDlnaProfile | DlnaProfile: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServiceAPI.PostDlnaProfilesById(context.Background(), id).DlnaProfilesDlnaProfile(dlnaProfilesDlnaProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServiceAPI.PostDlnaProfilesById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostDlnaProfilesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dlnaProfilesDlnaProfile** | [**DlnaProfilesDlnaProfile**](DlnaProfilesDlnaProfile.md) | DlnaProfile:  | 

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

