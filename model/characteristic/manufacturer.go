package characteristic

type Manufacturer struct {
    *String
}

func NewManufacturer(name string) *Manufacturer {
    str := NewString(name)
    str.Type = CharTypeManufacturer
    str.Permissions = PermsReadOnly()
    
    return &Manufacturer{str}
}

func (c *Manufacturer) SetManufacturer(name string) {
    c.SetString(name)
}

func (c *Manufacturer) Manufacturer() string {
    return c.StringValue()
}