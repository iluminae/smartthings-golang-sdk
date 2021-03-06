/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// PingData A request to ping a SmartApp to ensure connectivity. 
type PingData struct {
	// A challenge phrase that the SmartApp must echo back to validate itself.
	Challenge string `json:"challenge"`
}
