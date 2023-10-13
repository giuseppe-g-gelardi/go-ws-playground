package routes

import (
	"net/http"

	"playground.com/m/handler"

	"github.com/charmbracelet/log"
)

func LoadRoutes() {
	http.HandleFunc("/", handler.HelloHandler) // ol' faithful

	userRoutes()

	// ...otherRoutes()
	
	log.Info("Successfully loaded routes")
}

func userRoutes() {
	http.HandleFunc("/users", handler.QueryDBUsers)
	http.HandleFunc("/user", handler.QueryUser)
	http.HandleFunc("/newuser", handler.NewUser)
}

/** as an example
** func otherRoutes() {
** 	 http.HandleFunc("/posts", handler.PostsHandler)
** }
***/

