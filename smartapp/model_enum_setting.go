/*
 * SmartThings SmartApps API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartapp
// EnumSetting Enum Setting
type EnumSetting struct {
	// A developer defined configuration ID.
	Id string `json:"id,omitempty"`
	// Name of the setting to be configured.
	Name string `json:"name,omitempty"`
	// Description of the app to be configured.
	Description string `json:"description,omitempty"`
	// Indicates if this setting is required for configuration.
	Required bool `json:"required,omitempty"`
	// Indicates if this setting should be disabled
	Disabled bool `json:"disabled,omitempty"`
	Type SettingType `json:"type"`
	// A default value for the setting.
	DefaultValue string `json:"defaultValue,omitempty"`
	// Indicates if this enum setting can have multiple values.
	Multiple bool `json:"multiple,omitempty"`
	// Indicates if this input should close on selection.
	CloseOnSelection bool `json:"closeOnSelection,omitempty"`
	// Display the enum options as groups.
	GroupedOptions []GroupedOption `json:"groupedOptions,omitempty"`
	// The enum options.
	Options []Option `json:"options,omitempty"`
	// Indicates if this input should refresh configs after a change in value.
	SubmitOnChange bool `json:"submitOnChange,omitempty"`
	Style EnumStyleType `json:"style,omitempty"`
}
