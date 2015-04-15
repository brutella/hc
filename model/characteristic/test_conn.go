package characteristic

import (
	"net"
	"time"
)

var TestConn net.Conn = &fakeConn{}

type fakeConn struct {
}

func (f *fakeConn) Read(b []byte) (n int, err error) {
	return 0, nil
}

func (f *fakeConn) Write(b []byte) (n int, err error) {
	return 0, nil
}

func (f *fakeConn) Close() error {
	return nil
}

func (f *fakeConn) LocalAddr() net.Addr {
	return nil
}

func (f *fakeConn) RemoteAddr() net.Addr {
	return nil
}

func (f *fakeConn) SetDeadline(t time.Time) error {
	return nil
}

func (f *fakeConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (f *fakeConn) SetWriteDeadline(t time.Time) error {
	return nil
}
