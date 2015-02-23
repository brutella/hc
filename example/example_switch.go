package main

import (
	"github.com/brutella/hap/app"
	"github.com/brutella/hap/model"
	"github.com/brutella/hap/model/accessory"
	"github.com/brutella/log"
	_ "time"
)

func main() {
	log.Verbose = false

	conf := app.NewConfig()
	conf.DatabaseDir = "./data"

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

	sw := accessory.NewSwitch(info)
	sw.OnStateChanged(func(on bool) {
		if on == true {
			log.Println("[INFO] Switch on")
		} else {
			log.Println("[INFO] Switch off")
		}
	})

	// go func() {
	//   for {
	//       on := !sw.IsOn()
	//       if on == true {
	//           log.Println("[INFO] Switch on")
	//       } else {
	//           log.Println("[INFO] Switch off")
	//       }
	//       sw.SetOn(on)
	//       time.Sleep(5 * time.Second)
	//   }
	// }()

	// go func() {
	//   for {
	//       time.Sleep(20 * time.Second)
	//       log.Println("[VERB] Set unreachable")
	//       app.SetReachable(false)
	//       time.Sleep(20 * time.Second)
	//       log.Println("[VERB] Set reachable")
	//       app.SetReachable(true)
	//   }
	// }()

	app.AddAccessory(sw.Accessory)

	app.Run()
}
