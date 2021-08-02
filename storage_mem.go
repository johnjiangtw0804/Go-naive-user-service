package main

import (
	"fmt"
	"log"
)

type StorageMemory struct {
	users []User
}

func (s *StorageMemory) GetUser(id string) (User, error) {
	for _, user := range s.users {
		if user.ID == id {
			log.Println("Returned user with id: ", id)
			return user, nil
		}
	}
	log.Println("User with id: ", id, " not found")
	return User{ID: "", AGE: 0}, fmt.Errorf("no user found")
}
func (s *StorageMemory) CreateUser(user User) error {
	if user.AGE < 0 {
		return fmt.Errorf("age cant be negative")
	}
	s.users = append(s.users, user)
	log.Println("User with id: ", user.ID, " created")
	return nil
}

func (s *StorageMemory) DeleteUser(id string) error {
	for index, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:index], s.users[index+1:]...)
			log.Println("User with id: ", id, " deleted")
			return nil
		}
	}
	log.Println("User with id: ", id, " not found")
	return fmt.Errorf("no user found to delete")
}
