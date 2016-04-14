// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTunneledAccessoryAdvertising = "60"

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
