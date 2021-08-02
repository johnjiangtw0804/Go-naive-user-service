package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct (our user)
type User struct {
	ID  string `json:"user_id"`
	AGE int32  `json:"age"`
}

var users []User

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, user := range users {
		if user.ID == params["user_id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	// No user found
	w.WriteHeader(http.StatusNoContent)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.AGE < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users = append(users, user)

	w.WriteHeader(http.StatusNoContent)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, user := range users {
		if user.ID == params["user_id"] {
			users = append(users[:index], users[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}
