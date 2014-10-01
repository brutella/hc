package characteristic

type Model struct {
    *String
}

func NewModel(model string) *Model {
    str := NewString(model)
    str.Type = CharTypeModel
    str.Permissions = PermsReadOnly()
    
    return &Model{str}
}

func (c *Model) SetModel(model string) {
    c.SetString(model)
}

func (c *Model) Model() string {
    return c.StringValue()
}