package main

import (
	"net/http"
	"time"

	"github.com/facebookgo/httpdown"

	"github.com/najeira/grace_samples/handler"
)

const (
	httpAddr = ":8080"
)

func main() {
	server := &http.Server{Addr: httpAddr, Handler: http.HandlerFunc(handler.Handler)}
	config := &httpdown.HTTP{
		StopTimeout: 10 * time.Second,
		KillTimeout: 1 * time.Second,
	}
	httpdown.ListenAndServe(server, config)
}
