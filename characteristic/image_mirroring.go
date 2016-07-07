// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeImageMirroring = "11F"

type ImageMirroring struct {
	*Bool
}

func NewImageMirroring() *ImageMirroring {
	char := NewBool(TypeImageMirroring)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(false)

	return &ImageMirroring{char}
}
