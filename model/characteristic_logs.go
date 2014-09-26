package model

type LogsCharacteristic struct {
    *TLV8Characteristic
}

func NewLogsCharacteristic(logs string) *LogsCharacteristic {
    str := NewTLV8Characteristic([]byte(logs))
    str.Type = CharTypeLogs
    str.Permissions = PermsAll()
    
    return &LogsCharacteristic{str}
}