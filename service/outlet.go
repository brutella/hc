// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeOutlet = "47"

type Outlet struct {
	*Service

	On          *characteristic.On
	OutletInUse *characteristic.OutletInUse

	Name *characteristic.Name
}

func NewOutlet() *Outlet {
	svc := Outlet{}
	svc.Service = New(TypeOutlet)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.OutletInUse = characteristic.NewOutletInUse()
	svc.AddCharacteristic(svc.OutletInUse.Characteristic)

	return &svc
}

func (svc *Outlet) AddOptionalCharaterics() {

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
