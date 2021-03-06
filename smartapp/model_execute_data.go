/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// ExecuteData The data payload to an execution request with an AppLifecycle of EXECUTE.
type ExecuteData struct {
	// An OAuth token to use when calling into SmartThings API's.
	AuthToken string `json:"authToken"`
	InstalledApp InstalledApp `json:"installedApp"`
	// An arbitrary map of input parameters which the SmartApp can use to build a custom response.
	Parameters map[string]string `json:"parameters"`
}
