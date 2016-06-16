// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeMute = "11A"

type Mute struct {
	*Bool
}

func NewMute() *Mute {
	char := NewBool(TypeMute)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(false)

	return &Mute{char}
}
