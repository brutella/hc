package characteristic

type TemperatureCharacteristic struct {
    *Float
    value float64
}

func NewTemperatureTemperatureCharacteristic(value, min, max, steps float64, unit string, charType CharType, permissions []string) *TemperatureCharacteristic {
    c := TemperatureCharacteristic{NewFloatMinMaxSteps(value, min, max, steps), value}
    c.Unit = unit
    c.Type = charType
    c.Permissions = permissions
    return &c
}


func NewCurrentTemperatureCharacteristic(value, min, max, steps float64, unit string) *TemperatureCharacteristic {
    return NewTemperatureTemperatureCharacteristic(value, min, max, steps, unit, CharTypeTemperatureCurrent, PermsRead())
}

func NewTargetTemperatureCharacteristic(value, min, max, steps float64, unit string) *TemperatureCharacteristic {
    return NewTemperatureTemperatureCharacteristic(value, min, max, steps, unit, CharTypeTemperatureTarget, PermsAll())
}

func (c *TemperatureCharacteristic) SetTemperature(value float64) {
    c.SetFloat(value)
}

func (c *TemperatureCharacteristic) Temperature() float64 {
    return c.FloatValue()
}

func (c *TemperatureCharacteristic) MinTemperature() float64 {
    return c.Min()
}

func (c *TemperatureCharacteristic) MaxTemperature() float64 {
    return c.Max()
}

func (c *TemperatureCharacteristic) MinStepTemperature() float64 {
    return c.MinStep()
}