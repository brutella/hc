package characteristic

type Name struct {
    *String
}

func NewName(name string) *Name {
    str := NewString(name)
    str.Type = CharTypeName
    str.Permissions = PermsRead()
    
    return &Name{str}
}

func (c *Name) SetName(name string) {
    c.SetString(name)
}

func (c *Name) Name() string{
    return c.StringValue()
}