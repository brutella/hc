package http

import (
	"github.com/brutella/hc/hap"

	"net"
)

func (s *Server) Accept() (con net.Conn, err error) {
	con, err = s.listener.AcceptTCP()
	if err != nil {
		return
	}

	hapCon := hap.NewConnection(con, s.context)

	return hapCon, err
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func (s *Server) Addr() net.Addr {
	return s.listener.Addr()
}
