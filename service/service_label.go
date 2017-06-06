// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeServiceLabel = "CC"

type ServiceLabel struct {
	*Service

	ServiceLabelNamespace *characteristic.ServiceLabelNamespace
}

func NewServiceLabel() *ServiceLabel {
	svc := ServiceLabel{}
	svc.Service = New(TypeServiceLabel)

	svc.ServiceLabelNamespace = characteristic.NewServiceLabelNamespace()
	svc.AddCharacteristic(svc.ServiceLabelNamespace.Characteristic)

	return &svc
}
