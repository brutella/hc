// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CurrentHumidifierDehumidifierStateInactive      int = 0
	CurrentHumidifierDehumidifierStateIdle          int = 1
	CurrentHumidifierDehumidifierStateHumidifying   int = 2
	CurrentHumidifierDehumidifierStateDehumidifying int = 3
)

const TypeCurrentHumidifierDehumidifierState = "B3"

type CurrentHumidifierDehumidifierState struct {
	*Int
}

func NewCurrentHumidifierDehumidifierState() *CurrentHumidifierDehumidifierState {
	char := NewInt(TypeCurrentHumidifierDehumidifierState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &CurrentHumidifierDehumidifierState{char}
}
