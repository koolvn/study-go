package main

import (
	"flag"
	"github.com/koolvn/study-go.git/api"
	"log"
)

func main() {
	host := flag.String("host", "0.0.0.0", "Host to listen on")
	port := flag.String("port", "8080", "Port to listen on")
	cert := flag.String("ssl-cert", "", "Path to SSL certificate")
	key := flag.String("ssl-key", "", "Path to SSL key")
	flag.Parse()

	https := false
	if *cert != "" && *key != "" {
		https = true
	}
	server := api.NewAPI(*host, *port, https, *cert, *key)
	err := server.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
