package domain

import (
	"time"
)

type Tweet struct {
	User *User
	Text string
	Date *time.Time
	Id   int
}

func NewTweet(user *User, text string) *Tweet {
	t := time.Now()
	v := Tweet{User: user, Text: text, Date: &t}
	return &v
}
