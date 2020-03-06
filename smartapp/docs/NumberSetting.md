# NumberSetting

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
**Max** | **int32** | The maximum inclusive value the number can be set to. | [optional] 
**Min** | **int32** | The minumum inclusive value the number can be set to. | [optional] 
**Step** | **int32** | The increment to step by. | [optional] 
**Style** | **string** | The way to style the number input. | [optional] 
**Image** | **string** | The image url. | [optional] 
**PostMessage** | **string** | A string to be shown after the text input field. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


