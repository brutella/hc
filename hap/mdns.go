package hap

import (
	"github.com/brutella/log"
	"github.com/gosexy/to"
	"github.com/oleksandr/bonjour"

	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

// MDNSService represents a mDNS service.
type MDNSService struct {
	name               string
	port               int
	protocol           string // Protocol version (pv) (Default 1.0)
	id                 string
	configuration      int64 // c#
	state              int64 // s#
	mfiCompliant       bool  // ff
	status             int64 // sf
	categoryIdentifier int64 // ci; default 1 (Other)

	server *bonjour.Server
}

// NewMDNSService returns a new service based for the bridge name, id and port.
func NewMDNSService(name, id string, port int) *MDNSService {
	return &MDNSService{
		name:               name,
		port:               port,
		protocol:           "1.0",
		id:                 id,
		configuration:      1,
		state:              1,
		mfiCompliant:       false,
		status:             1,
		categoryIdentifier: 1,
	}
}

// IsPublished returns true when the HAP service is published.
func (s *MDNSService) IsPublished() bool {
	return s.server != nil
}

// Publish announces the HAP service for the machine's ip address on a random port used mDNS.
func (s *MDNSService) Publish() error {
	ip, err := GetFirstLocalIPAddress()
	if err != nil {
		return err
	}
	log.Println("[INFO] Bridge IP is", ip)

	// Host should end with '.'
	hostname, _ := os.Hostname()
	host := fmt.Sprintf("%s.", strings.Trim(hostname, "."))
	text := s.txtRecords()
	server, err := bonjour.RegisterProxy(s.name, "_hap._tcp.", "", s.port, host, ip.String(), text, nil)
	if err != nil {
		log.Fatal(err)
	}

	s.server = server
	return err
}

// Update updates the mDNS txt records.
func (s *MDNSService) Update() {
	if s.server != nil {
		s.server.SetText(s.txtRecords())
		log.Println("[INFO]", s.txtRecords())
	}
}

// Stop stops the running mDNS service.
func (s *MDNSService) Stop() {
	s.server.Shutdown()
	s.server = nil
}

// GetFirstLocalIPAddress returns the first available IP address of the local machine
// This is a fix for Beaglebone Black where net.LookupIP(hostname) return no IP address.
func GetFirstLocalIPAddress() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP, nil
			}
		}
	}

	return nil, errors.New("Could not determine ip address")
}

func (s *MDNSService) txtRecords() []string {
	return []string{
		fmt.Sprintf("pv=%s", s.protocol),
		fmt.Sprintf("id=%s", s.id),
		fmt.Sprintf("c#=%d", s.configuration),
		fmt.Sprintf("s#=%d", s.state),
		fmt.Sprintf("sf=%d", s.status),
		fmt.Sprintf("ff=%d", to.Int64(s.mfiCompliant)),
		fmt.Sprintf("md=%s", s.name),
		fmt.Sprintf("ci=%d", s.categoryIdentifier),
	}
}
