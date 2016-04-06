// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeAudioFeedback = "00000005-0000-1000-8000-0026BB765291"

type AudioFeedback struct {
	*Bool
}

func NewAudioFeedback() *AudioFeedback {
	char := NewBool(TypeAudioFeedback)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(false)

	return &AudioFeedback{char}
}
