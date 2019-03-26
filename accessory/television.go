package accessory

import (
	"github.com/brutella/hc/service"
)

type Television struct {
	*Accessory
	Television *service.Television
	Speaker    *service.Speaker
}

// NewTelevision returns a television accessory.
func NewTelevision(info Info) *Television {
	acc := Television{}
	acc.Accessory = New(info, TypeTelevision)
	acc.Television = service.NewTelevision()
	acc.Speaker = service.NewSpeaker()

	acc.AddService(acc.Television.Service)
	acc.AddService(acc.Speaker.Service)

	return &acc
}
