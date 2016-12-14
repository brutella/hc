// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CurrentFanStateInactive   int = 0
	CurrentFanStateIdle       int = 1
	CurrentFanStateBlowingAir int = 2
)

const TypeCurrentFanState = "AF"

type CurrentFanState struct {
	*Int
}

func NewCurrentFanState() *CurrentFanState {
	char := NewInt(TypeCurrentFanState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &CurrentFanState{char}
}
