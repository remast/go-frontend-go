package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Serve static assets
	fileServer := http.FileServer(http.Dir("."))
	router.Handle("/{wildcard:.*}", fileServer)

	log.Fatal(http.ListenAndServe(":3000", router))
}
