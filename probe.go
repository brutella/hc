package dnssd

import (
	"context"
	"fmt"
	"github.com/miekg/dns"
	"math/rand"
	"time"
)

// ProbeService probes for the hostname and service instance name of srv.
// If err == nil, the returned service is verified to be unique on the local network.
func ProbeService(ctx context.Context, srv Service) (Service, error) {
	conn, err := newMDNSConn()

	if err != nil {
		return srv, err
	}

	defer conn.close()

	// After one minute of probing, if the Multicast DNS responder has been
	// unable to find any unused name, it should log an error (RFC6762 9)
	probeCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	return probeService(probeCtx, conn, srv, 1*time.Millisecond)
}

func probeService(ctx context.Context, conn MDNSConn, srv Service, delay time.Duration) (s Service, e error) {
	candidate := srv.Copy()
	retry := 10
	for i := 1; i <= retry; i++ {
		res, err := probe(ctx, conn, *candidate)
		if err != nil {
			e = err
			return
		}

		if !res.includesConflict() {
			s = *candidate
			return
		}

		candidate = srv.Copy()

		if res.hostnameConflict {
			candidate.Host = fmt.Sprintf("%s-%d", srv.Host, i+1)
		}

		if res.serviceInstanceNameConflict {
			candidate.Name = fmt.Sprintf("%s (%d)", srv.Name, i+1)
		}

		time.Sleep(delay)
	}

	// [...] after the tenth try, the device correctly limits its rate to no more than one try per second. [...]
	// – Bonjour conformance test
	oneSec := 1 * time.Second
	return probeService(ctx, conn, srv, oneSec)
}

func probe(ctx context.Context, conn MDNSConn, srv Service) (res probeResult, err error) {

	msg := new(dns.Msg)

	instanceQ := dns.Question{
		Name:   srv.ServiceInstanceName(),
		Qtype:  dns.TypeANY,
		Qclass: dns.ClassINET,
	}

	hostQ := dns.Question{
		Name:   srv.Hostname(),
		Qtype:  dns.TypeANY,
		Qclass: dns.ClassINET,
	}

	// Responses to probe should be unicast
	setQuestionUnicast(&instanceQ)
	setQuestionUnicast(&hostQ)

	msg.Question = []dns.Question{instanceQ, hostQ}

	var authority = []dns.RR{SRV(srv)}
	for _, a := range A(srv, nil) {
		authority = append(authority, a)
	}
	for _, aaaa := range AAAA(srv, nil) {
		authority = append(authority, aaaa)
	}
	msg.Ns = authority

	readCtx, readCancel := context.WithCancel(ctx)
	defer readCancel()

	ch := conn.Read(readCtx)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	attempts := 0
	for {
		select {
		case req := <-ch:
			answers := allRecords(req.msg)
			for _, answer := range answers {

				if answer.Header().Ttl == 0 {
					// Ignore outdated records
					continue
				}

				switch rr := answer.(type) {
				case *dns.A:
					if rr.Hdr.Name == srv.Hostname() {
						res.hostnameConflict = true
					}
				case *dns.AAAA:
					if rr.Hdr.Name == srv.Hostname() {
						res.hostnameConflict = true
					}
				case *dns.SRV:
					if rr.Hdr.Name == srv.ServiceInstanceName() {
						res.serviceInstanceNameConflict = true
					}
				default:
					break
				}
			}

			if res.includesConflict() {
				return
			}

		case <-ctx.Done():
			err = ctx.Err()
			return

		default:
			if attempts >= 3 {
				// Stop after 3 messages
				return
			}

			q := &Query{msg: msg}
			conn.SendQuery(q)
			attempts++
			delay := r.Intn(250)
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}

	return
}

type probeResult struct {
	hostnameConflict            bool
	serviceInstanceNameConflict bool
}

func (pr probeResult) includesConflict() bool {
	return pr.hostnameConflict || pr.serviceInstanceNameConflict
}

// isLexicographicallyLater checks who is winning based on "lexicographically later" by
// comparing the RR.Header().Class, RR.Header().Rrtype and rdata(?)
// – RFC6763 Section 8.2.
//
// TODO(mah) Actually implement
func isLexicographicallyLater(this *dns.Msg, that *dns.Msg) bool {
	return true
}
