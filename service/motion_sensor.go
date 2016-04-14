// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeMotionSensor = "85"

type MotionSensor struct {
	*Service

	MotionDetected *characteristic.MotionDetected
}

func NewMotionSensor() *MotionSensor {
	svc := MotionSensor{}
	svc.Service = New(TypeMotionSensor)

	svc.MotionDetected = characteristic.NewMotionDetected()
	svc.AddCharacteristic(svc.MotionDetected.Characteristic)

	return &svc
}
