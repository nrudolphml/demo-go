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

func Equals(user1 *User, user2 *User) bool {
	return user1.Email == user2.Email && user1.Username == user2.Username
}
