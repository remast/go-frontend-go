package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UI server up and running!")
}

func main() {
	// Use default API url or read from environment variable
	apiURLParam := "http://localhost:8080"
	if os.Getenv("API_URL") != "" {
		apiURLParam = os.Getenv("API_URL")
	}

	apiURL, err := url.Parse(apiURLParam)
	if err != nil {
		log.Fatalf("Invalid api url (Error: %s).", err)
	}
	log.Printf("Starting UI server for %s ...\n", apiURL)

	router := mux.NewRouter()

	// Health check
	router.HandleFunc("/health", healthCheckHandler)

	// Proxy all requests to the API
	router.HandleFunc("/api/{wildcard:.*}", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(apiURL)
		proxy.ServeHTTP(w, r)
	})

	// Serve static assets (gzipped)
	fileServer := http.FileServer(http.Dir("."))
	fileServerWithGz := gziphandler.GzipHandler(fileServer)
	router.Handle("/{wildcard:.*}", fileServerWithGz)

	log.Fatal(http.ListenAndServe(":3000", router))
}
