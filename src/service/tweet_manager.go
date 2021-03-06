package service

import (
	"errors"
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/domain/user"
	"github.com/nrudolph/twitter/src/persistency"
	"strings"
)

type TweetManager struct {
	tweets        []domain.Tweet
	tweetsByOwner map[*user.User][]domain.Tweet
	id            int
	writer        persistency.TweeterWriter
}

func (tweetManager *TweetManager) PublishTweet(tweetToPublish domain.Tweet) (int, error) {
	if tweetToPublish.GetUser() == nil {
		return -1, errors.New("user is required")
	}
	if tweetToPublish.GetText() == "" {
		return -1, errors.New("text is required")
	}
	if len(tweetToPublish.GetText()) > 140 {
		return -1, errors.New("tweet over 140 characters")
	}
	tweetToPublish.SetId(tweetManager.id)
	tweetManager.id++
	tweetManager.tweets = append(tweetManager.tweets, tweetToPublish)

	listOfTweets, exists := tweetManager.tweetsByOwner[tweetToPublish.GetUser()]
	if exists {
		tweetManager.tweetsByOwner[tweetToPublish.GetUser()] = append(listOfTweets, tweetToPublish)
	} else {
		tweetManager.tweetsByOwner[tweetToPublish.GetUser()] = []domain.Tweet{tweetToPublish}
	}
	tweetManager.writer.WriteTweet(tweetToPublish)

	return tweetToPublish.GetId(), nil
}

func (tweetManager *TweetManager) GetTweets() []domain.Tweet {
	return tweetManager.tweets
}

func (tweetManager *TweetManager) GetTweetById(id int) (domain.Tweet, error) {
	for _, v := range tweetManager.tweets {
		if v.GetId() == id {
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

func (tweetManager *TweetManager) GetTweetsByUser(owner *user.User) []domain.Tweet {
	tweets, exists := tweetManager.tweetsByOwner[owner]
	if !exists {
		return nil
	}
	return tweets
}

func (tweetManager *TweetManager) DeleteTweet(owner *user.User, id int) (bool, error) {
	var index = -1
	var tweet domain.Tweet

	tweetList := tweetManager.tweetsByOwner[owner]
	for i, v := range tweetList {
		if v.GetId() == id {
			index = i
			tweet = v
			break
		}
	}

	if index == -1 {
		return false, errors.New("the tweet doesn't exist")
	}

	if tweet.GetUser() != owner {
		return false, errors.New("the tweet doesn't belong to the user")
	}
	tweetManager.tweetsByOwner[owner] = append(tweetList[:index], tweetList[index+1:]...)

	index = -1
	for i, v := range tweetManager.tweets {
		if v.GetId() == id {
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

func NewTweetManager(writer persistency.TweeterWriter) *TweetManager {
	tweetManager := TweetManager{make([]domain.Tweet, 0), make(map[*user.User][]domain.Tweet), 0, writer}
	return &tweetManager
}

func (tweetManager *TweetManager) SearchTweetsContaining(query string, searchResult chan domain.Tweet) {
	go func() {
		for _, v := range tweetManager.tweets {
			if strings.Contains(v.GetText(), query) {
				searchResult <- v
			}
		}
		searchResult <- nil
		return
	}()
}
