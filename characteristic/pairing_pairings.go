// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePairingPairings = "00000050-0000-1000-8000-0026BB765291"

type PairingPairings struct {
	*Bytes
}

func NewPairingPairings() *PairingPairings {
	char := NewBytes(TypePairingPairings)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermWrite}

	char.SetValue([]byte{})

	return &PairingPairings{char}
}
