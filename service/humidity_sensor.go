// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeHumiditySensor = "82"

type HumiditySensor struct {
	*Service

	CurrentRelativeHumidity *characteristic.CurrentRelativeHumidity
}

func NewHumiditySensor() *HumiditySensor {
	svc := HumiditySensor{}
	svc.Service = New(TypeHumiditySensor)

	svc.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	svc.AddCharacteristic(svc.CurrentRelativeHumidity.Characteristic)

	return &svc
}
