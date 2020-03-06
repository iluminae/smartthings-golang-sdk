# InstalledApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstalledAppId** | **string** | The ID of the installed app. | 
**InstalledAppType** | [**InstalledAppType**](InstalledAppType.md) |  | 
**InstalledAppStatus** | [**InstalledAppStatus**](InstalledAppStatus.md) |  | 
**DisplayName** | **string** | A user defined name for the installed app. May be null. | [optional] 
**AppId** | **string** | The ID of the app. | 
**ReferenceId** | **string** | A reference to an upstream system.  For example, Behaviors would reference the behaviorId. May be null.  | [optional] 
**LocationId** | **string** | The ID of the location to which the installed app may belong. | [optional] 
**Owner** | [**Owner**](Owner.md) |  | 
**Notices** | [**[]Notice**](Notice.md) |  | 
**CreatedDate** | [**time.Time**](time.Time.md) | A UTC ISO-8601 Date-Time String | 
**LastUpdatedDate** | [**time.Time**](time.Time.md) | A UTC ISO-8601 Date-Time String | 
**Ui** | [**InstalledAppUi**](InstalledApp_ui.md) |  | [optional] 
**IconImage** | [**InstalledAppIconImage**](InstalledApp_iconImage.md) |  | [optional] 
**Classifications** | **[]string** | An App maybe associated to many classifications.  A classification drives how the integration is presented to the user in the SmartThings mobile clients.  These classifications include: * AUTOMATION - Denotes an integration that should display under the \&quot;Automation\&quot; tab in mobile clients. * SERVICE - Denotes an integration that is classified as a \&quot;Service\&quot;. * DEVICE - Denotes an integration that should display under the \&quot;Device\&quot; tab in mobile clients. * CONNECTED_SERVICE - Denotes an integration that should display under the \&quot;Connected Services\&quot; menu in mobile clients. * HIDDEN - Denotes an integration that should not display in mobile clients  | 
**PrincipalType** | **string** | Denotes the principal type to be used with the app.  Default is LOCATION. | 
**SingleInstance** | **bool** | Inform the installation systems that the associated app can only be installed once within a user&#39;s account.  | [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


