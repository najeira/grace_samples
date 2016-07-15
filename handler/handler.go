package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	sleepDuration = time.Millisecond * 100
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "going to sleep %s with pid %d\n", sleepDuration, os.Getpid())
	w.(http.Flusher).Flush()
	time.Sleep(sleepDuration)
	fmt.Fprintf(w, "slept %s with pid %d\n", sleepDuration, os.Getpid())
}
