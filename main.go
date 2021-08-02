package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/user/jonathan/Go-naive-user-servive/handlers"
)

// init mux router
var router = mux.NewRouter()
var database Storage

func init() {
	var err error

	// pass in 1 => use memory as storage space
	database, err = NewStorage(Memory)
	if err != nil {
		log.Fatal(err)
	}
	router.HandleFunc("/user/{user_id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{user_id}", handlers.DeleteUser).Methods("DELETE")
}

func main() {
	// Route handler / Endpoints
	log.Println("Server starting... listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
