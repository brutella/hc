package accessory

import (
	"github.com/brutella/hc/service"
	"github.com/brutella/hc/characteristic"
)

type Info struct {
	Name         string
	SerialNumber string
	Manufacturer string
	Model        string
}

// Accessory implements the model.Accessory interface and contains the data
// structures to communicate with HomeKit.
//
// An accessory in consists of services, which consists of characteristics.
// Every accessory has the "accessory info" service by default which consists
// of characteristics to identify the accessory: name, model, manufacturer,...
type Accessory struct {
	ID       int64              `json:"aid"`
	Services []*service.Service `json:"services"`

	Type AccessoryType                 `json:"-"`
	Info *service.AccessoryInformation `json:"-"`

	idCount    int64
	onIdentify func()
}

// New returns an accessory which implements model.Accessory.
func New(info Info, typ AccessoryType) *Accessory {
	svc := service.NewAccessoryInformation()

	if name := info.Name; len(name) > 0 {
		svc.Name.SetValue(name)
	} else {
		svc.Name.SetValue("undefined")
	}

	if serial := info.SerialNumber; len(serial) > 0 {
		svc.SerialNumber.SetValue(serial)
	} else {
		svc.SerialNumber.SetValue("undefined")
	}

	if manufacturer := info.Manufacturer; len(manufacturer) > 0 {
		svc.Manufacturer.SetValue(manufacturer)
	} else {
		svc.Manufacturer.SetValue("undefined")
	}

	if model := info.Model; len(model) > 0 {
		svc.Model.SetValue(model)
	} else {
		svc.Model.SetValue("undefined")
	}

	acc := &Accessory{
		idCount: 1,
		Info:    svc,
		Type:    typ,
	}

	acc.AddService(acc.Info.Service)

	svc.Identify.OnValueRemoteUpdate(func(value bool) {
		acc.Identify()
	})

	return acc
}

// NewEx returns an accessory that contains a default AccessoryInformation.
func NewEx(typ AccessoryType) *Accessory {
	svc := service.NewAccessoryInformation()

	acc := &Accessory{
		idCount: 1,
		Info:    svc,
		Type:    typ,
	}

	acc.AddService(acc.Info.Service)

	svc.Identify.OnValueRemoteUpdate(func(value bool) {
		acc.Identify()
	})

	return acc	
}

func (a *Accessory) SetID(id int64) {
	a.ID = id
}

func (a *Accessory) GetID() int64 {
	return a.ID
}

func (a *Accessory) GetServices() []*service.Service {
	result := make([]*service.Service, 0)
	for _, s := range a.Services {
		result = append(result, s)
	}
	return result
}

func (a *Accessory) OnIdentify(fn func()) {
	a.onIdentify = fn
}

func (a *Accessory) Identify() {
	if a.onIdentify != nil {
		a.onIdentify()
	}
}

// Adds a service to the accessory and updates the ids of the service and the corresponding characteristics
func (a *Accessory) AddService(s *service.Service) {
	s.SetID(a.idCount)
	a.idCount++

	for _, c := range s.Characteristics {
		c.SetID(a.idCount)
		a.idCount++
	}

	a.Services = append(a.Services, s)
}

// Add a service to the accessory. IDs are not updated
func (a *Accessory) AddServiceEx(s *service.Service) {
	if s.Type == service.TypeAccessoryInformation {
		for _, char := range s.Characteristics {
			if char.Type == characteristic.TypeManufacturer {
				if val := char.Value.(string); len(val) > 0 {
					a.Info.Manufacturer.SetValue(val)
				} else {
					a.Info.Manufacturer.SetValue("undefined")
				}
			} else if char.Type == characteristic.TypeModel {
				if val := char.Value.(string); len(val) > 0 {
					a.Info.Model.SetValue(val)
				} else {
					a.Info.Model.SetValue("undefined")
				}
			} else if char.Type == characteristic.TypeName {
				if val := char.Value.(string); len(val) > 0 {
					a.Info.Name.SetValue(val)
				} else {
					a.Info.Name.SetValue("undefined")
				}
			} else if char.Type == characteristic.TypeSerialNumber {
				if val := char.Value.(string); len(val) > 0 {
					a.Info.SerialNumber.SetValue(val)
				} else {
					a.Info.SerialNumber.SetValue("undefined")
				}
			} else if char.Type == characteristic.TypeIdentify {
				// @mikejac: since this is a "remote" accessory, identify has no meaning here
			}
		}
	} else {
		a.Services = append(a.Services, s)
	}
}

// Equal returns true when receiver has the same services and id as the argument.
func (a *Accessory) Equal(other interface{}) bool {
	if accessory, ok := other.(*Accessory); ok == true {
		if len(a.Services) != len(accessory.Services) {
			return false
		}

		for i, s := range a.Services {
			if s.Equal(accessory.Services[i]) == false {
				return false
			}
		}

		return a.ID == accessory.ID
	}

	return false
}
