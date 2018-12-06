package service

import (
	"errors"
	"github.com/nrudolph/twitter/src/domain"
)

var tweets []*domain.Tweet

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
	tweets = append(tweets, tweetToPublish)
	return nil
}

func GetTweets() []*domain.Tweet {
	return tweets
}

func InitializeService() {
	tweets = make([]*domain.Tweet, 0)
	users = make([]*domain.User, 0)
}
