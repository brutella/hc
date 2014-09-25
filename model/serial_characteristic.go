package model

type SerialNumberCharacteristic struct {
    StringCharacteristic
}

func NewSerialNumberCharacteristic(serial string) SerialNumberCharacteristic {
    str := NewStringCharacteristic(serial)
    str.Type = CharTypeSerialNumber
    str.Permissions = []string{PermRead}
    
    return SerialNumberCharacteristic{str}
}