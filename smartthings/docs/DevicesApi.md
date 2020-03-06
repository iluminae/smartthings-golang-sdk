# \DevicesApi

All URIs are relative to *https://api.smartthings.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceEvents**](DevicesApi.md#CreateDeviceEvents) | **Post** /devices/{deviceId}/events | Create Device Events.
[**DeleteDevice**](DevicesApi.md#DeleteDevice) | **Delete** /devices/{deviceId} | Delete a Device.
[**ExecuteDeviceCommands**](DevicesApi.md#ExecuteDeviceCommands) | **Post** /devices/{deviceId}/commands | Execute commands on device.
[**GetDevice**](DevicesApi.md#GetDevice) | **Get** /devices/{deviceId} | Get a device&#39;s description.
[**GetDeviceComponentStatus**](DevicesApi.md#GetDeviceComponentStatus) | **Get** /devices/{deviceId}/components/{componentId}/status | Get a device component&#39;s status.
[**GetDeviceStatus**](DevicesApi.md#GetDeviceStatus) | **Get** /devices/{deviceId}/status | Get the full status of a device.
[**GetDeviceStatusByCapability**](DevicesApi.md#GetDeviceStatusByCapability) | **Get** /devices/{deviceId}/components/{componentId}/capabilities/{capabilityId}/status | Get a capability&#39;s status.
[**GetDevices**](DevicesApi.md#GetDevices) | **Get** /devices | List devices.
[**InstallDevice**](DevicesApi.md#InstallDevice) | **Post** /devices | Install a Device.
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Put** /devices/{deviceId} | Update a device.



## CreateDeviceEvents

> map[string]interface{} CreateDeviceEvents(ctx, authorization, deviceId, deviceEventRequest)

Create Device Events.

Create events for a device. When a device is managed by a SmartApp then it is responsible for creating events to update the attributes of the device in the SmartThings platform. The token must be for a SmartApp and it must be the SmartApp that created the Device. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**deviceId** | **string**| the device ID | 
**deviceEventRequest** | [**DeviceEventsRequest**](DeviceEventsRequest.md)|  | 

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


## DeleteDevice

> map[string]interface{} DeleteDevice(ctx, authorization, deviceId)

Delete a Device.

Delete a device by device id. If the token is for a SmartApp that created the device then it implicitly has permission for this api. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**deviceId** | **string**| the device ID | 

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


## ExecuteDeviceCommands

> map[string]interface{} ExecuteDeviceCommands(ctx, authorization, deviceId, executeCapabilityCommand)

Execute commands on device.

Execute commands on a device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**deviceId** | **string**| the device ID | 
**executeCapabilityCommand** | [**DeviceCommandsRequest**](DeviceCommandsRequest.md)|  | 

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


## GetDevice

> Device GetDevice(ctx, authorization, deviceId)

Get a device's description.

Get a device's description.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**deviceId** | **string**| the device ID | 

### Return type

[**Device**](Device.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceComponentStatus

> map[string]map[string]AttributeState GetDeviceComponentStatus(ctx, authorization, deviceId, componentId)

Get a device component's status.

Get the status of all attributes of a the component. The results may be filtered if the requester only has permission to view a subset of the component's capabilities. If the token is for a SmartApp that created the device then it implicitly has permission for this api. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**deviceId** | **string**| the device ID | 
**componentId** | **string**| The name of the component. | 

### Return type

[**map[string]map[string]AttributeState**](map.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceStatus

> DeviceStatus GetDeviceStatus(ctx, authorization, deviceId)

Get the full status of a device.

Get the current status of all of a device's component's attributes. The results may be filtered if the requester only has permission to view a subset of the device's components or capabilities. If the token is for a SmartApp that created the device then it implicitly has permission for this api. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**deviceId** | **string**| the device ID | 

### Return type

[**DeviceStatus**](DeviceStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceStatusByCapability

> map[string]AttributeState GetDeviceStatusByCapability(ctx, authorization, deviceId, componentId, capabilityId)

Get a capability's status.

Get the current status of a device component's capability. If the token is for a SmartApp that created the device then it implicitly has permission for this api. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**deviceId** | **string**| the device ID | 
**componentId** | **string**| The name of the component. | 
**capabilityId** | **string**| The ID of the capability | 

### Return type

[**map[string]AttributeState**](AttributeState.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> PagedDevices GetDevices(ctx, authorization, optional)

List devices.

Get a list of devices.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
 **optional** | ***GetDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **capability** | [**optional.Interface of []string**](string.md)| The device capabilities to filter the results by. The capabilities are treated as an \&quot;and\&quot; so all capabilities must be present.  | 
 **locationId** | [**optional.Interface of []string**](string.md)| The device locations to filter the results by.  | 
 **deviceId** | [**optional.Interface of []string**](string.md)| The device ids to filter the results by.  | 
 **capabilitiesMode** | **optional.String**| Treat all capability filter query params as a logical \&quot;or\&quot; or \&quot;and\&quot; with a default of \&quot;and\&quot;.  | [default to and]

### Return type

[**PagedDevices**](PagedDevices.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallDevice

> Device InstallDevice(ctx, authorization, installationRequest)

Install a Device.

Install a device. This is only available for SmartApp managed devices. The SmartApp that creates the device is responsible for handling commands for the device and updating the status of the device by creating events. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**installationRequest** | [**DeviceInstallRequest**](DeviceInstallRequest.md)| Installation Request | 

### Return type

[**Device**](Device.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> Device UpdateDevice(ctx, authorization, deviceId, updateDeviceRequest)

Update a device.

Update the properties of a device. If the token is for a SmartApp that created the device then it implicitly has permission for this api. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**deviceId** | **string**| the device ID | 
**updateDeviceRequest** | [**UpdateDeviceRequest**](UpdateDeviceRequest.md)|  | 

### Return type

[**Device**](Device.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

