package characteristic

type NumberCharacteristic struct {
    *Characteristic
}

func NewNumberCharacteristic(value, min, max, step interface{}, format string) *NumberCharacteristic {    
    c := NumberCharacteristic{NewCharacteristic(value, format, CharTypeUnknown, nil)}
    c.MinValue = min
    c.MaxValue = max
    c.MinStep = step
    
    return &c
}

func (c *NumberCharacteristic) SetValue(value interface{}) {
    c.Value = value
}

func (c *NumberCharacteristic) SetMinValue(value interface{}) {
    c.MinValue = value
}

func (c *NumberCharacteristic) SetMaxValue(value interface{}) {
    c.MaxValue = value
}

func (c *NumberCharacteristic) SetMinStepValue(value interface{}) {
    c.MinStep = value
}

func (c *NumberCharacteristic) GetValue() interface{} {
    return c.Value
}

func (c *NumberCharacteristic) GetMinValue() interface{} {
    return c.MinValue
}

func (c *NumberCharacteristic) GetMaxValue() interface{} {
    return c.MaxValue
}

func (c *NumberCharacteristic) GetMinStepValue() interface{} {
    return c.MinStep
}