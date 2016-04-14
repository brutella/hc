// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePairingPairings = "50"

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
