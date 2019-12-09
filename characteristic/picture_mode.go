// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	PictureModeOther          int = 0
	PictureModeStandard       int = 1
	PictureModeCalibrated     int = 2
	PictureModeCalibratedDark int = 3
	PictureModeVivid          int = 4
	PictureModeGame           int = 5
	PictureModeComputer       int = 6
	PictureModeCustom         int = 7
)

const TypePictureMode = "E2"

type PictureMode struct {
	*Int
}

func NewPictureMode() *PictureMode {
	char := NewInt(TypePictureMode)
	char.Format = FormatUInt16
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(13)
	char.SetStepValue(1)
	char.SetValue(0)

	return &PictureMode{char}
}
