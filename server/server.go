package server

import (
	"net/http"

	"playground.com/m/types"

	"github.com/charmbracelet/log"
)

type Server types.Server

func NewServer(port string) *Server {
	return &Server{Port: port}
}

func (s *Server) StartServer() {
	port := ":" + s.Port
	log.Info("Starting server on port:", s.Port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err, "Failed to start server")
	}
}
