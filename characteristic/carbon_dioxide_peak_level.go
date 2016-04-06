// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCarbonDioxidePeakLevel = "00000094-0000-1000-8000-0026BB765291"

type CarbonDioxidePeakLevel struct {
	*Float
}

func NewCarbonDioxidePeakLevel() *CarbonDioxidePeakLevel {
	char := NewFloat(TypeCarbonDioxidePeakLevel)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100000)
	char.SetStepValue(100)
	char.SetValue(0)

	return &CarbonDioxidePeakLevel{char}
}
