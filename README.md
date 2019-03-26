# HomeControl

[![Build Status](https://travis-ci.org/brutella/hc.svg)](https://travis-ci.org/brutella/hc)

`hc` is a lightweight framework to develop HomeKit accessories in Go. 
It abstracts the **H**omeKit **A**ccessory **P**rotocol and makes it easy to work with [services](service/README.md) and [characteristics](characteristic/README.md).

`hc` handles the underlying communication between HomeKit accessories and clients.
You can focus on implementing the business logic for your accessory, without having to worry about the underlying protocol.

I've already developed the following HomeKit bridges with in:

- [LIFX](https://github.com/brutella/hklifx/)
- [UVR1611](https://github.com/brutella/hkuvr)
- [Fronius Symo](https://github.com/brutella/hksymo)

**What is HomeKit?**

[HomeKit][homekit] is a set of protocols and libraries to communicate with smart home appliances. ~~The actual protocol documentation is only available to MFi members.~~ A non-commercial version of the documentation is now available on the [HomeKit developer website](https://developer.apple.com/homekit/).

**iOS**


HomeKit is fully integrated into iOS since iOS 8. Developers can use [HomeKit.framework](https://developer.apple.com/documentation/homekit) to communicate with accessories using high-level APIs.

[[https://github.com/brutella/hc/blob/master/_img/home-icon.png|alt=Home.app]]

I've developed the [Home][home] app to control HomeKit accessories from iPhone, iPad, and Apple Watch.
If you would like to support `hc`, please purchase Home from the [App Store](home-appstore). That would be awesome. ❤️

Once you've setup HomeKit on iOS, you can use Siri to interact with your accessories using voice command (*Hey Siri, turn off the lights in the living room*).

[home]: https://hochgatterer.me/home/
[home-appstore]: http://itunes.apple.com/app/id995994352

## Features

- Full implementation of the HAP in Go
- Supports all HomeKit [services and characteristics](service/README.md)
- Built-in service announcement via DNS-SD using [dnssd](http://github.com/brutella/dnssd)
- Runs on linux and macOS
- Documentation: http://godoc.org/github.com/brutella/hc

## Getting Started

1. [Install](http://golang.org/doc/install) and [set up](http://golang.org/doc/code.html#Organization) Go
2. Create your own HomeKit accessory or clone an existing one (e.g.  [hklight](https://github.com/brutella/hklight))

        cd $GOPATH/src
        
        # Clone project
        git clone https://github.com/brutella/hklight && cd hklight
        
        # Run the project
        make run

3. Pair with your HomeKit App of choice (e.g. [Home][home-appstore])

**Go Modules**

`hc` supports [Go module](https://github.com/golang/go/wiki/Modules) since `v1.0.0`.
Make sure to set the environment variable `GO111MODULE=on`.

## Example

See [_example](_example) for a variety of examples.

**Basic switch accessory**

Create a simple on/off switch, which is accessible via IP and secured using the pin *00102003*.

```go
package main

import (
    "log"
    "github.com/brutella/hc"
    "github.com/brutella/hc/accessory"
)

func main() {
    // create an accessory
	info := accessory.Info{Name: "Lamp"}
	ac := accessory.NewSwitch(info)
    
    // configure the ip transport
    config := hc.Config{Pin: "00102003"}
	t, err := hc.NewIPTransport(config, ac.Accessory)
	if err != nil {
		log.Panic(err)
	}
    
    hc.OnTermination(func(){
        <-t.Stop()
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

### Events

When a connected client changes the value of characteristics, you get a callback.
The following example shows how to get notified when the [On](characteristic/on.go) characteristic value changes.

```go
ac.Switch.On.OnValueRemoteUpdate(func(on bool) {
	if on == true {
		log.Println("Client changed switch to on")
	} else {
		log.Println("Client changed switch to off")
	}
})
```

When the switch is turned on "the analog way", you should set the state of the accessory.

```go
ac.Switch.On.SetValue(true)
```

## Accessory Architecture

HomeKit uses a hierarchical architecture for define accessories, services and characeristics.
At the root level there is an accessory.
Every accessory contains services.
And every service contains characteristics.

For example a [lightbulb accessory](accessory/lightbulb.go) contains a [lightbulb service](service/lightbulb.go).
This service contains characteristics like [on](characteristic/on.go) and [brightness](characteristic/brightness.go).

There are predefined accessories, services and characteristics available in HomeKit.
Those types are defined in the packages [accessory](accessory), [service](service), [characteristic](characteristic).

# Contact

Matthias Hochgatterer

Website: [http://hochgatterer.me](https://hochgatterer.me)

Github: [https://github.com/brutella](https://github.com/brutella/)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)


# License

`hc` is available under the Apache License 2.0 license. See the LICENSE file for more info.

[homekit]: https://developer.apple.com/homekit/
