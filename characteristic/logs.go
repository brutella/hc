// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeLogs = "1F"

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
