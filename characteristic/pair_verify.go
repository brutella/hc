// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePairVerify = "4E"

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
