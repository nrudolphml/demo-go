package service_test

import (
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/domain/user"
	"github.com/nrudolph/twitter/src/persistency"
	"github.com/nrudolph/twitter/src/service"
	"testing"
)

func BenchmarkPublishTweetWithFileTweetWriter(b *testing.B) {
	writer := persistency.NewFileTweetWritter()
	tweetManager := service.NewTweetManager(writer)
	u := user.NewUser("p", "p", "p", "p")

	tweet := domain.NewTextTweet(u, "Super tweet")

	for n := 0; n < b.N; n++ {
		tweetManager.PublishTweet(tweet)
	}
}

func BenchmarkPublishTweetWithMemTweetWriter(b *testing.B) {
	writer := persistency.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(writer)
	u := user.NewUser("p", "p", "p", "p")

	tweet := domain.NewTextTweet(u, "Super tweet")

	for n := 0; n < b.N; n++ {
		tweetManager.PublishTweet(tweet)
	}
}
