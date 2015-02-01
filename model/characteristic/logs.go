package characteristic

type Logs struct {
	*TLV8
}

func NewLogs(logs string) *Logs {
	str := NewTLV8([]byte(logs))
	str.Type = CharTypeLogs
	str.Permissions = PermsAll()

	return &Logs{str}
}

func (c *Logs) SetLogs(logs string) {
	// TODO
}

func (c *Logs) Logs() string {
	// TODO
	return ""
}
