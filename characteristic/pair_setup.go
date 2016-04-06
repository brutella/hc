// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePairSetup = "0000004C-0000-1000-8000-0026BB765291"

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
