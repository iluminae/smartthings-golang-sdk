# \SubscriptionsApi

All URIs are relative to *https://api.smartthings.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllSubscriptions**](SubscriptionsApi.md#DeleteAllSubscriptions) | **Delete** /installedapps/{installedAppId}/subscriptions | Delete all of an installed app&#39;s subscriptions.
[**DeleteSubscription**](SubscriptionsApi.md#DeleteSubscription) | **Delete** /installedapps/{installedAppId}/subscriptions/{subscriptionId} | Delete an installed app&#39;s subscription.
[**GetSubscription**](SubscriptionsApi.md#GetSubscription) | **Get** /installedapps/{installedAppId}/subscriptions/{subscriptionId} | Get an installed app&#39;s subscription.
[**ListSubscriptions**](SubscriptionsApi.md#ListSubscriptions) | **Get** /installedapps/{installedAppId}/subscriptions | List an installed app&#39;s subscriptions.
[**SaveSubscription**](SubscriptionsApi.md#SaveSubscription) | **Post** /installedapps/{installedAppId}/subscriptions | Create a subscription for an installed app.



## DeleteAllSubscriptions

> SubscriptionDelete DeleteAllSubscriptions(ctx, installedAppId, authorization, optional)

Delete all of an installed app's subscriptions.

Delete all subscriptions for the installed app. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAppId** | **string**| The ID of the installed application. | 
**authorization** | **string**| OAuth token | 
 **optional** | ***DeleteAllSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteAllSubscriptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceId** | **optional.String**| Limit deletion to subscriptions for a particular device. | 
 **modeId** | **optional.String**| Limit deletion to subscriptions for a particular mode or deletes parent if last mode | 

### Return type

[**SubscriptionDelete**](SubscriptionDelete.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> SubscriptionDelete DeleteSubscription(ctx, installedAppId, subscriptionId, authorization)

Delete an installed app's subscription.

Delete a specific subscription for the installed app. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAppId** | **string**| The ID of the installed application. | 
**subscriptionId** | **string**| The ID of the subscription | 
**authorization** | **string**| OAuth token | 

### Return type

[**SubscriptionDelete**](SubscriptionDelete.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> Subscription GetSubscription(ctx, installedAppId, subscriptionId, authorization)

Get an installed app's subscription.

Get a specific subscription for the installed app. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAppId** | **string**| The ID of the installed application. | 
**subscriptionId** | **string**| The ID of the subscription | 
**authorization** | **string**| OAuth token | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptions

> PagedSubscriptions ListSubscriptions(ctx, installedAppId, authorization)

List an installed app's subscriptions.

List the subscriptions for the installed app. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAppId** | **string**| The ID of the installed application. | 
**authorization** | **string**| OAuth token | 

### Return type

[**PagedSubscriptions**](PagedSubscriptions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveSubscription

> Subscription SaveSubscription(ctx, installedAppId, authorization, optional)

Create a subscription for an installed app.

Create a subscription to a type of event from the specified source. Both the source and the installed app must be in the location specified and the installed app must have read access to the event being subscribed to. An installed app is only permitted to created 20 subscriptions.  ### Authorization scopes For installed app principal: * installed app id matches the incoming request installed app id location must match the installed app location  | Subscription Type  | Scope required                                                                         | | ------------------ | ---------------------------------------------------------------------------------------| | DEVICE             | `r:devices:$deviceId`                                                                  | | CAPABILITY         | `r:devices:*:*:$capability` or `r:devices:*`,                                          | | MODE               | `r:locations:$locationId` or `r:locations:*`                                           | | DEVICE_LIFECYCLE   | `r:devices:$deviceId` or `r:devices:*`                                                 | | DEVICE_HEALTH      | `r:devices:$deviceId` or `r:devices:*`                                                 | | SECURITY_ARM_STATE | `r:security:locations:$locationId:armstate` or `r:security:locations:*:armstate`       | | HUB_HEALTH         | `r:hubs`                                                                               | | SCENE_LIFECYCLE    | `r:scenes:*`                                                                           | 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAppId** | **string**| The ID of the installed application. | 
**authorization** | **string**| OAuth token | 
 **optional** | ***SaveSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SaveSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**optional.Interface of SubscriptionRequest**](SubscriptionRequest.md)| The Subscription to be created. | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

