package characteristic

type NameCharacteristic struct {
    *StringCharacteristic
}

func NewNameCharacteristic(name string) *NameCharacteristic {
    str := NewStringCharacteristic(name)
    str.Type = CharTypeName
    str.Permissions = PermsRead()
    
    return &NameCharacteristic{str}
}

func (c *NameCharacteristic) SetName(name string) {
    c.SetString(name)
}

func (c *NameCharacteristic) Name() string{
    return c.String()
}