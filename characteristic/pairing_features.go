// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePairingFeatures = "0000004F-0000-1000-8000-0026BB765291"

type PairingFeatures struct {
	*Int
}

func NewPairingFeatures() *PairingFeatures {
	char := NewInt(TypePairingFeatures)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead}

	char.SetValue(0)

	return &PairingFeatures{char}
}
