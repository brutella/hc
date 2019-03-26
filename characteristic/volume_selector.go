// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	VolumeSelectorIncrement int = 0
	VolumeSelectorDecrement int = 1
)

const TypeVolumeSelector = "EA"

type VolumeSelector struct {
	*Int
}

func NewVolumeSelector() *VolumeSelector {
	char := NewInt(TypeVolumeSelector)
	char.Format = FormatUInt8
	char.Perms = []string{PermWrite}
	char.SetMinValue(0)
	char.SetMaxValue(1)
	char.SetStepValue(1)

	return &VolumeSelector{char}
}
