package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	nbcsports "github.com/kevinburke/nbcsports-replays"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	http.Handle("/", nbcsports.IndexServer)
	cleanupDone := make(chan bool, 1)
	go nbcsports.Fetch()

	go func() {
		http.ListenAndServe("127.0.0.1:8965", nil)
	}()
	go func() {
		time.Sleep(30 * time.Millisecond)
		fmt.Println("Listening on port 8965...")
	}()
	<-cleanupDone
}
