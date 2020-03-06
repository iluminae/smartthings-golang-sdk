# \RoomsApi

All URIs are relative to *https://api.smartthings.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoom**](RoomsApi.md#CreateRoom) | **Post** /locations/{locationId}/rooms | Create a Room.
[**DeleteRoom**](RoomsApi.md#DeleteRoom) | **Delete** /locations/{locationId}/rooms/{roomId} | Delete a Room.
[**GetRoom**](RoomsApi.md#GetRoom) | **Get** /locations/{locationId}/rooms/{roomId} | Get a Room.
[**ListRooms**](RoomsApi.md#ListRooms) | **Get** /locations/{locationId}/rooms | List Rooms.
[**UpdateRoom**](RoomsApi.md#UpdateRoom) | **Put** /locations/{locationId}/rooms/{roomId} | Update a Room.



## CreateRoom

> Room CreateRoom(ctx, authorization, locationId, createRoomRequest)

Create a Room.

Create a Room for the Location. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**locationId** | **string**| The ID of the location. | 
**createRoomRequest** | [**CreateRoomRequest**](CreateRoomRequest.md)|  | 

### Return type

[**Room**](Room.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoom

> map[string]interface{} DeleteRoom(ctx, authorization, locationId, roomId)

Delete a Room.

Delete a Room from a location.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**locationId** | **string**| The ID of the location. | 
**roomId** | **string**| The ID of the room. | 

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


## GetRoom

> Room GetRoom(ctx, authorization, locationId, roomId)

Get a Room.

Get a specific Room.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**locationId** | **string**| The ID of the location. | 
**roomId** | **string**| The ID of the room. | 

### Return type

[**Room**](Room.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRooms

> PagedRooms ListRooms(ctx, authorization, locationId)

List Rooms.

List all Rooms currently available in a Location.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**locationId** | **string**| The ID of the location. | 

### Return type

[**PagedRooms**](PagedRooms.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoom

> Room UpdateRoom(ctx, authorization, locationId, roomId, updateRoomRequest)

Update a Room.

All the fields in the request body are optional. Only the specified fields will be updated.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**locationId** | **string**| The ID of the location. | 
**roomId** | **string**| The ID of the room. | 
**updateRoomRequest** | [**UpdateRoomRequest**](UpdateRoomRequest.md)|  | 

### Return type

[**Room**](Room.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

