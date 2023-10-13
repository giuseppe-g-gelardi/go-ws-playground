package controllers

import (
	"context"

	"playground.com/m/types"
	e "playground.com/m/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User types.User

func CreateNewUser(user User) (User, error) {
	db := Collection("users")

	result, err := db.InsertOne(context.Background(), user)
	e.Err(err, "Failed to insert user")
	// if err != nil {
	// 	log.Error(err, "Failed to insert user")
	// 	return User{}, err
	// }

	objectID := result.InsertedID.(primitive.ObjectID)
	user.Id = objectID.Hex()

	return user, nil
}

// gets a single user (by _id) from the database
func GetUserFromDatabase(id string) (types.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	e.Err(err, "Invalid User ID format")
	// if err != nil {
	// 	log.Error(err, "Invalid User ID format")
	// 	return types.User{}, err
	// }

	var result types.User
	db := Collection("users")
	err = db.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&result)
	e.Err(err, "User not found")
	// if err != nil {
	// 	log.Error(err, "User not found")
	// 	return types.User{}, err
	// }

	return result, nil
}

// gets all users from the database
func GetAllUsersFromDatabase() ([]User, error) {
	db := Collection("users")

	var results []User
	cursor, err := db.Find(context.Background(), bson.D{{}})
	e.Err(err, "Failed to find users")
	// if err != nil {
	// 	log.Error(err, "Failed to find users")
	// 	return nil, err
	// }
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		e.Err(err, "Failed to decode user")
		// if err != nil {
		// 	log.Fatal(err, "Failed to decode user")
		// 	return nil, err
		// }
		results = append(results, user)
	}

	if err := cursor.Err(); err != nil {
		e.Err(err, "Failed to iterate over users")
		// log.Fatal(err, "Failed to iterate over users")
		// return nil, err
	}

	return results, nil
}
