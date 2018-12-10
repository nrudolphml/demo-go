package domain

import (
	"fmt"
	"github.com/nrudolph/twitter/src/domain/user"
	"time"
)

type QuoteTweet struct {
	TextTweet
	Quote Tweet `json:"quote"`
}

func NewQuoteTweet(u *user.User, text string, quote Tweet) *QuoteTweet {
	t := time.Now()
	return &QuoteTweet{TextTweet{User: u, Text: text, Date: &t}, quote}
}

func (tweet *QuoteTweet) String() string {
	return fmt.Sprintf("@%s: %s '@%s: %s'", tweet.User.Nickname, tweet.Text, tweet.Quote.GetUser().Nickname,
		tweet.Quote.GetText())
}
