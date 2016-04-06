// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeStatelessProgrammableSwitch = "00000089-0000-1000-8000-0026BB765291"

type StatelessProgrammableSwitch struct {
	*Service

	ProgrammableSwitchEvent *characteristic.ProgrammableSwitchEvent
}

func NewStatelessProgrammableSwitch() *StatelessProgrammableSwitch {
	svc := StatelessProgrammableSwitch{}
	svc.Service = New(TypeStatelessProgrammableSwitch)

	svc.ProgrammableSwitchEvent = characteristic.NewProgrammableSwitchEvent()
	svc.AddCharacteristic(svc.ProgrammableSwitchEvent.Characteristic)

	return &svc
}
