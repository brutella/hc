package model

type BoolCharacteristic struct {
    *NumberCharacteristic
}

func NewBoolCharacteristic(value bool) *BoolCharacteristic {
    number := NewNumberCharacteristic(Int(value), nil, nil, nil, FormatBool, )
    return &BoolCharacteristic{number}
}

func (c *BoolCharacteristic) SetBool(value bool) {
    c.SetValue(Int(value))
}

func (c *BoolCharacteristic) Bool() bool {
    return Bool(c.GetValue())
}

func Bool(value interface{}) bool {
    if value, ok := value.(int); ok == true {
        if value == 0 {
            return false
        }
    }
    
    return true
}

func Int(value bool) int {
    switch value {
    case true:
        return 1
    }
    
    return 0
}