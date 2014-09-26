package model

type IdentifyCharacteristic struct {
    *BoolCharacteristic
}

func NewIdentifyCharacteristic(identify bool) *IdentifyCharacteristic {
    b := NewBoolCharacteristic(identify)
    b.Type = CharTypeIdentify
    b.Permissions = PermsAll()
    
    return &IdentifyCharacteristic{b}
}