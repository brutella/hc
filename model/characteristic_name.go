package model

type NameCharacteristic struct {
    *StringCharacteristic
}

func NewNameCharacteristic(name string) *NameCharacteristic {
    str := NewStringCharacteristic(name)
    str.Type = CharTypeName
    str.Permissions = PermsRead()
    
    return &NameCharacteristic{str}
}