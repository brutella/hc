// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeAccessoryInformation = "3E"

type AccessoryInformation struct {
	*Service

	Identify     *characteristic.Identify
	Manufacturer *characteristic.Manufacturer
	Model        *characteristic.Model
	Name         *characteristic.Name
	SerialNumber *characteristic.SerialNumber
}

func NewAccessoryInformation() *AccessoryInformation {
	svc := AccessoryInformation{}
	svc.Service = New(TypeAccessoryInformation)

	svc.Identify = characteristic.NewIdentify()
	svc.AddCharacteristic(svc.Identify.Characteristic)

	svc.Manufacturer = characteristic.NewManufacturer()
	svc.AddCharacteristic(svc.Manufacturer.Characteristic)

	svc.Model = characteristic.NewModel()
	svc.AddCharacteristic(svc.Model.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.SerialNumber = characteristic.NewSerialNumber()
	svc.AddCharacteristic(svc.SerialNumber.Characteristic)

	return &svc
}
