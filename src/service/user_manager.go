package service

import (
	"errors"
	"github.com/nrudolph/twitter/src/domain/user"
)

var users []*user.User
var loggedUsers []*user.User

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
	if checkIfUserExists(username, email, nickname) {
		return nil, errors.New("el usuario ya existe")
	}
	newUser := user.NewUser(username, email, password, nickname)
	users = append(users, newUser)
	return newUser, nil
}

func checkIfUserExists(username string, email string, nickname string) bool {
	for _, v := range users {
		if username == v.Username || email == v.Email || nickname == v.Nickname {
			return true
		}
	}
	return false
}

func GetUser(identification string) (*user.User, error) {
	for _, v := range users {
		if user.IsUser(v, identification) {
			return v, nil
		}
	}
	return nil, errors.New("no user was found")
}

func LoginUser(identification string, password string) (bool, error) {
	currentUser, e := GetUser(identification)
	if e != nil {
		return false, e
	}
	if !user.CheckPassword(currentUser, password) {
		return false, errors.New("invalid credentials")
	}
	if isUserLogged(identification) {
		return false, errors.New("user is already logged in")
	}
	loggedUsers = append(loggedUsers, currentUser)
	return true, nil
}

func isUserLogged(identification string) bool {
	for _, v := range loggedUsers {
		if user.IsUser(v, identification) {
			return true
		}
	}
	return false
}

func IsUserLoggedIn(userToCheck *user.User) bool {
	return isUserLogged(userToCheck.Username)
}

func LogoutUser(userToLogout *user.User) bool {
	var index = -1
	for i, v := range loggedUsers {
		if v == userToLogout {
			index = i
			break
		}
	}
	if index == -1 {
		return false
	}
	loggedUsers = append(loggedUsers[:index], loggedUsers[index+1:]...)
	return true
}
