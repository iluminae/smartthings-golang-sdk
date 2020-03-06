# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **string** | The ID of the location. | [optional] 
**Name** | **string** | A name given for the location (eg. Home) | [optional] 
**CountryCode** | **string** | An ISO Alpha-3 country code.  (i.e. GBR, USA) | [optional] 
**Latitude** | **float32** | A geographical latitude. | [optional] 
**Longitude** | **float32** | A geographical longitude. | [optional] 
**RegionRadius** | **int32** | The radius in meters around latitude and longitude which defines this location. | [optional] 
**TemperatureScale** | **string** | The desired temperature scale used within location. Value can be F or C. | [optional] 
**TimeZoneId** | **string** | An ID matching the Java Time Zone ID of the location. Automatically generated if latitude and longitude have been provided.  | [optional] 
**Locale** | **string** | An IETF BCP 47 language tag representing the chosen locale for this location. | [optional] 
**BackgroundImage** | **string** | Not currently in use. | [optional] 
**AdditionalProperties** | **map[string]string** | Additional information about the location that allows SmartThings to further define your location. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


