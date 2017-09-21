# DNS-SD

This library implements [Multicast DNS](mDNS) and [DNS-Based Service Discovery](dnssd) to provide zero-configuration operations. It lets you announce and find services in a specific mDNS domain.

[mdns]: https://tools.ietf.org/html/rfc6762
[dnssd]: https://tools.ietf.org/html/rfc6763

## TODO

- [ ] Support negative responses (RFC6762 6.1)
- [ ] Handle txt records case insensitive
- [ ] Remove outdated services from cache regularly
- [ ] Implement lexicographically later algorithm
- [ ] Make sure that hostnames are FQDNs

# Contact

Matthias Hochgatterer

Github: [https://github.com/brutella](https://github.com/brutella/)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)


# License

*dnssd* is available under the MIT license. See the LICENSE file for more info.