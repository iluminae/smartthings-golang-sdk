/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// BasicBackgroundImage A background image.
type BasicBackgroundImage struct {
	// URL of image.  HTTPS url is required.
	Url string `json:"url,omitempty"`
	// some color code
	Color string `json:"color,omitempty"`
	// transparency of the color/image
	Alpha float32 `json:"alpha,omitempty"`
}