package main

import (
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"

	"log"
	"time"
)

func main() {
	switchInfo := model.Info{
		Name: "Lamp",
	}
	sw := accessory.NewSwitch(switchInfo)

	config := hap.Config{Pin: "12344321", Port: "12345", StoragePath: "db"}
	t, err := hap.NewIPTransport(config, sw.Accessory)

	if err != nil {
		log.Fatal(err)
	}

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

	hap.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
