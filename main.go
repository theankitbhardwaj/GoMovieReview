package main

import (
	"log"
)

func main() {
	pgstore, err := NewPostgresStore()
	if err != nil {
		log.Fatalf("P#1XK2NT: %v", err)
	}

	err = pgstore.init()

	if err != nil {
		log.Fatalf("P#1XK2P9: %v", err)
	}

	store, err := NewTempStore()

	if err != nil {
		log.Fatalf("P#1XK2PV: %v", err)
	}

	store.init()

	server := NewAPIServer(":8080", pgstore)

	server.Run()
}
