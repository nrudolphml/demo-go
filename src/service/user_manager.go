package service

import (
	"errors"
	"github.com/nrudolph/twitter/src/domain"
)

var users []*domain.User

func AddUser(username string, email string, nickname string, password string) (*domain.User, error) {
	if username == "" {
		return nil, errors.New("username is empty")
	}
	if email == "" {
		return nil, errors.New("email is empty")
	}
	if nickname == "" {
		return nil, errors.New("nickname is empty")
	}
	if password == "" {
		return nil, errors.New("password is empty")
	}
	if checkIfUserExists(username, email) {
		return nil, errors.New("el usuario ya existe")
	}
	user := domain.NewUser(username, email, password, nickname)
	users = append(users, user)
	return user, nil
}

func checkIfUserExists(username string, email string) bool {
	for _, v := range users {
		if username == v.Username && email == v.Email {
			return true
		}
	}
	return false
}

func GetUser(username string) (*domain.User, error) {
	for _, v := range users {
		if username == v.Username {
			return v, nil
		}
	}
	return nil, errors.New("no user was found")
}
