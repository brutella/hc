// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCurrentPositionStatus = "6D"

type CurrentPositionStatus struct {
	*Int
}

func NewCurrentPositionStatus() *CurrentPositionStatus {
	char := NewInt(TypeCurrentPositionStatus)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(0) // Изменил шаг с 1 на 100
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &CurrentPositionStatus{char}
}
