package domain

import (
	"github.com/nrudolph/twitter/src/domain/user"
	"time"
)

type Tweet struct {
	User *user.User
	Text string
	Date *time.Time
	Id   int
}

func (tweet *Tweet) PrintableTweet() string {
	return "@" + tweet.User.Nickname + ": " + tweet.Text
}

func NewTweet(user *user.User, text string) *Tweet {
	t := time.Now()
	v := Tweet{User: user, Text: text, Date: &t}
	return &v
}
