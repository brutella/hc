package hk

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