package characteristic

import(
    "fmt"
)

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

// TODO implement notifications
func (c *Characteristic) EnableEvents(enable bool) {
    c.Events = enable
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

// Compareable
func (c *Characteristic) Equal(other interface{}) bool {
    if characteristic, ok := other.(*Characteristic); ok == true {
        // The value type (e.g. float32, bool,...) of property `Value` may be different even though
        // they look the same. They are equal when they have the same string representation.
        value := fmt.Sprintf("%+v", c.Value)
        otherValue := fmt.Sprintf("%+v", characteristic.Value)
        
        return value == otherValue && c.Id == characteristic.Id && c.Type == characteristic.Type && len(c.Permissions) == len(characteristic.Permissions) && c.Description == characteristic.Description && c.Format == characteristic.Format && c.Unit == characteristic.Unit && c.MaxLen == characteristic.MaxLen && c.MaxValue == characteristic.MaxValue && c.MinValue == characteristic.MinValue && c.MinStep == characteristic.MinStep && c.Events == characteristic.Events && c.Bonjour == characteristic.Bonjour
    }
    
    return false
}

// model.Characteristic
func (c *Characteristic) SetId(id int) {
    c.Id = id
}

func (c *Characteristic) GetId() int {
    return c.Id
}

func (c *Characteristic) GetValue() interface{} {
    return c.Value
}

// Private

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