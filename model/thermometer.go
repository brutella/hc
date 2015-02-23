package model

type TempUnit string

const (
	TempUnitCelsius = TempUnit("celsius")
)

// A thermometer measures a temperature and let you set a target temperature
type Thermometer interface {
	Accessory

	SetTemperature(float64)
	Temperature() float64
	Unit() TempUnit
}
