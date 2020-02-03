package accessory

import (
	"github.com/brutella/hc/service"
)

//ContactSensor struct
type ContactSensor struct {
	*Accessory
	ContactSensor *service.ContactSensor
}

//NewContactSensor function
func NewContactSensor(info Info) *ContactSensor {
	acc := ContactSensor{}
	acc.Accessory = New(info, TypeSensor)
	acc.ContactSensor = service.NewContactSensor()

	acc.AddService(acc.ContactSensor.Service)

	return &acc
}
