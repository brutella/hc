# DNS-SD

[![Build Status](https://travis-ci.org/brutella/hc.svg)](https://travis-ci.org/brutella/dnssd)

This library implements [Multicast DNS][mdns] and [DNS-Based Service Discovery][dnssd] to provide zero-configuration operations. It lets you announce and find services in a specific link-local domain.

[mdns]: https://tools.ietf.org/html/rfc6762
[dnssd]: https://tools.ietf.org/html/rfc6763

## Usage

#### Create a mDNS responder

The following code creates a service with name "My Website._http._tcp.local." for the host "My Computer" which has the IP "192.168.0.123" on port "12345". The service is added to a responder.

```go
service := dnssd.NewService("My Website", "_http._tcp.", "local.", "My Computer", []net.IP{net.ParseIP("192.168.0.123")}, 12345)
resp, _ := dnssd.NewResponder()
handle, _ := resp.Add(service)

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

resp.Respond(ctx)
```

If the service should be published on all available network interfaces, you can provide an empty host name and no IP addresses. The service will then get the local host name and IP addresses assigned to it.

```go
service := dnssd.NewService("My Website", "_http._tcp.", "local.", "", nil, 12345)
```

When calling `Respond` the responder probes for the service instance name and host name to be unqiue in the network. Once probing is finished, the service will be announced.

#### Update TXT records

Once a service is added to a responder, you have to use the `handle` object to update properties.

```go
handle.UpdateText(map[string]string{"key1": "value1", "key2": "value2"}, resp)
```

## Examples

There are examples in the `_cmd` directory to register services, resolve service instances and to browse for service types.

## Conformance

This library passes the [multicast DNS tests](https://github.com/brutella/dnssd/blob/36a2d8c541aab14895fc5492d5ad8ec447a67c47/_cmd/bct/ConformanceTestResults) of Apple's Bonjour Conformance Test.

## TODO

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
