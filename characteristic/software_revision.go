// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSoftwareRevision = "54"

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
