package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/log"

	"fmt"
	"net/http"
)

var WebSwitchState bool
var accessoryDevice *accessory.Switch

func toggleHandler(w http.ResponseWriter, r *http.Request) {
	WebSwitchState = accessoryDevice.Switch.On.GetValue()
	if WebSwitchState == true {
		log.Debug.Println("Switch is on")
	} else {
		log.Debug.Println("Switch is off")
	}
	accessoryDevice.Switch.On.SetValue(!WebSwitchState)
	fmt.Fprintf(w, "Hello World\nSwitch state is now: %v\n", !WebSwitchState)
}

func turnOnHandler(w http.ResponseWriter, r *http.Request) {
	accessoryDevice.Switch.On.SetValue(true)
	fmt.Fprintf(w, "Hello World\nSwitch state is now: true\n")
}

func turnOffHandler(w http.ResponseWriter, r *http.Request) {
	accessoryDevice.Switch.On.SetValue(false)
	fmt.Fprintf(w, "Hello World\nSwitch state is now: false\n")
}

func main() {
	log.Debug.Enable()
	switchInfo := accessory.Info{
		Name: "WebSwitch",
		//SerialNumber:     "00001",
		//Manufacturer:     "",
		//Model:            "",
		//FirmwareRevision: "0.0.1",
		//ID:               123456,
	}
	accessoryDevice = accessory.NewSwitch(switchInfo)

	config := hc.Config{Pin: "12344321", Port: "12345", StoragePath: "./db"}
	t, err := hc.NewIPTransport(config, accessoryDevice.Accessory)

	if err != nil {
		log.Info.Panic(err)
	}

	// Log to console when client (e.g. iOS app) changes the value of the on characteristic
	accessoryDevice.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			log.Debug.Println("Client changed switch to on")
		} else {
			log.Debug.Println("Client changed switch to off")
		}
	})

	hc.OnTermination(func() {
		<-t.Stop()
	})

	//http handler to toggle the switch
	//send web request to http://localhost:8080/toggle to toggle the switch
	http.HandleFunc("/toggle", toggleHandler)

	//http handler to turn on the switch
	//send web request to http://localhost:8080/turnon to turn the switch on
	http.HandleFunc("/turnon", turnOnHandler)

	//http handler to turn off the switch
	//send web request to http://localhost:8080/turnoff to turn the switch off
	http.HandleFunc("/turnoff", turnOffHandler)

	//start the webserver and listen on port 8080
	go http.ListenAndServe(":8080", nil)

	t.Start()
}
