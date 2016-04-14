// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTunneledAccessoryStateNumber = "58"

type TunneledAccessoryStateNumber struct {
	*Float
}

func NewTunneledAccessoryStateNumber() *TunneledAccessoryStateNumber {
	char := NewFloat(TypeTunneledAccessoryStateNumber)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &TunneledAccessoryStateNumber{char}
}
