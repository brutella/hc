# GOHAP

This is an implementation of the **H**omeKit **A**ccessory **P**rotocol (HAP) to create HomeKit bridges for external accessories.

## Accessories

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

## Components

The library consists several packages

- `app` to get the bridge up and running
- `common` basic structs
- `crypto` contains helper methods for cryptography
- `db` for persistent storage
- `model` interfaces for HomeKit accessory types
    - `accessory` implementation of accessories
    - `service` implementation of services
    - `characteristic` implementation of characteristics
- `netio` contains the HTTP endpoint handlers
    - `controller`
- `server` contains the server implementation
- `example` contains example implementation of a HomeKit bridge and client

## Dependencies

Gohap depends on the following libraries

- `github.com/stretchr/testify` to get asserts in unit tests
- `github.com/tadglines/go-pkgs/crypto/srp` for *SRP* algorithm
- `github.com/codahale/chacha20` for *chacha20 poly1305* algorithm
- `github.com/tonnerre/golang-go.crypto/poly1305`for *chacha20 poly1305* algorithm
- `github.com/tonnerre/golang-go.crypto/curve25519` for *curve25519* key generation
- `github.com/tonnerre/golang-go.crypto/hkdf`
- `github.com/agl/ed25519` for *ed25519* signature
- `github.com/gosexy/to` for type conversion

## Installation

**TODO**

## Usage

The following code shows a minimal implementation of a Gohap bridge

	conf := app.NewConfig()
	
	// Path to folder where data is stored
    conf.DatabaseDir = "./data"
        
    // Creates a new app
    app, err := app.NewApp(conf)
    if err != nil {
        log.Fatal(err)
    }
    
    // Runs the app
    app.Run()

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

The following example adds a switch accessory to the bridge

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

**TODO: Update the bonjour txt records for the bridge when characteristics are changed.**

## TODOs

- Create interfaces for models to hide implementation details *WIP*
- On server close, wait until connections are closed
- Check invalid service and characteristics (e.g. temperature must not be smaller than min and bigger than max)
- Do not allow value changes of read-only characteristics

- Rename `on` characteristic to sth like `power state`
- Add fan, garage door opener, lock management and mechanism accessory types
- Put vendor packages into vendor dir
- Create pull requests for vendor package changes

## IO with Virtual Devices

**THIS IDEA IS ABANDONED**

On linux, IO is done via files (e.g. [GPIO](https://developer.ridgerun.com/wiki/index.php/How_to_use_GPIO_signals) input is written to a file). Files can read and write to those files.

The HomeKit brige is a gateway to the HomeKit universe. It could act like a daemon which watches for accessory changes and delivers that to the HomeKit clients. An acccesory can offer services which are specified by the *HAP*.

- Info service: info (name, manufacturer, ...) about accessory, every bridge must have it too
- Thermostat service: read and write to temperature
- Switch service: dis-/enable a switch
- ...

Those services are made of characteristics like

- name
- current temperature
- target temperature
- on
- ...

Characteristics can be read and write by clients, and the bridge is the mediator to the accessory. The interaction between accessory and bridge should be based on files, very similar to virtual devices on linux. The accessory and bridge watches the file system. This allows multiple accessories to use the same bridge as a gateway.

There must be a contract between accessory and bridge on the file system representation.

**Thermostat**

| file | value |
| ---- | ----- |
| type | thermostat |
| serial-number| string |
| name | string |
| manufacturer| string |
| model| string |
| unit | celsius |
| temperature| float |
| target-temperature| float |
| mode| off | heating | cooling |
| target-mode| off | heating | cooling |
| humidity| float |
| target-humidity | float |

**Switch**

| file | value |
| ---- | ----- |
| type | switch |
| serial-number| string |
| name | string |
| manufacturer| string |
| model| string |
| on | 0 | 1 |

The files must be located in a specific folder. On startup, the bridge scans the folder and create a model representation. On file changes, the bridge re-initialized the model and notifies the client for new values by updating the Bonjour TXT record *s#*.

The accessory 