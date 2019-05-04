package service

import (
	"github.com/brutella/hc/characteristic"
)

// Service is an HomeKit service consisting of characteristics.
type Service struct {
	ID              int64                            `json:"iid"`
	Type            string                           `json:"type"`
	Characteristics []*characteristic.Characteristic `json:"characteristics"`
	Hidden          *bool                            `json:"hidden,omitempty"`
	Primary         *bool                            `json:"primary,omitempty"`
	Linked          []int64                          `json:"linked,omitempty"`
}

// New returns a new service.
func New(typ string) *Service {
	s := Service{
		Type:            typ,
		Characteristics: []*characteristic.Characteristic{},
		Linked:          []int64{},
	}

	return &s
}

// GetType of the service.
func (s *Service) GetType() string {
	return s.Type
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

func (s *Service) SetHidden(b bool) {
	s.Hidden = &b
}

func (s *Service) IsHidden() bool {
	if s.Hidden != nil && *s.Hidden == true {
		return true
	}
	return false
}

func (s *Service) SetPrimary(b bool) {
	s.Primary = &b
}

func (s *Service) IsPrimary() bool {
	if s.Primary != nil && *s.Primary == true {
		return true
	}
	return false
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
	if other.ID == 0 {
		panic("adding a linked should be done after the server was added to an accessory")
	}
	s.Linked = append(s.Linked, other.ID)
}
