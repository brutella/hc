package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/brutella/dnssd"
	"github.com/brutella/dnssd/log"
	"net"
	"os"
	"os/signal"
	"strings"
)

func main() {
	log.Debug.Enable()

	if resp, err := dnssd.NewResponder(); err != nil {
		fmt.Println(err)
	} else {

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			for {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Enter name \nor\n\"exit\"\n>")
				name, _ := reader.ReadString('\n')
				name = strings.Trim(name, "\n")

				if name == "exit" {
					cancel()
					return
				}

				srv := dnssd.NewService(name, "_asdf._tcp.", "local.", "", []net.IP{net.ParseIP("192.168.0.123")}, 12345)
				h, _ := resp.Add(srv)

				<-stop
				resp.Remove(h)
			}
		}()

		resp.Respond(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}
}
