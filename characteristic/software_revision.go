// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSoftwareRevision = "00000054-0000-1000-8000-0026BB765291"

type SoftwareRevision struct {
	*String
}

func NewSoftwareRevision() *SoftwareRevision {
	char := NewString(TypeSoftwareRevision)
	char.Format = FormatString
	char.Perms = []string{PermRead}

	char.SetValue("")

	return &SoftwareRevision{char}
}
