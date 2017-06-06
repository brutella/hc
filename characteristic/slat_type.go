// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	SlatTypeHorizontal int = 0
	SlatTypeVertical   int = 1
)

const TypeSlatType = "C0"

type SlatType struct {
	*Int
}

func NewSlatType() *SlatType {
	char := NewInt(TypeSlatType)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead}

	char.SetValue(0)

	return &SlatType{char}
}
