// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TemperatureDisplayUnitsCelsius    int = 0
	TemperatureDisplayUnitsFahrenheit int = 1
)

const TypeTemperatureDisplayUnits = "00000036-0000-1000-8000-0026BB765291"

type TemperatureDisplayUnits struct {
	*Int
}

func NewTemperatureDisplayUnits() *TemperatureDisplayUnits {
	char := NewInt(TypeTemperatureDisplayUnits)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &TemperatureDisplayUnits{char}
}
