package characteristic

type Float struct {
    *Number
}

func NewFloat(value float64) *Float {
    number := NewNumber(value, nil, nil, nil, FormatFloat)
    return &Float{number}
}

func NewFloatMinMaxSteps(value, min, max, min_step float64) *Float {
    number := NewNumber(value, min, max, min_step, FormatFloat)
    return &Float{number}
}

func (c *Float) SetFloat(value float64) {
    c.SetValue(value)
}

func (c *Float) SetMin(value float64) {
    c.SetMinValue(value)
}

func (c *Float) SetMax(value float64) {
    c.SetMaxValue(value)
}

func (c *Float) SetMinStep(value float64) {
    c.SetMinStepValue(value)
}

func (c *Float) FloatValue() float64 {
    return c.GetValue().(float64)
}

func (c *Float) Min() float64 {
    return c.GetMinValue().(float64)
}

func (c *Float) Max() float64 {
    return c.GetMaxValue().(float64)
}

func (c *Float) MinStep() float64 {
    return c.GetMinStepValue().(float64)
}