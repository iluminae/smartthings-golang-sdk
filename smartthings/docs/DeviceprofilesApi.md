# \DeviceprofilesApi

All URIs are relative to *https://api.smartthings.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceProfile**](DeviceprofilesApi.md#CreateDeviceProfile) | **Post** /deviceprofiles | Create a device profile
[**DeleteDeviceProfile**](DeviceprofilesApi.md#DeleteDeviceProfile) | **Delete** /deviceprofiles/{deviceProfileId} | Delete a device profile
[**GetDeviceProfile**](DeviceprofilesApi.md#GetDeviceProfile) | **Get** /deviceprofiles/{deviceProfileId} | GET a device profile
[**ListDeviceProfiles**](DeviceprofilesApi.md#ListDeviceProfiles) | **Get** /deviceprofiles | List all device profiles for the authenticated user
[**UpdateDeviceProfile**](DeviceprofilesApi.md#UpdateDeviceProfile) | **Put** /deviceprofiles/{deviceProfileId} | Update a device profile.



## CreateDeviceProfile

> DeviceProfile CreateDeviceProfile(ctx, authorization, optional)

Create a device profile

Create a device profile.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
 **optional** | ***CreateDeviceProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDeviceProfileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**optional.Interface of CreateDeviceProfileRequest**](CreateDeviceProfileRequest.md)| The device profile to be created. | 

### Return type

[**DeviceProfile**](DeviceProfile.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceProfile

> map[string]interface{} DeleteDeviceProfile(ctx, authorization, deviceProfileId)

Delete a device profile

Delete a device profile by ID. Admin use only.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**deviceProfileId** | **string**| the device profile ID | 

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


## GetDeviceProfile

> DeviceProfile GetDeviceProfile(ctx, authorization, deviceProfileId, optional)

GET a device profile

GET a device profile.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**deviceProfileId** | **string**| the device profile ID | 
 **optional** | ***GetDeviceProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDeviceProfileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **optional.String**| Language header representing the client&#39;s preferred language. The format of the &#x60;Accept-Language&#x60; header follows what is defined in [RFC 7231, section 5.3.5](https://tools.ietf.org/html/rfc7231#section-5.3.5) | 

### Return type

[**DeviceProfile**](DeviceProfile.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceProfiles

> PagedDeviceProfiles ListDeviceProfiles(ctx, authorization, optional)

List all device profiles for the authenticated user

List device profiles.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
 **optional** | ***ListDeviceProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeviceProfilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileId** | [**optional.Interface of []string**](string.md)| The device profiles IDs to filter the results by.  | 

### Return type

[**PagedDeviceProfiles**](PagedDeviceProfiles.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceProfile

> DeviceProfile UpdateDeviceProfile(ctx, authorization, deviceProfileId, optional)

Update a device profile.

Update a device profile. The device profile has to be in development status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**deviceProfileId** | **string**| the device profile ID | 
 **optional** | ***UpdateDeviceProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDeviceProfileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**optional.Interface of UpdateDeviceProfileRequest**](UpdateDeviceProfileRequest.md)| The device profile to be updated. | 

### Return type

[**DeviceProfile**](DeviceProfile.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

