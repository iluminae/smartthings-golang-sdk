# DeviceActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** | Device ID | [optional] 
**DeviceName** | **string** | Device nick name | [optional] 
**LocationId** | **string** | Location ID | [optional] 
**LocationName** | **string** | Location name | [optional] 
**Time** | [**time.Time**](time.Time.md) | The IS0-8601 date time strings in UTC of the activity | [optional] 
**Text** | **string** | Translated human readable string (localized) | [optional] 
**Component** | **string** | device component ID. Not nullable. | 
**ComponentLabel** | **string** | device component label. Nullable. | [optional] 
**Capability** | **string** | capability name | [optional] 
**Attribute** | **string** | attribute name | [optional] 
**Value** | [**map[string]interface{}**](.md) | attribute value | [optional] 
**Unit** | **string** |  | [optional] 
**Data** | [**map[string]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


