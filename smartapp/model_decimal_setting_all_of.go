/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// DecimalSettingAllOf struct for DecimalSettingAllOf
type DecimalSettingAllOf struct {
	// A default value for the setting.
	DefaultValue string `json:"defaultValue,omitempty"`
	// The maximum inclusive value the decimal can be set to.
	Max int32 `json:"max,omitempty"`
	// The minumum inclusive value the decimal can be set to.
	Min int32 `json:"min,omitempty"`
	// The image url.
	Image string `json:"image,omitempty"`
	// A string to be shown after the text input field.
	PostMessage string `json:"postMessage,omitempty"`
}
