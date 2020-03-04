package main

type Request struct {
	Lifecycle         string            `json:"lifecycle,omitempty"`
	ExecutionId       string            `json:"executionId,omitempty"`
	Locale            string            `json:"locale,omitempty"`
	Version           string            `json:"version,omitempty"`
	PingData          PingData          `json:"pingData,omitempty"`
	ConfigurationData ConfigurationData `json:"configurationData,omitempty"`
	Settings          map[string]string `json:"settings,omitempty"`
	StatusCode        int               `json:"statusCode,omitempty"`
}
type PingData struct {
	Challenge string `json:"challenge,omitempty"`
}
type ConfigurationData struct {
	IsResubmit     bool        `json:"isResubmit,omitempty"`
	Phase          string      `json:"phase,omitempty"`
	InstalledAppId string      `json:"installedAppId,omitempty"`
	PageId         string      `json:"pageId,omitempty"`
	PreviousPageId string      `json:"previousPageId,omitempty"`
	Initialize     Initialize  `json:"initialize,omitempty"`
	Page           Page        `json:"page,omitempty"`
	Config         interface{} `json:"config,omitempty"`
}
type Page struct {
	Name           string    `json:"name"`
	PageId         string    `json:"pageId"`
	PreviousPageId string    `json:"previousPageId,omitempty"`
	NextPageId     string    `json:"nextPageId,omitempty"`
	Complete       bool      `json:"complete"`
	Style          string    `json:"style"`
	NextText       string    `json:"nextText,omitempty"`
	Sections       []Section `json:"sections"`
}
type Initialize struct {
	Id                       string   `json:"id,omitempty"`
	Name                     string   `json:"name,omitempty"`
	Description              string   `json:"description,omitempty"`
	FirstPageId              string   `json:"firstPageId,omitempty"`
	Permissions              []string `json:"permissions,omitempty"`
	DisableCustomDisplayName bool     `json:"disableCustomDisplayName,omitempty"`
	DisableRemoveApp         bool     `json:"disableRemoveApp,omitempty"`
}
type Section struct {
	Name     string            `json:"name"`
	Hideable bool              `json:"hideable"`
	Hidden   bool              `json:"hidden"`
	Style    string            `json:"style"`
	Settings []SectionSettings `json:"settings,omitempty"`
}
type SectionSettings interface{
	isSectionSettings()
}
type CommonSettings struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Disabled    bool   `json:"disabled"`
}
type DeviceSettings struct {
	CommonSettings
	Capabilities []string `json:"capabilities"`
	Permissions []string `json:"permissions"`
}
func(DeviceSettings) isSectionSettings(){}
