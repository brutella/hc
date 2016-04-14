// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CurrentDoorStateOpen    int = 0
	CurrentDoorStateClosed  int = 1
	CurrentDoorStateOpening int = 2
	CurrentDoorStateClosing int = 3
	CurrentDoorStateStopped int = 4
)

const TypeCurrentDoorState = "E"

type CurrentDoorState struct {
	*Int
}

func NewCurrentDoorState() *CurrentDoorState {
	char := NewInt(TypeCurrentDoorState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &CurrentDoorState{char}
}
