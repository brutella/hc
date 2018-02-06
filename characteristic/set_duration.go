// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSetDuration = "D3"

type SetDuration struct {
	*Int
}

func NewSetDuration() *SetDuration {
	char := NewInt(TypeSetDuration)
	char.Format = FormatUInt32
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(3600)
	char.SetStepValue(1)
	char.SetValue(0)

	return &SetDuration{char}
}
