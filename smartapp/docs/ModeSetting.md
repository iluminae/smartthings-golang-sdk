# ModeSetting

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
**Multiple** | **bool** | Indicates if this enum setting can have multiple values. | [optional] [default to false]
**CloseOnSelection** | **bool** | Indicates if this input should close on selection. | [optional] [default to true]
**SubmitOnChange** | **bool** | Indicates if this input should refresh configs after a change in value. | [optional] [default to false]
**Style** | [**StyleType**](StyleType.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


