package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/koolvn/study-go.git/handlers"
)

func main() {
	var host = flag.String("host", "0.0.0.0", "host to listen on")
	var port = flag.String("port", "8080", "port to listen on")
	flag.Parse()

	addr := *host + ":" + *port
	log.Printf("Listening on http://%s\n", addr)
	mux := createMux()
	log.Fatal(http.ListenAndServe(addr, mux))
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /favicon.ico",
		func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
		})
	mux.HandleFunc("GET /", handlers.ServeRoot)
	mux.HandleFunc("POST /image", handlers.ReceiveImage)

	return mux
}
