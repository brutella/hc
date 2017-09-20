package dnssd

import (
	"context"
	"net"
)

type ReadFunc func(*Request)

func (r *responder) Debug(ctx context.Context, fn ReadFunc) {
	conn := r.conn.(*mdnsConn)
	ifaces, _ := net.Interfaces()
	for _, ifi := range ifaces {
		conn.ipv4.JoinGroup(&ifi, &net.UDPAddr{IP: IPv4Group})
		conn.ipv6.JoinGroup(&ifi, &net.UDPAddr{IP: IPv6Group})
	}

	readCtx, readCancel := context.WithCancel(ctx)
	defer readCancel()

	ch := conn.read(readCtx)

	for {
		select {
		case req := <-ch:
			fn(req)
		case <-ctx.Done():
			return
		}
	}
}
