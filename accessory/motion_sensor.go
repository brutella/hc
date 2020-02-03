package accessory

import (
	"github.com/brutella/hc/service"
)

//MotionSensor struct
type MotionSensor struct {
	*Accessory
	MotionSensor *service.MotionSensor
}

//NewMotionSensor function
func NewMotionSensor(info Info) *MotionSensor {
	acc := MotionSensor{}
	acc.Accessory = New(info, TypeSensor)
	acc.MotionSensor = service.NewMotionSensor()

	acc.AddService(acc.MotionSensor.Service)

	return &acc
}
