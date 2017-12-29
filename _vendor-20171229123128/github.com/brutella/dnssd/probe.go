package dnssd

import (
	"context"
	"fmt"
	"github.com/brutella/dnssd/log"
	"github.com/miekg/dns"
	"math/rand"
	"net"
	"strings"
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

	// When ready to send its Multicast DNS probe packet(s) the host should
	// first wait for a short random delay time, uniformly distributed in
	// the range 0-250 ms. (RFC6762 8.1)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	delay := time.Duration(r.Intn(250)) * time.Millisecond
	log.Debug.Println("Probing delay", delay)
	time.Sleep(delay)

	return probeService(probeCtx, conn, srv, 1*time.Millisecond, false)
}

func ReprobeService(ctx context.Context, srv Service) (Service, error) {
	conn, err := newMDNSConn()

	if err != nil {
		return srv, err
	}

	defer conn.close()
	return probeService(ctx, conn, srv, 1*time.Millisecond, true)
}

func probeService(ctx context.Context, conn MDNSConn, srv Service, delay time.Duration, probeOnce bool) (s Service, e error) {
	candidate := srv.Copy()
	prevConflict := probeConflict{}

	// Keep track of the number of conflicts
	numHostConflicts := 0
	numNameConflicts := 0

	for i := 1; i <= 100; i++ {

		conflict, err := probe(ctx, conn, *candidate)
		if err != nil {
			e = err
			return
		}

		if conflict.hasNone() {
			s = *candidate
			return
		}

		candidate = candidate.Copy()

		if conflict.hostname && (prevConflict.hostname || probeOnce) {
			numHostConflicts++
			candidate.Host = fmt.Sprintf("%s-%d", srv.Host, numHostConflicts+1)
			conflict.hostname = false
		}

		if conflict.serviceName && (prevConflict.serviceName || probeOnce) {
			numNameConflicts++
			candidate.Name = fmt.Sprintf("%s-%d", srv.Name, numNameConflicts+1)
			conflict.serviceName = false
		}

		prevConflict = conflict

		if i == 10 {
			// [...] after the tenth try, the device correctly limits its rate to no more than one try per second. [...]
			// – Bonjour conformance test
			delay = 1 * time.Second
		}

		log.Debug.Println("Probing wait", delay)
		time.Sleep(delay)
	}

	return
}

func probe(ctx context.Context, conn MDNSConn, service Service) (conflict probeConflict, err error) {

	msg := new(dns.Msg)

	instanceQ := dns.Question{
		Name:   service.ServiceInstanceName(),
		Qtype:  dns.TypeANY,
		Qclass: dns.ClassINET,
	}

	hostQ := dns.Question{
		Name:   service.Hostname(),
		Qtype:  dns.TypeANY,
		Qclass: dns.ClassINET,
	}

	// Responses to probe should be unicast
	setQuestionUnicast(&instanceQ)
	setQuestionUnicast(&hostQ)

	msg.Question = []dns.Question{instanceQ, hostQ}

	srv := SRV(service)
	as := A(service, nil)
	aaaas := AAAA(service, nil)

	var authority = []dns.RR{srv}
	for _, a := range as {
		authority = append(authority, a)
	}
	for _, aaaa := range aaaas {
		authority = append(authority, aaaa)
	}
	msg.Ns = authority

	readCtx, readCancel := context.WithCancel(ctx)
	defer readCancel()

	ch := conn.Read(readCtx)

	queryTime := time.After(1 * time.Millisecond)
	queriesCount := 1

	for {
		select {
		case req := <-ch:
			answers := allRecords(req.msg)
			for _, answer := range answers {
				switch rr := answer.(type) {
				case *dns.A:
					if len(as) == 0 {
						continue
					}

					if isLexicographicalLaterA(rr, as[0]) {
						conflict.hostname = true
					}

				case *dns.AAAA:
					if len(aaaas) == 0 {
						continue
					}

					if isLexicographicalLaterAAAA(rr, aaaas[0]) {
						conflict.hostname = true
					}

				case *dns.SRV:
					if isLexicographicalLaterSRV(rr, srv) {
						conflict.serviceName = true
					}

				default:
					break
				}
			}

		case <-ctx.Done():
			err = ctx.Err()
			return

		case <-queryTime:
			// Stop on conflict
			if !conflict.hasNone() {
				return
			}

			// Stop after 3 probe queries
			if queriesCount > 3 {
				return
			}

			queriesCount++
			log.Debug.Println("Sending probe", msg)
			q := &Query{msg: msg}
			conn.SendQuery(q)

			delay := 250 * time.Millisecond
			log.Debug.Println("Sending wait", delay)
			queryTime = time.After(delay)
		}
	}

	return
}

