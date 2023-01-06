package main

import (
	"log"

	"github.com/paveltyukin/practicum-go-service-devops/internal/server"
	"github.com/paveltyukin/practicum-go-service-devops/internal/server/storage"
)

func main() {
	strg := storage.New()
	err := server.Serve("127.0.0.1:8080", strg)
	if err != nil {
		log.Fatal(err)
	}
}
