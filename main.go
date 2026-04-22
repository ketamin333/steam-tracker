package main

import (
	"log"

	"go.rest.api/internal/db"
	"go.rest.api/internal/router"
	"go.rest.api/internal/server"
)

func main() {
	db.Init()

	r := router.Setup()
	s := server.NewServer(":8080", r)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
