package main

import (
	"context"
	"fmt"
	"github.com/brutella/dnssd"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	if resp, err := dnssd.NewResponder(); err != nil {
		fmt.Println(err)
	} else {
		srv1 := dnssd.NewService("Service 1", "_asdf._tcp.", "local.", "My Computer", []net.IP{net.ParseIP("192.168.0.123")}, 12332)
		srv2 := dnssd.NewService("Service 2", "_asdf._tcp.", "local.", "", nil, 45614)

		go func() {
			stop := make(chan os.Signal, 1)
			signal.Notify(stop, os.Interrupt)

			select {
			case <-stop:
				cancel()
			}
		}()

		h1, err := resp.Add(srv1)

		if err != nil {
			fmt.Println(err)
			return
		}

		go func() {
			time.Sleep(5 * time.Second)
			h2, _ := resp.Add(srv2)

			time.Sleep(5 * time.Second)
			h2.UpdateText(map[string]string{"key1": "value1", "key2": "value2"}, resp)
			h1.UpdateText(map[string]string{"key3": "value3", "key4": "value4"}, resp)

			time.Sleep(5 * time.Second)
			h2.UpdateText(map[string]string{"key3": "value3", "key4": "value4"}, resp)

			time.Sleep(10 * time.Second)
			resp.Remove(h2)
		}()

		err = resp.Respond(ctx)

		if err != nil {
			fmt.Println(err)
		}
	}
}
