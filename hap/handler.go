package hap

import (
	"github.com/brutella/hc/util"
)

// A ContainerHandler abstracts request/response communication
type ContainerHandler interface {
	Handle(util.Container) (util.Container, error)
}

// A PairVerifyHandler is a ContainerHandler which negotations a shared key.
type PairVerifyHandler interface {
	ContainerHandler
	SharedKey() [32]byte
}
