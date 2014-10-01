package characteristic

type Number struct {
    *Characteristic
}

func NewNumber(value, min, max, step interface{}, format string) *Number {    
    c := Number{NewCharacteristic(value, format, CharTypeUnknown, nil)}
    c.MinValue = min
    c.MaxValue = max
    c.MinStep = step
    
    return &c
}

func (c *Number) SetValue(value interface{}) {
    c.Value = value
}

func (c *Number) SetMinValue(value interface{}) {
    c.MinValue = value
}

func (c *Number) SetMaxValue(value interface{}) {
    c.MaxValue = value
}

func (c *Number) SetMinStepValue(value interface{}) {
    c.MinStep = value
}

func (c *Number) GetValue() interface{} {
    return c.Value
}

func (c *Number) GetMinValue() interface{} {
    return c.MinValue
}

func (c *Number) GetMaxValue() interface{} {
    return c.MaxValue
}

func (c *Number) GetMinStepValue() interface{} {
    return c.MinStep
}