// Program assistant starts a virtual assistant.
// Supported arguments of the program are:
//
//	-help Print help for running the program.
//
//	-port Specify the port to start the server on.
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
