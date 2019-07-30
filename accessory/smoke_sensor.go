package accessory

import (
	"github.com/brutella/hc/service"
)

//SmokeSensor struct
type SmokeSensor struct {
	*Accessory
	SmokeSensor *service.SmokeSensor
}

//NewSmokeSensor function
func NewSmokeSensor(info Info) *SmokeSensor {
	acc := SmokeSensor{}
	acc.Accessory = New(info, TypeSensor)
	acc.SmokeSensor = service.NewSmokeSensor()

	acc.AddService(acc.SmokeSensor.Service)

	return &acc
}
