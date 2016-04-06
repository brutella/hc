// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCoolingThresholdTemperature = "0000000D-0000-1000-8000-0026BB765291"

type CoolingThresholdTemperature struct {
	*Float
}

func NewCoolingThresholdTemperature() *CoolingThresholdTemperature {
	char := NewFloat(TypeCoolingThresholdTemperature)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(10)
	char.SetMaxValue(35)
	char.SetStepValue(0.1)
	char.SetValue(10)
	char.Unit = UnitCelsius

	return &CoolingThresholdTemperature{char}
}
