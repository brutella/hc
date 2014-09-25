package model

import(
)

type Accessory struct {
    Id int `json:"aid"`
    Services []*Service `json:"services"`
    
    idCount int
}

func NewAccessory() *Accessory {
    return &Accessory{
        idCount: 1,
    }
}

func (a *Accessory) AddService(s *Service) {
    s.Id = a.idCount
    a.idCount += 1
    
    for _, c := range s.Characteristics {
        c.Id = a.idCount
        a.idCount += 1
    }
    
    a.Services = append(a.Services, s)
}