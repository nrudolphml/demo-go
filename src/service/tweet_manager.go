package service

import (
	"errors"
	"github.com/nrudolph/twitter/src/domain"
)

var tweet *domain.Tweet

func PublishTweet(tweetToPublish *domain.Tweet) error {
	if tweetToPublish.User == nil {
		return errors.New("user is required")
	}
	if tweetToPublish.Text == "" {
		return errors.New("text is required")
	}
	if len(tweetToPublish.Text) > 140 {
		return errors.New("tweet over 140 characters")
	}
	tweet = tweetToPublish
	return nil
}

func GetTweet() *domain.Tweet {
	return tweet
}
