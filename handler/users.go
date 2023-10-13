package handler

import (
	"encoding/json"
	"net/http"

	uc "playground.com/m/controllers"

	"github.com/charmbracelet/log"
)

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

// func QueryUser(w http.ResponseWriter, r *http.Response) {

// 	log.Print("QueryUser called")

// 	w.Header().Set("Content-Type", "application/json")

// }

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
