// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CurrentAirPurifierStateInactive     int = 0
	CurrentAirPurifierStateIdle         int = 1
	CurrentAirPurifierStatePurifyingAir int = 2
)

const TypeCurrentAirPurifierState = "A9"

type CurrentAirPurifierState struct {
	*Int
}

func NewCurrentAirPurifierState() *CurrentAirPurifierState {
	char := NewInt(TypeCurrentAirPurifierState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &CurrentAirPurifierState{char}
}
