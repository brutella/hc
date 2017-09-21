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

var instanceFlag = flag.String("Name", "Service", "Service name")
var serviceFlag = flag.String("Type", "_asdf._tcp.", "Service type")
var domainFlag = flag.String("Domain", "local.", "domain")
var portFlag = flag.Int("Port", 12345, "Port")

var timeFormat = "15:04:05.000"

func main() {
	flag.Parse()
	if len(*instanceFlag) == 0 || len(*serviceFlag) == 0 || len(*domainFlag) == 0 {
		flag.Usage()
		return
	}

	instance := fmt.Sprintf("%s.%s.%s.", strings.Trim(*instanceFlag, "."), strings.Trim(*serviceFlag, "."), strings.Trim(*domainFlag, "."))

	fmt.Printf("Registering Service %s port %d\n", instance, *portFlag)
	fmt.Printf("DATE: –––%s–––\n", time.Now().Format("Mon Jan 2 2006"))
	fmt.Printf("%s	...STARTING...\n", time.Now().Format(timeFormat))

	ctx, cancel := context.WithCancel(context.Background())

	if resp, err := dnssd.NewResponder(); err != nil {
		fmt.Println(err)
	} else {
		srv := dnssd.NewService(*instanceFlag, *serviceFlag, *domainFlag, "", nil, *portFlag)

		go func() {
			stop := make(chan os.Signal, 1)
			signal.Notify(stop, os.Interrupt)

			select {
			case <-stop:
				cancel()
			}
		}()

		go func() {
			time.Sleep(100 * time.Millisecond)
			handle, err := resp.Add(srv)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%s	Got a reply for service %s: Name now registered and active\n", time.Now().Format(timeFormat), handle.Service().ServiceInstanceName())
			}
		}()
		err = resp.Respond(ctx)

		if err != nil {
			fmt.Println(err)
		}
	}
}
