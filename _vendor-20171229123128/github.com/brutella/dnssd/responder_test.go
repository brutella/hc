package dnssd

import (
	"github.com/miekg/dns"
	"net"
	"testing"
)

func TestRemove(t *testing.T) {
	si := NewService("A", "._hap._tcp", "local.", "Computer", []net.IP{}, 1234)

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
