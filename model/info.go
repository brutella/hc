package model

// Info contains basic information about an accessory
type Info struct {
	Name         string
	SerialNumber string
	Manufacturer string
	Model        string
	Firmware     string
	Software     string
	Hardware     string
	LogFile      string // TODO remove
}
