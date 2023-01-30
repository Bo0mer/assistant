package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server is the main type that implements the http.Handler interface.
type Server struct {
	r *mux.Router
}

// NewServer returns a new initialized server.
func NewServer() *Server {
	s := &Server{}
	r := mux.NewRouter()
	s.r = r
	r.Handle("/weather", http.HandlerFunc(s.handleGetWeather))
	r.Handle("/meetings", http.HandlerFunc(s.handleGetMeetings))
	return s
}

// ServeHTTP handles HTTP reqeusts by assisting the remote user.
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.r.ServeHTTP(w, req)
}

func (s *Server) handleGetWeather(w http.ResponseWriter, req *http.Request) {

}

func (s *Server) handleGetMeetings(w http.ResponseWriter, req *http.Request) {

}
