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
func NewHumidifierDehumidifier(info Info, stateTHDS, minTHDS, maxTHDS, stepTHDS int) *HumidifierDehumidifier {
	acc := HumidifierDehumidifier{}
	acc.Accessory = New(info, TypeDehumidifier)
	acc.HumidifierDehumidifier = service.NewHumidifierDehumidifier()

	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetValue(stateTHDS)
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMinValue(minTHDS)
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMaxValue(maxTHDS)
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetStepValue(stepTHDS)

	acc.AddService(acc.HumidifierDehumidifier.Service)

	return &acc
}
