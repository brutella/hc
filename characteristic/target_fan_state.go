// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TargetFanStateManual int = 0
	TargetFanStateAuto   int = 1
)

const TypeTargetFanState = "BF"

type TargetFanState struct {
	*Int
}

func NewTargetFanState() *TargetFanState {
	char := NewInt(TypeTargetFanState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &TargetFanState{char}
}
