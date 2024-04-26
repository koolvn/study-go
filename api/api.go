package api

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	mux.HandleFunc("GET /", AuthPageHandler)
	mux.HandleFunc("POST /auth", AuthHandler)

	log.Printf("Listening on http://%v\n", s.addr)

	return http.ListenAndServe(s.addr, mux)
}
