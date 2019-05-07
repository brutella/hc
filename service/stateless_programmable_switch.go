// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeStatelessProgrammableSwitch = "89"

type StatelessProgrammableSwitch struct {
	*Service

	ProgrammableSwitchEvent *characteristic.ProgrammableSwitchEvent

	Name              *characteristic.Name
	ServiceLabelIndex *characteristic.ServiceLabelIndex
}

func NewStatelessProgrammableSwitch() *StatelessProgrammableSwitch {
	svc := StatelessProgrammableSwitch{}
	svc.Service = New(TypeStatelessProgrammableSwitch)

	svc.ProgrammableSwitchEvent = characteristic.NewProgrammableSwitchEvent()
	svc.AddCharacteristic(svc.ProgrammableSwitchEvent.Characteristic)

	return &svc
}

func (svc *StatelessProgrammableSwitch) AddOptionalCharacteristics() {

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.ServiceLabelIndex = characteristic.NewServiceLabelIndex()
	svc.AddCharacteristic(svc.ServiceLabelIndex.Characteristic)

}
