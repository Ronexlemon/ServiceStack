package main

import (
	"log"
	"net/http"
)

const (
	httpAddr = ":9090"
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoute(mux)
	log.Printf("Startig HTTP server at %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("failed to start server")
	}

}
