package characteristic

type Brightness struct {
    *Int
}

func NewBrightness(value int) *Brightness {
    integer := NewInt(value, 0, 100, 1)
    integer.Unit = UnitPercentage
    integer.Type = CharTypeBrightness
    
    return &Brightness{integer}
}