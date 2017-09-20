package dnssd

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// Service represents a DNS-SD service instance
type Service struct {
	Name     string // e.g. Accessory (no trailing dot)
	Type     string // e.g. _hap._tcp.
	Domain   string // e.g. local.
	Host     string // e.g. MacBook (no trailing dot)
	Text     map[string]string
	Ttl      time.Duration // Original time to live
	Port     int
	IPs      []net.IP
	IfaceIPs map[string][]net.IP

	expiration time.Time
}

func NewService(name, typ, domain, host string, ips []net.IP, port int) Service {
	if len(host) == 0 {
		if osHost, err := os.Hostname(); err != nil {
			host = "unknown"
		} else {
			host = osHost
		}
	}

	var ifIPs map[string][]net.IP

	// If no ips are specified, use the resolved ips
	if ips == nil {
		ips = []net.IP{}
		osIPs, err := net.LookupIP(host)
		if err == nil && osIPs != nil && len(osIPs) > 0 {
			ips = append(ips, osIPs...)
		}

		ifIPs = make(map[string][]net.IP, 0)
		for _, iface := range multicastInterfaces() {
			ipv4, ipv6 := addrsForInterface(&iface)
			ifIPs[iface.Name] = append(ipv4, ipv6...)
		}

		if len(ips) == 0 {
			for _, ip := range ifIPs {
				ips = append(ips, ip...)
			}
		}
	}

	host = sanitizeHostname(host)

	// Must not end with .
	name = strings.Trim(name, ".")
	host, _ = parseHostname(host)

	// Must end with .
	typ = fmt.Sprintf("%s.", strings.Trim(typ, "."))
	domain = fmt.Sprintf("%s.", strings.Trim(domain, "."))

	return Service{
		Name:     name,
		Type:     typ,
		Domain:   domain,
		Host:     host,
		Text:     make(map[string]string),
		Port:     port,
		IPs:      ips,
		IfaceIPs: ifIPs,
	}
}

func newService(instance string) *Service {
	name, typ, domain := parseServiceInstanceName(instance)
	return &Service{
		Name:   name,
		Type:   typ,
		Domain: domain,
		Host:   "",
		Text:   make(map[string]string),
	}
}

func (srv Service) Copy() *Service {
	return &Service{
		Name:       srv.Name,
		Type:       srv.Type,
		Domain:     srv.Domain,
		Host:       srv.Host,
		Text:       srv.Text,
		Ttl:        srv.Ttl,
		Port:       srv.Port,
		IPs:        srv.IPs,
		expiration: srv.expiration,
	}
}

func (srv Service) ServiceInstanceName() string {
	return fmt.Sprintf("%s.%s%s", srv.Name, srv.Type, srv.Domain)
}

func (srv Service) ServiceName() string {
	return fmt.Sprintf("%s%s", srv.Type, srv.Domain)
}

func (srv Service) Hostname() string {
	return fmt.Sprintf("%s.%s", srv.Host, srv.Domain)
}

func (srv *Service) SetHostname(hostname string) {
	name, domain := parseHostname(hostname)
	if domain == srv.Domain {
		srv.Host = name
	}
}

func (srv Service) ServicesMetaQueryName() string {
	return fmt.Sprintf("_services._dns-sd._udp.%s", srv.Domain)
}

func (srv Service) Equal(other Service) bool {
	if srv.Name == other.Name && srv.Type == other.Type && srv.Domain == other.Domain && srv.Host == other.Host && srv.Port == other.Port {
		for _, ip := range srv.IPs {
			var found = false
			for _, otherIP := range other.IPs {
				if ip.String() == otherIP.String() {
					found = true
				}
			}

			if !found {
				return false
			}
		}
		return true
	}

	return false
}

func parseServiceInstanceName(str string) (name string, service string, domain string) {
	elems := strings.Split(str, ".")
	for i, e := range elems {

		if len(e) == 0 {
			continue
		}

		if i == 0 {
			name = e
			continue
		}

		if strings.HasPrefix(e, "_") {
			if len(service) > 0 {
				service = fmt.Sprintf("%s%s.", service, e)
			} else {
				service = e + "."
			}
		} else {
			if len(domain) > 0 {
				domain = fmt.Sprintf("%s%s.", domain, e)
			} else {
				domain = e + "."
			}
		}
	}

	return
}

// Get Fully Qualified Domain Name
// returns "unknown" or hostanme in case of error
func hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}

	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return hostname
	}

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip, err := ipv4.MarshalText()
			if err != nil {
				return hostname
			}
			hosts, err := net.LookupAddr(string(ip))
			if err != nil || len(hosts) == 0 {
				return hostname
			}
			fqdn := hosts[0]
			return strings.TrimSuffix(fqdn, ".") // return fqdn without trailing dot
		}
	}
	return hostname
}

/// TODO make it a FQDN
func sanitizeHostname(name string) string {
	return strings.Replace(name, " ", "-", -1)
}

func parseHostname(str string) (name string, domain string) {
	elems := strings.Split(str, ".")
	for i, e := range elems {

		if len(e) == 0 {
			continue
		}

		if i == 0 {
			name = e
			continue
		}

		if len(domain) > 0 {
			domain = fmt.Sprintf("%s%s.", domain, e)
		} else {
			domain = e + "."
		}
	}

	return
}

func multicastInterfaces() []net.Interface {
	var interfaces []net.Interface
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, ifi := range ifaces {
		if (ifi.Flags & net.FlagUp) == 0 {
			continue
		}

		if (ifi.Flags & net.FlagLoopback) > 0 {
			continue
		}

		if (ifi.Flags & net.FlagMulticast) > 0 {
			interfaces = append(interfaces, ifi)
		}
	}

	return interfaces
}

// addrsForInterface returns ipv4 and ipv6 addresses for a specific interface.
func addrsForInterface(iface *net.Interface) ([]net.IP, []net.IP) {
	var v4, v6, v6local []net.IP

	addrs, _ := iface.Addrs()

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				v4 = append(v4, ipnet.IP)
			} else {
				switch ip := ipnet.IP.To16(); ip != nil {
				case ip.IsGlobalUnicast():
					v6 = append(v6, ipnet.IP)
				case ip.IsLinkLocalUnicast():
					v6local = append(v6local, ipnet.IP)
				}
			}
		}
	}
	if len(v6) == 0 {
		v6 = v6local
	}
	return v4, v6
}
