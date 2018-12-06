package service

import (
	"errors"
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/domain/user"
)

type TweetManager struct {
	tweets        []*domain.Tweet
	tweetsByOwner map[*user.User][]*domain.Tweet
}

func (tweetManager *TweetManager) PublishTweet(tweetToPublish *domain.Tweet) (int, error) {
	if tweetToPublish.User == nil {
		return -1, errors.New("user is required")
	}
	if tweetToPublish.Text == "" {
		return -1, errors.New("text is required")
	}
	if len(tweetToPublish.Text) > 140 {
		return -1, errors.New("tweet over 140 characters")
	}
	tweetToPublish.Id = len(tweetManager.tweets)
	tweetManager.tweets = append(tweetManager.tweets, tweetToPublish)

	listOfTweets, exists := tweetManager.tweetsByOwner[tweetToPublish.User]
	if exists {
		tweetManager.tweetsByOwner[tweetToPublish.User] = append(listOfTweets, tweetToPublish)
	} else {
		tweetManager.tweetsByOwner[tweetToPublish.User] = []*domain.Tweet{tweetToPublish}
	}

	return tweetToPublish.Id, nil
}

func (tweetManager *TweetManager) GetTweets() []*domain.Tweet {
	return tweetManager.tweets
}

func (tweetManager *TweetManager) GetTweetById(id int) (*domain.Tweet, error) {
	for _, v := range tweetManager.tweets {
		if v.Id == id {
			return v, nil
		}
	}
	return nil, errors.New("no tweet found with id")
}

func (tweetManager *TweetManager) CountTweetsByUser(owner *user.User) int {
	tweets, exists := tweetManager.tweetsByOwner[owner]
	if !exists {
		return 0
	}
	return len(tweets)
}

func (tweetManager *TweetManager) GetTweetsByUser(owner *user.User) []*domain.Tweet {
	tweets, exists := tweetManager.tweetsByOwner[owner]
	if !exists {
		return nil
	}
	return tweets
}

func (tweetManager *TweetManager) DeleteTweet(owner *user.User, id int) (bool, error) {
	var index = -1
	var tweet *domain.Tweet

	tweetList := tweetManager.tweetsByOwner[owner]
	for i, v := range tweetList {
		if v.Id == id {
			index = i
			tweet = v
			break
		}
	}

	if index == -1 {
		return false, errors.New("the tweet doesn't exist")
	}

	if tweet.User != owner {
		return false, errors.New("the tweet doesn't belong to the user")
	}
	tweetManager.tweetsByOwner[owner] = append(tweetList[:index], tweetList[index+1:]...)

	index = -1
	for i, v := range tweetManager.tweets {
		if v.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return false, errors.New("the tweet doesn't exist")
	}
	tweetManager.tweets = append(tweetManager.tweets[:index], tweetManager.tweets[index+1:]...)
	return true, nil

}

func NewTweetManager() *TweetManager {
	tweetManager := TweetManager{make([]*domain.Tweet, 0), make(map[*user.User][]*domain.Tweet)}
	return &tweetManager
}
