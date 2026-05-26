# \UserNotificationsServiceAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationsServicesDefaults**](UserNotificationsServiceAPI.md#GetNotificationsServicesDefaults) | **Get** /Notifications/Services/Defaults | Gets default notification info
[**PostNotificationsServicesTest**](UserNotificationsServiceAPI.md#PostNotificationsServicesTest) | **Post** /Notifications/Services/Test | Sends a test notification



## GetNotificationsServicesDefaults

> UserNotificationInfo GetNotificationsServicesDefaults(ctx).Execute()

Gets default notification info



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
	resp, r, err := apiClient.UserNotificationsServiceAPI.GetNotificationsServicesDefaults(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserNotificationsServiceAPI.GetNotificationsServicesDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsServicesDefaults`: UserNotificationInfo
	fmt.Fprintf(os.Stdout, "Response from `UserNotificationsServiceAPI.GetNotificationsServicesDefaults`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsServicesDefaultsRequest struct via the builder pattern


### Return type

[**UserNotificationInfo**](UserNotificationInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotificationsServicesTest

> PostNotificationsServicesTest(ctx).UserNotificationInfo(userNotificationInfo).Execute()

Sends a test notification



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
	userNotificationInfo := *openapiclient.NewUserNotificationInfo() // UserNotificationInfo | UserNotificationInfo: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserNotificationsServiceAPI.PostNotificationsServicesTest(context.Background()).UserNotificationInfo(userNotificationInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserNotificationsServiceAPI.PostNotificationsServicesTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNotificationsServicesTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userNotificationInfo** | [**UserNotificationInfo**](UserNotificationInfo.md) | UserNotificationInfo:  | 

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

