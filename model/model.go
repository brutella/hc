package model

import(
    "github.com/brutella/hap/model/accessory"
)

type Model struct {
    Accessories []*accessory.Accessory `json:"accessories"`
    
    idCount int
}

func NewModel() *Model {
    return &Model{
        Accessories: make([]*accessory.Accessory, 0),
        idCount: 1,
    }
}

func (m *Model) AddAccessory(a *accessory.Accessory) {
    a.Id = m.idCount
    m.idCount += 1
    m.Accessories = append(m.Accessories, a)
}

// TODO write tests
func (m *Model) RemoveAccessory(a *accessory.Accessory) {
    for i, accessory := range m.Accessories {
        if accessory == a {
            m.Accessories = append(m.Accessories[:i], m.Accessories[i+1:]...)
        }
    }
}

func (m *Model) Equal(other interface{}) bool {
    if model, ok := other.(*Model); ok == true {
        if len(m.Accessories) != len(model.Accessories) {
            return false
        }
        
        for i, a := range m.Accessories {
            if a.Equal(model.Accessories[i]) == false {
                return false
            }
        }
        
        return true
    }
    
    return false
}
