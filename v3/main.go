package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"v3/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProduct(l)
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  1 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	// ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil,
	// which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:
	go func() {
		if err := s.ListenAndServe(); err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("Received signal", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
