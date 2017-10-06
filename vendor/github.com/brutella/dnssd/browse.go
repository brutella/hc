package dnssd

import (
	"context"
	"github.com/miekg/dns"
	"net"
)

type AddServiceFunc func(Service)
type RmvServiceFunc func(Service)

func LookupType(ctx context.Context, service string, add AddServiceFunc, rmv RmvServiceFunc) (err error) {
	var conn *mdnsConn
	var cache = NewCache()
	conn, err = newMDNSConn()

	if err != nil {
		return
	}

	defer conn.close()

	ifaces, _ := net.Interfaces()
	for _, ifi := range ifaces {
		conn.ipv4.JoinGroup(&ifi, &net.UDPAddr{IP: IPv4LinkLocalMulticast})
		conn.ipv6.JoinGroup(&ifi, &net.UDPAddr{IP: IPv6LinkLocalMulticast})
	}

	m := new(dns.Msg)
	m.Question = []dns.Question{
		dns.Question{service, dns.TypePTR, dns.ClassINET},
	}
	// TODO include known answers which current ttl is more than half of the correct ttl (see TFC6772 7.1: Known-Answer Supression)
	// m.Answer = ...
	// m.Authoritive = false // because our answers are *believes*

	readCtx, readCancel := context.WithCancel(ctx)
	defer readCancel()

	ch := conn.read(readCtx)

	q := &Query{msg: m}
	conn.SendQuery(q)

	for {
		select {
		case req := <-ch:
			adds, rmvs := cache.UpdateFrom(req.msg)

			for _, srv := range adds {
				if srv.ServiceName() == service {
					add(*srv)
				}
			}

			for _, srv := range rmvs {
				if srv.ServiceName() == service {
					rmv(*srv)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
