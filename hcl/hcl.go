// The hcl package lets you upgrade your light bulb service to support
// adaptive lighting.
package hcl

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/service"
	"github.com/brutella/hc/tlv8"

	"encoding/base64"
	"net"
	"sync"
	"time"
)

// SetupHCL configures the light bulb service to support adaptive lighting.
// The brightness and color temperature characteristics are required.
// The hue and saturation characteristics are optional.
func SetupHCL(light *service.Lightbulb, brightness *characteristic.Brightness, colorTemperature *characteristic.ColorTemperature, hue *characteristic.Hue, saturation *characteristic.Saturation) {
	config := NewSupportedTransitionConfigurationCharacteristic()
	control := NewTransitionControlCharacteristic()
	count := NewActiveTransitionCountCharacteristic()

	light.AddCharacteristic(config.Characteristic)
	light.AddCharacteristic(control.Characteristic)
	light.AddCharacteristic(count.Characteristic)

	config.OnValueGet(func() interface{} {
		log.Info.Println("hcl: get supported configuration")
		v := SupportedTransitionConfiguration{
			Configurations: []ValueTransitionConfiguration{
				ValueTransitionConfiguration{
					CharacteristicID: uint8(brightness.ID), // dangerous casting
					TransitionType:   TransitionTypeBrightness,
				},
				ValueTransitionConfiguration{
					CharacteristicID: uint8(colorTemperature.ID), // dangerous casting
					TransitionType:   TransitionTypeColorTemperature,
				},
			},
		}

		tlv8, err := tlv8.Marshal(v)
		if err != nil {
			log.Info.Println("hcl:", err)
		}
		res := base64.StdEncoding.EncodeToString(tlv8)
		log.Info.Printf("hcl: %+v\n", v)
		log.Info.Printf("hcl: %+v\n", tlv8)
		log.Info.Println(res)
		return res
	})

	var hcl *HCL
	if hcl != nil {
		count.SetValue(1)
	}

	control.OnValueGet(func() interface{} {
		log.Info.Println("hcl: get transition control status")

		if hcl != nil {
			since := time.Now().Sub(hcl.Cfg().StartDate()) * time.Millisecond
			resp := TransitionControlResponse{
				Status: TransitionControlStatus{
					ColorTemperatureIID: uint8(colorTemperature.ID),
					Params:              hcl.cfg.Params,
					SinceStart:          uint64(since),
				},
			}

			tlv8, err := tlv8.Marshal(resp)
			if err != nil {
				log.Info.Println("hcl:", err)
			}
			res := base64.StdEncoding.EncodeToString(tlv8)
			log.Info.Printf("hcl: get control %+v\n", resp)
			log.Info.Printf("hcl: get control %+v\n", tlv8)
			log.Info.Println(res)
			return res
		}

		return ""
	})

	control.OnValueRemoteUpdate(func(buf []byte) {
		log.Info.Println("hcl: set control")
		log.Info.Println("hcl:", base64.StdEncoding.EncodeToString(buf))

		var transitionControl TransitionControl
		err := tlv8.Unmarshal(buf, &transitionControl)
		log.Info.Printf("hcl: %+v\n", transitionControl)

		if err != nil {
			log.Debug.Fatalf("TransitionControl: Could not unmarshal tlv8 data: %s\n", err)
		}

		if transitionControl.Read.ColorTemperatureIID != 0 {
			if hcl == nil {
				log.Info.Println("hcl: no active transition")
				return
			}

			if hcl.cfg.ColorTemperatureIID != transitionControl.Read.ColorTemperatureIID {
				log.Info.Println("hcl: invalid color temperature iid", transitionControl.Read.ColorTemperatureIID)
				return
			}

			resp := TransitionReadResponse{
				Configuration: hcl.cfg,
			}

			setTLV8Payload(control.Bytes, resp)
		} else if transitionControl.Update.Configuration.ColorTemperatureIID != 0 {
			if hcl != nil {
				hcl.Stop()
			}

			if transitionControl.Update.Configuration.Enabled {
				hcl = NewHCL(transitionControl.Update.Configuration, brightness, colorTemperature, count)
				go func() {
					hcl.Schedule()
					hcl = nil
				}()
			} else {
				hcl.Stop()
			}

			resp := &TransitionUpdateResponse{Error: 0}
			setTLV8Payload(control.Bytes, resp)
		}
	})

	brightness.OnValueRemoteUpdate(func(v int) {
		log.Info.Println("hcl: brightness", v)
		if hcl != nil {
			hcl.UpdateColorTemperatureForCurrentIndex()
		}
	})

	valueRemoteUpdateFunc := func(conn net.Conn, c *characteristic.Characteristic, v, old interface{}) {
		if hcl != nil {
			if conn == hcl.conn {
				// ignore the value update
				log.Info.Println("hcl: ignoring change")
			} else {
				log.Info.Println("hcl: disable hcl because color temperature, hue or saturation changed")
				// Stop the HCL schedule
				hcl.Stop()
				hcl = nil
				// Set the number of active transitions to 0.
				count.SetValue(0)
			}
		}
	}

	colorTemperature.OnValueUpdateFromConn(valueRemoteUpdateFunc)
	if hue != nil {
		hue.OnValueUpdateFromConn(valueRemoteUpdateFunc)
	}

	if saturation != nil {
		saturation.OnValueUpdateFromConn(valueRemoteUpdateFunc)
	}
}

