package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("starting")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		fmt.Printf("body: %s\n", string(body))
		var request Request
		if err := json.Unmarshal(body, &request); err != nil {
			fmt.Printf("Error decoding body: %v", err)
			http.Error(w, "can't decode body", http.StatusBadRequest)
			return
		}
		switch request.Lifecycle {
		case "PING":
		case "INSTALL":
			//TODO:: write down creds here
		case "UNINSTALL":
		case "CONFIGURATION":
			switch request.ConfigurationData.Phase {
			case "INITIALIZE":
				request.ConfigurationData.Initialize = Initialize{
					Id:                       "init",
					Name:                     "golang app",
					Description:              "sample golang app",
					Permissions:              []string{"r:devices:*"},
					FirstPageId:              "intro",
					DisableCustomDisplayName: false,
					DisableRemoveApp:         false,
				}
			case "PAGE":
				request.ConfigurationData.Page = Pages[request.ConfigurationData.PageId]
			default:
				panic("unhandled phase " + request.ConfigurationData.Phase)
			}
		default:
			panic("unhandled lifecycle " + request.Lifecycle)
		}
		if request.StatusCode == 0 {
			request.StatusCode = 200
		}
		//DEBUG
		jsonBuf, _ := json.Marshal(request)
		fmt.Printf("response: %s\n", string(jsonBuf))
		//GUBED
		json.NewEncoder(w).Encode(request)
	})

	if err := http.ListenAndServeTLS(":9999", "cert.pem", "key.pem", nil); err != nil {
		panic(err)
	}
}
