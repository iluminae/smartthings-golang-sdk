/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// UninstallData The data payload to an execution request with an AppLifecycle of UNINSTALL.
type UninstallData struct {
	InstalledApp InstalledApp `json:"installedApp"`
}
