// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	StatusTamperedNotTampered int = 0
	StatusTamperedTampered    int = 1
)

const TypeStatusTampered = "0000007A-0000-1000-8000-0026BB765291"

type StatusTampered struct {
	*Int
}

func NewStatusTampered() *StatusTampered {
	char := NewInt(TypeStatusTampered)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &StatusTampered{char}
}
