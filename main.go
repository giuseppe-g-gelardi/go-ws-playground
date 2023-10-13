package main

import (
	"os"

	"playground.com/m/database"
	e "playground.com/m/errors"
	"playground.com/m/routes"
	"playground.com/m/server"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

var (
	port string
	uri  string
	db   *database.Database
)

func init() {
	err := godotenv.Load() // load the .env
	e.Fatal(err, "Error loading .env file")
	log.Info("Successfully loaded .env file")

	// set the port and uri
	port = os.Getenv("PORT")
	uri = os.Getenv("MONGO_URI")

}

func main() {
	routes.LoadRoutes() // load route handlers

	db = database.NewDatabase(uri)
	db.Connect() // connect to DB before starting server or things will break

	s := server.NewServer(port)
	s.StartServer() // start server
}
