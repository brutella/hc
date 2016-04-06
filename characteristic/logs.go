// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeLogs = "0000001F-0000-1000-8000-0026BB765291"

type Logs struct {
	*Bytes
}

func NewLogs() *Logs {
	char := NewBytes(TypeLogs)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue([]byte{})

	return &Logs{char}
}
