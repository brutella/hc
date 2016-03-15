package hap

import (
	"github.com/brutella/log"
	"github.com/gosexy/to"
	"github.com/oleksandr/bonjour"

	"fmt"
	"os"
	"strings"
)

// MDNSService represents a mDNS service.
type MDNSService struct {
	name               string
	ip                 string
	port               int
	protocol           string // Protocol version (pv) (Default 1.0)
	id                 string
	configuration      int64 // c#
	state              int64 // s#
	mfiCompliant       bool  // ff
	reachable          bool  // sf
	categoryIdentifier int64 // ci (see AccessoryType)

	server *bonjour.Server
}

// NewMDNSService returns a new service based for the bridge name, id and port.
func NewMDNSService(name string, id string, ip string, port int, category int64) *MDNSService {
	return &MDNSService{
		name:               name,
		ip:                 ip,
		port:               port,
		protocol:           "1.0",
		id:                 id,
		configuration:      1,
		state:              1,
		mfiCompliant:       false,
		reachable:          true,
		categoryIdentifier: category,
	}
}

// IsPublished returns true when the service is published.
func (s *MDNSService) IsPublished() bool {
	return s.server != nil
}

func (s *MDNSService) SetReachable(r bool) {
	s.reachable = r
}

// Publish announces the service for the machine's ip address on a random port using mDNS.
func (s *MDNSService) Publish() error {
	// Host should end with '.'
	hostname, _ := os.Hostname()
	host := fmt.Sprintf("%s.", strings.Trim(hostname, "."))
	text := s.txtRecords()

	// 2016-03-14(brutella): Remove whitespaces from service name to fix
	// invalid http host header field value produces by iOS.
	//
	// [Radar] http://openradar.appspot.com/radar?id=4931940373233664
	stripped := strings.Replace(s.name, " ", "", -1)

	server, err := bonjour.RegisterProxy(stripped, "_hap._tcp.", "", s.port, host, s.ip, text, nil)
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

func (s *MDNSService) txtRecords() []string {
	return []string{
		fmt.Sprintf("pv=%s", s.protocol),
		fmt.Sprintf("id=%s", s.id),
		fmt.Sprintf("c#=%d", s.configuration),
		fmt.Sprintf("s#=%d", s.state),
		fmt.Sprintf("sf=%d", to.Int64(s.reachable)),
		fmt.Sprintf("ff=%d", to.Int64(s.mfiCompliant)),
		fmt.Sprintf("md=%s", s.name),
		fmt.Sprintf("ci=%d", s.categoryIdentifier),
	}
}
