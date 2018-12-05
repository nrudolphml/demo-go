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
