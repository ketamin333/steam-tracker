package main

import (
	"log"

	"go.rest.api/internal/app"
	"go.rest.api/internal/db"
)

func main() {
	db.Init()
	s := app.Bootstrap()

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
