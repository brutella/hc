package dnssd

import (
	"context"
	"github.com/miekg/dns"
	"net"
	"testing"
	"time"
)

func TestRemove(t *testing.T) {
	cfg := Config{
		Name: "Test",
		Type: "_asdf._tcp",
		Port: 1234,
	}
	si, err := NewService(cfg)
	if err != nil {
		t.Fatal(err)
	}

	msg := new(dns.Msg)
	msg.Answer = []dns.RR{SRV(si), TXT(si)}

	answers := []dns.RR{SRV(si), TXT(si), PTR(si)}
	unknown := remove(msg.Answer, answers)

	if x := len(unknown); x != 1 {
		t.Fatal(x)
	}

	rr := unknown[0]
	if _, ok := rr.(*dns.PTR); !ok {
		t.Fatal("Invalid type", rr)
	}
}

func TestRegisterServiceWithExplicitIP(t *testing.T) {
	cfg := Config{
		Host:   "My Computer",
		Name:   "Test",
		Type:   "_asdf._tcp",
		Domain: "local",
		IPs:    []net.IP{net.ParseIP("192.168.0.123")},
		Port:   12345,
	}
	sv, err := NewService(cfg)
	if err != nil {
		t.Fatal(err)
	}

	if is, want := len(sv.IPs), 1; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	conn := newTestConn()
	otherConn := newTestConn()
	conn.in = otherConn.out
	conn.out = otherConn.in

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-time.After(10 * time.Millisecond)

		lookupCtx, lookupCancel := context.WithTimeout(ctx, 2*time.Second)

		defer lookupCancel()
		defer cancel()

		srv, err := lookupInstance(lookupCtx, "Test._asdf._tcp.local.", otherConn)
		if err != nil {
			t.Fatal(err)
		}

		if is, want := srv.Name, "Test"; is != want {
			t.Fatalf("%v != %v", is, want)
		}

		if is, want := srv.Type, "_asdf._tcp"; is != want {
			t.Fatalf("%v != %v", is, want)
		}
	}()

	r := newResponder(conn)
	r.Add(sv)
	r.Respond(ctx)
}
