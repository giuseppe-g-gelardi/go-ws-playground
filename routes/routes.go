package routes

import (
	"net/http"

	"playground.com/m/handler"

	"github.com/charmbracelet/log"
)

func LoadRoutes() {
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/users", handler.QueryDBUsers)
	http.HandleFunc("/user", handler.QueryUser)
	
	log.Info("Successfully loaded routes")
}
