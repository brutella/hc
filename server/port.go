package server

import(
    "net"
    "strings"
)

func ExtractPort(addr net.Addr) string {
    comps := strings.Split(addr.String(), ":")
    return comps[len(comps) - 1]
}