/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// DeviceLifecycleEvent A device lifecycle event.
type DeviceLifecycleEvent struct {
	Lifecycle DeviceLifecycle `json:"lifecycle,omitempty"`
	// The id of the event.
	EventId string `json:"eventId,omitempty"`
	// The id of the location in which the event was triggered.
	LocationId string `json:"locationId,omitempty"`
	// The id of the device.
	DeviceId string `json:"deviceId,omitempty"`
	// The name of the device
	DeviceName string `json:"deviceName,omitempty"`
	// The principal that made the change
	Principal string `json:"principal,omitempty"`
	// Create device lifecycle. 
	Create map[string]interface{} `json:"create,omitempty"`
	// Delete device lifecycle. 
	Delete map[string]interface{} `json:"delete,omitempty"`
	// Update device lifecycle. 
	Update map[string]interface{} `json:"update,omitempty"`
	MoveFrom DeviceLifecycleMove `json:"moveFrom,omitempty"`
	MoveTo DeviceLifecycleMove `json:"moveTo,omitempty"`
}