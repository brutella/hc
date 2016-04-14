// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CarbonMonoxideDetectedCOLevelsNormal   int = 0
	CarbonMonoxideDetectedCOLevelsAbnormal int = 1
)

const TypeCarbonMonoxideDetected = "69"

type CarbonMonoxideDetected struct {
	*Int
}

func NewCarbonMonoxideDetected() *CarbonMonoxideDetected {
	char := NewInt(TypeCarbonMonoxideDetected)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &CarbonMonoxideDetected{char}
}
