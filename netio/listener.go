package netio

import (
	"net"
)

// HAPTCPListener listens for new connection and creates HAPConnections for new connections
type HAPTCPListener struct {
	*net.TCPListener
	context HAPContext
}

// NewHAPTCPListener returns a new hap tcp listener.
func NewHAPTCPListener(l *net.TCPListener, context HAPContext) *HAPTCPListener {
	return &HAPTCPListener{l, context}
}

// Accept creates and returns a HAPConnection.
func (l *HAPTCPListener) Accept() (c net.Conn, err error) {
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
