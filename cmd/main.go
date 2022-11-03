package main

import (
	"context"
	"firstMicroservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server start for go

func main() {
	l := log.New(os.Stdout, "firstMicroservice", log.LstdFlags)
	hh := handlers.NewRoot(l)
	gh := handlers.NewHello(l)
	mux := http.NewServeMux()
	mux.Handle("/", hh)
	mux.Handle("/hello", gh)
	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	// wrapping ListenAndServe in gofunc so it's not going to block
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// make a new channel to notify on os interrupt of server (ctrl + C)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// This blocks the code until the channel receives some message
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// Once message is consumed shut everything down
	// Gracefully shuts down all client requests. Makes server more reliable
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
