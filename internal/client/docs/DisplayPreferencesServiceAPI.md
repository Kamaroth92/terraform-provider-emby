# \DisplayPreferencesServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDisplaypreferencesById**](DisplayPreferencesServiceAPI.md#GetDisplaypreferencesById) | **Get** /DisplayPreferences/{Id} | Gets a user&#39;s display preferences for an item
[**GetUsersettingsByUserid**](DisplayPreferencesServiceAPI.md#GetUsersettingsByUserid) | **Get** /UserSettings/{UserId} | Gets user settings
[**PostDisplaypreferencesByDisplaypreferencesid**](DisplayPreferencesServiceAPI.md#PostDisplaypreferencesByDisplaypreferencesid) | **Post** /DisplayPreferences/{DisplayPreferencesId} | Updates a user&#39;s display preferences for an item
[**PostUsersettingsByUserid**](DisplayPreferencesServiceAPI.md#PostUsersettingsByUserid) | **Post** /UserSettings/{UserId} | Updates a user&#39;s display preferences for an item
[**PostUsersettingsByUseridPartial**](DisplayPreferencesServiceAPI.md#PostUsersettingsByUseridPartial) | **Post** /UserSettings/{UserId}/Partial | Updates a user&#39;s display preferences for an item



## GetDisplaypreferencesById

> DisplayPreferences GetDisplaypreferencesById(ctx, id).UserId(userId).Client(client).Execute()

Gets a user's display preferences for an item



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
	userId := "userId_example" // string | User Id
	client := "client_example" // string | Client

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisplayPreferencesServiceAPI.GetDisplaypreferencesById(context.Background(), id).UserId(userId).Client(client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayPreferencesServiceAPI.GetDisplaypreferencesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisplaypreferencesById`: DisplayPreferences
	fmt.Fprintf(os.Stdout, "Response from `DisplayPreferencesServiceAPI.GetDisplaypreferencesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisplaypreferencesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User Id | 
 **client** | **string** | Client | 

### Return type

[**DisplayPreferences**](DisplayPreferences.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersettingsByUserid

> map[string]string GetUsersettingsByUserid(ctx, userId).Execute()

Gets user settings



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisplayPreferencesServiceAPI.GetUsersettingsByUserid(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayPreferencesServiceAPI.GetUsersettingsByUserid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersettingsByUserid`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `DisplayPreferencesServiceAPI.GetUsersettingsByUserid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersettingsByUseridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDisplaypreferencesByDisplaypreferencesid

> PostDisplaypreferencesByDisplaypreferencesid(ctx, displayPreferencesId).UserId(userId).DisplayPreferences(displayPreferences).Execute()

Updates a user's display preferences for an item



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
	displayPreferencesId := "displayPreferencesId_example" // string | DisplayPreferences Id
	userId := "userId_example" // string | User Id
	displayPreferences := *openapiclient.NewDisplayPreferences() // DisplayPreferences | DisplayPreferences: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DisplayPreferencesServiceAPI.PostDisplaypreferencesByDisplaypreferencesid(context.Background(), displayPreferencesId).UserId(userId).DisplayPreferences(displayPreferences).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayPreferencesServiceAPI.PostDisplaypreferencesByDisplaypreferencesid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**displayPreferencesId** | **string** | DisplayPreferences Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDisplaypreferencesByDisplaypreferencesidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User Id | 
 **displayPreferences** | [**DisplayPreferences**](DisplayPreferences.md) | DisplayPreferences:  | 

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


## PostUsersettingsByUserid

> PostUsersettingsByUserid(ctx, userId).RequestBody(requestBody).Execute()

Updates a user's display preferences for an item



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
	requestBody := []string{"Property_example"} // []string | UserSettings: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DisplayPreferencesServiceAPI.PostUsersettingsByUserid(context.Background(), userId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayPreferencesServiceAPI.PostUsersettingsByUserid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersettingsByUseridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** | UserSettings:  | 

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


## PostUsersettingsByUseridPartial

> PostUsersettingsByUseridPartial(ctx, userId).Body(body).Execute()

Updates a user's display preferences for an item



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
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DisplayPreferencesServiceAPI.PostUsersettingsByUseridPartial(context.Background(), userId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayPreferencesServiceAPI.PostUsersettingsByUseridPartial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersettingsByUseridPartialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | Binary stream | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

