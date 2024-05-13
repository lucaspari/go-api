package types

import (
	"log"
	"net/http"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s Server) Run(mux *http.ServeMux) error {
	server := http.Server{
		Addr:    s.port,
		Handler: mux,
	}
	log.Printf("Listening on port %s", s.port)
	return server.ListenAndServe()
}
