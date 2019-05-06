// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeFaucet = "D7"

type Faucet struct {
	*Service

	Active *characteristic.Active

	Name        *characteristic.Name
	StatusFault *characteristic.StatusFault
}

func NewFaucet() *Faucet {
	svc := Faucet{}
	svc.Service = New(TypeFaucet)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	return &svc
}

func (svc *Faucet) addOptionalCharaterics() {

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddCharacteristic(svc.StatusFault.Characteristic)

}
