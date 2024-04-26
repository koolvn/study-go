package main

import (
	"flag"
	"github.com/koolvn/study-go.git/api"
	"log"
)

func main() {
	var host = flag.String("host", "0.0.0.0", "Host to listen on")
	var port = flag.String("port", "8443", "Port to listen on")
	var cert = flag.String("cert", "certs/cert.crt", "Path to SSL certificate")
	var key = flag.String("key", "certs/cert.key", "Path to SSL certificate key")
	addr := *host + ":" + *port
	flag.Parse()

	server := api.NewAuthServer(addr, *cert, *key)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
