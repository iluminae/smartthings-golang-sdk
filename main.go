package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antihax/optional"
	"github.com/iluminae/smartthings-golang-sdk/smartapp"
	"github.com/iluminae/smartthings-golang-sdk/smartthings"
)

func main() {
	fmt.Println("starting")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		//DEBUG
		fmt.Printf("body: %s\n", string(body))
		//GUBED
		var request smartapp.ExecutionRequest
		if err := json.Unmarshal(body, &request); err != nil {
			fmt.Printf("Error decoding body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		client := smartthings.NewAPIClient(smartthings.NewConfiguration())
		ctx := context.Background()
		response := smartapp.ExecutionResponse{}
		switch request.Lifecycle {
		case smartapp.APPLIFECYCLE_PING:
		case smartapp.APPLIFECYCLE_UPDATE:
			if _, _, err := client.SubscriptionsApi.DeleteAllSubscriptions(ctx, request.AppId, request.UpdateData.AuthToken, nil); err != nil {
				fmt.Printf("Error deleting subscriptions: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			request.InstallData = smartapp.InstallData{
				AuthToken:    request.UpdateData.AuthToken,
				RefreshToken: request.UpdateData.RefreshToken,
				InstalledApp: request.UpdateData.InstalledApp,
			}
			fallthrough
		case smartapp.APPLIFECYCLE_INSTALL:
			for groupName, deviceList := range request.InstallData.InstalledApp.Config {
				for _, device := range deviceList {
					opts := smartthings.SaveSubscriptionOpts{
						Request: optional.NewInterface(
							smartthings.SubscriptionRequest{
								SourceType: smartthings.SUBSCRIPTIONSOURCE_DEVICE,
								Device: smartthings.DeviceSubscriptionDetail{
									DeviceId:        device.DeviceConfig.DeviceId,
									ComponentId:     device.DeviceConfig.ComponentId,
									Capability:      "*", //TODO:: limit to what you need
									Attribute:       "*",
									StateChangeOnly: false,
									SubscriptionName: fmt.Sprintf("%s|%s|%s",
										request.AppId,
										groupName,
										device.DeviceConfig.DeviceId,
									),
									//Modes: []string{},
								},
							},
						),
					}
					if _, _, err := client.SubscriptionsApi.SaveSubscription(ctx, request.AppId, request.InstallData.AuthToken, &opts); err != nil {
						fmt.Printf("Error creating subscriptions: %v", err)
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				}
			}
		case smartapp.APPLIFECYCLE_UNINSTALL:
		case smartapp.APPLIFECYCLE_EVENT:
		case smartapp.APPLIFECYCLE_CONFIRMATION:
		case smartapp.APPLIFECYCLE_OAUTH_CALLBACK:
		case smartapp.APPLIFECYCLE_CONFIGURATION:
			switch request.ConfigurationData.Phase {
			case smartapp.CONFIGURATIONPHASE_INITIALIZE:
				response.ConfigurationData.Initialize = smartapp.InitializeSetting{
					Id:                       "init",
					Name:                     "golang app",
					Description:              "sample golang app",
					Permissions:              []string{"r:devices:*"},
					FirstPageId:              "intro",
					DisableCustomDisplayName: false,
					DisableRemoveApp:         false,
				}
			case smartapp.CONFIGURATIONPHASE_PAGE:
				response.ConfigurationData.Page = Pages[request.ConfigurationData.PageId]
			default:
				panic("unhandled phase " + request.ConfigurationData.Phase)
			}
		default:
			panic("unhandled lifecycle " + request.Lifecycle)
		}
		if response.StatusCode == 0 {
			response.StatusCode = 200
		}
		{ //DEBUG
			jsonBuf, _ := json.Marshal(response)
			fmt.Printf("response: %s\n", string(jsonBuf))
		} //GUBED
		json.NewEncoder(w).Encode(response)
	})

	if err := http.ListenAndServeTLS(":9999", "cert.pem", "key.pem", nil); err != nil {
		panic(err)
	}
}
