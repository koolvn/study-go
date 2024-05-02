package main

import (
	"log"

	"github.com/koolvn/study-go.git/api"
	"github.com/koolvn/study-go.git/config"
)

func main() {
	cfg := config.NewConfig()
	server := api.NewAPI(*cfg)
	err := server.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
