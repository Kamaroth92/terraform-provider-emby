# \BackupApiAPI

All URIs are relative to *http://emby.home.barriball.id.au/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackuprestoreBackupinfo**](BackupApiAPI.md#GetBackuprestoreBackupinfo) | **Get** /BackupRestore/BackupInfo | 
[**PostBackuprestoreRestore**](BackupApiAPI.md#PostBackuprestoreRestore) | **Post** /BackupRestore/Restore | 
[**PostBackuprestoreRestoredata**](BackupApiAPI.md#PostBackuprestoreRestoredata) | **Post** /BackupRestore/RestoreData | 



## GetBackuprestoreBackupinfo

> MBBackupApiAllBackupsInfo GetBackuprestoreBackupinfo(ctx).Execute()





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
	resp, r, err := apiClient.BackupApiAPI.GetBackuprestoreBackupinfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupApiAPI.GetBackuprestoreBackupinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackuprestoreBackupinfo`: MBBackupApiAllBackupsInfo
	fmt.Fprintf(os.Stdout, "Response from `BackupApiAPI.GetBackuprestoreBackupinfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackuprestoreBackupinfoRequest struct via the builder pattern


### Return type

[**MBBackupApiAllBackupsInfo**](MBBackupApiAllBackupsInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBackuprestoreRestore

> PostBackuprestoreRestore(ctx).MBBackupApiRestoreOptions(mBBackupApiRestoreOptions).Execute()





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
	mBBackupApiRestoreOptions := *openapiclient.NewMBBackupApiRestoreOptions() // MBBackupApiRestoreOptions | RestoreOptions: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupApiAPI.PostBackuprestoreRestore(context.Background()).MBBackupApiRestoreOptions(mBBackupApiRestoreOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupApiAPI.PostBackuprestoreRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBackuprestoreRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mBBackupApiRestoreOptions** | [**MBBackupApiRestoreOptions**](MBBackupApiRestoreOptions.md) | RestoreOptions:  | 

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


## PostBackuprestoreRestoredata

> PostBackuprestoreRestoredata(ctx).MBBackupApiDataRestoreOptions(mBBackupApiDataRestoreOptions).Execute()





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
	mBBackupApiDataRestoreOptions := *openapiclient.NewMBBackupApiDataRestoreOptions() // MBBackupApiDataRestoreOptions | DataRestoreOptions: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupApiAPI.PostBackuprestoreRestoredata(context.Background()).MBBackupApiDataRestoreOptions(mBBackupApiDataRestoreOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupApiAPI.PostBackuprestoreRestoredata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBackuprestoreRestoredataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mBBackupApiDataRestoreOptions** | [**MBBackupApiDataRestoreOptions**](MBBackupApiDataRestoreOptions.md) | DataRestoreOptions:  | 

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

