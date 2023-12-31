package server

import (
	"net/http"

	e "playground.com/m/errors"
	"playground.com/m/types"

	"github.com/charmbracelet/log"
)

type Server types.Server

func NewServer(port string) *Server {
	return &Server{Port: port}
}

func (s *Server) StartServer() {
	port := ":" + s.Port
	log.Infof("Starting server on port: %v", s.Port)
	err := http.ListenAndServe(port, nil)
	e.Fatal(err, "Failed to start server")
}
