package domain

import (
	"fmt"
	"github.com/nrudolph/twitter/src/domain/user"
	"time"
)

type Tweet struct {
	User *user.User
	Text string
	Date *time.Time
	Id   int
}

func (tweet *Tweet) String() string {
	return "@" + tweet.User.Nickname + ": " + tweet.Text
}

func (tweet *Tweet) PrintableFullTweet() string {
	return fmt.Sprintf("@%s: %s, %d-%02d-%02d %02d:%02d, (id: %d)", tweet.User.Nickname, tweet.Text, tweet.Date.Year(),
		tweet.Date.Month(), tweet.Date.Day(), tweet.Date.Hour(), tweet.Date.Minute(), tweet.Id)
}

func NewTweet(user *user.User, text string) *Tweet {
	t := time.Now()
	v := Tweet{User: user, Text: text, Date: &t}
	return &v
}
