package accessory

import (
	"github.com/brutella/hc/service"
)

// Doorbell provides RTP video streaming, Speaker and Mic controls
type Doorbell struct {
	*Accessory
	Control           *service.Doorbell
	StreamManagement1 *service.CameraRTPStreamManagement
	Speaker 	  *service.Speaker
	Microphone	  *service.Microphone
}

// NewDoorbell returns a Video Doorbell accessory.
func NewDoorbell(info Info) *Doorbell {
	acc := Doorbell{}
	acc.Accessory = New(info, TypeVideoDoorbell)
	acc.Control = service.NewDoorbell()
	acc.AddService(acc.Control.Service)

	acc.StreamManagement1 = service.NewCameraRTPStreamManagement()
	acc.AddService(acc.StreamManagement1.Service)

	acc.Speaker = service.NewSpeaker()
	acc.AddService(acc.Speaker.Service)

	acc.Microphone = service.NewMicrophone()
	acc.AddService(acc.Microphone.Service)

	return &acc
}
