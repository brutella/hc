// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeCameraRecordingManagement = "204"

type CameraRecordingManagement struct {
	*Service

	SupportedCameraRecordingConfiguration *characteristic.SupportedCameraRecordingConfiguration
	SupportedVideoRecordingConfiguration  *characteristic.SupportedVideoRecordingConfiguration
	SupportedAudioRecordingConfiguration  *characteristic.SupportedAudioRecordingConfiguration
	SelectedCameraRecordingConfiguration  *characteristic.SelectedCameraRecordingConfiguration
}

func NewCameraRecordingManagement() *CameraRecordingManagement {
	svc := CameraRecordingManagement{}
	svc.Service = New(TypeCameraRecordingManagement)

	svc.SupportedCameraRecordingConfiguration = characteristic.NewSupportedCameraRecordingConfiguration()
	svc.AddCharacteristic(svc.SupportedCameraRecordingConfiguration.Characteristic)

	svc.SupportedVideoRecordingConfiguration = characteristic.NewSupportedVideoRecordingConfiguration()
	svc.AddCharacteristic(svc.SupportedVideoRecordingConfiguration.Characteristic)

	svc.SupportedAudioRecordingConfiguration = characteristic.NewSupportedAudioRecordingConfiguration()
	svc.AddCharacteristic(svc.SupportedAudioRecordingConfiguration.Characteristic)

	svc.SelectedCameraRecordingConfiguration = characteristic.NewSelectedCameraRecordingConfiguration()
	svc.AddCharacteristic(svc.SelectedCameraRecordingConfiguration.Characteristic)

	return &svc
}