type probeConflict struct {
	hostname    bool
	serviceName bool
}

func (pr probeConflict) hasNone() bool {
	return !pr.hostname && !pr.serviceName
}

func isLexicographicalLaterA(this *dns.A, that *dns.A) bool {
	if strings.EqualFold(this.Hdr.Name, that.Hdr.Name) {
		log.Debug.Println("Conflicting hosts")
		if !isValidRR(this) {
			log.Debug.Println("Invalid record produces conflict")
			return true
		}

		switch compareIP(this.A.To4(), that.A.To4()) {
		case -1:
			log.Debug.Println("Lexicographical earlier")
			break
		case 1:
			log.Debug.Println("Lexicographical later")
			return true
		default:
			log.Debug.Println("Tiebreak")
			break
		}
	}

	return false
}

func isLexicographicalLaterAAAA(this *dns.AAAA, that *dns.AAAA) bool {
	if strings.EqualFold(this.Hdr.Name, that.Hdr.Name) {
		log.Debug.Println("Conflicting hosts")
		if !isValidRR(this) {
			log.Debug.Println("Invalid record produces conflict")
			return true
		}

		switch compareIP(this.AAAA.To16(), that.AAAA.To16()) {
		case -1:
			log.Debug.Println("Lexicographical earlier")
			break
		case 1:
			log.Debug.Println("Lexicographical later")
			return true
		default:
			log.Debug.Println("Tiebreak")
			break
		}
	}

	return false
}

func isLexicographicalLaterSRV(this *dns.SRV, that *dns.SRV) bool {
	if strings.EqualFold(this.Hdr.Name, that.Hdr.Name) {
		log.Debug.Println("Conflicting SRV")

		if !isValidRR(this) {
			log.Debug.Println("Invalid record produces conflict")
			return true
		}

		switch compareSRV(this, that) {
		case -1:
			log.Debug.Println("Lexicographical earlier")
			break
		case 1:
			log.Debug.Println("Lexicographical later")
			return true
		default:
			log.Debug.Println("Tiebreak")
			break
		}
	}

	return false
}

func isValidRR(rr dns.RR) bool {
	switch r := rr.(type) {
	case *dns.A:
		return !net.IPv4zero.Equal(r.A)
	case *dns.AAAA:
		return !net.IPv6zero.Equal(r.AAAA)
	case *dns.SRV:
		return len(r.Target) > 0 && r.Port != 0
	default:
		break
	}

	return true
}

func compareIP(this net.IP, that net.IP) int {
	count := len(this)
	if count > len(that) {
		count = len(that)
	}

	for i := 0; i < count; i++ {
		if this[i] < that[i] {
			return -1
		} else if this[i] > that[i] {
			return 1
		}
	}

	if len(this) < len(that) {
		return -1
	} else if len(this) > len(that) {
		return 1
	}
	return 0
}

func compareSRV(this *dns.SRV, that *dns.SRV) int {
	if this.Priority < that.Priority {
		return -1
	} else if this.Priority > that.Priority {
		return 1
	}

	if this.Weight < that.Weight {
		return -1
	} else if this.Weight > that.Weight {
		return 1
	}

	if this.Port < that.Port {
		return -1
	} else if this.Port > that.Port {
		return 1
	}

	return strings.Compare(this.Target, that.Target)
}
