# ExecutionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** |  | [optional] 
**EventData** | [**map[string]interface{}**](.md) | Empty object response for an Event lifecycle execution.  | [optional] 
**InstallData** | [**map[string]interface{}**](.md) | Empty object response for an Install lifecycle execution.  | [optional] 
**UpdateData** | [**map[string]interface{}**](.md) | Empty object response for an Update lifecycle execution.  | [optional] 
**UninstallData** | [**map[string]interface{}**](.md) | Empty object response for an Uninstall lifecycle execution.  | [optional] 
**ConfigurationData** | [**ConfigurationResponseData**](ConfigurationResponseData.md) |  | [optional] 
**PingData** | [**PingResponseData**](PingResponseData.md) |  | [optional] 
**OauthCallbackData** | [**map[string]interface{}**](.md) | Empty object response for an OAuth Callback lifecycle execution.  | [optional] 
**ExecuteData** | [**map[string]interface{}**](.md) | A custom JSON formatted response object with a maximum size of 64kb.  | [optional] 
**DashboardData** | [**DashboardResponseData**](DashboardResponseData.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


