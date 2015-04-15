package service

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
)

// AccessoryInfo is a service to describe an accessory.
type AccessoryInfo struct {
	*Service

	Identify     *characteristic.Identify
	Serial       *characteristic.SerialNumber
	Model        *characteristic.Model
	Manufacturer *characteristic.Manufacturer
	Name         *characteristic.Name

	// Optional
	Firmware *characteristic.Revision
	Hardware *characteristic.Revision
	Software *characteristic.Revision
}

// NewInfo returns a accessory info
func NewInfo(info model.Info) *AccessoryInfo {
	return NewAccessoryInfo(info.Name, info.SerialNumber, info.Manufacturer, info.Model, info.Firmware, info.Hardware, info.Software)
}

// NewAccessoryInfo returns a accessory info
func NewAccessoryInfo(accessoryName, serialNumber, manufacturerName, modelName, firmwareRevision, hardwareRevision, softwareRevision string) *AccessoryInfo {

	if len(accessoryName) == 0 {
		accessoryName = "Undefined"
	}
	if len(serialNumber) == 0 {
		serialNumber = "Undefined"
	}
	if len(manufacturerName) == 0 {
		manufacturerName = "Undefined"
	}
	if len(modelName) == 0 {
		modelName = "Undefined"
	}

	identify := characteristic.NewIdentify()
	serial := characteristic.NewSerialNumber(serialNumber)
	model := characteristic.NewModel(modelName)
	manufacturer := characteristic.NewManufacturer(manufacturerName)
	name := characteristic.NewName(accessoryName)

	service := New()
	service.Type = typeAccessoryInfo
	service.addCharacteristic(name.Characteristic)
	service.addCharacteristic(manufacturer.Characteristic)
	service.addCharacteristic(model.Characteristic)
	service.addCharacteristic(serial.Characteristic)
	service.addCharacteristic(identify.Characteristic)

	var firmware *characteristic.Revision
	if firmwareRevision != "" {
		firmware = characteristic.NewFirmwareRevision(firmwareRevision)
		service.addCharacteristic(firmware.Characteristic)
	}

	var hardware *characteristic.Revision
	if hardwareRevision != "" {
		hardware = characteristic.NewHardwareRevision(hardwareRevision)
		service.addCharacteristic(hardware.Characteristic)
	}

	var software *characteristic.Revision
	if softwareRevision != "" {
		software = characteristic.NewSoftwareRevision(softwareRevision)
		service.addCharacteristic(software.Characteristic)
	}

	return &AccessoryInfo{service, identify, serial, model, manufacturer, name, firmware, hardware, software}
}
