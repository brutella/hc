// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeIrrigationSystem = "CF"

type IrrigationSystem struct {
	*Service

	Active      *characteristic.Active
	ProgramMode *characteristic.ProgramMode
	InUse       *characteristic.InUse
}

func NewIrrigationSystem() *IrrigationSystem {
	svc := IrrigationSystem{}
	svc.Service = New(TypeIrrigationSystem)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.ProgramMode = characteristic.NewProgramMode()
	svc.AddCharacteristic(svc.ProgramMode.Characteristic)

	svc.InUse = characteristic.NewInUse()
	svc.AddCharacteristic(svc.InUse.Characteristic)

	return &svc
}
