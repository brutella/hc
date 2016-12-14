// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TargetAirPurifierStateManual int = 0
	TargetAirPurifierStateAuto   int = 1
)

const TypeTargetAirPurifierState = "A8"

type TargetAirPurifierState struct {
	*Int
}

func NewTargetAirPurifierState() *TargetAirPurifierState {
	char := NewInt(TypeTargetAirPurifierState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &TargetAirPurifierState{char}
}
