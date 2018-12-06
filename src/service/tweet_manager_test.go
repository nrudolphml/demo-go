package service_test

import (
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/domain/user"
	"github.com/nrudolph/twitter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	u, _ := service.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe")
	text := "mi super tweet"
	tweet = domain.NewTweet(u, text)

	// Operation
	_, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweets()[0]

	if publishedTweet.User.Username != u.Username && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbout is %s: %s", u.Username, text, publishedTweet.User.Username, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expeted date can't be nil")
	}
}

func TestWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet *domain.Tweet

	var u *user.User
	text := "super tweet"

	tweet = domain.NewTweet(u, text)

	// Operation

	_, err := service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet *domain.Tweet

	u, _ := service.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe")
	var text string

	tweet = domain.NewTweet(u, text)

	// Operation

	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet *domain.Tweet

	u, _ := service.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe")
	text := "super tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweettweetsuper tweet"

	tweet = domain.NewTweet(u, text)

	// Operation

	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "tweet over 140 characters" {
		t.Error("Expected error is user is required")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Init
	service.InitializeService()
	var tweet, secondTweet *domain.Tweet

	tweetUser, _ := service.AddUser("pepe", "pepe", "pepe", "pepe")
	tweet = domain.NewTweet(tweetUser, "Mi primer tweet")
	secondTweet = domain.NewTweet(tweetUser, "Mi segundo tweet")

	// Operations
	_, _ = service.PublishTweet(tweet)
	_, _ = service.PublishTweet(secondTweet)

	// Validation

	publishedTweets := service.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishTweet := publishedTweets[0]
	secondPublishTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishTweet, tweetUser.Username, tweet.Text) {
		return
	}

	if !isValidTweet(t, secondPublishTweet, tweetUser.Username, secondTweet.Text) {
		return
	}
}

func isValidTweet(t *testing.T, tweet *domain.Tweet, user string, text string) bool {
	if tweet.User.Username != user {
		t.Errorf("Expected %s but was %s", user, tweet.User.Username)
		return false
	}

	if tweet.Text != text {
		t.Errorf("Expected %s but was %s", text, tweet.Text)
		return false
	}

	return true
}

func TestCanRetrieveTweetById(t *testing.T) {
	// Init
	service.InitializeService()
	var tweet *domain.Tweet
	var id int

	tweetUser, _ := service.AddUser("pepe", "pepe@pepe.com", "Pepe", "cont")
	text := "Super tweet"

	tweet = domain.NewTweet(tweetUser, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet, _ := service.GetTweetById(id)
	isValidTweet(t, publishedTweet, tweetUser.Username, text)
}

func TestReturnsErrorWhenRetrieveTweetByIdInvalid(t *testing.T) {
	// Init
	service.InitializeService()
	var tweet *domain.Tweet

	tweetUser, _ := service.AddUser("pepe", "pepe@pepe.com", "Pepe", "cont")
	text := "Super tweet"

	tweet = domain.NewTweet(tweetUser, text)

	// Operation
	_, _ = service.PublishTweet(tweet)

	// Validation
	_, err := service.GetTweetById(2)
	if err == nil {
		t.Error("Expected an error but was not found")
		return
	}
}
