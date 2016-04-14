// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TargetDoorStateOpen   int = 0
	TargetDoorStateClosed int = 1
)

const TypeTargetDoorState = "32"

type TargetDoorState struct {
	*Int
}

func NewTargetDoorState() *TargetDoorState {
	char := NewInt(TypeTargetDoorState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &TargetDoorState{char}
}
