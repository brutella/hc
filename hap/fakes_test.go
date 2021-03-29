package hap

import (
	"github.com/brutella/hc/util"
	"io"
	"net"
	"time"
)

var testConn net.Conn = &fakeConn{}

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

type fakeCryptographer struct {
}

func (f fakeCryptographer) Encrypt(r io.Reader) (io.Reader, error) {
	return nil, nil
}

func (f fakeCryptographer) Decrypt(r io.Reader) (io.Reader, error) {
	return nil, nil
}

type fakeContainerHandler struct {
}

func (f fakeContainerHandler) Handle(container util.Container) (util.Container, error) {
	return nil, nil
}

type fakePairVerifyHandler struct {
}

func (p fakePairVerifyHandler) Handle(container util.Container) (util.Container, error) {
	return nil, nil
}

func (p fakePairVerifyHandler) SharedKey() [32]byte {
	return [32]byte{}
}
