// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeDayOfTheWeek = "98"

type DayOfTheWeek struct {
	*Int
}

func NewDayOfTheWeek() *DayOfTheWeek {
	char := NewInt(TypeDayOfTheWeek)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite}
	char.SetMinValue(1)
	char.SetMaxValue(7)
	char.SetStepValue(1)
	char.SetValue(1)

	return &DayOfTheWeek{char}
}
