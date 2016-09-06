package accessory

import (
	"crypto/md5"
	"encoding/json"
	"log"
)

// Container manages a list of accessories.
type Container struct {
	Accessories []*Accessory `json:"accessories"`

	idCount int64
}

// NewContainer returns a container.
func NewContainer() *Container {
	return &Container{
		Accessories: make([]*Accessory, 0),
		idCount:     1,
	}
}

// AddAccessory adds an accessory to the container.
// This method ensures that the accessory ids are valid and unique withing the container.
func (m *Container) AddAccessory(a *Accessory) {
	// set aid for all 'child' characteristics
	for _, s := range a.GetServices() {
		for _, c := range s.GetCharacteristics() {
			c.SetAccessoryID(m.idCount)
		}
	}
	
	a.SetID(m.idCount)
	m.idCount++
	m.Accessories = append(m.Accessories, a)
}

// Returns accessory that have the specified id
func (m *Container) GetAccessoryByID(aid int64) (*Accessory) {
	for _, accessory := range m.Accessories {
		if accessory.GetID() == aid {
			return accessory
		}
	}
	
	return nil
}

// RemoveAccessory removes an accessory from the container.
func (m *Container) RemoveAccessory(a *Accessory) {
	for i, accessory := range m.Accessories {
		if accessory == a {
			m.Accessories = append(m.Accessories[:i], m.Accessories[i+1:]...)
		}
	}
}

// Equal returns true when receiver has the same accessories as the argument.
func (m *Container) Equal(other interface{}) bool {
	if container, ok := other.(*Container); ok == true {
		if len(m.Accessories) != len(container.Accessories) {
			return false
		}

		for i, a := range m.Accessories {
			if a.Equal(container.Accessories[i]) == false {
				return false
			}
		}
		return true
	}

	return false
}

// AccessoryType returns the accessory type identifier for the accessories inside the container.
func (m *Container) AccessoryType() AccessoryType {
	if as := m.Accessories; len(as) > 0 {
		if len(as) > 1 {
			return TypeBridge
		}

		return as[0].Type
	}

	return TypeOther
}

// ContentHash returns a hash of the content (ignoring the value field).
func (m *Container) ContentHash() []byte {
	var b []byte
	var err error

	if b, err = json.Marshal(m); err != nil {
		log.Fatal(err)
	}

	val := map[string]interface{}{}
	if err := json.Unmarshal(b, &val); err != nil {
		log.Fatal(err)
	}

	deleteFieldFromDict(&val, "value")

	if b, err = json.Marshal(val); err != nil {
		log.Fatal(err)
	}

	h := md5.New()
	h.Write(b)
	return h.Sum(nil)
}

func deleteFieldFromDict(val *map[string]interface{}, field string) {
	for k, v := range *val {
		if k == field {
			delete(*val, k)
		} else {
			deleteFieldFromInterface(&v, field)
		}
	}
}

func deleteFieldFromArray(val *[]interface{}, field string) {
	for _, v := range *val {
		deleteFieldFromInterface(&v, field)
	}
}

func deleteFieldFromInterface(val *interface{}, field string) {
	v := *val

	if dict, ok := v.(map[string]interface{}); ok == true {
		deleteFieldFromDict(&dict, field)
	}

	if array, ok := v.([]interface{}); ok == true {
		deleteFieldFromArray(&array, field)
	}
}
