// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeBridgingState = "00000062-0000-1000-8000-0026BB765291"

type BridgingState struct {
	*Service

	Reachable           *characteristic.Reachable
	LinkQuality         *characteristic.LinkQuality
	AccessoryIdentifier *characteristic.AccessoryIdentifier
	Category            *characteristic.Category
}

func NewBridgingState() *BridgingState {
	svc := BridgingState{}
	svc.Service = New(TypeBridgingState)

	svc.Reachable = characteristic.NewReachable()
	svc.AddCharacteristic(svc.Reachable.Characteristic)

	svc.LinkQuality = characteristic.NewLinkQuality()
	svc.AddCharacteristic(svc.LinkQuality.Characteristic)

	svc.AccessoryIdentifier = characteristic.NewAccessoryIdentifier()
	svc.AddCharacteristic(svc.AccessoryIdentifier.Characteristic)

	svc.Category = characteristic.NewCategory()
	svc.AddCharacteristic(svc.Category.Characteristic)

	return &svc
}
