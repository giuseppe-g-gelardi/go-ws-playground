package handler

import (
	"encoding/json"
	"net/http"

	uc "playground.com/m/controllers"

	"github.com/charmbracelet/log"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "playground.com/m/database"
	// "playground.com/m/types"
)

type RequestBody struct {
	User_id string `json:"_id"`
}

func QueryUser(w http.ResponseWriter, r *http.Request) {

	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	userID := requestBody.User_id
	// log.Warn("userid", userID)




	// // Convert the string ID into an ObjectID
	// objectID, err := primitive.ObjectIDFromHex(userID)
	// if err != nil {
	// 	http.Error(w, "Invalid User ID format", http.StatusBadRequest)
	// 	return
	// }

	// var result types.User
	// collection := database.Client.Database("go-ws").Collection("users")

	// err = collection.FindOne(r.Context(), bson.M{"_id": objectID}).Decode(&result)
	// if err != nil {
	// 	http.Error(w, "User not found", http.StatusNotFound)
	// 	log.Error("Failed to get user from the database")
	// 	return
	// }

	// log.Warn("user result", result)


	user, err := uc.GetUserFromDatabase(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Failed to get user from database")
		return
	}





	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Failed to encode user")
		return
	}

	log.Printf("User: %v", user)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(userID))

}
