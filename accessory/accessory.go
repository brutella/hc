package accessory

import (
	"github.com/brutella/hc/service"
)

type Info struct {
	Name             string
	SerialNumber     string
	Manufacturer     string
	Model            string
	FirmwareRevision string
	ID               uint64
}

// Accessory is a HomeKit accessory.
//
// An accessory contains services, which themselves contain characteristics.
// Every accessory has the "accessory info" service by default which consists
// of characteristics to identify the accessory: name, model, manufacturer,...
type Accessory struct {
	ID       uint64             `json:"aid"`
	Services []*service.Service `json:"services"`

	Type AccessoryType                 `json:"-"`
	Info *service.AccessoryInformation `json:"-"`

	idCount    uint64
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

	if version := info.FirmwareRevision; len(version) > 0 {
		svc.FirmwareRevision.SetValue(version)
	} else {
		svc.FirmwareRevision.SetValue("undefined")
	}

	var id uint64 = 0
	if info.ID > id {
		id = info.ID
	}

	acc := &Accessory{
		idCount: 1,
		ID:      id,
		Info:    svc,
		Type:    typ,
	}

	acc.AddService(acc.Info.Service)

	svc.Identify.OnValueRemoteUpdate(func(value bool) {
		acc.Identify()
	})

	return acc
}

func (a *Accessory) GetServices() []*service.Service {
	result := make([]*service.Service, 0)
	result = append(result, a.Services...)
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
	a.Services = append(a.Services, s)
}

// UpdateIDs updates the service and characteirstic ids.
func (a *Accessory) UpdateIDs() {
	for _, s := range a.Services {
		s.ID = a.idCount
		a.idCount++

		for _, c := range s.Characteristics {
			c.ID = a.idCount
			a.idCount++
		}
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
