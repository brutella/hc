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

var instanceFlag = flag.String("Name", "A", "Service Name")
var serviceFlag = flag.String("Type", "_hap._tcp.", "Service type")
var domainFlag = flag.String("Domain", "local.", "Browsing domain")

var timeFormat = "15:04:05.000"

func main() {
	if len(*instanceFlag) == 0 {
		flag.Usage()
		return
	}
	service := fmt.Sprintf("%s.%s.", strings.Trim(*serviceFlag, "."), strings.Trim(*domainFlag, "."))
	instance := fmt.Sprintf("%s.%s.%s.", strings.Trim(*instanceFlag, "."), strings.Trim(*serviceFlag, "."), strings.Trim(*domainFlag, "."))

	fmt.Printf("Lookup %s\n", instance)
	fmt.Printf("DATE: –––%s–––\n", time.Now().Format("Mon Jan 2 2006"))
	fmt.Printf("%s	...STARTING...\n", time.Now().Format(timeFormat))

	resolve := func() {
		ctx, cancel := dnssd.NewResolveContext()
		defer cancel()

		if srv, err := dnssd.LookupInstance(ctx, instance); err != nil {
			fmt.Println("Could not resolve", err)
		} else {
			fmt.Printf("%s	%s can be reached at %s:%d\n%v\n", time.Now().Format(timeFormat), srv.ServiceInstanceName(), srv.Hostname, srv.Port, srv.Text)
		}
	}

	addFn := func(srv dnssd.Service) {
		if srv.ServiceInstanceName() != instance {
			return
		}
		go resolve()
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := dnssd.LookupType(ctx, service, addFn, func(dnssd.Service) {}); err != nil {
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
