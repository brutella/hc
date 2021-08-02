package hcl

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/tlv8"

	"encoding/base64"
	"log"
	"testing"
	"time"
)

func TestTransitionControl(t *testing.T) {
	str := "Av8B/wEBDAImARDE8z2pY8lMD4uJbDajPu+eAgjeCGQplwAAAAMIc9glI1nZPR0DAQEF/wEPAQQXbMG/AgQ5DrhDAwEAAAABEgEERETEvwIEq6q3QwMESNYEAAAAARIBBM3MzL8CBAAAtEMDBEB3GwAAAAESAQQ/6dO/AgQcx7FDAwRAdxsAAAABEgEEbMHWvwIEjmOwQwMEQHcbAAAAARIBBIMt2L8CBMdxr0MDBEB3GwAAAAESAQSDLdi/AgTH8a5DAwRAdxsAAAABGAEEbMHWvwIEjuOuQwMEQHcbAAQEgO42AAAAARIBBIMt2L8CBMfxrkMDBEB3GwAAAAEYAQQC/z/pAf/TvwIEHEevQwMEQHcbAAQEgO42AAAAARIBBOQ4zr8CBDkOsEMDBEB3GwAAAAESAQQF/3Icx78CBBzHsUMDBEB3GwAAAAESAQS8u7u/AgRVVbNDAwRAdxsAAAABEgEEBluwvwIEjmO2QwMEQHcbAAAAARIBBFD6pL8CBMfxuUMDBEB3GwAAAAESAQTe3Z2/AgSrKr9DAwRAdxsAAAABEgEE9UmfvwIE5LjGQwMEQHcbAAAAARIBBH3Sp78CBDkO0EMDBEB3GwAAAAESAQSlT7q/AgQcx9tDAwRAdxsAAAABEgEEERHRvwIEq6rnQwMEQHcbAAAAARIBBKuq6gL/vwIEqwH/qvNDAwRAdxsAAAABEgEEAAAAwAIEAAD+QwMEQHcbAAAAARIBBFuwBcACBOQ4AkQDBEAF/3cbAAAAARIBBCIiAsACBFUVAkQDBEB3GwAAAAESAQTSJ/2/AgTH8QFEAwRAdxsAAAABEgEEBlvwvwIEx/EARAMEQHcbAAAAARIBBNiC7b8CBI6jAUQDBEB3GwAAAAEYAQR90ue/AgQcBwBEAwRAdxsABASAy6QAAAABEgEEC7bgvwIEHMf7QwMEQHcbAAAAARIBBFVV1b8CBFVV9UMDBEB3GwAAAAESAQSf9Mm/AgSOY+5DAwRAdxsAAAABEgEE0ie9vwIEAtmOY+dDAwQB0UB3GwAAAAESAQTYgq2/AgQcx99DAwRAdxsAAAABEgEE9UmfvwIE5LjYQwMEQHcbAAAAARIFkQEEP+mTvwIEHMfRQwMEQHcbAAAAARIBBBERkb8CBKuqy0MDBEB3GwAAAAESAQSDLZi/AgTH8cVDAwRAdxsAAAABEgEEZmamvwIEAIDAQwMEQHcbAAAAARIBBGELtr8CBHKcu0MDBEB3GwAAAAESAQQXbMG/AgQ5DrhDAwT4oBYAAgELAwwBBAoAAAACBGQAAAAGAmDqCATAJwkA"

	buf, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%x\n", buf)

	var ctrl TransitionControl
	if err := tlv8.Unmarshal(buf, &ctrl); err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v\n", ctrl)

	if is, want := ctrl.Read.ColorTemperatureIID, uint8(0); is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := ctrl.Update.Configuration.ColorTemperatureIID, uint8(12); is != want {
		t.Fatalf("%v != %v", is, want)
	}

	config := ctrl.Update.Configuration
	st := config.StartDate()
	log.Println("Start Time:", st)
	for _, entry := range config.Curve.Entries {
		st = st.Add(time.Duration(entry.TimeOffset) * time.Millisecond)
		kelvin := 1_000_000 / (entry.ColorTemperature + 100*entry.BrightnessAdjustment)
		log.Printf("%s: %.2f Mired %.4f %% → %.0f Kelvin\n", st, entry.ColorTemperature, entry.BrightnessAdjustment, kelvin)
	}
}

func TestHCL(t *testing.T) {
	brightness := characteristic.NewBrightness()
	brightness.ID = 1
	colorTemp := characteristic.NewColorTemperature()
	colorTemp.ID = 2
	count := NewActiveTransitionCountCharacteristic()

	id := []byte{1, 2, 3, 4, 5, 6}

	timeoffset := uint32(100) // 100msec

	cfg := TransitionConfiguration{
		ColorTemperatureIID: uint8(colorTemp.ID),
		Params: TransitionParams{
			TransitionID: id,
			StartTime:    Timestamp(time.Now()),
		},
		Enabled: true,
		Curve: TransitionCurve{
			Entries: []TransitionCurveEntry{
				TransitionCurveEntry{
					BrightnessAdjustment: 0,
					ColorTemperature:     500,
					TimeOffset:           0,
					Duration:             0,
				},
				TransitionCurveEntry{
					BrightnessAdjustment: -1,
					ColorTemperature:     450,
					TimeOffset:           timeoffset,
					Duration:             0,
				},
				TransitionCurveEntry{
					BrightnessAdjustment: -1.3,
					ColorTemperature:     300,
					TimeOffset:           timeoffset,
					Duration:             0,
				},
			},
			BrightnessIID: uint8(brightness.ID),
			ValueRange: TransitionValueRange{
				Min: 10,
				Max: 100,
			},
		},
		UpdateInterval: 1000, // every second
		NotifyInterval: 1000,
	}
	hcl := NewHCL(cfg, brightness, colorTemp, count)

	brightness.SetValue(50)

	ch := make(chan int, 0)
	colorTemp.OnValueRemoteUpdate(func(v int) {
		ch <- v
	})

	done := make(chan struct{}, 0)
	go func() {
		hcl.Schedule()
		done <- struct{}{}
	}()

	for _, expected := range []int{500, 450 - 1*50, 300 - 1.3*50} {
		select {
		case v := <-ch:
			if is, want := v, expected; is != want {
				t.Fatalf("%v != %v", is, want)
			}
		case <-time.After(1 * time.Second):
			t.Fatal("color temperature transition expected")
		}
	}

	select {
	case <-ch:
		t.Fatal("no more color temperature changes expected")
	case <-time.After(200 * time.Millisecond):
		break
	}

	hcl.Stop()
	<-done
}
