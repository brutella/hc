// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CurrentHeatingCoolingStateOff  int = 0
	CurrentHeatingCoolingStateHeat int = 1
	CurrentHeatingCoolingStateCool int = 2
)

const TypeCurrentHeatingCoolingState = "0000000F-0000-1000-8000-0026BB765291"

type CurrentHeatingCoolingState struct {
	*Int
}

func NewCurrentHeatingCoolingState() *CurrentHeatingCoolingState {
	char := NewInt(TypeCurrentHeatingCoolingState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &CurrentHeatingCoolingState{char}
}
