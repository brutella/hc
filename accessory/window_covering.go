package accessory

import (
	"github.com/brutella/hc/service"
)

//WindowCovering struct
type WindowCovering struct {
	*Accessory
	WindowCovering *service.WindowCovering
}

//NewWindowCovering function
func NewWindowCovering(info Info, setTP, minTP, maxTP, stepTP int) *WindowCovering {
	acc := WindowCovering{}
	acc.Accessory = New(info, TypeWindowCovering)
	acc.WindowCovering = service.NewWindowCovering()

	acc.WindowCovering.TargetPosition.SetValue(setTP)
	acc.WindowCovering.TargetPosition.SetMinValue(minTP)
	acc.WindowCovering.TargetPosition.SetMaxValue(maxTP)
	acc.WindowCovering.TargetPosition.SetStepValue(stepTP)

	acc.AddService(acc.WindowCovering.Service)

	return &acc
}
