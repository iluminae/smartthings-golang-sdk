# InstalledAppLifecycleEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | The id of the event. | [optional] 
**LocationId** | **string** | The ID of the location in which the event was triggered. | [optional] 
**InstalledAppId** | **string** | The ID of the installed application. | [optional] 
**AppId** | **string** | The ID of the application. | [optional] 
**Lifecycle** | [**InstalledAppLifecycle**](InstalledAppLifecycle.md) |  | [optional] 
**Create** | [**map[string]interface{}**](.md) | Create installed app lifecycle.  | [optional] 
**Install** | [**map[string]interface{}**](.md) | Install installed app lifecycle.  | [optional] 
**Update** | [**map[string]interface{}**](.md) | Update installed app lifecycle.  | [optional] 
**Delete** | [**map[string]interface{}**](.md) | Delete installed app lifecycle.  | [optional] 
**Other** | [**map[string]interface{}**](.md) | Other installed app lifecycle.  | [optional] 
**Error** | [**InstalledAppLifecycleError**](InstalledAppLifecycleError.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


