# ExecutionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lifecycle** | [**AppLifecycle**](AppLifecycle.md) |  | [optional] 
**ExecutionId** | **string** | This is a correlation id that is assigned to the execution that is useful for support requests. | [optional] 
**AppId** | **string** | A globally unique identifier for the application being executed. It should match the ID of the application itself. | [optional] 
**Locale** | **string** | An IETF BCP 47 language tag representing the chosen locale for this account. | [optional] 
**Version** | **string** | The version of the execution&#39;s request model. | [optional] 
**Client** | [**ClientDetails**](ClientDetails.md) |  | [optional] 
**EventData** | [**EventData**](EventData.md) |  | [optional] 
**InstallData** | [**InstallData**](InstallData.md) |  | [optional] 
**UpdateData** | [**UpdateData**](UpdateData.md) |  | [optional] 
**UninstallData** | [**UninstallData**](UninstallData.md) |  | [optional] 
**ConfigurationData** | [**ConfigurationData**](ConfigurationData.md) |  | [optional] 
**PingData** | [**PingData**](PingData.md) |  | [optional] 
**OauthCallbackData** | [**OAuthCallbackData**](OAuthCallbackData.md) |  | [optional] 
**ExecuteData** | [**ExecuteData**](ExecuteData.md) |  | [optional] 
**DashboardData** | [**DashboardData**](DashboardData.md) |  | [optional] 
**ConfirmationData** | [**ConfirmationData**](ConfirmationData.md) |  | [optional] 
**Settings** | **map[string]string** | Global settings as defined on the App. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


