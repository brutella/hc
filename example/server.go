package main

import(
    "fmt"
    "gohap"
)

func main() {
    accessory, _ := gohap.NewAccessory("HAP Test", "123-45-678")    
    controller, _ := gohap.NewSetupController(accessory)
    fmt.Println(controller)
}