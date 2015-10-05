# HomeControl

[HomeControl][homecontrol] is an implementation of the [HomeKit][homekit] Accessory Protocol (HAP) to create your own HomeKit accessory and bridges. HomeKit bridges make non-HomeKit accessories available to HomeKit by acting as a middleman.

## Overview

[HomeKit][homekit] is a set of protocols and libraries to access accessories for Home Automation. Unfortunately the protocol is not open source and the official documentation is only available to MFi members. HomeControl is a complete implementation of the protocol in Go and does not depend on any OS.

## HomeKit Client

I've made an app for iPhone, iPad and Apple Watch called [Home][home] to control any HomeKit accessory. If you purchase Home on the [App Store][home-appstore], you not only support my work but also get an awesome iOS app. Thank you.

[home]: http://selfcoded.com/home/
[home-appstore]: http://itunes.apple.com/app/id995994352

## Features

- Full implementation of the HomeKit Accessory Protocol in pure Go
    - Support for switch, outlet, light bulb, thermostat accessory
- Built-in service announcement via mDNS using [bonjour](http://github.com/oleksandr/bonjour)
- Optional logging with https://github.com/brutella/log
- Runs on multiple platforms (already in use on Linux and OS X)
- Documentation: http://godoc.org/github.com/brutella/hc

## Getting Started

1. [Install Go](http://golang.org/doc/install)
2. [Setup Go workspace](http://golang.org/doc/code.html#Organization)
3. Create your own HomeKit bridge or clone an existing one (e.g.  [hklight](https://github.com/brutella/hklight))

        cd $GOPATH/src
        
        # Clone project
        git clone https://github.com/brutella/hklight && cd hklight
        
        # Install dependencies
        go get
        
        # Run the project
        go run hklightd.go

4. Pair with your HomeKit App of choice (e.g. [Home][home-appstore])

## API Example

Create a simple on/off switch which is accessible via IP and secured using the pin *00102003*.

```go
package main

import (
    "log"
    "github.com/brutella/hc/hap"
    "github.com/brutella/hc/model"
    "github.com/brutella/hc/model/accessory"
)

func main() {
	info := model.Info{
		Name: "Lamp",
	}
	sw := accessory.NewSwitch(info)
    
	t, err := hap.NewIPTransport("00102003", sw.Accessory)
	if err != nil {
		log.Fatal(err)
	}
    
    hap.OnTermination(func(){
        t.Stop()
    })
    
	t.Start()
}
```

You should change some default values for your own needs

```go
info := model.Info{
    Name: "Lamp",
    SerialNumber: "051AC-23AAM1",
	Manufacturer: "Apple",
    Model: "AB",
    Firmware: "1.0.1",
}
```

### Callbacks

You get a callback when the power state of a switch changed by a client.

```go
sw.OnStateChanged(func(on bool) {
	if on == true {
		log.Println("Client changed switch to on")
	} else {
		log.Println("Client changed switch to off")
	}
})
```

When the switch is turned on "the analog way", you should set the state of the accessory.

	sw.SetOn(true)

A complete example is available in `_example/example.go`.

## Dependencies

HomeControl depends on the following libraries

- `github.com/tadglines/go-pkgs/crypto/srp` for *SRP* algorithm
- `github.com/codahale/chacha20` for *chacha20 poly1305* algorithm
- `github.com/golang/crypto`for *chacha20 poly1305* algorithm and *curve25519* key generation
- `github.com/agl/ed25519` for *ed25519* signature
- `github.com/gosexy/to` for type conversion
- `github.com/oleksandr/bonjour` for mDNS

## HomeKit Accessories

HomeControl currently supports the following accessory types

- Switch
- Outlet
- Light Bulb
- Thermostat
- Thermometer (same as the Thermostat accessory which just readonly services)

The metdata dump in iOS 8.3 (found by [@KhaosT](https://twitter.com/khaost/status/567621750494474241)) includes a list of required and optional characteristics.

<table>
    <tr><th>Service</th><th>Required</th><th>Optional</th><tr>
    <tr><td>Accessory Information</td><td>name, manufacturer, model, serial-number, identify</td><td>firmware.revision, hardware.revision, software.revision</td><tr>
    <tr><td>Switch</td><td>on</td><td>name</td><tr>
    <tr><td>Outlet</td><td>on, outlet-in-use</td><td>name</td><tr>
    <tr><td>Fan</td><td>on</td><td>name, rotation.direction, rotation.speed</td><tr>
    <tr><td>Thermostat</td><td>heating-cooling.current, heating-cooling.target, temperature.current, temperature.target, temperature.units</td><td>name, relative-humidity.current, relative-humidity.target, temperature.cooling-threshold, temperature.heating-threshold</td><tr>
    <tr><td>Garage Door Opener</td><td>door-state.current, door-state.target, obstruction-detected</td><td>lock-mechanism.current-state, lock-mechanism.target-state, name</td><tr>
    <tr><td>Light Bulb</td><td>on</td><td>name, brightness, hue, saturation</td><tr>
    <tr><td>Lock Management</td><td>version, lock-management.control-point</td><td>administrator-only-access, audio-feedback, door-state.current, lock-management.auto-secure-timeout, lock-mechanism.last-known-action, logs, motion-detected</td><tr>
    <tr><td>Lock Mechanism</td><td>lock-mechanism.current-state, lock-mechanism.target-state</td><td>name</td><tr>
</table>

The HomeKit framework on iOS uses the same order as in the json. I assume that clients displays them in the same order to the user.

### iOS 9

iOS 9 supports new type of accessories and includes new service and characteristic types. The new types are already available in *model/service/constants.go* and *model/characteristic/constants.go*.

# Contact

Matthias Hochgatterer

Github: [https://github.com/brutella](https://github.com/brutella/)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)


# License

HomeControl is available under a non-commercial license. See the LICENSE file for more info.

[homecontrol]: http://selfcoded.com/homecontrol/
[homekit]: https://developer.apple.com/homekit/