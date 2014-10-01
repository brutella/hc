package characteristic

type LogsCharacteristic struct {
    *TLV8Characteristic
}

func NewLogsCharacteristic(logs string) *LogsCharacteristic {
    str := NewTLV8Characteristic([]byte(logs))
    str.Type = CharTypeLogs
    str.Permissions = PermsAll()
    
    return &LogsCharacteristic{str}
}

func (c *LogsCharacteristic) SetLogs(logs string) {
    // TODO
}

func (c *LogsCharacteristic) Logs() string {
    // TODO
    return ""
}