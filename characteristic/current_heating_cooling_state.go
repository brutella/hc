// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CurrentHeatingCoolingStateOff  int = 0
	CurrentHeatingCoolingStateHeat int = 1
	CurrentHeatingCoolingStateCool int = 2
)

const TypeCurrentHeatingCoolingState = "F"

type CurrentHeatingCoolingState struct {
	*UInt8
}

func NewCurrentHeatingCoolingState() *CurrentHeatingCoolingState {
	char := NewUInt8(TypeCurrentHeatingCoolingState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &CurrentHeatingCoolingState{char}
}
