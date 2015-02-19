package service

import (
	"github.com/brutella/hap/model"
	"github.com/brutella/hap/model/characteristic"
)

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

func NewInfo(info model.Info) *AccessoryInfo {
	return NewAccessoryInfo(info.Name, info.SerialNumber, info.Manufacturer, info.Model, info.Firmware, info.Hardware, info.Software)
}

func NewAccessoryInfo(accessoryName, serialNumber, manufacturerName, modelName, firmwareRevision, hardwareRevision, softwareRevision string) *AccessoryInfo {
	identify := characteristic.NewIdentify(false)
	serial := characteristic.NewSerialNumber(serialNumber)
	model := characteristic.NewModel(modelName)
	manufacturer := characteristic.NewManufacturer(manufacturerName)
	name := characteristic.NewName(accessoryName)

	service := New()
	service.Type = TypeAccessoryInfo
	service.AddCharacteristic(identify.Characteristic)
	service.AddCharacteristic(serial.Characteristic)
	service.AddCharacteristic(model.Characteristic)
	service.AddCharacteristic(manufacturer.Characteristic)
	service.AddCharacteristic(name.Characteristic)

	var firmware *characteristic.Revision
	if firmwareRevision != "" {
		firmware = characteristic.NewFirmwareRevision(firmwareRevision)
		service.AddCharacteristic(firmware.Characteristic)
	}

	var hardware *characteristic.Revision
	if hardwareRevision != "" {
		hardware = characteristic.NewHardwareRevision(hardwareRevision)
		service.AddCharacteristic(hardware.Characteristic)
	}

	var software *characteristic.Revision
	if softwareRevision != "" {
		software = characteristic.NewSoftwareRevision(softwareRevision)
		service.AddCharacteristic(software.Characteristic)
	}

	return &AccessoryInfo{service, identify, serial, model, manufacturer, name, firmware, hardware, software}
}
