package main

import "log"

func main() {
	store, err := NewTempStore()

	if err != nil {
		log.Fatal(err)
	}

	store.init()

	server := NewAPIServer(":8080", store)

	server.Run()
}
