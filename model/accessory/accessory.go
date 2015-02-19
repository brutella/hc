package accessory

import (
	"github.com/brutella/hap/model"
	"github.com/brutella/hap/model/characteristic"
	"github.com/brutella/hap/model/service"
)

// This clas holds the main data structure to communicate with HomeKit
//
// An accessory in general consists of services, which consists of characteristics.
// Every accessory has the "accessory info" service by default which consists
// of characteristics to identify the accessory: name, model, manufacturer,...
type Accessory struct {
	Id       int64              `json:"aid"`
	Services []*service.Service `json:"services"`

	Info    *service.AccessoryInfo `json:"-"`
	idCount int64

	onIdentify func()
}

func New(info model.Info) *Accessory {
	i := service.NewInfo(info)
	a := &Accessory{
		idCount: 1,
		Info:    i,
		Id:      model.InvalidId,
	}

	a.AddService(i.Service)

	i.Identify.OnRemoteChange(func(c *characteristic.Characteristic, v interface{}) {
		if a.onIdentify != nil && a.Info.Identify.Identify() == true {
			a.onIdentify()
		}
		c.SetValue(false)
	})

	return a
}

func (a *Accessory) SetId(id int64) {
	a.Id = id
}

func (a *Accessory) GetId() int64 {
	return a.Id
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
	s.SetId(a.idCount)
	a.idCount += 1

	for _, c := range s.Characteristics {
		c.SetId(a.idCount)
		a.idCount += 1
	}

	a.Services = append(a.Services, s)
}

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

		return a.Id == accessory.Id
	}

	return false
}
