package app

import (
    "github.com/armon/mdns"
    "github.com/gosexy/to"
    "fmt"
    "log"
    "net"
    "errors"
)

type Service struct {
    name string
    port int
    protocol string         // Protocol version (pv) (Default 1.0)
    id string
    configuration int64     // c#
    state int64             // s#
    mfiCompliant bool       // ff
    status int64            // sf
    
    server *mdns.Server
}

func NewService(name, id string, port int) *Service {
    return &Service{
        name: name,
        port: port,
        protocol: "1.0",
        id: id,
        configuration: 1,
        state: 1,
        mfiCompliant: false,
        status: 1,
    }
}

func (s *Service) Publish() error {
    ip, err := GetFirstLocalIPAddress()
    if err != nil {
        return err
    }
    
    service, err := mdns.NewMDNSService(s.name, "_hap._tcp.", "", "", s.port, []net.IP{ip}, s.txtRecords())
    if err != nil {
        log.Fatal(err)
    }
    
    server, err := mdns.NewServer(&mdns.Config{Zone: service})
    if err != nil {
        log.Fatal(err)
    }
    
    s.server = server
    return err
}

func (s *Service) Update() {
}

func (s *Service) Stop() {
    s.server.Shutdown()
    s.server = nil
}

// Returns the first available IP address of the local machine
// This is a fix for Beaglebone Black where net.LookupIP(hostname)
// return no IP address
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

func (s *Service) txtRecords() []string {
    return []string{
        fmt.Sprintf("pv=%s", s.protocol),
        fmt.Sprintf("id=%s", s.id),
        fmt.Sprintf("c#=%d", s.configuration),
        fmt.Sprintf("s#=%d", s.state),
        fmt.Sprintf("sf=%d", s.status),
        fmt.Sprintf("ff=%d", to.Int64(s.mfiCompliant)),
        fmt.Sprintf("md=%s", s.name),
    }
}