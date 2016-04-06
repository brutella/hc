// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TargetDoorStateOpen   int = 0
	TargetDoorStateClosed int = 1
)

const TypeTargetDoorState = "00000032-0000-1000-8000-0026BB765291"

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
