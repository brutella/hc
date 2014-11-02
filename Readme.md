# GOHAP

This is an implementation of the **H**omeKit **A**ccessory **P**rotocol (HAP), which enabled you to write HomeKit bridges for external accessories.

Gohap currently supports the following accessory types

- Switch
- Light Bulb
- Thermostat (custom `Thermometer` accessory which has readonly services)

## Components

The library consists several packages

- `model` defines the accessory types
- `app` to get the bridge up and running
- `server` contains the server implementation
- `netio` contains the HTTP endpoint handlers
- `crypto` contains helper methods for cryptography
- `common` some common structs
- `db` contains classes to read and write data
- `example` contains example implementation of a HomeKit bridge and client

## Dependencies

Gohap depends on the following libraries

- `github.com/stretchr/testify` to get asserts in unit tests
- `github.com/tadglines/go-pkgs/crypto/srp` for *SRP* algorithm
- `github.com/codahale/chacha20` for *chacha20 poly1305* algorithm
- `github.com/tonnerre/golang-go.crypto/poly1305`for *chacha20 poly1305* algorithm
- `github.com/tonnerre/golang-go.crypto/curve25519` for *curve25519* key generation
- `github.com/agl/ed25519` for *ed25519* signature
- `github.com/tonnerre/golang-go.crypto/hkdf`
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

	info := model.Info{
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
- Add test for new methods
- Check invalid service and characteristics (e.g. temperature must not be smaller than min and bigger than max)
- Check invalid request (aid or iid not found)
- Do not allow value changes of read-only characteristics

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