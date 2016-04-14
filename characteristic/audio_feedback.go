// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeAudioFeedback = "5"

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
