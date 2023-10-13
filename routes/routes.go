package routes

import (
	"net/http"

	"github.com/charmbracelet/log"
	"playground.com/m/handler"
)

func init() {
	log.Info("Successfully loaded routes")
}

func LoadRoutes() {
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/users", handler.QueryDBUsers)
}
