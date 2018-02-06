// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeValve = "D0"

type Valve struct {
	*Service

	Active    *characteristic.Active
	InUse     *characteristic.InUse
	ValveType *characteristic.ValveType
}

func NewValve() *Valve {
	svc := Valve{}
	svc.Service = New(TypeValve)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.InUse = characteristic.NewInUse()
	svc.AddCharacteristic(svc.InUse.Characteristic)

	svc.ValveType = characteristic.NewValveType()
	svc.AddCharacteristic(svc.ValveType.Characteristic)

	return &svc
}
