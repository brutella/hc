// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TargetVisibilityStateShown  int = 0
	TargetVisibilityStateHidden int = 1
)

const TypeTargetVisibilityState = "134"

type TargetVisibilityState struct {
	*Int
}

func NewTargetVisibilityState() *TargetVisibilityState {
	char := NewInt(TypeTargetVisibilityState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(2)
	char.SetStepValue(1)
	char.SetValue(0)

	return &TargetVisibilityState{char}
}
