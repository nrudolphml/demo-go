package user

type User struct {
	Username, Email, password, Nickname string
}

func NewUser(username string, email string, password string, nickname string) *User {
	u := User{username, email, password, nickname}
	return &u
}

func (user *User) CheckPassword(password string) bool {
	return user.password == password
}

func (user *User) IsUser(identification string) bool {
	return user.Username == identification || user.Email == identification || user.Nickname == identification
}
