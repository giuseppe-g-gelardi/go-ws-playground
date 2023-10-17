package controllers

import (
	"context"

	e "playground.com/m/errors"
	"playground.com/m/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User types.User

// creates a new user in the database
func CreateNewUser(user User) (User, error) {
	db := Collection("users")

	result, err := db.InsertOne(context.Background(), user)
	e.Err(err, "Failed to insert user")

	objectID := result.InsertedID.(primitive.ObjectID)
	user.Id = objectID.Hex()

	return user, nil
}

// deletes a single user (by _id) from the database
func DeleteUser(id string) (User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	e.Err(err, "Invalid User ID format")

	var result User
	db := Collection("users")
	err = db.FindOneAndDelete(context.Background(), bson.M{"_id": objectID}).Decode(&result)
	e.Err(err, "User not found")

	return result, nil
}

// gets a single user (by _id) from the database
func GetUserFromDatabase(id string) (User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	e.Err(err, "Invalid User ID format")

	var result User
	db := Collection("users")
	err = db.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&result)
	e.Err(err, "User not found")

	return result, nil
}

func GetUserByUsername(username string) (User, error) {

	var result User
	db := Collection("users")
	err := db.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	e.Err(err, "User not found")

	return result, nil
}

// updates a single user (by _id) from the database
func UpdateUser(id string, user User) (interface{}, error) {

	objectID, err := primitive.ObjectIDFromHex(user.Id)
	e.Err(err, "Invalid User ID format")

	db := Collection("users")
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	update := bson.M{
		"$set": bson.M{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"username":   user.Username,
			"email":      user.Email,
			"password":   user.Password,
		},
	}

	var result User

	db_err := db.FindOneAndUpdate(context.Background(), bson.M{"_id": objectID}, update, opts).Decode(&result)
	if db_err != nil {
		return User{}, db_err
	}

	return result, nil
}

// gets all users from the database
func GetAllUsersFromDatabase() ([]User, error) {
	db := Collection("users")

	var results []User
	cursor, err := db.Find(context.Background(), bson.D{{}})
	e.Err(err, "Failed to find users")

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		e.Err(err, "Failed to decode user")
		results = append(results, user)
	}

	if err := cursor.Err(); err != nil {
		e.Err(err, "Failed to iterate over users")
	}

	return results, nil
}
