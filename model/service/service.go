package service

import (
	"example-auth/model"
	"example-auth/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func ValidateUser(u model.User) bool {
	queryUser := repository.Get(u)

	if queryUser.Username != "" {
		err := bcrypt.CompareHashAndPassword([]byte(queryUser.Password), []byte(u.Password))
		if err != nil {
			log.Println("Invalid user")
			log.Println(err)
			return false
		}
		return true
	}
	return false
}

func NewUser(u model.User) {
	repository.Save(u)
}
