package accessory

import (
	"github.com/brutella/hc/service"
)

type ContactSensor struct {
	*Accessory
	ContactSensor *service.ContactSensor
}

// NewContactSensor returns an accessory containing a service.
func NewContactSensor(info Info) *ContactSensor {
	acc := ContactSensor{}
	acc.Accessory = New(info, TypeSensor)
	acc.ContactSensor = service.NewContactSensor()

	acc.AddService(acc.ContactSensor.Service)

	return &acc
}
