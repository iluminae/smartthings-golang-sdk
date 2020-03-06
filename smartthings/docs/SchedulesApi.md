# \SchedulesApi

All URIs are relative to *https://api.smartthings.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSchedule**](SchedulesApi.md#CreateSchedule) | **Post** /installedapps/{installedAppId}/schedules | Save an installed app schedule.
[**DeleteSchedule**](SchedulesApi.md#DeleteSchedule) | **Delete** /installedapps/{installedAppId}/schedules/{scheduleName} | Delete a schedule.
[**DeleteSchedules**](SchedulesApi.md#DeleteSchedules) | **Delete** /installedapps/{installedAppId}/schedules | Delete all of an installed app&#39;s schedules.
[**GetSchedule**](SchedulesApi.md#GetSchedule) | **Get** /installedapps/{installedAppId}/schedules/{scheduleName} | Get an installed app&#39;s schedule.
[**GetSchedules**](SchedulesApi.md#GetSchedules) | **Get** /installedapps/{installedAppId}/schedules | List installed app schedules.



## CreateSchedule

> Schedule CreateSchedule(ctx, installedAppId, authorization, optional)

Save an installed app schedule.

Create a schedule for an installed app. The installed app must be in the location specified and the installed app must have permission to create schedules. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAppId** | **string**| The ID of the installed application. | 
**authorization** | **string**| OAuth token | 
 **optional** | ***CreateScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateScheduleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**optional.Interface of ScheduleRequest**](ScheduleRequest.md)| The schedule to be created. | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSchedule

> map[string]interface{} DeleteSchedule(ctx, installedAppId, scheduleName, authorization)

Delete a schedule.

Delete a specific schedule for the installed app. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAppId** | **string**| The ID of the installed application. | 
**scheduleName** | **string**| The name of the schedule | 
**authorization** | **string**| OAuth token | 

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


## DeleteSchedules

> map[string]interface{} DeleteSchedules(ctx, installedAppId, authorization)

Delete all of an installed app's schedules.

Delete all schedules for the installed app. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAppId** | **string**| The ID of the installed application. | 
**authorization** | **string**| OAuth token | 

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


## GetSchedule

> Schedule GetSchedule(ctx, installedAppId, scheduleName, authorization)

Get an installed app's schedule.

Get a specific schedule for the installed app. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAppId** | **string**| The ID of the installed application. | 
**scheduleName** | **string**| The name of the schedule | 
**authorization** | **string**| OAuth token | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedules

> PagedSchedules GetSchedules(ctx, installedAppId, authorization)

List installed app schedules.

List the schedules for the installed app. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAppId** | **string**| The ID of the installed application. | 
**authorization** | **string**| OAuth token | 

### Return type

[**PagedSchedules**](PagedSchedules.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

