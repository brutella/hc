// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeOutlet = "00000047-0000-1000-8000-0026BB765291"

type Outlet struct {
	*Service

	On          *characteristic.On
	OutletInUse *characteristic.OutletInUse
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
