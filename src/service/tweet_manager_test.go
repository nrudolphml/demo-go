package service_test

import (
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user, _ := service.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe")
	text := "mi super tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	_ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()

	if publishedTweet.User.Username != user.Username && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbout is %s: %s", user.Username, text, publishedTweet.User.Username, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expeted date can't be nil")
	}
}

func TestWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	var user *domain.User
	text := "super tweet"

	tweet = domain.NewTweet(user, text)

	// Operation

	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	user, _ := service.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe")
	var text string

	tweet = domain.NewTweet(user, text)

	// Operation

	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	user, _ := service.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe")
	text := "super tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweettweetsuper tweet"

	tweet = domain.NewTweet(user, text)

	// Operation

	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "tweet over 140 characters" {
		t.Error("Expected error is user is required")
	}
}