type HCL struct {
	brightness            *characteristic.Brightness
	colorTemp             *characteristic.ColorTemperature
	supportedTransition   *SupportedTransitionConfigurationCharacteristic
	transitionControl     *TransitionControlCharacteristic
	activeTransitionCount *ActiveTransitionCountCharacteristic
	cfg                   TransitionConfiguration

	stopped bool
	stop    chan struct{}
	conn    net.Conn
	i       int
	mu      sync.Mutex
}

func NewHCL(cfg TransitionConfiguration, brightness *characteristic.Brightness, colorTemp *characteristic.ColorTemperature, count *ActiveTransitionCountCharacteristic) *HCL {
	return &HCL{
		brightness:            brightness,
		colorTemp:             colorTemp,
		activeTransitionCount: count,
		cfg:                   cfg,
		stop:                  make(chan struct{}, 0),
		conn:                  characteristic.TestConn,
		mu:                    sync.Mutex{},
	}
}

func (hcl *HCL) Cfg() TransitionConfiguration {
	return hcl.cfg
}

func (hcl *HCL) Conn() net.Conn {
	return hcl.conn
}

func (hcl *HCL) ColorTemperatureAt(index int) int {
	brightnessValue := hcl.brightness.GetValue()

	min := int(hcl.cfg.Curve.ValueRange.Min)
	max := int(hcl.cfg.Curve.ValueRange.Max)

	if brightnessValue < min {
		brightnessValue = min
	} else if brightnessValue > max {
		brightnessValue = max
	}

	for i, entry := range hcl.cfg.Curve.Entries {
		if i == index {
			return int(entry.ColorTemperature + entry.BrightnessAdjustment*float32(brightnessValue))
		}
	}

	return 0
}

func (hcl *HCL) WaitTimeAt(index int) time.Duration {
	date := hcl.cfg.StartDate()
	for i, entry := range hcl.cfg.Curve.Entries {
		date = date.Add(time.Duration(entry.TimeOffset) * time.Millisecond)
		if i == index {
			break
		}
	}

	return date.Sub(time.Now())
}

func (hcl *HCL) Stop() {
	hcl.mu.Lock()
	defer hcl.mu.Unlock()
	if !hcl.stopped {
		hcl.stop <- struct{}{}
		hcl.stopped = true
	}
}

func (hcl *HCL) UpdateColorTemperatureForCurrentIndex() {
	hcl.mu.Lock()
	i := hcl.i
	hcl.mu.Unlock()

	hcl.UpdateColorTemperatureForIndex(i)
}

func (hcl *HCL) UpdateColorTemperatureForIndex(i int) {
	newColorTemperatur := hcl.ColorTemperatureAt(i)
	log.Info.Printf("hcl: color temperature is now %d Mired\n", newColorTemperatur)
	hcl.colorTemp.UpdateValueFromConnection(newColorTemperatur, hcl.conn)
}

func (hcl *HCL) Schedule() {
	hcl.mu.Lock()
	hcl.i = 0
	hcl.mu.Unlock()

	date := hcl.cfg.StartDate()
	log.Info.Printf("hcl: starting at %s\n", date)

	diff := date.Sub(time.Now())
	if diff > 0 {
		// Wait until start time starts
		<-time.After(diff)
	}

	hcl.activeTransitionCount.SetValue(1)
	defer hcl.activeTransitionCount.SetValue(0)

	for {
		hcl.mu.Lock()
		i := hcl.i
		hcl.mu.Unlock()
		if i >= len(hcl.cfg.Curve.Entries) {
			log.Info.Println("hcl: end of curve reached")
			<-hcl.stop
			log.Info.Println("hcl: schedule ended")
			return
		}

		waitTime := hcl.WaitTimeAt(i)
		log.Info.Printf("hcl: waiting until %v\n", time.Now().Add(waitTime))

		select {
		case <-time.After(waitTime):
			hcl.UpdateColorTemperatureForIndex(i)

		case <-hcl.stop:
			log.Info.Println("hcl: schedule stopped")
			return
		}
		hcl.mu.Lock()
		hcl.i++
		hcl.mu.Unlock()
	}
}

func setTLV8Payload(c *characteristic.Bytes, v interface{}) {
	if tlv8, err := tlv8.Marshal(v); err == nil {
		c.SetValue(tlv8)
	} else {
		log.Debug.Fatal(err)
	}
}
