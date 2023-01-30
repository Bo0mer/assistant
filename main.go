package main

import (
	"log"
	"net/http"

	"github.com/Bo0mer/assistant/server"
)

func main() {
	srv := server.NewServer()

	err := http.ListenAndServe(":9191", srv)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
