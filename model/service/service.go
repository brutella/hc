package service

import(
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/characteristic"
)

// This class represents accessory services and constis of characteristics
type Service struct {
    Id int64            `json:"iid"`
    Type ServiceType    `json:"type"`
    Characteristics []*characteristic.Characteristic `json:"characteristics"`
}

func New() *Service {
    s := Service{
        Id: model.InvalidId,
        Characteristics: []*characteristic.Characteristic{},
    }
    
    return &s
}

func (s *Service) AddCharacteristic(c *characteristic.Characteristic) {
    s.Characteristics = append(s.Characteristics, c)
}

// model.Service
func (s *Service) SetId(id int64) {
    s.Id = id
}

func (s *Service) GetId() int64 {
    return s.Id
}

func (s *Service) GetCharacteristics()[]model.Characteristic {
    result := make([]model.Characteristic, 0)
    for _, c := range s.Characteristics {
        result = append(result, c)
    }
    return result
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
