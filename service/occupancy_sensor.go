// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeOccupancySensor = "00000086-0000-1000-8000-0026BB765291"

type OccupancySensor struct {
	*Service

	OccupancyDetected *characteristic.OccupancyDetected
}

func NewOccupancySensor() *OccupancySensor {
	svc := OccupancySensor{}
	svc.Service = New(TypeOccupancySensor)

	svc.OccupancyDetected = characteristic.NewOccupancyDetected()
	svc.AddCharacteristic(svc.OccupancyDetected.Characteristic)

	return &svc
}
