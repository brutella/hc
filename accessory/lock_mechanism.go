package accessory

import (
	"github.com/brutella/hc/service"
)

//LockMechanism Struct
type LockMechanism struct {
	*Accessory
	LockMechanism *service.LockMechanism
}

//NewLockMechanism returns an lock mechanism accessory
func NewLockMechanism(info Info) *LockMechanism {
	acc := LockMechanism{}
	acc.Accessory = New(info, TypeDoorLock)
	acc.LockMechanism = service.NewLockMechanism()

	acc.AddService(acc.LockMechanism.Service)

	return &acc
}
