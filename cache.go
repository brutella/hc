package dnssd

import (
	"github.com/miekg/dns"
	"sort"
	"strings"
	"time"
)

type Cache struct {
	services map[string]*Service
}

func NewCache() *Cache {
	return &Cache{
		services: make(map[string]*Service),
	}
}

// UpdateFrom updates the cache from resource records in msg.
// TODO consider the cache-flush bit to make records as to be deleted in one second
func (c *Cache) UpdateFrom(msg *dns.Msg) (adds []Service, rmvs []Service) {
	answers := allRecords(msg)
	sort.Sort(ByType(answers))

loop:
	for _, answer := range answers {
		switch rr := answer.(type) {
		case *dns.PTR:
			ttl := time.Duration(rr.Hdr.Ttl) * time.Second

			var entry *Service
			if e, ok := c.services[rr.Ptr]; !ok {
				if ttl == 0 {
					// Ignore new records with no ttl
					break
				}
				entry = newService(rr.Ptr)
				adds = append(adds, *entry)
				c.services[entry.ServiceInstanceName()] = entry
			} else {
				entry = e
			}

			entry.Ttl = ttl
			entry.expiration = time.Now().Add(ttl)

		case *dns.SRV:
			ttl := time.Duration(rr.Hdr.Ttl) * time.Second
			var entry *Service
			if e, ok := c.services[rr.Hdr.Name]; !ok {
				if ttl == 0 {
					// Ignore new records with no ttl
					break
				}
				entry = newService(rr.Hdr.Name)
				adds = append(adds, *entry)
				c.services[entry.ServiceInstanceName()] = entry
			} else {
				entry = e
			}

			entry.SetHostname(rr.Target)
			entry.Ttl = ttl
			entry.expiration = time.Now().Add(ttl)
			entry.Port = int(rr.Port)

		case *dns.A:
			for _, entry := range c.services {
				if entry.Hostname() == rr.Hdr.Name {
					for _, ip := range entry.IPs {
						if ip.Equal(rr.A) {
							continue loop
						}
					}

					entry.IPs = append(entry.IPs, rr.A)
				}
			}

		case *dns.AAAA:
			for _, entry := range c.services {
				if entry.Hostname() == rr.Hdr.Name {
					for _, ip := range entry.IPs {
						if ip.Equal(rr.AAAA) {
							continue loop
						}
					}

					entry.IPs = append(entry.IPs, rr.AAAA)
				}
			}

		case *dns.TXT:
			if entry, ok := c.services[rr.Hdr.Name]; ok {
				text := make(map[string]string)
				for _, txt := range rr.Txt {
					var pairs = strings.Split(txt, " ")
					for _, pair := range pairs {
						elems := strings.SplitN(pair, "=", 2)
						if len(elems) == 2 {
							key := elems[0]
							value := elems[1]

							// Don't override existing keys
							// TODO make txt records case insensitive
							if _, ok := text[key]; !ok {
								text[key] = value
							}

							text[key] = value
						}
					}
				}

				entry.Text = text
				entry.Ttl = time.Duration(rr.Hdr.Ttl) * time.Second
				entry.expiration = time.Now().Add(entry.Ttl)
			}
		default:
			// ignore
			break
		}
	}

	// TODO remove outdated services regularly
	rmvs = c.removeExpired()

	return
}

func (c *Cache) removeExpired() []Service {
	var outdated []Service
	var services = c.services
	for key, srv := range services {
		if time.Now().After(srv.expiration) {
			outdated = append(outdated, *srv)
			delete(c.services, key)
		}
	}

	return outdated
}

type ByType []dns.RR

func (a ByType) Len() int      { return len(a) }
func (a ByType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByType) Less(i, j int) bool {
	switch a[i].(type) {
	case *dns.SRV:
		return true
	case *dns.PTR:
		return true
	case *dns.TXT:
		return false
	case *dns.A:
		return false
	case *dns.AAAA:
		return false
	case *dns.NSEC:
		return false
	}

	return false
}

func allRecords(m *dns.Msg) []dns.RR {
	var answ []dns.RR
	answ = append(answ, m.Answer...)
	answ = append(answ, m.Ns...)
	answ = append(answ, m.Extra...)

	return answ
}
