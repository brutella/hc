package dnssd

import (
	"github.com/miekg/dns"
	"net"
	"time"
)

// ServiceHandle serves a middleman between a service and a responder.
type ServiceHandle interface {
	UpdateText(text map[string]string, r Responder)
	Service() Service
}

type serviceHandle struct {
	service *Service
}

func (h *serviceHandle) UpdateText(text map[string]string, r Responder) {
	h.service.Text = text

	msg := new(dns.Msg)
	msg.Answer = []dns.RR{TXT(*h.service)}
	msg.Response = true
	msg.Authoritative = true

	setAnswerCacheFlushBit(msg)

	resp := &Response{msg: msg}

	rr := r.(*responder)

	rr.conn.SendResponse(resp)
	time.Sleep(1 * time.Second)
	rr.conn.SendResponse(resp)
}

func (h *serviceHandle) Service() Service {
	return *h.service
}

func (h *serviceHandle) IPv4s() []net.IP {
	var result []net.IP

	for _, ip := range h.service.IPs {
		if ip.To4() != nil {
			result = append(result, ip)
		}
	}

	return result
}

func (h *serviceHandle) IPv6s() []net.IP {
	var result []net.IP

	for _, ip := range h.service.IPs {
		if ip.To16() != nil {
			result = append(result, ip)
		}
	}

	return result
}
