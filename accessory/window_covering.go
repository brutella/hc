// Created by mestafin
package accessory

import (
	"github.com/brutella/hc/service"
)

type WindowCovering struct {
	*Accessory

	WindowCovering *service.WindowCovering
}

// NewWindowCovering returns a WindowCovering which implements model.WindowCovering.
func NewWindowCovering(info Info) *WindowCovering {
	acc := WindowCovering{}
	acc.Accessory = New(info, TypeWindowCovering)
	
	acc.WindowCovering = service.NewWindowCovering()

	acc.AddService(acc.WindowCovering.Service)

	return &acc
}
