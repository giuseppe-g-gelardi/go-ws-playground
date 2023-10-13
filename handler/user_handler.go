package handler

import (
	"encoding/json"
	"net/http"

	uc "playground.com/m/controllers"

	"github.com/charmbracelet/log"
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
}

func QueryDBUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.GetAllUsersFromDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Failed to get users from database")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Failed to encode users")
		return
	}
}
