# UpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | **string** | An OAuth token to use when calling into SmartThings API&#39;s. | 
**RefreshToken** | **string** | A refresh token which maybe used to obtain authorization to SmartThings API after expiration of the authToken. An integration will need to use this refreshToken to support calling the SmartThings API outside the context of an event.  | [optional] 
**InstalledApp** | [**InstalledApp**](InstalledApp.md) |  | 
**PreviousConfig** | [**map[string][]ConfigEntry**](array.md) | A map of configurations for an Installed App.  The map &#39;key&#39; is the configuration name and the &#39;value&#39; is an array of strings.  | 
**PreviousPermissions** | **[]string** | A list of permissions associated with this execution. See &#x60;securityDefinitions&#x60; for more information. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


