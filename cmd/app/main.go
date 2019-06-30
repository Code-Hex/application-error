package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Code-Hex/application-error/internal/application"
	"github.com/Code-Hex/application-error/internal/handler"
	"github.com/Code-Hex/application-error/internal/health"
	"github.com/Code-Hex/application-error/internal/storage"
)

var version string

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "application run failed: %v", err)
		os.Exit(1)
	}
}

func run() error {
	port := 8080

	log.Printf("starting port: %d", port)

	h := handler.New()

	h.Handle("/_ah/health", health.NewHandler())

	s := &storage.InMemStore{}

	h.Handle("/get", application.NewGethandler(s))
	h.Handle("/put", application.NewPuthandler(s))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), h); err != nil {
		return err
	}
	return nil
}
