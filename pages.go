package main

import "github.com/iluminae/smartthings-golang-sdk/smartapp"

var Pages = map[string]smartapp.Page{"intro": {
	PageId:   "intro",
	Name:     "Install Golang App",
	Complete: true,
	Style:    "NORMAL",
	Sections: []smartapp.Section{
		{
			Name:  "Sensors",
			Style: "NORMAL",
			Settings: []smartapp.SectionSettingInterface{
				smartapp.DeviceSetting{
					Id:           "sensors",
					Name:         "Sensors",
					Description:  "Select Sensors",
					Type:         smartapp.SETTINGTYPE_DEVICE,
					Required:     false,
					Disabled:     false,
					Multiple:     true,
					Capabilities: []string{"switch"},
					Permissions:  []string{"r"},
				},
			},
		},
		{
			Name:  "Lights",
			Style: "NORMAL",
			Settings: []smartapp.SectionSettingInterface{
				smartapp.DeviceSetting{
					Id:           "lights",
					Name:         "Lights",
					Description:  "Select Lights",
					Type:         smartapp.SETTINGTYPE_DEVICE,
					Required:     false,
					Disabled:     false,
					Multiple:     true,
					Capabilities: []string{"light"},
					Permissions:  []string{"r"},
				},
			},
		},
	},
},
}
