# \LocationsApi

All URIs are relative to *https://api.smartthings.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLocation**](LocationsApi.md#CreateLocation) | **Post** /locations | Create a Location.
[**DeleteLocation**](LocationsApi.md#DeleteLocation) | **Delete** /locations/{locationId} | Delete a Location.
[**GetLocation**](LocationsApi.md#GetLocation) | **Get** /locations/{locationId} | Get a Location.
[**ListLocations**](LocationsApi.md#ListLocations) | **Get** /locations | List Locations.
[**UpdateLocation**](LocationsApi.md#UpdateLocation) | **Put** /locations/{locationId} | Update a Location.



## CreateLocation

> Location CreateLocation(ctx, authorization, createLocationRequest)

Create a Location.

Create a Location for a user. We will try and create the Location geographically close to the country code provided in the request body. If we do not support Location creation in the requested country code, then the API will return a 422 error response with an error code of UnsupportedGeoRegionError. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**createLocationRequest** | [**CreateLocationRequest**](CreateLocationRequest.md)|  | 

### Return type

[**Location**](Location.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLocation

> map[string]interface{} DeleteLocation(ctx, authorization, locationId)

Delete a Location.

Delete a Location from a user's account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**locationId** | **string**| The ID of the location. | 

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


## GetLocation

> Location GetLocation(ctx, authorization, locationId)

Get a Location.

Get a specific Location from a user's account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**locationId** | **string**| The ID of the location. | 

### Return type

[**Location**](Location.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocations

> PagedLocations ListLocations(ctx, authorization)

List Locations.

List all Locations currently available in a user account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 

### Return type

[**PagedLocations**](PagedLocations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocation

> Location UpdateLocation(ctx, authorization, locationId, updateLocationRequest)

Update a Location.

All the fields in the request body are optional. Only the specified fields will be updated.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**locationId** | **string**| The ID of the location. | 
**updateLocationRequest** | [**UpdateLocationRequest**](UpdateLocationRequest.md)|  | 

### Return type

[**Location**](Location.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

