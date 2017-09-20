package dnssd

import (
	"context"
	"github.com/miekg/dns"
)

// LookupInstance resolves a service by its service instance name.
func LookupInstance(ctx context.Context, instance string) (srv Service, err error) {
	var conn *mdnsConn
	var cache = NewCache()

	conn, err = newMDNSConn()

	if err != nil {
		return
	}

	defer conn.close()

	m := new(dns.Msg)
	m.Question = []dns.Question{
		dns.Question{instance, dns.TypeSRV, dns.ClassINET},
		dns.Question{instance, dns.TypeTXT, dns.ClassINET},
	}

	readCtx, readCancel := context.WithCancel(ctx)
	defer readCancel()

	ch := conn.read(readCtx)

	conn.sendQuery(m)

	for {
		select {
		case req := <-ch:
			cache.UpdateFrom(req.msg)
			if s, ok := cache.services[instance]; ok && s.IPs != nil && len(s.IPs) > 0 {
				srv = *s
				return
			}
		case <-ctx.Done():
			err = ctx.Err()
			return
		}
	}

	return
}
