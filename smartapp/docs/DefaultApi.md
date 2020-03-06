# \DefaultApi

All URIs are relative to *https://smartapps*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Execute**](DefaultApi.md#Execute) | **Post** / | Execute a third-party SmartApp.



## Execute

> ExecutionResponse Execute(ctx, executionRequest)

Execute a third-party SmartApp.

Execute a SmartThings SmartApp. Each SmartThings app execution is expected to contain the information outlined in the ExecutionRequest. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionRequest** | [**ExecutionRequest**](ExecutionRequest.md)|  | 

### Return type

[**ExecutionResponse**](ExecutionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

