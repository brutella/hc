package characteristic

type Name struct {
	*String
}

func NewName(name string) *Name {
	str := NewString(name)
	str.Type = TypeName
	str.Permissions = PermsRead()

	return &Name{str}
}

func (n *Name) SetName(name string) {
	n.SetString(name)
}

func (n *Name) Name() string {
	return n.StringValue()
}
