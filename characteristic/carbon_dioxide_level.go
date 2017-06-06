// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCarbonDioxideLevel = "93"

type CarbonDioxideLevel struct {
	*Float
}

func NewCarbonDioxideLevel() *CarbonDioxideLevel {
	char := NewFloat(TypeCarbonDioxideLevel)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100000)

	char.SetValue(0)

	return &CarbonDioxideLevel{char}
}
