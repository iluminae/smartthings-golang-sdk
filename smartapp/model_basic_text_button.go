/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// BasicTextButton A text button definition for a basic template.
type BasicTextButton struct {
	// The name of the button
	Name string `json:"name,omitempty"`
	Position BasicButtonPosition `json:"position,omitempty"`
	Action Action `json:"action,omitempty"`
}
