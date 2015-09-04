package characteristic

type Model struct {
	*String
}

func NewModel(model string) *Model {
	str := NewString(model)
	str.Type = TypeModel
	str.Permissions = PermsReadOnly()

	return &Model{str}
}

func (m *Model) SetModel(model string) {
	m.SetString(model)
}

func (m *Model) Model() string {
	return m.StringValue()
}
