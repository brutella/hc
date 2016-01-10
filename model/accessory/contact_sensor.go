package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/service"
)

type contactSensor struct {
	*Accessory
	contactSensor *service.ContactSensor
}

// NewContactSensor returns a contact sensor which implements model.ContactSensor.
func NewContactSensor(info model.Info) *contactSensor {
	accessory := New(info)
	s := service.NewContactSensor(info.Name)

	accessory.AddService(s.Service)

	cs := contactSensor{accessory, s}

	return &cs
}

func (cs *contactSensor) SetState(state model.ContactSensorStateType) {
	cs.contactSensor.ContactSensorState.SetContactSensorState(state)
}

func (cs *contactSensor) State() model.ContactSensorStateType {
	return cs.contactSensor.ContactSensorState.ContactSensorState()
}