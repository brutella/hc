# HomeControl

*HomeControl* is an implementation of the HomeKit Accessory Protocol (HAP) to create HomeKit accessories. *HomeControl* supports HomeKit bridges, which makes non-HomeKit accessories available to HomeKit by acting as a middleman.

### NOTICE

The HomeControl API is still subject to change because HomeKit is not ready for prime time.

## Overview

HomeKit is a set of protocols and libraries to access accessories used for Home Automation. Unfortunately the protocol is not open source and the official documentation is only available to MFi members. HomeControl is a complete implementation of the HAP in Go and does not depend on any OS.

Read the API documentation here: http://godoc.org/github.com/brutella/hc

## Features

- Complete implementation of HAP (only some accessory types are missing)
- Optional logging with https://github.com/brutella/log
- Runs on multiple platforms (already in use on Linux and OS X)

## Example

Create a simple on/off switch which is accessible via IP and secured using the password *00102003*.

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
    Firwmare: "1.0.1",
}
```

### Callbacks

When the power state is changed by a client, you get a callback.

```go
sw.OnStateChanged(func(on bool) {
	if on == true {
		log.Println("Client changed switch to on")
	} else {
		log.Println("Client changed switch to off")
	}
})
```

When the switch is turned on "the analog way", you should notify the clients by manually setting the state.

	sw.SetOn(true)

## Dependencies

HomeControl depends on the following libraries

- `github.com/stretchr/testify` to get asserts in unit tests
- `github.com/tadglines/go-pkgs/crypto/srp` for *SRP* algorithm
- `github.com/codahale/chacha20` for *chacha20 poly1305* algorithm
- `github.com/golang/crypto`for *chacha20 poly1305* algorithm and *curve25519* key generation
- `github.com/agl/ed25519` for *ed25519* signature
- `github.com/gosexy/to` for type conversion
- `github.com/oleksandr/bonjour` for mDNS

## TODOs

- Better random uuid

## HomeKit Accessories

The library supports the following accessory types

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

# Contact

Matthias Hochgatterer

Github: [https://github.com/brutella](https://github.com/brutella/)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)


# License

HomeControl is available under a non-commercial license. See the LICENSE file for more info.