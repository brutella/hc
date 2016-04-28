package hap

import (
	"net"
)

// TCPListener listens for new connection and creates Connections for new connections
type TCPListener struct {
	*net.TCPListener
	context Context
}

// NewTCPListener returns a new hap tcp listener.
func NewTCPListener(l *net.TCPListener, context Context) *TCPListener {
	return &TCPListener{l, context}
}

// Accept creates and returns a Connection.
func (l *TCPListener) Accept() (c net.Conn, err error) {
	conn, err := l.AcceptTCP()
	if err != nil {
		return
	}

	// TODO(brutella) Check if we should use tcp keepalive
	// conn.SetKeepAlive(true)
	// conn.SetKeepAlivePeriod(3 * time.Minute)
	hapConn := NewConnection(conn, l.context)

	return hapConn, err
}
