/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// BooleanSettingAllOf struct for BooleanSettingAllOf
type BooleanSettingAllOf struct {
	// A default value for the setting.
	DefaultValue string `json:"defaultValue,omitempty"`
	// The image url.
	Image string `json:"image,omitempty"`
	// Indicates if this input should refresh configs after a change in value.
	SubmitOnChange bool `json:"submitOnChange,omitempty"`
}
