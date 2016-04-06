// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTunneledAccessoryConnected = "00000059-0000-1000-8000-0026BB765291"

type TunneledAccessoryConnected struct {
	*Bool
}

func NewTunneledAccessoryConnected() *TunneledAccessoryConnected {
	char := NewBool(TypeTunneledAccessoryConnected)
	char.Format = FormatBool
	char.Perms = []string{PermWrite, PermRead, PermEvents}

	char.SetValue(false)

	return &TunneledAccessoryConnected{char}
}
