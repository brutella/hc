// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeMotionSensor = "00000085-0000-1000-8000-0026BB765291"

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
