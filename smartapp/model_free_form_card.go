/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// FreeFormCard struct for FreeFormCard
type FreeFormCard struct {
	// An arbitrary set of key / value pairs useful for passing any custom metadata.  * Supports a maximum of 5 entries. * Maximum key length: 36 Unicode characters in UTF-8 * Maximum value length: 1000 Unicode characters in UTF-8 * Allowed characters for *keys* are letters, plus the following special characters: `:`, `_` * Allowed characters for *values* are letters, whitespace, and numbers, plus the following special characters: `+`, `-`, `=`, `.`, `_`, `:`, `/` * If you need characters outside this allowed set, you can apply standard base-64 encoding. 
	Attributes map[string]string `json:"attributes,omitempty"`
}
