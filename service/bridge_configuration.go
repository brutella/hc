// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeBridgeConfiguration = "000000A1-0000-1000-8000-0026BB765291"

type BridgeConfiguration struct {
	*Service

	ConfigureBridgedAccessoryStatus *characteristic.ConfigureBridgedAccessoryStatus
	DiscoverBridgedAccessories      *characteristic.DiscoverBridgedAccessories
	DiscoveredBridgedAccessories    *characteristic.DiscoveredBridgedAccessories
	ConfigureBridgedAccessory       *characteristic.ConfigureBridgedAccessory
}

func NewBridgeConfiguration() *BridgeConfiguration {
	svc := BridgeConfiguration{}
	svc.Service = New(TypeBridgeConfiguration)

	svc.ConfigureBridgedAccessoryStatus = characteristic.NewConfigureBridgedAccessoryStatus()
	svc.AddCharacteristic(svc.ConfigureBridgedAccessoryStatus.Characteristic)

	svc.DiscoverBridgedAccessories = characteristic.NewDiscoverBridgedAccessories()
	svc.AddCharacteristic(svc.DiscoverBridgedAccessories.Characteristic)

	svc.DiscoveredBridgedAccessories = characteristic.NewDiscoveredBridgedAccessories()
	svc.AddCharacteristic(svc.DiscoveredBridgedAccessories.Characteristic)

	svc.ConfigureBridgedAccessory = characteristic.NewConfigureBridgedAccessory()
	svc.AddCharacteristic(svc.ConfigureBridgedAccessory.Characteristic)

	return &svc
}
