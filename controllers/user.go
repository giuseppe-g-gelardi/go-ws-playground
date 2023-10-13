package controllers

import (
	"context"

	"playground.com/m/database"
	"playground.com/m/types"

	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserFromDatabase(id string) (types.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error("Invalid User ID format")
		return types.User{}, err
	}

	var result types.User
	collection := database.Client.Database("go-ws").Collection("users")

	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		log.Error("User not found")
		return types.User{}, err
	}

	return result, nil

}
