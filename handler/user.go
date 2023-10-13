package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"playground.com/m/database"
	"playground.com/m/types"

	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/bson"
)

type User types.User

func QueryDBUsers(w http.ResponseWriter, r *http.Request) {
	users, err := GetAllUsersFromDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Failed to get users from database -- it keeps breaking here too")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Failed to encode users")
		return
	}
}

func GetAllUsersFromDatabase() ([]User, error) {
	collection := database.Client.Database("go-ws").Collection("users")

	var results []User
	// cursor, err := collection.Find(context.TODO(), bson.D{})
	// cursor, err := collection.Find(context.TODO(), filter)
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Error(err, "Failed to find users -- it keeps breaking here")
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

// func NewUser(w http.ResponseWriter, r *http.Request) *User {
// 	var user User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		log.Error("Failed to decode user")
// 		return nil
// 	}

// 	// user.Id = uuid.New()
// 	// add the new user to the database and return the user

// 	return &user
// }
