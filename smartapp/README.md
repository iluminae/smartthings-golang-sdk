# Go API client for smartapp

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1.0.3
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
import "./smartapp"
```

## Documentation for API Endpoints

All URIs are relative to *https://smartapps*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**Execute**](docs/DefaultApi.md#execute) | **Post** / | Execute a third-party SmartApp.


## Documentation For Models

 - [Action](docs/Action.md)
 - [ActionType](docs/ActionType.md)
 - [AppLifecycle](docs/AppLifecycle.md)
 - [BasicBackgroundImage](docs/BasicBackgroundImage.md)
 - [BasicBody](docs/BasicBody.md)
 - [BasicButton](docs/BasicButton.md)
 - [BasicButtonPosition](docs/BasicButtonPosition.md)
 - [BasicButtonType](docs/BasicButtonType.md)
 - [BasicCard](docs/BasicCard.md)
 - [BasicIconButton](docs/BasicIconButton.md)
 - [BasicImage](docs/BasicImage.md)
 - [BasicImagePosition](docs/BasicImagePosition.md)
 - [BasicText](docs/BasicText.md)
 - [BasicTextButton](docs/BasicTextButton.md)
 - [BasicV2Body](docs/BasicV2Body.md)
 - [BasicV2Button](docs/BasicV2Button.md)
 - [BasicV2Card](docs/BasicV2Card.md)
 - [BasicV2IconButton](docs/BasicV2IconButton.md)
 - [BasicV2Image](docs/BasicV2Image.md)
 - [BasicV2ImageItem](docs/BasicV2ImageItem.md)
 - [BasicV2Item](docs/BasicV2Item.md)
 - [BasicV2Text](docs/BasicV2Text.md)
 - [BasicV2TextButton](docs/BasicV2TextButton.md)
 - [BooleanSetting](docs/BooleanSetting.md)
 - [BooleanSettingAllOf](docs/BooleanSettingAllOf.md)
 - [ClientDetails](docs/ClientDetails.md)
 - [ConfigEntry](docs/ConfigEntry.md)
 - [ConfigurationData](docs/ConfigurationData.md)
 - [ConfigurationPhase](docs/ConfigurationPhase.md)
 - [ConfigurationResponseData](docs/ConfigurationResponseData.md)
 - [ConfirmationData](docs/ConfirmationData.md)
 - [DashboardCard](docs/DashboardCard.md)
 - [DashboardCardTemplate](docs/DashboardCardTemplate.md)
 - [DashboardData](docs/DashboardData.md)
 - [DashboardResponseData](docs/DashboardResponseData.md)
 - [DecimalSetting](docs/DecimalSetting.md)
 - [DecimalSettingAllOf](docs/DecimalSettingAllOf.md)
 - [DeviceCommandsEvent](docs/DeviceCommandsEvent.md)
 - [DeviceCommandsEventCommand](docs/DeviceCommandsEventCommand.md)
 - [DeviceConfig](docs/DeviceConfig.md)
 - [DeviceEvent](docs/DeviceEvent.md)
 - [DeviceHealthEvent](docs/DeviceHealthEvent.md)
 - [DeviceLifecycle](docs/DeviceLifecycle.md)
 - [DeviceLifecycleEvent](docs/DeviceLifecycleEvent.md)
 - [DeviceLifecycleMove](docs/DeviceLifecycleMove.md)
 - [DeviceSetting](docs/DeviceSetting.md)
 - [DeviceSettingAllOf](docs/DeviceSettingAllOf.md)
 - [EmailSetting](docs/EmailSetting.md)
 - [EnumSetting](docs/EnumSetting.md)
 - [EnumSettingAllOf](docs/EnumSettingAllOf.md)
 - [EnumStyleType](docs/EnumStyleType.md)
 - [Event](docs/Event.md)
 - [EventData](docs/EventData.md)
 - [EventType](docs/EventType.md)
 - [ExecuteAction](docs/ExecuteAction.md)
 - [ExecuteData](docs/ExecuteData.md)
 - [ExecutionRequest](docs/ExecutionRequest.md)
 - [ExecutionResponse](docs/ExecutionResponse.md)
 - [FreeFormCard](docs/FreeFormCard.md)
 - [GroupedOption](docs/GroupedOption.md)
 - [ImageSetting](docs/ImageSetting.md)
 - [ImagesSetting](docs/ImagesSetting.md)
 - [ImagesSettingAllOf](docs/ImagesSettingAllOf.md)
 - [InitializeSetting](docs/InitializeSetting.md)
 - [InitializeSettingAllOf](docs/InitializeSettingAllOf.md)
 - [InstallData](docs/InstallData.md)
 - [InstalledApp](docs/InstalledApp.md)
 - [LaunchPluginAction](docs/LaunchPluginAction.md)
 - [LinkSetting](docs/LinkSetting.md)
 - [LinkSettingAllOf](docs/LinkSettingAllOf.md)
 - [LinkStyleType](docs/LinkStyleType.md)
 - [MessageConfig](docs/MessageConfig.md)
 - [ModeConfig](docs/ModeConfig.md)
 - [ModeEvent](docs/ModeEvent.md)
 - [ModeSetting](docs/ModeSetting.md)
 - [ModeSettingAllOf](docs/ModeSettingAllOf.md)
 - [NumberSetting](docs/NumberSetting.md)
 - [NumberSettingAllOf](docs/NumberSettingAllOf.md)
 - [OAuthCallbackData](docs/OAuthCallbackData.md)
 - [OAuthSetting](docs/OAuthSetting.md)
 - [OAuthSettingAllOf](docs/OAuthSettingAllOf.md)
 - [Option](docs/Option.md)
 - [Page](docs/Page.md)
 - [PageSetting](docs/PageSetting.md)
 - [PageSettingAllOf](docs/PageSettingAllOf.md)
 - [ParagraphSetting](docs/ParagraphSetting.md)
 - [PasswordSetting](docs/PasswordSetting.md)
 - [PasswordSettingAllOf](docs/PasswordSettingAllOf.md)
 - [PhoneSetting](docs/PhoneSetting.md)
 - [PhoneSettingAllOf](docs/PhoneSettingAllOf.md)
 - [PingData](docs/PingData.md)
 - [PingResponseData](docs/PingResponseData.md)
 - [SceneConfig](docs/SceneConfig.md)
 - [SceneLifecycle](docs/SceneLifecycle.md)
 - [SceneLifecycleEvent](docs/SceneLifecycleEvent.md)
 - [SceneSetting](docs/SceneSetting.md)
 - [SceneSettingAllOf](docs/SceneSettingAllOf.md)
 - [Section](docs/Section.md)
 - [SectionSetting](docs/SectionSetting.md)
 - [SectionSettingAllOf](docs/SectionSettingAllOf.md)
 - [SecurityArmStateEvent](docs/SecurityArmStateEvent.md)
 - [Setting](docs/Setting.md)
 - [SettingType](docs/SettingType.md)
 - [SimpleValue](docs/SimpleValue.md)
 - [StringConfig](docs/StringConfig.md)
 - [StyleType](docs/StyleType.md)
 - [TextSetting](docs/TextSetting.md)
 - [TextSettingAllOf](docs/TextSettingAllOf.md)
 - [TimeSetting](docs/TimeSetting.md)
 - [TimeSettingAllOf](docs/TimeSettingAllOf.md)
 - [TimerEvent](docs/TimerEvent.md)
 - [TimerType](docs/TimerType.md)
 - [UninstallData](docs/UninstallData.md)
 - [UpdateData](docs/UpdateData.md)
 - [VideoSetting](docs/VideoSetting.md)
 - [VideoSettingAllOf](docs/VideoSettingAllOf.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



