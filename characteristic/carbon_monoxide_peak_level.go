// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCarbonMonoxidePeakLevel = "00000091-0000-1000-8000-0026BB765291"

type CarbonMonoxidePeakLevel struct {
	*Float
}

func NewCarbonMonoxidePeakLevel() *CarbonMonoxidePeakLevel {
	char := NewFloat(TypeCarbonMonoxidePeakLevel)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(0.1)
	char.SetValue(0)

	return &CarbonMonoxidePeakLevel{char}
}
