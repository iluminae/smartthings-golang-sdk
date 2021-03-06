/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// DashboardCard payload of dashboard card
type DashboardCard struct {
	TemplateId DashboardCardTemplate `json:"templateId"`
	// A unique identifier for a card of service
	CardId string `json:"cardId"`
	BasicV1 BasicCard `json:"basicV1,omitempty"`
	BasicV2 BasicV2Card `json:"basicV2,omitempty"`
	FreeForm FreeFormCard `json:"freeForm,omitempty"`
}
