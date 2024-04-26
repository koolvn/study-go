package main

import (
	"github.com/koolvn/study-go.git/api"
	"log"
)

func main() {
	server := api.NewAPIServer("0.0.0.0:8080")
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
