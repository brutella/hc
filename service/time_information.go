// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeTimeInformation = "00000099-0000-1000-8000-0026BB765291"

type TimeInformation struct {
	*Service

	CurrentTime  *characteristic.CurrentTime
	DayOfTheWeek *characteristic.DayOfTheWeek
	TimeUpdate   *characteristic.TimeUpdate
}

func NewTimeInformation() *TimeInformation {
	svc := TimeInformation{}
	svc.Service = New(TypeTimeInformation)

	svc.CurrentTime = characteristic.NewCurrentTime()
	svc.AddCharacteristic(svc.CurrentTime.Characteristic)

	svc.DayOfTheWeek = characteristic.NewDayOfTheWeek()
	svc.AddCharacteristic(svc.DayOfTheWeek.Characteristic)

	svc.TimeUpdate = characteristic.NewTimeUpdate()
	svc.AddCharacteristic(svc.TimeUpdate.Characteristic)

	return &svc
}
