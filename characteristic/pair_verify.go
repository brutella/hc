// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePairVerify = "0000004E-0000-1000-8000-0026BB765291"

type PairVerify struct {
	*Bytes
}

func NewPairVerify() *PairVerify {
	char := NewBytes(TypePairVerify)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermWrite}

	char.SetValue([]byte{})

	return &PairVerify{char}
}
