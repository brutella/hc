package accessory

import (
	"github.com/brutella/hc/service"
)

type Windows struct {
	*Accessory
	Window *service.Window
}

// NewWindow returns a window which implements model.NewWindow.
func NewWindow(info Info, currentState int) *Windows {
	acc := Windows{}
	acc.Accessory = New(info, TypeWindow)
	acc.Window = service.NewWindow()
	acc.Window.CurrentPosition.SetValue(currentState)
	acc.AddService(acc.Window.Service)

	return &acc
}
