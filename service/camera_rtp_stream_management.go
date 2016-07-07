// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeCameraRTPStreamManagement = "110"

type CameraRTPStreamManagement struct {
	*Service

	SupportedVideoStreamConfiguration *characteristic.SupportedVideoStreamConfiguration
	SupportedAudioStreamConfiguration *characteristic.SupportedAudioStreamConfiguration
	SupportedRTPConfiguration         *characteristic.SupportedRTPConfiguration
	SelectedStreamConfiguration       *characteristic.SelectedStreamConfiguration
	StreamingStatus                   *characteristic.StreamingStatus
	SetupEndpoints                    *characteristic.SetupEndpoints
}

func NewCameraRTPStreamManagement() *CameraRTPStreamManagement {
	svc := CameraRTPStreamManagement{}
	svc.Service = New(TypeCameraRTPStreamManagement)

	svc.SupportedVideoStreamConfiguration = characteristic.NewSupportedVideoStreamConfiguration()
	svc.AddCharacteristic(svc.SupportedVideoStreamConfiguration.Characteristic)

	svc.SupportedAudioStreamConfiguration = characteristic.NewSupportedAudioStreamConfiguration()
	svc.AddCharacteristic(svc.SupportedAudioStreamConfiguration.Characteristic)

	svc.SupportedRTPConfiguration = characteristic.NewSupportedRTPConfiguration()
	svc.AddCharacteristic(svc.SupportedRTPConfiguration.Characteristic)

	svc.SelectedStreamConfiguration = characteristic.NewSelectedStreamConfiguration()
	svc.AddCharacteristic(svc.SelectedStreamConfiguration.Characteristic)

	svc.StreamingStatus = characteristic.NewStreamingStatus()
	svc.AddCharacteristic(svc.StreamingStatus.Characteristic)

	svc.SetupEndpoints = characteristic.NewSetupEndpoints()
	svc.AddCharacteristic(svc.SetupEndpoints.Characteristic)

	return &svc
}
