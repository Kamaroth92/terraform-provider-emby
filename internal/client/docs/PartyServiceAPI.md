# \PartyServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetParties**](PartyServiceAPI.md#GetParties) | **Get** /Parties | Gets a list of active parties
[**GetPartiesInfo**](PartyServiceAPI.md#GetPartiesInfo) | **Get** /Parties/Info | Gets info about the session&#39;s current party
[**PostParties**](PartyServiceAPI.md#PostParties) | **Post** /Parties | Creates a party
[**PostPartiesByIdJoin**](PartyServiceAPI.md#PostPartiesByIdJoin) | **Post** /Parties/{Id}/Join | Joins a party
[**PostPartiesLeave**](PartyServiceAPI.md#PostPartiesLeave) | **Post** /Parties/Leave | Leaves a party



## GetParties

> GetParties(ctx).Execute()

Gets a list of active parties



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
	r, err := apiClient.PartyServiceAPI.GetParties(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartyServiceAPI.GetParties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartiesRequest struct via the builder pattern


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


## GetPartiesInfo

> SessionPartyInfoResult GetPartiesInfo(ctx).Execute()

Gets info about the session's current party



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
	resp, r, err := apiClient.PartyServiceAPI.GetPartiesInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartyServiceAPI.GetPartiesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPartiesInfo`: SessionPartyInfoResult
	fmt.Fprintf(os.Stdout, "Response from `PartyServiceAPI.GetPartiesInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartiesInfoRequest struct via the builder pattern


### Return type

[**SessionPartyInfoResult**](SessionPartyInfoResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostParties

> SessionPartyInfoResult PostParties(ctx).Execute()

Creates a party



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
	resp, r, err := apiClient.PartyServiceAPI.PostParties(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartyServiceAPI.PostParties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostParties`: SessionPartyInfoResult
	fmt.Fprintf(os.Stdout, "Response from `PartyServiceAPI.PostParties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostPartiesRequest struct via the builder pattern


### Return type

[**SessionPartyInfoResult**](SessionPartyInfoResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPartiesByIdJoin

> SessionPartyInfoResult PostPartiesByIdJoin(ctx, id).Execute()

Joins a party



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
	id := "id_example" // string | Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartyServiceAPI.PostPartiesByIdJoin(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartyServiceAPI.PostPartiesByIdJoin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPartiesByIdJoin`: SessionPartyInfoResult
	fmt.Fprintf(os.Stdout, "Response from `PartyServiceAPI.PostPartiesByIdJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPartiesByIdJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionPartyInfoResult**](SessionPartyInfoResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPartiesLeave

> PostPartiesLeave(ctx).Execute()

Leaves a party



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
	r, err := apiClient.PartyServiceAPI.PostPartiesLeave(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartyServiceAPI.PostPartiesLeave``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostPartiesLeaveRequest struct via the builder pattern


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

