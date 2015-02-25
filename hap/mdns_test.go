package hap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMDNS(t *testing.T) {
	mdns := NewService("My MDNS Service", "1234", 5010)
	assert.Equal(t, mdns.txtRecords(), []string{
		"pv=1.0",
		"id=1234",
		"c#=1",
		"s#=1",
		"sf=1",
		"ff=0",
		"md=My MDNS Service",
	})
}
