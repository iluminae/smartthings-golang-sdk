# Go API client for smartthings

# Overview

This is the reference documentation for the SmartThings API.

The SmartThings API supports [REST](https://en.wikipedia.org/wiki/Representational_state_transfer), resources are protected with [OAuth 2.0 Bearer Tokens](https://tools.ietf.org/html/rfc6750#section-2.1), and all responses are sent as [JSON](http://www.json.org/).

# Authentication

All SmartThings resources are protected with [OAuth 2.0 Bearer Tokens](https://tools.ietf.org/html/rfc6750#section-2.1) sent on the request as an `Authorization: Bearer <TOKEN>` header, and operations require specific OAuth scopes that specify the exact permissions authorized by the user.

## Token types

There are two types of tokens: SmartApp tokens, and personal access tokens.

SmartApp tokens are used to communicate between third-party integrations, or SmartApps, and the SmartThings API.
When a SmartApp is called by the SmartThings platform, it is sent an authorization token that can be used to interact with the SmartThings API.

Personal access tokens are used to interact with the API for non-SmartApp use cases. They can be created and managed on the [personal access tokens page](https://account.smartthings.com/tokens).

## OAuth2 scopes

Operations may be protected by one or more OAuth security schemes, which specify the required permissions.
Each scope for a given scheme is required.
If multiple schemes are specified (not common), you may use either scheme.

SmartApp token scopes are derived from the permissions requested by the SmartApp and granted by the end-user during installation.
Personal access token scopes are associated with the specific permissions authorized when creating them.

Scopes are generally in the form `permission:entity-type:entity-id`.

**An `*` used for the `entity-id` specifies that the permission may be applied to all entities that the token type has access to, or may be replaced with a specific ID.**

For more information about authrization and permissions, please see the [Authorization and permissions guide](https://smartthings.developer.samsung.com/develop/guides/smartapps/auth-and-permissions.html).

<!-- ReDoc-Inject: <security-definitions> -->

# Errors

The SmartThings API uses conventional HTTP response codes to indicate the success or failure of a request.
In general, a `2XX` response code indicates success, a `4XX` response code indicates an error given the inputs for the request, and a `5XX` response code indicates a failure on the SmartThings platform.

API errors will contain a JSON response body with more information about the error:

```json
{
  \"requestId\": \"031fec1a-f19f-470a-a7da-710569082846\"
  \"error\": {
    \"code\": \"ConstraintViolationError\",
    \"message\": \"Validation errors occurred while process your request.\",
    \"details\": [
      { \"code\": \"PatternError\", \"target\": \"latitude\", \"message\": \"Invalid format.\" },
      { \"code\": \"SizeError\", \"target\": \"name\", \"message\": \"Too small.\" },
      { \"code\": \"SizeError\", \"target\": \"description\", \"message\": \"Too big.\" }
    ]
  }
}
```

## Error Response Body

The error response attributes are:

| Property | Type | Required | Description |
| --- | --- | --- | --- |
| requestId | String | No | A request identifier that can be used to correlate an error to additional logging on the SmartThings servers.
| error | Error | **Yes** | The Error object, documented below.

## Error Object

The Error object contains the following attributes:

| Property | Type | Required | Description |
| --- | --- | --- | --- |
| code | String | **Yes** | A SmartThings-defined error code that serves as a more specific indicator of the error than the HTTP error code specified in the response. See [SmartThings Error Codes](#section/Errors/SmartThings-Error-Codes) for more information.
| message | String | **Yes** | A description of the error, intended to aid developers in debugging of error responses.
| target | String | No | The target of the particular error. For example, it could be the name of the property that caused the error.
| details | Error[] | No | An array of Error objects that typically represent distinct, related errors that occurred during the request. As an optional attribute, this may be null or an empty array.

## Standard HTTP Error Codes

The following table lists the most common HTTP error response:

| Code | Name | Description |
| --- | --- | --- |
| 400 | Bad Request | The client has issued an invalid request. This is commonly used to specify validation errors in a request payload.
| 401 | Unauthorized | Authorization for the API is required, but the request has not been authenticated.
| 403 | Forbidden | The request has been authenticated but does not have appropriate permissions, or a requested resource is not found.
| 404 | Not Found | Specifies the requested path does not exist.
| 406 | Not Acceptable | The client has requested a MIME type via the Accept header for a value not supported by the server.
| 415 | Unsupported Media Type | The client has defined a contentType header that is not supported by the server.
| 422 | Unprocessable Entity | The client has made a valid request, but the server cannot process it. This is often used for APIs for which certain limits have been exceeded.
| 429 | Too Many Requests | The client has exceeded the number of requests allowed for a given time window.
| 500 | Internal Server Error | An unexpected error on the SmartThings servers has occurred. These errors should be rare.
| 501 | Not Implemented | The client request was valid and understood by the server, but the requested feature has yet to be implemented. These errors should be rare.

## SmartThings Error Codes

SmartThings specifies several standard custom error codes.
These provide more information than the standard HTTP error response codes.
The following table lists the standard SmartThings error codes and their description:

| Code | Typical HTTP Status Codes | Description |
| --- | --- | --- |
| PatternError | 400, 422 | The client has provided input that does not match the expected pattern.
| ConstraintViolationError | 422 | The client has provided input that has violated one or more constraints.
| NotNullError | 422 | The client has provided a null input for a field that is required to be non-null.
| NullError | 422 | The client has provided an input for a field that is required to be null.
| NotEmptyError | 422 | The client has provided an empty input for a field that is required to be non-empty.
| SizeError | 400, 422 | The client has provided a value that does not meet size restrictions.
| Unexpected Error | 500 | A non-recoverable error condition has occurred. Indicates a problem occurred on the SmartThings server that is no fault of the client.
| UnprocessableEntityError | 422 | The client has sent a malformed request body.
| TooManyRequestError | 429 | The client issued too many requests too quickly.
| LimitError | 422 | The client has exceeded certain limits an API enforces.
| UnsupportedOperationError | 400, 422 | The client has issued a request to a feature that currently isn't supported by the SmartThings platform. These should be rare.

## Custom Error Codes

An API may define its own error codes where appropriate.
These custom error codes are documented as part of that specific API's documentation.

# Warnings
The SmartThings API issues warning messages via standard HTTP Warning headers. These messages do not represent a request failure, but provide additional information that the requester might want to act upon.
For instance a warning will be issued if you are using an old API version.

# API Versions

The SmartThings API supports both path and header-based versioning.
The following are equivalent:

- https://api.smartthings.com/v1/locations
- https://api.smartthings.com/locations with header `Accept: application/vnd.smartthings+json;v=1`

Currently, only version 1 is available.

# Paging

Operations that return a list of objects return a paginated response.
The `_links` object contains the items returned, and links to the next and previous result page, if applicable.

```json
{
  \"items\": [
    {
      \"locationId\": \"6b3d1909-1e1c-43ec-adc2-5f941de4fbf9\",
      \"name\": \"Home\"
    },
    {
      \"locationId\": \"6b3d1909-1e1c-43ec-adc2-5f94d6g4fbf9\",
      \"name\": \"Work\"
    }
    ....
  ],
  \"_links\": {
    \"next\": {
      \"href\": \"https://api.smartthings.com/v1/locations?page=3\"
    },
    \"previous\": {
      \"href\": \"https://api.smartthings.com/v1/locations?page=1\"
    }
  }
}
```

# Localization

Some SmartThings API's support localization.
Specific information regarding localization endpoints are documented in the API itself.
However, the following should apply to all endpoints that support localization.

## Fallback Patterns

When making a request with the `Accept-Language` header, this fallback pattern is observed.
* Response will be translated with exact locale tag.
* If a translation does not exist for the requested language and region, the translation for the language will be returned.
* If a translation does not exist for the language, English (en) will be returned.
* Finally, an untranslated response will be returned in the absense of the above translations.

## Accept-Language Header
The format of the `Accept-Language` header follows what is defined in [RFC 7231, section 5.3.5](https://tools.ietf.org/html/rfc7231#section-5.3.5)

## Content-Language
The `Content-Language` header should be set on the response from the server to indicate which translation was given back to the client.
The absense of the header indicates that the server did not recieve a request with the `Accept-Language` header set.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0-PREVIEW
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./smartthings"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.smartthings.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppsApi* | [**CreateApp**](docs/AppsApi.md#createapp) | **Post** /apps | Create an app.
*AppsApi* | [**DeleteApp**](docs/AppsApi.md#deleteapp) | **Delete** /apps/{appNameOrId} | Delete an app.
*AppsApi* | [**GenerateAppOauth**](docs/AppsApi.md#generateappoauth) | **Post** /apps/{appNameOrId}/oauth/generate | Generate an app&#39;s oauth client/secret.
*AppsApi* | [**GetApp**](docs/AppsApi.md#getapp) | **Get** /apps/{appNameOrId} | Get an app.
*AppsApi* | [**GetAppOauth**](docs/AppsApi.md#getappoauth) | **Get** /apps/{appNameOrId}/oauth | Get an app&#39;s oauth settings.
*AppsApi* | [**GetAppSettings**](docs/AppsApi.md#getappsettings) | **Get** /apps/{appNameOrId}/settings | Get settings.
*AppsApi* | [**ListApps**](docs/AppsApi.md#listapps) | **Get** /apps | List apps.
*AppsApi* | [**Register**](docs/AppsApi.md#register) | **Put** /apps/{appNameOrId}/register | Sends a confirmation request to App.
*AppsApi* | [**UpdateApp**](docs/AppsApi.md#updateapp) | **Put** /apps/{appNameOrId} | Update an app.
*AppsApi* | [**UpdateAppOauth**](docs/AppsApi.md#updateappoauth) | **Put** /apps/{appNameOrId}/oauth | Update an app&#39;s oauth settings.
*AppsApi* | [**UpdateAppSettings**](docs/AppsApi.md#updateappsettings) | **Put** /apps/{appNameOrId}/settings | Update settings.
*AppsApi* | [**UpdateSignatureType**](docs/AppsApi.md#updatesignaturetype) | **Put** /apps/{appNameOrId}/signature-type | Update an app&#39;s signature type.
*DeviceprofilesApi* | [**CreateDeviceProfile**](docs/DeviceprofilesApi.md#createdeviceprofile) | **Post** /deviceprofiles | Create a device profile
*DeviceprofilesApi* | [**DeleteDeviceProfile**](docs/DeviceprofilesApi.md#deletedeviceprofile) | **Delete** /deviceprofiles/{deviceProfileId} | Delete a device profile
*DeviceprofilesApi* | [**GetDeviceProfile**](docs/DeviceprofilesApi.md#getdeviceprofile) | **Get** /deviceprofiles/{deviceProfileId} | GET a device profile
*DeviceprofilesApi* | [**ListDeviceProfiles**](docs/DeviceprofilesApi.md#listdeviceprofiles) | **Get** /deviceprofiles | List all device profiles for the authenticated user
*DeviceprofilesApi* | [**UpdateDeviceProfile**](docs/DeviceprofilesApi.md#updatedeviceprofile) | **Put** /deviceprofiles/{deviceProfileId} | Update a device profile.
*DevicesApi* | [**CreateDeviceEvents**](docs/DevicesApi.md#createdeviceevents) | **Post** /devices/{deviceId}/events | Create Device Events.
*DevicesApi* | [**DeleteDevice**](docs/DevicesApi.md#deletedevice) | **Delete** /devices/{deviceId} | Delete a Device.
*DevicesApi* | [**ExecuteDeviceCommands**](docs/DevicesApi.md#executedevicecommands) | **Post** /devices/{deviceId}/commands | Execute commands on device.
*DevicesApi* | [**GetDevice**](docs/DevicesApi.md#getdevice) | **Get** /devices/{deviceId} | Get a device&#39;s description.
*DevicesApi* | [**GetDeviceComponentStatus**](docs/DevicesApi.md#getdevicecomponentstatus) | **Get** /devices/{deviceId}/components/{componentId}/status | Get a device component&#39;s status.
*DevicesApi* | [**GetDeviceStatus**](docs/DevicesApi.md#getdevicestatus) | **Get** /devices/{deviceId}/status | Get the full status of a device.
*DevicesApi* | [**GetDeviceStatusByCapability**](docs/DevicesApi.md#getdevicestatusbycapability) | **Get** /devices/{deviceId}/components/{componentId}/capabilities/{capabilityId}/status | Get a capability&#39;s status.
*DevicesApi* | [**GetDevices**](docs/DevicesApi.md#getdevices) | **Get** /devices | List devices.
*DevicesApi* | [**InstallDevice**](docs/DevicesApi.md#installdevice) | **Post** /devices | Install a Device.
*DevicesApi* | [**UpdateDevice**](docs/DevicesApi.md#updatedevice) | **Put** /devices/{deviceId} | Update a device.
*InstalledappsApi* | [**CreateInstalledAppEvents**](docs/InstalledappsApi.md#createinstalledappevents) | **Post** /installedapps/{installedAppId}/events | Create Installed App events.
*InstalledappsApi* | [**DeleteInstallation**](docs/InstalledappsApi.md#deleteinstallation) | **Delete** /installedapps/{installedAppId} | Delete an installed app.
*InstalledappsApi* | [**GetInstallation**](docs/InstalledappsApi.md#getinstallation) | **Get** /installedapps/{installedAppId} | Get an installed app.
*InstalledappsApi* | [**GetInstallationConfig**](docs/InstalledappsApi.md#getinstallationconfig) | **Get** /installedapps/{installedAppId}/configs/{configurationId} | Get an installed app configuration.
*InstalledappsApi* | [**ListInstallationConfig**](docs/InstalledappsApi.md#listinstallationconfig) | **Get** /installedapps/{installedAppId}/configs | List an installed app&#39;s configurations.
*InstalledappsApi* | [**ListInstallations**](docs/InstalledappsApi.md#listinstallations) | **Get** /installedapps | List installed apps.
*LocationsApi* | [**CreateLocation**](docs/LocationsApi.md#createlocation) | **Post** /locations | Create a Location.
*LocationsApi* | [**DeleteLocation**](docs/LocationsApi.md#deletelocation) | **Delete** /locations/{locationId} | Delete a Location.
*LocationsApi* | [**GetLocation**](docs/LocationsApi.md#getlocation) | **Get** /locations/{locationId} | Get a Location.
*LocationsApi* | [**ListLocations**](docs/LocationsApi.md#listlocations) | **Get** /locations | List Locations.
*LocationsApi* | [**UpdateLocation**](docs/LocationsApi.md#updatelocation) | **Put** /locations/{locationId} | Update a Location.
*RoomsApi* | [**CreateRoom**](docs/RoomsApi.md#createroom) | **Post** /locations/{locationId}/rooms | Create a Room.
*RoomsApi* | [**DeleteRoom**](docs/RoomsApi.md#deleteroom) | **Delete** /locations/{locationId}/rooms/{roomId} | Delete a Room.
*RoomsApi* | [**GetRoom**](docs/RoomsApi.md#getroom) | **Get** /locations/{locationId}/rooms/{roomId} | Get a Room.
*RoomsApi* | [**ListRooms**](docs/RoomsApi.md#listrooms) | **Get** /locations/{locationId}/rooms | List Rooms.
*RoomsApi* | [**UpdateRoom**](docs/RoomsApi.md#updateroom) | **Put** /locations/{locationId}/rooms/{roomId} | Update a Room.
*RulesApi* | [**CreateRule**](docs/RulesApi.md#createrule) | **Post** /rules | Create a rule
*RulesApi* | [**DeleteRule**](docs/RulesApi.md#deleterule) | **Delete** /rules/{ruleId} | Delete a rule
*RulesApi* | [**ExecuteRule**](docs/RulesApi.md#executerule) | **Post** /rules/execute/{ruleId} | Execute a rule
*RulesApi* | [**GetRule**](docs/RulesApi.md#getrule) | **Get** /rules/{ruleId} | Get a Rule
*RulesApi* | [**ListRules**](docs/RulesApi.md#listrules) | **Get** /rules | Rules list
*RulesApi* | [**UpdateRule**](docs/RulesApi.md#updaterule) | **Put** /rules/{ruleId} | Update a rule
*ScenesApi* | [**ExecuteScene**](docs/ScenesApi.md#executescene) | **Post** /scenes/{sceneId}/execute | Execute Scene
*ScenesApi* | [**ListScenes**](docs/ScenesApi.md#listscenes) | **Get** /scenes | List Scenes
*SchedulesApi* | [**CreateSchedule**](docs/SchedulesApi.md#createschedule) | **Post** /installedapps/{installedAppId}/schedules | Save an installed app schedule.
*SchedulesApi* | [**DeleteSchedule**](docs/SchedulesApi.md#deleteschedule) | **Delete** /installedapps/{installedAppId}/schedules/{scheduleName} | Delete a schedule.
*SchedulesApi* | [**DeleteSchedules**](docs/SchedulesApi.md#deleteschedules) | **Delete** /installedapps/{installedAppId}/schedules | Delete all of an installed app&#39;s schedules.
*SchedulesApi* | [**GetSchedule**](docs/SchedulesApi.md#getschedule) | **Get** /installedapps/{installedAppId}/schedules/{scheduleName} | Get an installed app&#39;s schedule.
*SchedulesApi* | [**GetSchedules**](docs/SchedulesApi.md#getschedules) | **Get** /installedapps/{installedAppId}/schedules | List installed app schedules.
*SubscriptionsApi* | [**DeleteAllSubscriptions**](docs/SubscriptionsApi.md#deleteallsubscriptions) | **Delete** /installedapps/{installedAppId}/subscriptions | Delete all of an installed app&#39;s subscriptions.
*SubscriptionsApi* | [**DeleteSubscription**](docs/SubscriptionsApi.md#deletesubscription) | **Delete** /installedapps/{installedAppId}/subscriptions/{subscriptionId} | Delete an installed app&#39;s subscription.
*SubscriptionsApi* | [**GetSubscription**](docs/SubscriptionsApi.md#getsubscription) | **Get** /installedapps/{installedAppId}/subscriptions/{subscriptionId} | Get an installed app&#39;s subscription.
*SubscriptionsApi* | [**ListSubscriptions**](docs/SubscriptionsApi.md#listsubscriptions) | **Get** /installedapps/{installedAppId}/subscriptions | List an installed app&#39;s subscriptions.
*SubscriptionsApi* | [**SaveSubscription**](docs/SubscriptionsApi.md#savesubscription) | **Post** /installedapps/{installedAppId}/subscriptions | Create a subscription for an installed app.


## Documentation For Models

 - [Action](docs/Action.md)
 - [AdhocMessage](docs/AdhocMessage.md)
 - [AdhocMessageTemplate](docs/AdhocMessageTemplate.md)
 - [App](docs/App.md)
 - [AppClassification](docs/AppClassification.md)
 - [AppDeviceDetails](docs/AppDeviceDetails.md)
 - [AppOAuth](docs/AppOAuth.md)
 - [AppTargetStatus](docs/AppTargetStatus.md)
 - [AppType](docs/AppType.md)
 - [AppUiSettings](docs/AppUiSettings.md)
 - [Argument](docs/Argument.md)
 - [ArrayOperand](docs/ArrayOperand.md)
 - [AttributeProperties](docs/AttributeProperties.md)
 - [AttributePropertiesData](docs/AttributePropertiesData.md)
 - [AttributePropertiesUnit](docs/AttributePropertiesUnit.md)
 - [AttributeSchema](docs/AttributeSchema.md)
 - [AttributeState](docs/AttributeState.md)
 - [BetweenCondition](docs/BetweenCondition.md)
 - [CapabilityAttribute](docs/CapabilityAttribute.md)
 - [CapabilityAttributeEnumCommands](docs/CapabilityAttributeEnumCommands.md)
 - [CapabilityCommand](docs/CapabilityCommand.md)
 - [CapabilityReference](docs/CapabilityReference.md)
 - [CapabilitySubscriptionDetail](docs/CapabilitySubscriptionDetail.md)
 - [CapabilitySummary](docs/CapabilitySummary.md)
 - [ChangesCondition](docs/ChangesCondition.md)
 - [CommandAction](docs/CommandAction.md)
 - [ComponentTranslations](docs/ComponentTranslations.md)
 - [Condition](docs/Condition.md)
 - [ConditionAggregationMode](docs/ConditionAggregationMode.md)
 - [ConfigEntry](docs/ConfigEntry.md)
 - [CreateAppRequest](docs/CreateAppRequest.md)
 - [CreateAppResponse](docs/CreateAppResponse.md)
 - [CreateDeviceProfileRequest](docs/CreateDeviceProfileRequest.md)
 - [CreateInstalledAppEventsRequest](docs/CreateInstalledAppEventsRequest.md)
 - [CreateLocationRequest](docs/CreateLocationRequest.md)
 - [CreateOrUpdateLambdaSmartAppRequest](docs/CreateOrUpdateLambdaSmartAppRequest.md)
 - [CreateOrUpdateWebhookSmartAppRequest](docs/CreateOrUpdateWebhookSmartAppRequest.md)
 - [CreateRoomRequest](docs/CreateRoomRequest.md)
 - [CronSchedule](docs/CronSchedule.md)
 - [DashboardCardLifecycle](docs/DashboardCardLifecycle.md)
 - [DateOperand](docs/DateOperand.md)
 - [DateTimeOperand](docs/DateTimeOperand.md)
 - [DayOfWeek](docs/DayOfWeek.md)
 - [DeleteInstalledAppResponse](docs/DeleteInstalledAppResponse.md)
 - [Device](docs/Device.md)
 - [DeviceActivity](docs/DeviceActivity.md)
 - [DeviceCategory](docs/DeviceCategory.md)
 - [DeviceCommand](docs/DeviceCommand.md)
 - [DeviceCommandsEvent](docs/DeviceCommandsEvent.md)
 - [DeviceCommandsEventCommand](docs/DeviceCommandsEventCommand.md)
 - [DeviceCommandsRequest](docs/DeviceCommandsRequest.md)
 - [DeviceComponent](docs/DeviceComponent.md)
 - [DeviceComponentReference](docs/DeviceComponentReference.md)
 - [DeviceConfig](docs/DeviceConfig.md)
 - [DeviceEvent](docs/DeviceEvent.md)
 - [DeviceEventsRequest](docs/DeviceEventsRequest.md)
 - [DeviceHealthDetail](docs/DeviceHealthDetail.md)
 - [DeviceHealthEvent](docs/DeviceHealthEvent.md)
 - [DeviceInstallRequest](docs/DeviceInstallRequest.md)
 - [DeviceInstallRequestApp](docs/DeviceInstallRequestApp.md)
 - [DeviceIntegrationType](docs/DeviceIntegrationType.md)
 - [DeviceLifecycle](docs/DeviceLifecycle.md)
 - [DeviceLifecycleDetail](docs/DeviceLifecycleDetail.md)
 - [DeviceLifecycleEvent](docs/DeviceLifecycleEvent.md)
 - [DeviceLifecycleMove](docs/DeviceLifecycleMove.md)
 - [DeviceNetworkSecurityLevel](docs/DeviceNetworkSecurityLevel.md)
 - [DeviceOperand](docs/DeviceOperand.md)
 - [DeviceProfile](docs/DeviceProfile.md)
 - [DeviceProfileReference](docs/DeviceProfileReference.md)
 - [DeviceProfileStatus](docs/DeviceProfileStatus.md)
 - [DeviceResults](docs/DeviceResults.md)
 - [DeviceStateEvent](docs/DeviceStateEvent.md)
 - [DeviceStatus](docs/DeviceStatus.md)
 - [DeviceSubscriptionDetail](docs/DeviceSubscriptionDetail.md)
 - [DthDeviceDetails](docs/DthDeviceDetails.md)
 - [EndpointApp](docs/EndpointApp.md)
 - [EqualsCondition](docs/EqualsCondition.md)
 - [Error](docs/Error.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [EventType](docs/EventType.md)
 - [EveryAction](docs/EveryAction.md)
 - [ExecutionResult](docs/ExecutionResult.md)
 - [GenerateAppOAuthRequest](docs/GenerateAppOAuthRequest.md)
 - [GenerateAppOAuthResponse](docs/GenerateAppOAuthResponse.md)
 - [GetAppSettingsResponse](docs/GetAppSettingsResponse.md)
 - [GreaterThanCondition](docs/GreaterThanCondition.md)
 - [GreaterThanOrEqualsCondition](docs/GreaterThanOrEqualsCondition.md)
 - [HubHealthDetail](docs/HubHealthDetail.md)
 - [HubHealthEvent](docs/HubHealthEvent.md)
 - [IconImage](docs/IconImage.md)
 - [IfAction](docs/IfAction.md)
 - [IfActionAllOf](docs/IfActionAllOf.md)
 - [InstallConfiguration](docs/InstallConfiguration.md)
 - [InstallConfigurationDetail](docs/InstallConfigurationDetail.md)
 - [InstallConfigurationStatus](docs/InstallConfigurationStatus.md)
 - [InstalledApp](docs/InstalledApp.md)
 - [InstalledAppIconImage](docs/InstalledAppIconImage.md)
 - [InstalledAppLifecycle](docs/InstalledAppLifecycle.md)
 - [InstalledAppLifecycleError](docs/InstalledAppLifecycleError.md)
 - [InstalledAppLifecycleEvent](docs/InstalledAppLifecycleEvent.md)
 - [InstalledAppStatus](docs/InstalledAppStatus.md)
 - [InstalledAppType](docs/InstalledAppType.md)
 - [InstalledAppUi](docs/InstalledAppUi.md)
 - [Interval](docs/Interval.md)
 - [IntervalUnit](docs/IntervalUnit.md)
 - [IrDeviceDetails](docs/IrDeviceDetails.md)
 - [IrDeviceDetailsFunctionCodes](docs/IrDeviceDetailsFunctionCodes.md)
 - [IsaResults](docs/IsaResults.md)
 - [LambdaSmartApp](docs/LambdaSmartApp.md)
 - [LessThanCondition](docs/LessThanCondition.md)
 - [LessThanOrEqualsCondition](docs/LessThanOrEqualsCondition.md)
 - [Link](docs/Link.md)
 - [Links](docs/Links.md)
 - [LocaleReference](docs/LocaleReference.md)
 - [LocaleVariables](docs/LocaleVariables.md)
 - [Location](docs/Location.md)
 - [LocationAction](docs/LocationAction.md)
 - [LocationAttribute](docs/LocationAttribute.md)
 - [LocationOperand](docs/LocationOperand.md)
 - [Message](docs/Message.md)
 - [MessageConfig](docs/MessageConfig.md)
 - [MessageTemplate](docs/MessageTemplate.md)
 - [MessageType](docs/MessageType.md)
 - [Mode](docs/Mode.md)
 - [ModeConfig](docs/ModeConfig.md)
 - [ModeEvent](docs/ModeEvent.md)
 - [ModeSubscriptionDetail](docs/ModeSubscriptionDetail.md)
 - [Notice](docs/Notice.md)
 - [NoticeAction](docs/NoticeAction.md)
 - [NoticeCode](docs/NoticeCode.md)
 - [OnceSchedule](docs/OnceSchedule.md)
 - [Operand](docs/Operand.md)
 - [OperandAggregationMode](docs/OperandAggregationMode.md)
 - [Owner](docs/Owner.md)
 - [PagedApp](docs/PagedApp.md)
 - [PagedApps](docs/PagedApps.md)
 - [PagedDeviceProfiles](docs/PagedDeviceProfiles.md)
 - [PagedDevices](docs/PagedDevices.md)
 - [PagedInstallConfigurations](docs/PagedInstallConfigurations.md)
 - [PagedInstalledApps](docs/PagedInstalledApps.md)
 - [PagedLocation](docs/PagedLocation.md)
 - [PagedLocations](docs/PagedLocations.md)
 - [PagedMessageTemplate](docs/PagedMessageTemplate.md)
 - [PagedRooms](docs/PagedRooms.md)
 - [PagedRules](docs/PagedRules.md)
 - [PagedSchedules](docs/PagedSchedules.md)
 - [PagedSubscriptions](docs/PagedSubscriptions.md)
 - [ParentType](docs/ParentType.md)
 - [PermissionConfig](docs/PermissionConfig.md)
 - [PredefinedMessage](docs/PredefinedMessage.md)
 - [PrincipalType](docs/PrincipalType.md)
 - [Room](docs/Room.md)
 - [Rule](docs/Rule.md)
 - [RuleAllOf](docs/RuleAllOf.md)
 - [RuleExecutionResponse](docs/RuleExecutionResponse.md)
 - [RuleMetadata](docs/RuleMetadata.md)
 - [RuleRequest](docs/RuleRequest.md)
 - [SceneAction](docs/SceneAction.md)
 - [SceneArgument](docs/SceneArgument.md)
 - [SceneCapability](docs/SceneCapability.md)
 - [SceneCommand](docs/SceneCommand.md)
 - [SceneComponent](docs/SceneComponent.md)
 - [SceneConfig](docs/SceneConfig.md)
 - [SceneDevice](docs/SceneDevice.md)
 - [SceneDeviceGroup](docs/SceneDeviceGroup.md)
 - [SceneDeviceGroupRequest](docs/SceneDeviceGroupRequest.md)
 - [SceneDeviceRequest](docs/SceneDeviceRequest.md)
 - [SceneLifecycle](docs/SceneLifecycle.md)
 - [SceneLifecycleDetail](docs/SceneLifecycleDetail.md)
 - [SceneLifecycleEvent](docs/SceneLifecycleEvent.md)
 - [SceneMode](docs/SceneMode.md)
 - [SceneModeRequest](docs/SceneModeRequest.md)
 - [ScenePagedResult](docs/ScenePagedResult.md)
 - [SceneRequest](docs/SceneRequest.md)
 - [SceneSecurityModeRequest](docs/SceneSecurityModeRequest.md)
 - [SceneSleepRequest](docs/SceneSleepRequest.md)
 - [SceneSummary](docs/SceneSummary.md)
 - [Schedule](docs/Schedule.md)
 - [ScheduleRequest](docs/ScheduleRequest.md)
 - [SecurityArmStateDetail](docs/SecurityArmStateDetail.md)
 - [SecurityArmStateEvent](docs/SecurityArmStateEvent.md)
 - [SignatureType](docs/SignatureType.md)
 - [SimpleCondition](docs/SimpleCondition.md)
 - [SimpleValue](docs/SimpleValue.md)
 - [SingleOperandCondition](docs/SingleOperandCondition.md)
 - [SleepAction](docs/SleepAction.md)
 - [SmartAppDashboardCardEventRequest](docs/SmartAppDashboardCardEventRequest.md)
 - [SmartAppEventRequest](docs/SmartAppEventRequest.md)
 - [StandardSuccessResponse](docs/StandardSuccessResponse.md)
 - [StringConfig](docs/StringConfig.md)
 - [Subscription](docs/Subscription.md)
 - [SubscriptionDelete](docs/SubscriptionDelete.md)
 - [SubscriptionFilter](docs/SubscriptionFilter.md)
 - [SubscriptionFilterTypes](docs/SubscriptionFilterTypes.md)
 - [SubscriptionMode](docs/SubscriptionMode.md)
 - [SubscriptionRequest](docs/SubscriptionRequest.md)
 - [SubscriptionSource](docs/SubscriptionSource.md)
 - [SubscriptionTarget](docs/SubscriptionTarget.md)
 - [TimeOperand](docs/TimeOperand.md)
 - [TimeReference](docs/TimeReference.md)
 - [TimerEvent](docs/TimerEvent.md)
 - [TimerType](docs/TimerType.md)
 - [UpdateAppOAuthRequest](docs/UpdateAppOAuthRequest.md)
 - [UpdateAppRequest](docs/UpdateAppRequest.md)
 - [UpdateAppSettingsRequest](docs/UpdateAppSettingsRequest.md)
 - [UpdateAppSettingsResponse](docs/UpdateAppSettingsResponse.md)
 - [UpdateDeviceProfileRequest](docs/UpdateDeviceProfileRequest.md)
 - [UpdateDeviceRequest](docs/UpdateDeviceRequest.md)
 - [UpdateDeviceRequestComponents](docs/UpdateDeviceRequestComponents.md)
 - [UpdateLocationRequest](docs/UpdateLocationRequest.md)
 - [UpdateRoomRequest](docs/UpdateRoomRequest.md)
 - [UpdateSignatureTypeRequest](docs/UpdateSignatureTypeRequest.md)
 - [ViperDeviceDetails](docs/ViperDeviceDetails.md)
 - [WebhookSmartApp](docs/WebhookSmartApp.md)


## Documentation For Authorization



## Basic

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Bearer


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://auth-global.api.smartthings.com
- **Scopes**: 
 - **r:installedapps:***: Read details about installed SmartApps, such as which devices have been configured for the installation. For SmartApp tokens, the scope is restricted to the location the SmartApp is installed into. For personal access tokens, the scope is limited to the account associated with the token. 
 - **l:installedapps**: View a list of installed SmartApps. For SmartApp tokens, the scope is restricted to the location the SmartApp is installed into. For personal access tokens, the scope is limited to the account associated with the token. 
 - **w:installedapps:***: Create, update, or delete installed SmartApps. For SmartApp tokens, the scope is restricted to the location the SmartApp is installed into. For personal access tokens, the scope is limited to the account associated with the token. 
 - **r:apps:***: Read details about a SmartApp. Only applicable for personal access tokens, and the scope is limited to the SmartApps associated with the token's account. 
 - **w:apps:***: Create, update, or delete SmartApps. For SmartApp tokens, the scope is restricted to the location the SmartApp is installed into. For personal access tokens, the scope is limited to the account associated with the token. 
 - **l:devices**: View a list of devices. For SmartApp tokens, the scope is restricted to the location the SmartApp is installed into. For personal access tokens, the scope is limited to the account associated with the token. 
 - **r:devices:***: Read details about a device, including device attribute state. For SmartApp tokens, the scope is restricted to the location the SmartApp is installed into. For personal access tokens, the scope is limited to the account associated with the token. This scope is required to create subscriptions. 
 - **w:devices:***: Update details such as the device name, or delete a device. For SmartApp tokens, the scope is restricted to the location the SmartApp is installed into. For personal access tokens, the scope is limited to the account associated with the token. 
 - **x:devices:***: Execute commands on a device. For SmartApp tokens, the scope is restricted to the location the SmartApp is installed into. For personal access tokens, the scope is limited to the account associated with the token. 
 - **r:deviceprofiles**: View details of device profiles associated with the account. Only applicable for personal access tokens. 
 - **w:deviceprofiles**: Create, update, or delete device profiles. Only applicable to personal access tokens, and the device profile must be owned by the same account associated with the token. 
 - **i:deviceprofiles**: Create devices of the type associated with the device profile. Only applicable for SmartApp tokens, and is requires the device profile and the SmartApp have the same account owner.
 - **r:scenes:***: Read details about a scene. For personal access tokens, the scope is limited to the account associated with the token. 
 - **x:scenes:***: Execute a scene. For personal access tokens, the scope is limited to the account associated with the token. 
 - **r:schedules**: Read details of scheduled executions. For SmartApp tokens, the scope is restricted to the installed SmartApp. For personal access tokens, the scope is limited to the account associated with the token. 
 - **w:schedules**: Create, update, or delete schedules. For SmartApp tokens, the scope is restricted to the installed SmartApp. For personal access tokens, the scope is limited to the account associated with the token. 
 - **l:locations**: View a list of locations. Only applicable for personal access tokens, and the scope is limited to the account associated with the token. 
 - **r:locations:***: Read details of a location, such as geocoordinates and temperature scale. For SmartApp tokens, the scope is restricted to the installed SmartApp. For personal access tokens, the scope is limited to the account associated with the token. 
 - **w:locations:***: Create, update, and delete locations. Only applicable for personal access tokens (the scope is limited to the account associated with the token). 
 - **r:hubs**: Read hubs.
 - **r:security:locations:*:armstate**: Read arm state in the given location.

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```



## Author



