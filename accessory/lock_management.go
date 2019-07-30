package accessory

import (
	"github.com/brutella/hc/service"
)

//LockManagement structure
type LockManagement struct {
	*Accessory
	LockManagement *service.LockManagement
}

// NewLockManagement returns an outlet accessory containing one outlet service.
func NewLockManagement(info Info) *LockManagement {
	acc := LockManagement{}
	acc.Accessory = New(info, TypeDoorLock)
	acc.LockManagement = service.NewLockManagement()

	acc.AddService(acc.LockManagement.Service)

	return &acc
}
