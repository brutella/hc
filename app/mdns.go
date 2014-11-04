package app

import (
    "github.com/armon/mdns"
    "github.com/gosexy/to"
    "fmt"
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
    }
}

func (s *Service) Publish() error {
    fmt.Println(s.txtRecords())
    service, err := mdns.NewMDNSService(s.name, "_hap._tcp.", "", "", s.port, nil, s.txtRecords())
    server, err := mdns.NewServer(&mdns.Config{Zone: service})
    
    if err == nil {
        s.server = server
    }
    
    return err
}

func (s *Service) Update() {
    // TODO publish new txt records
}

func (s *Service) Stop() {
    s.server.Shutdown()
    s.server = nil
}

func (s *Service) txtRecords() []string {
    return []string{
        fmt.Sprintf("pv=%s", s.protocol),
        fmt.Sprintf("id=%s", s.id),
        fmt.Sprintf("c#=%d", s.configuration),
        fmt.Sprintf("s#=%d", s.state),
        fmt.Sprintf("sf=%d", s.state),
        fmt.Sprintf("ff=%d", to.Int64(s.mfiCompliant)),
        fmt.Sprintf("md=%s", s.name),
    }
}