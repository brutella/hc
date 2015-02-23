package main

import (
	"github.com/brutella/hap/app"
	"github.com/brutella/hap/model"
	"github.com/brutella/hap/model/accessory"
	"github.com/brutella/log"
	"time"
)

// This sample demonstrates how to create a HomeKit bridge for a switch accessory
// which periodically toggles the switch's on state.
func main() {
	// Disable verbose logging
	log.Verbose = false

	conf := app.NewConfig()
	// Path to database directory to store bridge informations (serial number, crypto keys,...)
	conf.DatabaseDir = "./data"

	// Create a new app
	app, err := app.NewApp(conf)
	if err != nil {
		log.Fatal(err)
	}

	info := model.Info{
		Name:         "My Switch",
		SerialNumber: "001",
		Manufacturer: "Google",
		Model:        "Switchy",
	}

	// Create a switch accessory
	sw := accessory.NewSwitch(info)
	// Log to console when client (e.g. iOS app) changes the value of the on characteristic
	sw.OnStateChanged(func(on bool) {
		if on == true {
			log.Println("[INFO] Client changed switch to on")
		} else {
			log.Println("[INFO] Client changed switch to off")
		}
	})

	// Periodically toggle the switch's on characteristic
	go func() {
		for {
			on := !sw.IsOn()
			if on == true {
				log.Println("[INFO] Switch on")
			} else {
				log.Println("[INFO] Switch off")
			}
			sw.SetOn(on)
			time.Sleep(5 * time.Second)
		}
	}()

	// Add the switch to the app
	app.AddAccessory(sw.Accessory)

	// Run the app
	app.Run()
}
