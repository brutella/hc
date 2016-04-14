// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeStatelessProgrammableSwitch = "89"

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
