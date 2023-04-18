package main

import (
	"log"
	"net/http"
)

func main() {
	app := NewApp()

	err := http.ListenAndServe(":8080", app)
	if err != nil {
		log.Fatal("error in the server", err)
	}
}
