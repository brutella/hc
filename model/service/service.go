package service

import(
    "github.com/brutella/hap/model"
)

type Service struct {
    Id int                            `json:"iid"`
    Type ServiceType                  `json:"type"`
    Characteristics []model.Characteristic `json:"characteristics"`
}

func NewService() *Service {
    s := Service{
        Characteristics: []model.Characteristic{},
    }
    
    return &s
}

func (s *Service) AddCharacteristic(c model.Characteristic) {
    s.Characteristics = append(s.Characteristics, c)
}

// model.Service
func (s *Service) SetId(id int) {
    s.Id = id
}

func (s *Service) GetCharacteristics()[]model.Characteristic {
    return s.Characteristics
}

// Compareable
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
