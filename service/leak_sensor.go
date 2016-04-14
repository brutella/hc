// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLeakSensor = "83"

type LeakSensor struct {
	*Service

	LeakDetected *characteristic.LeakDetected
}

func NewLeakSensor() *LeakSensor {
	svc := LeakSensor{}
	svc.Service = New(TypeLeakSensor)

	svc.LeakDetected = characteristic.NewLeakDetected()
	svc.AddCharacteristic(svc.LeakDetected.Characteristic)

	return &svc
}
