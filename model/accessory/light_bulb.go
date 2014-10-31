package accessory

import(
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/service"
    "github.com/brutella/hap/model/characteristic"
)

type lightBulb struct {
    *Accessory
    bulb *service.LightBulb
    
    onChanged func(bool)
    brightnessChanged func(int)
}

func NewLightBulb(info model.Info) *lightBulb {
    accessory := New(info)
    bulb := service.NewLightBulb(info.Name, false) // off
    
    accessory.AddService(bulb.Service)
    
    lightBulb := lightBulb{accessory, bulb, nil, nil}
    bulb.On.AddRemoteChangeDelegate(&lightBulb)
    bulb.Brightness.AddRemoteChangeDelegate(&lightBulb)
        
    return &lightBulb
}

func (l *lightBulb) SetOn(on bool) {
    l.bulb.On.SetOn(on)
}

func (l *lightBulb) IsOn() bool {
    return l.bulb.On.On()
}

func (l *lightBulb) GetBrightness() int {
    return l.bulb.Brightness.IntValue()
}

func (l *lightBulb) GetHue() float64  {
    return l.bulb.Hue.FloatValue()
}

func (l *lightBulb) GetSaturation() float64 {
    return l.bulb.Saturation.FloatValue()
}

func (l *lightBulb) OnStateChanged(fn func(bool)){
    l.onChanged = fn
}

func (l *lightBulb) OnBrightnessChanged(fn func(int)) {
    l.brightnessChanged = fn
}

// CharacteristicDelegate
func (l *lightBulb) CharactericDidChangeValue(c *characteristic.Characteristic, change characteristic.CharacteristicChange) {
    switch c {
    case l.bulb.On.Characteristic:
        if l.onChanged != nil {
            l.onChanged(l.bulb.On.On())
        }
    case l.bulb.Brightness.Characteristic:
        if l.brightnessChanged != nil {
            l.brightnessChanged(l.bulb.Brightness.IntValue())
        }
    }
}