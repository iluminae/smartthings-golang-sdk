package main

var Pages = map[string]Page{"intro": {
	PageId:   "intro",
	Name:     "Install Golang App",
	Complete: true,
	Style:    "NORMAL",
	Sections: []Section{
		{
			Name:  "Sensors",
			Style: "NORMAL",
			Settings: []SectionSettings{
				DeviceSettings{
					CommonSettings: CommonSettings{
						Id:          "sensors",
						Name:        "Sensors",
						Description: "Select Sensors",
						Type:        "DEVICE",
						Required:    false,
						Disabled:    false,
					},
					Capabilities: []string{"switch"},
					Permissions:  []string{"r"},
				},
			},
		},
		{
			Name:  "Lights",
			Style: "NORMAL",
			Settings: []SectionSettings{
				DeviceSettings{
					CommonSettings: CommonSettings{
						Id:          "lights",
						Name:        "Lights",
						Description: "Select Lights",
						Type:        "DEVICE",
						Required:    false,
						Disabled:    false,
					},
					Capabilities: []string{"switch"},
					Permissions:  []string{"r"},
				},
			},
		},
	},
},
}
