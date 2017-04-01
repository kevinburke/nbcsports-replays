package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/kevinburke/handlers"
	nbcsports "github.com/kevinburke/nbcsports-replays"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ln, err := net.Listen("tcp", "127.0.0.1:8965")
	if err != nil {
		handlers.Logger.Error("Error listening on port", "err", err)
		os.Exit(2)
	}
	http.Handle("/", nbcsports.IndexServer)
	ctx, cancel := context.WithCancel(context.Background())
	go nbcsports.FetchBody(ctx)

	handlers.Logger.Info("Listening on port 8965...")
	go func() {
		http.Serve(ln, http.DefaultServeMux)
	}()
	<-c
	cancel()
}
