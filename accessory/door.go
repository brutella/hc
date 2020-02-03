package accessory

import (
	"github.com/brutella/hc/service"
)

//Door struct
type Door struct {
	*Accessory
	Door *service.Door
}

//NewDoor function
func NewDoor(info Info, setTP, minTP, maxTP, stepTP int) *Door {
	acc := Door{}
	acc.Accessory = New(info, TypeDoor)
	acc.Door = service.NewDoor()

	acc.Door.TargetPosition.SetValue(setTP)
	acc.Door.TargetPosition.SetMinValue(minTP)
	acc.Door.TargetPosition.SetMaxValue(maxTP)
	acc.Door.TargetPosition.SetStepValue(stepTP)

	acc.AddService(acc.Door.Service)
	return &acc
}
