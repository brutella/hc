package accessory

import (
	"github.com/brutella/hc/service"
)

//HumidifierDehumidifier struct
type HumidifierDehumidifier struct {
	*Accessory

	HumidifierDehumidifier *service.HumidifierDehumidifier
}

// NewHumidifierDehumidifier returns.
func NewHumidifierDehumidifier(info Info,
	thdState, thdStateMin, thdStateMax, thdStateStep int) *HumidifierDehumidifier {
	acc := HumidifierDehumidifier{}
	acc.Accessory = New(info, TypeDehumidifier)
	acc.HumidifierDehumidifier = service.NewHumidifierDehumidifier()

	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetValue(thdState)
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMinValue(thdStateMin)
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMaxValue(thdStateMax)
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetStepValue(thdStateStep)

	acc.AddService(acc.HumidifierDehumidifier.Service)

	return &acc
}
