// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeProgrammableSwitchOutputState = "74"

type ProgrammableSwitchOutputState struct {
	*Int
}

func NewProgrammableSwitchOutputState() *ProgrammableSwitchOutputState {
	char := NewInt(TypeProgrammableSwitchOutputState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(1)
	char.SetStepValue(1)
	char.SetValue(0)

	return &ProgrammableSwitchOutputState{char}
}
