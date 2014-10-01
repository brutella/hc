package characteristic

type OnCharacteristic struct {
    *BoolCharacteristic
}

func NewOnCharacteristic(value bool) *OnCharacteristic {
    char := NewBoolCharacteristic(value)
    char.Type = CharTypeOn
    return &OnCharacteristic{char}
}

func (c *OnCharacteristic) SetOn(value bool) {
    c.SetBool(value)
}

func (c *OnCharacteristic) On() bool {
    return c.Bool()
}