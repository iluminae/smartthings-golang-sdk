/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// Section struct for Section
type Section struct {
	// Name of the section.
	Name string `json:"name,omitempty"`
	// Whether or not the section can be collapsed
	Hideable bool `json:"hideable,omitempty"`
	// If section can be collapsed, whether or not it defaults to hidden
	Hidden bool `json:"hidden,omitempty"`
	// Change how the section is presented
	Style string `json:"style,omitempty"`
	// Configuration settings represent the questions asked to the end user installing an integration the answers to which provide the configuration for which the integration will use when executing. Settings follow an inheritance pattern.  The type field dictates the expected instance of setting that is provided. 
	Settings []SectionSettingInterface `json:"settings,omitempty"`
}
