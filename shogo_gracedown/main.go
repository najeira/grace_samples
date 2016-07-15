package main

import (
	"log"
	"net/http"
	"syscall"
	"os/signal"
	"os"

	"github.com/lestrrat/go-server-starter/listener"
	"github.com/shogo82148/go-gracedown"

	"github.com/najeira/grace_samples/handler"
)

func handleSignal() {
	signal_chan := make(chan os.Signal, 10)
	signal.Notify(signal_chan, syscall.SIGTERM)
	go func() {
		for {
			s := <-signal_chan
			if s == syscall.SIGTERM {
				gracedown.Close()
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
	gracedown.Serve(listeners[0], http.HandlerFunc(handler.Handler))
}
