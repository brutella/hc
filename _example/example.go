package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/log"

	"time"
)

func main() {
	log.Debug.Enable()
	switchInfo := accessory.Info{
		Name: "Lamp",
	}
	acc := accessory.NewSwitch(switchInfo)

	config := hc.Config{Pin: "12344321", Port: "12345", StoragePath: "./db"}
	t, err := hc.NewIPTransport(config, acc.Accessory)

	if err != nil {
		log.Info.Panic(err)
	}

	// Log to console when client (e.g. iOS app) changes the value of the on characteristic
	acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			log.Debug.Println("Client changed switch to on")
		} else {
			log.Debug.Println("Client changed switch to off")
		}
	})

	// Periodically toggle the switch's on characteristic
	go func() {
		for {
			on := !acc.Switch.On.GetValue()
			if on == true {
				log.Debug.Println("Switch is on")
			} else {
				log.Debug.Println("Switch is off")
			}
			acc.Switch.On.SetValue(on)
			time.Sleep(5 * time.Second)
		}
	}()

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
