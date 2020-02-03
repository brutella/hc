package characteristic

import (
	"fmt"
	"net"

	"github.com/xiam/to"
)

type ConnChangeFunc func(conn net.Conn, c *Characteristic, newValue, oldValue interface{})
type ChangeFunc func(c *Characteristic, newValue, oldValue interface{})
type GetFunc func() interface{}

// Characteristic is a HomeKit characteristic.
type Characteristic struct {
	ID          uint64   `json:"iid"` // managed by accessory
	Type        string   `json:"type"`
	Perms       []string `json:"perms"`
	Description string   `json:"description,omitempty"` // manufacturer description (optional)

	Value  interface{} `json:"value,omitempty"` // nil for write-only characteristics
	Format string      `json:"format"`
	Unit   string      `json:"unit,omitempty"`

	MaxLen    int         `json:"maxLen,omitempty"`
	MaxValue  interface{} `json:"maxValue,omitempty"`
	MinValue  interface{} `json:"minValue,omitempty"`
	StepValue interface{} `json:"minStep,omitempty"`

	// unused
	Events bool `json:"-"`

	connValueUpdateFuncs []ConnChangeFunc
	valueChangeFuncs     []ChangeFunc
	valueGetFunc         GetFunc
}

// readPerm returns true when perms only include read permission
func readPerm(perms []string) bool {
	for _, perm := range perms {
		if perm == PermRead {
			return true
		}
	}

	return false
}

// noWritePerms returns true when permissions include no write permission
func noWritePerms(permissions []string) bool {
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
// If permissions are write-only the setter methods (UpdateValue and UpdateValueFromRemote)
// don't set the Value field. The OnLocalChange and OnRemoteChange have the new
// value set as expect, but characteristics current and old value are nil.
func NewCharacteristic(typ string) *Characteristic {
	return &Characteristic{
		Type:                 typ,
		connValueUpdateFuncs: make([]ConnChangeFunc, 0),
		valueChangeFuncs:     make([]ChangeFunc, 0),
	}
}

func (c *Characteristic) GetValue() interface{} {
	return c.getValue(nil)
}

func (c *Characteristic) GetValueFromConnection(conn net.Conn) interface{} {
	return c.getValue(conn)
}

func (c *Characteristic) OnValueGet(fn GetFunc) {
	c.valueGetFunc = fn
}

func (c *Characteristic) UpdateValue(value interface{}) {
	c.updateValue(value, nil, false)
}

func (c *Characteristic) UpdateValueFromConnection(value interface{}, conn net.Conn) {
	c.updateValue(value, conn, true)
}

func (c *Characteristic) OnValueUpdate(fn ChangeFunc) {
	c.valueChangeFuncs = append(c.valueChangeFuncs, fn)
}

func (c *Characteristic) OnValueUpdateFromConn(fn ConnChangeFunc) {
	c.connValueUpdateFuncs = append(c.connValueUpdateFuncs, fn)
}

// Equal returns true when receiver has the values as the argument.
func (c *Characteristic) Equal(other interface{}) bool {
	if characteristic, ok := other.(*Characteristic); ok == true {
		// The value type (e.g. float32, bool,...) of property `Value` may be different even though
		// they look the same. They are equal when they have the same string representation.
		value := fmt.Sprintf("%+v", c.Value)
		otherValue := fmt.Sprintf("%+v", characteristic.Value)

		return value == otherValue && c.ID == characteristic.ID && c.Type == characteristic.Type && len(c.Perms) == len(characteristic.Perms) && c.Description == characteristic.Description && c.Format == characteristic.Format && c.Unit == characteristic.Unit && c.MaxLen == characteristic.MaxLen && c.MaxValue == characteristic.MaxValue && c.MinValue == characteristic.MinValue && c.StepValue == characteristic.StepValue && c.Events == characteristic.Events
	}

	return false
}

// Private

func (c *Characteristic) isReadable() bool {
	return readPerm(c.Perms)
}

func (c *Characteristic) hasWritePerms() bool {
	return noWritePerms(c.Perms) == false
}

func (c *Characteristic) getValue(conn net.Conn) interface{} {
	if c.valueGetFunc != nil {
		c.updateValue(c.valueGetFunc(), conn, false)
	}
	return c.Value
}

// Sets the value of the characteristic
// The implementation makes sure that the type of the value stays the same
// E.g. Type of characteristic value int, calling updateValue("10.5") sets the value to int(10)
//
// When permissions are write only and checkPerms is true, this methods does not set the Value field.
func (c *Characteristic) updateValue(value interface{}, conn net.Conn, checkPerms bool) {
	value = c.convert(value)

	// Value must be within min and max
	switch c.Format {
	case FormatFloat:
		value = c.boundFloat64Value(value.(float64))
	case FormatUInt8, FormatUInt16, FormatUInt32, FormatUInt64, FormatInt32:
		value = c.boundIntValue(value.(int))
	}

	// Ignore when new value is same
	if c.Value == value {
		return
	}

	// Ignore new values from remote when permissions don't allow write and checkPerms is true
	if checkPerms == true && c.hasWritePerms() == false {
		return
	}

	old := c.Value
	if c.isReadable() {
		c.Value = value
	}

	if conn != nil {
		c.onValueUpdateFromConn(c.connValueUpdateFuncs, conn, value, old)
	} else {
		c.onValueUpdate(c.valueChangeFuncs, value, old)
	}
}

func (c *Characteristic) onValueUpdate(funcs []ChangeFunc, newValue, oldValue interface{}) {
	for _, fn := range funcs {
		fn(c, newValue, oldValue)
	}
}

func (c *Characteristic) onValueUpdateFromConn(funcs []ConnChangeFunc, conn net.Conn, newValue, oldValue interface{}) {
	for _, fn := range funcs {
		fn(conn, c, newValue, oldValue)
	}
}

func (c *Characteristic) boundFloat64Value(value float64) interface{} {
	min, minOK := c.MinValue.(float64)
	max, maxOK := c.MaxValue.(float64)
	if maxOK == true && value > max {
		value = max
	} else if minOK == true && value < min {
		value = min
	}

	return value
}

func (c *Characteristic) boundIntValue(value int) interface{} {
	min, minOK := c.MinValue.(int)
	max, maxOK := c.MaxValue.(int)
	if maxOK == true && value > max {
		value = max
	} else if minOK == true && value < min {
		value = min
	}

	return value
}

func (c *Characteristic) convert(v interface{}) interface{} {
	switch c.Format {
	case FormatFloat:
		return to.Float64(v)
	case FormatUInt8:
		return int(to.Uint64(v))
	case FormatUInt16:
		return int(to.Uint64(v))
	case FormatUInt32:
		return int(to.Uint64(v))
	case FormatInt32:
		return int(to.Uint64(v))
	case FormatUInt64:
		return int(to.Uint64(v))
	case FormatBool:
		return to.Bool(v)
	default:
		return v
	}
}
