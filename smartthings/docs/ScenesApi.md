# \ScenesApi

All URIs are relative to *https://api.smartthings.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteScene**](ScenesApi.md#ExecuteScene) | **Post** /scenes/{sceneId}/execute | Execute Scene
[**ListScenes**](ScenesApi.md#ListScenes) | **Get** /scenes | List Scenes



## ExecuteScene

> StandardSuccessResponse ExecuteScene(ctx, authorization, sceneId, optional)

Execute Scene

Execute a Scene by id for the logged in user and given locationId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
**sceneId** | **string**| The ID of the Scene. | 
 **optional** | ***ExecuteSceneOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExecuteSceneOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locationId** | **optional.String**| The location of a scene. | 

### Return type

[**StandardSuccessResponse**](StandardSuccessResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScenes

> ScenePagedResult ListScenes(ctx, authorization, optional)

List Scenes

Fetch a list of Scenes for the logged in user filtered by locationIds. If no locationId is sent, return scenes for all available locations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| OAuth token | 
 **optional** | ***ListScenesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListScenesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **optional.String**| The location of a scene. | 

### Return type

[**ScenePagedResult**](ScenePagedResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.smartthings+json, application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

