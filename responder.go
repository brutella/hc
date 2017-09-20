package dnssd

import (
	"context"
	"fmt"
	"github.com/brutella/dnssd/log"
	"github.com/miekg/dns"
	"sync"
	"time"
)

// Responder represents a mDNS responder.
type Responder interface {
	// Add adds a service to the responder.
	// Use the returned service handle to update service properties.
	Add(srv Service) (ServiceHandle, error)
	// Remove removes the service associated with the service handle from the responder.
	Remove(srv ServiceHandle)
	// Respond makes the receiver announcing and managing services.
	Respond(ctx context.Context) error
}

type responder struct {
	isRunning bool

	conn      MDNSConn
	unmanaged []*serviceHandle
	managed   []*serviceHandle

	mutex *sync.Mutex
}

func NewResponder() (Responder, error) {
	conn, err := newMDNSConn()
	if err != nil {
		return nil, err
	}

	return newResponder(conn), nil
}

func newResponder(conn MDNSConn) *responder {
	return &responder{
		isRunning: false,
		conn:      conn,
		unmanaged: []*serviceHandle{},
		managed:   []*serviceHandle{},
		mutex:     &sync.Mutex{},
	}
}

func (r *responder) Remove(h ServiceHandle) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for i, s := range r.managed {
		if h == s {
			handle := h.(*serviceHandle)
			r.unannounce([]*Service{handle.service})
			r.managed = append(r.managed[:i], r.managed[i+1:]...)
			return
		}
	}
}

func (r *responder) Add(srv Service) (ServiceHandle, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.isRunning {
		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()

		if srv, err := r.register(ctx, srv); err != nil {
			return nil, err
		} else {
			return r.addManaged(srv), nil
		}
	}

	return r.addUnmanaged(srv), nil
}

// announce sends announcement messages including all services.
func (r *responder) announce(services []*Service) {
	var msgs []*dns.Msg
	for _, srv := range services {
		var answer []dns.RR
		answer = append(answer, SRV(*srv))
		answer = append(answer, PTR(*srv))
		answer = append(answer, TXT(*srv))
		for _, a := range A(*srv, nil) {
			answer = append(answer, a)
		}

		for _, aaaa := range AAAA(*srv, nil) {
			answer = append(answer, aaaa)
		}
		msg := new(dns.Msg)
		msg.Answer = answer
		msgs = append(msgs, msg)
	}
	msg := mergeMsgs(msgs)
	msg.Response = true
	msg.Authoritative = true

	setAnswerCacheFlushBit(msg)

	resp := &Response{msg: msg}

	r.conn.SendResponse(resp)
	time.Sleep(1 * time.Second)
	r.conn.SendResponse(resp)
}

func (r *responder) register(ctx context.Context, srv Service) (Service, error) {
	if !r.isRunning {
		return srv, fmt.Errorf("cannot register service when responder is not responding")
	}

	log.Debug.Printf("Probing for host %s and service %s…\n", srv.Hostname(), srv.ServiceInstanceName())
	probed, err := ProbeService(ctx, srv)
	if err != nil {
		return srv, err
	}

	srvs := []*Service{&probed}
	for _, h := range r.managed {
		srvs = append(srvs, h.service)
	}
	log.Debug.Println("Announcing services", srvs)
	go r.announce(srvs)

	return probed, nil
}

func (r *responder) addManaged(srv Service) ServiceHandle {
	h := &serviceHandle{&srv}
	r.managed = append(r.managed, h)
	return h
}

func (r *responder) addUnmanaged(srv Service) ServiceHandle {
	h := &serviceHandle{&srv}
	r.unmanaged = append(r.unmanaged, h)
	return h
}

func (r *responder) Respond(ctx context.Context) error {
	r.mutex.Lock()
	r.isRunning = true
	for _, h := range r.unmanaged {
		if srv, err := r.register(ctx, *h.service); err != nil {
			return err
		} else {
			h.service = &srv
			r.managed = append(r.managed, h)
		}
	}
	r.unmanaged = []*serviceHandle{}
	r.mutex.Unlock()

	return r.respond(ctx)
}

func (r *responder) respond(ctx context.Context) error {
	if !r.isRunning {
		return fmt.Errorf("isRunning should be true before calling respond()")
	}

	readCtx, readCancel := context.WithCancel(ctx)
	defer readCancel()

	ch := r.conn.Read(readCtx)

	for {
		select {
		case req := <-ch:
			r.handleRequest(req, services(r.managed))
		case <-ctx.Done():
			r.unannounce(services(r.managed))
			r.conn.Close()
			r.isRunning = false
			return ctx.Err()
		}
	}
}

