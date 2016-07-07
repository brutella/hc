// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	StatusFaultNoFault      int = 0
	StatusFaultGeneralFault int = 1
)

const TypeStatusFault = "77"

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
