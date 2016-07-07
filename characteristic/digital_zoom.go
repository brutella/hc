// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeDigitalZoom = "11D"

type DigitalZoom struct {
	*Float
}

func NewDigitalZoom() *DigitalZoom {
	char := NewFloat(TypeDigitalZoom)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &DigitalZoom{char}
}
