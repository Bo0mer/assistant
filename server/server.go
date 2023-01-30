package server

import (
	"fmt"
	"net/http"
)

// Server is the main type that implements the http.Handler interface.
type Server struct {
}

// NewServer returns a new initialized server.
func NewServer() *Server {
	return &Server{}
}

// ServeHTTP handles HTTP reqeusts by assisting the remote user.
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, How can I assist you?")
}
