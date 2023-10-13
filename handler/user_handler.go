package handler

import (
	"encoding/json"
	"net/http"

	uc "playground.com/m/controllers"
	e "playground.com/m/errors"

	"github.com/charmbracelet/log"
)

type RequestBody struct {
	User_id string `json:"_id"`
}

func NewUser(w http.ResponseWriter, r *http.Request) {

	var requestBody uc.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	e.BadRequest(w, err, http.StatusBadRequest, "Invalid JSON request")
	// if err != nil {
	// 	http.Error(w, "Invalid JSON request", http.StatusBadRequest)
	// 	return
	// }

	user, err := uc.CreateNewUser(requestBody)
	e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to create new user")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	log.Error("Failed to create new user")
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to encode user")
	} // err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	log.Error("Failed to encode user")
	// 	return
	// }

	log.Printf("User: %v", user)
}

func QueryUser(w http.ResponseWriter, r *http.Request) {

	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	e.BadRequest(w, err, http.StatusBadRequest, "Invalid JSON request")
	// if err != nil {
	// 	http.Error(w, "Invalid JSON request", http.StatusBadRequest)
	// 	return
	// }

	userID := requestBody.User_id

	user, err := uc.GetUserFromDatabase(userID)
	e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to get user from database")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	log.Error("Failed to get user from database")
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to encode user")
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// log.Error("Failed to encode user")
		// return
	}

	log.Printf("User: %v", user)
}

func QueryDBUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.GetAllUsersFromDatabase()
	e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to get users from database")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	log.Error("Failed to get users from database")
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to encode users")
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// log.Error("Failed to encode users")
		// return
	}
}
