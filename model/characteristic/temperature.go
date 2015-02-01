package characteristic

type TemperatureCharacteristic struct {
	*Float
	value float64
}

func NewTemperatureTemperatureCharacteristic(value, min, max, steps float64, unit string, charType CharType, permissions []string) *TemperatureCharacteristic {
	t := TemperatureCharacteristic{NewFloatMinMaxSteps(value, min, max, steps), value}
	t.Unit = unit
	t.Type = charType
	t.Permissions = permissions
	return &t
}

func NewCurrentTemperatureCharacteristic(value, min, max, steps float64, unit string) *TemperatureCharacteristic {
	return NewTemperatureTemperatureCharacteristic(value, min, max, steps, unit, CharTypeTemperatureCurrent, PermsRead())
}

func NewTargetTemperatureCharacteristic(value, min, max, steps float64, unit string) *TemperatureCharacteristic {
	return NewTemperatureTemperatureCharacteristic(value, min, max, steps, unit, CharTypeTemperatureTarget, PermsAll())
}

func (t *TemperatureCharacteristic) SetTemperature(value float64) {
	t.SetFloat(value)
}

func (t *TemperatureCharacteristic) Temperature() float64 {
	return t.FloatValue()
}

func (t *TemperatureCharacteristic) MinTemperature() float64 {
	return t.Min()
}

func (t *TemperatureCharacteristic) MaxTemperature() float64 {
	return t.Max()
}

func (t *TemperatureCharacteristic) MinStepTemperature() float64 {
	return t.MinStep()
}
