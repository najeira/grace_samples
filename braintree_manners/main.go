package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/lestrrat/go-server-starter/listener"

	"github.com/najeira/grace_samples/handler"
)

func handleSignal() {
	signal_chan := make(chan os.Signal, 10)
	signal.Notify(signal_chan, syscall.SIGTERM)
	go func() {
		for {
			s := <-signal_chan
			if s == syscall.SIGTERM {
				manners.Close()
			}
		}
	}()
}

func main() {
	listeners, err := listener.ListenAll()
	if err != nil && err != listener.ErrNoListeningTarget {
		log.Fatal(err)
	}
	handleSignal()
	manners.Serve(listeners[0], http.HandlerFunc(handler.Handler))
}
