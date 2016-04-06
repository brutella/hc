// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLeakSensor = "00000083-0000-1000-8000-0026BB765291"

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
