// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLockManagement = "44"

type LockManagement struct {
	*Service

	LockControlPoint *characteristic.LockControlPoint
	Version          *characteristic.Version

	Logs                              *characteristic.Logs
	AudioFeedback                     *characteristic.AudioFeedback
	LockManagementAutoSecurityTimeout *characteristic.LockManagementAutoSecurityTimeout
	AdministratorOnlyAccess           *characteristic.AdministratorOnlyAccess
	LockLastKnownAction               *characteristic.LockLastKnownAction
	CurrentDoorState                  *characteristic.CurrentDoorState
	MotionDetected                    *characteristic.MotionDetected
	Name                              *characteristic.Name
}

func NewLockManagement() *LockManagement {
	svc := LockManagement{}
	svc.Service = New(TypeLockManagement)

	svc.LockControlPoint = characteristic.NewLockControlPoint()
	svc.AddCharacteristic(svc.LockControlPoint.Characteristic)

	svc.Version = characteristic.NewVersion()
	svc.AddCharacteristic(svc.Version.Characteristic)

	return &svc
}

func (svc *LockManagement) addOptionalCharaterics() {

	svc.Logs = characteristic.NewLogs()
	svc.AddCharacteristic(svc.Logs.Characteristic)

	svc.AudioFeedback = characteristic.NewAudioFeedback()
	svc.AddCharacteristic(svc.AudioFeedback.Characteristic)

	svc.LockManagementAutoSecurityTimeout = characteristic.NewLockManagementAutoSecurityTimeout()
	svc.AddCharacteristic(svc.LockManagementAutoSecurityTimeout.Characteristic)

	svc.AdministratorOnlyAccess = characteristic.NewAdministratorOnlyAccess()
	svc.AddCharacteristic(svc.AdministratorOnlyAccess.Characteristic)

	svc.LockLastKnownAction = characteristic.NewLockLastKnownAction()
	svc.AddCharacteristic(svc.LockLastKnownAction.Characteristic)

	svc.CurrentDoorState = characteristic.NewCurrentDoorState()
	svc.AddCharacteristic(svc.CurrentDoorState.Characteristic)

	svc.MotionDetected = characteristic.NewMotionDetected()
	svc.AddCharacteristic(svc.MotionDetected.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
