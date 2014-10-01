package model


type Model struct {
    Accessories []*Accessory `json:"accessories"`
    
    idCount int
}

func NewModel() *Model {
    return &Model{
        Accessories: make([]*Accessory, 0),
        idCount: 1,
    }
}

func (m *Model) AddAccessory(a *Accessory) {
    a.Id = m.idCount
    m.idCount += 1
    m.Accessories = append(m.Accessories, a)
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
