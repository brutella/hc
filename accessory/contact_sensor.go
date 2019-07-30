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

/*
type ContactSensor struct {
	*Service

	ContactSensorState *characteristic.ContactSensorState
}

func NewContactSensor() *ContactSensor {
	svc := ContactSensor{}
	svc.Service = New(TypeContactSensor)

	svc.ContactSensorState = characteristic.NewContactSensorState()
	svc.AddCharacteristic(svc.ContactSensorState.Characteristic)

	return &svc
}*/
