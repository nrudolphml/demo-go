package domain

import "time"

type Tweet struct {
	User *User
	Text string
	Date *time.Time
}

func NewTweet(user *User, text string) *Tweet {
	t := time.Now()
	v := Tweet{user, text, &t}
	return &v
}
