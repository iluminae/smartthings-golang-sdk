# AdhocMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FallbackLocale** | **string** | The tag of the locale as defined in [RFC bcp47](http://www.rfc-editor.org/rfc/bcp/bcp47.txt). | 
**DefaultVariables** | **map[string]string** | A map&lt;string,string&gt; with the key representing the variable name, and the value representing the verbiage to be replaced in template string. &#x60;defaultVariables&#x60; will only be used if there are no matching locale-level (template) variables for that key.  | [optional] 
**Templates** | [**[]AdhocMessageTemplate**](AdhocMessageTemplate.md) | A list of templates representing the same message in different languages. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


