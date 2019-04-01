package dnssd

import (
	"context"
	"github.com/miekg/dns"
)

type AddServiceFunc func(Service)
type RmvServiceFunc func(Service)

func LookupType(ctx context.Context, service string, add AddServiceFunc, rmv RmvServiceFunc) (err error) {
	conn, err := newMDNSConn()
	if err != nil {
		return err
	}
	defer conn.close()

	return lookupType(ctx, service, conn, add, rmv)
}

func lookupType(ctx context.Context, service string, conn MDNSConn, add AddServiceFunc, rmv RmvServiceFunc) (err error) {
	var cache = NewCache()

	m := new(dns.Msg)
	m.Question = []dns.Question{
		dns.Question{service, dns.TypePTR, dns.ClassINET},
	}
	// TODO include known answers which current ttl is more than half of the correct ttl (see TFC6772 7.1: Known-Answer Supression)
	// m.Answer = ...
	// m.Authoritive = false // because our answers are *believes*

	readCtx, readCancel := context.WithCancel(ctx)
	defer readCancel()

	ch := conn.Read(readCtx)

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
			return ctx.Err()
		}
	}
}
