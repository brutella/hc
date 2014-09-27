package model

type SerialNumberCharacteristic struct {
    *StringCharacteristic
}

func NewSerialNumberCharacteristic(serial string) *SerialNumberCharacteristic {
    str := NewStringCharacteristic(serial)
    str.Type = CharTypeSerialNumber
    str.Permissions = PermsReadOnly()
    
    return &SerialNumberCharacteristic{str}
}