func (r *responder) unannounce(services []*Service) {
	if len(services) == 0 {
		return
	}

	log.Debug.Println("Send goodbye for", services)

	// Send goodbye packets
	var answer []dns.RR
	for _, srv := range services {
		answer = append(answer, PTR(*srv))
	}

	for _, a := range answer {
		a.Header().Ttl = 0
	}

	goodbye := new(dns.Msg)
	goodbye.Answer = answer
	goodbye.Response = true
	goodbye.Authoritative = true

	resp := &Response{msg: goodbye}

	r.conn.SendResponse(resp)
	time.Sleep(250 * time.Millisecond)
	r.conn.SendResponse(resp)
}

func (r *responder) handleRequest(req *Request, services []*Service) {
	if req.msg.Truncated {
		log.Info.Println("TODO(mah) Wait for additional messages to come if request was unicast (see 18.5)")
	}

	for _, q := range req.msg.Question {
		msgs := []*dns.Msg{}
		for _, srv := range services {
			log.Debug.Printf("%s tries to give response to question %v\n", srv.ServiceInstanceName(), q)
			if msg := r.handleQuestion(q, req, *srv); msg != nil {
				log.Debug.Println("Response", msg)
				msgs = append(msgs, msg)
			} else {
				log.Debug.Println("No response")
			}
		}

		msg := mergeMsgs(msgs)
		msg.SetReply(req.msg)
		msg.Question = nil
		msg.Response = true
		msg.Authoritative = true

		if len(msg.Answer) == 0 {
			continue
		}

		if isUnicastQuestion(q) {
			resp := &Response{msg: msg, addr: req.from}
			log.Debug.Printf("Send unicast response\n%v\n", msg)
			r.conn.SendResponse(resp)
		} else {
			resp := &Response{msg: msg}
			log.Debug.Printf("Send multicast response\n%v\n", msg)
			r.conn.SendResponse(resp)
		}
	}
}

func (r *responder) handleQuestion(q dns.Question, req *Request, srv Service) *dns.Msg {
	resp := new(dns.Msg)

	switch q.Name {
	case srv.ServiceName():
		ptr := PTR(srv)
		resp.Answer = []dns.RR{ptr}

		extra := []dns.RR{SRV(srv), TXT(srv)}

		for _, a := range A(srv, req.iface) {
			extra = append(extra, a)
		}

		for _, aaaa := range AAAA(srv, req.iface) {
			extra = append(extra, aaaa)
		}

		extra = append(extra, NSEC(ptr, srv, req.iface))
		resp.Extra = extra

	case srv.ServiceInstanceName():
		resp.Answer = []dns.RR{SRV(srv), TXT(srv), PTR(srv)}

		var extra []dns.RR

		for _, a := range A(srv, req.iface) {
			extra = append(extra, a)
		}

		for _, aaaa := range AAAA(srv, req.iface) {
			extra = append(extra, aaaa)
		}

		nsec := NSEC(SRV(srv), srv, req.iface)
		if nsec != nil {
			extra = append(extra, nsec)
		}

		resp.Extra = extra

	case srv.Hostname():
		var answer []dns.RR

		for _, a := range A(srv, req.iface) {
			answer = append(answer, a)
		}

		for _, aaaa := range AAAA(srv, req.iface) {
			answer = append(answer, aaaa)
		}

		resp.Answer = answer
		nsec := NSEC(SRV(srv), srv, req.iface)

		if nsec != nil {
			resp.Extra = []dns.RR{nsec}
		}

	case srv.ServicesMetaQueryName():
		resp.Answer = []dns.RR{DNSSDServicesPTR(srv)}

	default:
		return nil
	}

	log.Debug.Println("Answers\n", resp.Answer)
	log.Debug.Println("Extra\n", resp.Extra)
	log.Debug.Println("Known answers\n", req.msg.Answer)

	// Supress known answers
	resp.Answer = remove(req.msg.Answer, resp.Answer)
	log.Debug.Println("Unknown answers\n", resp.Answer)

	resp.SetReply(req.msg)
	resp.Question = nil
	resp.Response = true
	resp.Authoritative = true

	return resp
}

func services(hs []*serviceHandle) []*Service {
	var result []*Service
	for _, h := range hs {
		result = append(result, h.service)
	}

	return result
}
