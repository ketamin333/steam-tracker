package main

import (
	"log"

	"go.rest.api/internal/app"
)

func main() {
	s, err := app.Bootstrap()

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
