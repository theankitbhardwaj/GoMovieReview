package main

import "log"

func main() {
	store, err := NewMyStore()

	if err != nil {
		log.Fatal(err)
	}

	store.setupDB()

	server := NewAPIServer(":8080", store)

	server.Run()
}
