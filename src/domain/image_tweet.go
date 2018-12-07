package domain

import (
	"fmt"
	"github.com/nrudolph/twitter/src/domain/user"
	"time"
)

type ImageTweet struct {
	TextTweet
	Url string
}

func NewImageTweet(user *user.User, text string, url string) *ImageTweet {
	t := time.Now()
	return &ImageTweet{TextTweet: TextTweet{User: user, Text: text, Date: &t}, Url: url}
}

func (tweet *ImageTweet) String() string {
	return fmt.Sprintf("@%s: %s, %s", tweet.User.Nickname, tweet.Text, tweet.Url)
}
