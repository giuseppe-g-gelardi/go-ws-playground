package main

import (
	"os"

	"playground.com/m/database"
	"playground.com/m/routes"
	"playground.com/m/server"
	e "playground.com/m/errors"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

var (
	port string
	uri  string
	db   *database.Database
)

func init() {
	// load the .env
	err := godotenv.Load()
	e.Fatal(err, "Error loading .env file")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	log.Info("Successfully loaded .env file")

	// set the port and uri
	port = os.Getenv("PORT")
	uri = os.Getenv("MONGO_URI")

}

func main() {
	// load route handlers
	routes.LoadRoutes()

	// connect to DB before starting server or things will break
	db = database.NewDatabase(uri)
	db.Connect()

	// start server
	s := server.NewServer(port)
	s.StartServer()
}
