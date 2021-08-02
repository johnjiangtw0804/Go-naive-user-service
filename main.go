package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/user/jonathan/Go-naive-user-servive/handlers"
	"github.com/user/jonathan/Go-naive-user-servive/models"
)

// init mux router
var router = mux.NewRouter()

func init() {
	var err error

	err = models.NewStorage(models.Memory)
	if err != nil {
		log.Fatal(err)
	}
	router.HandleFunc("/user/{user_id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{user_id}", handlers.DeleteUser).Methods("DELETE")
}

func main() {
	// Route handler / Endpoints
	log.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}
