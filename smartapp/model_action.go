/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// Action A definition of the action to be taken when button is activated.
type Action struct {
	Type ActionType `json:"type,omitempty"`
	LaunchPlugin LaunchPluginAction `json:"launchPlugin,omitempty"`
	Execute ExecuteAction `json:"execute,omitempty"`
}