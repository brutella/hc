package dnssd

import (
	"github.com/miekg/dns"
	"time"
)

// ServiceHandle serves a middleman between a service and a responder.
type ServiceHandle interface {
	UpdateText(text map[string]string, r Responder)
	Handles(s *Service) bool
	Service() *Service
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

func (h *serviceHandle) Handles(s *Service) bool {
	return h.service == s
}

func (h *serviceHandle) Service() *Service {
	return h.service
}
