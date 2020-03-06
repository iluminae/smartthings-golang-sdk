# SceneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SceneName** | **string** | The user-defined name of the Scene | 
**SceneIcon** | **string** | The name of the icon | [optional] 
**SceneColor** | **string** | The color of the icon | [optional] 
**Devices** | [**[]SceneDeviceRequest**](SceneDeviceRequest.md) | Non-sequential list of device actions | 
**Sequences** | [**[][]SceneAction**](array.md) | List of parallel action sequences | [optional] 
**Mode** | [**SceneModeRequest**](SceneModeRequest.md) |  | [optional] 
**SecurityMode** | [**SceneSecurityModeRequest**](SceneSecurityModeRequest.md) |  | [optional] 
**Devicegroups** | [**[]SceneDeviceGroupRequest**](SceneDeviceGroupRequest.md) | List of device group actions | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


