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

	SetDuration       *characteristic.SetDuration
	RemainingDuration *characteristic.RemainingDuration
	IsConfigured      *characteristic.IsConfigured
	ServiceLabelIndex *characteristic.ServiceLabelIndex
	StatusFault       *characteristic.StatusFault
	Name              *characteristic.Name
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

func (svc *Valve) AddOptionalCharaterics() {

	svc.SetDuration = characteristic.NewSetDuration()
	svc.AddCharacteristic(svc.SetDuration.Characteristic)

	svc.RemainingDuration = characteristic.NewRemainingDuration()
	svc.AddCharacteristic(svc.RemainingDuration.Characteristic)

	svc.IsConfigured = characteristic.NewIsConfigured()
	svc.AddCharacteristic(svc.IsConfigured.Characteristic)

	svc.ServiceLabelIndex = characteristic.NewServiceLabelIndex()
	svc.AddCharacteristic(svc.ServiceLabelIndex.Characteristic)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddCharacteristic(svc.StatusFault.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
