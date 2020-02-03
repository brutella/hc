package accessory

import (
	"github.com/brutella/hc/service"
)

//CameraRTPStreamManagement struct
type CameraRTPStreamManagement struct {
	*Accessory
	CameraRTPStreamManagement *service.CameraRTPStreamManagement
}

//NewCameraRTPStreamManagement function
func NewCameraRTPStreamManagement(info Info) *CameraRTPStreamManagement {
	acc := CameraRTPStreamManagement{}

	acc.Accessory = New(info, TypeVideoDoorbell) //TypeIPCamera)
	acc.CameraRTPStreamManagement = service.NewCameraRTPStreamManagement()

	acc.AddService(acc.CameraRTPStreamManagement.Service)

	return &acc
}
