// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeCarbonMonoxideSensor = "7F"

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
