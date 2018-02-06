// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	ProgramModeNoProgramScheduled         int = 0
	ProgramModeProgramScheduled           int = 1
	ProgramModeProgramScheduledManualMode int = 2
)

const TypeProgramMode = "D1"

type ProgramMode struct {
	*Int
}

func NewProgramMode() *ProgramMode {
	char := NewInt(TypeProgramMode)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &ProgramMode{char}
}
