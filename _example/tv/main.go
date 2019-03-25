package main

import (
	"fmt"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/service"
)

func addInputSource(t *accessory.Television, id int, name string, inputSourceType int) {
	in := service.NewInputSource()

	in.Identifier.SetValue(id)
	in.ConfiguredName.SetValue(name)
	in.Name.SetValue(name)
	in.InputSourceType.SetValue(inputSourceType)
	in.IsConfigured.SetValue(characteristic.IsConfiguredConfigured)

	t.AddService(in.Service)
	t.Television.AddLinkedService(in.Service)

	in.ConfiguredName.OnValueRemoteUpdate(func(str string) {
		fmt.Printf(" %s configured name => %s\n", name, str)
	})
	in.InputSourceType.OnValueRemoteUpdate(func(v int) {
		fmt.Printf(" %s source type => %v\n", name, v)
	})
	in.IsConfigured.OnValueRemoteUpdate(func(v int) {
		fmt.Printf(" %s configured => %v\n", name, v)
	})
	in.CurrentVisibilityState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf(" %s current visibility => %v\n", name, v)
	})
	in.Identifier.OnValueRemoteUpdate(func(v int) {
		fmt.Printf(" %s identifier => %v\n", name, v)
	})
	in.InputDeviceType.OnValueRemoteUpdate(func(v int) {
		fmt.Printf(" %s device type => %v\n", name, v)
	})
	in.TargetVisibilityState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf(" %s target visibility => %v\n", name, v)
	})
	in.Name.OnValueRemoteUpdate(func(str string) {
		fmt.Printf(" %s name => %s\n", name, str)
	})
}

// TODO
// - Refactoring how to store characteristic values
func main() {
	log.Debug.Enable()

	info := accessory.Info{
		Name: "Television",
	}
	acc := accessory.NewTelevision(info)

	acc.Television.Active.SetValue(characteristic.ActiveActive)
	acc.Television.SleepDiscoveryMode.SetValue(characteristic.SleepDiscoveryModeAlwaysDiscoverable)
	acc.Television.ActiveIdentifier.SetValue(1)
	acc.Television.CurrentMediaState.SetValue(characteristic.CurrentMediaStatePause)
	acc.Television.TargetMediaState.SetValue(characteristic.TargetMediaStatePause)

	acc.Television.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("active => %d\n", v)
	})

	acc.Television.ActiveIdentifier.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("active identifier => %d\n", v)
	})

	acc.Television.ConfiguredName.OnValueRemoteUpdate(func(v string) {
		fmt.Printf("configured name => %s\n", v)
	})
	acc.Television.SleepDiscoveryMode.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("sleep discovery mode => %d\n", v)
	})
	acc.Television.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("brightness => %d\n", v)
	})
	acc.Television.ClosedCaptions.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("closed captions => %d\n", v)
	})
	acc.Television.DisplayOrder.OnValueRemoteUpdate(func(v []byte) {
		fmt.Printf("display order => %v\n", v)
	})
	acc.Television.CurrentMediaState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("current media state => %d\n", v)
	})
	acc.Television.TargetMediaState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("target media state => %d\n", v)
	})

	acc.Television.PowerModeSelection.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("power mode selection => %d\n", v)
	})

	acc.Television.PictureMode.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("PictureMode => %d\n", v)
	})

	acc.Television.RemoteKey.OnValueRemoteUpdate(func(v int) {
		switch v {
		case characteristic.RemoteKeyRewind:
			fmt.Println("Rewind")
		case characteristic.RemoteKeyFastForward:
			fmt.Println("Fast forward")
		case characteristic.RemoteKeyExit:
			fmt.Println("Exit")
		case characteristic.RemoteKeyPlayPause:
			fmt.Println("Play/Pause")
		case characteristic.RemoteKeyInfo:
			fmt.Println("Info")
		case characteristic.RemoteKeyNextTrack:
			fmt.Println("Next")
		case characteristic.RemoteKeyPrevTrack:
			fmt.Println("Prev")
		case characteristic.RemoteKeyArrowUp:
			fmt.Println("Up")
		case characteristic.RemoteKeyArrowDown:
			fmt.Println("Down")
		case characteristic.RemoteKeyArrowLeft:
			fmt.Println("Left")
		case characteristic.RemoteKeyArrowRight:
			fmt.Println("Right")
		case characteristic.RemoteKeySelect:
			fmt.Println("Select")
		case characteristic.RemoteKeyBack:
			fmt.Println("Back")
		}
	})

	config := hc.Config{Pin: "12344321", StoragePath: "./db"}
	t, err := hc.NewIPTransport(config, acc.Accessory)

	addInputSource(acc, 1, "HDMI 1", characteristic.InputSourceTypeHdmi)
	addInputSource(acc, 2, "HDMI 2", characteristic.InputSourceTypeHdmi)
	addInputSource(acc, 3, "Netflix", characteristic.InputSourceTypeApplication)
	addInputSource(acc, 4, "YouTube", characteristic.InputSourceTypeApplication)
	addInputSource(acc, 5, "Twitter", characteristic.InputSourceTypeApplication)
	addInputSource(acc, 6, "Composite Video", characteristic.InputSourceTypeCompositeVideo)
	addInputSource(acc, 7, "S-Video", characteristic.InputSourceTypeSVideo)
	addInputSource(acc, 8, "Component Video", characteristic.InputSourceTypeComponentVideo)
	addInputSource(acc, 9, "Dvi", characteristic.InputSourceTypeDvi)
	addInputSource(acc, 10, "Airplay", characteristic.InputSourceTypeAirplay)
	addInputSource(acc, 11, "Usb", characteristic.InputSourceTypeUsb)

	if err != nil {
		log.Info.Panic(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
