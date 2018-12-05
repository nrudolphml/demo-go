package service

import (
	"errors"
	"github.com/nrudolph/twitter/src/domain"
)

var users []*domain.User

func AddUser(username string, email string, nickname string, password string) error {
	if username == "" {
		return errors.New("username is empty")
	}
	if email == "" {
		return errors.New("email is empty")
	}
	if nickname == "" {
		return errors.New("nickname is empty")
	}
	if password == "" {
		return errors.New("password is empty")
	}
	if checkIfUserExists(username, email) {
		return errors.New("el usuario ya existe")
	}
	users = append(users, domain.NewUser(username, email, password, nickname))
	return nil
}

func checkIfUserExists(username string, email string) bool {
	for _, v := range users {
		if username == v.Username && email == v.Email {
			return true
		}
	}
	return false
}
