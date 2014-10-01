package model

type Service struct {
    Compareable
    Id int `json:"iid"`
    Type ServiceType `json:"type"`
    Characteristics []*Characteristic `json:"characteristics"`
}

func NewService() *Service {
    return &Service{}
}

func (s *Service) AddCharacteristic(c *Characteristic) {
    s.Characteristics = append(s.Characteristics, c)
}

func (s *Service) Equal(other interface{}) bool {
    if service, ok := other.(*Service); ok == true {
        if len(s.Characteristics) != len(service.Characteristics) {
            return false
        }
        
        for i, c := range s.Characteristics {
            if c.Equal(service.Characteristics[i]) == false {
                return false
            }
        }
        
        return s.Id == service.Id && s.Type == service.Type
    }
    
    return false
}
