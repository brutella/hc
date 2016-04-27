package hap

import (
	"github.com/brutella/log"
	"github.com/oleksandr/bonjour"

	"fmt"
	"os"
	"strings"
)

// MDNSService represents a mDNS service.
type MDNSService struct {
	config *Config
	server *bonjour.Server
}

// NewMDNSService returns a new service based for the bridge name, id and port.
func NewMDNSService(config *Config) *MDNSService {
	return &MDNSService{
		config: config,
	}
}

// IsPublished returns true when the service is published.
func (s *MDNSService) IsPublished() bool {
	return s.server != nil
}

// Publish announces the service for the machine's ip address on a random port using mDNS.
func (s *MDNSService) Publish() error {
	// Host should end with '.'
	hostname, _ := os.Hostname()
	host := fmt.Sprintf("%s.", strings.Trim(hostname, "."))
	text := s.config.txtRecords()

	// 2016-03-14(brutella): Replace whitespaces (" ") from service name
	// with underscores ("_")to fix invalid http host header field value
	// produces by iOS.
	//
	// [Radar] http://openradar.appspot.com/radar?id=4931940373233664
	stripped := strings.Replace(s.config.name, " ", "_", -1)

	server, err := bonjour.RegisterProxy(stripped, "_hap._tcp.", "", s.config.servePort, host, s.config.IP, text, nil)
	if err != nil {
		log.Fatal(err)
	}

	s.server = server
	return err
}

// Update updates the mDNS txt records.
func (s *MDNSService) Update() {
	if s.server != nil {
		txt := s.config.txtRecords()
		s.server.SetText(txt)
		log.Println("[INFO]", txt)
	}
}

// Stop stops the running mDNS service.
func (s *MDNSService) Stop() {
	s.server.Shutdown()
	s.server = nil
}
