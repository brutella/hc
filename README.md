# hc

[![GoDoc Widget]][GoDoc] [![Travis Widget]][Travis]

`hc` is a lightweight framework to develop HomeKit accessories in Go.
It abstracts the **H**omeKit **A**ccessory **P**rotocol (HAP) and makes it easy to work with [services](service/README.md) and [characteristics](characteristic/README.md).

`hc` handles the underlying communication between HomeKit accessories and clients.
You can focus on implementing the business logic for your accessory, without having to worry about the protocol.

Here are some projects which use `hc`.

- [hkcam](https://github.com/brutella/hkcam)
- [hklifx](https://github.com/brutella/hklifx/)
- [hkuvr](https://github.com/brutella/hkuvr)
- [hksymo](https://github.com/brutella/hksymo)

**What is HomeKit?**

[HomeKit][homekit] is a set of protocols and libraries from Apple. It is used by Apple's platforms to communicate with smart home appliances. A non-commercial version of the documentation is now available on the [HomeKit developer website](https://developer.apple.com/homekit/).

HomeKit is fully integrated into iOS since iOS 8. Developers can use [HomeKit.framework](https://developer.apple.com/documentation/homekit) to communicate with accessories using high-level APIs.

<img alt="Home+.app" src="_img/home-icon.png?raw=true" width="87" />

I've developed the [Home+][home+] app to control HomeKit accessories from iPhone, iPad, and Apple Watch.
If you want to support `hc`, please purchase Home from the [App Store][home-appstore]. That would be awesome. ❤️

Checkout the official [website][home+].

[home+]: https://hochgatterer.me/home/
[home-appstore]: http://itunes.apple.com/app/id995994352
[GoDoc]: https://godoc.org/github.com/brutella/hc
[GoDoc Widget]: https://godoc.org/github.com/brutella/hc?status.svg
[Travis]: https://travis-ci.org/brutella/hc
[Travis Widget]: https://travis-ci.org/brutella/hc.svg

## Features

- Supports Go modules (requires Go 1.13, or setting `GO111MODULE` to `on` when using Go 1.11/12)
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

You can define more specific accessory info, if you want.

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

The library provides callback functions, which let you know when a clients updates a characteristic value.
The following example shows how to get notified when the [On](characteristic/on.go) characteristic value changes.

```go
ac.Switch.On.OnValueRemoteUpdate(func(on bool) {
    if on == true {
        log.Println("Switch is on")
    } else {
        log.Println("Switch is off")
    }
})
```

When the switch is turned on "the analog way", you should set the state of the accessory.

```go
ac.Switch.On.SetValue(true)
```

## Multiple Accessories

When you create an IP transport, you can specify more than one accessory like this

```go
bridge := accessory.NewBridge(...)
outlet := accessory.NewOutlet(...)
lightbulb := accessory.NewColoredLightbulb(...)

hc.NewIPTransport(config, bridge, outlet.Accessory, lightbulb.Accessory)
```

By doing so, the *bridge* accessory will become a HomeKit bridge.
The *outlet* and *lightbulb* are the bridged accessories.

When adding the accessories to HomeKit, iOS only shows the *bridge* accessory.
Once the bridge was added, the other accessories appear automatically.

HomeKit requires that every accessory has a unique id, which must not change between system restarts.
`hc` automatically assigns the ids for you based on the order in which the accessories are added to the bridge.

But I recommend that you specify the accessory id yourself, via the [accessory.Config.ID](https://github.com/brutella/hc/blob/master/accessory/accessory.go#L13) field, like this.

```go
bridge := accessory.NewBridge(accessory.Info{Name: "Bridge", ID: 1})
outlet := accessory.NewOutlet(accessory.Info{Name: "Outlet", ID: 2})
lightbulb := accessory.NewColoredLightbulb(accessory.Info{Name: "Light", ID: 3})
```

## Accessory Architecture

HomeKit uses a hierarchical architecture to define accessories, services and characeristics.
At the root level there is an accessory.
Every accessory contains services.
And every service contains characteristics.

For example a [lightbulb accessory](accessory/lightbulb.go) contains a [lightbulb service](service/lightbulb.go).
This service contains characteristics like [on](characteristic/on.go) and [brightness](characteristic/brightness.go).

There are predefined accessories, services and characteristics available in HomeKit.
Those types are defined in the packages [accessory](accessory), [service](service), [characteristic](characteristic).

# Contact

Matthias Hochgatterer

Website: [https://hochgatterer.me](https://hochgatterer.me)

Github: [https://github.com/brutella](https://github.com/brutella/)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)


# License

`hc` is available under the Apache License 2.0 license. See the LICENSE file for more info.

[homekit]: https://developer.apple.com/homekit/
