package service

import (
	"errors"
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/domain/user"
)

var tweets []*domain.Tweet
var tweetsByOwner map[*user.User][]*domain.Tweet

func InitializeService() {
	tweets = make([]*domain.Tweet, 0)
	users = make([]*user.User, 0)
	loggedUsers = make([]*user.User, 0)
	tweetsByOwner = make(map[*user.User][]*domain.Tweet)
}

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

	listOfTweets, exists := tweetsByOwner[tweetToPublish.User]
	if exists {
		tweetsByOwner[tweetToPublish.User] = append(listOfTweets, tweetToPublish)
	} else {
		tweetsByOwner[tweetToPublish.User] = []*domain.Tweet{tweetToPublish}
	}

	return tweetToPublish.Id, nil
}

func GetTweets() []*domain.Tweet {
	return tweets
}

func GetTweetById(id int) (*domain.Tweet, error) {
	for _, v := range tweets {
		if v.Id == id {
			return v, nil
		}
	}
	return nil, errors.New("no tweet found with id")
}

func CountTweetsByUser(owner *user.User) int {
	tweets, exists := tweetsByOwner[owner]
	if !exists {
		return 0
	}
	return len(tweets)
}

func GetTweetsByUser(owner *user.User) []*domain.Tweet {
	tweets, exists := tweetsByOwner[owner]
	if !exists {
		return nil
	}
	return tweets
}
