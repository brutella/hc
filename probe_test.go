package dnssd

import (
	"context"
	"github.com/miekg/dns"
	"net"
	"testing"
	"time"
)

type testConn struct {
	read chan *Request

	in  chan *dns.Msg
	out chan *dns.Msg
}

func newTestConn() *testConn {
	c := &testConn{
		read: make(chan *Request),
		in:   make(chan *dns.Msg),
		out:  make(chan *dns.Msg),
	}

	return c
}

func (c *testConn) SendQuery(q *Query) error {
	go func() {
		c.out <- q.msg
	}()
	return nil
}

func (c *testConn) SendResponse(resp *Response) error {
	go func() {
		c.out <- resp.msg
	}()

	return nil
}

func (c *testConn) Read(ctx context.Context) <-chan *Request {
	go c.start(ctx)
	return c.read
}

func (c *testConn) Close() {
}

func (c *testConn) start(ctx context.Context) {
	for {
		select {
		case msg := <-c.in:
			req := &Request{msg: msg}
			c.read <- req
		case <-ctx.Done():
			return
		default:
			break
		}
	}
}

func TestProbing(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn := newTestConn()
	otherConn := newTestConn()
	conn.in = otherConn.out
	conn.out = otherConn.in

	srv := NewService("My Service", "._hap._tcp", "local.", "My Computer", []net.IP{net.ParseIP("192.168.0.122")}, 12334)

	go func() {
		otherSrv := NewService("My Service", "._hap._tcp", "local.", "My Computer", []net.IP{net.ParseIP("192.168.0.123")}, 43321)
		otherResp := newResponder(otherConn)
		otherResp.addManaged(otherSrv)
		otherResp.isRunning = true
		otherResp.respond(ctx)
	}()

	resolved, err := probeService(ctx, conn, srv, 1*time.Millisecond, true)

	if x := err; x != nil {
		t.Fatal(x)
	}

	if is, want := resolved.Host, "My-Computer-2"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := resolved.Name, "My Service-2"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestIsLexicographicLater(t *testing.T) {
	this := &dns.A{
		Hdr: dns.RR_Header{
			Name:   "MyPrinter.local.",
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    TtlHostname,
		},
		A: net.ParseIP("169.254.99.200"),
	}

	that := &dns.A{
		Hdr: dns.RR_Header{
			Name:   "MyPrinter.local.",
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    TtlHostname,
		},
		A: net.ParseIP("169.254.200.50"),
	}

	if is, want := compareIP(this.A.To4(), that.A.To4()), -1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := compareIP(that.A.To4(), this.A.To4()), 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
