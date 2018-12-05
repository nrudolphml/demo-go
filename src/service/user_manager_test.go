package service

import "testing"

func TestAddUser(t *testing.T) {
	err := AddUser("pepe", "pepe@pepe.com", "pepe", "pepe123")
	if err != nil {
		t.Error("No error was expected")
	}
}

func TestAddUserWithExistingUsernameReturnsError(t *testing.T) {
	_ = AddUser("pepe", "pepe@pepe.com", "pepe", "pepe123")
	err := AddUser("pepe", "pepe2@pepe.com", "pepe2", "pepe1231")
	if err != nil && err.Error() != "el usuario ya existe" {
		t.Error("Error was expected")
	}
}

func TestAddUserWithExistingEmailReturnsError(t *testing.T) {
	_ = AddUser("pepe", "pepe@pepe.com", "pepe", "pepe123")
	err := AddUser("pepe2", "pepe@pepe.com", "pepe2", "pepe1231")
	if err != nil && err.Error() != "el usuario ya existe" {
		t.Error("Error was expected")
	}
}

func TestAddUserWithoutUsernameReturnsError(t *testing.T) {
	err := AddUser("", "pepe@pepe.com", "pepe", "pepe123")
	if err != nil && err.Error() != "username is empty" {
		t.Error("No error was expected")
	}
}

func TestAddUserWithoutEmailReturnsError(t *testing.T) {
	err := AddUser("pepe", "", "pepe", "pepe123")
	if err != nil && err.Error() != "email is empty" {
		t.Error("No error was expected")
	}
}

func TestAddUserWithoutNicknameReturnsError(t *testing.T) {
	err := AddUser("pepe", "pepe@pepe.com", "", "pepe123")
	if err != nil && err.Error() != "nickname is empty" {
		t.Error("No error was expected")
	}
}

func TestAddUserWithoutPasswordReturnsError(t *testing.T) {
	err := AddUser("pepe", "pepe@pepe.com", "pepe", "")
	if err != nil && err.Error() != "password is empty" {
		t.Error("No error was expected")
	}
}
