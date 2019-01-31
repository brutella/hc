package dnssd

import (
	"context"
	"github.com/miekg/dns"
	"net"
	"testing"
	"time"
)

var testIface = &net.Interface{
	Index:        0,
	MTU:          0,
	Name:         "lo0",
	HardwareAddr: []byte{},
	Flags:        net.FlagUp,
}

var testAddr = net.UDPAddr{
	IP:   net.IP{},
	Port: 1234,
	Zone: "",
}

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

func (c *testConn) Drain(ctx context.Context) {}

func (c *testConn) Close() {}

func (c *testConn) start(ctx context.Context) {
	for {
		select {
		case msg := <-c.in:
			req := &Request{msg: msg, from: &testAddr, iface: testIface}
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

	cfg := Config{
		Name: "My Service",
		Type: "_hap._tcp",
		Host: "My-Computer",
		Port: 12334,
		ifaceIPs: map[string][]net.IP{
			testIface.Name: []net.IP{net.ParseIP("192.168.0.122")},
		},
	}
	srv, err := NewService(cfg)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		otherCfg := cfg.Copy()
		otherCfg.ifaceIPs = map[string][]net.IP{
			testIface.Name: []net.IP{net.ParseIP("192.168.0.123")},
		}
		otherSrv, otherErr := NewService(otherCfg)
		if otherErr != nil {
			t.Fatal(otherErr)
		}
		otherResp := newResponder(otherConn)
		otherResp.Add(otherSrv)
		otherResp.Respond(ctx)
	}()

	// Allow other responder to claim service instance
	<-time.After(1 * time.Second)
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
