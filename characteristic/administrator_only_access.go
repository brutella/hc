// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeAdministratorOnlyAccess = "1"

type AdministratorOnlyAccess struct {
	*Bool
}

func NewAdministratorOnlyAccess() *AdministratorOnlyAccess {
	char := NewBool(TypeAdministratorOnlyAccess)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(false)

	return &AdministratorOnlyAccess{char}
}
