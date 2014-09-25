package model

type ModelCharacteristic struct {
    StringCharacteristic
}

func NewModelCharacteristic(model string) ModelCharacteristic {
    str := NewStringCharacteristic(model)
    str.Type = CharTypeModel
    str.Permissions = []string{PermRead}
    
    return ModelCharacteristic{str}
}