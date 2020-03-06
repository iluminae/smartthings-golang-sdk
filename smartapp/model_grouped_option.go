/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// GroupedOption struct for GroupedOption
type GroupedOption struct {
	// The display name of this group of enum options.
	Name string `json:"name,omitempty"`
	// The enum options.
	Options []Option `json:"options,omitempty"`
}