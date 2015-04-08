package characteristic

import (
	"github.com/brutella/hc/model"
)

type TemperatureUnit struct {
	*ByteCharacteristic
}

func NewTemperatureUnit(unit model.TempUnit) *TemperatureUnit {
	b := ByteFromTempUnit(unit)
	c := TemperatureUnit{NewByteCharacteristic(b, PermsAll())}
	c.Type = CharTypeTemperatureUnits
	return &c
}

func (t *TemperatureUnit) Unit() model.TempUnit {
	return TempUnitFromByte(t.Byte())
}
