// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeServiceLabelIndex = "CB"

type ServiceLabelIndex struct {
	*Int
}

func NewServiceLabelIndex() *ServiceLabelIndex {
	char := NewInt(TypeServiceLabelIndex)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead}
	char.SetMinValue(1)
	char.SetMaxValue(255)
	char.SetStepValue(1)
	char.SetValue(1)

	return &ServiceLabelIndex{char}
}
