# DeviceLifecycleEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lifecycle** | [**DeviceLifecycle**](DeviceLifecycle.md) |  | [optional] 
**EventId** | **string** | The id of the event. | [optional] 
**LocationId** | **string** | The id of the location in which the event was triggered. | [optional] 
**DeviceId** | **string** | The id of the device. | [optional] 
**DeviceName** | **string** | The name of the device | [optional] 
**Principal** | **string** | The principal that made the change | [optional] 
**Create** | [**map[string]interface{}**](.md) | Create device lifecycle.  | [optional] 
**Delete** | [**map[string]interface{}**](.md) | Delete device lifecycle.  | [optional] 
**Update** | [**map[string]interface{}**](.md) | Update device lifecycle.  | [optional] 
**MoveFrom** | [**DeviceLifecycleMove**](DeviceLifecycleMove.md) |  | [optional] 
**MoveTo** | [**DeviceLifecycleMove**](DeviceLifecycleMove.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


