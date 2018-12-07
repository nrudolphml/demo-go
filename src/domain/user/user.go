package user

type User struct {
	Username, Email, password, Nickname string
	following                           []*User
}

func NewUser(username string, email string, password string, nickname string) *User {
	u := User{username, email, password, nickname, make([]*User, 0)}
	return &u
}

func (user *User) CheckPassword(password string) bool {
	return user.password == password
}

func (user *User) IsUser(identification string) bool {
	return user.Username == identification || user.Email == identification || user.Nickname == identification
}

func (user *User) isFollowingUser(follower *User) bool {
	for _, v := range user.following {
		if v.Username == follower.Username {
			return true
		}
	}
	return false
}

func (user *User) FollowUser(userToFollow *User) bool {
	if user.isFollowingUser(userToFollow) {
		return false
	}
	user.following = append(user.following, userToFollow)
	return true
}

func (user *User) GetFollowers() []*User {
	return user.following
}
