// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeResetFilterIndication = "AD"

type ResetFilterIndication struct {
	*Int
}

func NewResetFilterIndication() *ResetFilterIndication {
	char := NewInt(TypeResetFilterIndication)
	char.Format = FormatUInt8
	char.Perms = []string{PermWrite}
	char.SetMinValue(1)
	char.SetMaxValue(1)
	char.SetStepValue(1)

	return &ResetFilterIndication{char}
}
