/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// BasicV2ImageItem Body item object definition for a Basic V2 template
type BasicV2ImageItem struct {
	// URL of image.  HTTPS url is required.
	Url string `json:"url,omitempty"`
	// The text that would be shown under image
	Description string `json:"description,omitempty"`
	BgImage BasicBackgroundImage `json:"bgImage,omitempty"`
	Action Action `json:"action,omitempty"`
}
