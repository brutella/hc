package hap

import (
	"bytes"
	gocontext "context"
	"github.com/brutella/hc/log"
	"io/ioutil"
	"time"
)

// KeepAlive encapsulates sending notifications with no content to all
// connected clients. This way we can find abandoned connections and close them.
// Thise is also done in homebridge: https://github.com/KhaosT/HAP-NodeJS/blob/c3a8f989685b62515968278c81b86b744b968960/lib/HAPServer.js#L107
type KeepAlive struct {
	timeout time.Duration
	context Context
}

// NewKeepAlive returns a new keep alive for a specific timeout.
func NewKeepAlive(timeout time.Duration, context Context) *KeepAlive {
	k := KeepAlive{
		timeout: timeout,
		context: context,
	}

	return &k
}

// Start starts sending keep alive messages. This method blocks until the context is canceled.
func (k *KeepAlive) Start(ctx gocontext.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(k.timeout):
			k.sendKeepAlive()
		}
	}
}

func (k *KeepAlive) sendKeepAlive() {
	conns := k.context.ActiveConnections()
	var empty = new(bytes.Buffer)
	for _, conn := range conns {
		resp := NewNotification(empty)

		var buffer = new(bytes.Buffer)
		resp.Write(buffer)
		bytes, _ := ioutil.ReadAll(buffer)
		bytes = FixProtocolSpecifier(bytes)
		log.Debug.Printf("Keep alive %s <- %s", conn.RemoteAddr(), string(bytes))
		conn.Write(bytes)
	}
}
