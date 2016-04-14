// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	PositionStateDecreasing int = 0
	PositionStateIncreasing int = 1
	PositionStateStopped    int = 2
)

const TypePositionState = "72"

type PositionState struct {
	*Int
}

func NewPositionState() *PositionState {
	char := NewInt(TypePositionState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &PositionState{char}
}
