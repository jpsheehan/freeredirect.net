package main

import (
	"database/sql"
	"net/http"
)

// NewServer creates a new Server struct
func NewServer() Server {
	var s Server

	s.db = dbGet()
	s.router = http.NewServeMux()
	s.routes()

	return s
}

// Server handle things
type Server struct {
	db     *sql.DB
	router *http.ServeMux
}

func (s *Server) routes() {
	s.router.HandleFunc("/", s.handleDefaultHosts())
	s.router.HandleFunc("freeredirect.net/", s.handleStaticDocument(""))
}
