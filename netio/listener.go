package netio

import (
	"net"
	_ "time"
)

// TCPHAPListener listens for new connection and creates hapConnections for new connections
type TCPHAPListener struct {
	*net.TCPListener
	context HAPContext
}

func NewTCPHAPListener(l *net.TCPListener, context HAPContext) *TCPHAPListener {
	return &TCPHAPListener{l, context}
}

// Accept creates and returns a hapConnection
func (l *TCPHAPListener) Accept() (c net.Conn, err error) {
	conn, err := l.AcceptTCP()
	if err != nil {
		return
	}

	// TODO(brutella) Check if we should use tcp keepalive
	// conn.SetKeepAlive(true)
	// conn.SetKeepAlivePeriod(3 * time.Minute)
	hapConn := NewHAPConnection(conn, l.context)

	return hapConn, err
}
