// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSetupEndpoints = "118"

type SetupEndpoints struct {
	*Bytes
}

func NewSetupEndpoints() *SetupEndpoints {
	char := NewBytes(TypeSetupEndpoints)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue([]byte{})

	return &SetupEndpoints{char}
}
