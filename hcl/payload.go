package hcl

import (
	"time"
)

const (
	TransitionTypeBrightness       byte = 1 // not sure
	TransitionTypeColorTemperature byte = 2
)

type SupportedTransitionConfiguration struct {
	Configurations []ValueTransitionConfiguration `tlv8:"1"`
}

type ValueTransitionConfiguration struct {
	CharacteristicID uint8 `tlv8:"1"`
	TransitionType   uint8 `tlv8:"2"`
}

type TransitionControl struct {
	Read   TransitionRead   `tlv8:"1,optional"`
	Update TransitionUpdate `tlv8:"2,optional"`
}

type TransitionRead struct {
	ColorTemperatureIID uint8 `tlv8:"1"`
}

type TransitionReadResponse struct {
	Configuration TransitionConfiguration `tlv8:"1"`
}

type TransitionUpdate struct {
	Configuration TransitionConfiguration `tlv8:"1"`
}

type TransitionUpdateResponse struct {
	Error byte `tlv8:"2"`
}

type TransitionConfiguration struct {
	ColorTemperatureIID uint8            `tlv8:"1"`
	Params              TransitionParams `tlv8:"2"`
	Enabled             bool             `tlv8:"3,optional"` // not sure
	Unknown4            uint8            `tlv8:"4,optional"`
	Curve               TransitionCurve  `tlv8:"5"`
	UpdateInterval      uint16           `tlv8:"6"`
	Unknown7            byte             `tlv8:"7,optional"`
	NotifyInterval      uint32           `tlv8:"8"`
}

type TransitionControlResponse struct {
	Status TransitionControlStatus `tlv8:"1"`
}

type TransitionControlStatus struct {
	ColorTemperatureIID uint8            `tlv8:"1"`
	Params              TransitionParams `tlv8:"2"`
	SinceStart          uint64           `tlv8:"3"`
}

func (c TransitionConfiguration) StartDate() time.Time {
	return Date(c.Params.StartTime)
}

type TransitionParams struct {
	TransitionID []byte `tlv8:"1"`
	StartTime    uint64 `tlv8:"2"` // Milliseconds since 2001/01/01 00:00:00
	Unknown3     []byte `tlv8:"3,optional"`
}

type TransitionCurve struct {
	Entries       []TransitionCurveEntry `tlv8:"1"`
	BrightnessIID uint8                  `tlv8:"2"`
	ValueRange    TransitionValueRange   `tlv8:"3"`
}

type TransitionCurveEntry struct {
	BrightnessAdjustment float32 `tlv8:"1"`
	ColorTemperature     float32 `tlv8:"2"`
	TimeOffset           uint32  `tlv8:"3"`          // type ?
	Duration             uint32  `tlv8:"4,optional"` // type ?
}

type TransitionValueRange struct {
	Min uint32 `tlv8:"1"`
	Max uint32 `tlv8:"2"`
}

var RefDate = time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)

func Date(timestamp uint64) time.Time {
	return RefDate.Add(time.Duration(timestamp) * time.Millisecond)
}

func Timestamp(t time.Time) uint64 {
	return uint64(t.Sub(RefDate).Milliseconds())
}
