package routes

import (
	"net/http"

	h "playground.com/m/handler"

	"github.com/charmbracelet/log"
)

func LoadRoutes() {
	http.HandleFunc("/", h.HelloHandler) // ol' faithful

	userRoutes()

	// ...otherRoutes()

	log.Info("Successfully loaded routes")
}

func userRoutes() {
	http.HandleFunc("/users", h.QueryDBUsers)
	http.HandleFunc("/user", h.QueryUser)
	http.HandleFunc("/newuser", h.NewUser)
	http.HandleFunc("/deleteuser", h.DeleteUser)
	http.HandleFunc("/username", h.QueryUserByUsername)
	http.HandleFunc("/updateuser", h.QueryUserAndUpdate)
}

/** as an example
** func otherRoutes() {
** 	 http.HandleFunc("/posts", handler.PostsHandler)
** }
***/
