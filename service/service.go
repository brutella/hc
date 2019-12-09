package service

import (
	"github.com/brutella/hc/characteristic"

	"encoding/json"
)

// Service is an HomeKit service consisting of characteristics.
type Service struct {
	ID              int64
	Type            string
	Characteristics []*characteristic.Characteristic
	Hidden          bool
	Primary         bool
	Linked          []*Service
}

// New returns a new service.
func New(typ string) *Service {
	s := Service{
		Type:            typ,
		Characteristics: []*characteristic.Characteristic{},
		Linked:          []*Service{},
	}

	return &s
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

func (s *Service) AddLinkedService(other *Service) {
	s.Linked = append(s.Linked, other)
}

type servicePayload struct {
	ID              int64                            `json:"iid"`
	Type            string                           `json:"type"`
	Characteristics []*characteristic.Characteristic `json:"characteristics"`
	Hidden          *bool                            `json:"hidden,omitempty"`
	Primary         *bool                            `json:"primary,omitempty"`
	Linked          []int64                          `json:"linked,omitempty"`
}

func (s *Service) MarshalJSON() ([]byte, error) {
	ids := []int64{}
	for _, s := range s.Linked {
		ids = append(ids, s.ID)
	}

	p := servicePayload{
		ID:              s.ID,
		Type:            s.Type,
		Characteristics: s.Characteristics,
		Linked:          ids,
	}

	if s.Hidden {
		p.Hidden = &s.Hidden
	}

	if s.Primary {
		p.Primary = &s.Primary
	}

	return json.Marshal(p)
}
