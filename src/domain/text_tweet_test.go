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
