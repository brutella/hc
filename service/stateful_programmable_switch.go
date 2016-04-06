// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeStatefulProgrammableSwitch = "00000088-0000-1000-8000-0026BB765291"

type StatefulProgrammableSwitch struct {
	*Service

	ProgrammableSwitchEvent       *characteristic.ProgrammableSwitchEvent
	ProgrammableSwitchOutputState *characteristic.ProgrammableSwitchOutputState
}

func NewStatefulProgrammableSwitch() *StatefulProgrammableSwitch {
	svc := StatefulProgrammableSwitch{}
	svc.Service = New(TypeStatefulProgrammableSwitch)

	svc.ProgrammableSwitchEvent = characteristic.NewProgrammableSwitchEvent()
	svc.AddCharacteristic(svc.ProgrammableSwitchEvent.Characteristic)

	svc.ProgrammableSwitchOutputState = characteristic.NewProgrammableSwitchOutputState()
	svc.AddCharacteristic(svc.ProgrammableSwitchOutputState.Characteristic)

	return &svc
}
