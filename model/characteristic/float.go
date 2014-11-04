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

func (f *Float) SetFloat(value float64) {
    f.SetNumber(value)
}

func (f *Float) SetMin(value float64) {
    f.SetMinValue(value)
}

func (f *Float) SetMax(value float64) {
    f.SetMaxValue(value)
}

func (f *Float) SetMinStep(value float64) {
    f.SetMinStepValue(value)
}

func (f *Float) FloatValue() float64 {
    return f.GetValue().(float64)
}

func (f *Float) Min() float64 {
    return f.GetMinValue().(float64)
}

func (f *Float) Max() float64 {
    return f.GetMaxValue().(float64)
}

func (f *Float) MinStep() float64 {
    return f.GetMinStepValue().(float64)
}