package main

import (
	"log"
	"net/http"

	"github.com/daalfox/bookstore/internal/api"
	"github.com/daalfox/bookstore/internal/spec"
	"github.com/go-chi/chi/v5"
)

func main() {
	bookstoreApi := api.NewBookstoreApi()
	r := chi.NewMux()
	h := spec.HandlerFromMux(&bookstoreApi, r)

	server := http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	log.Printf("starting server on %v\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to serve application: %v\n", err)
	}
}
