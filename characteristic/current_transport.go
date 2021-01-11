package characteristic

// CurrentTransport
// this can't be right -- '0000021E-0000-1000-8000-0000022B'
// const TypeWifiCurrentTransport = "21E"
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
