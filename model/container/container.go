package container

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
)

// Container manages a list of accessories.
type Container struct {
	Accessories []*accessory.Accessory `json:"accessories"`

	idCount int64
}

// NewContainer returns a container.
func NewContainer() *Container {
	return &Container{
		Accessories: make([]*accessory.Accessory, 0),
		idCount:     1,
	}
}

// AddAccessory adds an accessory to the container.
// This method ensures that the accessory ids are valid and unique withing the container.
func (m *Container) AddAccessory(a *accessory.Accessory) {
	// Set accessory id when invalid
	if a.GetId() == model.InvalidId {
		a.SetId(m.idCount)
		m.idCount += 1
	}

	m.Accessories = append(m.Accessories, a)
}

// RemoveAccessory removes an accessory from the container.
func (m *Container) RemoveAccessory(a *accessory.Accessory) {
	for i, accessory := range m.Accessories {
		if accessory == a {
			m.Accessories = append(m.Accessories[:i], m.Accessories[i+1:]...)
		}
	}
}

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
