package database

import (
	"context"

	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

type Database struct {
	URI    string
}

func NewDatabase(URI string) *Database {
	return &Database{URI: URI}
}

func (db *Database) Connect() {
	d := NewDatabase(db.URI)

	// set the stable api to version 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(d.URI).SetServerAPIOptions(serverAPI)

	// create a new client and connect to the server
	var err error
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err, "Failed to connect to MongoDB")
	}

	// defer func() {
	// 	if err := Client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatal(err, "Failed to disconnect from MongoDB")
	// 	}
	// }()

	// send ping to confirm a successful connection
	if err := Client.Database("admin").RunCommand(context.TODO(), bson.D{{
		Key:   "ping",
		Value: 1,
	}}).Err(); err != nil {
		log.Fatal(err, "Failed to ping MongoDB")
	}
	log.Info("Successfully connected to MongoDB")
}

