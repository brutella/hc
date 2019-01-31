// Command browse browses for specific dns-sd service types.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/brutella/dnssd"
	"os"
	"os/signal"
	"strings"
	"time"
)

var serviceFlag = flag.String("Type", "_hap._tcp", "Service type")
var domainFlag = flag.String("Domain", "local.", "Browsing domain")

var timeFormat = "15:04:05.000"

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := fmt.Sprintf("%s.%s.", strings.Trim(*serviceFlag, "."), strings.Trim(*domainFlag, "."))

	fmt.Printf("Browsing for %s\n", service)
	fmt.Printf("DATE: –––%s–––\n", time.Now().Format("Mon Jan 2 2006"))
	fmt.Printf("%s  ...STARTING...\n", time.Now().Format(timeFormat))
	fmt.Printf("Timestamp	A/R	Domain	Service Type	Service Name\n")

	addFn := func(srv dnssd.Service) {
		fmt.Printf("%s	Add	%s	%s	%s\n", time.Now().Format(timeFormat), srv.Domain, srv.Type, srv.Name)
	}

	rmvFn := func(srv dnssd.Service) {
		fmt.Printf("%s	Rmv	%s	%s	%s\n", time.Now().Format(timeFormat), srv.Domain, srv.Type, srv.Name)
	}

	if err := dnssd.LookupType(ctx, service, addFn, rmvFn); err != nil {
		fmt.Println(err)
		return
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	select {
	case <-stop:
		cancel()
	}
}
