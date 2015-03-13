package characteristic

import (
	"fmt"
	"github.com/brutella/hc/model"
	"github.com/gosexy/to"
	"reflect"
)

type ChangeFunc func(c *Characteristic, newValue, oldValue interface{})

// Characteristic is a HomeKit characteristic.
type Characteristic struct {
	Id          int64    `json:"iid"` // managed by accessory
	Type        CharType `json:"type"`
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

	remoteChangeFuncs []ChangeFunc
	localChangeFuncs  []ChangeFunc
}

func WriteOnlyPermissions(permissions []string) bool {
	if len(permissions) == 1 {
		return permissions[0] == PermWrite
	}
	return false
}

// NewCharacteristic returns a characteristic
// If no permissions are specified, the value of PermsAll() is used.
//
// If permissions are write-only the setter methods (SetValue and SetValueFromRemote)
// don't set the Value field. The OnLocalChange and OnRemoteChange have the new
// value set as expect, but characteristics current and old value are nil.
func NewCharacteristic(value interface{}, format string, t CharType, permissions []string) *Characteristic {
	if len(permissions) == 0 {
		permissions = PermsAll()
	}

	if WriteOnlyPermissions(permissions) == true {
		value = nil
	}

	return &Characteristic{
		Id:                model.InvalidId,
		Value:             value,
		Format:            format,
		Type:              t,
		Permissions:       permissions,
		remoteChangeFuncs: make([]ChangeFunc, 0),
		localChangeFuncs:  make([]ChangeFunc, 0),
	}
}

func (c *Characteristic) SetValue(value interface{}) {
	c.setValue(value, false)
}

func (c *Characteristic) SetValueFromRemote(value interface{}) {
	// Make sure that the new value is of same type the old value
	c.setValue(value, true)
}

func (c *Characteristic) SetEventsEnabled(enable bool) {
	c.Events = enable
}

func (c *Characteristic) EventsEnabled() bool {
	return c.Events
}

func (c *Characteristic) OnLocalChange(fn ChangeFunc) {
	c.localChangeFuncs = append(c.localChangeFuncs, fn)
}

func (c *Characteristic) OnRemoteChange(fn ChangeFunc) {
	c.remoteChangeFuncs = append(c.remoteChangeFuncs, fn)
}

// Compareable
func (c *Characteristic) Equal(other interface{}) bool {
	if characteristic, ok := other.(*Characteristic); ok == true {
		// The value type (e.g. float32, bool,...) of property `Value` may be different even though
		// they look the same. They are equal when they have the same string representation.
		value := fmt.Sprintf("%+v", c.Value)
		otherValue := fmt.Sprintf("%+v", characteristic.Value)

		return value == otherValue && c.Id == characteristic.Id && c.Type == characteristic.Type && len(c.Permissions) == len(characteristic.Permissions) && c.Description == characteristic.Description && c.Format == characteristic.Format && c.Unit == characteristic.Unit && c.MaxLen == characteristic.MaxLen && c.MaxValue == characteristic.MaxValue && c.MinValue == characteristic.MinValue && c.MinStep == characteristic.MinStep && c.Events == characteristic.Events
	}

	return false
}

// model.Characteristic
func (c *Characteristic) SetId(id int64) {
	c.Id = id
}

func (c *Characteristic) GetId() int64 {
	return c.Id
}

func (c *Characteristic) GetValue() interface{} {
	return c.Value
}

// Private

func (c *Characteristic) isWriteOnly() bool {
	return WriteOnlyPermissions(c.Permissions)
}

// Sets the value of the characteristic
// The implementation makes sure that the type of the value stays the same
// E.g. Type of characteristic value int, calling setValue("10.5") sets the value to int(10)
//
// When permissions are write only, this methods does not set the Value field.
func (c *Characteristic) setValue(value interface{}, remote bool) {
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

	old := c.Value
	if c.isWriteOnly() == false {
		c.Value = value
	} else {
		c.Value = nil
	}

	if remote == true {
		c.onChange(c.remoteChangeFuncs, value, old)
	} else {
		c.onChange(c.localChangeFuncs, value, old)
	}
}

func (c *Characteristic) onChange(funcs []ChangeFunc, newValue, oldValue interface{}) {
	for _, fn := range funcs {
		fn(c, newValue, oldValue)
	}
}
