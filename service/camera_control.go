// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeCameraControl = "111"

type CameraControl struct {
	*Service

	On *characteristic.On
}

func NewCameraControl() *CameraControl {
	svc := CameraControl{}
	svc.Service = New(TypeCameraControl)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	return &svc
}
