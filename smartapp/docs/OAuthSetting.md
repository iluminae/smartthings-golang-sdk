# OAuthSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A developer defined configuration ID. | [optional] 
**Name** | **string** | Name of the setting to be configured. | [optional] 
**Description** | **string** | Description of the app to be configured. | [optional] 
**Required** | **bool** | Indicates if this setting is required for configuration. | [optional] [default to false]
**Disabled** | **bool** | Indicates if this setting should be disabled | [optional] [default to false]
**Type** | [**SettingType**](SettingType.md) |  | 
**DefaultValue** | **string** | A default value for the setting. | [optional] 
**UrlTemplate** | **string** | The url to use for the OAuth service. Use __SmartThingsOAuthCallback__ in the template for the callback/redirect url you need to provide to the OAuth service.  | [optional] 
**Style** | [**StyleType**](StyleType.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


