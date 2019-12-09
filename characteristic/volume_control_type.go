// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	VolumeControlTypeNone                int = 0
	VolumeControlTypeRelative            int = 1
	VolumeControlTypeRelativeWithCurrent int = 2
	VolumeControlTypeAbsolute            int = 3
)

const TypeVolumeControlType = "E9"

type VolumeControlType struct {
	*Int
}

func NewVolumeControlType() *VolumeControlType {
	char := NewInt(TypeVolumeControlType)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(3)
	char.SetStepValue(1)
	char.SetValue(0)

	return &VolumeControlType{char}
}
