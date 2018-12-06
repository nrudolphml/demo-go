package domain

import (
	"fmt"
	"github.com/nrudolph/twitter/src/domain/user"
	"time"
)

type TextTweet struct {
	User *user.User
	Text string
	Date *time.Time
	Id   int
}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) GetUser() *user.User {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *TextTweet) String() string {
	return "@" + tweet.User.Nickname + ": " + tweet.Text
}

func (tweet *TextTweet) FullString() string {
	return fmt.Sprintf("@%s: %s, %d-%02d-%02d %02d:%02d, (id: %d)", tweet.User.Nickname, tweet.Text, tweet.Date.Year(),
		tweet.Date.Month(), tweet.Date.Day(), tweet.Date.Hour(), tweet.Date.Minute(), tweet.Id)
}

func NewTextTweet(user *user.User, text string) *TextTweet {
	t := time.Now()
	v := TextTweet{User: user, Text: text, Date: &t}
	return &v
}
