package accessory

import (
	"github.com/brutella/hc/service"
)

// Camera provides RTP video streaming.
type Camera struct {
	*Accessory
	Control           *service.CameraControl
	StreamManagement1 *service.CameraRTPStreamManagement
	StreamManagement2 *service.CameraRTPStreamManagement
}

// NewCamera returns an IP camera accessory.
func NewCamera(info Info) *Camera {
	acc := Camera{}
	acc.Accessory = New(info, TypeIPCamera)
	acc.Control = service.NewCameraControl()
	acc.AddService(acc.Control.Service)

	// TODO (mah) a camera must support at least 2 rtp streams
	acc.StreamManagement1 = service.NewCameraRTPStreamManagement()
	acc.StreamManagement2 = service.NewCameraRTPStreamManagement()
	acc.AddService(acc.StreamManagement1.Service)
	// acc.AddService(acc.StreamManagement2.Service)

	return &acc
}
