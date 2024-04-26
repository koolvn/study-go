package api

import (
	"log"
	"net/http"
)

type AuthServer struct {
	addr     string
	certPath string
	keyPath  string
}

func NewAuthServer(addr string, certPath string, keyPath string) *AuthServer {
	return &AuthServer{
		addr:     addr,
		certPath: certPath,
		keyPath:  keyPath,
	}
}

func (s *AuthServer) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	mux.HandleFunc("GET /", AuthPageHandler)
	mux.HandleFunc("POST /auth", AuthHandler)

	log.Printf("Auth server is listening on https://%v\n", s.addr)

	return http.ListenAndServeTLS(s.addr, s.certPath, s.keyPath, mux)
}
