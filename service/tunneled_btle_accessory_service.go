// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeTunneledBTLEAccessoryService = "56"

type TunneledBTLEAccessoryService struct {
	*Service

	Name                         *characteristic.Name
	AccessoryIdentifier          *characteristic.AccessoryIdentifier
	TunneledAccessoryStateNumber *characteristic.TunneledAccessoryStateNumber
	TunneledAccessoryConnected   *characteristic.TunneledAccessoryConnected
	TunneledAccessoryAdvertising *characteristic.TunneledAccessoryAdvertising
	TunnelConnectionTimeout      *characteristic.TunnelConnectionTimeout
}

func NewTunneledBTLEAccessoryService() *TunneledBTLEAccessoryService {
	svc := TunneledBTLEAccessoryService{}
	svc.Service = New(TypeTunneledBTLEAccessoryService)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.AccessoryIdentifier = characteristic.NewAccessoryIdentifier()
	svc.AddCharacteristic(svc.AccessoryIdentifier.Characteristic)

	svc.TunneledAccessoryStateNumber = characteristic.NewTunneledAccessoryStateNumber()
	svc.AddCharacteristic(svc.TunneledAccessoryStateNumber.Characteristic)

	svc.TunneledAccessoryConnected = characteristic.NewTunneledAccessoryConnected()
	svc.AddCharacteristic(svc.TunneledAccessoryConnected.Characteristic)

	svc.TunneledAccessoryAdvertising = characteristic.NewTunneledAccessoryAdvertising()
	svc.AddCharacteristic(svc.TunneledAccessoryAdvertising.Characteristic)

	svc.TunnelConnectionTimeout = characteristic.NewTunnelConnectionTimeout()
	svc.AddCharacteristic(svc.TunnelConnectionTimeout.Characteristic)

	return &svc
}
