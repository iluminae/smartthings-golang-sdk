/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// UpdateData The data payload to an execution request with an AppLifecycle of UPDATE.
type UpdateData struct {
	// An OAuth token to use when calling into SmartThings API's.
	AuthToken string `json:"authToken"`
	// A refresh token which maybe used to obtain authorization to SmartThings API after expiration of the authToken. An integration will need to use this refreshToken to support calling the SmartThings API outside the context of an event. 
	RefreshToken string `json:"refreshToken,omitempty"`
	InstalledApp InstalledApp `json:"installedApp"`
	// A map of configurations for an Installed App.  The map 'key' is the configuration name and the 'value' is an array of strings. 
	PreviousConfig map[string][]ConfigEntry `json:"previousConfig"`
	// A list of permissions associated with this execution. See `securityDefinitions` for more information.
	PreviousPermissions []string `json:"previousPermissions"`
}
