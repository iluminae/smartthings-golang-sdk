/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// OAuthCallbackData Provides intergration with the result of a third party oauth attempt.  This will only be available for executions of type \"OAUTH_CALLBACK\". 
type OAuthCallbackData struct {
	// The id of the installed app.
	InstalledAppId string `json:"installedAppId,omitempty"`
	// A relative URL containing all of the query string parameters as returned by the third party oauth system. A SmartApp can parse the `urlPath` property to extract any senstive auth codes/tokens which can then be used to access the third party system. 
	UrlPath string `json:"urlPath,omitempty"`
}
