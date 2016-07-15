package main

import (
	"log"
	"net/http"
	"time"

	"github.com/lestrrat/go-server-starter/listener"
	"gopkg.in/tylerb/graceful.v1"

	"github.com/najeira/grace_samples/handler"
)

func main() {
	listeners, err := listener.ListenAll()
	if err != nil && err != listener.ErrNoListeningTarget {
		log.Fatal(err)
	}
	graceful.Serve(
		&http.Server{Handler: http.HandlerFunc(handler.Handler)},
		listeners[0],
		time.Second*10)
}
