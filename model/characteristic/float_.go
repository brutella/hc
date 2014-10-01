package characteristic

type FloatCharacteristic struct {
    *NumberCharacteristic
}

func NewFloatCharacteristic(value float64) *FloatCharacteristic {
    number := NewNumberCharacteristic(value, nil, nil, nil, FormatFloat)
    return &FloatCharacteristic{number}
}

func NewFloatCharacteristicMinMaxSteps(value, min, max, min_step float64) *FloatCharacteristic {
    number := NewNumberCharacteristic(value, min, max, min_step, FormatFloat)
    return &FloatCharacteristic{number}
}

func (c *FloatCharacteristic) SetFloat(value float64) {
    c.SetValue(value)
}

func (c *FloatCharacteristic) SetMin(value float64) {
    c.SetMinValue(value)
}

func (c *FloatCharacteristic) SetMax(value float64) {
    c.SetMaxValue(value)
}

func (c *FloatCharacteristic) SetMinStep(value float64) {
    c.SetMinStepValue(value)
}

func (c *FloatCharacteristic) Float() float64 {
    return c.GetValue().(float64)
}

func (c *FloatCharacteristic) Min() float64 {
    return c.GetMinValue().(float64)
}

func (c *FloatCharacteristic) Max() float64 {
    return c.GetMaxValue().(float64)
}

func (c *FloatCharacteristic) MinStep() float64 {
    return c.GetMinStepValue().(float64)
}