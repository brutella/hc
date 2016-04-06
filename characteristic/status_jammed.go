// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	StatusJammedNotJammed int = 0
	StatusJammedJammed    int = 1
)

const TypeStatusJammed = "00000078-0000-1000-8000-0026BB765291"

type StatusJammed struct {
	*Int
}

func NewStatusJammed() *StatusJammed {
	char := NewInt(TypeStatusJammed)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &StatusJammed{char}
}
