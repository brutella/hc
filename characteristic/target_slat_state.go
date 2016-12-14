// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TargetSlatStateManual int = 0
	TargetSlatStateAuto   int = 1
)

const TypeTargetSlatState = "BE"

type TargetSlatState struct {
	*Int
}

func NewTargetSlatState() *TargetSlatState {
	char := NewInt(TypeTargetSlatState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &TargetSlatState{char}
}
