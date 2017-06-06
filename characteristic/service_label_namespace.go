// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	ServiceLabelNamespaceDots           int = 0
	ServiceLabelNamespaceArabicNumerals int = 1
)

const TypeServiceLabelNamespace = "CD"

type ServiceLabelNamespace struct {
	*Int
}

func NewServiceLabelNamespace() *ServiceLabelNamespace {
	char := NewInt(TypeServiceLabelNamespace)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead}

	char.SetValue(0)

	return &ServiceLabelNamespace{char}
}
