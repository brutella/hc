package characteristic

import (
	"fmt"
	"github.com/brutella/hc/model"
	"github.com/gosexy/to"
	"net"
	"reflect"
)

type ConnChangeFunc func(conn net.Conn, c *Characteristic, newValue, oldValue interface{})
type ChangeFunc func(c *Characteristic, newValue, oldValue interface{})

// Characteristic is a HomeKit characteristic.
type Characteristic struct {
	ID          int64    `json:"iid"` // managed by accessory
	Type        CharacteristicType `json:"type"`
	Permissions []string `json:"perms"`
	Description string   `json:"description,omitempty"` // manufacturer description (optional)

	Value  interface{} `json:"value,omitempty"` // nil for write-only characteristics
	Format string      `json:"format"`
	Unit   string      `json:"unit,omitempty"`

	MaxLen   int         `json:"maxLen,omitempty"`
	MaxValue interface{} `json:"maxValue,omitempty"`
	MinValue interface{} `json:"minValue,omitempty"`
	MinStep  interface{} `json:"minStep,omitempty"`

	// unused
	Events bool `json:"-"`

	connChangeFuncs  []ConnChangeFunc
	localChangeFuncs []ChangeFunc
}

// writeOnlyPermissions returns true when permissions only include write permission
func writeOnlyPermissions(permissions []string) bool {
	if len(permissions) == 1 {
		return permissions[0] == PermWrite
	}
	return false
}

// noWritePermissions returns true when permissions include no write permission
func noWritePermissions(permissions []string) bool {
	for _, value := range permissions {
		if value == PermWrite {
			return false
		}
	}
	return true
}

// NewCharacteristic returns a characteristic
// If no permissions are specified, the value of PermsAll() is used.
//
// If permissions are write-only the setter methods (SetValue and SetValueFromRemote)
// don't set the Value field. The OnLocalChange and OnRemoteChange have the new
// value set as expect, but characteristics current and old value are nil.
func NewCharacteristic(value interface{}, format string, t CharacteristicType, permissions []string) *Characteristic {
	if len(permissions) == 0 {
		permissions = PermsAll()
	}

	if writeOnlyPermissions(permissions) == true {
		value = nil
	}

	return &Characteristic{
		ID:               model.InvalidID,
		Value:            value,
		Format:           format,
		Type:             t,
		Permissions:      permissions,
		connChangeFuncs:  make([]ConnChangeFunc, 0),
		localChangeFuncs: make([]ChangeFunc, 0),
	}
}

func (c *Characteristic) SetValue(value interface{}) {
	c.setValue(value, nil)
}

func (c *Characteristic) SetValueFromConnection(value interface{}, conn net.Conn) {
	c.setValue(value, conn)
}

func (c *Characteristic) SetEventsEnabled(enable bool) {
	c.Events = enable
}

func (c *Characteristic) EventsEnabled() bool {
	return c.Events
}

func (c *Characteristic) OnChange(fn ChangeFunc) {
	c.localChangeFuncs = append(c.localChangeFuncs, fn)
}

func (c *Characteristic) OnConnChange(fn ConnChangeFunc) {
	c.connChangeFuncs = append(c.connChangeFuncs, fn)
}

// Equal returns true when receiver has the values as the argument.
func (c *Characteristic) Equal(other interface{}) bool {
	if characteristic, ok := other.(*Characteristic); ok == true {
		// The value type (e.g. float32, bool,...) of property `Value` may be different even though
		// they look the same. They are equal when they have the same string representation.
		value := fmt.Sprintf("%+v", c.Value)
		otherValue := fmt.Sprintf("%+v", characteristic.Value)

		return value == otherValue && c.ID == characteristic.ID && c.Type == characteristic.Type && len(c.Permissions) == len(characteristic.Permissions) && c.Description == characteristic.Description && c.Format == characteristic.Format && c.Unit == characteristic.Unit && c.MaxLen == characteristic.MaxLen && c.MaxValue == characteristic.MaxValue && c.MinValue == characteristic.MinValue && c.MinStep == characteristic.MinStep && c.Events == characteristic.Events
	}

	return false
}

// model.Characteristic
func (c *Characteristic) SetID(id int64) {
	c.ID = id
}

func (c *Characteristic) GetID() int64 {
	return c.ID
}

func (c *Characteristic) GetValue() interface{} {
	return c.Value
}

// Private

func (c *Characteristic) isWriteOnly() bool {
	return writeOnlyPermissions(c.Permissions)
}

func (c *Characteristic) hasWritePermissions() bool {
	return noWritePermissions(c.Permissions) == false
}

// Sets the value of the characteristic
// The implementation makes sure that the type of the value stays the same
// E.g. Type of characteristic value int, calling setValue("10.5") sets the value to int(10)
//
// When permissions are write only, this methods does not set the Value field.
func (c *Characteristic) setValue(value interface{}, conn net.Conn) {
	if c.Value != nil {
		converted, err := to.Convert(value, reflect.TypeOf(c.Value).Kind())
		if err == nil {
			value = converted
		}
	}

	// Ignore when new value is same
	if c.Value == value {
		return
	}

	// Ignore new values from remote when permissions don't allow write
	if c.hasWritePermissions() == false && conn != nil {
		return
	}

	old := c.Value
	if c.isWriteOnly() == false {
		c.Value = value
	} else {
		c.Value = nil
	}

	if conn != nil {
		c.onConnChange(c.connChangeFuncs, conn, value, old)
	} else {
		c.onChange(c.localChangeFuncs, value, old)
	}
}

func (c *Characteristic) onChange(funcs []ChangeFunc, newValue, oldValue interface{}) {
	for _, fn := range funcs {
		fn(c, newValue, oldValue)
	}
}

func (c *Characteristic) onConnChange(funcs []ConnChangeFunc, conn net.Conn, newValue, oldValue interface{}) {
	for _, fn := range funcs {
		fn(conn, c, newValue, oldValue)
	}
}
