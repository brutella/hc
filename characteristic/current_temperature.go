// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCurrentTemperature = "11"

type CurrentTemperature struct {
	*Float
}

func NewCurrentTemperature() *CurrentTemperature {
	char := NewFloat(TypeCurrentTemperature)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(0.1)
	char.SetValue(0)
	char.Unit = UnitCelsius

	return &CurrentTemperature{char}
}
