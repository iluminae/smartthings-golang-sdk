/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// ActionType The type of action to take when UI element is activated.  * LAUNCH_PLUGIN - Launch a corresponding UI plugin. * EXECUTE - Execute an API call to the backing SmartApp. 
type ActionType string

// List of ActionType
const (
	ACTIONTYPE_LAUNCH_PLUGIN ActionType = "LAUNCH_PLUGIN"
	ACTIONTYPE_EXECUTE ActionType = "EXECUTE"
)
