package hk

type ModelCharacteristic struct {
    *StringCharacteristic
}

func NewModelCharacteristic(model string) *ModelCharacteristic {
    str := NewStringCharacteristic(model)
    str.Type = CharTypeModel
    str.Permissions = PermsReadOnly()
    
    return &ModelCharacteristic{str}
}

func (c *ModelCharacteristic) SetModel(model string) {
    c.SetString(model)
}

func (c *ModelCharacteristic) Model() string {
    return c.String()
}