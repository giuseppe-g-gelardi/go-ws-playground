package controllers

import (
	"context"

	"playground.com/m/database"
	"playground.com/m/types"

	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/bson"
)

type User types.User


func GetAllUsersFromDatabase() ([]User, error) {
	collection := database.Client.Database("go-ws").Collection("users")

	var results []User
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Error(err, "Failed to find users")
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err, "Failed to decode user")
			return nil, err
		}
		results = append(results, user)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err, "Failed to iterate over users")
		return nil, err
	}

	return results, nil
}
