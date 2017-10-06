package main

import (
	"context"
	"fmt"
	"github.com/brutella/dnssd"
	"os"
	"os/signal"
	"time"
)

var timeFormat = "15:04:05.000"

func main() {
	fmt.Printf("Debugging…\n")
	fmt.Printf("DATE: –––%s–––\n", time.Now().Format("Mon Jan 2 2006"))
	fmt.Printf("%s	...STARTING...\n", time.Now().Format(timeFormat))

	fn := func(req *dnssd.Request) {
		fmt.Println("-------------------------------------------")
		fmt.Printf("%s	%v\n", time.Now().Format(timeFormat), req)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if s, err := dnssd.NewResponder(); err != nil {
		fmt.Println(err)
	} else {
		s.Debug(ctx, fn)

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)

		select {
		case <-stop:
			cancel()
		}
	}
}
