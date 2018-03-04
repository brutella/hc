// Created by mestafin
package accessory

import (
	"github.com/brutella/hc/service"
)

type Fan struct {
	*Accessory
	Fan *service.Fan
}

// NewFan returns a fan accessory containing one fan service.
func NewFan(info Info) *Fan {
	acc := Fan{}
	acc.Accessory = New(info, TypeFan)
	acc.Fan = service.NewFan()

	acc.AddService(acc.Fan.Service)

	return &acc
}

