package service

import (
	"errors"
	"github.com/nrudolph/twitter/src/domain"
)

var tweets []*domain.Tweet

func PublishTweet(tweetToPublish *domain.Tweet) (int, error) {
	if tweetToPublish.User == nil {
		return -1, errors.New("user is required")
	}
	if tweetToPublish.Text == "" {
		return -1, errors.New("text is required")
	}
	if len(tweetToPublish.Text) > 140 {
		return -1, errors.New("tweet over 140 characters")
	}
	tweetToPublish.Id = len(tweets)
	tweets = append(tweets, tweetToPublish)
	return tweetToPublish.Id, nil
}

func GetTweets() []*domain.Tweet {
	return tweets
}

func InitializeService() {
	tweets = make([]*domain.Tweet, 0)
	users = make([]*domain.User, 0)
}

func GetTweetById(id int) (*domain.Tweet, error) {
	for _, v := range tweets {
		if v.Id == id {
			return v, nil
		}
	}
	return nil, errors.New("no tweet found with id")
}
