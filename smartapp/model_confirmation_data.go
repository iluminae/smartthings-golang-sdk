/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// ConfirmationData A request to send a confirmation link to a SmartApp for ownership validation. type: object 
type ConfirmationData struct {
	// A globally unique identifier for an app.
	AppId string `json:"appId,omitempty"`
	// An HTTPS url that may be visted to confirm / activate an App registration. 
	ConfirmationUrl string `json:"confirmationUrl,omitempty"`
}
