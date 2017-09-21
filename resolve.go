package dnssd

import (
	"context"
	"github.com/miekg/dns"
)

// LookupInstance resolves a service by its service instance name.
func LookupInstance(ctx context.Context, instance string) (srv Service, err error) {
	var conn MDNSConn
	var cache = NewCache()

	conn, err = NewMDNSConn()

	if err != nil {
		return
	}

	defer conn.Close()

	m := new(dns.Msg)

	srvQ := dns.Question{instance, dns.TypeSRV, dns.ClassINET}
	txtQ := dns.Question{instance, dns.TypeTXT, dns.ClassINET}
	setQuestionUnicast(&srvQ)
	setQuestionUnicast(&txtQ)

	m.Question = []dns.Question{srvQ, txtQ}

	readCtx, readCancel := context.WithCancel(ctx)
	defer readCancel()

	ch := conn.Read(readCtx)

	q := &Query{msg: m}
	conn.SendQuery(q)

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
