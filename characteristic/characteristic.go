package characteristic

import (
	"fmt"
	"github.com/gosexy/to"
	"net"
	"reflect"
)

type ConnChangeFunc func(conn net.Conn, c *Characteristic, newValue, oldValue interface{})
type ChangeFunc func(c *Characteristic, newValue, oldValue interface{})

// Characteristic is a HomeKit characteristic.
type Characteristic struct {
	ID          int64    `json:"iid"` // managed by accessory
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
	
	AccessoryID	int64 `json:"-"`	// 'parent'
}

// writeOnlyPerms returns true when permissions only include write permission
func writeOnlyPerms(permissions []string) bool {
	if len(permissions) == 1 {
		return permissions[0] == PermWrite
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

// hasEvents returns true when permissions include events
func hasEvents(permissions []string) bool {
	for _, value := range permissions {
		if value == PermEvents {
			return true
		}
	}
	return false
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

func (c *Characteristic) UpdateValue(value interface{}) {
	c.updateValue(value, nil)
}

func (c *Characteristic) UpdateValueFromConnection(value interface{}, conn net.Conn) {
	c.updateValue(value, conn)
}

func (c *Characteristic) SetEventsEnabled(enable bool) {
	c.Events = enable
}

func (c *Characteristic) EventsEnabled() bool {
	return c.Events
}


// Activate events for this characteristic if it contains events
func (c *Characteristic) ActivateEvents() {
	c.Events = hasEvents(c.Perms)
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

// model.Characteristic
func (c *Characteristic) SetID(id int64) {
	c.ID = id
}

// Set parent accessory ID
func (c *Characteristic) SetAccessoryID(id int64) {
	c.AccessoryID = id
}

// Set characteristic permissions
func (c *Characteristic) SetPerms(perms []string) {
	c.Perms = perms
}

// Set characteristic description
func (c *Characteristic) SetDescription(description string) {
	c.Description = description
}

// Set characteristic format
func (c *Characteristic) SetFormat(format string) {
	c.Format = format
}

// Set characteristic units
func (c *Characteristic) SetUnit(unit string) {
	if unit != "" {
		c.Unit = unit
	}
}

// Set characteristic minimum value
func (c *Characteristic) SetMinValue(minValue interface{}) {
	if minValue != nil {
		c.MinValue = minValue
	}
}

// Set characteristic maximum value
func (c *Characteristic) SetMaxValue(maxValue interface{}) {
	if maxValue != nil {
		c.MaxValue = maxValue
	}
}

// Set characteristic step value
func (c *Characteristic) SetStepValue(stepValue interface{}) {
	if stepValue != nil {
		c.StepValue = stepValue
	}
}

func (c *Characteristic) GetID() int64 {
	return c.ID
}

// Return the ID of the parent accessory
func (c *Characteristic) GetAccessoryID() int64 {
	return c.AccessoryID
}

// Private

func (c *Characteristic) isWriteOnly() bool {
	return writeOnlyPerms(c.Perms)
}

func (c *Characteristic) hasWritePerms() bool {
	return noWritePerms(c.Perms) == false
}

// Sets the value of the characteristic
// The implementation makes sure that the type of the value stays the same
// E.g. Type of characteristic value int, calling updateValue("10.5") sets the value to int(10)
//
// When permissions are write only, this methods does not set the Value field.
func (c *Characteristic) updateValue(value interface{}, conn net.Conn) {
	if c.Value != nil {
		if converted, err := to.Convert(value, reflect.TypeOf(c.Value).Kind()); err == nil {
			value = converted
		}
	}

	// TODO: All the int/uint's must be split into their 'true' data-types. Cannot gurantee that a UInt64 can fit into
	// an 'int' on all platforms
	// Value must be within min and max
	switch c.Format {
	case FormatFloat:
		value = c.boundFloat64Value(value.(float64))
	case FormatUInt8:
		value = c.boundUInt8Value(value.(uint8))
	case FormatUInt16:
		value = c.boundUInt16Value(value.(uint16))
	case FormatUInt32, FormatUInt64, FormatInt32:
		value = c.boundIntValue(value.(int))
	}


	// Ignore when new value is same
	if c.Value == value {
		return
	}

	// Ignore new values from remote when permissions don't allow write
	if c.hasWritePerms() == false && conn != nil {
		return
	}

	old := c.Value
	if c.isWriteOnly() == false {
		c.Value = value
	} else {
		c.Value = nil
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

func (c *Characteristic) boundUInt8Value(value uint8) interface{} {
	min, minOK := c.MinValue.(uint8)
	max, maxOK := c.MaxValue.(uint8)
	if maxOK == true && value > max {
		value = max
	} else if minOK == true && value < min {
		value = min
	}

	return value
}

func (c *Characteristic) boundUInt16Value(value uint16) interface{} {
	min, minOK := c.MinValue.(uint16)
	max, maxOK := c.MaxValue.(uint16)
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
