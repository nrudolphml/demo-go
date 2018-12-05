package service_test

import (
	"github.com/nrudolph/twitter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	var tweet = "this is my first tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}
