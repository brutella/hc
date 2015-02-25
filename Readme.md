# HomeControl

*HomeControl* is an implementation of the HomeKit Accessory Protocol (HAP) to create HomeKit bridges. A HomeKit bridge makes non-HomeKit accessories available to HomeKit by acting as a middleman.

### NOTICE

HomeKit is not ready for prime time yet. The library is still subject to change because I cannot guarantee that everything works with future iOS versions.

## Overview

HomeKit is a set of protocols and libraries to access accessories used for Home Automation. Unfortunately the protocol is not open source and the official documentation is only available to MFi members. HomeControl is a complete implementation of the HAP in Go and does not depend on any OS.

Read the API documentation here: http://godoc.org/github.com/brutella/hc

## Features

- Complete implementation of HAP (only some accessory types are missing)
- Very easy to get started
- Optional logging with https://github.com/brutella/log
- Runs on multiple platforms (already in use on Linux and OS X)

## Example

Here is a minimal implementation of a HomeKit bridge

    package main

    import (
      "log"
      "github.com/brutella/hc/hap"
    )

    func main() {
        conf := hap.NewConfig()

        // Path to folder where data is stored
        conf.DatabaseDir = "./data"

        // Create an app
        app, err := hap.NewApp(conf)
        if err != nil {
            log.Fatal(err)
        }

        // Run it
        app.Run()
    }

You should change some default values for your own needs

    // Name of the bridge which appears in the accessory browser on iOS; default "GoBridge"
    conf.BridgeName = "TestBridge"

    // Password the user has to enter when adding the accessory to HomeKit
    // Default "00102003"
    pwd, _ := server.NewPassword("11122333")
    conf.BridgePassword = pwd 
    
    // Bridge manufacturer name
    conf.BridgeManufacturer = "Apple Inc."


### Add Accessories

Now lets add a switch accessory which can be switched on and off.

    import "github.com/brutella/hc/model/accessory"
    
	info := accessory.Info{
        Name: "My Switch",
        SerialNumber: "001",
        Manufacturer: "Google",
        Model: "Switchy",
    }
    
    sw := accessory.NewSwitch(info)    
    app.AddAccessory(sw.Accessory)

You can use the `OnStateChanged` callback to get notified, when a HomeKit client (iOS device) changes the value of the `on` characteristic.

	sw.OnStateChanged(func(on bool) {
        if on == true {
            log.Println("Switch on")
        } else {
            log.Println("Switch off")
        }
    })

When the `on` characteristic is changed by the accessory itself e.g. when the switch is turned on "the analog way", you should notify the clients by manually setting the state.

	sw.SetOn(true)

## Dependencies

HomeControl depends on the following libraries

- `github.com/stretchr/testify` to get asserts in unit tests
- `github.com/tadglines/go-pkgs/crypto/srp` for *SRP* algorithm
- `github.com/codahale/chacha20` for *chacha20 poly1305* algorithm
- `github.com/tonnerre/golang-go.crypto/poly1305`for *chacha20 poly1305* algorithm
- `github.com/tonnerre/golang-go.crypto/curve25519` for *curve25519* key generation
- `github.com/tonnerre/golang-go.crypto/hkdf`
- `github.com/agl/ed25519` for *ed25519* signature
- `github.com/gosexy/to` for type conversion
- `github.com/oleksandr/bonjour` for mDNS

## TODOs

- Create interfaces for models to hide implementation details *WIP*
- On server close, wait until connections are closed
- Check invalid service and characteristics (e.g. temperature must not be smaller than min and bigger than max)
- Do not allow value changes of read-only characteristics
- Add fan, garage door opener, lock management and mechanism accessory types
- Put vendor packages into vendor dir

## HomeKit Accessories

The library supports the following accessory types

- Switch
- Outlet
- Light Bulb
- Thermostat
- Thermometer (same as the Thermostat accessory which just readonly services)

The metdata dump in iOS 8.3 (found by [@KhaosT](https://twitter.com/khaost/status/567621750494474241)) includes a list of required and optional characteristics.

<table>
    <tr><th>Accessory</th><th>Required</th><th>Optional</th><tr>
    <tr><td>Accessory Information</td><td>name, manufacturer, model, serial-number, identify</td><td>firmware.revision, hardware.revision, software.revision</td><tr>
    <tr><td>Switch</td><td>on</td><td>name</td><tr>
    <tr><td>Outlet</td><td>on, outlet-in-use</td><td>name</td><tr>
    <tr><td>Fan</td><td>on</td><td>name, rotatino.direction, rotation.speed</td><tr>
    <tr><td>Thermostat</td><td>heating-cooling.current, heating-cooling.target, temperature.current, temperature.target, temperature.units</td><td>name, relative-humidity.current, relative-humidity.target, temperature.cooling-threshold, temperature.heating-threshold</td><tr>
    <tr><td>Garage Door Opener</td><td>door-state.current, door-state.target, obstruction-detected</td><td>lock-mechanism.current-state, lock-mechanism.target-state, name</td><tr>
    <tr><td>Light Bulb</td><td>on</td><td>name, brightness, hue, saturation</td><tr>
    <tr><td>Lock Management</td><td>version, lock-management.control-point</td><td>administrator-only-access, audio-feedback, door-state.current, lock-management.auto-secure-timeout, lock-mechanism.last-known-action, logs, motion-detected</td><tr>
    <tr><td>Lock Mechanism</td><td>lock-mechanism.current-state, lock-mechanism.target-state</td><td>name</td><tr>
</table>