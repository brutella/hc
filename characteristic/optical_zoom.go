// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeOpticalZoom = "11C"

type OpticalZoom struct {
	*Float
}

func NewOpticalZoom() *OpticalZoom {
	char := NewFloat(TypeOpticalZoom)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &OpticalZoom{char}
}
