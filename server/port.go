package server

import (
	"net"
	"strings"
)

// ExtractPort returns the address's port as string
func ExtractPort(addr net.Addr) string {
	comps := strings.Split(addr.String(), ":")
	return comps[len(comps)-1]
}
