# \AppsApi

All URIs are relative to *https://api.smartthings.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApp**](AppsApi.md#CreateApp) | **Post** /apps | Create an app.
[**DeleteApp**](AppsApi.md#DeleteApp) | **Delete** /apps/{appNameOrId} | Delete an app.
[**GenerateAppOauth**](AppsApi.md#GenerateAppOauth) | **Post** /apps/{appNameOrId}/oauth/generate | Generate an app&#39;s oauth client/secret.
[**GetApp**](AppsApi.md#GetApp) | **Get** /apps/{appNameOrId} | Get an app.
[**GetAppOauth**](AppsApi.md#GetAppOauth) | **Get** /apps/{appNameOrId}/oauth | Get an app&#39;s oauth settings.
[**GetAppSettings**](AppsApi.md#GetAppSettings) | **Get** /apps/{appNameOrId}/settings | Get settings.
[**ListApps**](AppsApi.md#ListApps) | **Get** /apps | List apps.
[**Register**](AppsApi.md#Register) | **Put** /apps/{appNameOrId}/register | Sends a confirmation request to App.
[**UpdateApp**](AppsApi.md#UpdateApp) | **Put** /apps/{appNameOrId} | Update an app.
[**UpdateAppOauth**](AppsApi.md#UpdateAppOauth) | **Put** /apps/{appNameOrId}/oauth | Update an app&#39;s oauth settings.
[**UpdateAppSettings**](AppsApi.md#UpdateAppSettings) | **Put** /apps/{appNameOrId}/settings | Update settings.
[**UpdateSignatureType**](AppsApi.md#UpdateSignatureType) | **Put** /apps/{appNameOrId}/signature-type | Update an app&#39;s signature type.



## CreateApp

> CreateAppResponse CreateApp(ctx, authorization, createOrUpdateAppRequest, optional)

Create an app.

Create an app integration.  A single developer account is allowed to contain a maximum of 100 apps.  Upon hitting that limit a 422 error response is returned with an error code of LimitError. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**createOrUpdateAppRequest** | [**CreateAppRequest**](CreateAppRequest.md)|  | 
 **optional** | ***CreateAppOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAppOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **signatureType** | **optional.String**| The Signature Type of the application. For WEBHOOK_SMART_APP only.  | 
 **requireConfirmation** | **optional.Bool**| Override default configuration to use either PING or CONFIRMATION lifecycle. For WEBHOOK_SMART_APP only.  | 

### Return type

[**CreateAppResponse**](CreateAppResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApp

> map[string]interface{} DeleteApp(ctx, authorization, appNameOrId)

Delete an app.

Delete an app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**appNameOrId** | **string**| The appName or appId field of an app. | 

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


## GenerateAppOauth

> GenerateAppOAuthResponse GenerateAppOauth(ctx, authorization, appNameOrId, generateAppOAuthRequest)

Generate an app's oauth client/secret.

When an app is first created an OAuth client/secret are automatically generated for the integration.  However, there are times when it maybe useful to re-generate a client/secret.  Such as in cases where a secret becomes compromised. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**appNameOrId** | **string**| The appName or appId field of an app. | 
**generateAppOAuthRequest** | [**GenerateAppOAuthRequest**](GenerateAppOAuthRequest.md)|  | 

### Return type

[**GenerateAppOAuthResponse**](GenerateAppOAuthResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApp

> App GetApp(ctx, authorization, appNameOrId)

Get an app.

Get a single app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**appNameOrId** | **string**| The appName or appId field of an app. | 

### Return type

[**App**](App.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppOauth

> AppOAuth GetAppOauth(ctx, authorization, appNameOrId)

Get an app's oauth settings.

Get an app's oauth settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**appNameOrId** | **string**| The appName or appId field of an app. | 

### Return type

[**AppOAuth**](AppOAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppSettings

> GetAppSettingsResponse GetAppSettings(ctx, authorization, appNameOrId)

Get settings.

Get settings for an app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**appNameOrId** | **string**| The appName or appId field of an app. | 

### Return type

[**GetAppSettingsResponse**](GetAppSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApps

> PagedApps ListApps(ctx, authorization, optional)

List apps.

List all apps configured in an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
 **optional** | ***ListAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appType** | **optional.String**| The App Type of the application. | 
 **classification** | **optional.String**| An App maybe associated to many classifications.  A classification drives how the integration is presented to the user in the SmartThings mobile clients.  These classifications include: * AUTOMATION - Denotes an integration that should display under the \&quot;Automation\&quot; tab in mobile clients. * SERVICE - Denotes an integration that is classified as a \&quot;Service\&quot;. * DEVICE - Denotes an integration that should display under the \&quot;Device\&quot; tab in mobile clients. * CONNECTED_SERVICE - Denotes an integration that should display under the \&quot;Connected Services\&quot; menu in mobile clients.  | 
 **tag** | **optional.String**| May be used to filter a resource by it&#39;s assigned user-tags.  Multiple tag query params are automatically joined with OR.  Example usage in query string: &#x60;&#x60;&#x60; ?tag:key_name&#x3D;value1&amp;tag:key_name&#x3D;value2 &#x60;&#x60;&#x60;  | 

### Return type

[**PagedApps**](PagedApps.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> map[string]interface{} Register(ctx, authorization, appNameOrId, appRegisterRequest)

Sends a confirmation request to App.

Prepares to register an App by sending the endpoint a confirmation message.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**appNameOrId** | **string**| The appName or appId field of an app. | 
**appRegisterRequest** | **map[string]interface{}**|  | 

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


## UpdateApp

> App UpdateApp(ctx, authorization, appNameOrId, updateAppRequest, optional)

Update an app.

Update an app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**appNameOrId** | **string**| The appName or appId field of an app. | 
**updateAppRequest** | [**UpdateAppRequest**](UpdateAppRequest.md)|  | 
 **optional** | ***UpdateAppOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAppOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **signatureType** | **optional.String**| The Signature Type of the application. For WEBHOOK_SMART_APP only.  | 
 **requireConfirmation** | **optional.Bool**| Override default configuration to use either PING or CONFIRMATION lifecycle. For WEBHOOK_SMART_APP only.  | 

### Return type

[**App**](App.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppOauth

> AppOAuth UpdateAppOauth(ctx, authorization, appNameOrId, updateAppOAuthRequest)

Update an app's oauth settings.

Update an app's oauth settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**appNameOrId** | **string**| The appName or appId field of an app. | 
**updateAppOAuthRequest** | [**UpdateAppOAuthRequest**](UpdateAppOAuthRequest.md)|  | 

### Return type

[**AppOAuth**](AppOAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppSettings

> UpdateAppSettingsResponse UpdateAppSettings(ctx, authorization, appNameOrId, updateAppSettingsRequest)

Update settings.

Update settings for an app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**appNameOrId** | **string**| The appName or appId  field of an app. | 
**updateAppSettingsRequest** | [**UpdateAppSettingsRequest**](UpdateAppSettingsRequest.md)|  | 

### Return type

[**UpdateAppSettingsResponse**](UpdateAppSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSignatureType

> map[string]interface{} UpdateSignatureType(ctx, authorization, appNameOrId, updateSignatureTypeRequest)

Update an app's signature type.

Updates the signature type of an App.  Signature options:   * APP_RSA - Legacy signing mechanism comprised of a public / private key generated for an App during registration.  This mechanism requires an App to download the public key and deploy along side their integration to verify the signature in the authorization header.   * ST_PADLOCK - App callbacks are signed with a SmartThings certificate which is publicly available at https://key.smartthings.com.  App's authorize callbacks by fetching the certificate over HTTPS and using it to validate the signature in the authorization header.  Note that when upgrading an App from APP_RSA to ST_PADLOCK it is recommended to implement both verification methods. This will provide the ability to seamlessly switch between mechanisms in case a rollback is needed. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**appNameOrId** | **string**| The appName or appId field of an app. | 
**updateSignatureTypeRequest** | [**UpdateSignatureTypeRequest**](UpdateSignatureTypeRequest.md)|  | 

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

