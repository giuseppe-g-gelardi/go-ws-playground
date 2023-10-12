package main

import (
	"os"

	"playground.com/m/routes"
	"playground.com/m/server"
	"playground.com/m/database"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

var (
	port string
	uri string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Info("Successfully loaded .env file")

	port = os.Getenv("PORT")
	uri = os.Getenv("MONGO_URI")

	routes.LoadRoutes()
	database.NewDatabase(uri).Connect()
	
}

func main() {

	s := server.NewServer(port)
	s.StartServer()
}
