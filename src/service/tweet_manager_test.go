package service_test

import (
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "mi super tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()

	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbout is %s: %s", user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expeted date can't be nil")
	}
}

func TestWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	var user string
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

	user := "nico"
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

	user := "nico"
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
