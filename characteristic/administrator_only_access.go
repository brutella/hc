// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeAdministratorOnlyAccess = "00000001-0000-1000-8000-0026BB765291"

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
