# HomeControl

[![Build Status](https://travis-ci.org/brutella/hc.svg)](https://travis-ci.org/brutella/hc)

[HomeControl][homecontrol] is an implementation of the [HomeKit][homekit] Accessory Protocol (HAP) to create your own HomeKit accessory in [Go](https://golang.org). [HomeKit][homekit] is a set of protocols and libraries to access devices for Home Automation. The actual protocol documentation is only available to MFi members.

You can use this library to make existing Home Automation devices HomeKit compatible. I've already developed the following HomeKit bridges with in:

- [LIFX](https://github.com/brutella/hklifx/)
- [UVR1611](https://github.com/brutella/hkuvr)
- [Fronius Symo](https://github.com/brutella/hksymo)

## HomeKit on iOS

HomeKit is fully integrated since iOS 8. Developers can use the HomeKit framework to communicate with HomeKit using high-level APIs.
I've developed the [Home][home] app (for iPhone, iPad, Apple Watch) to control HomeKit accessories. If you [purchase Home][home-appstore] on the App Store, you not only support my work but also get an awesome iOS app. Thank you.

Once you've setup HomeKit, you can use Siri to interact with your accessories using voice command (*Hey Siri, turn off the lights in the living room*).

[home]: http://selfcoded.com/home/
[home-appstore]: http://itunes.apple.com/app/id995994352

## Features

- Full implementation of the HAP in Go
- Built-in service announcement via mDNS using [bonjour](http://github.com/oleksandr/bonjour)
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

Create a simple on/off switch, which is accessible via IP and secured using the pin *00102003*.

```go
package main

import (
    "log"
    "github.com/brutella/hc"
    "github.com/brutella/hc/accessory"
)

func main() {
	info := accessory.Info{
		Name: "Lamp",
	}
	acc := accessory.NewSwitch(info)
    
    config := hc.Config{Pin: "00102003"}
	t, err := hc.NewIPTransport(config, acc.Accessory)
	if err != nil {
		log.Panic(err)
	}
    
    hc.OnTermination(func(){
        t.Stop()
    })
    
	t.Start()
}
```

You should change some default values for your own needs

```go
info := accessory.Info{
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
acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
	if on == true {
		log.Println("Client changed switch to on")
	} else {
		log.Println("Client changed switch to off")
	}
})
```

When the switch is turned on "the analog way", you should set the state of the accessory.

	acc.Switch.On.SetValue(true)

A complete example is available in `_example/example.go`.

## Model

The HomeKit model hierarchy looks like this:

    Accessory
    |-- Accessory Info Service
    |   |-- Identify Characteristic
    |   |-- Manufacturer Characteristic
    |   |-- Model Characteristic
    |   |-- Name Characteristic
    |   |-- Serial Characteristic
    |   
    |-- * Service
    |   |-- * Characteristic

HomeKit accessories are container for services. Every accessory must provide the `Accessory Information Service`. Every service provides one or more characteristics (a characteristic might be the power state of an outlet). HomeKit has predefined service and characteristic types, which are supported by iOS. You can define your own service and characteristic types, but it's recommended to use predefined ones.

This library provides all HomeKit characteristics (see `characteristic` package) and services (see `service` package).
You can also find common accessory types like lightbulbs, outlets, thermostats in the `accessory` package.

## Dependencies

HomeControl uses vendor directories (`vendor/`) to integrate the following libraries

- `github.com/tadglines/go-pkgs/crypto/srp` for *SRP* algorithm
- `github.com/codahale/chacha20` for *chacha20 poly1305* algorithm
- `github.com/agl/ed25519` for *ed25519* signature
- `github.com/gosexy/to` for type conversion
- `github.com/oleksandr/bonjour` for mDNS

# Contact

Matthias Hochgatterer

Github: [https://github.com/brutella](https://github.com/brutella/)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)


# License

HomeControl is available under a non-commercial license. See the LICENSE file for more info.

[homecontrol]: http://selfcoded.com/homecontrol/
[homekit]: https://developer.apple.com/homekit/