package accessory

import (
	"github.com/brutella/hc/service"
)

//Window struct
type Window struct {
	*Accessory
	Window *service.Window
}

//NewWindow function
func NewWindow(info Info, setTP, minTP, maxTP, stepTP int) *Window {
	acc := Window{}
	acc.Accessory = New(info, TypeWindow)
	acc.Window = service.NewWindow()

	acc.Window.TargetPosition.SetValue(setTP)
	acc.Window.TargetPosition.SetMinValue(minTP)
	acc.Window.TargetPosition.SetMaxValue(maxTP)
	acc.Window.TargetPosition.SetStepValue(stepTP)

	acc.AddService(acc.Window.Service)

	return &acc
}
