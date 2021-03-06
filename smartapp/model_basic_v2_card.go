/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// BasicV2Card Data requirements for a Basic V2 card template.
type BasicV2Card struct {
	// An title icon url for card. A HTTPS URL is required.
	IconUrl string `json:"iconUrl,omitempty"`
	// name of the card
	Name string `json:"name,omitempty"`
	// sublined text of the card
	Description string `json:"description,omitempty"`
	BgImage BasicBackgroundImage `json:"bgImage,omitempty"`
	Body BasicV2Body `json:"body,omitempty"`
	// A list of buttons to render and buttons must be of the same type.
	Buttons []BasicV2Button `json:"buttons,omitempty"`
}
