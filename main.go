package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// init mux router
	// type inference :=
	r := mux.NewRouter()
	// Route handler / Endpoints
	log.Println("Server starting")
	r.HandleFunc("/user/{user_id}", getUser).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user/{user_id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
