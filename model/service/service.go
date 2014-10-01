package service

import(
    "github.com/brutella/hap/model/characteristic"
)

type Service struct {
    Id int                            `json:"iid"`
    Type ServiceType                  `json:"type"`
    Characteristics []*characteristic.Characteristic `json:"characteristics"`
}

func NewService() *Service {
    return &Service{}
}

func (s *Service) AddCharacteristic(c *characteristic.Characteristic) {
    s.Characteristics = append(s.Characteristics, c)
}

func (s *Service) Equal(other interface{}) bool {
    if service, ok := other.(*Service); ok == true {
        if len(s.Characteristics) != len(service.Characteristics) {
            println("Number of chars wrong")
            return false
        }
        
        for i, c := range s.Characteristics {
            other := service.Characteristics[i]
            if c.Equal(other) == false {
                return false
            }
        }
        
        return s.Id == service.Id && s.Type == service.Type
    }
    
    return false
}
