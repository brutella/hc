package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
)

// Accessory implements the model.Accessory interface and contains the data
// structures to communicate with HomeKit.
//
// An accessory in consists of services, which consists of characteristics.
// Every accessory has the "accessory info" service by default which consists
// of characteristics to identify the accessory: name, model, manufacturer,...
type Accessory struct {
	ID       int64              `json:"aid"`
	Services []*service.Service `json:"services"`

	Info    *service.AccessoryInfo `json:"-"`
	idCount int64

	onIdentify func()
}

// New returns an accessory which implements model.Accessory.
func New(info model.Info) *Accessory {
	i := service.NewInfo(info)
	a := &Accessory{
		idCount: 1,
		Info:    i,
		ID:      model.InvalidID,
	}

	a.AddService(i.Service)

	i.Identify.OnRemoteChange(func(c *characteristic.Characteristic, new, old interface{}) {
		if a.onIdentify != nil {
			a.onIdentify()
		}
	})

	return a
}

func (a *Accessory) SetID(id int64) {
	a.ID = id
}

func (a *Accessory) GetID() int64 {
	return a.ID
}

func (a *Accessory) GetServices() []model.Service {
	result := make([]model.Service, 0)
	for _, s := range a.Services {
		result = append(result, s)
	}
	return result
}

func (a *Accessory) Name() string {
	return a.Info.Name.Name()
}

func (a *Accessory) SerialNumber() string {
	return a.Info.Serial.SerialNumber()
}

func (a *Accessory) Manufacturer() string {
	return a.Info.Manufacturer.Manufacturer()
}

func (a *Accessory) Model() string {
	return a.Info.Model.Model()
}

func (a *Accessory) Firmware() string {
	firmware := a.Info.Firmware
	if firmware != nil {
		return firmware.Revision()
	}
	return ""
}

func (a *Accessory) Hardware() string {
	hardware := a.Info.Hardware
	if hardware != nil {
		return hardware.Revision()
	}
	return ""
}

func (a *Accessory) Software() string {
	software := a.Info.Software
	if software != nil {
		return software.Revision()
	}
	return ""
}

func (a *Accessory) OnIdentify(fn func()) {
	a.onIdentify = fn
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
