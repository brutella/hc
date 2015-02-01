package netio

import (
	"net"
	_ "time"
)

// TCP listener listens for new connection and creates
// a new TCP HAP connection for new connections
type TCPHAPListener struct {
	*net.TCPListener
	context HAPContext
}

func NewTCPHAPListener(l *net.TCPListener, context HAPContext) *TCPHAPListener {
	return &TCPHAPListener{l, context}
}

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
