package main

import (
	"log"

	"github.com/HellfastUSMC/go_api"
)

func main() {
	s := go_api.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
