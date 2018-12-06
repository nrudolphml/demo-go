package service

import (
	"errors"
	"github.com/nrudolph/twitter/src/domain/user"
)

type UserManager struct {
	users       []*user.User
	loggedUsers []*user.User
}

func (userManager *UserManager) AddUser(username string, email string, nickname string, password string) (*user.User, error) {
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
	if userManager.checkIfUserExists(username, email, nickname) {
		return nil, errors.New("el usuario ya existe")
	}
	newUser := user.NewUser(username, email, password, nickname)
	userManager.users = append(userManager.users, newUser)
	return newUser, nil
}

func (userManager *UserManager) checkIfUserExists(username string, email string, nickname string) bool {
	for _, v := range userManager.users {
		if username == v.Username || email == v.Email || nickname == v.Nickname {
			return true
		}
	}
	return false
}

func (userManager *UserManager) GetUser(identification string) (*user.User, error) {
	for _, v := range userManager.users {
		if v.IsUser(identification) {
			return v, nil
		}
	}
	return nil, errors.New("no user was found")
}

func (userManager *UserManager) LoginUser(identification string, password string) (bool, error) {
	currentUser, e := userManager.GetUser(identification)
	if e != nil {
		return false, e
	}
	if !currentUser.CheckPassword(password) {
		return false, errors.New("invalid credentials")
	}
	if userManager.isUserLogged(identification) {
		return false, errors.New("user is already logged in")
	}
	userManager.loggedUsers = append(userManager.loggedUsers, currentUser)
	return true, nil
}

func (userManager *UserManager) isUserLogged(identification string) bool {
	for _, v := range userManager.loggedUsers {
		if v.IsUser(identification) {
			return true
		}
	}
	return false
}

func (userManager *UserManager) IsUserLoggedIn(userToCheck *user.User) bool {
	return userManager.isUserLogged(userToCheck.Username)
}

func (userManager *UserManager) LogoutUser(userToLogout *user.User) bool {
	var index = -1
	for i, v := range userManager.loggedUsers {
		if v == userToLogout {
			index = i
			break
		}
	}
	if index == -1 {
		return false
	}
	userManager.loggedUsers = append(userManager.loggedUsers[:index], userManager.loggedUsers[index+1:]...)
	return true
}

func NewUserManager() *UserManager {
	userManager := UserManager{make([]*user.User, 0), make([]*user.User, 0)}
	return &userManager
}
