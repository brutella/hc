package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeWifiTransport = "22A"

type WifiTransport struct {
	*Service

	CurrentTransport         *characteristic.CurrentTransport
	WifiCapabilities         *characteristic.WifiCapabilities
	WifiConfigurationControl *characteristic.WifiConfigurationControl
}

func NewWifiTransport() *WifiTransport {
	svc := WifiTransport{}
	svc.Service = New(TypeWifiTransport)

	svc.CurrentTransport = characteristic.NewCurrentTransport()
	svc.AddCharacteristic(svc.CurrentTransport.Characteristic)

	svc.WifiCapabilities = characteristic.NewWifiCapabilities()
	svc.AddCharacteristic(svc.WifiCapabilities.Characteristic)

	svc.WifiConfigurationControl = characteristic.NewWifiConfigurationControl()
	svc.AddCharacteristic(svc.WifiConfigurationControl.Characteristic)

	return &svc
}
