# TextSetting

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
**MaxLength** | **int32** | The maximum length the text can have. | [optional] 
**MinLength** | **int32** | The minimum length the text can have. | [optional] 
**Image** | **string** | The image url. | [optional] 
**PostMessage** | **string** | A string to be shown after the text input field. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


