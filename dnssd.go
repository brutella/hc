package dnssd

import (
	"net"
)

var (
	IPv4Group = net.ParseIP("224.0.0.251")
	IPv6Group = net.ParseIP("ff02::fb")

	AddrIPv4 = &net.UDPAddr{
		IP:   IPv4Group,
		Port: 5353,
	}

	AddrIPv6 = &net.UDPAddr{
		IP:   IPv6Group,
		Port: 5353,
	}

	TtlHostname = 120
	TtlDefault  = 75 * 60
)

var (
	ttlHostname = uint32(TtlHostname)
	ttlDefault  = uint32(TtlDefault)
)
