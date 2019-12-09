// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TargetMediaStatePlay  int = 0
	TargetMediaStatePause int = 1
	TargetMediaStateStop  int = 2
)

const TypeTargetMediaState = "137"

type TargetMediaState struct {
	*Int
}

func NewTargetMediaState() *TargetMediaState {
	char := NewInt(TypeTargetMediaState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(2)
	char.SetStepValue(1)
	char.SetValue(0)

	return &TargetMediaState{char}
}
