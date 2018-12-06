package service

import (
	"errors"
	"github.com/nrudolph/twitter/src/domain/user"
)

var users []*user.User

func AddUser(username string, email string, nickname string, password string) (*user.User, error) {
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
	newUser := user.NewUser(username, email, password, nickname)
	users = append(users, newUser)
	return newUser, nil
}

func checkIfUserExists(username string, email string) bool {
	for _, v := range users {
		if username == v.Username && email == v.Email {
			return true
		}
	}
	return false
}

func GetUser(username string) (*user.User, error) {
	for _, v := range users {
		if username == v.Username {
			return v, nil
		}
	}
	return nil, errors.New("no user was found")
}
