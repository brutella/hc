package service

import (
	"github.com/brutella/hc/characteristic"
)

// Service is an HomeKit service consisting of characteristics.
type Service struct {
	ID              int64                            `json:"iid"`
	Type            string                           `json:"type"`
	Characteristics []*characteristic.Characteristic `json:"characteristics"`
}

// New returns a new service.
func New(typ string) *Service {
	s := Service{
		Type:            typ,
		Characteristics: []*characteristic.Characteristic{},
	}

	return &s
}

// SetID sets the service id.
func (s *Service) SetID(id int64) {
	s.ID = id
}

// GetID returns the service id.
func (s *Service) GetID() int64 {
	return s.ID
}

// GetCharacteristics returns the characteristics which represent the service.
func (s *Service) GetCharacteristics() []*characteristic.Characteristic {
	var result []*characteristic.Characteristic
	for _, c := range s.Characteristics {
		result = append(result, c)
	}
	return result
}

// Equal returns true when receiver has the same characteristics, service id and service type as the argument.
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

		return s.ID == service.ID && s.Type == service.Type
	}

	return false
}

func (s *Service) AddCharacteristic(c *characteristic.Characteristic) {
	s.Characteristics = append(s.Characteristics, c)
}
