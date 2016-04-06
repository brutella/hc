// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeStatusFault = "00000077-0000-1000-8000-0026BB765291"

type StatusFault struct {
	*Int
}

func NewStatusFault() *StatusFault {
	char := NewInt(TypeStatusFault)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &StatusFault{char}
}
