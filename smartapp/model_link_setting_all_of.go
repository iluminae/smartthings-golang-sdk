/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// LinkSettingAllOf struct for LinkSettingAllOf
type LinkSettingAllOf struct {
	// A default value for the setting.
	DefaultValue string `json:"defaultValue,omitempty"`
	// The page to navigate to.
	Url string `json:"url,omitempty"`
	// The image url.
	Image string `json:"image,omitempty"`
	Style LinkStyleType `json:"style,omitempty"`
}
