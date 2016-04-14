// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePairSetup = "4C"

type PairSetup struct {
	*Bytes
}

func NewPairSetup() *PairSetup {
	char := NewBytes(TypePairSetup)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermWrite}

	char.SetValue([]byte{})

	return &PairSetup{char}
}
