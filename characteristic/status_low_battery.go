// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	StatusLowBatteryBatteryLevelNormal int = 0
	StatusLowBatteryBatteryLevelLow    int = 1
)

const TypeStatusLowBattery = "00000079-0000-1000-8000-0026BB765291"

type StatusLowBattery struct {
	*Int
}

func NewStatusLowBattery() *StatusLowBattery {
	char := NewInt(TypeStatusLowBattery)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &StatusLowBattery{char}
}
