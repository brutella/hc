package model

type Service struct {
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