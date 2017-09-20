package dnssd

import (
	"context"
	//"fmt"
	//"github.com/brutella/dnssd/log"
	"github.com/miekg/dns"
	"net"
	"sync"
	"testing"
	"time"
)

type testConn struct {
	read chan *Request

	in  chan *dns.Msg
	out chan *dns.Msg

	once sync.Once
}

func newTestConn() *testConn {
	c := &testConn{
		read: make(chan *Request),
		in:   make(chan *dns.Msg),
		out:  make(chan *dns.Msg),
		once: sync.Once{},
	}

	return c
}

func (c *testConn) SendQuery(q *Query) error {
	c.out <- q.msg
	return nil
}

func (c *testConn) SendResponse(resp *Response) error {
	c.out <- resp.msg
	return nil
}

func (c *testConn) Read(ctx context.Context) <-chan *Request {
	c.once.Do(func() {
		go c.start(ctx)
	})
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
	ctx, cancel := context.WithCancel(context.Background())
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

	resolved, err := probeService(ctx, conn, srv, 1*time.Millisecond)

	if x := err; x != nil {
		t.Fatal(x)
	}

	if is, want := resolved.Host, "My-Computer-2"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := resolved.Name, "My Service (2)"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
