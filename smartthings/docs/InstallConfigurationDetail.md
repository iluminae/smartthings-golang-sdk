# InstallConfigurationDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstalledAppId** | **string** | The ID of the installed app. | [optional] 
**ConfigurationId** | **string** | The ID to this configration instance. | [optional] 
**ConfigurationStatus** | [**InstallConfigurationStatus**](InstallConfigurationStatus.md) |  | [optional] 
**Config** | [**map[string][]ConfigEntry**](array.md) | A multi-map of configurations for an Installed App.  The map &#39;key&#39; is the configuration name and the &#39;value&#39; is an array of ConfigEntry models each containing a value and associated permissions.  The config key is alpha-numeric, may contain dashes, underscores, periods, and must be less then 50 characters long.  | [optional] 
**CreatedDate** | [**time.Time**](time.Time.md) | A UTC ISO-8601 Date-Time String | [optional] 
**LastUpdatedDate** | [**time.Time**](time.Time.md) | A UTC ISO-8601 Date-Time String | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


