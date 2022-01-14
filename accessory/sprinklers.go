package accessory

import (
	"github.com/brutella/hc/service"
)

type Sprinklers struct {
	*Accessory
	IrrigationSystem *service.IrrigationSystem
}

func NewSprinklers(info Info) *Sprinklers {
	acc := Sprinklers{}
	acc.Accessory = New(info, TypeSprinklers)
	acc.IrrigationSystem = service.NewIrrigationSystem()

	acc.AddService(acc.IrrigationSystem.Service)

	return &acc
}
