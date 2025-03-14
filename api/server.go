package api

import (
	"net/http"
)

type Server struct {
	mux http.ServeMux
}

func NewServer() *Server {
	server := new(Server)
	server.mux = *http.NewServeMux()

	return server
}

// func (s *Server) HandleFunc(path string, handle())
