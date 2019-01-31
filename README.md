# DNS-SD

[![Build Status](https://travis-ci.org/brutella/hc.svg)](https://travis-ci.org/brutella/dnssd)

This library implements [Multicast DNS][mdns] and [DNS-Based Service Discovery][dnssd] to provide zero-configuration operations. It lets you announce and find services in a specific link-local domain.

[mdns]: https://tools.ietf.org/html/rfc6762
[dnssd]: https://tools.ietf.org/html/rfc6763

## Usage

#### Create a mDNS responder

The following code creates a service with name "My Website._http._tcp.local." for the host "My Computer" which has the IP "192.168.0.123" on port "12345". The service is added to a responder.

```go
cfg := dnssd.Config{
    Name:   "My Website",
    Type:   "_http._tcp",
    Domain: "local",
    Host:   "My Computer",
    IPs:    []net.IP{net.ParseIP("192.168.0.123")},
    Port:   12345,
}
sv, _ := dnssd.NewService(cfg)
```

In most cases you only need to specify the name, type and port of the service.

```go
cfg := dnssd.Config{
    Name:   "My Website",
    Type:   "_http._tcp",
    Port:   12345,
}
sv, _ := dnssd.NewService(cfg)
```

Then you create a responder and add the service to it.
```go
rp, _ := dnssd.NewResponder()
hdl, _ := rp.Add(sv)

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

rp.Respond(ctx)
```

When calling `Respond` the responder probes for the service instance name and host name to be unqiue in the network. 
Once probing is finished, the service will be announced.

#### Update TXT records

Once a service is added to a responder, you can use the `hdl` to update properties.

```go
hdl.UpdateText(map[string]string{"key1": "value1", "key2": "value2"}, rsp)
```

## Examples

There are examples in the `_cmd` directory to register services, resolve service instances and to browse for service types.

## Conformance

This library passes the [multicast DNS tests](https://github.com/brutella/dnssd/blob/36a2d8c541aab14895fc5492d5ad8ec447a67c47/_cmd/bct/ConformanceTestResults) of Apple's Bonjour Conformance Test.

## TODO

- [ ] Support hot plugging
- [ ] Support negative responses (RFC6762 6.1)
- [ ] Handle txt records case insensitive
- [ ] Remove outdated services from cache regularly
- [ ] Make sure that hostnames are FQDNs

# Contact

Matthias Hochgatterer

Github: [https://github.com/brutella](https://github.com/brutella/)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)


# License

*dnssd* is available under the MIT license. See the LICENSE file for more info.
