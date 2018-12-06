package user

type User struct {
	Username, Email, password, Nickname string
}

func NewUser(username string, email string, password string, nickname string) *User {
	u := User{username, email, password, nickname}
	return &u
}

func CheckPassword(user *User, password string) bool {
	return user.password == password
}

func IsUser(user *User, identification string) bool {
	return user.Username == identification || user.Email == identification || user.Nickname == identification
}
