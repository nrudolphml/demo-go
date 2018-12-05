package service

import "testing"

func TestAddUser(t *testing.T) {
	_, err := AddUser("pepe", "pepe@pepe.com", "pepe", "pepe123")
	if err != nil {
		t.Error("No error was expected")
	}
}

func TestAddUserWithExistingUsernameReturnsError(t *testing.T) {
	_, _ = AddUser("pepe", "pepe@pepe.com", "pepe", "pepe123")
	_, err := AddUser("pepe", "pepe2@pepe.com", "pepe2", "pepe1231")
	if err != nil && err.Error() != "el usuario ya existe" {
		t.Error("Error was expected")
	}
}

func TestAddUserWithExistingEmailReturnsError(t *testing.T) {
	_, _ = AddUser("pepe", "pepe@pepe.com", "pepe", "pepe123")
	_, err := AddUser("pepe2", "pepe@pepe.com", "pepe2", "pepe1231")
	if err != nil && err.Error() != "el usuario ya existe" {
		t.Error("Error was expected")
	}
}

func TestAddUserWithoutUsernameReturnsError(t *testing.T) {
	_, err := AddUser("", "pepe@pepe.com", "pepe", "pepe123")
	if err != nil && err.Error() != "username is empty" {
		t.Error("No error was expected")
	}
}

func TestAddUserWithoutEmailReturnsError(t *testing.T) {
	_, err := AddUser("pepe", "", "pepe", "pepe123")
	if err != nil && err.Error() != "email is empty" {
		t.Error("No error was expected")
	}
}

func TestAddUserWithoutNicknameReturnsError(t *testing.T) {
	_, err := AddUser("pepe", "pepe@pepe.com", "", "pepe123")
	if err != nil && err.Error() != "nickname is empty" {
		t.Error("No error was expected")
	}
}

func TestAddUserWithoutPasswordReturnsError(t *testing.T) {
	_, err := AddUser("pepe", "pepe@pepe.com", "pepe", "")
	if err != nil && err.Error() != "password is empty" {
		t.Error("No error was expected")
	}
}

func TestGetUserFromUsername(t *testing.T) {
	_, _ = AddUser("pepe", "pepe@pepe.com", "pepe", "ppp")
	_, _ = AddUser("manolo", "pepe2@pepe.com", "pepe", "ppp")
	_, _ = AddUser("juan carlos", "pepe3@pepe.com", "pepe", "ppp")

	user, _ := GetUser("pepe")

	if user.Username != "pepe" {
		t.Error("No user was found with username 'pepe'")
	}
}

func TestGetUserFromUsernameReturnsErrorForInexistingUsername(t *testing.T) {
	_, _ = AddUser("pepe", "pepe@pepe.com", "pepe", "ppp")
	_, _ = AddUser("manolo", "pepe2@pepe.com", "pepe", "ppp")
	_, _ = AddUser("juan carlos", "pepe3@pepe.com", "pepe", "ppp")

	user, err := GetUser("p")

	if user != nil {
		t.Error("a user was found, it was expected not to find users")
	}
	if err != nil && err.Error() != "no user was found" {
		t.Error("An error was expected")
	}
}