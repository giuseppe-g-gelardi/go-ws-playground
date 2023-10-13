package controllers

import (
	"context"

	"playground.com/m/database"
	"playground.com/m/types"

	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User types.User

func user_collection() *mongo.Collection {
	return database.Client.Database("go-ws").Collection("users")
}




// gets a single user (by _id) from the database
func GetUserFromDatabase(id string) (types.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error("Invalid User ID format")
		return types.User{}, err
	}

	var result types.User
	db := user_collection()
	// collection := database.Client.Database("go-ws").Collection("users")
	// err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&result)
	err = db.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		log.Error("User not found")
		return types.User{}, err
	}

	return result, nil
}

// gets all users from the database
func GetAllUsersFromDatabase() ([]User, error) {
	// collection := database.Client.Database("go-ws").Collection("users")
	db := user_collection()

	var results []User
	// cursor, err := collection.Find(context.Background(), bson.D{{}})
	cursor, err := db.Find(context.Background(), bson.D{{}})

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
