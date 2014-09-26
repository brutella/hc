package model

type ModelCharacteristic struct {
    *StringCharacteristic
}

func NewModelCharacteristic(model string) *ModelCharacteristic {
    str := NewStringCharacteristic(model)
    str.Type = CharTypeModel
    str.Permissions = PermsReadOnly()
    
    return &ModelCharacteristic{str}
}