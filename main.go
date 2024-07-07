package main

import (
	"log"
)

func main() {
	// pgstore, err := NewPostgresStore()
	tempStore, err := NewTempStore()
	if err != nil {
		log.Fatalf("P#1XK2NT: %v", err)
	}

	// err = pgstore.init()
	err = tempStore.init()

	if err != nil {
		log.Fatalf("P#1XK2P9: %v", err)
	}

	server := NewAPIServer(":8080", tempStore)

	server.Run()
}
