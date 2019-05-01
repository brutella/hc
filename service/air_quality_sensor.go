// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeAirQualitySensor = "8D"

type AirQualitySensor struct {
	*Service

	AirQuality *characteristic.AirQuality

	StatusActive           *characteristic.StatusActive
	StatusFault            *characteristic.StatusFault
	StatusTampered         *characteristic.StatusTampered
	StatusLowBattery       *characteristic.StatusLowBattery
	Name                   *characteristic.Name
	OzoneDensity           *characteristic.OzoneDensity
	NitrogenDioxideDensity *characteristic.NitrogenDioxideDensity
	SulphurDioxideDensity  *characteristic.SulphurDioxideDensity
	PM2_5Density           *characteristic.PM2_5Density
	PM10Density            *characteristic.PM10Density
	VOCDensity             *characteristic.VOCDensity
	CarbonMonoxideLevel    *characteristic.CarbonMonoxideLevel
	CarbonDioxideLevel     *characteristic.CarbonDioxideLevel
}

func NewAirQualitySensor() *AirQualitySensor {
	svc := AirQualitySensor{}
	svc.Service = New(TypeAirQualitySensor)

	svc.AirQuality = characteristic.NewAirQuality()
	svc.AddCharacteristic(svc.AirQuality.Characteristic)

	svc.StatusActive = characteristic.NewStatusActive()
	svc.AddCharacteristic(svc.StatusActive.Characteristic)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddCharacteristic(svc.StatusFault.Characteristic)

	svc.StatusTampered = characteristic.NewStatusTampered()
	svc.AddCharacteristic(svc.StatusTampered.Characteristic)

	svc.StatusLowBattery = characteristic.NewStatusLowBattery()
	svc.AddCharacteristic(svc.StatusLowBattery.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.OzoneDensity = characteristic.NewOzoneDensity()
	svc.AddCharacteristic(svc.OzoneDensity.Characteristic)

	svc.NitrogenDioxideDensity = characteristic.NewNitrogenDioxideDensity()
	svc.AddCharacteristic(svc.NitrogenDioxideDensity.Characteristic)

	svc.SulphurDioxideDensity = characteristic.NewSulphurDioxideDensity()
	svc.AddCharacteristic(svc.SulphurDioxideDensity.Characteristic)

	svc.PM2_5Density = characteristic.NewPM2_5Density()
	svc.AddCharacteristic(svc.PM2_5Density.Characteristic)

	svc.PM10Density = characteristic.NewPM10Density()
	svc.AddCharacteristic(svc.PM10Density.Characteristic)

	svc.VOCDensity = characteristic.NewVOCDensity()
	svc.AddCharacteristic(svc.VOCDensity.Characteristic)

	svc.CarbonMonoxideLevel = characteristic.NewCarbonMonoxideLevel()
	svc.AddCharacteristic(svc.CarbonMonoxideLevel.Characteristic)

	svc.CarbonDioxideLevel = characteristic.NewCarbonDioxideLevel()
	svc.AddCharacteristic(svc.CarbonDioxideLevel.Characteristic)

	return &svc
}
