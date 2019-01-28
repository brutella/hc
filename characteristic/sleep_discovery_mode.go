// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	SleepDiscoveryModeNotDiscoverable    int = 0
	SleepDiscoveryModeAlwaysDiscoverable int = 1
)

const TypeSleepDiscoveryMode = "E8"

type SleepDiscoveryMode struct {
	*Int
}

func NewSleepDiscoveryMode() *SleepDiscoveryMode {
	char := NewInt(TypeSleepDiscoveryMode)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(1)

	char.SetValue(0)

	return &SleepDiscoveryMode{char}
}
