package dnssd

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	"reflect"
)

func PTR(srv Service) *dns.PTR {
	return &dns.PTR{
		Hdr: dns.RR_Header{
			Name:   srv.ServiceName(),
			Rrtype: dns.TypePTR,
			Class:  dns.ClassINET,
			Ttl:    TtlDefault,
		},
		Ptr: srv.ServiceInstanceName(),
	}
}

func DNSSDServicesPTR(srv Service) *dns.PTR {
	return &dns.PTR{
		Hdr: dns.RR_Header{
			Name:   srv.ServicesMetaQueryName(),
			Rrtype: dns.TypePTR,
			Class:  dns.ClassINET,
			Ttl:    TtlDefault,
		},
		Ptr: srv.ServiceName(),
	}
}

func SRV(srv Service) *dns.SRV {
	return &dns.SRV{
		Hdr: dns.RR_Header{
			Name:   srv.ServiceInstanceName(),
			Rrtype: dns.TypeSRV,
			Class:  dns.ClassINET,
			Ttl:    TtlHostname,
		},
		Priority: 0,
		Weight:   0,
		Port:     uint16(srv.Port),
		Target:   srv.Hostname(),
	}
}

func TXT(srv Service) *dns.TXT {
	txts := []string{}
	for key, value := range srv.Text {
		txts = append(txts, fmt.Sprintf("%s=%s", key, value))
	}

	return &dns.TXT{
		Hdr: dns.RR_Header{
			Name:   srv.ServiceInstanceName(),
			Rrtype: dns.TypeTXT,
			Class:  dns.ClassINET,
			Ttl:    TtlDefault,
		},
		Txt: txts,
	}
}

func NSEC(rr dns.RR, srv Service, iface *net.Interface) *dns.NSEC {
	switch r := rr.(type) {
	case *dns.PTR:
		return &dns.NSEC{
			Hdr: dns.RR_Header{
				Name:   r.Ptr,
				Rrtype: dns.TypeNSEC,
				Class:  dns.ClassINET,
				Ttl:    TtlDefault,
			},
			NextDomain: r.Ptr,
			TypeBitMap: []uint16{dns.TypeTXT, dns.TypeSRV},
		}
	case *dns.SRV:
		types := []uint16{}
		var ips []net.IP
		if iface != nil {
			if srv.IfaceIPs[iface.Name] != nil {
				ips = srv.IfaceIPs[iface.Name]
			}
		} else {
			ips = srv.IPs
		}
		if includesIPv4(ips) {
			types = append(types, dns.TypeA)
		}
		if includesIPv6(ips) {
			types = append(types, dns.TypeAAAA)
		}

		if len(types) > 0 {
			return &dns.NSEC{
				Hdr: dns.RR_Header{
					Name:   r.Target,
					Rrtype: dns.TypeNSEC,
					Class:  dns.ClassINET,
					Ttl:    TtlDefault,
				},
				NextDomain: r.Target,
				TypeBitMap: types,
			}
		}
	default:
		break
	}

	return nil
}

func A(srv Service, iface *net.Interface) []*dns.A {
	var ips []net.IP
	if iface != nil && srv.IfaceIPs != nil {
		if srv.IfaceIPs[iface.Name] != nil {
			ips = srv.IfaceIPs[iface.Name]
		}
	} else {
		ips = srv.IPs
	}

	var as []*dns.A
	for _, ip := range ips {
		if ip.To4() != nil {
			a := &dns.A{
				Hdr: dns.RR_Header{
					Name:   srv.Hostname(),
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    TtlHostname,
				},
				A: ip,
			}
			as = append(as, a)
		}
	}

	return as
}

func AAAA(srv Service, iface *net.Interface) []*dns.AAAA {
	var ips []net.IP
	if iface != nil && srv.IfaceIPs != nil {
		if srv.IfaceIPs[iface.Name] != nil {
			ips = srv.IfaceIPs[iface.Name]
		}
	} else {
		ips = srv.IPs
	}

	var aaaas []*dns.AAAA
	for _, ip := range ips {
		if ip.To4() == nil && ip.To16() != nil {
			aaaa := &dns.AAAA{
				Hdr: dns.RR_Header{
					Name:   srv.Hostname(),
					Rrtype: dns.TypeAAAA,
					Class:  dns.ClassINET,
					Ttl:    TtlHostname,
				},
				AAAA: ip,
			}
			aaaas = append(aaaas, aaaa)
		}
	}

	return aaaas
}

// Returns true if ips contains IPv4 addresses.
func includesIPv4(ips []net.IP) bool {
	for _, ip := range ips {
		if ip.To4() != nil {
			return true
		}
	}

	return false
}

// Returns true if ips contains IPv6 addresses.
func includesIPv6(ips []net.IP) bool {
	for _, ip := range ips {
		if ip.To4() == nil && ip.To16() != nil {
			return true
		}
	}

	return false
}

// Removes this from that.
func remove(this []dns.RR, that []dns.RR) []dns.RR {
	var result []dns.RR
	for _, thatRr := range that {
		isUnknown := true
		for _, thisRr := range this {
			switch a := thisRr.(type) {
			case *dns.PTR:
				if ptr, ok := thatRr.(*dns.PTR); ok {
					if a.Ptr == ptr.Ptr && a.Hdr.Name == ptr.Hdr.Name && a.Hdr.Ttl > ptr.Hdr.Ttl/2 {
						isUnknown = false
					}
				}
			case *dns.SRV:
				if srv, ok := thatRr.(*dns.SRV); ok {
					if a.Target == srv.Target && a.Port == srv.Port && a.Hdr.Name == srv.Hdr.Name && a.Hdr.Ttl > srv.Hdr.Ttl/2 {
						isUnknown = false
					}
				}
			case *dns.TXT:
				if txt, ok := thatRr.(*dns.TXT); ok {
					if reflect.DeepEqual(a.Txt, txt.Txt) && a.Hdr.Ttl > txt.Hdr.Ttl/2 {
						isUnknown = false
					}
				}
			}
		}

		if isUnknown {
			result = append(result, thatRr)
		}
	}

	return result
}

// mergeMsgs merges the records in msgs into one message.
func mergeMsgs(msgs []*dns.Msg) *dns.Msg {
	resp := new(dns.Msg)
	resp.Answer = []dns.RR{}
	resp.Ns = []dns.RR{}
	resp.Extra = []dns.RR{}
	resp.Question = []dns.Question{}

	for _, msg := range msgs {
		if msg.Answer != nil {
			resp.Answer = append(resp.Answer, remove(resp.Answer, msg.Answer)...)
		}
		if msg.Ns != nil {
			resp.Ns = append(resp.Ns, remove(resp.Ns, msg.Ns)...)
		}
		if msg.Extra != nil {
			resp.Extra = append(resp.Extra, remove(resp.Extra, msg.Extra)...)
		}

		if msg.Question != nil {
			resp.Question = append(resp.Question, msg.Question...)
		}
	}

	return resp
}
