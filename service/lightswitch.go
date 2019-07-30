// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLightswitch = "43"

type Lightswitch struct {
	*Service

	On *characteristic.On
}

func NewLightswitch() *Lightswitch {
	svc := Lightswitch{}
	svc.Service = New(TypeLightswitch)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	return &svc
}
