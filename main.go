package main

import (
	"devbook-api/pkg/database"
	"devbook-api/pkg/server"
	"log"
)

func main() {

	// Start the database
	database.StartDB()

	// Start the server
	s := server.NewServer("3000")

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
