package model

type TemperatureUnitCharacteristic struct {
    *ByteCharacteristic
}

func NewTemperatureUnitCharacteristic(unit string) *TemperatureUnitCharacteristic {
    b := ByteFromUnit(unit)
    c := TemperatureUnitCharacteristic{NewByteCharacteristic(b)}
    c.Type = CharTypeTemperatureUnits
    c.Permissions = PermsAll()
    return &c
}

func (c *TemperatureUnitCharacteristic) Unit() byte {
    return c.Byte()
}