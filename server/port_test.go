package server

import (            
	"testing"
    "github.com/stretchr/testify/assert"
)
type testAddr struct {
    addr string
}

func NewAddr(addr string) testAddr {
    return testAddr{addr: addr}
}

func (a testAddr) Network() string {
    return "foo"
}
func (a testAddr) String() string {
    return a.addr
}

func TestPortFromAddr(t *testing.T) {
    port := ExtractPort(NewAddr("[::]:12345"))
    assert.Equal(t, port, "12345")
}