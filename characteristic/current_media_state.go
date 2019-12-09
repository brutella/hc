// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CurrentMediaStatePlay    int = 0
	CurrentMediaStatePause   int = 1
	CurrentMediaStateStop    int = 2
	CurrentMediaStateUnknown int = 3
)

const TypeCurrentMediaState = "E0"

type CurrentMediaState struct {
	*Int
}

func NewCurrentMediaState() *CurrentMediaState {
	char := NewInt(TypeCurrentMediaState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(3)
	char.SetStepValue(1)
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &CurrentMediaState{char}
}
