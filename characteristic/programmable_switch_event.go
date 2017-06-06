// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	ProgrammableSwitchEventSinglePress int = 0
	ProgrammableSwitchEventDoublePress int = 1
	ProgrammableSwitchEventLongPress   int = 2
)

const TypeProgrammableSwitchEvent = "73"

type ProgrammableSwitchEvent struct {
	*Int
}

func NewProgrammableSwitchEvent() *ProgrammableSwitchEvent {
	char := NewInt(TypeProgrammableSwitchEvent)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &ProgrammableSwitchEvent{char}
}
