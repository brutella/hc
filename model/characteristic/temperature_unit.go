package characteristic

type TemperatureUnit struct {
    *ByteCharacteristic
}

func NewTemperatureUnit(unit string) *TemperatureUnit {
    b := ByteFromUnit(unit)
    c := TemperatureUnit{NewByteCharacteristic(b)}
    c.Type = CharTypeTemperatureUnits
    c.Permissions = PermsAll()
    return &c
}

func (t *TemperatureUnit) Unit() byte {
    return t.Byte()
}