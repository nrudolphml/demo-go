package domain_test

import (
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/domain/user"
	"testing"
)

func TestTextTweetCanGetAStringTweet(t *testing.T) {
	// init

	u := user.NewUser("pepe", "pepe@pepe.com", "ppp", "Pepe")
	tweet := domain.NewTextTweet(u, "this is my tweet")

	text := tweet.String()

	expectedText := "@Pepe: this is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}

func TestImageTweetPrintsUserTextAndImageUrl(t *testing.T) {
	// init

	u := user.NewUser("p", "p", "p", "p")
	tweet := domain.NewImageTweet(u, "first tweet", "url.png")

	//operation

	text := tweet.String()

	// Validation

	expectedText := "@p: first tweet, url.png"

	if text != expectedText {
		t.Errorf("Was expected %s but is %s", expectedText, text)
	}
}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	// Initialization
	u := user.NewUser("p", "p", "p", "p")
	quotedTweet := domain.NewTextTweet(u, "This is my tweet")

	u2 := user.NewUser("n", "n", "p", "n")
	tweet := domain.NewQuoteTweet(u2, "Awesome", quotedTweet)
	// Validation
	expectedText := `@n: Awesome '@p: This is my tweet'`
	if tweet.String() != expectedText {
		t.Errorf("Was expected %s but is %s", expectedText, tweet.String())
	}
}
