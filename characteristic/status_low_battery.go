// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	StatusLowBatteryBatteryLevelNormal int = 0
	StatusLowBatteryBatteryLevelLow    int = 1
)

const TypeStatusLowBattery = "79"

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
