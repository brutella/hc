// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTunnelConnectionTimeout = "61"

type TunnelConnectionTimeout struct {
	*Int
}

func NewTunnelConnectionTimeout() *TunnelConnectionTimeout {
	char := NewInt(TypeTunnelConnectionTimeout)
	char.Format = FormatUInt32
	char.Perms = []string{PermWrite, PermRead, PermEvents}

	char.SetValue(0)

	return &TunnelConnectionTimeout{char}
}
