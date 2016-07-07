// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeNightVision = "11B"

type NightVision struct {
	*Bool
}

func NewNightVision() *NightVision {
	char := NewBool(TypeNightVision)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(false)

	return &NightVision{char}
}
