package characteristic

import(
    "github.com/brutella/hap/model"
)

type TemperatureUnit struct {
    *ByteCharacteristic
}

func NewTemperatureUnit(unit model.TempUnit) *TemperatureUnit {
    b := ByteFromUnit(unit)
    c := TemperatureUnit{NewByteCharacteristic(b)}
    c.Type = CharTypeTemperatureUnits
    c.Permissions = PermsAll()
    return &c
}

func (t *TemperatureUnit) Unit() model.TempUnit {
    return TempUnitFromByte(t.Byte())
}