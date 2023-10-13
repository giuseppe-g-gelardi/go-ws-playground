package controllers

import (
	"playground.com/m/database"
	"go.mongodb.org/mongo-driver/mongo"	
)

func Collection(collection string) *mongo.Collection {
	return database.Client.Database("go-ws").Collection(collection)
}
