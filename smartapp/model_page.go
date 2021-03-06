/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// Page struct for Page
type Page struct {
	// Name of the page to be configured.
	Name string `json:"name,omitempty"`
	// A developer defined page ID. Must be URL safe characters.
	PageId string `json:"pageId,omitempty"`
	// A developer defined page ID for the next page in the configuration process. Must be URL safe characters.
	NextPageId string `json:"nextPageId,omitempty"`
	// A developer defined page ID for the previous page in the configuration process. Must be URL safe characters.
	PreviousPageId string `json:"previousPageId,omitempty"`
	// Indicates if this is the last page in the configuration process.
	Complete bool `json:"complete,omitempty"`
	// Change how the page is presented
	Style string `json:"style,omitempty"`
	// The text for the next button. Only applies if style is `SPLASH`
	NextText string `json:"nextText,omitempty"`
	// The display sections for user defined settings.
	Sections []Section `json:"sections,omitempty"`
}
