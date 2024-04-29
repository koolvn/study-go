package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var host = getEnv("OPENAI_PROXY_HOST", "0.0.0.0")
var port = getEnv("OPENAI_PROXY_PORT", "28082")
var certFile = getEnv("OPENAI_PROXY_CERT_PATH", "")
var keyFile = getEnv("OPENAI_PROXY_CERT_KEY_PATH", "")

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func ReverseProxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] receive a request from %s, request header:\n %s: \n", r.RemoteAddr, r.Header)
	target := "api.openai.com"
	director := func(req *http.Request) {
		req.URL.Scheme = "https"
		req.URL.Host = target
		req.Host = target
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
	log.Printf("[*] receive the destination website response header: %s\n", w.Header())
}

func main() {
	log.Printf("[INFO] PID: %d PPID: %d\n", os.Getpid(), os.Getppid())
	servingAddr := fmt.Sprintf("%s:%s", host, port)
	if certFile != "" && keyFile != "" {
		log.Printf("[INFO] Starting server at  https://%s\n", servingAddr)
		err := http.ListenAndServeTLS(
			servingAddr,
			certFile,
			keyFile,
			http.HandlerFunc(ReverseProxyHandler))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("[INFO] Starting server at  http://%s\n", servingAddr)
		err := http.ListenAndServe(servingAddr, http.HandlerFunc(ReverseProxyHandler))
		if err != nil {
			log.Fatal(err)
		}
	}
}
