# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the page to be configured. | [optional] 
**PageId** | **string** | A developer defined page ID. Must be URL safe characters. | [optional] 
**NextPageId** | **string** | A developer defined page ID for the next page in the configuration process. Must be URL safe characters. | [optional] 
**PreviousPageId** | **string** | A developer defined page ID for the previous page in the configuration process. Must be URL safe characters. | [optional] 
**Complete** | **bool** | Indicates if this is the last page in the configuration process. | [optional] [default to false]
**Style** | **string** | Change how the page is presented | [optional] [default to STYLE_NORMAL]
**NextText** | **string** | The text for the next button. Only applies if style is &#x60;SPLASH&#x60; | [optional] 
**Sections** | [**[]Section**](Section.md) | The display sections for user defined settings. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


