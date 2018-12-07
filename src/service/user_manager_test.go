package service_test

import (
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/domain/user"
	"github.com/nrudolph/twitter/src/service"
	"testing"
)

func TestAddUser(t *testing.T) {
	_ = service.NewTweetManager()
	userManager := service.NewUserManager()
	_, err := userManager.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe123")
	if err != nil {
		t.Error("No error was expected")
	}
}

func TestAddUserWithExistingUsernameReturnsError(t *testing.T) {
	_ = service.NewTweetManager()
	userManager := service.NewUserManager()
	_, _ = userManager.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe123")
	_, err := userManager.AddUser("pepe", "pepe2@pepe.com", "pepe2", "pepe1231")
	if err != nil && err.Error() != "el usuario ya existe" {
		t.Error("Error was expected")
	}
}

func TestAddUserWithExistingEmailReturnsError(t *testing.T) {
	_ = service.NewTweetManager()
	userManager := service.NewUserManager()
	_, _ = userManager.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe123")
	_, err := userManager.AddUser("pepe2", "pepe@pepe.com", "pepe2", "pepe1231")
	if err != nil && err.Error() != "el usuario ya existe" {
		t.Error("Error was expected")
	}
}

func TestAddUserWithoutUsernameReturnsError(t *testing.T) {
	_ = service.NewTweetManager()
	userManager := service.NewUserManager()
	_, err := userManager.AddUser("", "pepe@pepe.com", "pepe", "pepe123")
	if err != nil && err.Error() != "username is empty" {
		t.Error("No error was expected")
	}
}

func TestAddUserWithoutEmailReturnsError(t *testing.T) {
	_ = service.NewTweetManager()
	userManager := service.NewUserManager()
	_, err := userManager.AddUser("pepe", "", "pepe", "pepe123")
	if err != nil && err.Error() != "email is empty" {
		t.Error("No error was expected")
	}
}

func TestAddUserWithoutNicknameReturnsError(t *testing.T) {
	_ = service.NewTweetManager()
	userManager := service.NewUserManager()
	_, err := userManager.AddUser("pepe", "pepe@pepe.com", "", "pepe123")
	if err != nil && err.Error() != "nickname is empty" {
		t.Error("No error was expected")
	}
}

func TestAddUserWithoutPasswordReturnsError(t *testing.T) {
	_ = service.NewTweetManager()
	userManager := service.NewUserManager()
	_, err := userManager.AddUser("pepe", "pepe@pepe.com", "pepe", "")
	if err != nil && err.Error() != "password is empty" {
		t.Error("No error was expected")
	}
}

func TestGetUserFromUsername(t *testing.T) {
	_ = service.NewTweetManager()
	userManager := service.NewUserManager()
	_, _ = userManager.AddUser("pepe", "pepe@pepe.com", "pepe", "ppp")
	_, _ = userManager.AddUser("manolo", "pepe2@pepe.com", "pepe", "ppp")
	_, _ = userManager.AddUser("juan carlos", "pepe3@pepe.com", "pepe", "ppp")

	user, _ := userManager.GetUser("pepe")

	if user.Username != "pepe" {
		t.Error("No user was found with username 'pepe'")
	}
}

func TestGetUserFromUsernameReturnsErrorForInexistingUsername(t *testing.T) {
	_ = service.NewTweetManager()
	userManager := service.NewUserManager()
	_, _ = userManager.AddUser("pepe", "pepe@pepe.com", "pepe", "ppp")
	_, _ = userManager.AddUser("manolo", "pepe2@pepe.com", "pepe", "ppp")
	_, _ = userManager.AddUser("juan carlos", "pepe3@pepe.com", "pepe", "ppp")

	user, err := userManager.GetUser("p")

	if user != nil {
		t.Error("a user was found, it was expected not to find users")
	}
	if err != nil && err.Error() != "no user was found" {
		t.Error("An error was expected")
	}
}

func TestLoginUser(t *testing.T) {
	_ = service.NewTweetManager()
	userManager := service.NewUserManager()
	user, _ := userManager.AddUser("pepe", "pepe@pepe.com", "pepe", "ppp")
	_, _ = userManager.LoginUser(user.Username, "ppp")

	if !userManager.IsUserLoggedIn(user) {
		t.Error("user could not be logged in")
	}
}

func TestLogout(t *testing.T) {
	_ = service.NewTweetManager()
	userManager := service.NewUserManager()
	user, _ := userManager.AddUser("pepe", "pepe@pepe.com", "pepe", "ppp")
	_, _ = userManager.LoginUser(user.Username, "ppp")

	userManager.LogoutUser(user)

	if userManager.IsUserLoggedIn(user) {
		t.Error("user could not be logged out")
	}
}

func TestDeleteTweet(t *testing.T) {

	// init
	_ = service.NewUserManager()
	tweetManager := service.NewTweetManager()

	newUser := user.NewUser("p", "p", "p", "p")

	tweet := domain.NewTextTweet(newUser, "SuperTweet")

	i, _ := tweetManager.PublishTweet(tweet)

	// operation

	if _, err := tweetManager.DeleteTweet(newUser, i); err != nil {
		t.Errorf("error al eliminar : %s\n", err.Error())
	}

	if tweetManager.CountTweetsByUser(newUser) != 0 {
		t.Errorf("Was expected 0 but was %d\n", tweetManager.CountTweetsByUser(newUser))
	}

}
