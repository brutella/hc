// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeHeatingThresholdTemperature = "00000012-0000-1000-8000-0026BB765291"

type HeatingThresholdTemperature struct {
	*Float
}

func NewHeatingThresholdTemperature() *HeatingThresholdTemperature {
	char := NewFloat(TypeHeatingThresholdTemperature)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(25)
	char.SetStepValue(0.1)
	char.SetValue(0)
	char.Unit = UnitCelsius

	return &HeatingThresholdTemperature{char}
}
