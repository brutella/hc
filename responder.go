package dnssd

import (
	"context"
	"fmt"
	"github.com/brutella/dnssd/log"
	"github.com/miekg/dns"
	"math/rand"
	"net"
	"strings"
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

	mutex     *sync.Mutex
	truncated *Request
	random    *rand.Rand
	upIfaces  []string
}

var DefaultResponder, _ = NewResponder()

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
		random:    rand.New(rand.NewSource(time.Now().UnixNano())),
		upIfaces:  []string{},
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

// announce sends announcement messages including all services.
func (r *responder) announce(services []*Service) {
	names := ifaceNames(services)

	if len(names) == 0 {
		r.announceAtInterface(services, nil)
		return
	}

	for _, name := range names {
		if iface, err := net.InterfaceByName(name); err != nil {
			log.Debug.Println("Unable to find interface", name)
		} else {
			r.announceAtInterface(services, iface)
		}
	}
}

func (r *responder) announceAtInterface(services []*Service, iface *net.Interface) {
	var msgs []*dns.Msg
	for _, srv := range services {
		var ips []net.IP
		if iface != nil {
			ips = srv.IPsAtInterface(iface)
		} else {
			ips = srv.IPs
		}
		if len(ips) == 0 {
			log.Debug.Println("No IPs for service", srv.ServiceInstanceName())
			continue
		}

		var answer []dns.RR
		answer = append(answer, SRV(*srv))
		answer = append(answer, PTR(*srv))
		answer = append(answer, TXT(*srv))
		for _, a := range A(*srv, iface) {
			answer = append(answer, a)
		}

		for _, aaaa := range AAAA(*srv, iface) {
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

	resp := &Response{msg: msg, iface: iface}

	log.Debug.Println("Sending 1st announcement", msg)
	r.conn.SendResponse(resp)
	time.Sleep(1 * time.Second)
	log.Debug.Println("Sending 2nd announcement", msg)
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
			if len(r.managed) == 0 {
				// Ignore requests when no services are managed
				break
			}

			// If messages is truncated, we wait for the next message to come (RFC6762 18.5)
			if req.msg.Truncated {
				r.truncated = req
				log.Debug.Println("Waiting for additional answers...")
				break
			}

			// append request
			if r.truncated != nil && r.truncated.from.IP.Equal(req.from.IP) {
				log.Debug.Println("Add answers to truncated message")
				msgs := []*dns.Msg{r.truncated.msg, req.msg}
				r.truncated = nil
				req.msg = mergeMsgs(msgs)
			}

			// Conflicting records remove managed services from
			// the responder and trigger reprobing
			conflicts := r.findConflicts(req, r.managed)
			for _, h := range conflicts {
				log.Debug.Println("Reprobe for", h.service)
				go r.reprobe(h)
				for i, m := range r.managed {
					if h == m {
						r.managed = append(r.managed[:i], r.managed[i+1:]...)
						break
					}
				}
			}

			r.handleQuery(req, services(r.managed))

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

func (r *responder) handleQuery(req *Request, services []*Service) {
	for _, q := range req.msg.Question {
		msgs := []*dns.Msg{}
		for _, srv := range services {
			log.Debug.Printf("%s tries to give response to question %v\n", srv.ServiceInstanceName(), q)
			if msg := r.handleQuestion(q, req, *srv); msg != nil {
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
			log.Debug.Println("No answers")
			continue
		}

		if isUnicastQuestion(q) {
			resp := &Response{msg: msg, addr: req.from, iface: req.iface}
			log.Debug.Printf("Send unicast response\n%v\n", msg)
			r.conn.SendResponse(resp)
		} else {
			resp := &Response{msg: msg, iface: req.iface}
			log.Debug.Printf("Send multicast response\n%v\n", msg)
			r.conn.SendResponse(resp)
		}
	}
}

func (r *responder) findConflicts(req *Request, hs []*serviceHandle) []*serviceHandle {
	var conflicts []*serviceHandle
	for _, h := range hs {
		if containsConflictingAnswers(req, h) {
			log.Debug.Println("Received conflicting record", req.msg)
			conflicts = append(conflicts, h)
		}
	}

	return conflicts
}

func (r *responder) reprobe(h *serviceHandle) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	probed, err := ReprobeService(ctx, *h.service)
	if err != nil {
		return
	}

	h.service = &probed
	r.managed = append(r.managed, h)
	log.Debug.Println("Reannouncing services", r.managed)
	go r.announce(services(r.managed))
}

func (r *responder) handleQuestion(q dns.Question, req *Request, srv Service) *dns.Msg {
	resp := new(dns.Msg)

	switch strings.ToLower(q.Name) {
	case strings.ToLower(srv.ServiceName()):
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

		// Wait 20-125 msec for shared resource responses
		delay := time.Duration(r.random.Intn(105)+20) * time.Millisecond
		log.Debug.Println("Shared record response wait", delay)
		time.Sleep(delay)

	case strings.ToLower(srv.ServiceInstanceName()):
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

		// Set cache flush bit for non-shared records
		setAnswerCacheFlushBit(resp)

	case strings.ToLower(srv.Hostname()):
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

		// Set cache flush bit for non-shared records
		setAnswerCacheFlushBit(resp)

	case strings.ToLower(srv.ServicesMetaQueryName()):
		resp.Answer = []dns.RR{DNSSDServicesPTR(srv)}

	default:
		return nil
	}

	// Supress known answers
	resp.Answer = remove(req.msg.Answer, resp.Answer)

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

func ifaceNames(svs []*Service) []string {
	var names []string
	for _, sv := range svs {
		for name, _ := range sv.IfaceIPs {
			names = append(names, name)
		}
	}

	return names
}

func containsConflictingAnswers(req *Request, handle *serviceHandle) bool {
	answers := allRecords(req.msg)

	as := A(*handle.service, req.iface)
	aaaas := AAAA(*handle.service, req.iface)
	srv := SRV(*handle.service)

	for _, answer := range answers {
		switch rr := answer.(type) {
		case *dns.A:
			for _, a := range as {
				if isDenyingA(rr, a) {
					return true
				}
			}

		case *dns.AAAA:
			for _, aaaa := range aaaas {
				if isDenyingAAAA(rr, aaaa) {
					return true
				}
			}

		case *dns.SRV:
			if isDenyingSRV(rr, srv) {
				return true
			}

		default:
			break
		}
	}

	return false
}
