# \InstalledappsApi

All URIs are relative to *https://api.smartthings.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstalledAppEvents**](InstalledappsApi.md#CreateInstalledAppEvents) | **Post** /installedapps/{installedAppId}/events | Create Installed App events.
[**DeleteInstallation**](InstalledappsApi.md#DeleteInstallation) | **Delete** /installedapps/{installedAppId} | Delete an installed app.
[**GetInstallation**](InstalledappsApi.md#GetInstallation) | **Get** /installedapps/{installedAppId} | Get an installed app.
[**GetInstallationConfig**](InstalledappsApi.md#GetInstallationConfig) | **Get** /installedapps/{installedAppId}/configs/{configurationId} | Get an installed app configuration.
[**ListInstallationConfig**](InstalledappsApi.md#ListInstallationConfig) | **Get** /installedapps/{installedAppId}/configs | List an installed app&#39;s configurations.
[**ListInstallations**](InstalledappsApi.md#ListInstallations) | **Get** /installedapps | List installed apps.



## CreateInstalledAppEvents

> map[string]interface{} CreateInstalledAppEvents(ctx, authorization, installedAppId, createInstalledAppEventsRequest)

Create Installed App events.

Create events for an installed app.  This API allows Apps to create events to trigger custom behavior in installed apps. Requires a SmartApp token. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**installedAppId** | **string**| The ID of the installed application. | 
**createInstalledAppEventsRequest** | [**CreateInstalledAppEventsRequest**](CreateInstalledAppEventsRequest.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstallation

> DeleteInstalledAppResponse DeleteInstallation(ctx, authorization, installedAppId)

Delete an installed app.

Delete an Installed App.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**installedAppId** | **string**| The ID of the installed application. | 

### Return type

[**DeleteInstalledAppResponse**](DeleteInstalledAppResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstallation

> InstalledApp GetInstallation(ctx, authorization, installedAppId)

Get an installed app.

Fetch a single installed application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**installedAppId** | **string**| The ID of the installed application. | 

### Return type

[**InstalledApp**](InstalledApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstallationConfig

> InstallConfigurationDetail GetInstallationConfig(ctx, authorization, installedAppId, configurationId)

Get an installed app configuration.

Fetch a detailed install configuration model containing actual config entries / values.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**installedAppId** | **string**| The ID of the installed application. | 
**configurationId** | [**string**](.md)| The ID of the install configuration. | 

### Return type

[**InstallConfigurationDetail**](InstallConfigurationDetail.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstallationConfig

> PagedInstallConfigurations ListInstallationConfig(ctx, authorization, installedAppId, optional)

List an installed app's configurations.

List all configurations potentially filtered by status for an installed app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**installedAppId** | **string**| The ID of the installed application. | 
 **optional** | ***ListInstallationConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInstallationConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configurationStatus** | **optional.String**| Filter for configuration status. | 

### Return type

[**PagedInstallConfigurations**](PagedInstallConfigurations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstallations

> PagedInstalledApps ListInstallations(ctx, authorization, optional)

List installed apps.

List all installed applications within the specified locations. If no locations are provided, then list all installed apps accessible by the principle. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
 **optional** | ***ListInstallationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInstallationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **optional.String**| The ID of the location that both the installed smart app and source are associated with. | 
 **installedAppStatus** | **optional.String**| State of the Installed App. | 
 **installedAppType** | **optional.String**| Denotes the type of installed app. | 
 **tag** | **optional.String**| May be used to filter a resource by it&#39;s assigned user-tags.  Multiple tag query params are automatically joined with OR.  Example usage in query string: &#x60;&#x60;&#x60; ?tag:key_name&#x3D;value1&amp;tag:key_name&#x3D;value2 &#x60;&#x60;&#x60;  | 
 **appId** | **optional.String**| The ID of an App | 
 **modeId** | **optional.String**| The ID of the mode | 
 **deviceId** | **optional.String**| The ID of the device | 

### Return type

[**PagedInstalledApps**](PagedInstalledApps.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

