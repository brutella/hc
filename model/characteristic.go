package hk

type CharacteristicChange struct {
    OldValue interface{}
    NewValue interface{}
}

type CharacteristicDelegate interface {
    CharactericDidChangeValue(c *Characteristic, change CharacteristicChange)
}

type ValueChangedFunc func(CharacteristicChange)
type Characteristic struct {
    Id int                  `json:"iid"` // managed by accessory
    Type CharType           `json:"type"`
    Permissions []string    `json:"perms"`
    Description string      `json:"description,omitempty"`// manufacturer description (optional) 
    
    Value interface{}   `json:"value"` // any
    Format string       `json:"format"`
    Unit string         `json:"unit,omitempty"`
    
    MaxLen int              `json:"maxLen,omitempty"`
    MaxValue interface{}    `json:"maxValue,omitempty"`
    MinValue interface{}    `json:"minValue,omitempty"`
    MinStep interface{}     `json:"minStep,omitempty"`
    
    // unused
    Events bool     `json:",omitempty"`
    Bonjour bool    `json:",omitempty"`
    
    remoteDelegates []CharacteristicDelegate
    localDelegates []CharacteristicDelegate
}

// Creates a new characteristic
// If no permissions are specified, read and write will be added
func NewCharacteristic(value interface{}, format string, t CharType,  permissions []string) *Characteristic {
    if len(permissions) == 0 {
        permissions = PermsAll()
    }
    
    return &Characteristic{
        Value: value,
        Format: format,
        Type: t,
        Permissions: permissions,
        remoteDelegates: make([]CharacteristicDelegate, 0),
        localDelegates: make([]CharacteristicDelegate, 0),
    }
}

func (c *Characteristic) SetValue(value interface{}) {
    c.setValue(value, false)
}

func (c *Characteristic) SetValueFromRemote(value interface{}) {
    c.setValue(value, true)
}

func (c *Characteristic) RemoveDelegate(delegate CharacteristicDelegate) {
    for i, d := range c.localDelegates {
        if d == delegate {
            c.localDelegates = append(c.localDelegates[:i], c.localDelegates[i+1:]...)
            break
        }
    }
    
    for i, d := range c.remoteDelegates {
        if d == delegate {
            c.remoteDelegates = append(c.remoteDelegates[:i], c.remoteDelegates[i+1:]...)
            break
        }
    }
}

func (c *Characteristic) AddLocalChangeDelegate(d CharacteristicDelegate) {
    c.localDelegates = append(c.localDelegates, d)
}

func (c *Characteristic) AddRemoteChangeDelegate(d CharacteristicDelegate) {
    c.remoteDelegates = append(c.remoteDelegates, d)
}


func (c *Characteristic) setValue(value interface{}, remote bool) {
    old := c.Value
    c.Value = value

    change := CharacteristicChange{
            OldValue:old,
            NewValue:c.Value,
    }
    
    if remote == true {
        c.callValueChangeOnDelegates(change, c.remoteDelegates)
    } else {
        c.callValueChangeOnDelegates(change, c.localDelegates)
    }
}

func (c *Characteristic) callValueChangeOnDelegates(change CharacteristicChange, delegates []CharacteristicDelegate) {
    for _, d := range delegates {
        d.CharactericDidChangeValue(c, change)
    }
}