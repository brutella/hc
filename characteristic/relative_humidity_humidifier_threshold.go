// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeRelativeHumidityHumidifierThreshold = "CA"

type RelativeHumidityHumidifierThreshold struct {
	*Float
}

func NewRelativeHumidityHumidifierThreshold() *RelativeHumidityHumidifierThreshold {
	char := NewFloat(TypeRelativeHumidityHumidifierThreshold)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(1)
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &RelativeHumidityHumidifierThreshold{char}
}
