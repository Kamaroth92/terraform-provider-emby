# \UserServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsersById**](UserServiceAPI.md#DeleteUsersById) | **Delete** /Users/{Id} | Deletes a user
[**DeleteUsersByIdTrackselectionsByTracktype**](UserServiceAPI.md#DeleteUsersByIdTrackselectionsByTracktype) | **Delete** /Users/{Id}/TrackSelections/{TrackType} | Clears audio or subtitle track selections for a user
[**GetUsersById**](UserServiceAPI.md#GetUsersById) | **Get** /Users/{Id} | Gets a user by Id
[**GetUsersByUseridTypedsettingsByKey**](UserServiceAPI.md#GetUsersByUseridTypedsettingsByKey) | **Get** /Users/{UserId}/TypedSettings/{Key} | Gets a typed user setting
[**GetUsersItemaccess**](UserServiceAPI.md#GetUsersItemaccess) | **Get** /Users/ItemAccess | Gets a list of users
[**GetUsersPrefixes**](UserServiceAPI.md#GetUsersPrefixes) | **Get** /Users/Prefixes | Gets a list of users
[**GetUsersPublic**](UserServiceAPI.md#GetUsersPublic) | **Get** /Users/Public | Gets a list of publicly visible users for display on a login screen.
[**GetUsersQuery**](UserServiceAPI.md#GetUsersQuery) | **Get** /Users/Query | Gets a list of users
[**PostUsersAuthenticatebyname**](UserServiceAPI.md#PostUsersAuthenticatebyname) | **Post** /Users/AuthenticateByName | Authenticates a user
[**PostUsersById**](UserServiceAPI.md#PostUsersById) | **Post** /Users/{Id} | Updates a user
[**PostUsersByIdAuthenticate**](UserServiceAPI.md#PostUsersByIdAuthenticate) | **Post** /Users/{Id}/Authenticate | Authenticates a user
[**PostUsersByIdConfiguration**](UserServiceAPI.md#PostUsersByIdConfiguration) | **Post** /Users/{Id}/Configuration | Updates a user configuration
[**PostUsersByIdConfigurationPartial**](UserServiceAPI.md#PostUsersByIdConfigurationPartial) | **Post** /Users/{Id}/Configuration/Partial | Updates a user configuration
[**PostUsersByIdDelete**](UserServiceAPI.md#PostUsersByIdDelete) | **Post** /Users/{Id}/Delete | Deletes a user
[**PostUsersByIdPassword**](UserServiceAPI.md#PostUsersByIdPassword) | **Post** /Users/{Id}/Password | Updates a user&#39;s password
[**PostUsersByIdPolicy**](UserServiceAPI.md#PostUsersByIdPolicy) | **Post** /Users/{Id}/Policy | Updates a user policy
[**PostUsersByIdTrackselectionsByTracktypeDelete**](UserServiceAPI.md#PostUsersByIdTrackselectionsByTracktypeDelete) | **Post** /Users/{Id}/TrackSelections/{TrackType}/Delete | Clears audio or subtitle track selections for a user
[**PostUsersByUseridTypedsettingsByKey**](UserServiceAPI.md#PostUsersByUseridTypedsettingsByKey) | **Post** /Users/{UserId}/TypedSettings/{Key} | Updates a typed user setting
[**PostUsersForgotpassword**](UserServiceAPI.md#PostUsersForgotpassword) | **Post** /Users/ForgotPassword | Initiates the forgot password process for a local user
[**PostUsersForgotpasswordPin**](UserServiceAPI.md#PostUsersForgotpasswordPin) | **Post** /Users/ForgotPassword/Pin | Redeems a forgot password pin
[**PostUsersNew**](UserServiceAPI.md#PostUsersNew) | **Post** /Users/New | Creates a user



## DeleteUsersById

> DeleteUsersById(ctx, id).Execute()

Deletes a user



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.DeleteUsersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.DeleteUsersById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUsersByIdRequest struct via the builder pattern


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


## DeleteUsersByIdTrackselectionsByTracktype

> DeleteUsersByIdTrackselectionsByTracktype(ctx, id, trackType).Execute()

Clears audio or subtitle track selections for a user



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
	id := "id_example" // string | 
	trackType := "trackType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.DeleteUsersByIdTrackselectionsByTracktype(context.Background(), id, trackType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.DeleteUsersByIdTrackselectionsByTracktype``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**trackType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersByIdTrackselectionsByTracktypeRequest struct via the builder pattern


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


## GetUsersById

> UserDto GetUsersById(ctx, id).Execute()

Gets a user by Id



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.GetUsersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.GetUsersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersById`: UserDto
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.GetUsersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserDto**](UserDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersByUseridTypedsettingsByKey

> GetUsersByUseridTypedsettingsByKey(ctx, key, userId).Execute()

Gets a typed user setting



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
	key := "key_example" // string | Key
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.GetUsersByUseridTypedsettingsByKey(context.Background(), key, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.GetUsersByUseridTypedsettingsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByUseridTypedsettingsByKeyRequest struct via the builder pattern


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


## GetUsersItemaccess

> QueryResultUserDto GetUsersItemaccess(ctx).IsHidden(isHidden).IsDisabled(isDisabled).StartIndex(startIndex).Limit(limit).NameStartsWithOrGreater(nameStartsWithOrGreater).SortOrder(sortOrder).Execute()

Gets a list of users



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
	isHidden := true // bool | Optional filter by IsHidden=true or false (optional)
	isDisabled := true // bool | Optional filter by IsDisabled=true or false (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.GetUsersItemaccess(context.Background()).IsHidden(isHidden).IsDisabled(isDisabled).StartIndex(startIndex).Limit(limit).NameStartsWithOrGreater(nameStartsWithOrGreater).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.GetUsersItemaccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersItemaccess`: QueryResultUserDto
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.GetUsersItemaccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersItemaccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **bool** | Optional filter by IsHidden&#x3D;true or false | 
 **isDisabled** | **bool** | Optional filter by IsDisabled&#x3D;true or false | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 

### Return type

[**QueryResultUserDto**](QueryResultUserDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersPrefixes

> []NameIdPair GetUsersPrefixes(ctx).IsHidden(isHidden).IsDisabled(isDisabled).StartIndex(startIndex).Limit(limit).NameStartsWithOrGreater(nameStartsWithOrGreater).SortOrder(sortOrder).Execute()

Gets a list of users



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
	isHidden := true // bool | Optional filter by IsHidden=true or false (optional)
	isDisabled := true // bool | Optional filter by IsDisabled=true or false (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.GetUsersPrefixes(context.Background()).IsHidden(isHidden).IsDisabled(isDisabled).StartIndex(startIndex).Limit(limit).NameStartsWithOrGreater(nameStartsWithOrGreater).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.GetUsersPrefixes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersPrefixes`: []NameIdPair
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.GetUsersPrefixes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersPrefixesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **bool** | Optional filter by IsHidden&#x3D;true or false | 
 **isDisabled** | **bool** | Optional filter by IsDisabled&#x3D;true or false | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 

### Return type

[**[]NameIdPair**](NameIdPair.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersPublic

> []UserDto GetUsersPublic(ctx).Execute()

Gets a list of publicly visible users for display on a login screen.



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
	resp, r, err := apiClient.UserServiceAPI.GetUsersPublic(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.GetUsersPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersPublic`: []UserDto
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.GetUsersPublic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersPublicRequest struct via the builder pattern


### Return type

[**[]UserDto**](UserDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersQuery

> QueryResultUserDto GetUsersQuery(ctx).IsHidden(isHidden).IsDisabled(isDisabled).StartIndex(startIndex).Limit(limit).NameStartsWithOrGreater(nameStartsWithOrGreater).SortOrder(sortOrder).Execute()

Gets a list of users



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
	isHidden := true // bool | Optional filter by IsHidden=true or false (optional)
	isDisabled := true // bool | Optional filter by IsDisabled=true or false (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.GetUsersQuery(context.Background()).IsHidden(isHidden).IsDisabled(isDisabled).StartIndex(startIndex).Limit(limit).NameStartsWithOrGreater(nameStartsWithOrGreater).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.GetUsersQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersQuery`: QueryResultUserDto
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.GetUsersQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **bool** | Optional filter by IsHidden&#x3D;true or false | 
 **isDisabled** | **bool** | Optional filter by IsDisabled&#x3D;true or false | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 

### Return type

[**QueryResultUserDto**](QueryResultUserDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersAuthenticatebyname

> AuthenticationAuthenticationResult PostUsersAuthenticatebyname(ctx).XEmbyAuthorization(xEmbyAuthorization).AuthenticateUserByName(authenticateUserByName).Execute()

Authenticates a user



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
	xEmbyAuthorization := "xEmbyAuthorization_example" // string | The authorization header can be either named 'Authorization' or 'X-Emby-Authorization'.    It must be of the following schema:     Emby UserId=\"(guid)\", Client=\"(string)\", Device=\"(string)\", DeviceId=\"(string)\", Version=\"string\", Token=\"(string)\"     Please consult the documentation for further details.
	authenticateUserByName := *openapiclient.NewAuthenticateUserByName() // AuthenticateUserByName | AuthenticateUserByName

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.PostUsersAuthenticatebyname(context.Background()).XEmbyAuthorization(xEmbyAuthorization).AuthenticateUserByName(authenticateUserByName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersAuthenticatebyname``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersAuthenticatebyname`: AuthenticationAuthenticationResult
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.PostUsersAuthenticatebyname`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersAuthenticatebynameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEmbyAuthorization** | **string** | The authorization header can be either named &#39;Authorization&#39; or &#39;X-Emby-Authorization&#39;.    It must be of the following schema:     Emby UserId&#x3D;\&quot;(guid)\&quot;, Client&#x3D;\&quot;(string)\&quot;, Device&#x3D;\&quot;(string)\&quot;, DeviceId&#x3D;\&quot;(string)\&quot;, Version&#x3D;\&quot;string\&quot;, Token&#x3D;\&quot;(string)\&quot;     Please consult the documentation for further details. | 
 **authenticateUserByName** | [**AuthenticateUserByName**](AuthenticateUserByName.md) | AuthenticateUserByName | 

### Return type

[**AuthenticationAuthenticationResult**](AuthenticationAuthenticationResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersById

> PostUsersById(ctx, id).UserDto(userDto).Execute()

Updates a user



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
	id := "id_example" // string | 
	userDto := *openapiclient.NewUserDto() // UserDto | UserDto: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersById(context.Background(), id).UserDto(userDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userDto** | [**UserDto**](UserDto.md) | UserDto:  | 

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


## PostUsersByIdAuthenticate

> AuthenticationAuthenticationResult PostUsersByIdAuthenticate(ctx, id).AuthenticateUser(authenticateUser).Execute()

Authenticates a user



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
	id := "id_example" // string | 
	authenticateUser := *openapiclient.NewAuthenticateUser() // AuthenticateUser | AuthenticateUser

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.PostUsersByIdAuthenticate(context.Background(), id).AuthenticateUser(authenticateUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersByIdAuthenticate`: AuthenticationAuthenticationResult
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.PostUsersByIdAuthenticate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByIdAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticateUser** | [**AuthenticateUser**](AuthenticateUser.md) | AuthenticateUser | 

### Return type

[**AuthenticationAuthenticationResult**](AuthenticationAuthenticationResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersByIdConfiguration

> PostUsersByIdConfiguration(ctx, id).UserConfiguration(userConfiguration).Execute()

Updates a user configuration



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
	id := "id_example" // string | 
	userConfiguration := *openapiclient.NewUserConfiguration() // UserConfiguration | UserConfiguration: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByIdConfiguration(context.Background(), id).UserConfiguration(userConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userConfiguration** | [**UserConfiguration**](UserConfiguration.md) | UserConfiguration:  | 

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


## PostUsersByIdConfigurationPartial

> PostUsersByIdConfigurationPartial(ctx, id).Body(body).Execute()

Updates a user configuration



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
	id := "id_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByIdConfigurationPartial(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdConfigurationPartial``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdConfigurationPartialRequest struct via the builder pattern


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


## PostUsersByIdDelete

> PostUsersByIdDelete(ctx, id).Execute()

Deletes a user



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdDeleteRequest struct via the builder pattern


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


## PostUsersByIdPassword

> PostUsersByIdPassword(ctx, id).UpdateUserPassword(updateUserPassword).Execute()

Updates a user's password



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
	id := "id_example" // string | 
	updateUserPassword := *openapiclient.NewUpdateUserPassword() // UpdateUserPassword | UpdateUserPassword

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByIdPassword(context.Background(), id).UpdateUserPassword(updateUserPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdPassword``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserPassword** | [**UpdateUserPassword**](UpdateUserPassword.md) | UpdateUserPassword | 

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


## PostUsersByIdPolicy

> PostUsersByIdPolicy(ctx, id).UserPolicy(userPolicy).Execute()

Updates a user policy



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
	id := "id_example" // string | 
	userPolicy := *openapiclient.NewUserPolicy() // UserPolicy | UserPolicy: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByIdPolicy(context.Background(), id).UserPolicy(userPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdPolicy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userPolicy** | [**UserPolicy**](UserPolicy.md) | UserPolicy:  | 

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


## PostUsersByIdTrackselectionsByTracktypeDelete

> PostUsersByIdTrackselectionsByTracktypeDelete(ctx, id, trackType).Execute()

Clears audio or subtitle track selections for a user



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
	id := "id_example" // string | 
	trackType := "trackType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByIdTrackselectionsByTracktypeDelete(context.Background(), id, trackType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdTrackselectionsByTracktypeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**trackType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByIdTrackselectionsByTracktypeDeleteRequest struct via the builder pattern


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


## PostUsersByUseridTypedsettingsByKey

> PostUsersByUseridTypedsettingsByKey(ctx, userId, key).Body(body).Execute()

Updates a typed user setting



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
	userId := "userId_example" // string | 
	key := "key_example" // string | Key
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByUseridTypedsettingsByKey(context.Background(), userId, key).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByUseridTypedsettingsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**key** | **string** | Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridTypedsettingsByKeyRequest struct via the builder pattern


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


## PostUsersForgotpassword

> ForgotPasswordResult PostUsersForgotpassword(ctx).ForgotPassword(forgotPassword).Execute()

Initiates the forgot password process for a local user



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
	forgotPassword := *openapiclient.NewForgotPassword() // ForgotPassword | ForgotPassword

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.PostUsersForgotpassword(context.Background()).ForgotPassword(forgotPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersForgotpassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersForgotpassword`: ForgotPasswordResult
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.PostUsersForgotpassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersForgotpasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgotPassword** | [**ForgotPassword**](ForgotPassword.md) | ForgotPassword | 

### Return type

[**ForgotPasswordResult**](ForgotPasswordResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersForgotpasswordPin

> PinRedeemResult PostUsersForgotpasswordPin(ctx).ForgotPasswordPin(forgotPasswordPin).Execute()

Redeems a forgot password pin



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
	forgotPasswordPin := *openapiclient.NewForgotPasswordPin() // ForgotPasswordPin | ForgotPasswordPin

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.PostUsersForgotpasswordPin(context.Background()).ForgotPasswordPin(forgotPasswordPin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersForgotpasswordPin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersForgotpasswordPin`: PinRedeemResult
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.PostUsersForgotpasswordPin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersForgotpasswordPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgotPasswordPin** | [**ForgotPasswordPin**](ForgotPasswordPin.md) | ForgotPasswordPin | 

### Return type

[**PinRedeemResult**](PinRedeemResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersNew

> UserDto PostUsersNew(ctx).CreateUserByName(createUserByName).Execute()

Creates a user



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
	createUserByName := *openapiclient.NewCreateUserByName() // CreateUserByName | CreateUserByName

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.PostUsersNew(context.Background()).CreateUserByName(createUserByName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersNew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersNew`: UserDto
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.PostUsersNew`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserByName** | [**CreateUserByName**](CreateUserByName.md) | CreateUserByName | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

