package main

import (
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/accessory"

	"log"
	"time"
)

func main() {
	switchInfo := accessory.Info{
		Name: "Lamp",
	}
	acc := accessory.NewSwitch(switchInfo)

	config := hap.Config{Pin: "12344321", Port: "12345", StoragePath: "./db"}
	t, err := hap.NewIPTransport(config, acc.Accessory)

	if err != nil {
		log.Fatal(err)
	}

	// Log to console when client (e.g. iOS app) changes the value of the on characteristic
	acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			log.Println("[INFO] Client changed switch to on")
		} else {
			log.Println("[INFO] Client changed switch to off")
		}
	})

	// Periodically toggle the switch's on characteristic
	go func() {
		for {
			on := !acc.Switch.On.GetValue()
			if on == true {
				log.Println("[INFO] Switch is on")
			} else {
				log.Println("[INFO] Switch is off")
			}
			acc.Switch.On.SetValue(on)
			time.Sleep(5 * time.Second)
		}
	}()

	hap.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
