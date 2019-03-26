// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CurrentVisibilityStateShown  int = 0
	CurrentVisibilityStateHidden int = 1
)

const TypeCurrentVisibilityState = "135"

type CurrentVisibilityState struct {
	*Int
}

func NewCurrentVisibilityState() *CurrentVisibilityState {
	char := NewInt(TypeCurrentVisibilityState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(3)
	char.SetStepValue(1)
	char.SetValue(0)

	return &CurrentVisibilityState{char}
}
