package accessory

import (
	"github.com/brutella/hc/service"
)

//LeakSensor struct
type LeakSensor struct {
	*Accessory

	LeakSensor *service.LeakSensor
}

// NewLeakSensor returns a Thermometer which implements model.Thermometer.
func NewLeakSensor(info Info) *LeakSensor {
	acc := LeakSensor{}
	acc.Accessory = New(info, TypeSensor)
	acc.LeakSensor = service.NewLeakSensor()

	acc.AddService(acc.LeakSensor.Service)

	return &acc
}
