# PagedApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | A user defined unique identifier for an app.  It is alpha-numeric, may contain dashes, underscores, periods, and be less then 250 characters long.  It must be unique within your account.  | [optional] 
**AppId** | **string** | A globally unique identifier for an app. | [optional] 
**AppType** | [**AppType**](AppType.md) |  | [optional] 
**Classifications** | [**[]AppClassification**](AppClassification.md) | An App maybe associated to many classifications.  A classification drives how the integration is presented to the user in the SmartThings mobile clients.  These classifications include: * AUTOMATION - Denotes an integration that should display under the \&quot;Automation\&quot; tab in mobile clients. * SERVICE - Denotes an integration that is classified as a \&quot;Service\&quot;. * DEVICE - Denotes an integration that should display under the \&quot;Device\&quot; tab in mobile clients. * CONNECTED_SERVICE - Denotes an integration that should display under the \&quot;Connected Services\&quot; menu in mobile clients. * HIDDEN - Denotes an integration that should not display in mobile clients  | [optional] 
**DisplayName** | **string** | A default display name for an app.  | [optional] 
**Description** | **string** | A default description for an app.  | [optional] 
**IconImage** | [**IconImage**](IconImage.md) |  | [optional] 
**Owner** | [**Owner**](Owner.md) |  | [optional] 
**CreatedDate** | [**time.Time**](time.Time.md) | A UTC ISO-8601 Date-Time String | [optional] 
**LastUpdatedDate** | [**time.Time**](time.Time.md) | A UTC ISO-8601 Date-Time String | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


