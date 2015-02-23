package container

import (
	"github.com/brutella/hap/model"
	"github.com/brutella/hap/model/accessory"
)

type Container struct {
	Accessories []*accessory.Accessory `json:"accessories"`

	idCount int64
}

func NewContainer() *Container {
	return &Container{
		Accessories: make([]*accessory.Accessory, 0),
		idCount:     1,
	}
}

func (m *Container) AddAccessory(a *accessory.Accessory) {
	// Set accessory id when invalid
	if a.GetId() == model.InvalidId {
		a.SetId(m.idCount)
		m.idCount += 1
	}

	m.Accessories = append(m.Accessories, a)
}

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
