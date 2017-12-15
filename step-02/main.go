package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	apiURL, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatalf("Invalid api url (Error: %s).", err)
	}

	router := mux.NewRouter()

	// Proxy all requests to the API
	router.HandleFunc("/api/{wildcard:.*}", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(apiURL)
		proxy.ServeHTTP(w, r)
	})

	// Serve static assets
	fileServer := http.FileServer(http.Dir("."))
	router.Handle("/{wildcard:.*}", fileServer)

	log.Fatal(http.ListenAndServe(":3000", router))
}
