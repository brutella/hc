package hap

import (
	"os"
	"os/signal"
)

// TermFunc defines the function which is executed on termination.
type TermFunc func()

// OnTermination calls a function when the app receives an interrupt of kill signal.
func OnTermination(fn TermFunc) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	go func() {
		select {
		case <-c:
			if fn != nil {
				fn()
			}
		}
	}()
}
