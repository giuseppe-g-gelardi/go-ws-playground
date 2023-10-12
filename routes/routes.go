package routes

import (
	"net/http"

	"github.com/charmbracelet/log"
	"playground.com/m/handler"
)

func init() {
	log.Info("Routes loaded")
}

func LoadRoutes() {
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/users", handler.GetUsers)
	http.HandleFunc("/users/", handler.GetUser)
}
