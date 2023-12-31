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

	// ! check if the user already exists in the database

	user, err := uc.CreateNewUser(requestBody)
	e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to create new user")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to encode user")
	}
	log.Infof("User: %v", user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	e.BadRequest(w, err, http.StatusBadRequest, "Invalid JSON request")

	userID := requestBody.User_id

	user, err := uc.DeleteUser(userID)
	e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to delete user")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to encode user")
	}
	log.Infof("User: %v", user)
}

func QueryUser(w http.ResponseWriter, r *http.Request) {

	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	e.BadRequest(w, err, http.StatusBadRequest, "Invalid JSON request")

	userID := requestBody.User_id

	user, err := uc.GetUserFromDatabase(userID)
	e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to get user from database")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to encode user")
	}
	log.Infof("User: %v", user)
}

func QueryUserByUsername(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]string //username
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	e.BadRequest(w, err, http.StatusBadRequest, "Invalid JSON request")

	username := requestBody["username"]

	user, err := uc.GetUserByUsername(username)
	e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to get user from database")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to encode user")
	}
	log.Infof("User: %v", user)
}

func QueryUserAndUpdate(w http.ResponseWriter, r *http.Request) {
	var requestBody uc.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	e.BadRequest(w, err, http.StatusBadRequest, "Invalid JSON request")

	userID := r.URL.Query().Get("userid")

	// ? check if there is an existing user, if not, create a new user

	updatedUser, err := uc.UpdateUser(userID, requestBody)
	e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to update user")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&updatedUser); err != nil {
		e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to encode user")
	}

	log.Infof("UpdatedUser: %v", updatedUser)
}

func QueryDBUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.GetAllUsersFromDatabase()
	e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to get users from database")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		e.InternalServerError(w, err, http.StatusInternalServerError, "Failed to encode users")
	}
	log.Infof("Users: %v", users)
}
