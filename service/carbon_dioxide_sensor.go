// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeCarbonDioxideSensor = "97"

type CarbonDioxideSensor struct {
	*Service

	CarbonDioxideDetected *characteristic.CarbonDioxideDetected
}

func NewCarbonDioxideSensor() *CarbonDioxideSensor {
	svc := CarbonDioxideSensor{}
	svc.Service = New(TypeCarbonDioxideSensor)

	svc.CarbonDioxideDetected = characteristic.NewCarbonDioxideDetected()
	svc.AddCharacteristic(svc.CarbonDioxideDetected.Characteristic)

	return &svc
}
