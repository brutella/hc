// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeCarbonMonoxideSensor = "0000007F-0000-1000-8000-0026BB765291"

type CarbonMonoxideSensor struct {
	*Service

	CarbonMonoxideDetected *characteristic.CarbonMonoxideDetected
}

func NewCarbonMonoxideSensor() *CarbonMonoxideSensor {
	svc := CarbonMonoxideSensor{}
	svc.Service = New(TypeCarbonMonoxideSensor)

	svc.CarbonMonoxideDetected = characteristic.NewCarbonMonoxideDetected()
	svc.AddCharacteristic(svc.CarbonMonoxideDetected.Characteristic)

	return &svc
}
