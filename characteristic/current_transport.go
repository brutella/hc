package characteristic

const TypeCurrentTransport = "22B"

type CurrentTransport struct {
	*Bool
}

func NewCurrentTransport() *CurrentTransport {
	char := NewBool(TypeCurrentTransport)
	char.Format = FormatBool
	char.Perms = []string{PermRead}

	char.SetValue(false)

	return &CurrentTransport{char}
}
