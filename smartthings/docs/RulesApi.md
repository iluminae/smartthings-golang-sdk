# \RulesApi

All URIs are relative to *https://api.smartthings.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRule**](RulesApi.md#CreateRule) | **Post** /rules | Create a rule
[**DeleteRule**](RulesApi.md#DeleteRule) | **Delete** /rules/{ruleId} | Delete a rule
[**ExecuteRule**](RulesApi.md#ExecuteRule) | **Post** /rules/execute/{ruleId} | Execute a rule
[**GetRule**](RulesApi.md#GetRule) | **Get** /rules/{ruleId} | Get a Rule
[**ListRules**](RulesApi.md#ListRules) | **Get** /rules | Rules list
[**UpdateRule**](RulesApi.md#UpdateRule) | **Put** /rules/{ruleId} | Update a rule



## CreateRule

> map[string]interface{} CreateRule(ctx, authorization, locationId, request)

Create a rule

Create a rule for the location and token principal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**locationId** | **string**| The ID of the location in which to create the rule in. | 
**request** | [**RuleRequest**](RuleRequest.md)| The rule to be created. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> Rule DeleteRule(ctx, authorization, ruleId, locationId)

Delete a rule

Delete a rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**ruleId** | **string**| The rule ID | 
**locationId** | **string**| The ID of the location in which to delete the rule in. | 

### Return type

[**Rule**](Rule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteRule

> RuleExecutionResponse ExecuteRule(ctx, authorization, ruleId, locationId)

Execute a rule

Trigger Rule execution given a rule ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**ruleId** | **string**| The rule ID | 
**locationId** | **string**| The ID of the location that both the installed smart app and source are associated with. | 

### Return type

[**RuleExecutionResponse**](RuleExecutionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRule

> Rule GetRule(ctx, authorization, ruleId, locationId)

Get a Rule

Get a rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**ruleId** | **string**| The rule ID | 
**locationId** | **string**| The ID of the location to list the rules for. | 

### Return type

[**Rule**](Rule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRules

> PagedRules ListRules(ctx, authorization, locationId, optional)

Rules list

List of rules for the location for the given token principal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**locationId** | **string**| The ID of the location to list the rules for. | 
 **optional** | ***ListRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **max** | **optional.Int32**| The max number of rules to fetch | 
 **offset** | **optional.Int32**| The start index of rules to fetch | 

### Return type

[**PagedRules**](PagedRules.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRule

> Rule UpdateRule(ctx, authorization, ruleId, locationId, request)

Update a rule

Update a rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**ruleId** | **string**| The rule ID | 
**locationId** | **string**| The ID of the location in which to update the rule in. | 
**request** | [**RuleRequest**](RuleRequest.md)| The rule to be updated. | 

### Return type

[**Rule**](Rule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

