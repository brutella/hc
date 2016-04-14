// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePairingFeatures = "4F"

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
