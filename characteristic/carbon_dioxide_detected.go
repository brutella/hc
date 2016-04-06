// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CarbonDioxideDetectedCO2LevelsNormal   int = 0
	CarbonDioxideDetectedCO2LevelsAbnormal int = 1
)

const TypeCarbonDioxideDetected = "00000092-0000-1000-8000-0026BB765291"

type CarbonDioxideDetected struct {
	*Int
}

func NewCarbonDioxideDetected() *CarbonDioxideDetected {
	char := NewInt(TypeCarbonDioxideDetected)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &CarbonDioxideDetected{char}
}
