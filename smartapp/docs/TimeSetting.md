# TimeSetting

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
**Image** | **string** | The image url. | [optional] 
**Max** | [**time.Time**](time.Time.md) | The maximum inclusive value the time can be set to (only the time will be used out of the date time). | [optional] 
**Min** | [**time.Time**](time.Time.md) | The minimum inclusive value the time can be set to (only the time will be used out of the date time). | [optional] 
**SubmitOnChange** | **bool** | Indicates if this input should refresh configs after a change in value. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


