package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct (our models)
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
			fmt.Println(params["user_id"])
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)

	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, user := range users {
		if user.ID == params["user_id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}
