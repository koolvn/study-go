package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

const defaultHost = "0.0.0.0"
const defaultPort = "28082"
const apiTarget = "api.openai.com"

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// ReverseProxyHandler forwards requests to `apiTarget` and logs the request and response headers.
func ReverseProxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] Received a request from %s, request header: %v\n", r.RemoteAddr, r.Header)

	director := func(req *http.Request) {
		req.URL.Scheme = "https"
		req.URL.Host = apiTarget
		req.Host = apiTarget
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)

	log.Printf("[INFO] Received the destination website response header: %v\n", w.Header())
}

func main() {
	host := getEnv("OPENAI_PROXY_HOST", defaultHost)
	port := getEnv("OPENAI_PROXY_PORT", defaultPort)
	certFile := getEnv("OPENAI_PROXY_CERT_PATH", "")
	keyFile := getEnv("OPENAI_PROXY_CERT_KEY_PATH", "")

	log.Printf("[INFO] PID: %d, PPID: %d", os.Getpid(), os.Getppid())

	servingAddr := fmt.Sprintf("%s:%s", host, port)
	if certFile != "" && keyFile != "" {
		log.Printf("[INFO] Starting server at https://%s\n", servingAddr)
		if err := http.ListenAndServeTLS(servingAddr, certFile, keyFile, http.HandlerFunc(ReverseProxyHandler)); err != nil {
			log.Fatalf("[ERROR] Failed to start TLS server: %v\n", err)
		}
	} else {
		log.Printf("[INFO] Starting server at http://%s\n", servingAddr)
		if err := http.ListenAndServe(servingAddr, http.HandlerFunc(ReverseProxyHandler)); err != nil {
			log.Fatalf("[ERROR] Failed to start server: %v\n", err)
		}
	}
}
