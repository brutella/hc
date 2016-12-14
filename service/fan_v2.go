// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeFanV2 = "B7"

type FanV2 struct {
	*Service

	Active *characteristic.Active
}

func NewFanV2() *FanV2 {
	svc := FanV2{}
	svc.Service = New(TypeFanV2)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	return &svc
}
