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

var instanceFlag = flag.String("Name", "My Service", "Service Name")
var serviceFlag = flag.String("Type", "_http._tcp.", "Service type")
var domainFlag = flag.String("Domain", "local.", "Browsing domain")

var timeFormat = "15:04:05.000"

func main() {
	flag.Parse()
	if len(*instanceFlag) == 0 || len(*serviceFlag) == 0 || len(*domainFlag) == 0 {
		flag.Usage()
		return
	}
	service := fmt.Sprintf("%s.%s.", strings.Trim(*serviceFlag, "."), strings.Trim(*domainFlag, "."))
	instance := fmt.Sprintf("%s.%s.%s.", strings.Trim(*instanceFlag, "."), strings.Trim(*serviceFlag, "."), strings.Trim(*domainFlag, "."))

	fmt.Printf("Lookup %s\n", instance)
	fmt.Printf("DATE: –––%s–––\n", time.Now().Format("Mon Jan 2 2006"))
	fmt.Printf("%s	...STARTING...\n", time.Now().Format(timeFormat))

	addFn := func(srv dnssd.Service) {
		if srv.ServiceInstanceName() == instance {
			text := ""
			for key, value := range srv.Text {
				text += fmt.Sprintf("%s=%s", key, value)
			}
			fmt.Printf("%s	%s can be reached at %s:%d\n	%v\n", time.Now().Format(timeFormat), srv.ServiceInstanceName(), srv.Hostname(), srv.Port, text)
		}
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
