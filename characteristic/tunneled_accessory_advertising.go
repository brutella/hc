// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTunneledAccessoryAdvertising = "00000060-0000-1000-8000-0026BB765291"

type TunneledAccessoryAdvertising struct {
	*Bool
}

func NewTunneledAccessoryAdvertising() *TunneledAccessoryAdvertising {
	char := NewBool(TypeTunneledAccessoryAdvertising)
	char.Format = FormatBool
	char.Perms = []string{PermWrite, PermRead, PermEvents}

	char.SetValue(false)

	return &TunneledAccessoryAdvertising{char}
}
