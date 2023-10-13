package handler



// var users = mocks.GetMockUsers()

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	userJSON, err := json.Marshal(users)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Error("Failed to marshal users")
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(userJSON)
// }

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	segments := strings.Split(r.URL.Path, "/")

// 	if len(segments) != 3 {
// 		http.NotFound(w, r)
// 		log.Error("Invalid URL")
// 		return
// 	}

// 	userID, err := uuid.Parse(segments[2])
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		log.Error("Invalid user ID")
// 		return
// 	}

// 	var user User
// 	for _, u := range users {
// 		if u.Id == userID {
// 			user = User(u)
// 			break
// 		}
// 	}

// 	if user.Id == uuid.Nil {
// 		userJSON, err := json.Marshal(user)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			log.Error("Failed to marshal user")
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(userJSON)
// 	} else {
// 		http.NotFound(w, r)
// 		log.Error("User not found")
// 		return
// 	}
// }

